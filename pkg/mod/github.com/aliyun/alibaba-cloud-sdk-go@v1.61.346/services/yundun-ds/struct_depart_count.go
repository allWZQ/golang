package yundun_ds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DepartCount is a nested struct in yundun_ds response
type DepartCount struct {
	DtCount          int64  `json:"DtCount" xml:"DtCount"`
	DepartName       string `json:"DepartName" xml:"DepartName"`
	UserConfirmCount int64  `json:"UserConfirmCount" xml:"UserConfirmCount"`
	LastEventCount   int64  `json:"LastEventCount" xml:"LastEventCount"`
	TopDisplayName   int64  `json:"TopDisplayName" xml:"TopDisplayName"`
	UserCount        int64  `json:"UserCount" xml:"UserCount"`
	SensitiveCount   int64  `json:"SensitiveCount" xml:"SensitiveCount"`
	TopSubTypeName   int64  `json:"TopSubTypeName" xml:"TopSubTypeName"`
	SubCount         int64  `json:"SubCount" xml:"SubCount"`
	Count            int64  `json:"Count" xml:"Count"`
	EventTotalCount  int64  `json:"EventTotalCount" xml:"EventTotalCount"`
	ConfirmCount     int64  `json:"ConfirmCount" xml:"ConfirmCount"`
	DepartCount      int64  `json:"DepartCount" xml:"DepartCount"`
}
