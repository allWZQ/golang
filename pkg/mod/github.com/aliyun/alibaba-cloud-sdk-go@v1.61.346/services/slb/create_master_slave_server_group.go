package slb

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

// CreateMasterSlaveServerGroup invokes the slb.CreateMasterSlaveServerGroup API synchronously
// api document: https://help.aliyun.com/api/slb/createmasterslaveservergroup.html
func (client *Client) CreateMasterSlaveServerGroup(request *CreateMasterSlaveServerGroupRequest) (response *CreateMasterSlaveServerGroupResponse, err error) {
	response = CreateCreateMasterSlaveServerGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMasterSlaveServerGroupWithChan invokes the slb.CreateMasterSlaveServerGroup API asynchronously
// api document: https://help.aliyun.com/api/slb/createmasterslaveservergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMasterSlaveServerGroupWithChan(request *CreateMasterSlaveServerGroupRequest) (<-chan *CreateMasterSlaveServerGroupResponse, <-chan error) {
	responseChan := make(chan *CreateMasterSlaveServerGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMasterSlaveServerGroup(request)
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

// CreateMasterSlaveServerGroupWithCallback invokes the slb.CreateMasterSlaveServerGroup API asynchronously
// api document: https://help.aliyun.com/api/slb/createmasterslaveservergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateMasterSlaveServerGroupWithCallback(request *CreateMasterSlaveServerGroupRequest, callback func(response *CreateMasterSlaveServerGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMasterSlaveServerGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateMasterSlaveServerGroup(request)
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

// CreateMasterSlaveServerGroupRequest is the request struct for api CreateMasterSlaveServerGroup
type CreateMasterSlaveServerGroupRequest struct {
	*requests.RpcRequest
	AccessKeyId                string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId            requests.Integer `position:"Query" name:"ResourceOwnerId"`
	MasterSlaveBackendServers  string           `position:"Query" name:"MasterSlaveBackendServers"`
	ResourceOwnerAccount       string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount               string           `position:"Query" name:"OwnerAccount"`
	MasterSlaveServerGroupName string           `position:"Query" name:"MasterSlaveServerGroupName"`
	OwnerId                    requests.Integer `position:"Query" name:"OwnerId"`
	Tags                       string           `position:"Query" name:"Tags"`
	LoadBalancerId             string           `position:"Query" name:"LoadBalancerId"`
}

// CreateMasterSlaveServerGroupResponse is the response struct for api CreateMasterSlaveServerGroup
type CreateMasterSlaveServerGroupResponse struct {
	*responses.BaseResponse
	RequestId                 string                                                  `json:"RequestId" xml:"RequestId"`
	MasterSlaveServerGroupId  string                                                  `json:"MasterSlaveServerGroupId" xml:"MasterSlaveServerGroupId"`
	MasterSlaveBackendServers MasterSlaveBackendServersInCreateMasterSlaveServerGroup `json:"MasterSlaveBackendServers" xml:"MasterSlaveBackendServers"`
}

// CreateCreateMasterSlaveServerGroupRequest creates a request to invoke CreateMasterSlaveServerGroup API
func CreateCreateMasterSlaveServerGroupRequest() (request *CreateMasterSlaveServerGroupRequest) {
	request = &CreateMasterSlaveServerGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "CreateMasterSlaveServerGroup", "slb", "openAPI")
	return
}

// CreateCreateMasterSlaveServerGroupResponse creates a response to parse from CreateMasterSlaveServerGroup response
func CreateCreateMasterSlaveServerGroupResponse() (response *CreateMasterSlaveServerGroupResponse) {
	response = &CreateMasterSlaveServerGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
