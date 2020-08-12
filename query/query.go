package query

import (
	"encoding/json"

	"github.com/spiegel-im-spiegel/errs"
	paapi5 "github.com/spiegel-im-spiegel/pa-api"
)

//Query is a query data class for PA-API v5
type Query struct {
	OpeCode paapi5.Operation `json:"Operation"`
	request
	Resources       []string `json:",omitempty"`
	enableResources map[resource]bool
}

var _ paapi5.Query = (*Query)(nil) //Query is compatible with paapi5.Query interface

//New creates a new Query instance
func New(opeCode paapi5.Operation) *Query {
	return &Query{OpeCode: opeCode, enableResources: map[resource]bool{}}
}

//Operation returns the type of the PA-API operation
func (q *Query) Operation() paapi5.Operation {
	if q == nil {
		return paapi5.NullOperation
	}
	return q.OpeCode
}

//Payload defines the resources to be returned
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
	b, err := json.Marshal(q)
	return b, errs.Wrap(err)
}

//Stringer interface
func (q *Query) String() string {
	b, err := q.Payload()
	if err != nil {
		return ""
	}
	return string(b)
}

//With returns this instance
func (q *Query) With() *Query {
	if q == nil {
		q = New(paapi5.NullOperation)
	}
	return q
}

//RequestMap is mapping data for RequestFilter
type RequestMap map[RequestFilter]interface{}

//RequestFilters adds RequestFilter to Query instance
func (q *Query) RequestFilters(requests ...RequestMap) *Query {
	for _, request := range requests {
		for name, value := range request {
			q.mapFilter(name, value)
		}
	}
	return q
}

//BrowseNodeInfo sets the resource of BrowseNodeInfo
func (q *Query) BrowseNodeInfo() *Query {
	q.enableResources[resourceBrowseNodeInfo] = true
	return q
}

//Images sets the resource of Images
func (q *Query) Images() *Query {
	q.enableResources[resourceImages] = true
	return q
}

//ItemInfo sets the resource of ItemInfo
func (q *Query) ItemInfo() *Query {
	q.enableResources[resourceItemInfo] = true
	return q
}

//Offers sets the resource of Offers
func (q *Query) Offers() *Query {
	q.enableResources[resourceOffers] = true
	return q
}

//SearchRefinements sets the resource of SearchRefinements
func (q *Query) SearchRefinements() *Query {
	q.enableResources[resourceSearchRefinements] = true
	return q
}

//ParentASIN sets the resource of ParentASIN
func (q *Query) ParentASIN() *Query {
	q.enableResources[resourceParentASIN] = true
	return q
}

//CustomerReviews sets the resource of CustomerReviews resource
func (q *Query) CustomerReviews() *Query {
	q.enableResources[resourceCustomerReviews] = true
	return q
}

/* Copyright 2019,2020 Spiegel and contributors
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
