package query

import "strconv"

// RequestFilter signals the types of filters to use
type RequestFilter int

// Constants for RequestFilter
const (
	Actor RequestFilter = iota + 1
	Artist
	ASIN
	Author
	Availability
	Brand
	BrowseNodeID
	Condition
	CurrencyOfPreference
	DeliveryFlags
	ItemIds
	ItemIdType
	ItemCount
	ItemPage
	Keywords
	BrowseNodeIds
	LanguagesOfPreference
	Marketplace // Deprecated: forwarded as the `x-marketplace` request header by the client; not transmitted in the body.
	MaxPrice
	Merchant // Deprecated: removed in the Creators API; values are silently ignored.
	MinPrice
	MinReviewsRating
	MinSavingPercent
	OfferCount // Deprecated: removed in the Creators API; values are silently ignored.
	PartnerTag
	PartnerType // Deprecated: not transmitted by the Creators API.
	Properties
	SearchIndex
	SortBy
	Title
	VariationCount
	VariationPage
)

func (f RequestFilter) findIn(list []RequestFilter) bool {
	for _, elm := range list {
		if f == elm {
			return true
		}
	}
	return false
}

// available (valid) filter parameters
var (
	validationMap = map[RequestFilter][]string{
		Availability:  {"Available", "IncludeOutOfStock"},
		Condition:     {"Any", "New", "Used", "Collectible", "Refurbished"},
		DeliveryFlags: {"AmazonGlobal", "FreeShipping", "FulfilledByAmazon", "Prime"},
		ItemIdType:    {"ASIN"},
		PartnerType:   {"Associates"},
		SearchIndex:   {"All", "AmazonVideo", "Apparel", "Appliances", "ArtsAndCrafts", "Automotive", "Baby", "Beauty", "Books", "Classical", "Collectibles", "Computers", "DigitalMusic", "Electronics", "EverythingElse", "Fashion", "FashionBaby", "FashionBoys", "FashionGirls", "FashionMen", "FashionWomen", "GardenAndOutdoor", "GiftCards", "GroceryAndGourmetFood", "Handmade", "HealthPersonalCare", "HomeAndKitchen", "Industrial", "Jewelry", "KindleStore", "LocalServices", "Luggage", "LuxuryBeauty", "Magazines", "MobileAndAccessories", "MobileApps", "MoviesAndTV", "Music", "MusicalInstruments", "OfficeProducts", "PetSupplies", "Photo", "Shoes", "Software", "SportsAndOutdoors", "ToolsAndHomeImprovement", "ToysAndGames", "VHS", "VideoGames", "Watches"},
		SortBy:        {"AvgCustomerReviews", "Featured", "NewestArrivals", "Price:HighToLow", "Price:LowToHigh", "Relevance"},
	}
)

// isVlidString methos checks if the given parameter is valid for the chosen filter option
func (f RequestFilter) isVlidString(value string) bool {
	switch f {
	case BrowseNodeID, BrowseNodeIds:
		if _, err := strconv.ParseInt(value, 10, 64); err == nil {
			return true
		}
	case Availability, Condition, DeliveryFlags, ItemIdType, PartnerType, SearchIndex, SortBy:
		for _, param := range validationMap[f] {
			if value == param {
				return true
			}
		}
	default:
		if len(value) > 0 {
			return true
		}
	}
	return false
}

// request is the private and anonymously imported struct, which selects the
// filters to be used. Field names use lowerCamelCase JSON tags to match the
// Amazon Creators API request body shape.
type request struct {
	Actor                 string            `json:"actor,omitempty"`
	Artist                string            `json:"artist,omitempty"`
	ASIN                  string            `json:"asin,omitempty"`
	Availability          string            `json:"availability,omitempty"`
	Author                string            `json:"author,omitempty"`
	Brand                 string            `json:"brand,omitempty"`
	BrowseNodeID          string            `json:"browseNodeId,omitempty"`
	Condition             string            `json:"condition,omitempty"`
	CurrencyOfPreference  string            `json:"currencyOfPreference,omitempty"`
	DeliveryFlags         []string          `json:"deliveryFlags,omitempty"`
	ItemIds               []string          `json:"itemIds,omitempty"`
	ItemIdType            string            `json:"itemIdType,omitempty"`
	ItemCount             int               `json:"itemCount,omitempty"`
	ItemPage              int               `json:"itemPage,omitempty"`
	Keywords              string            `json:"keywords,omitempty"`
	BrowseNodeIds         []string          `json:"browseNodeIds,omitempty"`
	LanguagesOfPreference []string          `json:"languagesOfPreference,omitempty"`
	MaxPrice              int               `json:"maxPrice,omitempty"`
	MinPrice              int               `json:"minPrice,omitempty"`
	MinReviewsRating      int               `json:"minReviewsRating,omitempty"`
	MinSavingPercent      int               `json:"minSavingPercent,omitempty"`
	PartnerTag            string            `json:"partnerTag,omitempty"`
	Properties            map[string]string `json:"properties,omitempty"`
	SearchIndex           string            `json:"searchIndex,omitempty"`
	SortBy                string            `json:"sortBy,omitempty"`
	Title                 string            `json:"title,omitempty"`
	VariationCount        int               `json:"variationCount,omitempty"`
	VariationPage         int               `json:"variationPage,omitempty"`
}

