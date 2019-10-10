package query

import (
	"encoding/json"
	"strconv"

	paapi5 "github.com/hackmac89/pa-api"
	"github.com/spiegel-im-spiegel/errs"
)

// the available (valid) filter parameters
var (
	validAvailabilityParameters []string = []string{"Available", "IncludeOutOfStock"}
	validConditionParameters    []string = []string{"Any", "New", "Used", "Collectible", "Refurbished"}
	validDeliveryParameters     []string = []string{"AmazonGlobal", "FreeShipping", "FulfilledByAmazon", "Prime"}
	validMerchantParameters     []string = []string{"All", "Amazon"}
	validSortByParameters       []string = []string{"AvgCustomerReviews", "Featured", "NewestArrivals", "Price:HighToLow", "Price:LowToHigh", "Relevance"}
)

//RequestType signals the type of request to search for
type RequestType int

//Constants for RequestType
const (
	Actor RequestType = iota + 1
	Artist
	Author
	Brand
	Keywords
	Title
)

//requestTypes is the private and anonymously imported struct, which selects the request type to be used
type requestTypes struct {
	Actor    string `json:",omitempty"`
	Artist   string `json:",omitempty"`
	Author   string `json:",omitempty"`
	Brand    string `json:",omitempty"`
	Keywords string `json:",omitempty"`
	Title    string `json:",omitempty"`
}

//RequestFilter signals the types of filters to use
type RequestFilter int

//Constants for RequestFilter
const (
	Availability RequestFilter = iota + 1
	BrowseNodeID
	Condition
	CurrencyOfPreference
	DeliveryFlags
	ItemCount
	ItemPage
	LanguagesOfPreference
	MaxPrice
	Merchant
	MinPrice
	MinReviewsRating
	MinSavingPercent
	OfferCount
	SearchIndex
	SortBy
)

//filters is the private and anonymously imported struct, which selects the filters to be used
type filters struct {
	Availability          string   `json:",omitempty"`
	BrowseNodeID          string   `json:"BrowseNodeId,omitempty"`
	Condition             string   `json:",omitempty"`
	CurrencyOfPreference  string   `json:",omitempty"`
	DeliveryFlags         []string `json:",omitempty"`
	ItemCount             int      `json:",omitempty"`
	ItemPage              int      `json:",omitempty"`
	LanguagesOfPreference []string `json:",omitempty"`
	MaxPrice              int      `json:",omitempty"`
	Merchant              string   `json:",omitempty"`
	MinPrice              int      `json:",omitempty"`
	MinReviewsRating      int      `json:",omitempty"`
	MinSavingPercent      int      `json:",omitempty"`
	OfferCount            int      `json:",omitempty"`
	SearchIndex           string   `json:",omitempty"`
	SortBy                string   `json:",omitempty"`
}

//SearchItems is a query data class for PA-API v5
type SearchItems struct {
	requestTypes
	filters
	OpeCode                 paapi5.Operation `json:"Operation"`
	Marketplace             string
	PartnerTag              string
	PartnerType             string
	Resources               []string `json:",omitempty"`
	enableBrowseNodeInfo    bool
	enableImages            bool
	enableItemInfo          bool
	enableOffers            bool
	enableSearchRefinements bool
	enableParentASIN        bool
}

var _ paapi5.Query = (*SearchItems)(nil) //SearchItems is compatible with paapi5.Query interface

//NewSearchItems creates a new SearchItems instance
func NewSearchItems(marketplace, partnerTag, partnerType string) *SearchItems {
	q := &SearchItems{
		OpeCode:        paapi5.SearchItems,
		Marketplace:    marketplace,
		PartnerTag:     partnerTag,
		PartnerType:    partnerType,
		enableItemInfo: true,
	}
	return q
}
func newNilSearchItems() *SearchItems { return NewSearchItems("", "", "") }

//Search is a generic search query funtion to obtain informations from the "SearchItems"-operation
func (q *SearchItems) Search(searchParam string, searchType RequestType) *SearchItems {
	if q.Operation() == paapi5.SearchItems {
		if q == nil {
			q = newNilSearchItems()
		}
		switch searchType {
		case Actor:
			q.Actor = searchParam
		case Artist:
			q.Artist = searchParam
		case Author:
			q.Author = searchParam
		case Brand:
			q.Brand = searchParam
		case Keywords:
			q.Keywords = searchParam
		case Title:
			q.Title = searchParam
		}
	}
	return q
}

