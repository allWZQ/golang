package ons

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

// OnsWarnDelete invokes the ons.OnsWarnDelete API synchronously
// api document: https://help.aliyun.com/api/ons/onswarndelete.html
func (client *Client) OnsWarnDelete(request *OnsWarnDeleteRequest) (response *OnsWarnDeleteResponse, err error) {
	response = CreateOnsWarnDeleteResponse()
	err = client.DoAction(request, response)
	return
}

// OnsWarnDeleteWithChan invokes the ons.OnsWarnDelete API asynchronously
// api document: https://help.aliyun.com/api/ons/onswarndelete.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsWarnDeleteWithChan(request *OnsWarnDeleteRequest) (<-chan *OnsWarnDeleteResponse, <-chan error) {
	responseChan := make(chan *OnsWarnDeleteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsWarnDelete(request)
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

// OnsWarnDeleteWithCallback invokes the ons.OnsWarnDelete API asynchronously
// api document: https://help.aliyun.com/api/ons/onswarndelete.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsWarnDeleteWithCallback(request *OnsWarnDeleteRequest, callback func(response *OnsWarnDeleteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsWarnDeleteResponse
		var err error
		defer close(result)
		response, err = client.OnsWarnDelete(request)
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

// OnsWarnDeleteRequest is the request struct for api OnsWarnDelete
type OnsWarnDeleteRequest struct {
	*requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Topic      string `position:"Query" name:"Topic"`
}

// OnsWarnDeleteResponse is the response struct for api OnsWarnDelete
type OnsWarnDeleteResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
}

// CreateOnsWarnDeleteRequest creates a request to invoke OnsWarnDelete API
func CreateOnsWarnDeleteRequest() (request *OnsWarnDeleteRequest) {
	request = &OnsWarnDeleteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsWarnDelete", "ons", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOnsWarnDeleteResponse creates a response to parse from OnsWarnDelete response
func CreateOnsWarnDeleteResponse() (response *OnsWarnDeleteResponse) {
	response = &OnsWarnDeleteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}