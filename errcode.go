package paapi5

import "fmt"

//Error is error codes for paapi5 package
type Error int

const (
	ErrNullPointer Error = iota + 1 //Null reference instance
	ErrHTTPStatus                   //Bad HTTP status
	ErrNoData                       //No response data
)

var errMessages = map[Error]string{
	ErrNullPointer: "Null reference instance",
	ErrHTTPStatus:  "Bad HTTP status",
	ErrNoData:      "No response data",
}

//Error method returns error message.
//This method is a implementation of error interface.
func (e Error) Error() string {
	if s, ok := errMessages[e]; ok {
		return s
	}
	return fmt.Sprintf("unknown error (%d)", int(e))
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
