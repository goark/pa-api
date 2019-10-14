package paapi5

import (
	"encoding/json"
	"path"
	"strconv"
	"strings"
)

//Operation is enumeration of PA-API operation
type Operation int

var _ json.Marshaler = Operation(0)        //Operation type is compatible with json.Marshaler interface
var _ json.Unmarshaler = (*Operation)(nil) //Operation type is compatible with json.Unmarshaler interface

const (
	NullOperation Operation = iota //Unknown
	GetVariations                  //GetVariations
	GetItems                       //GetItems
	SearchItems                    //SearchItems
)

var nameMap = map[Operation]string{
	GetVariations: "GetVariations",
	GetItems:      "GetItems",
	SearchItems:   "SearchItems",
}

//String method is a implementation of fmt.Stringer interface.
func (c Operation) String() string {
	if s, ok := nameMap[c]; ok {
		return s
	}
	return ""
}

//Path method returns URL path of PA-API operation
func (c Operation) Path() string {
	cmd := c.String()
	if len(cmd) == 0 {
		return ""
	}
	return path.Join("/paapi5", strings.ToLower(cmd))
}

//Target method returns taget name of PA-API operation
func (c Operation) Target() string {
	cmd := c.String()
	if len(cmd) == 0 {
		return ""
	}
	return strings.Join([]string{"com.amazon.paapi5.v1.ProductAdvertisingAPIv1", cmd}, ".")
}

//UnmarshalJSON method implements json.Unmarshaler interface.
func (c *Operation) UnmarshalJSON(b []byte) error {
	s := string(b)
	if ss, err := strconv.Unquote(s); err == nil {
		s = ss
	}
	for k, v := range nameMap {
		if s == v {
			*c = k
		}
	}
	return nil
}

//MarshalJSON method implements the json.Marshaler interface.
func (c Operation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(c.String())), nil
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
