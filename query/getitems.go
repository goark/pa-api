package query

import (
	paapi5 "github.com/goark/pa-api"
)

//GetItems type is embedded Query for GetItems operation in PA-API v5
type GetItems struct {
	Query
}

var _ paapi5.Query = (*GetItems)(nil) //GetItems is compatible with paapi5.Query interface

//New creates new GetItems instance
func NewGetItems(marketplace, partnerTag, partnerType string) *GetItems {
	q := &GetItems{*(New(paapi5.GetItems))}
	q.Request(Marketplace, marketplace).Request(PartnerTag, partnerTag).Request(PartnerType, partnerType)
	return q
}

var requestsOfGetItems = []RequestFilter{
	Condition,
	CurrencyOfPreference,
	ItemIdType,
	ItemIds,
	LanguagesOfPreference,
	Marketplace,
	Merchant,
	OfferCount,
	PartnerTag,
	PartnerType,
}

//RequestFilters adds RequestFilter to Query instance
func (q *GetItems) Request(request RequestFilter, value interface{}) *GetItems {
	if request.findIn(requestsOfGetItems) {
		q.With().RequestFilters(RequestMap{request: value})
	}
	return q
}

//ASINs sets ItemIds in GetItems instance
func (q *GetItems) ASINs(itms []string) *GetItems {
	return q.Request(ItemIds, itms).Request(ItemIdType, "ASIN")
}

//EnableBrowseNodeInfo sets the resource of BrowseNodeInfo
func (q *GetItems) EnableBrowseNodeInfo() *GetItems {
	q.With().BrowseNodeInfo()
	return q
}

//EnableImages sets the resource of Images
func (q *GetItems) EnableImages() *GetItems {
	q.With().Images()
	return q
}

//EnableItemInfo sets the resource of ItemInfo
func (q *GetItems) EnableItemInfo() *GetItems {
	q.With().ItemInfo()
	return q
}

//EnableOffers sets the resource of Offers
func (q *GetItems) EnableOffers() *GetItems {
	q.With().Offers()
	return q
}

//EnableParentASIN sets the resource of ParentASIN
func (q *GetItems) EnableParentASIN() *GetItems {
	q.With().ParentASIN()
	return q
}

//EnableCustomerReviews sets the resource of CustomerReviews
func (q *GetItems) EnableCustomerReviews() *GetItems {
	q.With().CustomerReviews()
	return q
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
