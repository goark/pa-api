package entity

import "testing"

func TestDecodeResponseItemResultsOffersV2AndBrowseNodeInfo(t *testing.T) {
	body := []byte(`{
  "itemResults": {
    "items": [
      {
        "asin": "A1",
        "images": {
          "primary": {
            "small": {"url": "https://example/s.jpg", "height": 75, "width": 75},
            "large": {"url": "https://example/l.jpg", "height": 500, "width": 500}
          },
          "variants": [
            {
              "small": {"url": "https://example/vs.jpg", "height": 75, "width": 75},
              "large": {"url": "https://example/vl.jpg", "height": 500, "width": 500}
            }
          ]
        },
        "browseNodeInfo": {
          "browseNodes": [
            {
              "id": "123",
              "displayName": "Books",
              "contextFreeName": "Books",
              "websiteSalesRank": {
                "displayName": "Books",
                "contextFreeName": "Books",
                "salesRank": 7
              }
            }
          ]
        },
        "offersV2": {
          "listings": [
            {
              "type": "New",
              "violatesMAP": true
            }
          ]
        }
      }
    ]
  }
}`)

	resp, err := DecodeResponse(body)
	if err != nil {
		t.Fatalf("DecodeResponse: %+v", err)
	}
	if resp.ItemsResult == nil {
		t.Fatal("ItemsResult is nil")
	}
	if got, want := len(resp.ItemsResult.Items), 1; got != want {
		t.Fatalf("len(ItemsResult.Items) = %d, want %d", got, want)
	}

	item := resp.ItemsResult.Items[0]
	if item.Images == nil || item.Images.Primary == nil || item.Images.Primary.Large == nil {
		t.Fatal("Images.Primary.Large is nil")
	}
	if got, want := item.Images.Primary.Large.URL, "https://example/l.jpg"; got != want {
		t.Errorf("Images.Primary.Large.URL = %q, want %q", got, want)
	}
	if item.Images.Variants == nil || len(item.Images.Variants) != 1 || item.Images.Variants[0].Large == nil {
		t.Fatal("Images.Variants[0].Large is nil")
	}
	if got, want := item.Images.Variants[0].Large.URL, "https://example/vl.jpg"; got != want {
		t.Errorf("Images.Variants[0].Large.URL = %q, want %q", got, want)
	}

	if item.BrowseNodeInfo == nil || len(item.BrowseNodeInfo.BrowseNodes) != 1 || item.BrowseNodeInfo.BrowseNodes[0].WebsiteSalesRank == nil {
		t.Fatal("BrowseNodeInfo.BrowseNodes[0].WebsiteSalesRank is nil")
	}
	if got, want := item.BrowseNodeInfo.BrowseNodes[0].WebsiteSalesRank.SalesRank, 7; got != want {
		t.Errorf("WebsiteSalesRank.SalesRank = %d, want %d", got, want)
	}

	if item.OffersV2 == nil || item.OffersV2.Listings == nil || len(*item.OffersV2.Listings) != 1 {
		t.Fatal("OffersV2.Listings is empty")
	}
	if !(*item.OffersV2.Listings)[0].ViolatesMAP {
		t.Errorf("ViolatesMAP = false, want true")
	}
}

func TestDecodeResponseVariationSummary(t *testing.T) {
	body := []byte(`{
  "variationsResult": {
    "variationSummary": {
      "pageCount": 2,
      "variationCount": 13,
      "price": {
        "highestPrice": {"amount": 30.87, "currency": "GBP", "displayAmount": "GBP 30.87"},
        "lowestPrice": {"amount": 17.03, "currency": "GBP", "displayAmount": "GBP 17.03"}
      },
      "variationDimensions": [
        {
          "displayName": "Color",
          "name": "color_name",
          "values": ["Red", "Blue"]
        }
      ]
    }
  }
}`)

	resp, err := DecodeResponse(body)
	if err != nil {
		t.Fatalf("DecodeResponse: %+v", err)
	}
	if resp.VariationsResult == nil || resp.VariationsResult.VariationSummary == nil {
		t.Fatal("VariationsResult.VariationSummary is nil")
	}
	vs := resp.VariationsResult.VariationSummary
	if got, want := vs.PageCount, 2; got != want {
		t.Errorf("PageCount = %d, want %d", got, want)
	}
	if got, want := vs.VariationCount, 13; got != want {
		t.Errorf("VariationCount = %d, want %d", got, want)
	}
	if vs.Price == nil || vs.Price.HighestPrice == nil || vs.Price.LowestPrice == nil {
		t.Fatal("VariationSummary.Price is nil")
	}
	if got, want := vs.Price.HighestPrice.Currency, "GBP"; got != want {
		t.Errorf("HighestPrice.Currency = %q, want %q", got, want)
	}
	if got, want := len(vs.VariationDimensions), 1; got != want {
		t.Fatalf("len(VariationDimensions) = %d, want %d", got, want)
	}
	if got, want := vs.VariationDimensions[0].Name, "color_name"; got != want {
		t.Errorf("VariationDimensions[0].Name = %q, want %q", got, want)
	}
}
