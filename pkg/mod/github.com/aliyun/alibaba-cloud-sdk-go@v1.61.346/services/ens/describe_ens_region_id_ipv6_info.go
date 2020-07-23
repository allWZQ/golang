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

// DescribeEnsRegionIdIpv6Info invokes the ens.DescribeEnsRegionIdIpv6Info API synchronously
// api document: https://help.aliyun.com/api/ens/describeensregionidipv6info.html
func (client *Client) DescribeEnsRegionIdIpv6Info(request *DescribeEnsRegionIdIpv6InfoRequest) (response *DescribeEnsRegionIdIpv6InfoResponse, err error) {
	response = CreateDescribeEnsRegionIdIpv6InfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEnsRegionIdIpv6InfoWithChan invokes the ens.DescribeEnsRegionIdIpv6Info API asynchronously
// api document: https://help.aliyun.com/api/ens/describeensregionidipv6info.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEnsRegionIdIpv6InfoWithChan(request *DescribeEnsRegionIdIpv6InfoRequest) (<-chan *DescribeEnsRegionIdIpv6InfoResponse, <-chan error) {
	responseChan := make(chan *DescribeEnsRegionIdIpv6InfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEnsRegionIdIpv6Info(request)
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

// DescribeEnsRegionIdIpv6InfoWithCallback invokes the ens.DescribeEnsRegionIdIpv6Info API asynchronously
// api document: https://help.aliyun.com/api/ens/describeensregionidipv6info.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEnsRegionIdIpv6InfoWithCallback(request *DescribeEnsRegionIdIpv6InfoRequest, callback func(response *DescribeEnsRegionIdIpv6InfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEnsRegionIdIpv6InfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeEnsRegionIdIpv6Info(request)
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

// DescribeEnsRegionIdIpv6InfoRequest is the request struct for api DescribeEnsRegionIdIpv6Info
type DescribeEnsRegionIdIpv6InfoRequest struct {
	*requests.RpcRequest
	EnsRegionId string `position:"Query" name:"EnsRegionId"`
	Version     string `position:"Query" name:"Version"`
}

// DescribeEnsRegionIdIpv6InfoResponse is the response struct for api DescribeEnsRegionIdIpv6Info
type DescribeEnsRegionIdIpv6InfoResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	SupportIpv6Info SupportIpv6Info `json:"SupportIpv6Info" xml:"SupportIpv6Info"`
}

// CreateDescribeEnsRegionIdIpv6InfoRequest creates a request to invoke DescribeEnsRegionIdIpv6Info API
func CreateDescribeEnsRegionIdIpv6InfoRequest() (request *DescribeEnsRegionIdIpv6InfoRequest) {
	request = &DescribeEnsRegionIdIpv6InfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeEnsRegionIdIpv6Info", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeEnsRegionIdIpv6InfoResponse creates a response to parse from DescribeEnsRegionIdIpv6Info response
func CreateDescribeEnsRegionIdIpv6InfoResponse() (response *DescribeEnsRegionIdIpv6InfoResponse) {
	response = &DescribeEnsRegionIdIpv6InfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}