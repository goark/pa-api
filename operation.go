package paapi5

import (
	"path"
	"strconv"
	"strings"
)

type Operation int

const (
	NullOperation Operation = iota
	GetVariations
	GetItems
	SearchItems
)

var nameMap = map[Operation]string{
	GetVariations: "GetVariations",
	GetItems:      "GetItems",
	SearchItems:   "SearchItems",
}

//Stringer interface
func (c Operation) String() string {
	if s, ok := nameMap[c]; ok {
		return s
	}
	return ""
}

//Path returns URL path for PA-API command
func (c Operation) Path() string {
	cmd := c.String()
	if len(cmd) == 0 {
		return ""
	}
	return path.Join("/paapi5", strings.ToLower(cmd))
}

//Target returns taget name for PA-API command
func (c Operation) Target() string {
	cmd := c.String()
	if len(cmd) == 0 {
		return ""
	}
	return strings.Join([]string{"com.amazon.paapi5.v1.ProductAdvertisingAPIv1", cmd}, ".")
}

//UnmarshalJSON returns result of Unmarshal for json.Unmarshal()
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

//MarshalJSON returns time string with RFC3339 format
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
