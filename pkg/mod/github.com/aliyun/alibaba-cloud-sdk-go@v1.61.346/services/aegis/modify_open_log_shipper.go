package aegis

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

// ModifyOpenLogShipper invokes the aegis.ModifyOpenLogShipper API synchronously
// api document: https://help.aliyun.com/api/aegis/modifyopenlogshipper.html
func (client *Client) ModifyOpenLogShipper(request *ModifyOpenLogShipperRequest) (response *ModifyOpenLogShipperResponse, err error) {
	response = CreateModifyOpenLogShipperResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyOpenLogShipperWithChan invokes the aegis.ModifyOpenLogShipper API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifyopenlogshipper.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyOpenLogShipperWithChan(request *ModifyOpenLogShipperRequest) (<-chan *ModifyOpenLogShipperResponse, <-chan error) {
	responseChan := make(chan *ModifyOpenLogShipperResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyOpenLogShipper(request)
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

// ModifyOpenLogShipperWithCallback invokes the aegis.ModifyOpenLogShipper API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifyopenlogshipper.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyOpenLogShipperWithCallback(request *ModifyOpenLogShipperRequest, callback func(response *ModifyOpenLogShipperResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyOpenLogShipperResponse
		var err error
		defer close(result)
		response, err = client.ModifyOpenLogShipper(request)
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

// ModifyOpenLogShipperRequest is the request struct for api ModifyOpenLogShipper
type ModifyOpenLogShipperRequest struct {
	*requests.RpcRequest
	SourceIp string `position:"Query" name:"SourceIp"`
	From     string `position:"Query" name:"From"`
	Lang     string `position:"Query" name:"Lang"`
}

// ModifyOpenLogShipperResponse is the response struct for api ModifyOpenLogShipper
type ModifyOpenLogShipperResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyOpenLogShipperRequest creates a request to invoke ModifyOpenLogShipper API
func CreateModifyOpenLogShipperRequest() (request *ModifyOpenLogShipperRequest) {
	request = &ModifyOpenLogShipperRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "ModifyOpenLogShipper", "vipaegis", "openAPI")
	return
}

// CreateModifyOpenLogShipperResponse creates a response to parse from ModifyOpenLogShipper response
func CreateModifyOpenLogShipperResponse() (response *ModifyOpenLogShipperResponse) {
	response = &ModifyOpenLogShipperResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
