package vod

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

// SubmitWorkflowJob invokes the vod.SubmitWorkflowJob API synchronously
// api document: https://help.aliyun.com/api/vod/submitworkflowjob.html
func (client *Client) SubmitWorkflowJob(request *SubmitWorkflowJobRequest) (response *SubmitWorkflowJobResponse, err error) {
	response = CreateSubmitWorkflowJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitWorkflowJobWithChan invokes the vod.SubmitWorkflowJob API asynchronously
// api document: https://help.aliyun.com/api/vod/submitworkflowjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitWorkflowJobWithChan(request *SubmitWorkflowJobRequest) (<-chan *SubmitWorkflowJobResponse, <-chan error) {
	responseChan := make(chan *SubmitWorkflowJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitWorkflowJob(request)
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

// SubmitWorkflowJobWithCallback invokes the vod.SubmitWorkflowJob API asynchronously
// api document: https://help.aliyun.com/api/vod/submitworkflowjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitWorkflowJobWithCallback(request *SubmitWorkflowJobRequest, callback func(response *SubmitWorkflowJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitWorkflowJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitWorkflowJob(request)
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

// SubmitWorkflowJobRequest is the request struct for api SubmitWorkflowJob
type SubmitWorkflowJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	WorkflowId           string           `position:"Query" name:"WorkflowId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	MediaId              string           `position:"Query" name:"MediaId"`
	FileUrl              string           `position:"Query" name:"FileUrl"`
}

// SubmitWorkflowJobResponse is the response struct for api SubmitWorkflowJob
type SubmitWorkflowJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSubmitWorkflowJobRequest creates a request to invoke SubmitWorkflowJob API
func CreateSubmitWorkflowJobRequest() (request *SubmitWorkflowJobRequest) {
	request = &SubmitWorkflowJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "SubmitWorkflowJob", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitWorkflowJobResponse creates a response to parse from SubmitWorkflowJob response
func CreateSubmitWorkflowJobResponse() (response *SubmitWorkflowJobResponse) {
	response = &SubmitWorkflowJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
