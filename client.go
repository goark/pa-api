package paapi5

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/spiegel-im-spiegel/errs"
)

const (
	defaultPartnerType = "Associates"
)

//Query interface for Client type
type Query interface {
	Operation() Operation
	Payload() ([]byte, error)
}

//Client is http.Client for Aozora API Server
type Client struct {
	server     *Server
	client     *http.Client
	ctx        context.Context
	partnerTag string
	accessKey  string
	secretKey  string
}

//Marketplace returns name of Marketplace parameter for PA-API v5
func (c *Client) Marketplace() string {
	return c.server.Marketplace()
}

//PartnerTag returns PartnerTag parameter for PA-API v5
func (c *Client) PartnerTag() string {
	return c.partnerTag
}

//PartnerType returns PartnerType parameter for PA-API v5
func (c *Client) PartnerType() string {
	return defaultPartnerType
}

func (c *Client) Request(q Query) ([]byte, error) {
	payload, err := q.Payload()
	if err != nil {
		return nil, errs.Wrap(err, "", errs.WithContext("Operation", q.Operation().String()))
	}
	b, err := c.post(q.Operation(), payload)
	if err != nil {
		return nil, errs.Wrap(err, "", errs.WithContext("Operation", q.Operation().String()), errs.WithContext("payload", string(payload)))
	}
	return b, nil
}

func (c *Client) post(cmd Operation, payload []byte) ([]byte, error) {
	dt := NewTimeStamp(time.Now())
	u := c.server.URL(cmd.Path())
	hds := newHeaders(c.server, cmd, dt)
	sig := c.signiture(c.signedString(hds, payload), hds)
	req, err := http.NewRequestWithContext(c.ctx, "POST", u.String(), bytes.NewReader(payload))
	if err != nil {
		return nil, errs.Wrap(err, "", errs.WithContext("url", u.String()), errs.WithContext("payload", string(payload)))
	}
	req.Header.Add("Accept", c.server.Accept())
	req.Header.Add("Accept-Language", c.server.AcceptLanguage())
	req.Header.Add("Content-Type", c.server.ContentType())
	req.Header.Add("Content-Encoding", hds.get("Content-Encoding"))
	req.Header.Add("Host", hds.get("Host"))
	req.Header.Add("X-Amz-Date", hds.get("X-Amz-Date"))
	req.Header.Add("X-Amz-Target", hds.get("X-Amz-Target"))
	req.Header.Add("Authorization", c.authorization(sig, hds))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errs.Wrap(err, "", errs.WithContext("url", u.String()), errs.WithContext("payload", string(payload)))
	}
	defer resp.Body.Close()

	if !(resp.StatusCode != 0 && resp.StatusCode < http.StatusBadRequest) {
		return nil, errs.Wrap(ErrHTTPStatus, "", errs.WithContext("url", u.String()), errs.WithContext("payload", string(payload)), errs.WithContext("status", resp.Status))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errs.Wrap(err, "", errs.WithContext("url", u.String()), errs.WithContext("payload", string(payload)))
	}
	return body, nil
}

func (c *Client) authorization(sig string, hds *headers) string {
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

func (c *Client) signiture(signed string, hds *headers) string {
	dateKey := hmacSHA256([]byte("AWS4"+c.secretKey), []byte(hds.dt.StringDate()))
	regionKey := hmacSHA256(dateKey, []byte(c.server.Region()))
	serviceKey := hmacSHA256(regionKey, []byte(c.server.ServiceName()))
	requestKey := hmacSHA256(serviceKey, []byte(c.server.AWS4Request()))
	return hex.EncodeToString(hmacSHA256(requestKey, []byte(signed)))
}

func (c *Client) signedString(hds *headers, payload []byte) string {
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

func (c *Client) canonicalRequest(hds *headers, payload []byte) string {
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
