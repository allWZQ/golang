package live

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

// AddCasterVideoResource invokes the live.AddCasterVideoResource API synchronously
// api document: https://help.aliyun.com/api/live/addcastervideoresource.html
func (client *Client) AddCasterVideoResource(request *AddCasterVideoResourceRequest) (response *AddCasterVideoResourceResponse, err error) {
	response = CreateAddCasterVideoResourceResponse()
	err = client.DoAction(request, response)
	return
}

// AddCasterVideoResourceWithChan invokes the live.AddCasterVideoResource API asynchronously
// api document: https://help.aliyun.com/api/live/addcastervideoresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddCasterVideoResourceWithChan(request *AddCasterVideoResourceRequest) (<-chan *AddCasterVideoResourceResponse, <-chan error) {
	responseChan := make(chan *AddCasterVideoResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddCasterVideoResource(request)
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

// AddCasterVideoResourceWithCallback invokes the live.AddCasterVideoResource API asynchronously
// api document: https://help.aliyun.com/api/live/addcastervideoresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddCasterVideoResourceWithCallback(request *AddCasterVideoResourceRequest, callback func(response *AddCasterVideoResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddCasterVideoResourceResponse
		var err error
		defer close(result)
		response, err = client.AddCasterVideoResource(request)
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

// AddCasterVideoResourceRequest is the request struct for api AddCasterVideoResource
type AddCasterVideoResourceRequest struct {
	*requests.RpcRequest
	EndOffset           requests.Integer `position:"Query" name:"EndOffset"`
	MaterialId          string           `position:"Query" name:"MaterialId"`
	VodUrl              string           `position:"Query" name:"VodUrl"`
	CasterId            string           `position:"Query" name:"CasterId"`
	OwnerId             requests.Integer `position:"Query" name:"OwnerId"`
	BeginOffset         requests.Integer `position:"Query" name:"BeginOffset"`
	LiveStreamUrl       string           `position:"Query" name:"LiveStreamUrl"`
	LocationId          string           `position:"Query" name:"LocationId"`
	PtsCallbackInterval requests.Integer `position:"Query" name:"PtsCallbackInterval"`
	ResourceName        string           `position:"Query" name:"ResourceName"`
	RepeatNum           requests.Integer `position:"Query" name:"RepeatNum"`
}

// AddCasterVideoResourceResponse is the response struct for api AddCasterVideoResource
type AddCasterVideoResourceResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	ResourceId string `json:"ResourceId" xml:"ResourceId"`
}

// CreateAddCasterVideoResourceRequest creates a request to invoke AddCasterVideoResource API
func CreateAddCasterVideoResourceRequest() (request *AddCasterVideoResourceRequest) {
	request = &AddCasterVideoResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "AddCasterVideoResource", "live", "openAPI")
	return
}

// CreateAddCasterVideoResourceResponse creates a response to parse from AddCasterVideoResource response
func CreateAddCasterVideoResourceResponse() (response *AddCasterVideoResourceResponse) {
	response = &AddCasterVideoResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
