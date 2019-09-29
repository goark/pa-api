package paapi5

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultMarketplace     = "www.amazon.co.jp"
	defaultAccept          = "application/json, text/javascript"
	defaultAcceptLanguage  = "en-US"
	defaultContentType     = "application/json; charset=UTF-8"
	defaultHMACAlgorithm   = "AWS4-HMAC-SHA256"
	defaultServiceName     = "ProductAdvertisingAPI"
	defaultRegion          = "us-west-2"
	defaultContentEncoding = "amz-1.0"
	defaultAWS4Request     = "aws4_request"
)

//Server is informations of Aozora API
type Server struct {
	scheme      string
	marketplace string
}

//ServerOptFunc is self-referential function for functional options pattern
type ServerOptFunc func(*Server)

//New returns new Server instance
func New(opts ...ServerOptFunc) *Server {
	server := &Server{scheme: "https", marketplace: defaultMarketplace}
	for _, opt := range opts {
		opt(server)
	}
	return server
}

//WithMarketplace returns function for setting hostname
func WithMarketplace(marketplace string) ServerOptFunc {
	return func(s *Server) {
		if s != nil {
			s.marketplace = marketplace
		}
	}
}

//Marketplace returns name of Marketplace parameter for PA-API v5
func (s *Server) Marketplace() string {
	if s == nil {
		return ""
	}
	return s.marketplace
}

//HostName returns host name for PA-API v5
func (s *Server) HostName() string {
	if s == nil || len(s.marketplace) == 0 {
		return ""
	}
	if strings.HasPrefix(s.marketplace, "www.") {
		return strings.Replace(s.marketplace, "www", "webservices", 1)
	}
	return ""
}

//URL returns url.URL instance
func (s *Server) URL(path string) *url.URL {
	if s == nil {
		s = New()
	}
	return &url.URL{Scheme: s.scheme, Host: s.HostName(), Path: path}
}

//Accept returns Accept parameter for PA-API v5
func (s *Server) Accept() string {
	return defaultAccept
}

//AcceptLanguage returns Accept-Language parameter for PA-API v5
func (s *Server) AcceptLanguage() string {
	return defaultAcceptLanguage
}

//ContentType returns Content-Type parameter for PA-API v5
func (s *Server) ContentType() string {
	return defaultContentType
}

//HMACAlgorithm returns HMAC-Algorithm parameter for PA-API v5
func (s *Server) HMACAlgorithm() string {
	return defaultHMACAlgorithm
}

//ServiceName returns ServiceName parameter for PA-API v5
func (s *Server) ServiceName() string {
	return defaultServiceName
}

//Region returns Region parameter for PA-API v5
func (s *Server) Region() string {
	return defaultRegion
}

//AWS4Request returns AWS4Request parameter for PA-API v5
func (s *Server) AWS4Request() string {
	return defaultAWS4Request
}

//ContentEncoding returns Content-Encoding parameter for PA-API v5
func (s *Server) ContentEncoding() string {
	return defaultContentEncoding
}

//CreateClient returns new Client instance
func (s *Server) CreateClient(ctx context.Context, client *http.Client, associateTag, accessKey, secretKey string) *Client {
	if s == nil {
		s = New()
	}
	if client == nil {
		client = http.DefaultClient
	}
	if ctx == nil {
		ctx = context.Background()
	}
	return &Client{server: s, client: client, ctx: ctx, partnerTag: associateTag, accessKey: accessKey, secretKey: secretKey}
}

//DefaultClient returns new Client instance with default setting
func DefaultClient(associateTag, accessKey, secretKey string) *Client {
	return New().CreateClient(nil, nil, associateTag, accessKey, secretKey)
}

/* Copyright 2019 Spiegel
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* 	http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */
