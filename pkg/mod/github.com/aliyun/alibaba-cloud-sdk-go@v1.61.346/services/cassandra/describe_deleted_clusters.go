package cassandra

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

// DescribeDeletedClusters invokes the cassandra.DescribeDeletedClusters API synchronously
// api document: https://help.aliyun.com/api/cassandra/describedeletedclusters.html
func (client *Client) DescribeDeletedClusters(request *DescribeDeletedClustersRequest) (response *DescribeDeletedClustersResponse, err error) {
	response = CreateDescribeDeletedClustersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDeletedClustersWithChan invokes the cassandra.DescribeDeletedClusters API asynchronously
// api document: https://help.aliyun.com/api/cassandra/describedeletedclusters.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDeletedClustersWithChan(request *DescribeDeletedClustersRequest) (<-chan *DescribeDeletedClustersResponse, <-chan error) {
	responseChan := make(chan *DescribeDeletedClustersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDeletedClusters(request)
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

// DescribeDeletedClustersWithCallback invokes the cassandra.DescribeDeletedClusters API asynchronously
// api document: https://help.aliyun.com/api/cassandra/describedeletedclusters.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDeletedClustersWithCallback(request *DescribeDeletedClustersRequest, callback func(response *DescribeDeletedClustersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDeletedClustersResponse
		var err error
		defer close(result)
		response, err = client.DescribeDeletedClusters(request)
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

// DescribeDeletedClustersRequest is the request struct for api DescribeDeletedClusters
type DescribeDeletedClustersRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeDeletedClustersResponse is the response struct for api DescribeDeletedClusters
type DescribeDeletedClustersResponse struct {
	*responses.BaseResponse
	RequestId  string                            `json:"RequestId" xml:"RequestId"`
	TotalCount int64                             `json:"TotalCount" xml:"TotalCount"`
	PageNumber int                               `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                               `json:"PageSize" xml:"PageSize"`
	Clusters   ClustersInDescribeDeletedClusters `json:"Clusters" xml:"Clusters"`
}

// CreateDescribeDeletedClustersRequest creates a request to invoke DescribeDeletedClusters API
func CreateDescribeDeletedClustersRequest() (request *DescribeDeletedClustersRequest) {
	request = &DescribeDeletedClustersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cassandra", "2019-01-01", "DescribeDeletedClusters", "Cassandra", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDeletedClustersResponse creates a response to parse from DescribeDeletedClusters response
func CreateDescribeDeletedClustersResponse() (response *DescribeDeletedClustersResponse) {
	response = &DescribeDeletedClustersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
