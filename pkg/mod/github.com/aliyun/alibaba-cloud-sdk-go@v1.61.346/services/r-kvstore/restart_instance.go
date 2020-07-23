package r_kvstore

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

// RestartInstance invokes the r_kvstore.RestartInstance API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/restartinstance.html
func (client *Client) RestartInstance(request *RestartInstanceRequest) (response *RestartInstanceResponse, err error) {
	response = CreateRestartInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// RestartInstanceWithChan invokes the r_kvstore.RestartInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/restartinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RestartInstanceWithChan(request *RestartInstanceRequest) (<-chan *RestartInstanceResponse, <-chan error) {
	responseChan := make(chan *RestartInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RestartInstance(request)
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

// RestartInstanceWithCallback invokes the r_kvstore.RestartInstance API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/restartinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RestartInstanceWithCallback(request *RestartInstanceRequest, callback func(response *RestartInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RestartInstanceResponse
		var err error
		defer close(result)
		response, err = client.RestartInstance(request)
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

// RestartInstanceRequest is the request struct for api RestartInstance
type RestartInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	UpgradeMinorVersion  requests.Boolean `position:"Query" name:"UpgradeMinorVersion"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	EffectiveTime        string           `position:"Query" name:"EffectiveTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// RestartInstanceResponse is the response struct for api RestartInstance
type RestartInstanceResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	InstanceId string `json:"InstanceId" xml:"InstanceId"`
	TaskId     string `json:"TaskId" xml:"TaskId"`
}

// CreateRestartInstanceRequest creates a request to invoke RestartInstance API
func CreateRestartInstanceRequest() (request *RestartInstanceRequest) {
	request = &RestartInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "RestartInstance", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRestartInstanceResponse creates a response to parse from RestartInstance response
func CreateRestartInstanceResponse() (response *RestartInstanceResponse) {
	response = &RestartInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
