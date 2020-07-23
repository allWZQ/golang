package iot

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

// GetEdgeInstanceDeployment invokes the iot.GetEdgeInstanceDeployment API synchronously
// api document: https://help.aliyun.com/api/iot/getedgeinstancedeployment.html
func (client *Client) GetEdgeInstanceDeployment(request *GetEdgeInstanceDeploymentRequest) (response *GetEdgeInstanceDeploymentResponse, err error) {
	response = CreateGetEdgeInstanceDeploymentResponse()
	err = client.DoAction(request, response)
	return
}

// GetEdgeInstanceDeploymentWithChan invokes the iot.GetEdgeInstanceDeployment API asynchronously
// api document: https://help.aliyun.com/api/iot/getedgeinstancedeployment.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetEdgeInstanceDeploymentWithChan(request *GetEdgeInstanceDeploymentRequest) (<-chan *GetEdgeInstanceDeploymentResponse, <-chan error) {
	responseChan := make(chan *GetEdgeInstanceDeploymentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetEdgeInstanceDeployment(request)
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

// GetEdgeInstanceDeploymentWithCallback invokes the iot.GetEdgeInstanceDeployment API asynchronously
// api document: https://help.aliyun.com/api/iot/getedgeinstancedeployment.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetEdgeInstanceDeploymentWithCallback(request *GetEdgeInstanceDeploymentRequest, callback func(response *GetEdgeInstanceDeploymentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetEdgeInstanceDeploymentResponse
		var err error
		defer close(result)
		response, err = client.GetEdgeInstanceDeployment(request)
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

// GetEdgeInstanceDeploymentRequest is the request struct for api GetEdgeInstanceDeployment
type GetEdgeInstanceDeploymentRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	DeploymentId  string `position:"Query" name:"DeploymentId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// GetEdgeInstanceDeploymentResponse is the response struct for api GetEdgeInstanceDeployment
type GetEdgeInstanceDeploymentResponse struct {
	*responses.BaseResponse
	RequestId    string                          `json:"RequestId" xml:"RequestId"`
	Success      bool                            `json:"Success" xml:"Success"`
	Code         string                          `json:"Code" xml:"Code"`
	ErrorMessage string                          `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInGetEdgeInstanceDeployment `json:"Data" xml:"Data"`
}

// CreateGetEdgeInstanceDeploymentRequest creates a request to invoke GetEdgeInstanceDeployment API
func CreateGetEdgeInstanceDeploymentRequest() (request *GetEdgeInstanceDeploymentRequest) {
	request = &GetEdgeInstanceDeploymentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "GetEdgeInstanceDeployment", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetEdgeInstanceDeploymentResponse creates a response to parse from GetEdgeInstanceDeployment response
func CreateGetEdgeInstanceDeploymentResponse() (response *GetEdgeInstanceDeploymentResponse) {
	response = &GetEdgeInstanceDeploymentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
