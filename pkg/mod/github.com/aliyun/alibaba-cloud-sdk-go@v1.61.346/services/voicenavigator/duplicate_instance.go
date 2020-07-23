package voicenavigator

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

// DuplicateInstance invokes the voicenavigator.DuplicateInstance API synchronously
// api document: https://help.aliyun.com/api/voicenavigator/duplicateinstance.html
func (client *Client) DuplicateInstance(request *DuplicateInstanceRequest) (response *DuplicateInstanceResponse, err error) {
	response = CreateDuplicateInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DuplicateInstanceWithChan invokes the voicenavigator.DuplicateInstance API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/duplicateinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DuplicateInstanceWithChan(request *DuplicateInstanceRequest) (<-chan *DuplicateInstanceResponse, <-chan error) {
	responseChan := make(chan *DuplicateInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DuplicateInstance(request)
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

// DuplicateInstanceWithCallback invokes the voicenavigator.DuplicateInstance API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/duplicateinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DuplicateInstanceWithCallback(request *DuplicateInstanceRequest, callback func(response *DuplicateInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DuplicateInstanceResponse
		var err error
		defer close(result)
		response, err = client.DuplicateInstance(request)
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

// DuplicateInstanceRequest is the request struct for api DuplicateInstance
type DuplicateInstanceRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DuplicateInstanceResponse is the response struct for api DuplicateInstance
type DuplicateInstanceResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	InstanceId string `json:"InstanceId" xml:"InstanceId"`
}

// CreateDuplicateInstanceRequest creates a request to invoke DuplicateInstance API
func CreateDuplicateInstanceRequest() (request *DuplicateInstanceRequest) {
	request = &DuplicateInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VoiceNavigator", "2018-06-12", "DuplicateInstance", "voicebot", "openAPI")
	return
}

// CreateDuplicateInstanceResponse creates a response to parse from DuplicateInstance response
func CreateDuplicateInstanceResponse() (response *DuplicateInstanceResponse) {
	response = &DuplicateInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
