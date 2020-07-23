package imm

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

// GetImageJob invokes the imm.GetImageJob API synchronously
// api document: https://help.aliyun.com/api/imm/getimagejob.html
func (client *Client) GetImageJob(request *GetImageJobRequest) (response *GetImageJobResponse, err error) {
	response = CreateGetImageJobResponse()
	err = client.DoAction(request, response)
	return
}

// GetImageJobWithChan invokes the imm.GetImageJob API asynchronously
// api document: https://help.aliyun.com/api/imm/getimagejob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetImageJobWithChan(request *GetImageJobRequest) (<-chan *GetImageJobResponse, <-chan error) {
	responseChan := make(chan *GetImageJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetImageJob(request)
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

// GetImageJobWithCallback invokes the imm.GetImageJob API asynchronously
// api document: https://help.aliyun.com/api/imm/getimagejob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetImageJobWithCallback(request *GetImageJobRequest, callback func(response *GetImageJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetImageJobResponse
		var err error
		defer close(result)
		response, err = client.GetImageJob(request)
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

// GetImageJobRequest is the request struct for api GetImageJob
type GetImageJobRequest struct {
	*requests.RpcRequest
	Project string `position:"Query" name:"Project"`
	JobId   string `position:"Query" name:"JobId"`
	JobType string `position:"Query" name:"JobType"`
}

// GetImageJobResponse is the response struct for api GetImageJob
type GetImageJobResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	JobId           string `json:"JobId" xml:"JobId"`
	JobType         string `json:"JobType" xml:"JobType"`
	Parameters      string `json:"Parameters" xml:"Parameters"`
	Result          string `json:"Result" xml:"Result"`
	Status          string `json:"Status" xml:"Status"`
	StartTime       string `json:"StartTime" xml:"StartTime"`
	EndTime         string `json:"EndTime" xml:"EndTime"`
	ErrorMessage    string `json:"ErrorMessage" xml:"ErrorMessage"`
	NotifyEndpoint  string `json:"NotifyEndpoint" xml:"NotifyEndpoint"`
	NotifyTopicName string `json:"NotifyTopicName" xml:"NotifyTopicName"`
	Progress        int    `json:"Progress" xml:"Progress"`
}

// CreateGetImageJobRequest creates a request to invoke GetImageJob API
func CreateGetImageJobRequest() (request *GetImageJobRequest) {
	request = &GetImageJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "GetImageJob", "imm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetImageJobResponse creates a response to parse from GetImageJob response
func CreateGetImageJobResponse() (response *GetImageJobResponse) {
	response = &GetImageJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
