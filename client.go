package paapi5

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/fetch"
)

const (
	defaultPartnerType = "Associates"
)

//Query interface for Client type
type Query interface {
	Operation() Operation
	Payload() ([]byte, error)
}

//Client interface
type Client interface {
	Marketplace() string
	PartnerTag() string
	PartnerType() string
	Request(Query) ([]byte, error)
	RequestContext(context.Context, Query) ([]byte, error)
}

//client is http.Client for Aozora API Server
type client struct {
	server     *Server
	client     fetch.Client
	partnerTag string
	accessKey  string
	secretKey  string
}

//Marketplace returns name of Marketplace parameter for PA-API v5
func (c *client) Marketplace() string {
	return c.server.Marketplace()
}

//PartnerTag returns PartnerTag parameter for PA-API v5
func (c *client) PartnerTag() string {
	return c.partnerTag
}

//PartnerType returns PartnerType parameter for PA-API v5
func (c *client) PartnerType() string {
	return defaultPartnerType
}

//Request method returns response data (JSON format) by PA-APIv5.
func (c *client) Request(q Query) ([]byte, error) {
	return c.RequestContext(context.Background(), q)
}

//RequestContext method returns response data (JSON format) by PA-APIv5. (with context.Context)
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
	dt := NewTimeStamp(time.Now())
	u := c.server.URL(cmd.Path())
	hds := newHeaders(c.server, cmd, dt)
	sig := c.signiture(c.signedString(hds, payload), hds)
	resp, err := c.client.Post(
		u,
		bytes.NewReader(payload),
		fetch.WithContext(ctx),
		fetch.WithRequestHeaderSet("Accept", c.server.Accept()),
		fetch.WithRequestHeaderSet("Accept-Language", c.server.AcceptLanguage()),
		fetch.WithRequestHeaderSet("Content-Type", c.server.ContentType()),
		fetch.WithRequestHeaderSet("Content-Encoding", hds.get("Content-Encoding")),
		fetch.WithRequestHeaderSet("Host", hds.get("Host")),
		fetch.WithRequestHeaderSet("X-Amz-Date", hds.get("X-Amz-Date")),
		fetch.WithRequestHeaderSet("X-Amz-Target", hds.get("X-Amz-Target")),
		fetch.WithRequestHeaderSet("Authorization", c.authorization(sig, hds)),
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

func (c *client) authorization(sig string, hds *headers) string {
	buf := bytes.Buffer{}
	buf.WriteString(c.server.HMACAlgorithm())
	buf.WriteString(" Credential=")
	buf.WriteString(strings.Join([]string{c.accessKey, hds.dt.StringDate(), c.server.Region(), c.server.ServiceName(), c.server.AWS4Request()}, "/"))
	buf.WriteString(",SignedHeaders=")
	buf.WriteString(hds.list())
	buf.WriteString(",Signature=")
	buf.WriteString(sig)
	return buf.String()
}

func (c *client) signiture(signed string, hds *headers) string {
	dateKey := hmacSHA256([]byte("AWS4"+c.secretKey), []byte(hds.dt.StringDate()))
	regionKey := hmacSHA256(dateKey, []byte(c.server.Region()))
	serviceKey := hmacSHA256(regionKey, []byte(c.server.ServiceName()))
	requestKey := hmacSHA256(serviceKey, []byte(c.server.AWS4Request()))
	return hex.EncodeToString(hmacSHA256(requestKey, []byte(signed)))
}

func (c *client) signedString(hds *headers, payload []byte) string {
	return strings.Join(
		[]string{
			c.server.HMACAlgorithm(),
			hds.dt.String(),
			strings.Join([]string{hds.dt.StringDate(), c.server.Region(), c.server.ServiceName(), c.server.AWS4Request()}, "/"),
			hashedString([]byte(c.canonicalRequest(hds, payload))),
		},
		"\n",
	)
}

func (c *client) canonicalRequest(hds *headers, payload []byte) string {
	request := []string{"POST", hds.cmd.Path(), "", hds.values(), "", hds.list(), hashedString(payload)}
	return strings.Join(request, "\n")
}

func hmacSHA256(key, data []byte) []byte {
	hasher := hmac.New(sha256.New, key)
	_, err := hasher.Write(data)
	if err != nil {
		return []byte{}
	}
	return hasher.Sum(nil)
}

func hashedString(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

type headers struct {
	cmd       Operation
	dt        TimeStamp
	headers   []string
	valueList map[string]string
}

func newHeaders(svr *Server, cmd Operation, dt TimeStamp) *headers {
	hds := &headers{cmd: cmd, dt: dt, headers: []string{"content-encoding", "host", "x-amz-date", "x-amz-target"}, valueList: map[string]string{}}
	hds.valueList["content-encoding"] = svr.ContentEncoding()
	hds.valueList["host"] = svr.HostName()
	hds.valueList["x-amz-date"] = dt.String()
	hds.valueList["x-amz-target"] = cmd.Target()
	return hds
}

func (h *headers) get(name string) string {
	if s, ok := h.valueList[strings.ToLower(name)]; ok {
		return s
	}
	return ""
}

func (h *headers) list() string {
	return strings.Join(h.headers, ";")
}

func (h *headers) values() string {
	list := []string{}
	for _, name := range h.headers {
		list = append(list, name+":"+h.get(name))
	}
	return strings.Join(list, "\n")
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
