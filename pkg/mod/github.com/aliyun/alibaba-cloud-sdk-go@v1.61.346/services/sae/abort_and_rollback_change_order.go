package sae

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

// AbortAndRollbackChangeOrder invokes the sae.AbortAndRollbackChangeOrder API synchronously
// api document: https://help.aliyun.com/api/sae/abortandrollbackchangeorder.html
func (client *Client) AbortAndRollbackChangeOrder(request *AbortAndRollbackChangeOrderRequest) (response *AbortAndRollbackChangeOrderResponse, err error) {
	response = CreateAbortAndRollbackChangeOrderResponse()
	err = client.DoAction(request, response)
	return
}

// AbortAndRollbackChangeOrderWithChan invokes the sae.AbortAndRollbackChangeOrder API asynchronously
// api document: https://help.aliyun.com/api/sae/abortandrollbackchangeorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AbortAndRollbackChangeOrderWithChan(request *AbortAndRollbackChangeOrderRequest) (<-chan *AbortAndRollbackChangeOrderResponse, <-chan error) {
	responseChan := make(chan *AbortAndRollbackChangeOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AbortAndRollbackChangeOrder(request)
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

// AbortAndRollbackChangeOrderWithCallback invokes the sae.AbortAndRollbackChangeOrder API asynchronously
// api document: https://help.aliyun.com/api/sae/abortandrollbackchangeorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AbortAndRollbackChangeOrderWithCallback(request *AbortAndRollbackChangeOrderRequest, callback func(response *AbortAndRollbackChangeOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AbortAndRollbackChangeOrderResponse
		var err error
		defer close(result)
		response, err = client.AbortAndRollbackChangeOrder(request)
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

// AbortAndRollbackChangeOrderRequest is the request struct for api AbortAndRollbackChangeOrder
type AbortAndRollbackChangeOrderRequest struct {
	*requests.RoaRequest
	ChangeOrderId string `position:"Query" name:"ChangeOrderId"`
}

// AbortAndRollbackChangeOrderResponse is the response struct for api AbortAndRollbackChangeOrder
type AbortAndRollbackChangeOrderResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateAbortAndRollbackChangeOrderRequest creates a request to invoke AbortAndRollbackChangeOrder API
func CreateAbortAndRollbackChangeOrderRequest() (request *AbortAndRollbackChangeOrderRequest) {
	request = &AbortAndRollbackChangeOrderRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "AbortAndRollbackChangeOrder", "/pop/v1/sam/changeorder/AbortAndRollbackChangeOrder", "serverless", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateAbortAndRollbackChangeOrderResponse creates a response to parse from AbortAndRollbackChangeOrder response
func CreateAbortAndRollbackChangeOrderResponse() (response *AbortAndRollbackChangeOrderResponse) {
	response = &AbortAndRollbackChangeOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
