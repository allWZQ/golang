package emr

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

// ListFlowJobHistory invokes the emr.ListFlowJobHistory API synchronously
// api document: https://help.aliyun.com/api/emr/listflowjobhistory.html
func (client *Client) ListFlowJobHistory(request *ListFlowJobHistoryRequest) (response *ListFlowJobHistoryResponse, err error) {
	response = CreateListFlowJobHistoryResponse()
	err = client.DoAction(request, response)
	return
}

// ListFlowJobHistoryWithChan invokes the emr.ListFlowJobHistory API asynchronously
// api document: https://help.aliyun.com/api/emr/listflowjobhistory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListFlowJobHistoryWithChan(request *ListFlowJobHistoryRequest) (<-chan *ListFlowJobHistoryResponse, <-chan error) {
	responseChan := make(chan *ListFlowJobHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFlowJobHistory(request)
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

// ListFlowJobHistoryWithCallback invokes the emr.ListFlowJobHistory API asynchronously
// api document: https://help.aliyun.com/api/emr/listflowjobhistory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListFlowJobHistoryWithCallback(request *ListFlowJobHistoryRequest, callback func(response *ListFlowJobHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFlowJobHistoryResponse
		var err error
		defer close(result)
		response, err = client.ListFlowJobHistory(request)
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

// ListFlowJobHistoryRequest is the request struct for api ListFlowJobHistory
type ListFlowJobHistoryRequest struct {
	*requests.RpcRequest
	TimeRange  string           `position:"Query" name:"TimeRange"`
	StatusList *[]string        `position:"Query" name:"StatusList"  type:"Repeated"`
	JobType    string           `position:"Query" name:"JobType"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Id         string           `position:"Query" name:"Id"`
	ProjectId  string           `position:"Query" name:"ProjectId"`
}

// ListFlowJobHistoryResponse is the response struct for api ListFlowJobHistory
type ListFlowJobHistoryResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	PageNumber    int           `json:"PageNumber" xml:"PageNumber"`
	PageSize      int           `json:"PageSize" xml:"PageSize"`
	Total         int           `json:"Total" xml:"Total"`
	NodeInstances NodeInstances `json:"NodeInstances" xml:"NodeInstances"`
}

// CreateListFlowJobHistoryRequest creates a request to invoke ListFlowJobHistory API
func CreateListFlowJobHistoryRequest() (request *ListFlowJobHistoryRequest) {
	request = &ListFlowJobHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListFlowJobHistory", "emr", "openAPI")
	return
}

// CreateListFlowJobHistoryResponse creates a response to parse from ListFlowJobHistory response
func CreateListFlowJobHistoryResponse() (response *ListFlowJobHistoryResponse) {
	response = &ListFlowJobHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
