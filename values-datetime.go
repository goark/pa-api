package paapi5

import "time"

//TimeStamp is wrapper class of time.Time
type TimeStamp struct {
	time.Time
}

//NewTimeStamp returns TimeStamp instance
func NewTimeStamp(tm time.Time) TimeStamp {
	return TimeStamp{tm}
}

func (t TimeStamp) StringDate() string {
	return t.UTC().Format("20060102")
}

func (t TimeStamp) String() string {
	return t.UTC().Format("20060102T150405Z")
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
