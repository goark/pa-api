package paapi5

import (
	"bytes"
	"context"
	"net/http"

	"github.com/goark/errs"
	"github.com/goark/fetch"
)

const (
	defaultPartnerType = "Associates"
	// marketplaceHeader is the request header used by the Creators API to
	// select the target Amazon marketplace (e.g. www.amazon.co.jp).
	marketplaceHeader = "x-marketplace"
)

// Query interface for Client type
type Query interface {
	Operation() Operation
	Payload() ([]byte, error)
}

// Client interface
type Client interface {
	Marketplace() string
	PartnerTag() string
	PartnerType() string
	Request(Query) ([]byte, error)
	RequestContext(context.Context, Query) ([]byte, error)
}

// client is the HTTP client used to call the Amazon Creators API.
type client struct {
	server           *Server
	httpClient       fetch.Client
	tokenHTTPClient  *http.Client
	partnerTag       string
	credentialID     string
	credentialSecret string
	version          string
	authEndpoint     string
	auth             *tokenManager
}

// Marketplace returns the marketplace name (e.g. www.amazon.com).
func (c *client) Marketplace() string {
	return c.server.Marketplace()
}

// PartnerTag returns the configured partner (associate) tag.
func (c *client) PartnerTag() string {
	return c.partnerTag
}

// PartnerType returns the partner type. The Creators API does not
// transmit this value over the wire; it is retained for backward
// compatibility with code that inspects the client.
func (c *client) PartnerType() string {
	return defaultPartnerType
}

// Request issues the supplied query against the Creators API using a
// background context.
func (c *client) Request(q Query) ([]byte, error) {
	return c.RequestContext(context.Background(), q)
}

// RequestContext issues the supplied query against the Creators API,
// honouring cancellation on the supplied context.
func (c *client) RequestContext(ctx context.Context, q Query) ([]byte, error) {
	payload, err := q.Payload()
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("Operation", q.Operation().String()))
	}
	b, err := c.post(ctx, q.Operation(), payload)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("Operation", q.Operation().String()), errs.WithContext("payload", string(payload)))
	}
	return b, nil
}

func (c *client) post(ctx context.Context, cmd Operation, payload []byte) ([]byte, error) {
	u := c.server.URL(cmd.Path())
	token, err := c.auth.Token(ctx)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", u.String()))
	}
	resp, err := c.httpClient.PostWithContext(
		ctx,
		u,
		bytes.NewReader(payload),
		fetch.WithRequestHeaderSet("Accept", c.server.Accept()),
		fetch.WithRequestHeaderSet("Content-Type", c.server.ContentType()),
		fetch.WithRequestHeaderSet(marketplaceHeader, c.server.Marketplace()),
		fetch.WithRequestHeaderSet("Authorization", authorizationHeader(token, c.version)),
	)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", u.String()), errs.WithContext("payload", string(payload)))
	}
	body, err := resp.DumpBodyAndClose()
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", u.String()), errs.WithContext("payload", string(payload)))
	}
	return body, nil
}

/* Copyright 2019-2021 Spiegel
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
