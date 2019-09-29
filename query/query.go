package query

import (
	"encoding/json"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/pa-api/errcode"
)

//Query is query data class for PA-API v5
type Query struct {
	Operation               Operation
	Marketplace             string
	PartnerTag              string
	PartnerType             string
	ASIN                    string   `json:",omitempty"`
	VariationPage           *int     `json:",omitempty"`
	ItemIds                 []string `json:",omitempty"`
	ItemIdType              string   `json:",omitempty"`
	Keywords                string   `json:",omitempty"`
	Resources               []string `json:",omitempty"`
	enableBrowseNodeInfo    bool
	enableImages            bool
	enableItemInfo          bool
	enableOffers            bool
	enableParentASIN        bool
	enableSearchRefinements bool
	enableVariationSummary  bool
}

//New creates new Query instance
func New(operation Operation, marketplace, partnerTag, partnerType string) *Query {
	q := &Query{
		Operation:      operation,
		Marketplace:    marketplace,
		PartnerTag:     partnerTag,
		PartnerType:    partnerType,
		enableItemInfo: true,
	}
	switch operation {
	case GetVariations:
		q.enableVariationSummary = true
	case GetItems:
	case SearchItems:
		q.enableSearchRefinements = true
	default:
	}
	return q
}
func newNil() *Query { return New(Operation(0), "", "", "") }

//ASINs sets ItemIds in Query instance
func (q *Query) ASINcode(asin string, pageNum int) *Query {
	if q.Operation == GetVariations {
		if q == nil {
			q = newNil()
		}
		q.ASIN = asin
		*q.VariationPage = pageNum
		q.ItemIds = nil
		q.ItemIdType = ""
		q.Keywords = ""
	}
	return q
}

//ASINs sets ItemIds in Query instance
func (q *Query) ASINs(itms []string) *Query {
	if q.Operation == GetItems {
		if q == nil {
			q = newNil()
		}
		q.ASIN = ""
		q.VariationPage = nil
		q.ItemIds = itms
		q.ItemIdType = "ASIN"
		q.Keywords = ""
	}
	return q
}

//Keyword sets Keywords in Query instance
func (q *Query) Keyword(k string) *Query {
	if q.Operation == SearchItems {
		if q == nil {
			q = newNil()
		}
		q.ASIN = ""
		q.VariationPage = nil
		q.ItemIds = nil
		q.ItemIdType = ""
		q.Keywords = k
	}
	return q
}

//EnableBrowseNodeInfo sets enableBrowseNodeInfo flag in Query instance
func (q *Query) EnableBrowseNodeInfo(flag bool) *Query {
	if q == nil {
		q = newNil()
	}
	q.enableBrowseNodeInfo = flag
	return q
}

//EnableImages sets enableImages flag in Query instance
func (q *Query) EnableImages(flag bool) *Query {
	if q == nil {
		q = newNil()
	}
	q.enableImages = flag
	return q
}

//EnableItemInfo sets enableItemInfo flag in Query instance
func (q *Query) EnableItemInfo(flag bool) *Query {
	if q == nil {
		q = newNil()
	}
	q.enableItemInfo = flag
	return q
}

//EnableOffers sets enableOffers flag in Query instance
func (q *Query) EnableOffers(flag bool) *Query {
	if q == nil {
		q = newNil()
	}
	q.enableOffers = flag
	return q
}

//EnableParentASIN sets enableParentASIN flag in Query instance
func (q *Query) EnableParentASIN(flag bool) *Query {
	if q == nil {
		q = newNil()
	}
	q.enableParentASIN = flag
	return q
}

//EnableSearchRefinements sets enableSearchRefinements flag in Query instance
func (q *Query) EnableSearchRefinements(flag bool) *Query {
	if q == nil {
		q = newNil()
	}
	q.enableSearchRefinements = flag
	return q
}

//EnableVariationSummary sets enableVariationSummary flag in Query instance
func (q *Query) EnableVariationSummary(flag bool) *Query {
	if q == nil {
		q = newNil()
	}
	q.enableVariationSummary = flag
	return q
}

func (q *Query) String() string {
	b, err := q.JSON()
	if err != nil {
		return ""
	}
	return string(b)
}

func (q *Query) JSON() ([]byte, error) {
	if q == nil {
		return nil, errs.Wrap(errcode.ErrNullPointer, "")
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
			"ItemInfo.ContentInfo",
			"ItemInfo.ContentRating",
			"ItemInfo.Classifications",
			"ItemInfo.ExternalIds",
			"ItemInfo.Features",
			"ItemInfo.ManufactureInfo",
			"ItemInfo.ProductInfo",
			"ItemInfo.TechnicalInfo",
			"ItemInfo.Title",
			"ItemInfo.TradeInInfo",
		)
	}
	// if q.enableOffers {
	// 	q.Resources = append(
	// 		q.Resources,
	// 		"Offers.Listings.Availability.MaxOrderQuantity",
	// 		"Offers.Listings.Availability.Message",
	// 		"Offers.Listings.Availability.MinOrderQuantity",
	// 		"Offers.Listings.Availability.Type",
	// 		"Offers.Listings.Condition",
	// 		"Offers.Listings.Condition.SubCondition",
	// 		"Offers.Listings.DeliveryInfo.IsAmazonFulfilled",
	// 		"Offers.Listings.DeliveryInfo.IsFreeShippingEligible",
	// 		"Offers.Listings.DeliveryInfo.IsPrimeEligible",
	// 		"Offers.Listings.DeliveryInfo.ShippingCharges",
	// 		"Offers.Listings.IsBuyBoxWinner",
	// 		"Offers.Listings.LoyaltyPoints.Points",
	// 		"Offers.Listings.MerchantInfo",
	// 		"Offers.Listings.Price",
	// 		"Offers.Listings.ProgramEligibility.IsPrimeExclusive",
	// 		"Offers.Listings.ProgramEligibility.IsPrimePantry",
	// 		"Offers.Listings.Promotions",
	// 		"Offers.Listings.SavingBasis",
	// 		"Offers.Summaries.HighestPrice",
	// 		"Offers.Summaries.LowestPrice",
	// 		"Offers.Summaries.OfferCount",
	// 	)
	// }
	if q.enableParentASIN {
		q.Resources = append(
			q.Resources,
			"ParentASIN",
		)
	}
	if q.enableSearchRefinements && q.Operation == SearchItems {
		q.Resources = append(
			q.Resources,
			"SearchRefinements.SearchIndex",
			"SearchRefinements.BrowseNode",
			"SearchRefinements.OtherRefinements",
		)
	}
	if q.enableVariationSummary && q.Operation == GetVariations {
		q.Resources = append(
			q.Resources,
			"VariationSummary.Price.HighestPrice",
			"VariationSummary.Price.LowestPrice",
			"VariationSummary.VariationDimension",
		)
	}
	b, err := json.Marshal(q)
	return b, errs.Wrap(err, "")
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
