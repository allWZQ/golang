package vod

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

// GetVideoList invokes the vod.GetVideoList API synchronously
// api document: https://help.aliyun.com/api/vod/getvideolist.html
func (client *Client) GetVideoList(request *GetVideoListRequest) (response *GetVideoListResponse, err error) {
	response = CreateGetVideoListResponse()
	err = client.DoAction(request, response)
	return
}

// GetVideoListWithChan invokes the vod.GetVideoList API asynchronously
// api document: https://help.aliyun.com/api/vod/getvideolist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetVideoListWithChan(request *GetVideoListRequest) (<-chan *GetVideoListResponse, <-chan error) {
	responseChan := make(chan *GetVideoListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetVideoList(request)
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

// GetVideoListWithCallback invokes the vod.GetVideoList API asynchronously
// api document: https://help.aliyun.com/api/vod/getvideolist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetVideoListWithCallback(request *GetVideoListRequest, callback func(response *GetVideoListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetVideoListResponse
		var err error
		defer close(result)
		response, err = client.GetVideoList(request)
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

// GetVideoListRequest is the request struct for api GetVideoList
type GetVideoListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	StorageLocation      string           `position:"Query" name:"StorageLocation"`
	CateId               requests.Integer `position:"Query" name:"CateId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PageNo               requests.Integer `position:"Query" name:"PageNo"`
	SortBy               string           `position:"Query" name:"SortBy"`
	Status               string           `position:"Query" name:"Status"`
}

// GetVideoListResponse is the response struct for api GetVideoList
type GetVideoListResponse struct {
	*responses.BaseResponse
	RequestId string                  `json:"RequestId" xml:"RequestId"`
	Total     int                     `json:"Total" xml:"Total"`
	VideoList VideoListInGetVideoList `json:"VideoList" xml:"VideoList"`
}

// CreateGetVideoListRequest creates a request to invoke GetVideoList API
func CreateGetVideoListRequest() (request *GetVideoListRequest) {
	request = &GetVideoListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetVideoList", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetVideoListResponse creates a response to parse from GetVideoList response
func CreateGetVideoListResponse() (response *GetVideoListResponse) {
	response = &GetVideoListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
