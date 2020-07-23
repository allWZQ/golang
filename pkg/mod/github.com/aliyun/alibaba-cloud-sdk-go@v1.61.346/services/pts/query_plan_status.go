package pts

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

// QueryPlanStatus invokes the pts.QueryPlanStatus API synchronously
// api document: https://help.aliyun.com/api/pts/queryplanstatus.html
func (client *Client) QueryPlanStatus(request *QueryPlanStatusRequest) (response *QueryPlanStatusResponse, err error) {
	response = CreateQueryPlanStatusResponse()
	err = client.DoAction(request, response)
	return
}

// QueryPlanStatusWithChan invokes the pts.QueryPlanStatus API asynchronously
// api document: https://help.aliyun.com/api/pts/queryplanstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPlanStatusWithChan(request *QueryPlanStatusRequest) (<-chan *QueryPlanStatusResponse, <-chan error) {
	responseChan := make(chan *QueryPlanStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryPlanStatus(request)
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

// QueryPlanStatusWithCallback invokes the pts.QueryPlanStatus API asynchronously
// api document: https://help.aliyun.com/api/pts/queryplanstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryPlanStatusWithCallback(request *QueryPlanStatusRequest, callback func(response *QueryPlanStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryPlanStatusResponse
		var err error
		defer close(result)
		response, err = client.QueryPlanStatus(request)
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

// QueryPlanStatusRequest is the request struct for api QueryPlanStatus
type QueryPlanStatusRequest struct {
	*requests.RpcRequest
	ReportId requests.Integer `position:"Query" name:"ReportId"`
	SceneId  requests.Integer `position:"Query" name:"SceneId"`
}

// QueryPlanStatusResponse is the response struct for api QueryPlanStatus
type QueryPlanStatusResponse struct {
	*responses.BaseResponse
	RequestId           string         `json:"RequestId" xml:"RequestId"`
	Code                string         `json:"Code" xml:"Code"`
	Message             string         `json:"Message" xml:"Message"`
	Success             bool           `json:"Success" xml:"Success"`
	Tips                string         `json:"Tips" xml:"Tips"`
	RequestCount        string         `json:"RequestCount" xml:"RequestCount"`
	Vum                 int            `json:"Vum" xml:"Vum"`
	BpsRequest          string         `json:"BpsRequest" xml:"BpsRequest"`
	BpsResponse         string         `json:"BpsResponse" xml:"BpsResponse"`
	FailedRequestCount  int            `json:"FailedRequestCount" xml:"FailedRequestCount"`
	FailedBusinessCount int            `json:"FailedBusinessCount" xml:"FailedBusinessCount"`
	Concurrency         int            `json:"Concurrency" xml:"Concurrency"`
	ConcurrencyLimit    int            `json:"ConcurrencyLimit" xml:"ConcurrencyLimit"`
	Tps                 int            `json:"Tps" xml:"Tps"`
	TpsLimit            int            `json:"TpsLimit" xml:"TpsLimit"`
	AliveAgentCount     int            `json:"AliveAgentCount" xml:"AliveAgentCount"`
	TotalAgentCount     int            `json:"TotalAgentCount" xml:"TotalAgentCount"`
	Seg90Rt             int            `json:"Seg90Rt" xml:"Seg90Rt"`
	AverageRt           int            `json:"AverageRt" xml:"AverageRt"`
	ReportId            int64          `json:"ReportId" xml:"ReportId"`
	StartTime           int64          `json:"StartTime" xml:"StartTime"`
	CurrentTime         int64          `json:"CurrentTime" xml:"CurrentTime"`
	MonitorData         MonitorData    `json:"MonitorData" xml:"MonitorData"`
	AgentLocations      AgentLocations `json:"AgentLocations" xml:"AgentLocations"`
}

// CreateQueryPlanStatusRequest creates a request to invoke QueryPlanStatus API
func CreateQueryPlanStatusRequest() (request *QueryPlanStatusRequest) {
	request = &QueryPlanStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PTS", "2018-11-11", "QueryPlanStatus", "1.0.0", "openAPI")
	return
}

// CreateQueryPlanStatusResponse creates a response to parse from QueryPlanStatus response
func CreateQueryPlanStatusResponse() (response *QueryPlanStatusResponse) {
	response = &QueryPlanStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
