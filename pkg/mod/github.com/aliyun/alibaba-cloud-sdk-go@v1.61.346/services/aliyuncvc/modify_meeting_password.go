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

// ModifyMeetingPassword invokes the aliyuncvc.ModifyMeetingPassword API synchronously
// api document: https://help.aliyun.com/api/aliyuncvc/modifymeetingpassword.html
func (client *Client) ModifyMeetingPassword(request *ModifyMeetingPasswordRequest) (response *ModifyMeetingPasswordResponse, err error) {
	response = CreateModifyMeetingPasswordResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyMeetingPasswordWithChan invokes the aliyuncvc.ModifyMeetingPassword API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/modifymeetingpassword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyMeetingPasswordWithChan(request *ModifyMeetingPasswordRequest) (<-chan *ModifyMeetingPasswordResponse, <-chan error) {
	responseChan := make(chan *ModifyMeetingPasswordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyMeetingPassword(request)
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

// ModifyMeetingPasswordWithCallback invokes the aliyuncvc.ModifyMeetingPassword API asynchronously
// api document: https://help.aliyun.com/api/aliyuncvc/modifymeetingpassword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyMeetingPasswordWithCallback(request *ModifyMeetingPasswordRequest, callback func(response *ModifyMeetingPasswordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyMeetingPasswordResponse
		var err error
		defer close(result)
		response, err = client.ModifyMeetingPassword(request)
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

// ModifyMeetingPasswordRequest is the request struct for api ModifyMeetingPassword
type ModifyMeetingPasswordRequest struct {
	*requests.RpcRequest
	UserId           string           `position:"Body" name:"UserId"`
	OpenPasswordFlag requests.Boolean `position:"Body" name:"OpenPasswordFlag"`
	MeetingUUID      string           `position:"Body" name:"MeetingUUID"`
	Password         string           `position:"Body" name:"Password"`
}

// ModifyMeetingPasswordResponse is the response struct for api ModifyMeetingPassword
type ModifyMeetingPasswordResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyMeetingPasswordRequest creates a request to invoke ModifyMeetingPassword API
func CreateModifyMeetingPasswordRequest() (request *ModifyMeetingPasswordRequest) {
	request = &ModifyMeetingPasswordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "ModifyMeetingPassword", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyMeetingPasswordResponse creates a response to parse from ModifyMeetingPassword response
func CreateModifyMeetingPasswordResponse() (response *ModifyMeetingPasswordResponse) {
	response = &ModifyMeetingPasswordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
