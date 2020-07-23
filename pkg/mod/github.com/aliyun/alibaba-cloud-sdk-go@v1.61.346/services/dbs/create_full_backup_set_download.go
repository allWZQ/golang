package dbs

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

// CreateFullBackupSetDownload invokes the dbs.CreateFullBackupSetDownload API synchronously
// api document: https://help.aliyun.com/api/dbs/createfullbackupsetdownload.html
func (client *Client) CreateFullBackupSetDownload(request *CreateFullBackupSetDownloadRequest) (response *CreateFullBackupSetDownloadResponse, err error) {
	response = CreateCreateFullBackupSetDownloadResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFullBackupSetDownloadWithChan invokes the dbs.CreateFullBackupSetDownload API asynchronously
// api document: https://help.aliyun.com/api/dbs/createfullbackupsetdownload.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateFullBackupSetDownloadWithChan(request *CreateFullBackupSetDownloadRequest) (<-chan *CreateFullBackupSetDownloadResponse, <-chan error) {
	responseChan := make(chan *CreateFullBackupSetDownloadResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFullBackupSetDownload(request)
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

// CreateFullBackupSetDownloadWithCallback invokes the dbs.CreateFullBackupSetDownload API asynchronously
// api document: https://help.aliyun.com/api/dbs/createfullbackupsetdownload.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateFullBackupSetDownloadWithCallback(request *CreateFullBackupSetDownloadRequest, callback func(response *CreateFullBackupSetDownloadResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFullBackupSetDownloadResponse
		var err error
		defer close(result)
		response, err = client.CreateFullBackupSetDownload(request)
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

// CreateFullBackupSetDownloadRequest is the request struct for api CreateFullBackupSetDownload
type CreateFullBackupSetDownloadRequest struct {
	*requests.RpcRequest
	ClientToken         string `position:"Query" name:"ClientToken"`
	BackupSetId         string `position:"Query" name:"BackupSetId"`
	OwnerId             string `position:"Query" name:"OwnerId"`
	BackupSetDataFormat string `position:"Query" name:"BackupSetDataFormat"`
}

// CreateFullBackupSetDownloadResponse is the response struct for api CreateFullBackupSetDownload
type CreateFullBackupSetDownloadResponse struct {
	*responses.BaseResponse
	Success                 bool   `json:"Success" xml:"Success"`
	ErrCode                 string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage              string `json:"ErrMessage" xml:"ErrMessage"`
	HttpStatusCode          int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId               string `json:"RequestId" xml:"RequestId"`
	BackupSetDownloadTaskId string `json:"BackupSetDownloadTaskId" xml:"BackupSetDownloadTaskId"`
}

// CreateCreateFullBackupSetDownloadRequest creates a request to invoke CreateFullBackupSetDownload API
func CreateCreateFullBackupSetDownloadRequest() (request *CreateFullBackupSetDownloadRequest) {
	request = &CreateFullBackupSetDownloadRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "CreateFullBackupSetDownload", "cbs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateFullBackupSetDownloadResponse creates a response to parse from CreateFullBackupSetDownload response
func CreateCreateFullBackupSetDownloadResponse() (response *CreateFullBackupSetDownloadResponse) {
	response = &CreateFullBackupSetDownloadResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}