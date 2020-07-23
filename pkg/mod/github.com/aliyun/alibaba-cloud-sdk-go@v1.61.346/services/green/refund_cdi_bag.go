package green

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

// RefundCdiBag invokes the green.RefundCdiBag API synchronously
// api document: https://help.aliyun.com/api/green/refundcdibag.html
func (client *Client) RefundCdiBag(request *RefundCdiBagRequest) (response *RefundCdiBagResponse, err error) {
	response = CreateRefundCdiBagResponse()
	err = client.DoAction(request, response)
	return
}

// RefundCdiBagWithChan invokes the green.RefundCdiBag API asynchronously
// api document: https://help.aliyun.com/api/green/refundcdibag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RefundCdiBagWithChan(request *RefundCdiBagRequest) (<-chan *RefundCdiBagResponse, <-chan error) {
	responseChan := make(chan *RefundCdiBagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RefundCdiBag(request)
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

// RefundCdiBagWithCallback invokes the green.RefundCdiBag API asynchronously
// api document: https://help.aliyun.com/api/green/refundcdibag.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RefundCdiBagWithCallback(request *RefundCdiBagRequest, callback func(response *RefundCdiBagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RefundCdiBagResponse
		var err error
		defer close(result)
		response, err = client.RefundCdiBag(request)
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

// RefundCdiBagRequest is the request struct for api RefundCdiBag
type RefundCdiBagRequest struct {
	*requests.RpcRequest
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
}

// RefundCdiBagResponse is the response struct for api RefundCdiBag
type RefundCdiBagResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRefundCdiBagRequest creates a request to invoke RefundCdiBag API
func CreateRefundCdiBagRequest() (request *RefundCdiBagRequest) {
	request = &RefundCdiBagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "RefundCdiBag", "green", "openAPI")
	return
}

// CreateRefundCdiBagResponse creates a response to parse from RefundCdiBag response
func CreateRefundCdiBagResponse() (response *RefundCdiBagResponse) {
	response = &RefundCdiBagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
