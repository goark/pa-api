package query

import (
	"fmt"
	"testing"
)

func TestQueryNil(t *testing.T) {
	q := (*Query)(nil)
	if _, err := q.JSON(); err == nil {
		t.Errorf("(*Query)(nil) is nil, not want nil")
	} else {
		fmt.Printf("Info: %+v\n", err)
	}
	res := `{"Operation":"","Marketplace":"","PartnerTag":"","PartnerType":"","Resources":["ItemInfo.ByLineInfo","ItemInfo.ContentInfo","ItemInfo.ContentRating","ItemInfo.Classifications","ItemInfo.ExternalIds","ItemInfo.Features","ItemInfo.ManufactureInfo","ItemInfo.ProductInfo","ItemInfo.TechnicalInfo","ItemInfo.Title","ItemInfo.TradeInInfo"]}`
	q = newNil()
	b, err := q.JSON()
	if err != nil {
		t.Errorf("nil-Query is \"%v\", want nil", err)
		fmt.Printf("Info: %+v\n", err)
	} else if string(b) != res {
		t.Errorf("nil-Query = \"%v\", want \"%v\"", string(b), res)
	}
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
