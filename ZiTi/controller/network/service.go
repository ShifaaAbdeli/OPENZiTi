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

package network

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/openziti/storage/boltz"
	"github.com/openziti/ziti/common/pb/cmd_pb"
	"github.com/openziti/ziti/controller/change"
	"github.com/openziti/ziti/controller/command"
	"github.com/openziti/ziti/controller/db"
	"github.com/openziti/ziti/controller/fields"
	"github.com/openziti/ziti/controller/models"
	"github.com/orcaman/concurrent-map/v2"
	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
	"google.golang.org/protobuf/proto"
	"reflect"
	"time"
)

type Service struct {
	models.BaseEntity
	Name               string
	TerminatorStrategy string
	Terminators        []*Terminator
	MaxIdleTime        time.Duration
}

func (self *Service) GetName() string {
	return self.Name
}

func (entity *Service) toBolt() *db.Service {
	return &db.Service{
		BaseExtEntity:      *boltz.NewExtEntity(entity.Id, entity.Tags),
		Name:               entity.Name,
		MaxIdleTime:        entity.MaxIdleTime,
		TerminatorStrategy: entity.TerminatorStrategy,
	}
}

func newServiceManager(managers *Managers) *ServiceManager {
	result := &ServiceManager{
		baseEntityManager: newBaseEntityManager[*Service, *db.Service](managers, managers.stores.Service, func() *Service {
			return &Service{}
		}),
		cache: cmap.New[*Service](),
		store: managers.stores.Service,
	}
	result.populateEntity = result.populateService

	managers.stores.Service.AddEntityIdListener(result.RemoveFromCache, boltz.EntityUpdated, boltz.EntityDeleted)

	return result
}

type ServiceManager struct {
	baseEntityManager[*Service, *db.Service]
	cache cmap.ConcurrentMap[string, *Service]
	store db.ServiceStore
}

func (self *ServiceManager) NotifyTerminatorChanged(terminator *db.Terminator) *db.Terminator {
	// patched entities may not have all fields, if service is blank, load terminator
	serviceId := terminator.Service
	if serviceId == "" {
		err := self.db.View(func(tx *bbolt.Tx) error {
			t, _, err := self.stores.Terminator.FindById(tx, terminator.Id)
			if t != nil {
				terminator = t
			}
			return err
		})
		if err != nil {
			self.clearCache()
			return terminator
		}
		serviceId = terminator.Service
	}
	pfxlog.Logger().Debugf("clearing service from cache: %v", serviceId)
	self.RemoveFromCache(serviceId)
	return terminator
}

func (self *ServiceManager) Create(entity *Service, ctx *change.Context) error {
	return DispatchCreate[*Service](self, entity, ctx)
}

func (self *ServiceManager) ApplyCreate(cmd *command.CreateEntityCommand[*Service], ctx boltz.MutateContext) error {
	s := cmd.Entity
	err := self.db.Update(ctx, func(ctx boltz.MutateContext) error {
		if err := self.ValidateNameOnCreate(ctx.Tx(), s); err != nil {
			return err
		}
		if err := self.store.Create(ctx, s.toBolt()); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	// don't cache, wait for first read. entity may not match data store as data store may have set defaults
	return nil
}

func (self *ServiceManager) Update(entity *Service, updatedFields fields.UpdatedFields, ctx *change.Context) error {
	return DispatchUpdate[*Service](self, entity, updatedFields, ctx)
}

func (self *ServiceManager) ApplyUpdate(cmd *command.UpdateEntityCommand[*Service], ctx boltz.MutateContext) error {
	if err := self.updateGeneral(ctx, cmd.Entity, cmd.UpdatedFields); err != nil {
		return err
	}
	self.RemoveFromCache(cmd.Entity.Id)
	return nil
}

func (self *ServiceManager) Read(id string) (entity *Service, err error) {
	err = self.db.View(func(tx *bbolt.Tx) error {
		entity, err = self.readInTx(tx, id)
		return err
	})
	if err != nil {
		return nil, err
	}
	return entity, err
}

func (self *ServiceManager) GetIdForName(id string) (string, error) {
	var result []byte
	err := self.db.View(func(tx *bbolt.Tx) error {
		result = self.store.GetNameIndex().Read(tx, []byte(id))
		return nil
	})
	return string(result), err
}

func (self *ServiceManager) readInTx(tx *bbolt.Tx, id string) (*Service, error) {
	if service, found := self.cache.Get(id); found {
		return service, nil
	}

	entity := &Service{}
	if err := self.readEntityInTx(tx, id, entity); err != nil {
		return nil, err
	}

	self.cacheService(entity)
	return entity, nil
}

func (self *ServiceManager) populateService(entity *Service, tx *bbolt.Tx, boltEntity boltz.Entity) error {
	boltService, ok := boltEntity.(*db.Service)
	if !ok {
		return errors.Errorf("unexpected type %v when filling model service", reflect.TypeOf(boltEntity))
	}
	entity.Name = boltService.Name
	entity.MaxIdleTime = boltService.MaxIdleTime
	entity.TerminatorStrategy = boltService.TerminatorStrategy
	entity.FillCommon(boltService)

	terminatorIds := self.store.GetRelatedEntitiesIdList(tx, entity.Id, db.EntityTypeTerminators)
	for _, terminatorId := range terminatorIds {
		if terminator, _ := self.Terminators.readInTx(tx, terminatorId); terminator != nil {
			entity.Terminators = append(entity.Terminators, terminator)
		}
	}

	return nil
}

func (self *ServiceManager) cacheService(service *Service) {
	pfxlog.Logger().Tracef("updated service cache: %v", service.Id)
	self.cache.Set(service.Id, service)
}

func (self *ServiceManager) RemoveFromCache(id string) {
	pfxlog.Logger().Debugf("removed service from cache: %v", id)
	self.cache.Remove(id)
}

func (self *ServiceManager) clearCache() {
	pfxlog.Logger().Debugf("clearing all services from cache")
	for _, key := range self.cache.Keys() {
		self.cache.Remove(key)
	}
}

func (self *ServiceManager) Marshall(entity *Service) ([]byte, error) {
	tags, err := cmd_pb.EncodeTags(entity.Tags)
	if err != nil {
		return nil, err
	}

	msg := &cmd_pb.Service{
		Id:                 entity.Id,
		Name:               entity.Name,
		MaxIdleTime:        int64(entity.MaxIdleTime),
		TerminatorStrategy: entity.TerminatorStrategy,
		Tags:               tags,
	}

	return proto.Marshal(msg)
}

func (self *ServiceManager) Unmarshall(bytes []byte) (*Service, error) {
	msg := &cmd_pb.Service{}
	if err := proto.Unmarshal(bytes, msg); err != nil {
		return nil, err
	}

	return &Service{
		BaseEntity: models.BaseEntity{
			Id:   msg.Id,
			Tags: cmd_pb.DecodeTags(msg.Tags),
		},
		Name:               msg.Name,
		MaxIdleTime:        time.Duration(msg.MaxIdleTime),
		TerminatorStrategy: msg.TerminatorStrategy,
	}, nil
}
