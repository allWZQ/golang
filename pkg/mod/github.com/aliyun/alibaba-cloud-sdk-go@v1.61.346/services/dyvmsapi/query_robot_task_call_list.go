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

// QueryRobotTaskCallList invokes the dyvmsapi.QueryRobotTaskCallList API synchronously
// api document: https://help.aliyun.com/api/dyvmsapi/queryrobottaskcalllist.html
func (client *Client) QueryRobotTaskCallList(request *QueryRobotTaskCallListRequest) (response *QueryRobotTaskCallListResponse, err error) {
	response = CreateQueryRobotTaskCallListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryRobotTaskCallListWithChan invokes the dyvmsapi.QueryRobotTaskCallList API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/queryrobottaskcalllist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryRobotTaskCallListWithChan(request *QueryRobotTaskCallListRequest) (<-chan *QueryRobotTaskCallListResponse, <-chan error) {
	responseChan := make(chan *QueryRobotTaskCallListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryRobotTaskCallList(request)
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

// QueryRobotTaskCallListWithCallback invokes the dyvmsapi.QueryRobotTaskCallList API asynchronously
// api document: https://help.aliyun.com/api/dyvmsapi/queryrobottaskcalllist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryRobotTaskCallListWithCallback(request *QueryRobotTaskCallListRequest, callback func(response *QueryRobotTaskCallListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryRobotTaskCallListResponse
		var err error
		defer close(result)
		response, err = client.QueryRobotTaskCallList(request)
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

// QueryRobotTaskCallListRequest is the request struct for api QueryRobotTaskCallList
type QueryRobotTaskCallListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Called               string           `position:"Query" name:"Called"`
	DialogCountTo        string           `position:"Query" name:"DialogCountTo"`
	DurationFrom         string           `position:"Query" name:"DurationFrom"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	TaskId               string           `position:"Query" name:"TaskId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DialogCountFrom      string           `position:"Query" name:"DialogCountFrom"`
	DurationTo           string           `position:"Query" name:"DurationTo"`
	HangupDirection      string           `position:"Query" name:"HangupDirection"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PageNo               requests.Integer `position:"Query" name:"PageNo"`
	CallResult           string           `position:"Query" name:"CallResult"`
}

// QueryRobotTaskCallListResponse is the response struct for api QueryRobotTaskCallList
type QueryRobotTaskCallListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateQueryRobotTaskCallListRequest creates a request to invoke QueryRobotTaskCallList API
func CreateQueryRobotTaskCallListRequest() (request *QueryRobotTaskCallListRequest) {
	request = &QueryRobotTaskCallListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "QueryRobotTaskCallList", "dyvms", "openAPI")
	return
}

// CreateQueryRobotTaskCallListResponse creates a response to parse from QueryRobotTaskCallList response
func CreateQueryRobotTaskCallListResponse() (response *QueryRobotTaskCallListResponse) {
	response = &QueryRobotTaskCallListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
