package query

import (
	paapi5 "github.com/goark/pa-api"
)

//GetItems type is embedded Query for GetItems operation in PA-API v5
type GetBrowseNodes struct {
	Query
}

var _ paapi5.Query = (*GetItems)(nil) //GetItems is compatible with paapi5.Query interface

//New creates new GetBrowseNodes instance
func NewGetBrowseNodes(marketplace, partnerTag, partnerType string) *GetBrowseNodes {
	q := &GetBrowseNodes{*(New(paapi5.GetBrowseNodes))}
	q.Request(Marketplace, marketplace).Request(PartnerTag, partnerTag).Request(PartnerType, partnerType)
	return q
}

var requestsOfGetBrowseNodes = []RequestFilter{
	BrowseNodeIds,
	LanguagesOfPreference,
	Marketplace,
	PartnerTag,
	PartnerType,
}

//RequestFilters adds RequestFilter to Query instance
func (q *GetBrowseNodes) Request(request RequestFilter, value interface{}) *GetBrowseNodes {
	if request.findIn(requestsOfGetBrowseNodes) {
		q.With().RequestFilters(RequestMap{request: value})
	}
	return q
}

//BrowseNodeIds sets ItemIds in GetItems instance
func (q *GetBrowseNodes) BrowseNodeIds(itms []string) *GetBrowseNodes {
	return q.Request(BrowseNodeIds, itms)
}

//EnableBrowseNodes sets the resource of EnableBrowseNodes
func (q *GetBrowseNodes) EnableBrowseNodes() *GetBrowseNodes {
	q.With().BrowseNodes()
	return q
}

/* Copyright 2022 Spiegel and contributors
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
