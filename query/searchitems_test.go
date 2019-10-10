package query

import (
	"encoding/json"
	"reflect"
	"testing"

	paapi5 "github.com/spiegel-im-spiegel/pa-api"
)

func TestNewSearchItems(t *testing.T) {
	type args struct {
		marketplace string
		partnerTag  string
		partnerType string
	}
	tests := []struct {
		name string
		args args
		want *SearchItems
	}{
		{
			"TestNewSearchItems",
			args{
				"www.amazon.co.jp",
				"wwwyourpartnertag-20",
				"Associate",
			},
			&SearchItems{
				OpeCode:        paapi5.SearchItems,
				Marketplace:    "www.amazon.co.jp",
				PartnerTag:     "wwwyourpartnertag-20",
				PartnerType:    "Associate",
				enableItemInfo: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSearchItems(tt.args.marketplace, tt.args.partnerTag, tt.args.partnerType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSearchItems() = %[1]v (%[1]T), want %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func Test_newNilSearchItems(t *testing.T) {
	tests := []struct {
		name    string
		want    *SearchItems
		wantErr bool
	}{
		{
			"Test_newNilSearchItems #1",
			newStandardSearchItem(),
			false,
		},
		{
			"Test_newNilSearchItems #2",
			&SearchItems{
				OpeCode:     paapi5.SearchItems,
				Marketplace: "",
				PartnerTag:  "",
				PartnerType: "",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newNilSearchItems(); !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("newNilSearchItems() = %[1]v (%[1]T), want %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func TestSearchItems_Search(t *testing.T) {
	finalTestStruct := newStandardSearchItem()
	finalTestStruct.Actor = "Tom Cruise"

	type args struct {
		searchParam string
		searchType  RequestType
	}
	tests := []struct {
		name    string
		q       *SearchItems
		args    args
		want    *SearchItems
		wantErr bool
	}{
		{
			"TestSearchItems_Search #1",
			newStandardSearchItem(),
			args{
				"Tom Cruise",
				Actor,
			},
			finalTestStruct,
			false,
		},
		{
			"TestSearchItems_Search #2",
			newStandardSearchItem(),
			args{
				"Muse",
				Artist,
			},
			finalTestStruct,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Search(tt.args.searchParam, tt.args.searchType); !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("SearchItems.Search() = %[1]v (%[1]T), want %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func TestSearchItems_WithFilters(t *testing.T) {
	originStruct := newStandardSearchItem()
	finalTestStruct := newStandardSearchItem()
	type args struct {
		filters []map[RequestFilter]interface{}
	}

	originStruct.Keywords = "Sony PlayStation 4 Pro"
	finalTestStruct.Keywords = "Sony PlayStation 4 Pro"
	finalTestStruct.Condition = "New"
	finalTestStruct.DeliveryFlags = []string{"Prime"}

	tests := []struct {
		name string
		q    *SearchItems
		args args
		want *SearchItems
	}{
		{
			"TestSearchItems_WithFilters",
			originStruct,
			args{
				[]map[RequestFilter]interface{}{
					{
						Condition: "New",
					},
					{
						DeliveryFlags: "Prime",
					},
				},
			},
			finalTestStruct,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.WithFilters(tt.args.filters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchItems.WithFilters() = %[1]v (%[1]T), want %[2]v (%[2]T)", got, tt.want)
			}
		})
	}
}

func TestSearchItems_mapFilter(t *testing.T) {
	inputStruct := newStandardSearchItem()

	type args struct {
		filter      RequestFilter
		filterValue interface{}
	}
	tests := []struct {
		name    string
		q       *SearchItems
		args    args
		wantErr bool
	}{
		{
			"TestSearchItems_mapFilter #1",
			inputStruct,
			args{
				BrowseNodeID,
				"290060",
			},
			false,
		},
		{
			"TestSearchItems_mapFilter #2",
			inputStruct,
			args{
				BrowseNodeID,
				"XXX844857",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.mapFilter(tt.args.filter, tt.args.filterValue)
			if tt.q.BrowseNodeID != tt.args.filterValue.(string) && !tt.wantErr {
				t.Errorf("SearchItems.WithFilters() = %[1]v (%[1]T), want %[2]v (%[2]T)", tt.q.BrowseNodeID, tt.args.filterValue)
			}
		})
	}
}

func TestSearchItems_EnableBrowseNodeInfo(t *testing.T) {
	type args struct {
		flag bool
	}
	tests := []struct {
		name string
		q    *SearchItems
		args args
		want *SearchItems
	}{
		{
			"TestSearchItems_EnableBrowseNodeInfo #1",
			newStandardSearchItem(),
			args{
				false,
			},
			newStandardSearchItem().EnableBrowseNodeInfo(false),
		},
		{
			"TestSearchItems_EnableBrowseNodeInfo #2",
			newStandardSearchItem(),
			args{
				true,
			},
			newStandardSearchItem().EnableBrowseNodeInfo(true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.EnableBrowseNodeInfo(tt.args.flag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchItems.EnableBrowseNodeInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchItems_EnableImages(t *testing.T) {
	type args struct {
		flag bool
	}
	tests := []struct {
		name string
		q    *SearchItems
		args args
		want *SearchItems
	}{
		{
			"TestSearchItems_EnableImages #1",
			newStandardSearchItem(),
			args{
				false,
			},
			newStandardSearchItem().EnableImages(false),
		},
		{
			"TestSearchItems_EnableImages #2",
			newStandardSearchItem(),
			args{
				true,
			},
			newStandardSearchItem().EnableImages(true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.EnableImages(tt.args.flag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchItems.EnableImages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchItems_EnableItemInfo(t *testing.T) {
	type args struct {
		flag bool
	}
	tests := []struct {
		name string
		q    *SearchItems
		args args
		want *SearchItems
	}{
		{
			"TestSearchItems_EnableItemInfo #1",
			newStandardSearchItem(),
			args{
				false,
			},
			newStandardSearchItem().EnableItemInfo(false),
		},
		{
			"TestSearchItems_EnableItemInfo #2",
			newStandardSearchItem(),
			args{
				true,
			},
			newStandardSearchItem(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.EnableItemInfo(tt.args.flag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchItems.EnableItemInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchItems_EnableOffers(t *testing.T) {
	type args struct {
		flag bool
	}
	tests := []struct {
		name string
		q    *SearchItems
		args args
		want *SearchItems
	}{
		{
			"TestSearchItems_EnableOffers #1",
			newStandardSearchItem(),
			args{
				false,
			},
			newStandardSearchItem().EnableOffers(false),
		},
		{
			"TestSearchItems_EnableOffers #2",
			newStandardSearchItem(),
			args{
				true,
			},
			newStandardSearchItem().EnableOffers(true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.EnableOffers(tt.args.flag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchItems.EnableOffers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchItems_EnableSearchRefinements(t *testing.T) {
	type args struct {
		flag bool
	}
	tests := []struct {
		name string
		q    *SearchItems
		args args
		want *SearchItems
	}{
		{
			"TestSearchItems_EnableSearchRefinements #1",
			newStandardSearchItem(),
			args{
				false,
			},
			newStandardSearchItem().EnableSearchRefinements(false),
		},
		{
			"TestSearchItems_EnableSearchRefinements #2",
			newStandardSearchItem(),
			args{
				true,
			},
			newStandardSearchItem().EnableSearchRefinements(true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.EnableSearchRefinements(tt.args.flag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchItems.EnableSearchRefinements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchItems_EnableParentASIN(t *testing.T) {
	type args struct {
		flag bool
	}
	tests := []struct {
		name string
		q    *SearchItems
		args args
		want *SearchItems
	}{
		{
			"TestSearchItems_EnableParentASIN #1",
			newStandardSearchItem(),
			args{
				false,
			},
			newStandardSearchItem().EnableParentASIN(false),
		},
		{
			"TestSearchItems_EnableParentASIN #2",
			newStandardSearchItem(),
			args{
				true,
			},
			newStandardSearchItem().EnableParentASIN(true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.EnableParentASIN(tt.args.flag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchItems.EnableParentASIN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchItems_Operation(t *testing.T) {
	tests := []struct {
		name    string
		q       *SearchItems
		want    paapi5.Operation
		wantErr bool
	}{
		{
			"TestSearchItems_Operation #1",
			newStandardSearchItem(),
			paapi5.SearchItems,
			false,
		},
		{
			"TestSearchItems_Operation #2",
			newStandardSearchItem(),
			paapi5.GetItems,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Operation(); !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("SearchItems.Operation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchItems_Payload(t *testing.T) {
	originSearchItem := newStandardSearchItem().EnableItemInfo(false)
	byteResult, _ := json.Marshal(originSearchItem)

	tests := []struct {
		name    string
		q       *SearchItems
		want    []byte
		wantErr bool
	}{
		{
			"TestSearchItems_Payload #1",
			nil,
			nil,
			true,
		},
		{
			"TestSearchItems_Payload #2",
			newStandardSearchItem().EnableItemInfo(false),
			byteResult,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.Payload()
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchItems.Payload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchItems.Payload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchItems_String(t *testing.T) {
	tests := []struct {
		name string
		q    *SearchItems
		want string
	}{
		{
			"TestSearchItems_String",
			newStandardSearchItem(),
			newStandardSearchItem().String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.String(); got != tt.want {
				t.Errorf("SearchItems.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isFilterParamValid(t *testing.T) {
	type args struct {
		param  string
		params []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test_isFilterParamValid #1",
			args{
				"Prime",
				validDeliveryParameters,
			},
			true,
		},
		{
			"Test_isFilterParamValid #2",
			args{
				"Pigeon",
				validDeliveryParameters,
			},
			false,
		},
		{
			"Test_isFilterParamValid #3",
			args{
				"Amazon",
				validMerchantParameters,
			},
			true,
		},
		{
			"Test_isFilterParamValid #4",
			args{
				"AvgCustomerReviews",
				validAvailabilityParameters,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFilterParamValid(tt.args.param, tt.args.params); got != tt.want {
				t.Errorf("isFilterParamValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringIsNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test_stringIsNumber #1",
			args{
				"0123456789",
			},
			true,
		},
		{
			"Test_stringIsNumber #2",
			args{
				"01234S6T89",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringIsNumber(tt.args.s); got != tt.want {
				t.Errorf("stringIsNumber(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}

// helper function to instantiate a standard SearchItem with the ItemInfo resource
func newStandardSearchItem() *SearchItems {
	return &SearchItems{
		OpeCode:        paapi5.SearchItems,
		Marketplace:    "",
		PartnerTag:     "",
		PartnerType:    "",
		enableItemInfo: true,
	}
}
