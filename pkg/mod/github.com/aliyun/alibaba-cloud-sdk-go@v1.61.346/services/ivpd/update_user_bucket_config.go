package ivpd

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

// UpdateUserBucketConfig invokes the ivpd.UpdateUserBucketConfig API synchronously
// api document: https://help.aliyun.com/api/ivpd/updateuserbucketconfig.html
func (client *Client) UpdateUserBucketConfig(request *UpdateUserBucketConfigRequest) (response *UpdateUserBucketConfigResponse, err error) {
	response = CreateUpdateUserBucketConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateUserBucketConfigWithChan invokes the ivpd.UpdateUserBucketConfig API asynchronously
// api document: https://help.aliyun.com/api/ivpd/updateuserbucketconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateUserBucketConfigWithChan(request *UpdateUserBucketConfigRequest) (<-chan *UpdateUserBucketConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateUserBucketConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateUserBucketConfig(request)
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

// UpdateUserBucketConfigWithCallback invokes the ivpd.UpdateUserBucketConfig API asynchronously
// api document: https://help.aliyun.com/api/ivpd/updateuserbucketconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateUserBucketConfigWithCallback(request *UpdateUserBucketConfigRequest, callback func(response *UpdateUserBucketConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateUserBucketConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateUserBucketConfig(request)
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

// UpdateUserBucketConfigRequest is the request struct for api UpdateUserBucketConfig
type UpdateUserBucketConfigRequest struct {
	*requests.RpcRequest
	Data *[]UpdateUserBucketConfigData `position:"Body" name:"Data"  type:"Repeated"`
}

// UpdateUserBucketConfigData is a repeated param struct in UpdateUserBucketConfigRequest
type UpdateUserBucketConfigData struct {
	Bucket string `name:"Bucket"`
	Region string `name:"Region"`
}

// UpdateUserBucketConfigResponse is the response struct for api UpdateUserBucketConfig
type UpdateUserBucketConfigResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Code      string     `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateUpdateUserBucketConfigRequest creates a request to invoke UpdateUserBucketConfig API
func CreateUpdateUserBucketConfigRequest() (request *UpdateUserBucketConfigRequest) {
	request = &UpdateUserBucketConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivpd", "2019-06-25", "UpdateUserBucketConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateUserBucketConfigResponse creates a response to parse from UpdateUserBucketConfig response
func CreateUpdateUserBucketConfigResponse() (response *UpdateUserBucketConfigResponse) {
	response = &UpdateUserBucketConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
