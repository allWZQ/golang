package iot

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

// QueryDevice invokes the iot.QueryDevice API synchronously
// api document: https://help.aliyun.com/api/iot/querydevice.html
func (client *Client) QueryDevice(request *QueryDeviceRequest) (response *QueryDeviceResponse, err error) {
	response = CreateQueryDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceWithChan invokes the iot.QueryDevice API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceWithChan(request *QueryDeviceRequest) (<-chan *QueryDeviceResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDevice(request)
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

// QueryDeviceWithCallback invokes the iot.QueryDevice API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceWithCallback(request *QueryDeviceRequest, callback func(response *QueryDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceResponse
		var err error
		defer close(result)
		response, err = client.QueryDevice(request)
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

// QueryDeviceRequest is the request struct for api QueryDevice
type QueryDeviceRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	ProductKey    string           `position:"Query" name:"ProductKey"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
}

// QueryDeviceResponse is the response struct for api QueryDevice
type QueryDeviceResponse struct {
	*responses.BaseResponse
	RequestId    string            `json:"RequestId" xml:"RequestId"`
	Success      bool              `json:"Success" xml:"Success"`
	Code         string            `json:"Code" xml:"Code"`
	ErrorMessage string            `json:"ErrorMessage" xml:"ErrorMessage"`
	Total        int               `json:"Total" xml:"Total"`
	PageSize     int               `json:"PageSize" xml:"PageSize"`
	PageCount    int               `json:"PageCount" xml:"PageCount"`
	Page         int               `json:"Page" xml:"Page"`
	Data         DataInQueryDevice `json:"Data" xml:"Data"`
}

// CreateQueryDeviceRequest creates a request to invoke QueryDevice API
func CreateQueryDeviceRequest() (request *QueryDeviceRequest) {
	request = &QueryDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDevice", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryDeviceResponse creates a response to parse from QueryDevice response
func CreateQueryDeviceResponse() (response *QueryDeviceResponse) {
	response = &QueryDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
