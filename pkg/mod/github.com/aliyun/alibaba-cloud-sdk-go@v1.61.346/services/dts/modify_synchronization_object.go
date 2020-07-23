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

package dts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifySynchronizationObject invokes the dts.ModifySynchronizationObject API synchronously
// api document: https://help.aliyun.com/api/dts/modifysynchronizationobject.html
func (client *Client) ModifySynchronizationObject(request *ModifySynchronizationObjectRequest) (response *ModifySynchronizationObjectResponse, err error) {
	response = CreateModifySynchronizationObjectResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySynchronizationObjectWithChan invokes the dts.ModifySynchronizationObject API asynchronously
// api document: https://help.aliyun.com/api/dts/modifysynchronizationobject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySynchronizationObjectWithChan(request *ModifySynchronizationObjectRequest) (<-chan *ModifySynchronizationObjectResponse, <-chan error) {
	responseChan := make(chan *ModifySynchronizationObjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySynchronizationObject(request)
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

// ModifySynchronizationObjectWithCallback invokes the dts.ModifySynchronizationObject API asynchronously
// api document: https://help.aliyun.com/api/dts/modifysynchronizationobject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySynchronizationObjectWithCallback(request *ModifySynchronizationObjectRequest, callback func(response *ModifySynchronizationObjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySynchronizationObjectResponse
		var err error
		defer close(result)
		response, err = client.ModifySynchronizationObject(request)
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

// ModifySynchronizationObjectRequest is the request struct for api ModifySynchronizationObject
type ModifySynchronizationObjectRequest struct {
	*requests.RpcRequest
	SynchronizationJobId     string `position:"Query" name:"SynchronizationJobId"`
	SynchronizationObjects   string `position:"Query" name:"SynchronizationObjects"`
	SynchronizationDirection string `position:"Query" name:"SynchronizationDirection"`
	OwnerId                  string `position:"Query" name:"OwnerId"`
}

// ModifySynchronizationObjectResponse is the response struct for api ModifySynchronizationObject
type ModifySynchronizationObjectResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Success    string `json:"Success" xml:"Success"`
	ErrCode    string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
	TaskId     string `json:"TaskId" xml:"TaskId"`
}

// CreateModifySynchronizationObjectRequest creates a request to invoke ModifySynchronizationObject API
func CreateModifySynchronizationObjectRequest() (request *ModifySynchronizationObjectRequest) {
	request = &ModifySynchronizationObjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2018-08-01", "ModifySynchronizationObject", "dts", "openAPI")
	return
}

// CreateModifySynchronizationObjectResponse creates a response to parse from ModifySynchronizationObject response
func CreateModifySynchronizationObjectResponse() (response *ModifySynchronizationObjectResponse) {
	response = &ModifySynchronizationObjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
