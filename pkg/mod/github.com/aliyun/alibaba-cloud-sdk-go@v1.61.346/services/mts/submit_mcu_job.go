package mts

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

// SubmitMcuJob invokes the mts.SubmitMcuJob API synchronously
// api document: https://help.aliyun.com/api/mts/submitmcujob.html
func (client *Client) SubmitMcuJob(request *SubmitMcuJobRequest) (response *SubmitMcuJobResponse, err error) {
	response = CreateSubmitMcuJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitMcuJobWithChan invokes the mts.SubmitMcuJob API asynchronously
// api document: https://help.aliyun.com/api/mts/submitmcujob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitMcuJobWithChan(request *SubmitMcuJobRequest) (<-chan *SubmitMcuJobResponse, <-chan error) {
	responseChan := make(chan *SubmitMcuJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitMcuJob(request)
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

// SubmitMcuJobWithCallback invokes the mts.SubmitMcuJob API asynchronously
// api document: https://help.aliyun.com/api/mts/submitmcujob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitMcuJobWithCallback(request *SubmitMcuJobRequest, callback func(response *SubmitMcuJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitMcuJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitMcuJob(request)
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

// SubmitMcuJobRequest is the request struct for api SubmitMcuJob
type SubmitMcuJobRequest struct {
	*requests.RpcRequest
	Template             string           `position:"Query" name:"Template"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	UserData             string           `position:"Query" name:"UserData"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TemplateId           string           `position:"Query" name:"TemplateId"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
	Input                string           `position:"Query" name:"Input"`
}

// SubmitMcuJobResponse is the response struct for api SubmitMcuJob
type SubmitMcuJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateSubmitMcuJobRequest creates a request to invoke SubmitMcuJob API
func CreateSubmitMcuJobRequest() (request *SubmitMcuJobRequest) {
	request = &SubmitMcuJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SubmitMcuJob", "", "")
	return
}

// CreateSubmitMcuJobResponse creates a response to parse from SubmitMcuJob response
func CreateSubmitMcuJobResponse() (response *SubmitMcuJobResponse) {
	response = &SubmitMcuJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
