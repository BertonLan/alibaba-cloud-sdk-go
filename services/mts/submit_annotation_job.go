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

// SubmitAnnotationJob invokes the mts.SubmitAnnotationJob API synchronously
// api document: https://help.aliyun.com/api/mts/submitannotationjob.html
func (client *Client) SubmitAnnotationJob(request *SubmitAnnotationJobRequest) (response *SubmitAnnotationJobResponse, err error) {
	response = CreateSubmitAnnotationJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitAnnotationJobWithChan invokes the mts.SubmitAnnotationJob API asynchronously
// api document: https://help.aliyun.com/api/mts/submitannotationjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitAnnotationJobWithChan(request *SubmitAnnotationJobRequest) (<-chan *SubmitAnnotationJobResponse, <-chan error) {
	responseChan := make(chan *SubmitAnnotationJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitAnnotationJob(request)
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

// SubmitAnnotationJobWithCallback invokes the mts.SubmitAnnotationJob API asynchronously
// api document: https://help.aliyun.com/api/mts/submitannotationjob.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SubmitAnnotationJobWithCallback(request *SubmitAnnotationJobRequest, callback func(response *SubmitAnnotationJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitAnnotationJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitAnnotationJob(request)
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

// SubmitAnnotationJobRequest is the request struct for api SubmitAnnotationJob
type SubmitAnnotationJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AnnotationConfig     string           `position:"Query" name:"AnnotationConfig"`
	UserData             string           `position:"Query" name:"UserData"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
	Input                string           `position:"Query" name:"Input"`
}

// SubmitAnnotationJobResponse is the response struct for api SubmitAnnotationJob
type SubmitAnnotationJobResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateSubmitAnnotationJobRequest creates a request to invoke SubmitAnnotationJob API
func CreateSubmitAnnotationJobRequest() (request *SubmitAnnotationJobRequest) {
	request = &SubmitAnnotationJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SubmitAnnotationJob", "", "")
	return
}

// CreateSubmitAnnotationJobResponse creates a response to parse from SubmitAnnotationJob response
func CreateSubmitAnnotationJobResponse() (response *SubmitAnnotationJobResponse) {
	response = &SubmitAnnotationJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
