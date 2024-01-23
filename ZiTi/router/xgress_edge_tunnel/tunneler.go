/*
	Copyright NetFoundry Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package xgress_edge_tunnel

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/ziti/router/fabric"
	"github.com/openziti/ziti/router/xgress"
	"github.com/openziti/ziti/tunnel/dns"
	"github.com/openziti/ziti/tunnel/intercept"
	"github.com/openziti/ziti/tunnel/intercept/host"
	"github.com/openziti/ziti/tunnel/intercept/proxy"
	"github.com/openziti/ziti/tunnel/intercept/tproxy"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/pkg/errors"
	"math"
	"net"
	"strings"
	"time"
)

type tunneler struct {
	dialOptions   *Options
	listenOptions *Options
	stateManager  fabric.StateManager
	bindHandler   xgress.BindHandler

	interceptor     intercept.Interceptor
	servicePoller   *servicePoller
	fabricProvider  *fabricProvider
	terminators     cmap.ConcurrentMap[string, *tunnelTerminator]
	notifyReconnect chan struct{}
}

func newTunneler(factory *Factory, stateManager fabric.StateManager) *tunneler {
	result := &tunneler{
		stateManager:    stateManager,
		terminators:     cmap.New[*tunnelTerminator](),
		notifyReconnect: make(chan struct{}, 1),
	}

	result.fabricProvider = newProvider(factory, result)
	result.servicePoller = newServicePoller(result.fabricProvider)

	go result.ReestablishmentRunner()

	return result
}

func (self *tunneler) Start(notifyClose <-chan struct{}) error {
	self.servicePoller.serviceListenerLock.Lock()
	defer self.servicePoller.serviceListenerLock.Unlock()

	var err error

	log := pfxlog.Logger()
	log.WithField("mode", self.listenOptions.mode).Info("creating interceptor")

	if strings.HasPrefix(self.listenOptions.mode, "tproxy") {
		tproxyConfig := tproxy.Config{
			LanIf:            self.listenOptions.lanIf,
			UDPIdleTimeout:   self.listenOptions.udpIdleTimeout,
			UDPCheckInterval: self.listenOptions.udpCheckInterval,
		}

		if strings.HasPrefix(self.listenOptions.mode, "tproxy:") {
			tproxyConfig.Diverter = strings.TrimPrefix(self.listenOptions.mode, "tproxy:")
		}

		if self.interceptor, err = tproxy.New(tproxyConfig); err != nil {
			return errors.Wrap(err, "failed to initialize tproxy interceptor")
		}
	} else if self.listenOptions.mode == "host" {
		self.listenOptions.resolver = ""
		self.interceptor = host.New()
	} else if self.listenOptions.mode == "proxy" {
		self.listenOptions.resolver = ""
		if self.interceptor, err = proxy.New(net.IPv4zero, self.listenOptions.services); err != nil {
			return errors.Wrap(err, "failed to initialize tproxy interceptor")
		}
	} else {
		return errors.Errorf("unsupported tunnel mode '%v'", self.listenOptions.mode)
	}

	resolver, err := dns.NewResolver(self.listenOptions.resolver)
	if err != nil {
		pfxlog.Logger().WithError(err).Error("failed to start DNS resolver")
	}

	if err = intercept.SetDnsInterceptIpRange(self.listenOptions.dnsSvcIpRange); err != nil {
		pfxlog.Logger().Errorf("invalid dns service IP range %s: %v", self.listenOptions.dnsSvcIpRange, err)
		return err
	}

	self.servicePoller.serviceListener = intercept.NewServiceListener(self.interceptor, resolver)
	self.servicePoller.serviceListener.HandleProviderReady(self.fabricProvider)

	go self.servicePoller.pollServices(self.listenOptions.svcPollRate, notifyClose)
	go self.removeStaleConnections(notifyClose)

	return nil
}

func (self *tunneler) removeStaleConnections(notifyClose <-chan struct{}) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			var toRemove []string
			self.terminators.IterCb(func(key string, t *tunnelTerminator) {

				if t.closed.Load() {
					toRemove = append(toRemove, key)
				}

			})

			for _, key := range toRemove {
				if t, found := self.terminators.Get(key); found {
					self.terminators.Remove(key)
					pfxlog.Logger().Debugf("removed closed tunnel terminator %v for service %v", key, t.context.ServiceName())
				}
			}
		case <-notifyClose:
			return
		}
	}
}

func (self *tunneler) Listen(_ string, bindHandler xgress.BindHandler) error {
	self.bindHandler = bindHandler
	return nil
}

func (self *tunneler) Close() error {
	self.interceptor.Stop()
	return nil
}

func (self *tunneler) HandleReconnect() {
	terminators := self.terminators.Items()
	for _, terminator := range terminators {
		terminator.created.Store(false)
	}

	select {
	case self.notifyReconnect <- struct{}{}:
	default:
	}
}

func (self *tunneler) ReestablishmentRunner() {
	for {
		<-self.notifyReconnect
		self.ReestablishTerminators()
	}
}

func (self *tunneler) ReestablishTerminators() {
	log := pfxlog.Logger()
	terminators := self.terminators.Items()

	time.Sleep(10 * time.Second) // wait for validate terminator messages to come in first

	if len(terminators) > 0 {
		pfxlog.Logger().Debugf("reestablishing %v terminators", len(terminators))
	}

	count := 1
	for _, terminator := range terminators {
		t := terminator
		if !terminator.created.Load() && !terminator.closed.Load() {
			err := self.fabricProvider.factory.env.GetRateLimiterPool().QueueWithTimeout(func() {
				self.fabricProvider.establishTerminatorWithRetry(t)
			}, math.MaxInt64)

			if err != nil {
				// should only happen if router is shutting down
				log.WithField("terminatorId", terminator.id).WithError(err).
					Error("unable to queue terminator reestablishment")
			}
		} else if terminator.created.Load() {
			log.Infof("terminator %v already verified", terminator.id)
		} else if terminator.closed.Load() {
			log.Infof("terminator %v closed, can't reestablish", terminator.id)
		}
		count++
	}

	if len(terminators) > 0 {
		log.Debug("finished queueing terminator reestablishment")
	}
}
