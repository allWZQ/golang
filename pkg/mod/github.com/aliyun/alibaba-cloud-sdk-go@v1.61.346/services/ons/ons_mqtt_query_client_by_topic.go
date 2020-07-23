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

// OnsMqttQueryClientByTopic invokes the ons.OnsMqttQueryClientByTopic API synchronously
// api document: https://help.aliyun.com/api/ons/onsmqttqueryclientbytopic.html
func (client *Client) OnsMqttQueryClientByTopic(request *OnsMqttQueryClientByTopicRequest) (response *OnsMqttQueryClientByTopicResponse, err error) {
	response = CreateOnsMqttQueryClientByTopicResponse()
	err = client.DoAction(request, response)
	return
}

// OnsMqttQueryClientByTopicWithChan invokes the ons.OnsMqttQueryClientByTopic API asynchronously
// api document: https://help.aliyun.com/api/ons/onsmqttqueryclientbytopic.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsMqttQueryClientByTopicWithChan(request *OnsMqttQueryClientByTopicRequest) (<-chan *OnsMqttQueryClientByTopicResponse, <-chan error) {
	responseChan := make(chan *OnsMqttQueryClientByTopicResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsMqttQueryClientByTopic(request)
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

// OnsMqttQueryClientByTopicWithCallback invokes the ons.OnsMqttQueryClientByTopic API asynchronously
// api document: https://help.aliyun.com/api/ons/onsmqttqueryclientbytopic.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsMqttQueryClientByTopicWithCallback(request *OnsMqttQueryClientByTopicRequest, callback func(response *OnsMqttQueryClientByTopicResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsMqttQueryClientByTopicResponse
		var err error
		defer close(result)
		response, err = client.OnsMqttQueryClientByTopic(request)
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

// OnsMqttQueryClientByTopicRequest is the request struct for api OnsMqttQueryClientByTopic
type OnsMqttQueryClientByTopicRequest struct {
	*requests.RpcRequest
	ParentTopic string `position:"Query" name:"ParentTopic"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	SubTopic    string `position:"Query" name:"SubTopic"`
}

// OnsMqttQueryClientByTopicResponse is the response struct for api OnsMqttQueryClientByTopic
type OnsMqttQueryClientByTopicResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	HelpUrl         string          `json:"HelpUrl" xml:"HelpUrl"`
	MqttClientSetDo MqttClientSetDo `json:"MqttClientSetDo" xml:"MqttClientSetDo"`
}

// CreateOnsMqttQueryClientByTopicRequest creates a request to invoke OnsMqttQueryClientByTopic API
func CreateOnsMqttQueryClientByTopicRequest() (request *OnsMqttQueryClientByTopicRequest) {
	request = &OnsMqttQueryClientByTopicRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsMqttQueryClientByTopic", "ons", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOnsMqttQueryClientByTopicResponse creates a response to parse from OnsMqttQueryClientByTopic response
func CreateOnsMqttQueryClientByTopicResponse() (response *OnsMqttQueryClientByTopicResponse) {
	response = &OnsMqttQueryClientByTopicResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}