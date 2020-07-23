package alikafka

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

// DeleteAcl invokes the alikafka.DeleteAcl API synchronously
// api document: https://help.aliyun.com/api/alikafka/deleteacl.html
func (client *Client) DeleteAcl(request *DeleteAclRequest) (response *DeleteAclResponse, err error) {
	response = CreateDeleteAclResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteAclWithChan invokes the alikafka.DeleteAcl API asynchronously
// api document: https://help.aliyun.com/api/alikafka/deleteacl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteAclWithChan(request *DeleteAclRequest) (<-chan *DeleteAclResponse, <-chan error) {
	responseChan := make(chan *DeleteAclResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteAcl(request)
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

// DeleteAclWithCallback invokes the alikafka.DeleteAcl API asynchronously
// api document: https://help.aliyun.com/api/alikafka/deleteacl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteAclWithCallback(request *DeleteAclRequest, callback func(response *DeleteAclResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteAclResponse
		var err error
		defer close(result)
		response, err = client.DeleteAcl(request)
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

// DeleteAclRequest is the request struct for api DeleteAcl
type DeleteAclRequest struct {
	*requests.RpcRequest
	AclResourcePatternType string `position:"Query" name:"AclResourcePatternType"`
	AclResourceType        string `position:"Query" name:"AclResourceType"`
	AclOperationType       string `position:"Query" name:"AclOperationType"`
	AclResourceName        string `position:"Query" name:"AclResourceName"`
	InstanceId             string `position:"Query" name:"InstanceId"`
	Username               string `position:"Query" name:"Username"`
}

// DeleteAclResponse is the response struct for api DeleteAcl
type DeleteAclResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateDeleteAclRequest creates a request to invoke DeleteAcl API
func CreateDeleteAclRequest() (request *DeleteAclRequest) {
	request = &DeleteAclRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "DeleteAcl", "alikafka", "openAPI")
	return
}

// CreateDeleteAclResponse creates a response to parse from DeleteAcl response
func CreateDeleteAclResponse() (response *DeleteAclResponse) {
	response = &DeleteAclResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
