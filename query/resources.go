package query

type resource int

const (
	resourceBrowseNodeInfo    resource = 1 + iota //BrowseNodeInfo resource
	resourceImages                                //Images resource
	resourceItemInfo                              //ItemInfo resource
	resourceOffers                                //Offers resource
	resourceOffersV2                              //OffersV2 resource
	resourceSearchRefinements                     //SearchRefinements resource
	resourceParentASIN                            //ParentASIN resource
	resourceCustomerReviews                       //CustomerReviews resource
	resourceBrowseNodes                           //BrowseNodes resource
	resourceVariationSummary                      //VariationSummary resource
)

var (
	//BrowseNodeInfo resource
	resourcesBrowseNodeInfo = []string{
		"BrowseNodeInfo.BrowseNodes",
		"BrowseNodeInfo.BrowseNodes.Ancestor",
		"BrowseNodeInfo.BrowseNodes.SalesRank",
		"BrowseNodeInfo.WebsiteSalesRank",
	}
	//Images resource
	resourcesImages = []string{
		"Images.Primary.Small",
		"Images.Primary.Medium",
		"Images.Primary.Large",
		"Images.Variants.Small",
		"Images.Variants.Medium",
		"Images.Variants.Large",
	}
	//ItemInfo resource
	resourcesItemInfo = []string{
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
	}
	//Offers resource
	resourcesOffers = []string{
		"Offers.Listings.Availability.MaxOrderQuantity",
		"Offers.Listings.Availability.Message",
		"Offers.Listings.Availability.MinOrderQuantity",
		"Offers.Listings.Availability.Type",
		"Offers.Listings.Condition",
		"Offers.Listings.Condition.SubCondition",
		"Offers.Listings.DeliveryInfo.IsAmazonFulfilled",
		"Offers.Listings.DeliveryInfo.IsFreeShippingEligible",
		"Offers.Listings.DeliveryInfo.IsPrimeEligible",
		"Offers.Listings.DeliveryInfo.ShippingCharges",
		"Offers.Listings.IsBuyBoxWinner",
		"Offers.Listings.LoyaltyPoints.Points",
		"Offers.Listings.MerchantInfo",
		"Offers.Listings.Price",
		"Offers.Listings.ProgramEligibility.IsPrimeExclusive",
		"Offers.Listings.ProgramEligibility.IsPrimePantry",
		"Offers.Listings.Promotions",
		"Offers.Listings.SavingBasis",
		"Offers.Summaries.HighestPrice",
		"Offers.Summaries.LowestPrice",
		"Offers.Summaries.OfferCount",
	}
	//OffersV2 resource
	resourcesOffersV2 = []string{
		"OffersV2.Listings.Availability",
		"OffersV2.Listings.Condition",
		"OffersV2.Listings.DealDetails",
		"OffersV2.Listings.IsBuyBoxWinner",
		"OffersV2.Listings.LoyaltyPoints",
		"OffersV2.Listings.MerchantInfo",
		"OffersV2.Listings.Price",
		"OffersV2.Listings.Type",
	}
	//SearchRefinements resource
	resourcesSearchRefinements = []string{
		"SearchRefinements",
	}
	//ParentASIN resource
	resourcesParentASIN = []string{
		"ParentASIN",
	}
	//CustomerReviews resource
	resourcesCustomerReviews = []string{
		"CustomerReviews.Count",
		"CustomerReviews.StarRating",
	}
	//BrowseNodes resource
	resourcesBrowseNodes = []string{
		"BrowseNodes.Ancestor",
		"BrowseNodes.Children",
	}

	//VariationSummary resource
	resourcesVariationSummary = []string{
		"VariationSummary.Price.HighestPrice",
		"VariationSummary.Price.LowestPrice",
		"VariationSummary.VariationDimension",
	}

	resourcesMap = map[resource][]string{
		resourceBrowseNodeInfo:    resourcesBrowseNodeInfo,    //BrowseNodeInfo resource
		resourceImages:            resourcesImages,            //Images resource
		resourceItemInfo:          resourcesItemInfo,          //ItemInfo resource
		resourceOffers:            resourcesOffers,            //Offers resource
		resourceOffersV2:          resourcesOffersV2,          //OffersV2 resource
		resourceSearchRefinements: resourcesSearchRefinements, //SearchRefinements resource
		resourceParentASIN:        resourcesParentASIN,        //ParentASIN resource
		resourceCustomerReviews:   resourcesCustomerReviews,   //CustomerReviews resource
		resourceBrowseNodes:       resourcesBrowseNodes,       //BrowseNodes resource
		resourceVariationSummary:  resourcesVariationSummary,  //VariationSummary resource
	}
)

func (r resource) Strings() []string {
	if ss, ok := resourcesMap[r]; ok {
		return ss
	}
	return []string{}
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
