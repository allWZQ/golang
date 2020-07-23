package ens

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

// DescribeInstanceSpec invokes the ens.DescribeInstanceSpec API synchronously
// api document: https://help.aliyun.com/api/ens/describeinstancespec.html
func (client *Client) DescribeInstanceSpec(request *DescribeInstanceSpecRequest) (response *DescribeInstanceSpecResponse, err error) {
	response = CreateDescribeInstanceSpecResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceSpecWithChan invokes the ens.DescribeInstanceSpec API asynchronously
// api document: https://help.aliyun.com/api/ens/describeinstancespec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceSpecWithChan(request *DescribeInstanceSpecRequest) (<-chan *DescribeInstanceSpecResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceSpec(request)
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

// DescribeInstanceSpecWithCallback invokes the ens.DescribeInstanceSpec API asynchronously
// api document: https://help.aliyun.com/api/ens/describeinstancespec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceSpecWithCallback(request *DescribeInstanceSpecRequest, callback func(response *DescribeInstanceSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceSpecResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceSpec(request)
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

// DescribeInstanceSpecRequest is the request struct for api DescribeInstanceSpec
type DescribeInstanceSpecRequest struct {
	*requests.RpcRequest
	Version string `position:"Query" name:"Version"`
}

// DescribeInstanceSpecResponse is the response struct for api DescribeInstanceSpec
type DescribeInstanceSpecResponse struct {
	*responses.BaseResponse
	RequestId         string        `json:"RequestId" xml:"RequestId"`
	Code              int           `json:"Code" xml:"Code"`
	DataDiskMinSize   int           `json:"DataDiskMinSize" xml:"DataDiskMinSize"`
	DataDiskMaxSize   int           `json:"DataDiskMaxSize" xml:"DataDiskMaxSize"`
	SystemDiskMaxSize int           `json:"SystemDiskMaxSize" xml:"SystemDiskMaxSize"`
	BandwidthLimit    int           `json:"BandwidthLimit" xml:"BandwidthLimit"`
	InstanceSpecs     InstanceSpecs `json:"InstanceSpecs" xml:"InstanceSpecs"`
}

// CreateDescribeInstanceSpecRequest creates a request to invoke DescribeInstanceSpec API
func CreateDescribeInstanceSpecRequest() (request *DescribeInstanceSpecRequest) {
	request = &DescribeInstanceSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeInstanceSpec", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceSpecResponse creates a response to parse from DescribeInstanceSpec response
func CreateDescribeInstanceSpecResponse() (response *DescribeInstanceSpecResponse) {
	response = &DescribeInstanceSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}