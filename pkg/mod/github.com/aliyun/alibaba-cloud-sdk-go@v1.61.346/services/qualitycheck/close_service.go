package qualitycheck

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

// CloseService invokes the qualitycheck.CloseService API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/closeservice.html
func (client *Client) CloseService(request *CloseServiceRequest) (response *CloseServiceResponse, err error) {
	response = CreateCloseServiceResponse()
	err = client.DoAction(request, response)
	return
}

// CloseServiceWithChan invokes the qualitycheck.CloseService API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/closeservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CloseServiceWithChan(request *CloseServiceRequest) (<-chan *CloseServiceResponse, <-chan error) {
	responseChan := make(chan *CloseServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CloseService(request)
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

// CloseServiceWithCallback invokes the qualitycheck.CloseService API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/closeservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CloseServiceWithCallback(request *CloseServiceRequest, callback func(response *CloseServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CloseServiceResponse
		var err error
		defer close(result)
		response, err = client.CloseService(request)
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

// CloseServiceRequest is the request struct for api CloseService
type CloseServiceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// CloseServiceResponse is the response struct for api CloseService
type CloseServiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateCloseServiceRequest creates a request to invoke CloseService API
func CreateCloseServiceRequest() (request *CloseServiceRequest) {
	request = &CloseServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "CloseService", "", "")
	return
}

// CreateCloseServiceResponse creates a response to parse from CloseService response
func CreateCloseServiceResponse() (response *CloseServiceResponse) {
	response = &CloseServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
