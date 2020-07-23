package domain

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

// SaveSingleTaskForDomainNameProxyService invokes the domain.SaveSingleTaskForDomainNameProxyService API synchronously
// api document: https://help.aliyun.com/api/domain/savesingletaskfordomainnameproxyservice.html
func (client *Client) SaveSingleTaskForDomainNameProxyService(request *SaveSingleTaskForDomainNameProxyServiceRequest) (response *SaveSingleTaskForDomainNameProxyServiceResponse, err error) {
	response = CreateSaveSingleTaskForDomainNameProxyServiceResponse()
	err = client.DoAction(request, response)
	return
}

// SaveSingleTaskForDomainNameProxyServiceWithChan invokes the domain.SaveSingleTaskForDomainNameProxyService API asynchronously
// api document: https://help.aliyun.com/api/domain/savesingletaskfordomainnameproxyservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForDomainNameProxyServiceWithChan(request *SaveSingleTaskForDomainNameProxyServiceRequest) (<-chan *SaveSingleTaskForDomainNameProxyServiceResponse, <-chan error) {
	responseChan := make(chan *SaveSingleTaskForDomainNameProxyServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveSingleTaskForDomainNameProxyService(request)
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

// SaveSingleTaskForDomainNameProxyServiceWithCallback invokes the domain.SaveSingleTaskForDomainNameProxyService API asynchronously
// api document: https://help.aliyun.com/api/domain/savesingletaskfordomainnameproxyservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SaveSingleTaskForDomainNameProxyServiceWithCallback(request *SaveSingleTaskForDomainNameProxyServiceRequest, callback func(response *SaveSingleTaskForDomainNameProxyServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveSingleTaskForDomainNameProxyServiceResponse
		var err error
		defer close(result)
		response, err = client.SaveSingleTaskForDomainNameProxyService(request)
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

// SaveSingleTaskForDomainNameProxyServiceRequest is the request struct for api SaveSingleTaskForDomainNameProxyService
type SaveSingleTaskForDomainNameProxyServiceRequest struct {
	*requests.RpcRequest
	DomainName   string           `position:"Query" name:"DomainName"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	Lang         string           `position:"Query" name:"Lang"`
	Status       requests.Boolean `position:"Query" name:"Status"`
}

// SaveSingleTaskForDomainNameProxyServiceResponse is the response struct for api SaveSingleTaskForDomainNameProxyService
type SaveSingleTaskForDomainNameProxyServiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveSingleTaskForDomainNameProxyServiceRequest creates a request to invoke SaveSingleTaskForDomainNameProxyService API
func CreateSaveSingleTaskForDomainNameProxyServiceRequest() (request *SaveSingleTaskForDomainNameProxyServiceRequest) {
	request = &SaveSingleTaskForDomainNameProxyServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SaveSingleTaskForDomainNameProxyService", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSaveSingleTaskForDomainNameProxyServiceResponse creates a response to parse from SaveSingleTaskForDomainNameProxyService response
func CreateSaveSingleTaskForDomainNameProxyServiceResponse() (response *SaveSingleTaskForDomainNameProxyServiceResponse) {
	response = &SaveSingleTaskForDomainNameProxyServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
