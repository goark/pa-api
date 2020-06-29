package query

import (
	paapi5 "github.com/spiegel-im-spiegel/pa-api"
)

//SearchItems type is embedded Query for SearchItems operation in PA-API v5
type SearchItems struct {
	Query
}

var _ paapi5.Query = (*SearchItems)(nil) //SearchItems is compatible with paapi5.Query interface

//NewSearchItems creates a new SearchItems instance
func NewSearchItems(marketplace, partnerTag, partnerType string) *SearchItems {
	q := &SearchItems{*(New(paapi5.SearchItems))}
	q.Request(Marketplace, marketplace).Request(PartnerTag, partnerTag).Request(PartnerType, partnerType)
	return q
}

var (
	requestsOfSearchItems = []RequestFilter{
		Actor,
		Artist,
		Author,
		Availability,
		Brand,
		BrowseNodeID,
		Condition,
		CurrencyOfPreference,
		DeliveryFlags,
		ItemCount,
		ItemPage,
		Keywords,
		LanguagesOfPreference,
		Marketplace,
		MaxPrice,
		Merchant,
		MinPrice,
		MinReviewsRating,
		MinSavingPercent,
		OfferCount,
		PartnerTag,
		PartnerType,
		Properties,
		SearchIndex,
		SortBy,
		Title,
	}
	searchTypes = []RequestFilter{Actor, Artist, Author, Brand, Keywords, Title}
)

//Request adds RequestFilter to Query instance
func (q *SearchItems) Request(request RequestFilter, value interface{}) *SearchItems {
	if request.findIn(requestsOfSearchItems) {
		q.With().RequestFilters(RequestMap{request: value})
	}
	return q
}

//Search is a generic search query funtion to obtain informations from the "SearchItems"-operation
func (q *SearchItems) Search(searchType RequestFilter, searchParam string) *SearchItems {
	if searchType.findIn(searchTypes) {
		return q.Request(searchType, searchParam)
	}
	return q
}

//EnableBrowseNodeInfo sets the enableBrowseNodeInfo flag in SearchItems instance
func (q *SearchItems) EnableBrowseNodeInfo() *SearchItems {
	q.With().BrowseNodeInfo()
	return q
}

//EnableImages sets the enableImages flag in SearchItems instance
func (q *SearchItems) EnableImages() *SearchItems {
	q.With().Images()
	return q
}

//EnableItemInfo sets the enableItemInfo flag in SearchItems instance
func (q *SearchItems) EnableItemInfo() *SearchItems {
	q.With().ItemInfo()
	return q
}

//EnableOffers sets the enableOffers flag in SearchItems instance
func (q *SearchItems) EnableOffers() *SearchItems {
	q.With().Offers()
	return q
}

//EnableSearchRefinements sets the enableOffers flag in SearchItems instance
func (q *SearchItems) EnableSearchRefinements() *SearchItems {
	q.With().SearchRefinements()
	return q
}

//EnableParentASIN sets the enableParentASIN flag in SearchItems instance
func (q *SearchItems) EnableParentASIN() *SearchItems {
	q.With().ParentASIN()
	return q
}

//EnableCustomerReviews sets the enableCustomerReviews flag in SearchItems instance
func (q *SearchItems) EnableCustomerReviews() *SearchItems {
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