// mapFilter is a helper function for (*filters).WithFilters
// This function does not check, if the filters to be used match the chosen searchParam/searchType (Actor, Artist etc.pp.)
func (r *request) mapFilter(filter RequestFilter, filterValue interface{}) {
	switch filter {
	case Actor:
		if param, ok := filterValue.(string); ok && filter.isVlidString(param) {
			r.Actor = param
		}
	case Artist:
		if param, ok := filterValue.(string); ok && filter.isVlidString(param) {
			r.Artist = param
		}
	case ASIN:
		if param, ok := filterValue.(string); ok && filter.isVlidString(param) {
			r.ASIN = param
		}
	case Availability:
		if param, ok := filterValue.(string); ok && filter.isVlidString(param) {
			r.Availability = param
		}
	case Author:
		if param, ok := filterValue.(string); ok && filter.isVlidString(param) {
			r.Author = param
		}
	case Brand:
		if param, ok := filterValue.(string); ok && filter.isVlidString(param) {
			r.Brand = param
		}
	case BrowseNodeID:
		if param, ok := filterValue.(string); ok && filter.isVlidString(param) {
			r.BrowseNodeID = param
		}
	case Condition:
		if param, ok := filterValue.(string); ok && filter.isVlidString(param) {
			r.Condition = param
		}
	case CurrencyOfPreference:
		if param, ok := filterValue.(string); ok && filter.isVlidString(param) {
			r.CurrencyOfPreference = param
		}
	case DeliveryFlags:
		switch v := filterValue.(type) {
		case []string:
			r.DeliveryFlags = []string{}
			for _, param := range v {
				if filter.isVlidString(param) {
					r.DeliveryFlags = append(r.DeliveryFlags, param)
				}
			}
		case string:
			if filter.isVlidString(v) {
				r.DeliveryFlags = []string{v}
			}
		}
	case ItemIds:
		switch v := filterValue.(type) {
		case []string:
			r.ItemIds = []string{}
			for _, param := range v {
				if filter.isVlidString(param) {
					r.ItemIds = append(r.ItemIds, param)
				}
			}
		case string:
			if filter.isVlidString(v) {
				r.ItemIds = []string{v}
			}
		}
	case ItemIdType:
		if param, ok := filterValue.(string); ok && filter.isVlidString(param) {
			r.ItemIdType = param
		}
	case ItemCount:
		if count, ok := filterValue.(int); ok && 0 < count && count < 11 {
			r.ItemCount = count
		}
	case ItemPage:
		if page, ok := filterValue.(int); ok && 0 < page && page < 11 {
			r.ItemPage = page
		}
	case Keywords:
		if param, ok := filterValue.(string); ok && filter.isVlidString(param) {
			r.Keywords = param
		}
	case BrowseNodeIds:
		switch v := filterValue.(type) {
		case []string:
			r.BrowseNodeIds = []string{}
			for _, param := range v {
				if filter.isVlidString(param) {
					r.BrowseNodeIds = append(r.BrowseNodeIds, param)
				}
			}
		case string:
			if filter.isVlidString(v) {
				r.BrowseNodeIds = []string{v}
			}
		}
	case LanguagesOfPreference:
		switch v := filterValue.(type) {
		case []string:
			r.LanguagesOfPreference = []string{}
			for _, param := range v {
				if filter.isVlidString(param) {
					r.LanguagesOfPreference = append(r.LanguagesOfPreference, param)
				}
			}
		case string:
			if filter.isVlidString(v) {
				r.LanguagesOfPreference = []string{v}
			}
		}
	case Marketplace, Merchant, OfferCount, PartnerType:
		// Removed in the Creators API: Marketplace travels in the
		// `x-marketplace` header, PartnerType is implicit, and Merchant /
		// OfferCount are no longer accepted.
	case MaxPrice:
		if price, ok := filterValue.(int); ok && price > 0 {
			r.MaxPrice = price
		}
	case MinPrice:
		if price, ok := filterValue.(int); ok && price > 0 {
			r.MinPrice = price
		}
	case MinReviewsRating:
		if minRating, ok := filterValue.(int); ok && 0 < minRating && minRating < 5 {
			r.MinReviewsRating = minRating
		}
	case MinSavingPercent:
		if minSaving, ok := filterValue.(int); ok && 0 < minSaving && minSaving < 100 {
			r.MinSavingPercent = minSaving
		}
	case PartnerTag:
		if param, ok := filterValue.(string); ok && filter.isVlidString(param) {
			r.PartnerTag = param
		}
	case Properties:
		if params, ok := filterValue.(map[string]string); ok && len(params) > 0 {
			r.Properties = params
		}
	case SearchIndex:
		if param, ok := filterValue.(string); ok && filter.isVlidString(param) {
			r.SearchIndex = param
		}
	case SortBy:
		if param, ok := filterValue.(string); ok && filter.isVlidString(param) {
			r.SortBy = param
		}
	case Title:
		if param, ok := filterValue.(string); ok && filter.isVlidString(param) {
			r.Title = param
		}
	case VariationCount:
		if count, ok := filterValue.(int); ok && 0 < count && count < 11 {
			r.VariationCount = count
		}
	case VariationPage:
		if count, ok := filterValue.(int); ok && 0 < count {
			r.VariationPage = count
		}
	}

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
