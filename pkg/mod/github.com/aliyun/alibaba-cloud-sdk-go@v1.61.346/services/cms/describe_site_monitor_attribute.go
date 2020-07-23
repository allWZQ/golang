package cms

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeSiteMonitorAttribute invokes the cms.DescribeSiteMonitorAttribute API synchronously
// api document: https://help.aliyun.com/api/cms/describesitemonitorattribute.html
func (client *Client) DescribeSiteMonitorAttribute(request *DescribeSiteMonitorAttributeRequest) (response *DescribeSiteMonitorAttributeResponse, err error) {
	response = CreateDescribeSiteMonitorAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSiteMonitorAttributeWithChan invokes the cms.DescribeSiteMonitorAttribute API asynchronously
// api document: https://help.aliyun.com/api/cms/describesitemonitorattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSiteMonitorAttributeWithChan(request *DescribeSiteMonitorAttributeRequest) (<-chan *DescribeSiteMonitorAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeSiteMonitorAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSiteMonitorAttribute(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeSiteMonitorAttributeWithCallback invokes the cms.DescribeSiteMonitorAttribute API asynchronously
// api document: https://help.aliyun.com/api/cms/describesitemonitorattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSiteMonitorAttributeWithCallback(request *DescribeSiteMonitorAttributeRequest, callback func(response *DescribeSiteMonitorAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSiteMonitorAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeSiteMonitorAttribute(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeSiteMonitorAttributeRequest is the request struct for api DescribeSiteMonitorAttribute
type DescribeSiteMonitorAttributeRequest struct {
	*requests.RpcRequest
	IncludeAlert requests.Boolean `position:"Query" name:"IncludeAlert"`
	TaskId       string           `position:"Query" name:"TaskId"`
}

// DescribeSiteMonitorAttributeResponse is the response struct for api DescribeSiteMonitorAttribute
type DescribeSiteMonitorAttributeResponse struct {
	*responses.BaseResponse
	Code         string       `json:"Code" xml:"Code"`
	Message      string       `json:"Message" xml:"Message"`
	Success      bool         `json:"Success" xml:"Success"`
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	SiteMonitors SiteMonitors `json:"SiteMonitors" xml:"SiteMonitors"`
	MetricRules  MetricRules  `json:"MetricRules" xml:"MetricRules"`
}

// CreateDescribeSiteMonitorAttributeRequest creates a request to invoke DescribeSiteMonitorAttribute API
func CreateDescribeSiteMonitorAttributeRequest() (request *DescribeSiteMonitorAttributeRequest) {
	request = &DescribeSiteMonitorAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeSiteMonitorAttribute", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSiteMonitorAttributeResponse creates a response to parse from DescribeSiteMonitorAttribute response
func CreateDescribeSiteMonitorAttributeResponse() (response *DescribeSiteMonitorAttributeResponse) {
	response = &DescribeSiteMonitorAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}