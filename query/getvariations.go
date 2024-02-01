package query

import (
	paapi5 "github.com/goark/pa-api"
)

// GetVariations type is embedded Query for GetVariations operation in PA-API v5
type GetVariations struct {
	Query
}

var _ paapi5.Query = (*GetVariations)(nil) //GetVariations is compatible with paapi5.Query interface

// New creates new GetVariations instance
func NewGetVariations(marketplace, partnerTag, partnerType string) *GetVariations {
	q := &GetVariations{*(New(paapi5.GetVariations))}
	q.Request(Marketplace, marketplace).Request(PartnerTag, partnerTag).Request(PartnerType, partnerType)
	return q
}

var requestsOfGetVariations = []RequestFilter{
	ASIN,
	Condition,
	CurrencyOfPreference,
	LanguagesOfPreference,
	Marketplace,
	Merchant,
	OfferCount,
	PartnerTag,
	PartnerType,
}

// RequestFilters adds RequestFilter to Query instance
func (q *GetVariations) Request(request RequestFilter, value interface{}) *GetVariations {
	if request.findIn(requestsOfGetVariations) {
		q.With().RequestFilters(RequestMap{request: value})
	}
	return q
}

// ASIN sets ASIN in GetVariations instance
func (q *GetVariations) ASIN(itm string) *GetVariations {
	return q.Request(ASIN, itm)
}

// EnableBrowseNodeInfo sets the resource of BrowseNodeInfo
func (q *GetVariations) EnableBrowseNodeInfo() *GetVariations {
	q.With().BrowseNodeInfo()
	return q
}

// EnableImages sets the resource of Images
func (q *GetVariations) EnableImages() *GetVariations {
	q.With().Images()
	return q
}

// EnableItemInfo sets the resource of ItemInfo
func (q *GetVariations) EnableItemInfo() *GetVariations {
	q.With().ItemInfo()
	return q
}

// EnableOffers sets the resource of Offers
func (q *GetVariations) EnableOffers() *GetVariations {
	q.With().Offers()
	return q
}

// EnableVariationSummary sets the resource of VariationSummary
func (q *GetVariations) EnableVariationSummary() *GetVariations {
	q.With().VariationSummary()
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
