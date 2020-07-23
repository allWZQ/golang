package teslamaxcompute

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

// GetQuotaInstance invokes the teslamaxcompute.GetQuotaInstance API synchronously
// api document: https://help.aliyun.com/api/teslamaxcompute/getquotainstance.html
func (client *Client) GetQuotaInstance(request *GetQuotaInstanceRequest) (response *GetQuotaInstanceResponse, err error) {
	response = CreateGetQuotaInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// GetQuotaInstanceWithChan invokes the teslamaxcompute.GetQuotaInstance API asynchronously
// api document: https://help.aliyun.com/api/teslamaxcompute/getquotainstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetQuotaInstanceWithChan(request *GetQuotaInstanceRequest) (<-chan *GetQuotaInstanceResponse, <-chan error) {
	responseChan := make(chan *GetQuotaInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetQuotaInstance(request)
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

// GetQuotaInstanceWithCallback invokes the teslamaxcompute.GetQuotaInstance API asynchronously
// api document: https://help.aliyun.com/api/teslamaxcompute/getquotainstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetQuotaInstanceWithCallback(request *GetQuotaInstanceRequest, callback func(response *GetQuotaInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetQuotaInstanceResponse
		var err error
		defer close(result)
		response, err = client.GetQuotaInstance(request)
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

// GetQuotaInstanceRequest is the request struct for api GetQuotaInstance
type GetQuotaInstanceRequest struct {
	*requests.RpcRequest
	Cluster   string           `position:"Query" name:"Cluster"`
	PageSize  requests.Integer `position:"Query" name:"PageSize"`
	QuotaId   string           `position:"Query" name:"QuotaId"`
	PageNum   requests.Integer `position:"Query" name:"PageNum"`
	Region    string           `position:"Query" name:"Region"`
	QuotaName string           `position:"Query" name:"QuotaName"`
	Status    string           `position:"Query" name:"Status"`
}

// GetQuotaInstanceResponse is the response struct for api GetQuotaInstance
type GetQuotaInstanceResponse struct {
	*responses.BaseResponse
	Code      int                    `json:"Code" xml:"Code"`
	Message   string                 `json:"Message" xml:"Message"`
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Data      DataInGetQuotaInstance `json:"Data" xml:"Data"`
}

// CreateGetQuotaInstanceRequest creates a request to invoke GetQuotaInstance API
func CreateGetQuotaInstanceRequest() (request *GetQuotaInstanceRequest) {
	request = &GetQuotaInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("TeslaMaxCompute", "2018-01-04", "GetQuotaInstance", "teslamaxcompute", "openAPI")
	return
}

// CreateGetQuotaInstanceResponse creates a response to parse from GetQuotaInstance response
func CreateGetQuotaInstanceResponse() (response *GetQuotaInstanceResponse) {
	response = &GetQuotaInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
