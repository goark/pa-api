package query

import "strconv"

//RequestFilter signals the types of filters to use
type RequestFilter int

//Constants for RequestFilter
const (
	Actor RequestFilter = iota + 1
	Artist
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
	LanguagesOfPreference
	Marketplace
	MaxPrice
	Merchant
	MinPrice
	MinReviewsRating
	MinSavingPercent
	OfferCount
	PartnerTag
	PartnerType
	Properties
	SearchIndex
	SortBy
	Title
)

func (f RequestFilter) findIn(list []RequestFilter) bool {
	for _, elm := range list {
		if f == elm {
			return true
		}
	}
	return false
}

//available (valid) filter parameters
var (
	validationMap = map[RequestFilter][]string{
		Availability:  []string{"Available", "IncludeOutOfStock"},
		Condition:     []string{"Any", "New", "Used", "Collectible", "Refurbished"},
		DeliveryFlags: []string{"AmazonGlobal", "FreeShipping", "FulfilledByAmazon", "Prime"},
		ItemIdType:    []string{"ASIN"},
		Merchant:      []string{"All", "Amazon"},
		PartnerType:   []string{"Associates"},
		SearchIndex:   []string{"All", "AmazonVideo", "Apparel", "Appliances", "ArtsAndCrafts", "Automotive", "Baby", "Beauty", "Books", "Classical", "Collectibles", "Computers", "DigitalMusic", "Electronics", "EverythingElse", "Fashion", "FashionBaby", "FashionBoys", "FashionGirls", "FashionMen", "FashionWomen", "GardenAndOutdoor", "GiftCards", "GroceryAndGourmetFood", "Handmade", "HealthPersonalCare", "HomeAndKitchen", "Industrial", "Jewelry", "KindleStore", "LocalServices", "Luggage", "LuxuryBeauty", "Magazines", "MobileAndAccessories", "MobileApps", "MoviesAndTV", "Music", "MusicalInstruments", "OfficeProducts", "PetSupplies", "Photo", "Shoes", "Software", "SportsAndOutdoors", "ToolsAndHomeImprovement", "ToysAndGames", "VHS", "VideoGames", "Watches"},
		SortBy:        []string{"AvgCustomerReviews", "Featured", "NewestArrivals", "Price:HighToLow", "Price:LowToHigh", "Relevance"},
	}
)

//isVlidString methos checks if the given parameter is valid for the chosen filter option
func (f RequestFilter) isVlidString(value string) bool {
	switch f {
	case BrowseNodeID:
		if _, err := strconv.ParseInt(value, 10, 64); err == nil {
			return true
		}
	default:
		if params, ok := validationMap[f]; ok {
			for _, param := range params {
				if value == param {
					return true
				}
			}
		} else if len(value) > 0 {
			return true
		}
	}
	return false
}

//request is the private and anonymously imported struct, which selects the filters to be used
type request struct {
	Actor                 string            `json:",omitempty"`
	Artist                string            `json:",omitempty"`
	Availability          string            `json:",omitempty"`
	Author                string            `json:",omitempty"`
	Brand                 string            `json:",omitempty"`
	BrowseNodeID          string            `json:"BrowseNodeId,omitempty"`
	Condition             string            `json:",omitempty"`
	CurrencyOfPreference  string            `json:",omitempty"`
	DeliveryFlags         []string          `json:",omitempty"`
	ItemIds               []string          `json:",omitempty"`
	ItemIdType            string            `json:",omitempty"`
	ItemCount             int               `json:",omitempty"`
	ItemPage              int               `json:",omitempty"`
	Keywords              string            `json:",omitempty"`
	LanguagesOfPreference []string          `json:",omitempty"`
	Marketplace           string            `json:",omitempty"`
	MaxPrice              int               `json:",omitempty"`
	Merchant              string            `json:",omitempty"`
	MinPrice              int               `json:",omitempty"`
	MinReviewsRating      int               `json:",omitempty"`
	MinSavingPercent      int               `json:",omitempty"`
	OfferCount            int               `json:",omitempty"`
	PartnerTag            string            `json:",omitempty"`
	PartnerType           string            `json:",omitempty"`
	Properties            map[string]string `json:",omitempty"`
	SearchIndex           string            `json:",omitempty"`
	SortBy                string            `json:",omitempty"`
	Title                 string            `json:",omitempty"`
}

