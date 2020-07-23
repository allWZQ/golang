package edas

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

// StartK8sApplication invokes the edas.StartK8sApplication API synchronously
// api document: https://help.aliyun.com/api/edas/startk8sapplication.html
func (client *Client) StartK8sApplication(request *StartK8sApplicationRequest) (response *StartK8sApplicationResponse, err error) {
	response = CreateStartK8sApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// StartK8sApplicationWithChan invokes the edas.StartK8sApplication API asynchronously
// api document: https://help.aliyun.com/api/edas/startk8sapplication.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartK8sApplicationWithChan(request *StartK8sApplicationRequest) (<-chan *StartK8sApplicationResponse, <-chan error) {
	responseChan := make(chan *StartK8sApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartK8sApplication(request)
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

// StartK8sApplicationWithCallback invokes the edas.StartK8sApplication API asynchronously
// api document: https://help.aliyun.com/api/edas/startk8sapplication.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) StartK8sApplicationWithCallback(request *StartK8sApplicationRequest, callback func(response *StartK8sApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartK8sApplicationResponse
		var err error
		defer close(result)
		response, err = client.StartK8sApplication(request)
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

// StartK8sApplicationRequest is the request struct for api StartK8sApplication
type StartK8sApplicationRequest struct {
	*requests.RoaRequest
	Replicas requests.Integer `position:"Query" name:"Replicas"`
	AppId    string           `position:"Query" name:"AppId"`
	Timeout  requests.Integer `position:"Query" name:"Timeout"`
}

// StartK8sApplicationResponse is the response struct for api StartK8sApplication
type StartK8sApplicationResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Code          int    `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	ChangeOrderId string `json:"ChangeOrderId" xml:"ChangeOrderId"`
}

// CreateStartK8sApplicationRequest creates a request to invoke StartK8sApplication API
func CreateStartK8sApplicationRequest() (request *StartK8sApplicationRequest) {
	request = &StartK8sApplicationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "StartK8sApplication", "/pop/v5/k8s/acs/start_k8s_app", "edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartK8sApplicationResponse creates a response to parse from StartK8sApplication response
func CreateStartK8sApplicationResponse() (response *StartK8sApplicationResponse) {
	response = &StartK8sApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
