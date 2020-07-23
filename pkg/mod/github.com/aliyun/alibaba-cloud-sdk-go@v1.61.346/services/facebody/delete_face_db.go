package facebody

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

// DeleteFaceDb invokes the facebody.DeleteFaceDb API synchronously
// api document: https://help.aliyun.com/api/facebody/deletefacedb.html
func (client *Client) DeleteFaceDb(request *DeleteFaceDbRequest) (response *DeleteFaceDbResponse, err error) {
	response = CreateDeleteFaceDbResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFaceDbWithChan invokes the facebody.DeleteFaceDb API asynchronously
// api document: https://help.aliyun.com/api/facebody/deletefacedb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFaceDbWithChan(request *DeleteFaceDbRequest) (<-chan *DeleteFaceDbResponse, <-chan error) {
	responseChan := make(chan *DeleteFaceDbResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFaceDb(request)
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

// DeleteFaceDbWithCallback invokes the facebody.DeleteFaceDb API asynchronously
// api document: https://help.aliyun.com/api/facebody/deletefacedb.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteFaceDbWithCallback(request *DeleteFaceDbRequest, callback func(response *DeleteFaceDbResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFaceDbResponse
		var err error
		defer close(result)
		response, err = client.DeleteFaceDb(request)
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

// DeleteFaceDbRequest is the request struct for api DeleteFaceDb
type DeleteFaceDbRequest struct {
	*requests.RpcRequest
	Name string `position:"Body" name:"Name"`
}

// DeleteFaceDbResponse is the response struct for api DeleteFaceDb
type DeleteFaceDbResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteFaceDbRequest creates a request to invoke DeleteFaceDb API
func CreateDeleteFaceDbRequest() (request *DeleteFaceDbRequest) {
	request = &DeleteFaceDbRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "DeleteFaceDb", "facebody", "openAPI")
	return
}

// CreateDeleteFaceDbResponse creates a response to parse from DeleteFaceDb response
func CreateDeleteFaceDbResponse() (response *DeleteFaceDbResponse) {
	response = &DeleteFaceDbResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
