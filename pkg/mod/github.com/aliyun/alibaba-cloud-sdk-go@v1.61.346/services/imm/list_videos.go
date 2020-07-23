package imm

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

// ListVideos invokes the imm.ListVideos API synchronously
// api document: https://help.aliyun.com/api/imm/listvideos.html
func (client *Client) ListVideos(request *ListVideosRequest) (response *ListVideosResponse, err error) {
	response = CreateListVideosResponse()
	err = client.DoAction(request, response)
	return
}

// ListVideosWithChan invokes the imm.ListVideos API asynchronously
// api document: https://help.aliyun.com/api/imm/listvideos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListVideosWithChan(request *ListVideosRequest) (<-chan *ListVideosResponse, <-chan error) {
	responseChan := make(chan *ListVideosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVideos(request)
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

// ListVideosWithCallback invokes the imm.ListVideos API asynchronously
// api document: https://help.aliyun.com/api/imm/listvideos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListVideosWithCallback(request *ListVideosRequest, callback func(response *ListVideosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVideosResponse
		var err error
		defer close(result)
		response, err = client.ListVideos(request)
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

// ListVideosRequest is the request struct for api ListVideos
type ListVideosRequest struct {
	*requests.RpcRequest
	Project         string `position:"Query" name:"Project"`
	Marker          string `position:"Query" name:"Marker"`
	SetId           string `position:"Query" name:"SetId"`
	CreateTimeStart string `position:"Query" name:"CreateTimeStart"`
}

// ListVideosResponse is the response struct for api ListVideos
type ListVideosResponse struct {
	*responses.BaseResponse
	SetId      string       `json:"SetId" xml:"SetId"`
	NextMarker string       `json:"NextMarker" xml:"NextMarker"`
	RequestId  string       `json:"RequestId" xml:"RequestId"`
	Videos     []VideosItem `json:"Videos" xml:"Videos"`
}

// CreateListVideosRequest creates a request to invoke ListVideos API
func CreateListVideosRequest() (request *ListVideosRequest) {
	request = &ListVideosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "ListVideos", "imm", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListVideosResponse creates a response to parse from ListVideos response
func CreateListVideosResponse() (response *ListVideosResponse) {
	response = &ListVideosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
