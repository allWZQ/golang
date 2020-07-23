package ehpc

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

// ListInstalledSoftware invokes the ehpc.ListInstalledSoftware API synchronously
// api document: https://help.aliyun.com/api/ehpc/listinstalledsoftware.html
func (client *Client) ListInstalledSoftware(request *ListInstalledSoftwareRequest) (response *ListInstalledSoftwareResponse, err error) {
	response = CreateListInstalledSoftwareResponse()
	err = client.DoAction(request, response)
	return
}

// ListInstalledSoftwareWithChan invokes the ehpc.ListInstalledSoftware API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listinstalledsoftware.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListInstalledSoftwareWithChan(request *ListInstalledSoftwareRequest) (<-chan *ListInstalledSoftwareResponse, <-chan error) {
	responseChan := make(chan *ListInstalledSoftwareResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListInstalledSoftware(request)
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

// ListInstalledSoftwareWithCallback invokes the ehpc.ListInstalledSoftware API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listinstalledsoftware.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListInstalledSoftwareWithCallback(request *ListInstalledSoftwareRequest, callback func(response *ListInstalledSoftwareResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListInstalledSoftwareResponse
		var err error
		defer close(result)
		response, err = client.ListInstalledSoftware(request)
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

// ListInstalledSoftwareRequest is the request struct for api ListInstalledSoftware
type ListInstalledSoftwareRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

// ListInstalledSoftwareResponse is the response struct for api ListInstalledSoftware
type ListInstalledSoftwareResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	SoftwareList SoftwareList `json:"SoftwareList" xml:"SoftwareList"`
}

// CreateListInstalledSoftwareRequest creates a request to invoke ListInstalledSoftware API
func CreateListInstalledSoftwareRequest() (request *ListInstalledSoftwareRequest) {
	request = &ListInstalledSoftwareRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "ListInstalledSoftware", "", "")
	request.Method = requests.GET
	return
}

// CreateListInstalledSoftwareResponse creates a response to parse from ListInstalledSoftware response
func CreateListInstalledSoftwareResponse() (response *ListInstalledSoftwareResponse) {
	response = &ListInstalledSoftwareResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
