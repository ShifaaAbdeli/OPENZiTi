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

package rest_util

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

// NewHttpClientWithTlsConfig provides a default HTTP client with generous default timeouts.
func NewHttpClientWithTlsConfig(tlsClientConfig *tls.Config) (*http.Client, error) {
	httpClientTransport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		}).DialContext,

		ForceAttemptHTTP2:     true,
		MaxIdleConns:          10,
		IdleConnTimeout:       10 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	httpClientTransport.TLSClientConfig = tlsClientConfig

	httpClient := &http.Client{
		Transport: httpClientTransport,
		Timeout:   10 * time.Second,
	}
	return httpClient, nil
}

// NewTlsConfig creates a tls.Config with default min/max TSL versions.
func NewTlsConfig() (*tls.Config, error) {
	return &tls.Config{
		MinVersion: tls.VersionTLS12,
		MaxVersion: tls.VersionTLS13,
	}, nil
}
