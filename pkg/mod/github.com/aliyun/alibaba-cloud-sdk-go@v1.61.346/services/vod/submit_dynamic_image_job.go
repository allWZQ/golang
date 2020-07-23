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

// SubmitDynamicImageJob invokes the vod.SubmitDynamicImageJob API synchronously
// api document: https://help.aliyun.com/api/vod/submitdynamicimagejob.html
func (client *Client) SubmitDynamicImageJob(request *SubmitDynamicImageJobRequest) (response *SubmitDynamicImageJobResponse, err error) {
	response = CreateSubmitDynamicImageJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitDynamicImageJobWithChan invokes the vod.SubmitDynamicImageJob API asynchronously
// api document: https://help.aliyun.com/api/vod/submitdynamicimagejob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitDynamicImageJobWithChan(request *SubmitDynamicImageJobRequest) (<-chan *SubmitDynamicImageJobResponse, <-chan error) {
	responseChan := make(chan *SubmitDynamicImageJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitDynamicImageJob(request)
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

// SubmitDynamicImageJobWithCallback invokes the vod.SubmitDynamicImageJob API asynchronously
// api document: https://help.aliyun.com/api/vod/submitdynamicimagejob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitDynamicImageJobWithCallback(request *SubmitDynamicImageJobRequest, callback func(response *SubmitDynamicImageJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitDynamicImageJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitDynamicImageJob(request)
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

// SubmitDynamicImageJobRequest is the request struct for api SubmitDynamicImageJob
type SubmitDynamicImageJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DynamicImageTemplateId string           `position:"Query" name:"DynamicImageTemplateId"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	VideoId                string           `position:"Query" name:"VideoId"`
	OverrideParams         string           `position:"Query" name:"OverrideParams"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
}

// SubmitDynamicImageJobResponse is the response struct for api SubmitDynamicImageJob
type SubmitDynamicImageJobResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	DynamicImageJob DynamicImageJob `json:"DynamicImageJob" xml:"DynamicImageJob"`
}

// CreateSubmitDynamicImageJobRequest creates a request to invoke SubmitDynamicImageJob API
func CreateSubmitDynamicImageJobRequest() (request *SubmitDynamicImageJobRequest) {
	request = &SubmitDynamicImageJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "SubmitDynamicImageJob", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitDynamicImageJobResponse creates a response to parse from SubmitDynamicImageJob response
func CreateSubmitDynamicImageJobResponse() (response *SubmitDynamicImageJobResponse) {
	response = &SubmitDynamicImageJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
