package mts

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

// QueryMediaListByURL invokes the mts.QueryMediaListByURL API synchronously
// api document: https://help.aliyun.com/api/mts/querymedialistbyurl.html
func (client *Client) QueryMediaListByURL(request *QueryMediaListByURLRequest) (response *QueryMediaListByURLResponse, err error) {
	response = CreateQueryMediaListByURLResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMediaListByURLWithChan invokes the mts.QueryMediaListByURL API asynchronously
// api document: https://help.aliyun.com/api/mts/querymedialistbyurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMediaListByURLWithChan(request *QueryMediaListByURLRequest) (<-chan *QueryMediaListByURLResponse, <-chan error) {
	responseChan := make(chan *QueryMediaListByURLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMediaListByURL(request)
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

// QueryMediaListByURLWithCallback invokes the mts.QueryMediaListByURL API asynchronously
// api document: https://help.aliyun.com/api/mts/querymedialistbyurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMediaListByURLWithCallback(request *QueryMediaListByURLRequest, callback func(response *QueryMediaListByURLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMediaListByURLResponse
		var err error
		defer close(result)
		response, err = client.QueryMediaListByURL(request)
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

// QueryMediaListByURLRequest is the request struct for api QueryMediaListByURL
type QueryMediaListByURLRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	IncludeSummaryList   requests.Boolean `position:"Query" name:"IncludeSummaryList"`
	FileURLs             string           `position:"Query" name:"FileURLs"`
	IncludePlayList      requests.Boolean `position:"Query" name:"IncludePlayList"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	IncludeSnapshotList  requests.Boolean `position:"Query" name:"IncludeSnapshotList"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	IncludeMediaInfo     requests.Boolean `position:"Query" name:"IncludeMediaInfo"`
}

// QueryMediaListByURLResponse is the response struct for api QueryMediaListByURL
type QueryMediaListByURLResponse struct {
	*responses.BaseResponse
	RequestId        string                         `json:"RequestId" xml:"RequestId"`
	NonExistFileURLs NonExistFileURLs               `json:"NonExistFileURLs" xml:"NonExistFileURLs"`
	MediaList        MediaListInQueryMediaListByURL `json:"MediaList" xml:"MediaList"`
}

// CreateQueryMediaListByURLRequest creates a request to invoke QueryMediaListByURL API
func CreateQueryMediaListByURLRequest() (request *QueryMediaListByURLRequest) {
	request = &QueryMediaListByURLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaListByURL", "", "")
	return
}

// CreateQueryMediaListByURLResponse creates a response to parse from QueryMediaListByURL response
func CreateQueryMediaListByURLResponse() (response *QueryMediaListByURLResponse) {
	response = &QueryMediaListByURLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}