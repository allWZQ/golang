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

// BatchGetEdgeInstanceDriverConfigs invokes the iot.BatchGetEdgeInstanceDriverConfigs API synchronously
// api document: https://help.aliyun.com/api/iot/batchgetedgeinstancedriverconfigs.html
func (client *Client) BatchGetEdgeInstanceDriverConfigs(request *BatchGetEdgeInstanceDriverConfigsRequest) (response *BatchGetEdgeInstanceDriverConfigsResponse, err error) {
	response = CreateBatchGetEdgeInstanceDriverConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// BatchGetEdgeInstanceDriverConfigsWithChan invokes the iot.BatchGetEdgeInstanceDriverConfigs API asynchronously
// api document: https://help.aliyun.com/api/iot/batchgetedgeinstancedriverconfigs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchGetEdgeInstanceDriverConfigsWithChan(request *BatchGetEdgeInstanceDriverConfigsRequest) (<-chan *BatchGetEdgeInstanceDriverConfigsResponse, <-chan error) {
	responseChan := make(chan *BatchGetEdgeInstanceDriverConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchGetEdgeInstanceDriverConfigs(request)
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

// BatchGetEdgeInstanceDriverConfigsWithCallback invokes the iot.BatchGetEdgeInstanceDriverConfigs API asynchronously
// api document: https://help.aliyun.com/api/iot/batchgetedgeinstancedriverconfigs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchGetEdgeInstanceDriverConfigsWithCallback(request *BatchGetEdgeInstanceDriverConfigsRequest, callback func(response *BatchGetEdgeInstanceDriverConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchGetEdgeInstanceDriverConfigsResponse
		var err error
		defer close(result)
		response, err = client.BatchGetEdgeInstanceDriverConfigs(request)
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

// BatchGetEdgeInstanceDriverConfigsRequest is the request struct for api BatchGetEdgeInstanceDriverConfigs
type BatchGetEdgeInstanceDriverConfigsRequest struct {
	*requests.RpcRequest
	DriverIds     *[]string `position:"Query" name:"DriverIds"  type:"Repeated"`
	IotInstanceId string    `position:"Query" name:"IotInstanceId"`
	InstanceId    string    `position:"Query" name:"InstanceId"`
	ApiProduct    string    `position:"Body" name:"ApiProduct"`
	ApiRevision   string    `position:"Body" name:"ApiRevision"`
}

// BatchGetEdgeInstanceDriverConfigsResponse is the response struct for api BatchGetEdgeInstanceDriverConfigs
type BatchGetEdgeInstanceDriverConfigsResponse struct {
	*responses.BaseResponse
	RequestId        string         `json:"RequestId" xml:"RequestId"`
	Success          bool           `json:"Success" xml:"Success"`
	Code             string         `json:"Code" xml:"Code"`
	ErrorMessage     string         `json:"ErrorMessage" xml:"ErrorMessage"`
	DriverConfigList []DriverConfig `json:"DriverConfigList" xml:"DriverConfigList"`
}

// CreateBatchGetEdgeInstanceDriverConfigsRequest creates a request to invoke BatchGetEdgeInstanceDriverConfigs API
func CreateBatchGetEdgeInstanceDriverConfigsRequest() (request *BatchGetEdgeInstanceDriverConfigsRequest) {
	request = &BatchGetEdgeInstanceDriverConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "BatchGetEdgeInstanceDriverConfigs", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBatchGetEdgeInstanceDriverConfigsResponse creates a response to parse from BatchGetEdgeInstanceDriverConfigs response
func CreateBatchGetEdgeInstanceDriverConfigsResponse() (response *BatchGetEdgeInstanceDriverConfigsResponse) {
	response = &BatchGetEdgeInstanceDriverConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
