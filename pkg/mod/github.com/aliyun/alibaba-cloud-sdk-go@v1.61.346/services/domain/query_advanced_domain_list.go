package domain

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

// QueryAdvancedDomainList invokes the domain.QueryAdvancedDomainList API synchronously
// api document: https://help.aliyun.com/api/domain/queryadvanceddomainlist.html
func (client *Client) QueryAdvancedDomainList(request *QueryAdvancedDomainListRequest) (response *QueryAdvancedDomainListResponse, err error) {
	response = CreateQueryAdvancedDomainListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryAdvancedDomainListWithChan invokes the domain.QueryAdvancedDomainList API asynchronously
// api document: https://help.aliyun.com/api/domain/queryadvanceddomainlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryAdvancedDomainListWithChan(request *QueryAdvancedDomainListRequest) (<-chan *QueryAdvancedDomainListResponse, <-chan error) {
	responseChan := make(chan *QueryAdvancedDomainListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryAdvancedDomainList(request)
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

// QueryAdvancedDomainListWithCallback invokes the domain.QueryAdvancedDomainList API asynchronously
// api document: https://help.aliyun.com/api/domain/queryadvanceddomainlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryAdvancedDomainListWithCallback(request *QueryAdvancedDomainListRequest, callback func(response *QueryAdvancedDomainListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryAdvancedDomainListResponse
		var err error
		defer close(result)
		response, err = client.QueryAdvancedDomainList(request)
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

// QueryAdvancedDomainListRequest is the request struct for api QueryAdvancedDomainList
type QueryAdvancedDomainListRequest struct {
	*requests.RpcRequest
	ProductDomainType     string           `position:"Query" name:"ProductDomainType"`
	PageNum               requests.Integer `position:"Query" name:"PageNum"`
	Excluded              string           `position:"Query" name:"Excluded"`
	StartLength           requests.Integer `position:"Query" name:"StartLength"`
	ExcludedSuffix        requests.Boolean `position:"Query" name:"ExcludedSuffix"`
	PageSize              requests.Integer `position:"Query" name:"PageSize"`
	Lang                  string           `position:"Query" name:"Lang"`
	ExcludedPrefix        requests.Boolean `position:"Query" name:"ExcludedPrefix"`
	KeyWord               string           `position:"Query" name:"KeyWord"`
	ProductDomainTypeSort requests.Boolean `position:"Query" name:"ProductDomainTypeSort"`
	EndExpirationDate     requests.Integer `position:"Query" name:"EndExpirationDate"`
	Suffixs               string           `position:"Query" name:"Suffixs"`
	DomainNameSort        requests.Boolean `position:"Query" name:"DomainNameSort"`
	ExpirationDateSort    requests.Boolean `position:"Query" name:"ExpirationDateSort"`
	StartExpirationDate   requests.Integer `position:"Query" name:"StartExpirationDate"`
	DomainStatus          requests.Integer `position:"Query" name:"DomainStatus"`
	DomainGroupId         requests.Integer `position:"Query" name:"DomainGroupId"`
	KeyWordSuffix         requests.Boolean `position:"Query" name:"KeyWordSuffix"`
	KeyWordPrefix         requests.Boolean `position:"Query" name:"KeyWordPrefix"`
	TradeType             requests.Integer `position:"Query" name:"TradeType"`
	EndRegistrationDate   requests.Integer `position:"Query" name:"EndRegistrationDate"`
	Form                  requests.Integer `position:"Query" name:"Form"`
	UserClientIp          string           `position:"Query" name:"UserClientIp"`
	RegistrationDateSort  requests.Boolean `position:"Query" name:"RegistrationDateSort"`
	StartRegistrationDate requests.Integer `position:"Query" name:"StartRegistrationDate"`
	EndLength             requests.Integer `position:"Query" name:"EndLength"`
}

// QueryAdvancedDomainListResponse is the response struct for api QueryAdvancedDomainList
type QueryAdvancedDomainListResponse struct {
	*responses.BaseResponse
	RequestId      string                        `json:"RequestId" xml:"RequestId"`
	TotalItemNum   int                           `json:"TotalItemNum" xml:"TotalItemNum"`
	CurrentPageNum int                           `json:"CurrentPageNum" xml:"CurrentPageNum"`
	TotalPageNum   int                           `json:"TotalPageNum" xml:"TotalPageNum"`
	PageSize       int                           `json:"PageSize" xml:"PageSize"`
	PrePage        bool                          `json:"PrePage" xml:"PrePage"`
	NextPage       bool                          `json:"NextPage" xml:"NextPage"`
	Data           DataInQueryAdvancedDomainList `json:"Data" xml:"Data"`
}

// CreateQueryAdvancedDomainListRequest creates a request to invoke QueryAdvancedDomainList API
func CreateQueryAdvancedDomainListRequest() (request *QueryAdvancedDomainListRequest) {
	request = &QueryAdvancedDomainListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryAdvancedDomainList", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryAdvancedDomainListResponse creates a response to parse from QueryAdvancedDomainList response
func CreateQueryAdvancedDomainListResponse() (response *QueryAdvancedDomainListResponse) {
	response = &QueryAdvancedDomainListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
