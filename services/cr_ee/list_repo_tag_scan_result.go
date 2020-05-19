package cr_ee

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

// ListRepoTagScanResult invokes the cr.ListRepoTagScanResult API synchronously
// api document: https://help.aliyun.com/api/cr/listrepotagscanresult.html
func (client *Client) ListRepoTagScanResult(request *ListRepoTagScanResultRequest) (response *ListRepoTagScanResultResponse, err error) {
	response = CreateListRepoTagScanResultResponse()
	err = client.DoAction(request, response)
	return
}

// ListRepoTagScanResultWithChan invokes the cr.ListRepoTagScanResult API asynchronously
// api document: https://help.aliyun.com/api/cr/listrepotagscanresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRepoTagScanResultWithChan(request *ListRepoTagScanResultRequest) (<-chan *ListRepoTagScanResultResponse, <-chan error) {
	responseChan := make(chan *ListRepoTagScanResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRepoTagScanResult(request)
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

// ListRepoTagScanResultWithCallback invokes the cr.ListRepoTagScanResult API asynchronously
// api document: https://help.aliyun.com/api/cr/listrepotagscanresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRepoTagScanResultWithCallback(request *ListRepoTagScanResultRequest, callback func(response *ListRepoTagScanResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRepoTagScanResultResponse
		var err error
		defer close(result)
		response, err = client.ListRepoTagScanResult(request)
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

// ListRepoTagScanResultRequest is the request struct for api ListRepoTagScanResult
type ListRepoTagScanResultRequest struct {
	*requests.RpcRequest
	Severity   string           `position:"Query" name:"Severity"`
	RepoId     string           `position:"Query" name:"RepoId"`
	ScanTaskId string           `position:"Query" name:"ScanTaskId"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageNo     requests.Integer `position:"Query" name:"PageNo"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Tag        string           `position:"Query" name:"Tag"`
}

// ListRepoTagScanResultResponse is the response struct for api ListRepoTagScanResult
type ListRepoTagScanResultResponse struct {
	*responses.BaseResponse
	ListRepoTagScanResultIsSuccess bool                  `json:"IsSuccess" xml:"IsSuccess"`
	Code                           string                `json:"Code" xml:"Code"`
	RequestId                      string                `json:"RequestId" xml:"RequestId"`
	PageNo                         int                   `json:"PageNo" xml:"PageNo"`
	PageSize                       int                   `json:"PageSize" xml:"PageSize"`
	TotalCount                     int                   `json:"TotalCount" xml:"TotalCount"`
	Vulnerabilities                []VulnerabilitiesItem `json:"Vulnerabilities" xml:"Vulnerabilities"`
}

// CreateListRepoTagScanResultRequest creates a request to invoke ListRepoTagScanResult API
func CreateListRepoTagScanResultRequest() (request *ListRepoTagScanResultRequest) {
	request = &ListRepoTagScanResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cr", "2018-12-01", "ListRepoTagScanResult", "cr", "openAPI")
	return
}

// CreateListRepoTagScanResultResponse creates a response to parse from ListRepoTagScanResult response
func CreateListRepoTagScanResultResponse() (response *ListRepoTagScanResultResponse) {
	response = &ListRepoTagScanResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
