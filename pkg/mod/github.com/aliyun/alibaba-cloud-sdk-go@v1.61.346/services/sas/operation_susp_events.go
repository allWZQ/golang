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

// OperationSuspEvents invokes the sas.OperationSuspEvents API synchronously
// api document: https://help.aliyun.com/api/sas/operationsuspevents.html
func (client *Client) OperationSuspEvents(request *OperationSuspEventsRequest) (response *OperationSuspEventsResponse, err error) {
	response = CreateOperationSuspEventsResponse()
	err = client.DoAction(request, response)
	return
}

// OperationSuspEventsWithChan invokes the sas.OperationSuspEvents API asynchronously
// api document: https://help.aliyun.com/api/sas/operationsuspevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OperationSuspEventsWithChan(request *OperationSuspEventsRequest) (<-chan *OperationSuspEventsResponse, <-chan error) {
	responseChan := make(chan *OperationSuspEventsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OperationSuspEvents(request)
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

// OperationSuspEventsWithCallback invokes the sas.OperationSuspEvents API asynchronously
// api document: https://help.aliyun.com/api/sas/operationsuspevents.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OperationSuspEventsWithCallback(request *OperationSuspEventsRequest, callback func(response *OperationSuspEventsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OperationSuspEventsResponse
		var err error
		defer close(result)
		response, err = client.OperationSuspEvents(request)
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

// OperationSuspEventsRequest is the request struct for api OperationSuspEvents
type OperationSuspEventsRequest struct {
	*requests.RpcRequest
	SuspiciousEventIds string `position:"Query" name:"SuspiciousEventIds"`
	SubOperation       string `position:"Query" name:"SubOperation"`
	SourceIp           string `position:"Query" name:"SourceIp"`
	WarnType           string `position:"Query" name:"WarnType"`
	From               string `position:"Query" name:"From"`
	Operation          string `position:"Query" name:"Operation"`
}

// OperationSuspEventsResponse is the response struct for api OperationSuspEvents
type OperationSuspEventsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Success    bool   `json:"Success" xml:"Success"`
	AccessCode string `json:"AccessCode" xml:"AccessCode"`
}

// CreateOperationSuspEventsRequest creates a request to invoke OperationSuspEvents API
func CreateOperationSuspEventsRequest() (request *OperationSuspEventsRequest) {
	request = &OperationSuspEventsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "OperationSuspEvents", "sas", "openAPI")
	return
}

// CreateOperationSuspEventsResponse creates a response to parse from OperationSuspEvents response
func CreateOperationSuspEventsResponse() (response *OperationSuspEventsResponse) {
	response = &OperationSuspEventsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
