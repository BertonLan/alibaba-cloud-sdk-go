package ess

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

func (client *Client) DescribeCapacityHistory(request *DescribeCapacityHistoryRequest) (response *DescribeCapacityHistoryResponse, err error) {
	response = CreateDescribeCapacityHistoryResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeCapacityHistoryWithChan(request *DescribeCapacityHistoryRequest) (<-chan *DescribeCapacityHistoryResponse, <-chan error) {
	responseChan := make(chan *DescribeCapacityHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCapacityHistory(request)
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

func (client *Client) DescribeCapacityHistoryWithCallback(request *DescribeCapacityHistoryRequest, callback func(response *DescribeCapacityHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCapacityHistoryResponse
		var err error
		defer close(result)
		response, err = client.DescribeCapacityHistory(request)
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

type DescribeCapacityHistoryRequest struct {
	*requests.RpcRequest
	EndTime              string           `position:"Query" name:"EndTime"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	StartTime            string           `position:"Query" name:"StartTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
}

type DescribeCapacityHistoryResponse struct {
	*responses.BaseResponse
	TotalCount           int `json:"TotalCount" xml:"TotalCount"`
	PageNumber           int `json:"PageNumber" xml:"PageNumber"`
	PageSize             int `json:"PageSize" xml:"PageSize"`
	CapacityHistoryItems struct {
		CapacityHistoryModel []struct {
			ScalingGroupId      string `json:"ScalingGroupId" xml:"ScalingGroupId"`
			TotalCapacity       int    `json:"TotalCapacity" xml:"TotalCapacity"`
			AttachedCapacity    int    `json:"AttachedCapacity" xml:"AttachedCapacity"`
			AutoCreatedCapacity int    `json:"AutoCreatedCapacity" xml:"AutoCreatedCapacity"`
			Timestamp           string `json:"Timestamp" xml:"Timestamp"`
		} `json:"CapacityHistoryModel" xml:"CapacityHistoryModel"`
	} `json:"CapacityHistoryItems" xml:"CapacityHistoryItems"`
}

func CreateDescribeCapacityHistoryRequest() (request *DescribeCapacityHistoryRequest) {
	request = &DescribeCapacityHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "DescribeCapacityHistory", "", "")
	return
}

func CreateDescribeCapacityHistoryResponse() (response *DescribeCapacityHistoryResponse) {
	response = &DescribeCapacityHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
