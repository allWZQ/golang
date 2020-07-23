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

// DescribeJobLog invokes the retailcloud.DescribeJobLog API synchronously
// api document: https://help.aliyun.com/api/retailcloud/describejoblog.html
func (client *Client) DescribeJobLog(request *DescribeJobLogRequest) (response *DescribeJobLogResponse, err error) {
	response = CreateDescribeJobLogResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeJobLogWithChan invokes the retailcloud.DescribeJobLog API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/describejoblog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeJobLogWithChan(request *DescribeJobLogRequest) (<-chan *DescribeJobLogResponse, <-chan error) {
	responseChan := make(chan *DescribeJobLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeJobLog(request)
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

// DescribeJobLogWithCallback invokes the retailcloud.DescribeJobLog API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/describejoblog.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeJobLogWithCallback(request *DescribeJobLogRequest, callback func(response *DescribeJobLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeJobLogResponse
		var err error
		defer close(result)
		response, err = client.DescribeJobLog(request)
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

// DescribeJobLogRequest is the request struct for api DescribeJobLog
type DescribeJobLogRequest struct {
	*requests.RpcRequest
	AppId   requests.Integer `position:"Query" name:"AppId"`
	PodName string           `position:"Query" name:"PodName"`
	EnvId   requests.Integer `position:"Query" name:"EnvId"`
}

// DescribeJobLogResponse is the response struct for api DescribeJobLog
type DescribeJobLogResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrMsg    string `json:"ErrMsg" xml:"ErrMsg"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDescribeJobLogRequest creates a request to invoke DescribeJobLog API
func CreateDescribeJobLogRequest() (request *DescribeJobLogRequest) {
	request = &DescribeJobLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "DescribeJobLog", "retailcloud", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeJobLogResponse creates a response to parse from DescribeJobLog response
func CreateDescribeJobLogResponse() (response *DescribeJobLogResponse) {
	response = &DescribeJobLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
