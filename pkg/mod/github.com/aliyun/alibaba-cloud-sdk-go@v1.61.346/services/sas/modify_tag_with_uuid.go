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

// ModifyTagWithUuid invokes the sas.ModifyTagWithUuid API synchronously
// api document: https://help.aliyun.com/api/sas/modifytagwithuuid.html
func (client *Client) ModifyTagWithUuid(request *ModifyTagWithUuidRequest) (response *ModifyTagWithUuidResponse, err error) {
	response = CreateModifyTagWithUuidResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyTagWithUuidWithChan invokes the sas.ModifyTagWithUuid API asynchronously
// api document: https://help.aliyun.com/api/sas/modifytagwithuuid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyTagWithUuidWithChan(request *ModifyTagWithUuidRequest) (<-chan *ModifyTagWithUuidResponse, <-chan error) {
	responseChan := make(chan *ModifyTagWithUuidResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyTagWithUuid(request)
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

// ModifyTagWithUuidWithCallback invokes the sas.ModifyTagWithUuid API asynchronously
// api document: https://help.aliyun.com/api/sas/modifytagwithuuid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyTagWithUuidWithCallback(request *ModifyTagWithUuidRequest, callback func(response *ModifyTagWithUuidResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyTagWithUuidResponse
		var err error
		defer close(result)
		response, err = client.ModifyTagWithUuid(request)
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

// ModifyTagWithUuidRequest is the request struct for api ModifyTagWithUuid
type ModifyTagWithUuidRequest struct {
	*requests.RpcRequest
	TagId        string `position:"Query" name:"TagId"`
	MachineTypes string `position:"Query" name:"MachineTypes"`
	TagList      string `position:"Query" name:"TagList"`
	SourceIp     string `position:"Query" name:"SourceIp"`
	UuidList     string `position:"Query" name:"UuidList"`
}

// ModifyTagWithUuidResponse is the response struct for api ModifyTagWithUuid
type ModifyTagWithUuidResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyTagWithUuidRequest creates a request to invoke ModifyTagWithUuid API
func CreateModifyTagWithUuidRequest() (request *ModifyTagWithUuidRequest) {
	request = &ModifyTagWithUuidRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ModifyTagWithUuid", "sas", "openAPI")
	return
}

// CreateModifyTagWithUuidResponse creates a response to parse from ModifyTagWithUuid response
func CreateModifyTagWithUuidResponse() (response *ModifyTagWithUuidResponse) {
	response = &ModifyTagWithUuidResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
