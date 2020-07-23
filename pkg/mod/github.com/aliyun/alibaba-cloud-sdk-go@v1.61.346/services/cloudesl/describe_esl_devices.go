package cloudesl

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

// DescribeEslDevices invokes the cloudesl.DescribeEslDevices API synchronously
// api document: https://help.aliyun.com/api/cloudesl/describeesldevices.html
func (client *Client) DescribeEslDevices(request *DescribeEslDevicesRequest) (response *DescribeEslDevicesResponse, err error) {
	response = CreateDescribeEslDevicesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEslDevicesWithChan invokes the cloudesl.DescribeEslDevices API asynchronously
// api document: https://help.aliyun.com/api/cloudesl/describeesldevices.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEslDevicesWithChan(request *DescribeEslDevicesRequest) (<-chan *DescribeEslDevicesResponse, <-chan error) {
	responseChan := make(chan *DescribeEslDevicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEslDevices(request)
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

// DescribeEslDevicesWithCallback invokes the cloudesl.DescribeEslDevices API asynchronously
// api document: https://help.aliyun.com/api/cloudesl/describeesldevices.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEslDevicesWithCallback(request *DescribeEslDevicesRequest, callback func(response *DescribeEslDevicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEslDevicesResponse
		var err error
		defer close(result)
		response, err = client.DescribeEslDevices(request)
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

// DescribeEslDevicesRequest is the request struct for api DescribeEslDevices
type DescribeEslDevicesRequest struct {
	*requests.RpcRequest
	Type             string           `position:"Body" name:"Type"`
	StoreId          string           `position:"Body" name:"StoreId"`
	PageNumber       requests.Integer `position:"Body" name:"PageNumber"`
	EslBarCode       string           `position:"Body" name:"EslBarCode"`
	PageSize         requests.Integer `position:"Body" name:"PageSize"`
	EslStatus        string           `position:"Body" name:"EslStatus"`
	ToBatteryLevel   requests.Integer `position:"Body" name:"ToBatteryLevel"`
	FromBatteryLevel requests.Integer `position:"Body" name:"FromBatteryLevel"`
}

// DescribeEslDevicesResponse is the response struct for api DescribeEslDevices
type DescribeEslDevicesResponse struct {
	*responses.BaseResponse
	ErrorMessage   string          `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string          `json:"ErrorCode" xml:"ErrorCode"`
	PageSize       int             `json:"PageSize" xml:"PageSize"`
	Message        string          `json:"Message" xml:"Message"`
	TotalCount     int             `json:"TotalCount" xml:"TotalCount"`
	DynamicCode    string          `json:"DynamicCode" xml:"DynamicCode"`
	Code           string          `json:"Code" xml:"Code"`
	PageNumber     int             `json:"PageNumber" xml:"PageNumber"`
	DynamicMessage string          `json:"DynamicMessage" xml:"DynamicMessage"`
	RequestId      string          `json:"RequestId" xml:"RequestId"`
	Success        bool            `json:"Success" xml:"Success"`
	EslDevices     []EslDeviceInfo `json:"EslDevices" xml:"EslDevices"`
}

// CreateDescribeEslDevicesRequest creates a request to invoke DescribeEslDevices API
func CreateDescribeEslDevicesRequest() (request *DescribeEslDevicesRequest) {
	request = &DescribeEslDevicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "DescribeEslDevices", "cloudesl", "openAPI")
	return
}

// CreateDescribeEslDevicesResponse creates a response to parse from DescribeEslDevices response
func CreateDescribeEslDevicesResponse() (response *DescribeEslDevicesResponse) {
	response = &DescribeEslDevicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}