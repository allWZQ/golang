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

// UpdateEnvironment invokes the retailcloud.UpdateEnvironment API synchronously
// api document: https://help.aliyun.com/api/retailcloud/updateenvironment.html
func (client *Client) UpdateEnvironment(request *UpdateEnvironmentRequest) (response *UpdateEnvironmentResponse, err error) {
	response = CreateUpdateEnvironmentResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateEnvironmentWithChan invokes the retailcloud.UpdateEnvironment API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/updateenvironment.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateEnvironmentWithChan(request *UpdateEnvironmentRequest) (<-chan *UpdateEnvironmentResponse, <-chan error) {
	responseChan := make(chan *UpdateEnvironmentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateEnvironment(request)
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

// UpdateEnvironmentWithCallback invokes the retailcloud.UpdateEnvironment API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/updateenvironment.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateEnvironmentWithCallback(request *UpdateEnvironmentRequest, callback func(response *UpdateEnvironmentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateEnvironmentResponse
		var err error
		defer close(result)
		response, err = client.UpdateEnvironment(request)
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

// UpdateEnvironmentRequest is the request struct for api UpdateEnvironment
type UpdateEnvironmentRequest struct {
	*requests.RpcRequest
	Replicas    requests.Integer `position:"Query" name:"Replicas"`
	AppId       requests.Integer `position:"Query" name:"AppId"`
	AppSchemaId requests.Integer `position:"Query" name:"AppSchemaId"`
	AppEnvId    requests.Integer `position:"Query" name:"AppEnvId"`
}

// UpdateEnvironmentResponse is the response struct for api UpdateEnvironment
type UpdateEnvironmentResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateUpdateEnvironmentRequest creates a request to invoke UpdateEnvironment API
func CreateUpdateEnvironmentRequest() (request *UpdateEnvironmentRequest) {
	request = &UpdateEnvironmentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "UpdateEnvironment", "retailcloud", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateEnvironmentResponse creates a response to parse from UpdateEnvironment response
func CreateUpdateEnvironmentResponse() (response *UpdateEnvironmentResponse) {
	response = &UpdateEnvironmentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}