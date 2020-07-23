package r_kvstore

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

// DescribeInstanceSSL invokes the r_kvstore.DescribeInstanceSSL API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeinstancessl.html
func (client *Client) DescribeInstanceSSL(request *DescribeInstanceSSLRequest) (response *DescribeInstanceSSLResponse, err error) {
	response = CreateDescribeInstanceSSLResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceSSLWithChan invokes the r_kvstore.DescribeInstanceSSL API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeinstancessl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceSSLWithChan(request *DescribeInstanceSSLRequest) (<-chan *DescribeInstanceSSLResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceSSLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceSSL(request)
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

// DescribeInstanceSSLWithCallback invokes the r_kvstore.DescribeInstanceSSL API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describeinstancessl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceSSLWithCallback(request *DescribeInstanceSSLRequest, callback func(response *DescribeInstanceSSLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceSSLResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceSSL(request)
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

// DescribeInstanceSSLRequest is the request struct for api DescribeInstanceSSL
type DescribeInstanceSSLRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DescribeInstanceSSLResponse is the response struct for api DescribeInstanceSSL
type DescribeInstanceSSLResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	InstanceId     string `json:"InstanceId" xml:"InstanceId"`
	SSLEnabled     string `json:"SSLEnabled" xml:"SSLEnabled"`
	CertCommonName string `json:"CertCommonName" xml:"CertCommonName"`
	SSLExpiredTime string `json:"SSLExpiredTime" xml:"SSLExpiredTime"`
}

// CreateDescribeInstanceSSLRequest creates a request to invoke DescribeInstanceSSL API
func CreateDescribeInstanceSSLRequest() (request *DescribeInstanceSSLRequest) {
	request = &DescribeInstanceSSLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeInstanceSSL", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceSSLResponse creates a response to parse from DescribeInstanceSSL response
func CreateDescribeInstanceSSLResponse() (response *DescribeInstanceSSLResponse) {
	response = &DescribeInstanceSSLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
