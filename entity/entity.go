package entity

import (
	"bytes"
	"encoding/json"

	"github.com/spiegel-im-spiegel/errs"
)

type Image struct {
	URL    string
	Height int
	Width  int
}

type GenInfo struct {
	DisplayValue string
	Label        string `json:",omitempty"`
	Locale       string `json:",omitempty"`
}

type GenInfoInt struct {
	DisplayValue int
	Label        string `json:",omitempty"`
	Locale       string `json:",omitempty"`
}

type GenInfoFloat struct {
	DisplayValue float64
	Label        string `json:",omitempty"`
	Locale       string `json:",omitempty"`
	Unit         string `json:",omitempty"`
}

type GenInfoTime struct {
	DisplayValue Date
	Label        string `json:",omitempty"`
	Locale       string `json:",omitempty"`
}

type IdInfo struct {
	DisplayValues []string
	Label         string `json:",omitempty"`
	Locale        string `json:",omitempty"`
}

type Ancestor struct {
	Id              string
	DisplayName     string
	ContextFreeName string
	Ancestor        *Ancestor `json:",omitempty"`
}

type Item struct {
	ASIN           string
	DetailPageURL  string
	BrowseNodeInfo *struct {
		BrowseNodes []struct {
			Id               string
			DisplayName      string
			ContextFreeName  string
			IsRoot           bool
			SalesRank        *int      `json:",omitempty"`
			Ancestor         *Ancestor `json:",omitempty"`
			WebsiteSalesRank *struct {
				DisplayName     string
				ContextFreeName string
				SalesRank       int
			} `json:",omitempty"`
		} `json:",omitempty"`
	} `json:",omitempty"`
	Images *struct {
		Primary *struct {
			Large  *Image `json:",omitempty"`
			Medium *Image `json:",omitempty"`
			Small  *Image `json:",omitempty"`
		} `json:",omitempty"`
		Variants []*struct {
			Large  *Image `json:",omitempty"`
			Medium *Image `json:",omitempty"`
			Small  *Image `json:",omitempty"`
		} `json:",omitempty"`
	} `json:",omitempty"`
	ItemInfo *struct {
		ByLineInfo *struct {
			Brand        *GenInfo `json:",omitempty"`
			Manufacturer *GenInfo `json:",omitempty"`
			Contributors []struct {
				Name   string
				Locale string
				Role   string
			}
		} `json:",omitempty"`
		Classifications *struct {
			Binding      GenInfo
			ProductGroup GenInfo
		} `json:",omitempty"`
		ContentInfo *struct {
			Edition   *GenInfo `json:",omitempty"`
			Languages struct {
				DisplayValues []struct {
					DisplayValue string
					Type         string
				}
				Label  string
				Locale string
			}
			PagesCount struct {
				DisplayValue int
				Label        string
				Locale       string
			}
			PublicationDate GenInfoTime
		} `json:",omitempty"`
		ContentRating *struct {
			AudienceRating GenInfo
		} `json:",omitempty"`
		ExternalIds *struct {
			EANs  *IdInfo `json:",omitempty"`
			ISBNs *IdInfo `json:",omitempty"`
			UPCs  *IdInfo `json:",omitempty"`
		} `json:",omitempty"`
		Features        *IdInfo `json:",omitempty"`
		ManufactureInfo *struct {
			ItemPartNumber *GenInfo `json:",omitempty"`
			Model          *GenInfo `json:",omitempty"`
			Warranty       *GenInfo `json:",omitempty"`
		} `json:",omitempty"`
		ProductInfo *struct {
			Color          *GenInfo `json:",omitempty"`
			IsAdultProduct struct {
				DisplayValue bool
				Label        string
				Locale       string
			}
			ItemDimensions *struct {
				Height *GenInfoFloat `json:",omitempty"`
				Length *GenInfoFloat `json:",omitempty"`
				Weight *GenInfoFloat `json:",omitempty"`
				Width  *GenInfoFloat `json:",omitempty"`
			} `json:",omitempty"`
			ReleaseDate *GenInfoTime `json:",omitempty"`
			Size        *GenInfo     `json:",omitempty"`
			UnitCount   *GenInfoInt  `json:",omitempty"`
		} `json:",omitempty"`
		TechnicalInfo *struct {
			Formats IdInfo
		} `json:",omitempty"`
		Title       *GenInfo `json:",omitempty"`
		TradeInInfo *struct {
			IsEligibleForTradeIn bool
			Price                struct {
				DisplayAmount string
				Amount        float64
				Currency      string
			}
		} `json:",omitempty"`
	}
}

type Refinement struct {
	Id          string
	DisplayName string
	Bins        []struct {
		Id          string
		DisplayName string
	} `json:",omitempty"`
}

type Price struct {
	DisplayAmount string
	Amount        float64
	Currency      string
}

type VariationDimension struct {
	DisplayName string
	Name        string
	Values      []string
}

type Response struct {
	Errors []struct {
		Code    string
		Message string
	} `json:",omitempty"`
	ItemsResult *struct {
		Items []Item `json:",omitempty"`
	} `json:",omitempty"`
	SearchResult *struct {
		Items             []Item `json:",omitempty"`
		SearchRefinements *struct {
			SearchIndex      *Refinement  `json:",omitempty"`
			BrowseNode       *Refinement  `json:",omitempty"`
			OtherRefinements []Refinement `json:",omitempty"`
		} `json:",omitempty"`
		SearchURL        string
		TotalResultCount int
	} `json:",omitempty"`
	VariationSummary *struct {
		PageCount      int
		VariationCount int
		Price          *struct {
			HighestPrice *Price `json:",omitempty"`
			LowestPrice  *Price `json:",omitempty"`
		} `json:",omitempty"`
		VariationDimensions []VariationDimension `json:",omitempty"`
	} `json:",omitempty"`
}

//DecodeResponse returns array of Response instance from byte buffer
func DecodeResponse(b []byte) (*Response, error) {
	rsp := Response{}
	if err := json.NewDecoder(bytes.NewReader(b)).Decode(&rsp); err != nil {
		return &rsp, errs.Wrap(err, "", errs.WithContext("JSON", string(b)))
	}
	return &rsp, nil
}

//JSON returns JSON data from Response instance
func (r *Response) JSON() ([]byte, error) {
	b, err := json.Marshal(r)
	return b, errs.Wrap(err, "")
}

//Stringer
func (r *Response) String() string {
	b, err := r.JSON()
	if err != nil {
		return ""
	}
	return string(b)
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
