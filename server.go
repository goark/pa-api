package paapi5

import (
	"context"
	"net/http"
	"net/url"
)

const (
	defaultMarketplace     Marketplace = LocaleUnitedStates
	defaultScheme                      = "https"
	defaultAccept                      = "application/json, text/javascript"
	defaultContentType                 = "application/json; charset=UTF-8"
	defaultHMACAlgorithm               = "AWS4-HMAC-SHA256"
	defaultServiceName                 = "ProductAdvertisingAPI"
	defaultContentEncoding             = "amz-1.0"
	defaultAWS4Request                 = "aws4_request"
)

//Server type is a implementation of PA-API service.
type Server struct {
	scheme      string
	marketplace Marketplace
	language    string
}

//ServerOptFunc type is self-referential function type for New functions. (functional options pattern)
type ServerOptFunc func(*Server)

//New function returns an Server instance with options.
func New(opts ...ServerOptFunc) *Server {
	server := &Server{scheme: defaultScheme, marketplace: defaultMarketplace, language: ""}
	for _, opt := range opts {
		opt(server)
	}
	return server
}

//WithMarketplace function returns ServerOptFunc function value.
//This function is used in New functions that represents Marketplace data.
func WithMarketplace(marketplace Marketplace) ServerOptFunc {
	return func(s *Server) {
		if s != nil {
			s.marketplace = marketplace
		}
	}
}

//WithLanguage function returns ServerOptFunc function value.
//This function is used in New functions that represents Accept-Language parameter.
func WithLanguage(language string) ServerOptFunc {
	return func(s *Server) {
		if s != nil {
			s.language = language
		}
	}
}

//URL method returns url of service server information for PA-API v5.
func (s *Server) URL(path string) *url.URL {
	if s == nil {
		s = New()
	}
	return &url.URL{Scheme: s.scheme, Host: s.HostName(), Path: path}
}

//Marketplace method returns marketplace name for PA-API v5.
func (s *Server) Marketplace() string {
	if s == nil {
		s = New()
	}
	return s.marketplace.String()
}

//HostName method returns hostname for PA-API v5.
func (s *Server) HostName() string {
	if s == nil {
		s = New()
	}
	return s.marketplace.HostName()
}

//Region method returns region name for PA-API v5
func (s *Server) Region() string {
	if s == nil {
		s = New()
	}
	return s.marketplace.Region()
}

//Accept method returns Accept parameter for PA-API v5
func (s *Server) Accept() string {
	return defaultAccept
}

//AcceptLanguage method returns Accept-Language parameter for PA-API v5
func (s *Server) AcceptLanguage() string {
	if s == nil {
		s = New()
	}
	if len(s.language) > 0 {
		return s.language
	}
	return s.marketplace.Language() //default language
}

//ContentType method returns Content-Type parameter for PA-API v5
func (s *Server) ContentType() string {
	return defaultContentType
}

//HMACAlgorithm method returns HMAC-Algorithm parameter for PA-API v5
func (s *Server) HMACAlgorithm() string {
	return defaultHMACAlgorithm
}

//ServiceName method returns ServiceName parameter for PA-API v5
func (s *Server) ServiceName() string {
	return defaultServiceName
}

//AWS4Request method returns AWS4Request parameter for PA-API v5
func (s *Server) AWS4Request() string {
	return defaultAWS4Request
}

//ContentEncoding method returns Content-Encoding parameter for PA-API v5
func (s *Server) ContentEncoding() string {
	return defaultContentEncoding
}

//ClientOptFunc type is self-referential function type for Server.CreateClient method. (functional options pattern)
type ClientOptFunc func(*Client)

//CreateClient method returns an Client instance with associate-tag, access-key, secret-key, and other options.
func (s *Server) CreateClient(associateTag, accessKey, secretKey string, opts ...ClientOptFunc) *Client {
	if s == nil {
		s = New()
	}
	cli := &Client{
		server:     s,
		client:     nil,
		ctx:        nil,
		partnerTag: associateTag,
		accessKey:  accessKey,
		secretKey:  secretKey,
	}
	for _, opt := range opts {
		opt(cli)
	}
	if cli.client == nil {
		cli.client = http.DefaultClient
	}
	if cli.ctx == nil {
		cli.ctx = context.Background()
	}
	return cli
}

//WithContext function returns ClientOptFunc function value.
//This function is used in Server.CreateClient method that represents context.Context.
func WithContext(ctx context.Context) ClientOptFunc {
	return func(c *Client) {
		if c != nil {
			c.ctx = ctx
		}
	}
}

//WithHttpClient function returns ClientOptFunc function value.
//This function is used in Server.CreateClient method that represents http.Client.
func WithHttpClient(client *http.Client) ClientOptFunc {
	return func(c *Client) {
		if c != nil {
			c.client = client
		}
	}
}

//DefaultClient function returns an default Client instance with associate-tag, access-key, and secret-key parameters.
func DefaultClient(associateTag, accessKey, secretKey string) *Client {
	return New().CreateClient(associateTag, accessKey, secretKey)
}

/* Copyright 2019 Spiegel and contributors
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