//WithFilters adds filtered requests to the search query
func (q *SearchItems) WithFilters(filters ...map[RequestFilter]interface{}) *SearchItems {
	if q.Operation() == paapi5.SearchItems {
		if q == nil {
			q = newNilSearchItems()
		}
		for _, currentFilter := range filters {
			for filter, filterValue := range currentFilter {
				q.mapFilter(filter, filterValue)
			}
		}
	}
	return q
}

//mapFilter is a helper function for (*SearchItems).WithFilters
// This function does not check, if the filters to be used match the chosen searchParam/searchType (Actor, Artist etc.pp.)
// TODO: 	- reduce nesting
func (q *SearchItems) mapFilter(filter RequestFilter, filterValue interface{}) {
	switch filter {
	case Availability:
		if param := filterValue.(string); isFilterParamValid(param, validAvailabilityParameters) {
			q.Availability = param
		}
	case BrowseNodeID:
		if param := filterValue.(string); stringIsNumber(param) {
			q.BrowseNodeID = param
		}
	case Condition:
		if param := filterValue.(string); isFilterParamValid(param, validConditionParameters) {
			q.Condition = param
		}
	case CurrencyOfPreference:
		if currency := filterValue.(string); len(currency) > 0 {
			q.CurrencyOfPreference = currency
		}
	case DeliveryFlags:
		q.DeliveryFlags = []string{}
		switch filterValue.(type) {
		case []string:
			for _, deliveryFlag := range filterValue.([]string) {
				if !isFilterParamValid(deliveryFlag, validDeliveryParameters) {
					continue
				}
				q.DeliveryFlags = append(q.DeliveryFlags, deliveryFlag)
			}
		case string:
			if param := filterValue.(string); isFilterParamValid(param, validDeliveryParameters) {
				q.DeliveryFlags = append(q.DeliveryFlags, param)
			}
		}
	case ItemCount:
		if count := filterValue.(int); 0 < count && count < 11 {
			q.ItemCount = count
		}
	case ItemPage:
		if page := filterValue.(int); 0 < page && page < 11 {
			q.ItemPage = page
		}
	case LanguagesOfPreference:
		q.LanguagesOfPreference = []string{}
		switch filterValue.(type) {
		case []string:
			for _, language := range filterValue.([]string) {
				q.LanguagesOfPreference = append(q.LanguagesOfPreference, language)
			}
		case string:
			q.LanguagesOfPreference = append(q.LanguagesOfPreference, filterValue.(string))
		}
	case MaxPrice: // Yet, here is not further check if the given price is meaningful (it is assumed to already be the lowest currency denomination, e.g 3241 => 31.41)
		if price := filterValue.(int); price > 0 {
			q.MaxPrice = price
		}
	case Merchant:
		if param := filterValue.(string); isFilterParamValid(param, validMerchantParameters) {
			q.Merchant = param
		}
	case MinPrice: // Yet, here is not further check if the given price is meaningful (it is assumed to already be the lowest currency denomination, e.g 3241 => 31.41)
		if price := filterValue.(int); price > 0 {
			q.MinPrice = price
		}
	case MinReviewsRating:
		if minRating := filterValue.(int); 0 < minRating && minRating < 5 {
			q.MinReviewsRating = minRating
		}
	case MinSavingPercent:
		if minSaving := filterValue.(int); 0 < minSaving && minSaving < 100 {
			q.MinSavingPercent = minSaving
		}
	case OfferCount:
		if oCount := filterValue.(int); oCount > 0 {
			q.OfferCount = oCount
		}
	case SearchIndex:
		if sIdx := filterValue.(string); len(sIdx) > 0 {
			q.SearchIndex = sIdx
		}
	case SortBy:
		if param := filterValue.(string); isFilterParamValid(param, validSortByParameters) {
			q.SortBy = param
		}
	}
}

