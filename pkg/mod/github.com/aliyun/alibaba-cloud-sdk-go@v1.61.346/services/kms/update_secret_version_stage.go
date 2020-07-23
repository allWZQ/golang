package kms

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

// UpdateSecretVersionStage invokes the kms.UpdateSecretVersionStage API synchronously
// api document: https://help.aliyun.com/api/kms/updatesecretversionstage.html
func (client *Client) UpdateSecretVersionStage(request *UpdateSecretVersionStageRequest) (response *UpdateSecretVersionStageResponse, err error) {
	response = CreateUpdateSecretVersionStageResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateSecretVersionStageWithChan invokes the kms.UpdateSecretVersionStage API asynchronously
// api document: https://help.aliyun.com/api/kms/updatesecretversionstage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateSecretVersionStageWithChan(request *UpdateSecretVersionStageRequest) (<-chan *UpdateSecretVersionStageResponse, <-chan error) {
	responseChan := make(chan *UpdateSecretVersionStageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateSecretVersionStage(request)
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

// UpdateSecretVersionStageWithCallback invokes the kms.UpdateSecretVersionStage API asynchronously
// api document: https://help.aliyun.com/api/kms/updatesecretversionstage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateSecretVersionStageWithCallback(request *UpdateSecretVersionStageRequest, callback func(response *UpdateSecretVersionStageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateSecretVersionStageResponse
		var err error
		defer close(result)
		response, err = client.UpdateSecretVersionStage(request)
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

// UpdateSecretVersionStageRequest is the request struct for api UpdateSecretVersionStage
type UpdateSecretVersionStageRequest struct {
	*requests.RpcRequest
	RemoveFromVersion string `position:"Query" name:"RemoveFromVersion"`
	MoveToVersion     string `position:"Query" name:"MoveToVersion"`
	VersionStage      string `position:"Query" name:"VersionStage"`
	SecretName        string `position:"Query" name:"SecretName"`
}

// UpdateSecretVersionStageResponse is the response struct for api UpdateSecretVersionStage
type UpdateSecretVersionStageResponse struct {
	*responses.BaseResponse
	SecretName string `json:"SecretName" xml:"SecretName"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateSecretVersionStageRequest creates a request to invoke UpdateSecretVersionStage API
func CreateUpdateSecretVersionStageRequest() (request *UpdateSecretVersionStageRequest) {
	request = &UpdateSecretVersionStageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "UpdateSecretVersionStage", "kms-service", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateSecretVersionStageResponse creates a response to parse from UpdateSecretVersionStage response
func CreateUpdateSecretVersionStageResponse() (response *UpdateSecretVersionStageResponse) {
	response = &UpdateSecretVersionStageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
