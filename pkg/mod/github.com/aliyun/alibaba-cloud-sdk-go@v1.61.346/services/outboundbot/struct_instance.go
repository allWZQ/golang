package outboundbot

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

// Instance is a nested struct in outboundbot response
type Instance struct {
	NluServiceType            string     `json:"NluServiceType" xml:"NluServiceType"`
	InstanceId                string     `json:"InstanceId" xml:"InstanceId"`
	InstanceName              string     `json:"InstanceName" xml:"InstanceName"`
	CallCenterInstanceId      string     `json:"CallCenterInstanceId" xml:"CallCenterInstanceId"`
	CreationTime              int64      `json:"CreationTime" xml:"CreationTime"`
	Owner                     string     `json:"Owner" xml:"Owner"`
	MaxConcurrentConversation int        `json:"MaxConcurrentConversation" xml:"MaxConcurrentConversation"`
	InstanceDescription       string     `json:"InstanceDescription" xml:"InstanceDescription"`
	IsTemplateContainer       bool       `json:"IsTemplateContainer" xml:"IsTemplateContainer"`
	NluProfile                NluProfile `json:"NluProfile" xml:"NluProfile"`
}