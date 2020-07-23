package aliyuncvc

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

// JoinMeeting invokes the aliyuncvc.JoinMeeting API synchronously
// api document: https://help.aliyun.com/api/aliyuncvc/joinmeeting.html
func (client *Client) JoinMeeting(request *JoinMeetingRequest) (response *JoinMeetingResponse, err error) {
	response = CreateJoinMeetingResponse()
	err = client.DoAction(request, response)
	return
}

// JoinMeetingWithChan invokes the aliyuncvc.JoinMeeting API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/joinmeeting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) JoinMeetingWithChan(request *JoinMeetingRequest) (<-chan *JoinMeetingResponse, <-chan error) {
	responseChan := make(chan *JoinMeetingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.JoinMeeting(request)
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

// JoinMeetingWithCallback invokes the aliyuncvc.JoinMeeting API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/joinmeeting.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) JoinMeetingWithCallback(request *JoinMeetingRequest, callback func(response *JoinMeetingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *JoinMeetingResponse
		var err error
		defer close(result)
		response, err = client.JoinMeeting(request)
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

// JoinMeetingRequest is the request struct for api JoinMeeting
type JoinMeetingRequest struct {
	*requests.RpcRequest
	UserId      string `position:"Body" name:"UserId"`
	Password    string `position:"Body" name:"Password"`
	MeetingCode string `position:"Body" name:"MeetingCode"`
}

// JoinMeetingResponse is the response struct for api JoinMeeting
type JoinMeetingResponse struct {
	*responses.BaseResponse
	ErrorCode   int         `json:"ErrorCode" xml:"ErrorCode"`
	Success     bool        `json:"Success" xml:"Success"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	Message     string      `json:"Message" xml:"Message"`
	MeetingInfo MeetingInfo `json:"MeetingInfo" xml:"MeetingInfo"`
}

// CreateJoinMeetingRequest creates a request to invoke JoinMeeting API
func CreateJoinMeetingRequest() (request *JoinMeetingRequest) {
	request = &JoinMeetingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "JoinMeeting", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateJoinMeetingResponse creates a response to parse from JoinMeeting response
func CreateJoinMeetingResponse() (response *JoinMeetingResponse) {
	response = &JoinMeetingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
