package rds

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

func (client *Client) DescribeRealtimeDiagnoses(request *DescribeRealtimeDiagnosesRequest) (response *DescribeRealtimeDiagnosesResponse, err error) {
	response = CreateDescribeRealtimeDiagnosesResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeRealtimeDiagnosesWithChan(request *DescribeRealtimeDiagnosesRequest) (<-chan *DescribeRealtimeDiagnosesResponse, <-chan error) {
	responseChan := make(chan *DescribeRealtimeDiagnosesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRealtimeDiagnoses(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeRealtimeDiagnosesWithCallback(request *DescribeRealtimeDiagnosesRequest, callback func(response *DescribeRealtimeDiagnosesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRealtimeDiagnosesResponse
		var err error
		defer close(result)
		response, err = client.DescribeRealtimeDiagnoses(request)
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

type DescribeRealtimeDiagnosesRequest struct {
	*requests.RpcRequest
	EndTime              string           `position:"Query" name:"EndTime"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeRealtimeDiagnosesResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	Engine           string `json:"Engine" xml:"Engine"`
	TotalRecordCount int    `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       int    `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int    `json:"PageRecordCount" xml:"PageRecordCount"`
	Tasks            struct {
		RealtimeDiagnoseTasks []struct {
			CreateTime  string `json:"CreateTime" xml:"CreateTime"`
			TaskId      string `json:"TaskId" xml:"TaskId"`
			HealthScore string `json:"HealthScore" xml:"HealthScore"`
			Status      string `json:"Status" xml:"Status"`
		} `json:"RealtimeDiagnoseTasks" xml:"RealtimeDiagnoseTasks"`
	} `json:"Tasks" xml:"Tasks"`
}

func CreateDescribeRealtimeDiagnosesRequest() (request *DescribeRealtimeDiagnosesRequest) {
	request = &DescribeRealtimeDiagnosesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeRealtimeDiagnoses", "", "")
	return
}

func CreateDescribeRealtimeDiagnosesResponse() (response *DescribeRealtimeDiagnosesResponse) {
	response = &DescribeRealtimeDiagnosesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
