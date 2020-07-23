package cms

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

// CreateMonitoringAgentProcess invokes the cms.CreateMonitoringAgentProcess API synchronously
// api document: https://help.aliyun.com/api/cms/createmonitoringagentprocess.html
func (client *Client) CreateMonitoringAgentProcess(request *CreateMonitoringAgentProcessRequest) (response *CreateMonitoringAgentProcessResponse, err error) {
	response = CreateCreateMonitoringAgentProcessResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMonitoringAgentProcessWithChan invokes the cms.CreateMonitoringAgentProcess API asynchronously
// api document: https://help.aliyun.com/api/cms/createmonitoringagentprocess.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMonitoringAgentProcessWithChan(request *CreateMonitoringAgentProcessRequest) (<-chan *CreateMonitoringAgentProcessResponse, <-chan error) {
	responseChan := make(chan *CreateMonitoringAgentProcessResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMonitoringAgentProcess(request)
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

// CreateMonitoringAgentProcessWithCallback invokes the cms.CreateMonitoringAgentProcess API asynchronously
// api document: https://help.aliyun.com/api/cms/createmonitoringagentprocess.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMonitoringAgentProcessWithCallback(request *CreateMonitoringAgentProcessRequest, callback func(response *CreateMonitoringAgentProcessResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMonitoringAgentProcessResponse
		var err error
		defer close(result)
		response, err = client.CreateMonitoringAgentProcess(request)
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

// CreateMonitoringAgentProcessRequest is the request struct for api CreateMonitoringAgentProcess
type CreateMonitoringAgentProcessRequest struct {
	*requests.RpcRequest
	ProcessName string `position:"Query" name:"ProcessName"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	ProcessUser string `position:"Query" name:"ProcessUser"`
}

// CreateMonitoringAgentProcessResponse is the response struct for api CreateMonitoringAgentProcess
type CreateMonitoringAgentProcessResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Id        int64  `json:"Id" xml:"Id"`
}

// CreateCreateMonitoringAgentProcessRequest creates a request to invoke CreateMonitoringAgentProcess API
func CreateCreateMonitoringAgentProcessRequest() (request *CreateMonitoringAgentProcessRequest) {
	request = &CreateMonitoringAgentProcessRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "CreateMonitoringAgentProcess", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateMonitoringAgentProcessResponse creates a response to parse from CreateMonitoringAgentProcess response
func CreateCreateMonitoringAgentProcessResponse() (response *CreateMonitoringAgentProcessResponse) {
	response = &CreateMonitoringAgentProcessResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