//EnableBrowseNodeInfo sets the enableBrowseNodeInfo flag in SearchItems instance
func (q *SearchItems) EnableBrowseNodeInfo(flag bool) *SearchItems {
	if q == nil {
		q = newNilSearchItems()
	}
	q.enableBrowseNodeInfo = flag
	return q
}

//EnableImages sets the enableImages flag in SearchItems instance
func (q *SearchItems) EnableImages(flag bool) *SearchItems {
	if q == nil {
		q = newNilSearchItems()
	}
	q.enableImages = flag
	return q
}

//EnableItemInfo sets the enableItemInfo flag in SearchItems instance
func (q *SearchItems) EnableItemInfo(flag bool) *SearchItems {
	if q == nil {
		q = newNilSearchItems()
	}
	q.enableItemInfo = flag
	return q
}

//EnableOffers sets the enableOffers flag in SearchItems instance
func (q *SearchItems) EnableOffers(flag bool) *SearchItems {
	if q == nil {
		q = newNilSearchItems()
	}
	q.enableOffers = flag
	return q
}

//EnableSearchRefinements sets the enableOffers flag in SearchItems instance
func (q *SearchItems) EnableSearchRefinements(flag bool) *SearchItems {
	if q == nil {
		q = newNilSearchItems()
	}
	q.enableSearchRefinements = flag
	return q
}

//EnableParentASIN sets the enableParentASIN flag in SearchItems instance
func (q *SearchItems) EnableParentASIN(flag bool) *SearchItems {
	if q == nil {
		q = newNilSearchItems()
	}
	q.enableParentASIN = flag
	return q
}

//Operation returns the type of the PA-API operation
func (q *SearchItems) Operation() paapi5.Operation {
	if q == nil {
		return paapi5.NullOperation
	}
	return q.OpeCode
}

//Payload defines the resources to be returned
func (q *SearchItems) Payload() ([]byte, error) {
	if q == nil {
		return nil, errs.Wrap(paapi5.ErrNullPointer, "")
	}
	q.Resources = []string{}
	if q.enableBrowseNodeInfo {
		q.Resources = append(
			q.Resources,
			"BrowseNodeInfo.BrowseNodes",
			"BrowseNodeInfo.BrowseNodes.Ancestor",
			"BrowseNodeInfo.BrowseNodes.SalesRank",
			"BrowseNodeInfo.WebsiteSalesRank",
		)
	}
	if q.enableImages {
		q.Resources = append(
			q.Resources,
			"Images.Primary.Small",
			"Images.Primary.Medium",
			"Images.Primary.Large",
			"Images.Variants.Small",
			"Images.Variants.Medium",
			"Images.Variants.Large",
		)
	}
	if q.enableItemInfo {
		q.Resources = append(
			q.Resources,
			"ItemInfo.ByLineInfo",
			"ItemInfo.Classifications",
			"ItemInfo.ContentInfo",
			"ItemInfo.ContentRating",
			"ItemInfo.ExternalIds",
			"ItemInfo.Features",
			"ItemInfo.ManufactureInfo",
			"ItemInfo.ProductInfo",
			"ItemInfo.TechnicalInfo",
			"ItemInfo.Title",
			"ItemInfo.TradeInInfo",
		)
	}
	// Offers currently skipped (as in "getitems.go")
	//...
	if q.enableSearchRefinements {
		q.Resources = append(
			q.Resources,
			"SearchRefinements",
		)
	}
	if q.enableParentASIN {
		q.Resources = append(
			q.Resources,
			"ParentASIN",
		)
	}
	b, err := json.Marshal(q)
	return b, errs.Wrap(err, "")
}

//Stringer interface
func (q *SearchItems) String() string {
	b, err := q.Payload()
	if err != nil {
		return ""
	}
	return string(b)
}

//isFilterParamValid checks if the given parameter is valid for the chosen filter option
func isFilterParamValid(param string, params []string) bool {
	for _, currentParam := range params {
		if param == currentParam {
			return true
		}
	}

	return false
}

//stringIsNumber checks if the given string can be parsed and used as an int
func stringIsNumber(s string) bool {
	if _, convErr := strconv.ParseInt(s, 10, 64); convErr != nil {
		return false
	}
	return true
}
