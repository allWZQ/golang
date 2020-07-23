package rtc

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

// EnableMAURule invokes the rtc.EnableMAURule API synchronously
// api document: https://help.aliyun.com/api/rtc/enablemaurule.html
func (client *Client) EnableMAURule(request *EnableMAURuleRequest) (response *EnableMAURuleResponse, err error) {
	response = CreateEnableMAURuleResponse()
	err = client.DoAction(request, response)
	return
}

// EnableMAURuleWithChan invokes the rtc.EnableMAURule API asynchronously
// api document: https://help.aliyun.com/api/rtc/enablemaurule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableMAURuleWithChan(request *EnableMAURuleRequest) (<-chan *EnableMAURuleResponse, <-chan error) {
	responseChan := make(chan *EnableMAURuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableMAURule(request)
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

// EnableMAURuleWithCallback invokes the rtc.EnableMAURule API asynchronously
// api document: https://help.aliyun.com/api/rtc/enablemaurule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableMAURuleWithCallback(request *EnableMAURuleRequest, callback func(response *EnableMAURuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableMAURuleResponse
		var err error
		defer close(result)
		response, err = client.EnableMAURule(request)
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

// EnableMAURuleRequest is the request struct for api EnableMAURule
type EnableMAURuleRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	AppId   string           `position:"Query" name:"AppId"`
	RuleId  requests.Integer `position:"Query" name:"RuleId"`
}

// EnableMAURuleResponse is the response struct for api EnableMAURule
type EnableMAURuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableMAURuleRequest creates a request to invoke EnableMAURule API
func CreateEnableMAURuleRequest() (request *EnableMAURuleRequest) {
	request = &EnableMAURuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "EnableMAURule", "rtc", "openAPI")
	return
}

// CreateEnableMAURuleResponse creates a response to parse from EnableMAURule response
func CreateEnableMAURuleResponse() (response *EnableMAURuleResponse) {
	response = &EnableMAURuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
