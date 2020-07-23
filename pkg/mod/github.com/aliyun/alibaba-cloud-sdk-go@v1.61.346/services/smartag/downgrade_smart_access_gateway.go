package smartag

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

// DowngradeSmartAccessGateway invokes the smartag.DowngradeSmartAccessGateway API synchronously
// api document: https://help.aliyun.com/api/smartag/downgradesmartaccessgateway.html
func (client *Client) DowngradeSmartAccessGateway(request *DowngradeSmartAccessGatewayRequest) (response *DowngradeSmartAccessGatewayResponse, err error) {
	response = CreateDowngradeSmartAccessGatewayResponse()
	err = client.DoAction(request, response)
	return
}

// DowngradeSmartAccessGatewayWithChan invokes the smartag.DowngradeSmartAccessGateway API asynchronously
// api document: https://help.aliyun.com/api/smartag/downgradesmartaccessgateway.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DowngradeSmartAccessGatewayWithChan(request *DowngradeSmartAccessGatewayRequest) (<-chan *DowngradeSmartAccessGatewayResponse, <-chan error) {
	responseChan := make(chan *DowngradeSmartAccessGatewayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DowngradeSmartAccessGateway(request)
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

// DowngradeSmartAccessGatewayWithCallback invokes the smartag.DowngradeSmartAccessGateway API asynchronously
// api document: https://help.aliyun.com/api/smartag/downgradesmartaccessgateway.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DowngradeSmartAccessGatewayWithCallback(request *DowngradeSmartAccessGatewayRequest, callback func(response *DowngradeSmartAccessGatewayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DowngradeSmartAccessGatewayResponse
		var err error
		defer close(result)
		response, err = client.DowngradeSmartAccessGateway(request)
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

// DowngradeSmartAccessGatewayRequest is the request struct for api DowngradeSmartAccessGateway
type DowngradeSmartAccessGatewayRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	BandWidthSpec        requests.Integer `position:"Query" name:"BandWidthSpec"`
	UserCount            requests.Integer `position:"Query" name:"UserCount"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	DataPlan             requests.Integer `position:"Query" name:"DataPlan"`
}

// DowngradeSmartAccessGatewayResponse is the response struct for api DowngradeSmartAccessGateway
type DowngradeSmartAccessGatewayResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateDowngradeSmartAccessGatewayRequest creates a request to invoke DowngradeSmartAccessGateway API
func CreateDowngradeSmartAccessGatewayRequest() (request *DowngradeSmartAccessGatewayRequest) {
	request = &DowngradeSmartAccessGatewayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DowngradeSmartAccessGateway", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDowngradeSmartAccessGatewayResponse creates a response to parse from DowngradeSmartAccessGateway response
func CreateDowngradeSmartAccessGatewayResponse() (response *DowngradeSmartAccessGatewayResponse) {
	response = &DowngradeSmartAccessGatewayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
