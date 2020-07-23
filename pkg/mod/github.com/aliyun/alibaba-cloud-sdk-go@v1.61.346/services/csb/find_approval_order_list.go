package csb

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

// FindApprovalOrderList invokes the csb.FindApprovalOrderList API synchronously
// api document: https://help.aliyun.com/api/csb/findapprovalorderlist.html
func (client *Client) FindApprovalOrderList(request *FindApprovalOrderListRequest) (response *FindApprovalOrderListResponse, err error) {
	response = CreateFindApprovalOrderListResponse()
	err = client.DoAction(request, response)
	return
}

// FindApprovalOrderListWithChan invokes the csb.FindApprovalOrderList API asynchronously
// api document: https://help.aliyun.com/api/csb/findapprovalorderlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindApprovalOrderListWithChan(request *FindApprovalOrderListRequest) (<-chan *FindApprovalOrderListResponse, <-chan error) {
	responseChan := make(chan *FindApprovalOrderListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FindApprovalOrderList(request)
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

// FindApprovalOrderListWithCallback invokes the csb.FindApprovalOrderList API asynchronously
// api document: https://help.aliyun.com/api/csb/findapprovalorderlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindApprovalOrderListWithCallback(request *FindApprovalOrderListRequest, callback func(response *FindApprovalOrderListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FindApprovalOrderListResponse
		var err error
		defer close(result)
		response, err = client.FindApprovalOrderList(request)
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

// FindApprovalOrderListRequest is the request struct for api FindApprovalOrderList
type FindApprovalOrderListRequest struct {
	*requests.RpcRequest
	ProjectName string           `position:"Query" name:"ProjectName"`
	CsbId       requests.Integer `position:"Query" name:"CsbId"`
	PageNum     requests.Integer `position:"Query" name:"PageNum"`
	Alias       string           `position:"Query" name:"Alias"`
	ServiceName string           `position:"Query" name:"ServiceName"`
	ServiceId   requests.Integer `position:"Query" name:"ServiceId"`
	OnlyPending requests.Boolean `position:"Query" name:"OnlyPending"`
}

// FindApprovalOrderListResponse is the response struct for api FindApprovalOrderList
type FindApprovalOrderListResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateFindApprovalOrderListRequest creates a request to invoke FindApprovalOrderList API
func CreateFindApprovalOrderListRequest() (request *FindApprovalOrderListRequest) {
	request = &FindApprovalOrderListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CSB", "2017-11-18", "FindApprovalOrderList", "", "")
	return
}

// CreateFindApprovalOrderListResponse creates a response to parse from FindApprovalOrderList response
func CreateFindApprovalOrderListResponse() (response *FindApprovalOrderListResponse) {
	response = &FindApprovalOrderListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}