package sas

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

// DescribeSuspEvents invokes the sas.DescribeSuspEvents API synchronously
// api document: https://help.aliyun.com/api/sas/describesuspevents.html
func (client *Client) DescribeSuspEvents(request *DescribeSuspEventsRequest) (response *DescribeSuspEventsResponse, err error) {
	response = CreateDescribeSuspEventsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSuspEventsWithChan invokes the sas.DescribeSuspEvents API asynchronously
// api document: https://help.aliyun.com/api/sas/describesuspevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSuspEventsWithChan(request *DescribeSuspEventsRequest) (<-chan *DescribeSuspEventsResponse, <-chan error) {
	responseChan := make(chan *DescribeSuspEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSuspEvents(request)
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

// DescribeSuspEventsWithCallback invokes the sas.DescribeSuspEvents API asynchronously
// api document: https://help.aliyun.com/api/sas/describesuspevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSuspEventsWithCallback(request *DescribeSuspEventsRequest, callback func(response *DescribeSuspEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSuspEventsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSuspEvents(request)
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

// DescribeSuspEventsRequest is the request struct for api DescribeSuspEvents
type DescribeSuspEventsRequest struct {
	*requests.RpcRequest
	Remark           string `position:"Query" name:"Remark"`
	SourceIp         string `position:"Query" name:"SourceIp"`
	PageSize         string `position:"Query" name:"PageSize"`
	From             string `position:"Query" name:"From"`
	Lang             string `position:"Query" name:"Lang"`
	AlarmUniqueInfo  string `position:"Query" name:"AlarmUniqueInfo"`
	Dealed           string `position:"Query" name:"Dealed"`
	CurrentPage      string `position:"Query" name:"CurrentPage"`
	Name             string `position:"Query" name:"Name"`
	Levels           string `position:"Query" name:"Levels"`
	ParentEventTypes string `position:"Query" name:"ParentEventTypes"`
}

// DescribeSuspEventsResponse is the response struct for api DescribeSuspEvents
type DescribeSuspEventsResponse struct {
	*responses.BaseResponse
	RequestId   string           `json:"RequestId" xml:"RequestId"`
	Count       int              `json:"Count" xml:"Count"`
	PageSize    int              `json:"PageSize" xml:"PageSize"`
	TotalCount  int              `json:"TotalCount" xml:"TotalCount"`
	CurrentPage int              `json:"CurrentPage" xml:"CurrentPage"`
	SuspEvents  []WarningSummary `json:"SuspEvents" xml:"SuspEvents"`
}

// CreateDescribeSuspEventsRequest creates a request to invoke DescribeSuspEvents API
func CreateDescribeSuspEventsRequest() (request *DescribeSuspEventsRequest) {
	request = &DescribeSuspEventsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeSuspEvents", "sas", "openAPI")
	return
}

// CreateDescribeSuspEventsResponse creates a response to parse from DescribeSuspEvents response
func CreateDescribeSuspEventsResponse() (response *DescribeSuspEventsResponse) {
	response = &DescribeSuspEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
