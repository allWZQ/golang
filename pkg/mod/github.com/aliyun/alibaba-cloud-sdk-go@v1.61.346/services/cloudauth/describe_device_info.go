package cloudauth

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

// DescribeDeviceInfo invokes the cloudauth.DescribeDeviceInfo API synchronously
// api document: https://help.aliyun.com/api/cloudauth/describedeviceinfo.html
func (client *Client) DescribeDeviceInfo(request *DescribeDeviceInfoRequest) (response *DescribeDeviceInfoResponse, err error) {
	response = CreateDescribeDeviceInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDeviceInfoWithChan invokes the cloudauth.DescribeDeviceInfo API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/describedeviceinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDeviceInfoWithChan(request *DescribeDeviceInfoRequest) (<-chan *DescribeDeviceInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeDeviceInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDeviceInfo(request)
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

// DescribeDeviceInfoWithCallback invokes the cloudauth.DescribeDeviceInfo API asynchronously
// api document: https://help.aliyun.com/api/cloudauth/describedeviceinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDeviceInfoWithCallback(request *DescribeDeviceInfoRequest, callback func(response *DescribeDeviceInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDeviceInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeDeviceInfo(request)
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

// DescribeDeviceInfoRequest is the request struct for api DescribeDeviceInfo
type DescribeDeviceInfoRequest struct {
	*requests.RpcRequest
	UserDeviceId    string           `position:"Query" name:"UserDeviceId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	Lang            string           `position:"Query" name:"Lang"`
	ExpiredStartDay string           `position:"Query" name:"ExpiredStartDay"`
	TotalCount      requests.Integer `position:"Query" name:"TotalCount"`
	CurrentPage     requests.Integer `position:"Query" name:"CurrentPage"`
	DeviceId        string           `position:"Query" name:"DeviceId"`
	BizType         string           `position:"Query" name:"BizType"`
	ExpiredEndDay   string           `position:"Query" name:"ExpiredEndDay"`
}

// DescribeDeviceInfoResponse is the response struct for api DescribeDeviceInfo
type DescribeDeviceInfoResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	PageSize       int            `json:"PageSize" xml:"PageSize"`
	CurrentPage    int            `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount     int            `json:"TotalCount" xml:"TotalCount"`
	DeviceInfoList DeviceInfoList `json:"DeviceInfoList" xml:"DeviceInfoList"`
}

// CreateDescribeDeviceInfoRequest creates a request to invoke DescribeDeviceInfo API
func CreateDescribeDeviceInfoRequest() (request *DescribeDeviceInfoRequest) {
	request = &DescribeDeviceInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2018-09-16", "DescribeDeviceInfo", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDeviceInfoResponse creates a response to parse from DescribeDeviceInfo response
func CreateDescribeDeviceInfoResponse() (response *DescribeDeviceInfoResponse) {
	response = &DescribeDeviceInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
