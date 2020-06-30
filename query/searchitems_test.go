package query

import "testing"

func TestSearchItems(t *testing.T) {
	testCases := []struct {
		q   *SearchItems
		str string
	}{
		{q: NewSearchItems("foo.bar", "mytag-20", "Associates"), str: `{"Operation":"SearchItems","Marketplace":"foo.bar","PartnerTag":"mytag-20","PartnerType":"Associates"}`},
		{q: NewSearchItems("foo.bar", "mytag-20", "Associates").Search(ItemIds, "foo"), str: `{"Operation":"SearchItems","Marketplace":"foo.bar","PartnerTag":"mytag-20","PartnerType":"Associates"}`},
		{q: NewSearchItems("foo.bar", "mytag-20", "Associates").Search(Actor, "foo"), str: `{"Operation":"SearchItems","Actor":"foo","Marketplace":"foo.bar","PartnerTag":"mytag-20","PartnerType":"Associates"}`},
		{q: NewSearchItems("foo.bar", "mytag-20", "Associates").Search(Artist, "foo"), str: `{"Operation":"SearchItems","Artist":"foo","Marketplace":"foo.bar","PartnerTag":"mytag-20","PartnerType":"Associates"}`},
		{q: NewSearchItems("foo.bar", "mytag-20", "Associates").Search(Author, "foo"), str: `{"Operation":"SearchItems","Author":"foo","Marketplace":"foo.bar","PartnerTag":"mytag-20","PartnerType":"Associates"}`},
		{q: NewSearchItems("foo.bar", "mytag-20", "Associates").Search(Brand, "foo"), str: `{"Operation":"SearchItems","Brand":"foo","Marketplace":"foo.bar","PartnerTag":"mytag-20","PartnerType":"Associates"}`},
		{q: NewSearchItems("foo.bar", "mytag-20", "Associates").Search(Keywords, "foo"), str: `{"Operation":"SearchItems","Keywords":"foo","Marketplace":"foo.bar","PartnerTag":"mytag-20","PartnerType":"Associates"}`},
		{q: NewSearchItems("foo.bar", "mytag-20", "Associates").Search(Title, "foo"), str: `{"Operation":"SearchItems","Marketplace":"foo.bar","PartnerTag":"mytag-20","PartnerType":"Associates","Title":"foo"}`},
	}
	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("SearchItems.String() is \"%v\", want \"%v\"", str, tc.str)
		}
	}
}

