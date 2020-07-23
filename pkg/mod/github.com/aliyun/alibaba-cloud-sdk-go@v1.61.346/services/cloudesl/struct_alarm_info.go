package cloudesl

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

// AlarmInfo is a nested struct in cloudesl response
type AlarmInfo struct {
	Remark           string `json:"Remark" xml:"Remark"`
	StoreId          string `json:"StoreId" xml:"StoreId"`
	DeviceBarCode    string `json:"DeviceBarCode" xml:"DeviceBarCode"`
	DealUserId       string `json:"DealUserId" xml:"DealUserId"`
	DeviceType       string `json:"DeviceType" xml:"DeviceType"`
	ItemTitle        string `json:"ItemTitle" xml:"ItemTitle"`
	AlarmTime        string `json:"AlarmTime" xml:"AlarmTime"`
	ErrorType        string `json:"ErrorType" xml:"ErrorType"`
	DeviceMac        string `json:"DeviceMac" xml:"DeviceMac"`
	AlarmId          string `json:"AlarmId" xml:"AlarmId"`
	RetryGmtCreate   string `json:"RetryGmtCreate" xml:"RetryGmtCreate"`
	RetryGmtModified string `json:"RetryGmtModified" xml:"RetryGmtModified"`
	ItemBarCode      string `json:"ItemBarCode" xml:"ItemBarCode"`
	RetryTimes       int64  `json:"RetryTimes" xml:"RetryTimes"`
	DealTime         string `json:"DealTime" xml:"DealTime"`
	AlarmType        string `json:"AlarmType" xml:"AlarmType"`
	AlarmStatus      string `json:"AlarmStatus" xml:"AlarmStatus"`
}
