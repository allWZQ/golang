package vpc

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

// CreateBgpGroup invokes the vpc.CreateBgpGroup API synchronously
// api document: https://help.aliyun.com/api/vpc/createbgpgroup.html
func (client *Client) CreateBgpGroup(request *CreateBgpGroupRequest) (response *CreateBgpGroupResponse, err error) {
	response = CreateCreateBgpGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateBgpGroupWithChan invokes the vpc.CreateBgpGroup API asynchronously
// api document: https://help.aliyun.com/api/vpc/createbgpgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateBgpGroupWithChan(request *CreateBgpGroupRequest) (<-chan *CreateBgpGroupResponse, <-chan error) {
	responseChan := make(chan *CreateBgpGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateBgpGroup(request)
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

// CreateBgpGroupWithCallback invokes the vpc.CreateBgpGroup API asynchronously
// api document: https://help.aliyun.com/api/vpc/createbgpgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateBgpGroupWithCallback(request *CreateBgpGroupRequest, callback func(response *CreateBgpGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateBgpGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateBgpGroup(request)
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

// CreateBgpGroupRequest is the request struct for api CreateBgpGroup
type CreateBgpGroupRequest struct {
	*requests.RpcRequest
	AuthKey              string           `position:"Query" name:"AuthKey"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Description          string           `position:"Query" name:"Description"`
	PeerAsn              requests.Integer `position:"Query" name:"PeerAsn"`
	IsFakeAsn            requests.Boolean `position:"Query" name:"IsFakeAsn"`
	IpVersion            string           `position:"Query" name:"IpVersion"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	RouterId             string           `position:"Query" name:"RouterId"`
	Name                 string           `position:"Query" name:"Name"`
	LocalAsn             requests.Integer `position:"Query" name:"LocalAsn"`
}

// CreateBgpGroupResponse is the response struct for api CreateBgpGroup
type CreateBgpGroupResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	BgpGroupId string `json:"BgpGroupId" xml:"BgpGroupId"`
}

// CreateCreateBgpGroupRequest creates a request to invoke CreateBgpGroup API
func CreateCreateBgpGroupRequest() (request *CreateBgpGroupRequest) {
	request = &CreateBgpGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateBgpGroup", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateBgpGroupResponse creates a response to parse from CreateBgpGroup response
func CreateCreateBgpGroupResponse() (response *CreateBgpGroupResponse) {
	response = &CreateBgpGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
