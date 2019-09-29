package paapi5

import (
	"testing"
	"time"
)

func TestTimeStamp(t *testing.T) {
	testCases := []struct {
		tm      time.Time
		str     string
		strDate string
	}{
		{tm: time.Time{}, str: "00010101T000000Z", strDate: "00010101"},
		{tm: time.Date(2019, time.September, 30, 8, 31, 54, 0, time.UTC), str: "20190930T083154Z", strDate: "20190930"},
		{tm: time.Date(2019, time.September, 30, 8, 31, 54, 0, time.FixedZone("JST", 9*60*60)), str: "20190929T233154Z", strDate: "20190929"},
	}
	for _, tc := range testCases {
		dt := NewTimeStamp(tc.tm)
		str := dt.String()
		if str != tc.str {
			t.Errorf("TimeStamp.String() = \"%v\", want \"%v\".", str, tc.str)
		}
		strDate := dt.StringDate()
		if strDate != tc.strDate {
			t.Errorf("TimeStamp.StringDate() = \"%v\", want \"%v\".", strDate, tc.strDate)
		}
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
