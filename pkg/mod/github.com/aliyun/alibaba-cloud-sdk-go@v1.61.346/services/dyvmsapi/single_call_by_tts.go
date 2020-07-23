package dyvmsapi

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

// SingleCallByTts invokes the dyvmsapi.SingleCallByTts API synchronously
// api document: https://help.aliyun.com/api/dyvmsapi/singlecallbytts.html
func (client *Client) SingleCallByTts(request *SingleCallByTtsRequest) (response *SingleCallByTtsResponse, err error) {
	response = CreateSingleCallByTtsResponse()
	err = client.DoAction(request, response)
	return
}

// SingleCallByTtsWithChan invokes the dyvmsapi.SingleCallByTts API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/singlecallbytts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SingleCallByTtsWithChan(request *SingleCallByTtsRequest) (<-chan *SingleCallByTtsResponse, <-chan error) {
	responseChan := make(chan *SingleCallByTtsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SingleCallByTts(request)
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

// SingleCallByTtsWithCallback invokes the dyvmsapi.SingleCallByTts API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/singlecallbytts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SingleCallByTtsWithCallback(request *SingleCallByTtsRequest, callback func(response *SingleCallByTtsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SingleCallByTtsResponse
		var err error
		defer close(result)
		response, err = client.SingleCallByTts(request)
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

// SingleCallByTtsRequest is the request struct for api SingleCallByTts
type SingleCallByTtsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TtsParam             string           `position:"Query" name:"TtsParam"`
	Speed                requests.Integer `position:"Query" name:"Speed"`
	CalledNumber         string           `position:"Query" name:"CalledNumber"`
	CalledShowNumber     string           `position:"Query" name:"CalledShowNumber"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	TtsCode              string           `position:"Query" name:"TtsCode"`
	PlayTimes            requests.Integer `position:"Query" name:"PlayTimes"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Volume               requests.Integer `position:"Query" name:"Volume"`
	OutId                string           `position:"Query" name:"OutId"`
}

// SingleCallByTtsResponse is the response struct for api SingleCallByTts
type SingleCallByTtsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	CallId    string `json:"CallId" xml:"CallId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateSingleCallByTtsRequest creates a request to invoke SingleCallByTts API
func CreateSingleCallByTtsRequest() (request *SingleCallByTtsRequest) {
	request = &SingleCallByTtsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "SingleCallByTts", "dyvms", "openAPI")
	return
}

// CreateSingleCallByTtsResponse creates a response to parse from SingleCallByTts response
func CreateSingleCallByTtsResponse() (response *SingleCallByTtsResponse) {
	response = &SingleCallByTtsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