func TestRequestInSearchItems(t *testing.T) {
	testCases := []struct {
		q   *SearchItems
		str string
	}{
		{q: NewSearchItems("", "", ""), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(Actor, "foo"), str: `{"Operation":"SearchItems","Actor":"foo"}`},
		{q: NewSearchItems("", "", "").Request(Artist, "foo"), str: `{"Operation":"SearchItems","Artist":"foo"}`},
		{q: NewSearchItems("", "", "").Request(Availability, "foo"), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(Availability, "Available"), str: `{"Operation":"SearchItems","Availability":"Available"}`},
		{q: NewSearchItems("", "", "").Request(Availability, "IncludeOutOfStock"), str: `{"Operation":"SearchItems","Availability":"IncludeOutOfStock"}`},
		{q: NewSearchItems("", "", "").Request(Author, "foo"), str: `{"Operation":"SearchItems","Author":"foo"}`},
		{q: NewSearchItems("", "", "").Request(Brand, "foo"), str: `{"Operation":"SearchItems","Brand":"foo"}`},
		{q: NewSearchItems("", "", "").Request(BrowseNodeID, "foo"), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(BrowseNodeID, "123"), str: `{"Operation":"SearchItems","BrowseNodeId":"123"}`},
		{q: NewSearchItems("", "", "").Request(Condition, "foo"), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(Condition, "Any"), str: `{"Operation":"SearchItems","Condition":"Any"}`},
		{q: NewSearchItems("", "", "").Request(Condition, "New"), str: `{"Operation":"SearchItems","Condition":"New"}`},
		{q: NewSearchItems("", "", "").Request(Condition, "Used"), str: `{"Operation":"SearchItems","Condition":"Used"}`},
		{q: NewSearchItems("", "", "").Request(Condition, "Collectible"), str: `{"Operation":"SearchItems","Condition":"Collectible"}`},
		{q: NewSearchItems("", "", "").Request(Condition, "Refurbished"), str: `{"Operation":"SearchItems","Condition":"Refurbished"}`},
		{q: NewSearchItems("", "", "").Request(CurrencyOfPreference, "foo"), str: `{"Operation":"SearchItems","CurrencyOfPreference":"foo"}`},
		{q: NewSearchItems("", "", "").Request(DeliveryFlags, "foo"), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(DeliveryFlags, "AmazonGlobal"), str: `{"Operation":"SearchItems","DeliveryFlags":["AmazonGlobal"]}`},
		{q: NewSearchItems("", "", "").Request(DeliveryFlags, []string{"AmazonGlobal", "FreeShipping", "FulfilledByAmazon", "Prime"}), str: `{"Operation":"SearchItems","DeliveryFlags":["AmazonGlobal","FreeShipping","FulfilledByAmazon","Prime"]}`},
		{q: NewSearchItems("", "", "").Request(ItemIds, "4900900028"), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(ItemIdType, "ASIN"), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(ItemCount, -1), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(ItemCount, 0), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(ItemCount, 1), str: `{"Operation":"SearchItems","ItemCount":1}`},
		{q: NewSearchItems("", "", "").Request(ItemCount, 10), str: `{"Operation":"SearchItems","ItemCount":10}`},
		{q: NewSearchItems("", "", "").Request(ItemCount, 11), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(ItemPage, -1), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(ItemPage, 0), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(ItemPage, 1), str: `{"Operation":"SearchItems","ItemPage":1}`},
		{q: NewSearchItems("", "", "").Request(ItemPage, 10), str: `{"Operation":"SearchItems","ItemPage":10}`},
		{q: NewSearchItems("", "", "").Request(ItemPage, 11), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(Keywords, "foo"), str: `{"Operation":"SearchItems","Keywords":"foo"}`},
		{q: NewSearchItems("", "", "").Request(LanguagesOfPreference, "foo"), str: `{"Operation":"SearchItems","LanguagesOfPreference":["foo"]}`},
		{q: NewSearchItems("", "", "").Request(LanguagesOfPreference, []string{"foo", "bar"}), str: `{"Operation":"SearchItems","LanguagesOfPreference":["foo","bar"]}`},
		{q: NewSearchItems("", "", "").Request(Marketplace, "foo.bar"), str: `{"Operation":"SearchItems","Marketplace":"foo.bar"}`},
		{q: NewSearchItems("", "", "").Request(MaxPrice, -1), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(MaxPrice, 0), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(MaxPrice, 1), str: `{"Operation":"SearchItems","MaxPrice":1}`},
		{q: NewSearchItems("", "", "").Request(MaxPrice, 123), str: `{"Operation":"SearchItems","MaxPrice":123}`},
		{q: NewSearchItems("", "", "").Request(Merchant, "foo"), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(Merchant, "All"), str: `{"Operation":"SearchItems","Merchant":"All"}`},
		{q: NewSearchItems("", "", "").Request(Merchant, "Amazon"), str: `{"Operation":"SearchItems","Merchant":"Amazon"}`},
		{q: NewSearchItems("", "", "").Request(MinPrice, -1), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(MinPrice, 0), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(MinPrice, 1), str: `{"Operation":"SearchItems","MinPrice":1}`},
		{q: NewSearchItems("", "", "").Request(MinPrice, 123), str: `{"Operation":"SearchItems","MinPrice":123}`},
		{q: NewSearchItems("", "", "").Request(MinReviewsRating, -1), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(MinReviewsRating, 0), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(MinReviewsRating, 1), str: `{"Operation":"SearchItems","MinReviewsRating":1}`},
		{q: NewSearchItems("", "", "").Request(MinReviewsRating, 4), str: `{"Operation":"SearchItems","MinReviewsRating":4}`},
		{q: NewSearchItems("", "", "").Request(MinReviewsRating, 5), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(MinSavingPercent, -1), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(MinSavingPercent, 0), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(MinSavingPercent, 1), str: `{"Operation":"SearchItems","MinSavingPercent":1}`},
		{q: NewSearchItems("", "", "").Request(MinSavingPercent, 99), str: `{"Operation":"SearchItems","MinSavingPercent":99}`},
		{q: NewSearchItems("", "", "").Request(MinSavingPercent, 100), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(OfferCount, -1), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(OfferCount, 0), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(OfferCount, 1), str: `{"Operation":"SearchItems","OfferCount":1}`},
		{q: NewSearchItems("", "", "").Request(OfferCount, 123), str: `{"Operation":"SearchItems","OfferCount":123}`},
		{q: NewSearchItems("", "", "").Request(PartnerTag, "foo"), str: `{"Operation":"SearchItems","PartnerTag":"foo"}`},
		{q: NewSearchItems("", "", "").Request(PartnerType, "foo"), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(PartnerType, "Associates"), str: `{"Operation":"SearchItems","PartnerType":"Associates"}`},
		{q: NewSearchItems("", "", "").Request(Properties, map[string]string{"foo": "bar"}), str: `{"Operation":"SearchItems","Properties":{"foo":"bar"}}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "foo"), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "All"), str: `{"Operation":"SearchItems","SearchIndex":"All"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "AmazonVideo"), str: `{"Operation":"SearchItems","SearchIndex":"AmazonVideo"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Apparel"), str: `{"Operation":"SearchItems","SearchIndex":"Apparel"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Appliances"), str: `{"Operation":"SearchItems","SearchIndex":"Appliances"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "ArtsAndCrafts"), str: `{"Operation":"SearchItems","SearchIndex":"ArtsAndCrafts"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Automotive"), str: `{"Operation":"SearchItems","SearchIndex":"Automotive"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Baby"), str: `{"Operation":"SearchItems","SearchIndex":"Baby"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Beauty"), str: `{"Operation":"SearchItems","SearchIndex":"Beauty"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Books"), str: `{"Operation":"SearchItems","SearchIndex":"Books"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Classical"), str: `{"Operation":"SearchItems","SearchIndex":"Classical"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Collectibles"), str: `{"Operation":"SearchItems","SearchIndex":"Collectibles"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Computers"), str: `{"Operation":"SearchItems","SearchIndex":"Computers"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "DigitalMusic"), str: `{"Operation":"SearchItems","SearchIndex":"DigitalMusic"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Electronics"), str: `{"Operation":"SearchItems","SearchIndex":"Electronics"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "EverythingElse"), str: `{"Operation":"SearchItems","SearchIndex":"EverythingElse"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Fashion"), str: `{"Operation":"SearchItems","SearchIndex":"Fashion"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "FashionBaby"), str: `{"Operation":"SearchItems","SearchIndex":"FashionBaby"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "FashionBoys"), str: `{"Operation":"SearchItems","SearchIndex":"FashionBoys"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "FashionGirls"), str: `{"Operation":"SearchItems","SearchIndex":"FashionGirls"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "FashionMen"), str: `{"Operation":"SearchItems","SearchIndex":"FashionMen"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "FashionWomen"), str: `{"Operation":"SearchItems","SearchIndex":"FashionWomen"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "GardenAndOutdoor"), str: `{"Operation":"SearchItems","SearchIndex":"GardenAndOutdoor"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "GiftCards"), str: `{"Operation":"SearchItems","SearchIndex":"GiftCards"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "GroceryAndGourmetFood"), str: `{"Operation":"SearchItems","SearchIndex":"GroceryAndGourmetFood"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Handmade"), str: `{"Operation":"SearchItems","SearchIndex":"Handmade"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "HealthPersonalCare"), str: `{"Operation":"SearchItems","SearchIndex":"HealthPersonalCare"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "HomeAndKitchen"), str: `{"Operation":"SearchItems","SearchIndex":"HomeAndKitchen"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Industrial"), str: `{"Operation":"SearchItems","SearchIndex":"Industrial"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Jewelry"), str: `{"Operation":"SearchItems","SearchIndex":"Jewelry"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "KindleStore"), str: `{"Operation":"SearchItems","SearchIndex":"KindleStore"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "LocalServices"), str: `{"Operation":"SearchItems","SearchIndex":"LocalServices"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Luggage"), str: `{"Operation":"SearchItems","SearchIndex":"Luggage"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "LuxuryBeauty"), str: `{"Operation":"SearchItems","SearchIndex":"LuxuryBeauty"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Magazines"), str: `{"Operation":"SearchItems","SearchIndex":"Magazines"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "MobileAndAccessories"), str: `{"Operation":"SearchItems","SearchIndex":"MobileAndAccessories"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "MobileApps"), str: `{"Operation":"SearchItems","SearchIndex":"MobileApps"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "MoviesAndTV"), str: `{"Operation":"SearchItems","SearchIndex":"MoviesAndTV"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Music"), str: `{"Operation":"SearchItems","SearchIndex":"Music"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "MusicalInstruments"), str: `{"Operation":"SearchItems","SearchIndex":"MusicalInstruments"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "OfficeProducts"), str: `{"Operation":"SearchItems","SearchIndex":"OfficeProducts"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "PetSupplies"), str: `{"Operation":"SearchItems","SearchIndex":"PetSupplies"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Photo"), str: `{"Operation":"SearchItems","SearchIndex":"Photo"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Shoes"), str: `{"Operation":"SearchItems","SearchIndex":"Shoes"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Software"), str: `{"Operation":"SearchItems","SearchIndex":"Software"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "SportsAndOutdoors"), str: `{"Operation":"SearchItems","SearchIndex":"SportsAndOutdoors"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "ToolsAndHomeImprovement"), str: `{"Operation":"SearchItems","SearchIndex":"ToolsAndHomeImprovement"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "ToysAndGames"), str: `{"Operation":"SearchItems","SearchIndex":"ToysAndGames"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "VHS"), str: `{"Operation":"SearchItems","SearchIndex":"VHS"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "VideoGames"), str: `{"Operation":"SearchItems","SearchIndex":"VideoGames"}`},
		{q: NewSearchItems("", "", "").Request(SearchIndex, "Watches"), str: `{"Operation":"SearchItems","SearchIndex":"Watches"}`},
		{q: NewSearchItems("", "", "").Request(SortBy, "foo"), str: `{"Operation":"SearchItems"}`},
		{q: NewSearchItems("", "", "").Request(SortBy, "AvgCustomerReviews"), str: `{"Operation":"SearchItems","SortBy":"AvgCustomerReviews"}`},
		{q: NewSearchItems("", "", "").Request(SortBy, "Featured"), str: `{"Operation":"SearchItems","SortBy":"Featured"}`},
		{q: NewSearchItems("", "", "").Request(SortBy, "NewestArrivals"), str: `{"Operation":"SearchItems","SortBy":"NewestArrivals"}`},
		{q: NewSearchItems("", "", "").Request(SortBy, "Price:HighToLow"), str: `{"Operation":"SearchItems","SortBy":"Price:HighToLow"}`},
		{q: NewSearchItems("", "", "").Request(SortBy, "Price:LowToHigh"), str: `{"Operation":"SearchItems","SortBy":"Price:LowToHigh"}`},
		{q: NewSearchItems("", "", "").Request(SortBy, "Relevance"), str: `{"Operation":"SearchItems","SortBy":"Relevance"}`},
		{q: NewSearchItems("", "", "").Request(Title, "foo"), str: `{"Operation":"SearchItems","Title":"foo"}`},
	}

	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("Query.String() is \"%v\", want \"%v\"", str, tc.str)
		}
	}
}

func TestResourcesInSearchItems(t *testing.T) {
	testCases := []struct {
		q   *SearchItems
		str string
	}{
		{q: NewSearchItems("", "", "").EnableBrowseNodeInfo(), str: `{"Operation":"SearchItems","Resources":["BrowseNodeInfo.BrowseNodes","BrowseNodeInfo.BrowseNodes.Ancestor","BrowseNodeInfo.BrowseNodes.SalesRank","BrowseNodeInfo.WebsiteSalesRank"]}`},
		{q: NewSearchItems("", "", "").EnableImages(), str: `{"Operation":"SearchItems","Resources":["Images.Primary.Small","Images.Primary.Medium","Images.Primary.Large","Images.Variants.Small","Images.Variants.Medium","Images.Variants.Large"]}`},
		{q: NewSearchItems("", "", "").EnableItemInfo(), str: `{"Operation":"SearchItems","Resources":["ItemInfo.ByLineInfo","ItemInfo.ContentInfo","ItemInfo.ContentRating","ItemInfo.Classifications","ItemInfo.ExternalIds","ItemInfo.Features","ItemInfo.ManufactureInfo","ItemInfo.ProductInfo","ItemInfo.TechnicalInfo","ItemInfo.Title","ItemInfo.TradeInInfo"]}`},
		{q: NewSearchItems("", "", "").EnableOffers(), str: `{"Operation":"SearchItems","Resources":["Offers.Listings.Availability.MaxOrderQuantity","Offers.Listings.Availability.Message","Offers.Listings.Availability.MinOrderQuantity","Offers.Listings.Availability.Type","Offers.Listings.Condition","Offers.Listings.Condition.SubCondition","Offers.Listings.DeliveryInfo.IsAmazonFulfilled","Offers.Listings.DeliveryInfo.IsFreeShippingEligible","Offers.Listings.DeliveryInfo.IsPrimeEligible","Offers.Listings.DeliveryInfo.ShippingCharges","Offers.Listings.IsBuyBoxWinner","Offers.Listings.LoyaltyPoints.Points","Offers.Listings.MerchantInfo","Offers.Listings.Price","Offers.Listings.ProgramEligibility.IsPrimeExclusive","Offers.Listings.ProgramEligibility.IsPrimePantry","Offers.Listings.Promotions","Offers.Listings.SavingBasis","Offers.Summaries.HighestPrice","Offers.Summaries.LowestPrice","Offers.Summaries.OfferCount"]}`},
		{q: NewSearchItems("", "", "").EnableSearchRefinements(), str: `{"Operation":"SearchItems","Resources":["SearchRefinements"]}`},
		{q: NewSearchItems("", "", "").EnableParentASIN(), str: `{"Operation":"SearchItems","Resources":["ParentASIN"]}`},
		{q: NewSearchItems("", "", "").EnableCustomerReviews(), str: `{"Operation":"SearchItems","Resources":["CustomerReviews.Count","CustomerReviews.StarRating"]}`},
	}

	for _, tc := range testCases {
		if str := tc.q.String(); str != tc.str {
			t.Errorf("Query.String() is \"%v\", want \"%v\"", str, tc.str)
		}
	}
}

// func TestNewSearchItems(t *testing.T) {
// 	type args struct {
// 		marketplace string
// 		partnerTag  string
// 		partnerType string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *SearchItems
// 	}{
// 		{
// 			"TestNewSearchItems",
// 			args{
// 				"www.amazon.co.jp",
// 				"wwwyourpartnertag-20",
// 				"Associate",
// 			},
// 			&SearchItems{
// 				Query: &Query{
// 					OpeCode: paapi5.SearchItems,
// 					request: request{
// 						Marketplace: "www.amazon.co.jp",
// 						PartnerTag:  "wwwyourpartnertag-20",
// 						PartnerType: "Associate",
// 					},
// 					//enableResources: map[resource]bool{resourceItemInfo: true},
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewSearchItems(tt.args.marketplace, tt.args.partnerTag, tt.args.partnerType); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewSearchItems() = %[1]v (%[1]T), want %[2]v (%[2]T)", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func newNilSearchItems() *SearchItems {
// 	return &SearchItems{
// 		Query: &Query{
// 			OpeCode: paapi5.SearchItems,
// 			request: request{
// 				Marketplace: "",
// 				PartnerTag:  "",
// 				PartnerType: "",
// 			},
// 			enableResources: map[resource]bool{},
// 		},
// 	}
// }
//
// func Test_newNilSearchItems(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		want    *SearchItems
// 		wantErr bool
// 	}{
// 		{
// 			"Test_newNilSearchItems #1",
// 			newStandardSearchItem(),
// 			false,
// 		},
// 		{
// 			"Test_newNilSearchItems #2",
// 			&SearchItems{
// 				Query: &Query{
// 					OpeCode: paapi5.SearchItems,
// 					request: request{
// 						Marketplace: "",
// 						PartnerTag:  "",
// 						PartnerType: "",
// 					},
// 					enableResources: map[resource]bool{resourceItemInfo: true},
// 				},
// 			},
// 			true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := newNilSearchItems(); !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
// 				t.Errorf("newNilSearchItems() = %[1]v (%[1]T), want %[2]v (%[2]T)", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestSearchItems_Search(t *testing.T) {
// 	finalTestStruct := newStandardSearchItem()
// 	finalTestStruct.Actor = "Tom Cruise"
//
// 	type args struct {
// 		searchParam string
// 		searchType  RequestFilter
// 	}
// 	tests := []struct {
// 		name    string
// 		q       *SearchItems
// 		args    args
// 		want    *SearchItems
// 		wantErr bool
// 	}{
// 		{
// 			"TestSearchItems_Search #1",
// 			newStandardSearchItem(),
// 			args{
// 				"Tom Cruise",
// 				Actor,
// 			},
// 			finalTestStruct,
// 			false,
// 		},
// 		{
// 			"TestSearchItems_Search #2",
// 			newStandardSearchItem(),
// 			args{
// 				"Muse",
// 				Artist,
// 			},
// 			finalTestStruct,
// 			true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.q.Search(tt.args.searchType, tt.args.searchParam); !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
// 				t.Errorf("SearchItems.Search() = %[1]v (%[1]T), want %[2]v (%[2]T)", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestSearchItems_WithFilters(t *testing.T) {
// 	originStruct := newStandardSearchItem()
// 	finalTestStruct := newStandardSearchItem()
// 	type args struct {
// 		filters []map[RequestFilter]interface{}
// 	}
//
// 	originStruct.Keywords = "Sony PlayStation 4 Pro"
// 	finalTestStruct.Keywords = "Sony PlayStation 4 Pro"
// 	finalTestStruct.Condition = "New"
// 	finalTestStruct.DeliveryFlags = []string{"Prime"}
//
// 	tests := []struct {
// 		name string
// 		q    *SearchItems
// 		args args
// 		want *SearchItems
// 	}{
// 		{
// 			"TestSearchItems_WithFilters",
// 			originStruct,
// 			args{
// 				[]map[RequestFilter]interface{}{
// 					{
// 						Condition: "New",
// 					},
// 					{
// 						DeliveryFlags: "Prime",
// 					},
// 				},
// 			},
// 			finalTestStruct,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.q.WithFilters(tt.args.filters...); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("SearchItems.WithFilters() = %[1]v (%[1]T), want %[2]v (%[2]T)", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestSearchItems_mapFilter(t *testing.T) {
// 	inputStruct := newStandardSearchItem()
//
// 	type args struct {
// 		filter      RequestFilter
// 		filterValue interface{}
// 	}
// 	tests := []struct {
// 		name    string
// 		q       *SearchItems
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			"TestSearchItems_mapFilter #1",
// 			inputStruct,
// 			args{
// 				BrowseNodeID,
// 				"290060",
// 			},
// 			false,
// 		},
// 		{
// 			"TestSearchItems_mapFilter #2",
// 			inputStruct,
// 			args{
// 				BrowseNodeID,
// 				"XXX844857",
// 			},
// 			true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.q.mapFilter(tt.args.filter, tt.args.filterValue)
// 			if tt.q.BrowseNodeID != tt.args.filterValue.(string) && !tt.wantErr {
// 				t.Errorf("SearchItems.WithFilters() = %[1]v (%[1]T), want %[2]v (%[2]T)", tt.q.BrowseNodeID, tt.args.filterValue)
// 			}
// 		})
// 	}
// }
//
// func TestSearchItems_EnableBrowseNodeInfo(t *testing.T) {
// 	type args struct {
// 		flag bool
// 	}
// 	tests := []struct {
// 		name string
// 		q    *SearchItems
// 		args args
// 		want *SearchItems
// 	}{
// 		{
// 			"TestSearchItems_EnableBrowseNodeInfo #1",
// 			newStandardSearchItem(),
// 			args{
// 				false,
// 			},
// 			newStandardSearchItem().EnableBrowseNodeInfo(false),
// 		},
// 		{
// 			"TestSearchItems_EnableBrowseNodeInfo #2",
// 			newStandardSearchItem(),
// 			args{
// 				true,
// 			},
// 			newStandardSearchItem().EnableBrowseNodeInfo(true),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.q.EnableBrowseNodeInfo(tt.args.flag); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("SearchItems.EnableBrowseNodeInfo() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestSearchItems_EnableImages(t *testing.T) {
// 	type args struct {
// 		flag bool
// 	}
// 	tests := []struct {
// 		name string
// 		q    *SearchItems
// 		args args
// 		want *SearchItems
// 	}{
// 		{
// 			"TestSearchItems_EnableImages #1",
// 			newStandardSearchItem(),
// 			args{
// 				false,
// 			},
// 			newStandardSearchItem().EnableImages(false),
// 		},
// 		{
// 			"TestSearchItems_EnableImages #2",
// 			newStandardSearchItem(),
// 			args{
// 				true,
// 			},
// 			newStandardSearchItem().EnableImages(true),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.q.EnableImages(tt.args.flag); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("SearchItems.EnableImages() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestSearchItems_EnableItemInfo(t *testing.T) {
// 	type args struct {
// 		flag bool
// 	}
// 	tests := []struct {
// 		name string
// 		q    *SearchItems
// 		args args
// 		want *SearchItems
// 	}{
// 		{
// 			"TestSearchItems_EnableItemInfo #1",
// 			newStandardSearchItem(),
// 			args{
// 				false,
// 			},
// 			newStandardSearchItem().EnableItemInfo(false),
// 		},
// 		{
// 			"TestSearchItems_EnableItemInfo #2",
// 			newStandardSearchItem(),
// 			args{
// 				true,
// 			},
// 			newStandardSearchItem(),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.q.EnableItemInfo(tt.args.flag); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("SearchItems.EnableItemInfo() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestSearchItems_EnableOffers(t *testing.T) {
// 	type args struct {
// 		flag bool
// 	}
// 	tests := []struct {
// 		name string
// 		q    *SearchItems
// 		args args
// 		want *SearchItems
// 	}{
// 		{
// 			"TestSearchItems_EnableOffers #1",
// 			newStandardSearchItem(),
// 			args{
// 				false,
// 			},
// 			newStandardSearchItem().EnableOffers(false),
// 		},
// 		{
// 			"TestSearchItems_EnableOffers #2",
// 			newStandardSearchItem(),
// 			args{
// 				true,
// 			},
// 			newStandardSearchItem().EnableOffers(true),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.q.EnableOffers(tt.args.flag); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("SearchItems.EnableOffers() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestSearchItems_EnableSearchRefinements(t *testing.T) {
// 	type args struct {
// 		flag bool
// 	}
// 	tests := []struct {
// 		name string
// 		q    *SearchItems
// 		args args
// 		want *SearchItems
// 	}{
// 		{
// 			"TestSearchItems_EnableSearchRefinements #1",
// 			newStandardSearchItem(),
// 			args{
// 				false,
// 			},
// 			newStandardSearchItem().EnableSearchRefinements(false),
// 		},
// 		{
// 			"TestSearchItems_EnableSearchRefinements #2",
// 			newStandardSearchItem(),
// 			args{
// 				true,
// 			},
// 			newStandardSearchItem().EnableSearchRefinements(true),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.q.EnableSearchRefinements(tt.args.flag); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("SearchItems.EnableSearchRefinements() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestSearchItems_EnableParentASIN(t *testing.T) {
// 	type args struct {
// 		flag bool
// 	}
// 	tests := []struct {
// 		name string
// 		q    *SearchItems
// 		args args
// 		want *SearchItems
// 	}{
// 		{
// 			"TestSearchItems_EnableParentASIN #1",
// 			newStandardSearchItem(),
// 			args{
// 				false,
// 			},
// 			newStandardSearchItem().EnableParentASIN(false),
// 		},
// 		{
// 			"TestSearchItems_EnableParentASIN #2",
// 			newStandardSearchItem(),
// 			args{
// 				true,
// 			},
// 			newStandardSearchItem().EnableParentASIN(true),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.q.EnableParentASIN(tt.args.flag); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("SearchItems.EnableParentASIN() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestSearchItems_Operation(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		q       *SearchItems
// 		want    paapi5.Operation
// 		wantErr bool
// 	}{
// 		{
// 			"TestSearchItems_Operation #1",
// 			newStandardSearchItem(),
// 			paapi5.SearchItems,
// 			false,
// 		},
// 		{
// 			"TestSearchItems_Operation #2",
// 			newStandardSearchItem(),
// 			paapi5.GetItems,
// 			true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.q.Operation(); !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
// 				t.Errorf("SearchItems.Operation() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestSearchItems_Payload(t *testing.T) {
// 	originSearchItem := newStandardSearchItem().EnableItemInfo(false)
// 	byteResult, _ := json.Marshal(originSearchItem)
//
// 	tests := []struct {
// 		name    string
// 		q       *SearchItems
// 		want    []byte
// 		wantErr bool
// 	}{
// 		{
// 			"TestSearchItems_Payload #1",
// 			&SearchItems{nil},
// 			nil,
// 			true,
// 		},
// 		{
// 			"TestSearchItems_Payload #2",
// 			newStandardSearchItem().EnableItemInfo(false),
// 			byteResult,
// 			false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.q.Payload()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("SearchItems.Payload() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("SearchItems.Payload() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func TestSearchItems_String(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		q    *SearchItems
// 		want string
// 	}{
// 		{
// 			"TestSearchItems_String",
// 			newStandardSearchItem(),
// 			newStandardSearchItem().String(),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.q.String(); got != tt.want {
// 				t.Errorf("SearchItems.String() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func Test_isFilterParamValid(t *testing.T) {
// 	type args struct {
// 		param  string
// 		params []string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		{
// 			"Test_isFilterParamValid #1",
// 			args{
// 				"Prime",
// 				validDeliveryParameters,
// 			},
// 			true,
// 		},
// 		{
// 			"Test_isFilterParamValid #2",
// 			args{
// 				"Pigeon",
// 				validDeliveryParameters,
// 			},
// 			false,
// 		},
// 		{
// 			"Test_isFilterParamValid #3",
// 			args{
// 				"Amazon",
// 				validMerchantParameters,
// 			},
// 			true,
// 		},
// 		{
// 			"Test_isFilterParamValid #4",
// 			args{
// 				"AvgCustomerReviews",
// 				validAvailabilityParameters,
// 			},
// 			false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := isFilterParamValid(tt.args.param, tt.args.params); got != tt.want {
// 				t.Errorf("isFilterParamValid() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func Test_stringIsNumber(t *testing.T) {
// 	type args struct {
// 		s string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		{
// 			"Test_stringIsNumber #1",
// 			args{
// 				"0123456789",
// 			},
// 			true,
// 		},
// 		{
// 			"Test_stringIsNumber #2",
// 			args{
// 				"01234S6T89",
// 			},
// 			false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := stringIsNumber(tt.args.s); got != tt.want {
// 				t.Errorf("stringIsNumber(%v) = %v, want %v", tt.args.s, got, tt.want)
// 			}
// 		})
// 	}
// }
//
// // helper function to instantiate a standard SearchItem with the ItemInfo resource
// func newStandardSearchItem() *SearchItems {
// 	return &SearchItems{
// 		Query: &Query{
// 			OpeCode: paapi5.SearchItems,
// 			request: request{
// 				Marketplace: "",
// 				PartnerTag:  "",
// 				PartnerType: "",
// 			},
// 			enableResources: map[resource]bool{},
// 		},
// 	}
// }

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
