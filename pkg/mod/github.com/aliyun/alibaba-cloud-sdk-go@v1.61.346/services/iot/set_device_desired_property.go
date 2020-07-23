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

// SetDeviceDesiredProperty invokes the iot.SetDeviceDesiredProperty API synchronously
// api document: https://help.aliyun.com/api/iot/setdevicedesiredproperty.html
func (client *Client) SetDeviceDesiredProperty(request *SetDeviceDesiredPropertyRequest) (response *SetDeviceDesiredPropertyResponse, err error) {
	response = CreateSetDeviceDesiredPropertyResponse()
	err = client.DoAction(request, response)
	return
}

// SetDeviceDesiredPropertyWithChan invokes the iot.SetDeviceDesiredProperty API asynchronously
// api document: https://help.aliyun.com/api/iot/setdevicedesiredproperty.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetDeviceDesiredPropertyWithChan(request *SetDeviceDesiredPropertyRequest) (<-chan *SetDeviceDesiredPropertyResponse, <-chan error) {
	responseChan := make(chan *SetDeviceDesiredPropertyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDeviceDesiredProperty(request)
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

// SetDeviceDesiredPropertyWithCallback invokes the iot.SetDeviceDesiredProperty API asynchronously
// api document: https://help.aliyun.com/api/iot/setdevicedesiredproperty.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetDeviceDesiredPropertyWithCallback(request *SetDeviceDesiredPropertyRequest, callback func(response *SetDeviceDesiredPropertyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDeviceDesiredPropertyResponse
		var err error
		defer close(result)
		response, err = client.SetDeviceDesiredProperty(request)
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

// SetDeviceDesiredPropertyRequest is the request struct for api SetDeviceDesiredProperty
type SetDeviceDesiredPropertyRequest struct {
	*requests.RpcRequest
	IotId         string `position:"Query" name:"IotId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ProductKey    string `position:"Query" name:"ProductKey"`
	Versions      string `position:"Query" name:"Versions"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	DeviceName    string `position:"Query" name:"DeviceName"`
	Items         string `position:"Query" name:"Items"`
}

// SetDeviceDesiredPropertyResponse is the response struct for api SetDeviceDesiredProperty
type SetDeviceDesiredPropertyResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	Success          bool   `json:"Success" xml:"Success"`
	ErrorMessage     string `json:"ErrorMessage" xml:"ErrorMessage"`
	Code             string `json:"Code" xml:"Code"`
	MessageArguments string `json:"MessageArguments" xml:"MessageArguments"`
	Data             Data   `json:"Data" xml:"Data"`
}

// CreateSetDeviceDesiredPropertyRequest creates a request to invoke SetDeviceDesiredProperty API
func CreateSetDeviceDesiredPropertyRequest() (request *SetDeviceDesiredPropertyRequest) {
	request = &SetDeviceDesiredPropertyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "SetDeviceDesiredProperty", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetDeviceDesiredPropertyResponse creates a response to parse from SetDeviceDesiredProperty response
func CreateSetDeviceDesiredPropertyResponse() (response *SetDeviceDesiredPropertyResponse) {
	response = &SetDeviceDesiredPropertyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
