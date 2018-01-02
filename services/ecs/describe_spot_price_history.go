package ecs

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

func (client *Client) DescribeSpotPriceHistory(request *DescribeSpotPriceHistoryRequest) (response *DescribeSpotPriceHistoryResponse, err error) {
	response = CreateDescribeSpotPriceHistoryResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeSpotPriceHistoryWithChan(request *DescribeSpotPriceHistoryRequest) (<-chan *DescribeSpotPriceHistoryResponse, <-chan error) {
	responseChan := make(chan *DescribeSpotPriceHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSpotPriceHistory(request)
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

func (client *Client) DescribeSpotPriceHistoryWithCallback(request *DescribeSpotPriceHistoryRequest, callback func(response *DescribeSpotPriceHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSpotPriceHistoryResponse
		var err error
		defer close(result)
		response, err = client.DescribeSpotPriceHistory(request)
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

type DescribeSpotPriceHistoryRequest struct {
	*requests.RpcRequest
	EndTime              string           `position:"Query" name:"EndTime"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
	OSType               string           `position:"Query" name:"OSType"`
	NetworkType          string           `position:"Query" name:"NetworkType"`
	InstanceType         string           `position:"Query" name:"InstanceType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Offset               requests.Integer `position:"Query" name:"Offset"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	StartTime            string           `position:"Query" name:"StartTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	IoOptimized          string           `position:"Query" name:"IoOptimized"`
}

type DescribeSpotPriceHistoryResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	NextOffset int    `json:"NextOffset" xml:"NextOffset"`
	Currency   string `json:"Currency" xml:"Currency"`
	SpotPrices struct {
		SpotPriceType []struct {
			ZoneId       string  `json:"ZoneId" xml:"ZoneId"`
			InstanceType string  `json:"InstanceType" xml:"InstanceType"`
			IoOptimized  string  `json:"IoOptimized" xml:"IoOptimized"`
			Timestamp    string  `json:"Timestamp" xml:"Timestamp"`
			NetworkType  string  `json:"NetworkType" xml:"NetworkType"`
			SpotPrice    float64 `json:"SpotPrice" xml:"SpotPrice"`
			OriginPrice  float64 `json:"OriginPrice" xml:"OriginPrice"`
		} `json:"SpotPriceType" xml:"SpotPriceType"`
	} `json:"SpotPrices" xml:"SpotPrices"`
}

func CreateDescribeSpotPriceHistoryRequest() (request *DescribeSpotPriceHistoryRequest) {
	request = &DescribeSpotPriceHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSpotPriceHistory", "", "")
	return
}

func CreateDescribeSpotPriceHistoryResponse() (response *DescribeSpotPriceHistoryResponse) {
	response = &DescribeSpotPriceHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