//mapFilter is a helper function for (*filters).WithFilters
// This function does not check, if the filters to be used match the chosen searchParam/searchType (Actor, Artist etc.pp.)
// TODO: 	- reduce nesting
func (r *request) mapFilter(filter RequestFilter, filterValue interface{}) {
	switch filter {
	case Actor:
		if param := filterValue.(string); filter.isVlidString(param) {
			r.Actor = param
		}
	case Artist:
		if param := filterValue.(string); filter.isVlidString(param) {
			r.Artist = param
		}
	case Availability:
		if param := filterValue.(string); filter.isVlidString(param) {
			r.Availability = param
		}
	case Author:
		if param := filterValue.(string); filter.isVlidString(param) {
			r.Author = param
		}
	case Brand:
		if param := filterValue.(string); filter.isVlidString(param) {
			r.Brand = param
		}
	case BrowseNodeID:
		if param := filterValue.(string); filter.isVlidString(param) {
			r.BrowseNodeID = param
		}
	case Condition:
		if param := filterValue.(string); filter.isVlidString(param) {
			r.Condition = param
		}
	case CurrencyOfPreference:
		if param := filterValue.(string); filter.isVlidString(param) {
			r.CurrencyOfPreference = param
		}
	case DeliveryFlags:
		switch filterValue.(type) {
		case []string:
			r.DeliveryFlags = []string{}
			for _, param := range filterValue.([]string) {
				if filter.isVlidString(param) {
					r.DeliveryFlags = append(r.DeliveryFlags, param)
				}
			}
		case string:
			if param := filterValue.(string); filter.isVlidString(param) {
				r.DeliveryFlags = []string{param}
			}
		}
	case ItemIds:
		switch filterValue.(type) {
		case []string:
			r.ItemIds = []string{}
			for _, param := range filterValue.([]string) {
				if filter.isVlidString(param) {
					r.ItemIds = append(r.ItemIds, param)
				}
			}
		case string:
			if param := filterValue.(string); filter.isVlidString(param) {
				r.ItemIds = []string{param}
			}
		}
	case ItemIdType:
		if param := filterValue.(string); filter.isVlidString(param) {
			r.ItemIdType = param
		}
	case ItemCount:
		if count := filterValue.(int); 0 < count && count < 11 {
			r.ItemCount = count
		}
	case ItemPage:
		if page := filterValue.(int); 0 < page && page < 11 {
			r.ItemPage = page
		}
	case Keywords:
		if param := filterValue.(string); filter.isVlidString(param) {
			r.Keywords = param
		}
	case LanguagesOfPreference:
		switch filterValue.(type) {
		case []string:
			r.LanguagesOfPreference = []string{}
			for _, param := range filterValue.([]string) {
				if filter.isVlidString(param) {
					r.LanguagesOfPreference = append(r.LanguagesOfPreference, param)
				}
			}
		case string:
			if param := filterValue.(string); filter.isVlidString(param) {
				r.LanguagesOfPreference = []string{param}
			}
		}
	case Marketplace:
		if param := filterValue.(string); filter.isVlidString(param) {
			r.Marketplace = param
		}
	case MaxPrice: // Yet, here is not further check if the given price is meaningful (it is assumed to already be the lowest currency denomination, e.g 3241 => 31.41)
		if price := filterValue.(int); price > 0 {
			r.MaxPrice = price
		}
	case Merchant:
		if param := filterValue.(string); filter.isVlidString(param) {
			r.Merchant = param
		}
	case MinPrice: // Yet, here is not further check if the given price is meaningful (it is assumed to already be the lowest currency denomination, e.g 3241 => 31.41)
		if price := filterValue.(int); price > 0 {
			r.MinPrice = price
		}
	case MinReviewsRating:
		if minRating := filterValue.(int); 0 < minRating && minRating < 5 {
			r.MinReviewsRating = minRating
		}
	case MinSavingPercent:
		if minSaving := filterValue.(int); 0 < minSaving && minSaving < 100 {
			r.MinSavingPercent = minSaving
		}
	case OfferCount:
		if oCount := filterValue.(int); oCount > 0 {
			r.OfferCount = oCount
		}
	case PartnerTag:
		if param := filterValue.(string); filter.isVlidString(param) {
			r.PartnerTag = param
		}
	case PartnerType:
		if param := filterValue.(string); filter.isVlidString(param) {
			r.PartnerType = param
		}
	case Properties:
		if params := filterValue.(map[string]string); len(params) > 0 {
			r.Properties = params
		}
	case SearchIndex:
		if param := filterValue.(string); filter.isVlidString(param) {
			r.SearchIndex = param
		}
	case SortBy:
		if param := filterValue.(string); filter.isVlidString(param) {
			r.SortBy = param
		}
	case Title:
		if param := filterValue.(string); filter.isVlidString(param) {
			r.Title = param
		}
	}
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
