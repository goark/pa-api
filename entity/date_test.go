package entity

import (
	"encoding/json"
	"testing"
)

type ForTestStruct struct {
	DateTaken Date `json:"date_taken,omitempty"`
}

func TestUnmarshal(t *testing.T) {
	testCases := []struct {
		s   string
		jsn string
	}{
		{s: `{"date_taken": "2005-03-26T15:04:05Z"}`, jsn: `{"date_taken":"2005-03-26T15:04:05Z"}`},
		{s: `{"date_taken": "2005-03T"}`, jsn: `{"date_taken":"2005-03-01T00:00:00Z"}`},
		{s: `{"date_taken": "2005T"}`, jsn: `{"date_taken":"2005-01-01T00:00:00Z"}`},
		{s: `{"date_taken": "2005-03-26"}`, jsn: `{"date_taken":"2005-03-26T00:00:00Z"}`},
		{s: `{"date_taken": "2005-03"}`, jsn: `{"date_taken":"2005-03-01T00:00:00Z"}`},
		{s: `{"date_taken": "2005"}`, jsn: `{"date_taken":"2005-01-01T00:00:00Z"}`},
		{s: `{"date_taken": "2005/03/26"}`, jsn: `{"date_taken":"2005-03-26T00:00:00Z"}`},
		{s: `{"date_taken": "2005/03"}`, jsn: `{"date_taken":"2005-03-01T00:00:00Z"}`},
		{s: `{}`, jsn: `{"date_taken":"0001-01-01T00:00:00Z"}`},
	}

	for _, tc := range testCases {
		tst := &ForTestStruct{}
		if err := json.Unmarshal([]byte(tc.s), tst); err != nil {
			t.Errorf("json.Unmarshal() is \"%v\", want nil.", err)
			continue
		}
		b, err := json.Marshal(tst)
		if err != nil {
			t.Errorf("json.Marshal() is \"%v\", want nil.", err)
			continue
		}
		str := string(b)
		if str != tc.jsn {
			t.Errorf("ForTestStruct = \"%v\", want \"%v\".", str, tc.jsn)
		}
	}
}

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
