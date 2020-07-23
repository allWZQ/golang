package retailcloud

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

// DeleteService invokes the retailcloud.DeleteService API synchronously
// api document: https://help.aliyun.com/api/retailcloud/deleteservice.html
func (client *Client) DeleteService(request *DeleteServiceRequest) (response *DeleteServiceResponse, err error) {
	response = CreateDeleteServiceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteServiceWithChan invokes the retailcloud.DeleteService API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/deleteservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteServiceWithChan(request *DeleteServiceRequest) (<-chan *DeleteServiceResponse, <-chan error) {
	responseChan := make(chan *DeleteServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteService(request)
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

// DeleteServiceWithCallback invokes the retailcloud.DeleteService API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/deleteservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteServiceWithCallback(request *DeleteServiceRequest, callback func(response *DeleteServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteServiceResponse
		var err error
		defer close(result)
		response, err = client.DeleteService(request)
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

// DeleteServiceRequest is the request struct for api DeleteService
type DeleteServiceRequest struct {
	*requests.RpcRequest
	ServiceId requests.Integer `position:"Query" name:"ServiceId"`
}

// DeleteServiceResponse is the response struct for api DeleteService
type DeleteServiceResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteServiceRequest creates a request to invoke DeleteService API
func CreateDeleteServiceRequest() (request *DeleteServiceRequest) {
	request = &DeleteServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "DeleteService", "retailcloud", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteServiceResponse creates a response to parse from DeleteService response
func CreateDeleteServiceResponse() (response *DeleteServiceResponse) {
	response = &DeleteServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
