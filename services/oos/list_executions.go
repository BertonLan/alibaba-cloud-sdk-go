package oos

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

// ListExecutions invokes the oos.ListExecutions API synchronously
// api document: https://help.aliyun.com/api/oos/listexecutions.html
func (client *Client) ListExecutions(request *ListExecutionsRequest) (response *ListExecutionsResponse, err error) {
	response = CreateListExecutionsResponse()
	err = client.DoAction(request, response)
	return
}

// ListExecutionsWithChan invokes the oos.ListExecutions API asynchronously
// api document: https://help.aliyun.com/api/oos/listexecutions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListExecutionsWithChan(request *ListExecutionsRequest) (<-chan *ListExecutionsResponse, <-chan error) {
	responseChan := make(chan *ListExecutionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListExecutions(request)
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

// ListExecutionsWithCallback invokes the oos.ListExecutions API asynchronously
// api document: https://help.aliyun.com/api/oos/listexecutions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListExecutionsWithCallback(request *ListExecutionsRequest, callback func(response *ListExecutionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListExecutionsResponse
		var err error
		defer close(result)
		response, err = client.ListExecutions(request)
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

// ListExecutionsRequest is the request struct for api ListExecutions
type ListExecutionsRequest struct {
	*requests.RpcRequest
	ExecutedBy            string                 `position:"Query" name:"ExecutedBy"`
	IncludeChildExecution requests.Boolean       `position:"Query" name:"IncludeChildExecution"`
	Mode                  string                 `position:"Query" name:"Mode"`
	ExecutionId           string                 `position:"Query" name:"ExecutionId"`
	RamRole               string                 `position:"Query" name:"RamRole"`
	NextToken             string                 `position:"Query" name:"NextToken"`
	TemplateName          string                 `position:"Query" name:"TemplateName"`
	EndDateBefore         string                 `position:"Query" name:"EndDateBefore"`
	SortOrder             string                 `position:"Query" name:"SortOrder"`
	StartDateAfter        string                 `position:"Query" name:"StartDateAfter"`
	StartDateBefore       string                 `position:"Query" name:"StartDateBefore"`
	Tags                  map[string]interface{} `position:"Query" name:"Tags"`
	ParentExecutionId     string                 `position:"Query" name:"ParentExecutionId"`
	EndDateAfter          string                 `position:"Query" name:"EndDateAfter"`
	MaxResults            requests.Integer       `position:"Query" name:"MaxResults"`
	SortField             string                 `position:"Query" name:"SortField"`
	Category              string                 `position:"Query" name:"Category"`
	Status                string                 `position:"Query" name:"Status"`
}

// ListExecutionsResponse is the response struct for api ListExecutions
type ListExecutionsResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	MaxResults int         `json:"MaxResults" xml:"MaxResults"`
	NextToken  string      `json:"NextToken" xml:"NextToken"`
	Executions []Execution `json:"Executions" xml:"Executions"`
}

// CreateListExecutionsRequest creates a request to invoke ListExecutions API
func CreateListExecutionsRequest() (request *ListExecutionsRequest) {
	request = &ListExecutionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "ListExecutions", "oos", "openAPI")
	return
}

// CreateListExecutionsResponse creates a response to parse from ListExecutions response
func CreateListExecutionsResponse() (response *ListExecutionsResponse) {
	response = &ListExecutionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
