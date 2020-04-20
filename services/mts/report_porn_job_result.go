package mts

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

// ReportPornJobResult invokes the mts.ReportPornJobResult API synchronously
// api document: https://help.aliyun.com/api/mts/reportpornjobresult.html
func (client *Client) ReportPornJobResult(request *ReportPornJobResultRequest) (response *ReportPornJobResultResponse, err error) {
	response = CreateReportPornJobResultResponse()
	err = client.DoAction(request, response)
	return
}

// ReportPornJobResultWithChan invokes the mts.ReportPornJobResult API asynchronously
// api document: https://help.aliyun.com/api/mts/reportpornjobresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReportPornJobResultWithChan(request *ReportPornJobResultRequest) (<-chan *ReportPornJobResultResponse, <-chan error) {
	responseChan := make(chan *ReportPornJobResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReportPornJobResult(request)
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

// ReportPornJobResultWithCallback invokes the mts.ReportPornJobResult API asynchronously
// api document: https://help.aliyun.com/api/mts/reportpornjobresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReportPornJobResultWithCallback(request *ReportPornJobResultRequest, callback func(response *ReportPornJobResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReportPornJobResultResponse
		var err error
		defer close(result)
		response, err = client.ReportPornJobResult(request)
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

// ReportPornJobResultRequest is the request struct for api ReportPornJobResult
type ReportPornJobResultRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JobId                string           `position:"Query" name:"JobId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Label                string           `position:"Query" name:"Label"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Detail               string           `position:"Query" name:"Detail"`
}

// ReportPornJobResultResponse is the response struct for api ReportPornJobResult
type ReportPornJobResultResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateReportPornJobResultRequest creates a request to invoke ReportPornJobResult API
func CreateReportPornJobResultRequest() (request *ReportPornJobResultRequest) {
	request = &ReportPornJobResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ReportPornJobResult", "", "")
	return
}

// CreateReportPornJobResultResponse creates a response to parse from ReportPornJobResult response
func CreateReportPornJobResultResponse() (response *ReportPornJobResultResponse) {
	response = &ReportPornJobResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
