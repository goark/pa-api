package query

import (
	"encoding/json"

	"github.com/goark/errs"
	paapi5 "github.com/goark/pa-api"
)

// Query is a query data class for the Amazon Creators API.
//
// The operation type is no longer transmitted in the body (the Creators
// API routes operations by URL path), so OpeCode is excluded from the
// JSON payload. It is retained as a Go field so the client can resolve
// the correct path to POST against.
type Query struct {
	OpeCode paapi5.Operation `json:"-"`
	request
	Resources       []string `json:"resources,omitempty"`
	enableResources map[resource]bool
}

var _ paapi5.Query = (*Query)(nil) //Query is compatible with paapi5.Query interface

// New creates a new Query instance
func New(opeCode paapi5.Operation) *Query {
	return &Query{OpeCode: opeCode, enableResources: map[resource]bool{}}
}

// Operation returns the type of the Creators API operation.
func (q *Query) Operation() paapi5.Operation {
	if q == nil {
		return paapi5.NullOperation
	}
	return q.OpeCode
}

// Payload defines the resources to be returned and renders the request
// body that will be POSTed to the Creators API.
func (q *Query) Payload() ([]byte, error) {
	if q == nil {
		return nil, errs.Wrap(paapi5.ErrNullPointer)
	}
	q.Resources = []string{}
	for r, flag := range q.enableResources {
		if flag {
			q.Resources = append(q.Resources, r.Strings()...)
		}
	}
	if len(q.Resources) == 0 {
		q.Resources = nil
	}
	b, err := json.Marshal(q)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return b, nil
}

// Stringer interface
func (q *Query) String() string {
	b, err := q.Payload()
	if err != nil {
		return ""
	}
	return string(b)
}

// With returns this instance
func (q *Query) With() *Query {
	if q == nil {
		q = New(paapi5.NullOperation)
	}
	return q
}

// RequestMap is mapping data for RequestFilter
type RequestMap map[RequestFilter]interface{}

// RequestFilters adds RequestFilter to Query instance
func (q *Query) RequestFilters(requests ...RequestMap) *Query {
	for _, request := range requests {
		for name, value := range request {
			q.mapFilter(name, value)
		}
	}
	return q
}

// BrowseNodeInfo sets the resource of BrowseNodeInfo
func (q *Query) BrowseNodeInfo() *Query {
	q.enableResources[resourceBrowseNodeInfo] = true
	return q
}

// Images sets the resource of Images
func (q *Query) Images() *Query {
	q.enableResources[resourceImages] = true
	return q
}

// ItemInfo sets the resource of ItemInfo
func (q *Query) ItemInfo() *Query {
	q.enableResources[resourceItemInfo] = true
	return q
}

// Offers selects the V1 Offers resource.
//
// Deprecated: the Creators API does not expose V1 Offers (PA-API V1 Offers
// retired on 2026-01-31). Calling this method now selects OffersV2 to keep
// existing call sites working; migrate to OffersV2 explicitly.
func (q *Query) Offers() *Query {
	q.enableResources[resourceOffersV2] = true
	return q
}

// OffersV2 sets the resource of OffersV2
func (q *Query) OffersV2() *Query {
	q.enableResources[resourceOffersV2] = true
	return q
}

// SearchRefinements sets the resource of SearchRefinements
func (q *Query) SearchRefinements() *Query {
	q.enableResources[resourceSearchRefinements] = true
	return q
}

// ParentASIN sets the resource of ParentASIN
func (q *Query) ParentASIN() *Query {
	q.enableResources[resourceParentASIN] = true
	return q
}

// CustomerReviews sets the resource of CustomerReviews resource
func (q *Query) CustomerReviews() *Query {
	q.enableResources[resourceCustomerReviews] = true
	return q
}

// BrowseNodes sets the resource of BrowseNodes resource
func (q *Query) BrowseNodes() *Query {
	q.enableResources[resourceBrowseNodes] = true
	return q
}

// VariationSummary sets the resource of VariationSummary. The Creators API
// returns the variation summary (page count, total variation count, price
// range and variation dimensions) under the VariationsResult.VariationSummary
// container of the response.
func (q *Query) VariationSummary() *Query {
	q.enableResources[resourceVariationSummary] = true
	return q
}

/* Copyright 2019-2022 Spiegel and contributors
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
