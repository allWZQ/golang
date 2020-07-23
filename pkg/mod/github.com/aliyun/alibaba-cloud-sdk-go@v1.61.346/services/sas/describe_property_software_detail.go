package sas

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

// DescribePropertySoftwareDetail invokes the sas.DescribePropertySoftwareDetail API synchronously
// api document: https://help.aliyun.com/api/sas/describepropertysoftwaredetail.html
func (client *Client) DescribePropertySoftwareDetail(request *DescribePropertySoftwareDetailRequest) (response *DescribePropertySoftwareDetailResponse, err error) {
	response = CreateDescribePropertySoftwareDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePropertySoftwareDetailWithChan invokes the sas.DescribePropertySoftwareDetail API asynchronously
// api document: https://help.aliyun.com/api/sas/describepropertysoftwaredetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePropertySoftwareDetailWithChan(request *DescribePropertySoftwareDetailRequest) (<-chan *DescribePropertySoftwareDetailResponse, <-chan error) {
	responseChan := make(chan *DescribePropertySoftwareDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePropertySoftwareDetail(request)
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

// DescribePropertySoftwareDetailWithCallback invokes the sas.DescribePropertySoftwareDetail API asynchronously
// api document: https://help.aliyun.com/api/sas/describepropertysoftwaredetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePropertySoftwareDetailWithCallback(request *DescribePropertySoftwareDetailRequest, callback func(response *DescribePropertySoftwareDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePropertySoftwareDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribePropertySoftwareDetail(request)
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

// DescribePropertySoftwareDetailRequest is the request struct for api DescribePropertySoftwareDetail
type DescribePropertySoftwareDetailRequest struct {
	*requests.RpcRequest
	SoftwareVersion string           `position:"Query" name:"SoftwareVersion"`
	Remark          string           `position:"Query" name:"Remark"`
	Uuid            string           `position:"Query" name:"Uuid"`
	Path            string           `position:"Query" name:"Path"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage     requests.Integer `position:"Query" name:"CurrentPage"`
	Name            string           `position:"Query" name:"Name"`
}

// DescribePropertySoftwareDetailResponse is the response struct for api DescribePropertySoftwareDetail
type DescribePropertySoftwareDetailResponse struct {
	*responses.BaseResponse
	RequestId string             `json:"RequestId" xml:"RequestId"`
	PageInfo  PageInfo           `json:"PageInfo" xml:"PageInfo"`
	Propertys []PropertySoftware `json:"Propertys" xml:"Propertys"`
}

// CreateDescribePropertySoftwareDetailRequest creates a request to invoke DescribePropertySoftwareDetail API
func CreateDescribePropertySoftwareDetailRequest() (request *DescribePropertySoftwareDetailRequest) {
	request = &DescribePropertySoftwareDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribePropertySoftwareDetail", "sas", "openAPI")
	return
}

// CreateDescribePropertySoftwareDetailResponse creates a response to parse from DescribePropertySoftwareDetail response
func CreateDescribePropertySoftwareDetailResponse() (response *DescribePropertySoftwareDetailResponse) {
	response = &DescribePropertySoftwareDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
