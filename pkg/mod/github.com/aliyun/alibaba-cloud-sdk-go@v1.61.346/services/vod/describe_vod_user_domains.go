package vod

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

// DescribeVodUserDomains invokes the vod.DescribeVodUserDomains API synchronously
// api document: https://help.aliyun.com/api/vod/describevoduserdomains.html
func (client *Client) DescribeVodUserDomains(request *DescribeVodUserDomainsRequest) (response *DescribeVodUserDomainsResponse, err error) {
	response = CreateDescribeVodUserDomainsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVodUserDomainsWithChan invokes the vod.DescribeVodUserDomains API asynchronously
// api document: https://help.aliyun.com/api/vod/describevoduserdomains.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVodUserDomainsWithChan(request *DescribeVodUserDomainsRequest) (<-chan *DescribeVodUserDomainsResponse, <-chan error) {
	responseChan := make(chan *DescribeVodUserDomainsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVodUserDomains(request)
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

// DescribeVodUserDomainsWithCallback invokes the vod.DescribeVodUserDomains API asynchronously
// api document: https://help.aliyun.com/api/vod/describevoduserdomains.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVodUserDomainsWithCallback(request *DescribeVodUserDomainsRequest, callback func(response *DescribeVodUserDomainsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVodUserDomainsResponse
		var err error
		defer close(result)
		response, err = client.DescribeVodUserDomains(request)
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

// DescribeVodUserDomainsRequest is the request struct for api DescribeVodUserDomains
type DescribeVodUserDomainsRequest struct {
	*requests.RpcRequest
	PageNumber       requests.Integer `position:"Query" name:"PageNumber"`
	CheckDomainShow  requests.Boolean `position:"Query" name:"CheckDomainShow"`
	SecurityToken    string           `position:"Query" name:"SecurityToken"`
	CdnType          string           `position:"Query" name:"CdnType"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
	FuncFilter       string           `position:"Query" name:"FuncFilter"`
	DomainName       string           `position:"Query" name:"DomainName"`
	OwnerId          requests.Integer `position:"Query" name:"OwnerId"`
	FuncId           string           `position:"Query" name:"FuncId"`
	DomainStatus     string           `position:"Query" name:"DomainStatus"`
	DomainSearchType string           `position:"Query" name:"DomainSearchType"`
}

// DescribeVodUserDomainsResponse is the response struct for api DescribeVodUserDomains
type DescribeVodUserDomainsResponse struct {
	*responses.BaseResponse
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	PageNumber int64   `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64   `json:"PageSize" xml:"PageSize"`
	TotalCount int64   `json:"TotalCount" xml:"TotalCount"`
	Domains    Domains `json:"Domains" xml:"Domains"`
}

// CreateDescribeVodUserDomainsRequest creates a request to invoke DescribeVodUserDomains API
func CreateDescribeVodUserDomainsRequest() (request *DescribeVodUserDomainsRequest) {
	request = &DescribeVodUserDomainsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DescribeVodUserDomains", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVodUserDomainsResponse creates a response to parse from DescribeVodUserDomains response
func CreateDescribeVodUserDomainsResponse() (response *DescribeVodUserDomainsResponse) {
	response = &DescribeVodUserDomainsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
