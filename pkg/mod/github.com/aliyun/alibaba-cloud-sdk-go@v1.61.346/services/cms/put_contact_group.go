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

// PutContactGroup invokes the cms.PutContactGroup API synchronously
// api document: https://help.aliyun.com/api/cms/putcontactgroup.html
func (client *Client) PutContactGroup(request *PutContactGroupRequest) (response *PutContactGroupResponse, err error) {
	response = CreatePutContactGroupResponse()
	err = client.DoAction(request, response)
	return
}

// PutContactGroupWithChan invokes the cms.PutContactGroup API asynchronously
// api document: https://help.aliyun.com/api/cms/putcontactgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutContactGroupWithChan(request *PutContactGroupRequest) (<-chan *PutContactGroupResponse, <-chan error) {
	responseChan := make(chan *PutContactGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutContactGroup(request)
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

// PutContactGroupWithCallback invokes the cms.PutContactGroup API asynchronously
// api document: https://help.aliyun.com/api/cms/putcontactgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PutContactGroupWithCallback(request *PutContactGroupRequest, callback func(response *PutContactGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutContactGroupResponse
		var err error
		defer close(result)
		response, err = client.PutContactGroup(request)
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

// PutContactGroupRequest is the request struct for api PutContactGroup
type PutContactGroupRequest struct {
	*requests.RpcRequest
	ContactGroupName string    `position:"Query" name:"ContactGroupName"`
	Describe         string    `position:"Query" name:"Describe"`
	ContactNames     *[]string `position:"Query" name:"ContactNames"  type:"Repeated"`
}

// PutContactGroupResponse is the response struct for api PutContactGroup
type PutContactGroupResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreatePutContactGroupRequest creates a request to invoke PutContactGroup API
func CreatePutContactGroupRequest() (request *PutContactGroupRequest) {
	request = &PutContactGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "PutContactGroup", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePutContactGroupResponse creates a response to parse from PutContactGroup response
func CreatePutContactGroupResponse() (response *PutContactGroupResponse) {
	response = &PutContactGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
