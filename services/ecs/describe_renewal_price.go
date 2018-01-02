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

func (client *Client) DescribeRenewalPrice(request *DescribeRenewalPriceRequest) (response *DescribeRenewalPriceResponse, err error) {
	response = CreateDescribeRenewalPriceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeRenewalPriceWithChan(request *DescribeRenewalPriceRequest) (<-chan *DescribeRenewalPriceResponse, <-chan error) {
	responseChan := make(chan *DescribeRenewalPriceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRenewalPrice(request)
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

func (client *Client) DescribeRenewalPriceWithCallback(request *DescribeRenewalPriceRequest, callback func(response *DescribeRenewalPriceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRenewalPriceResponse
		var err error
		defer close(result)
		response, err = client.DescribeRenewalPrice(request)
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

type DescribeRenewalPriceRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceType         string           `position:"Query" name:"ResourceType"`
	PriceUnit            string           `position:"Query" name:"PriceUnit"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Period               requests.Integer `position:"Query" name:"Period"`
	ResourceId           string           `position:"Query" name:"ResourceId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeRenewalPriceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	PriceInfo struct {
		Price struct {
			OriginalPrice float64 `json:"OriginalPrice" xml:"OriginalPrice"`
			DiscountPrice float64 `json:"DiscountPrice" xml:"DiscountPrice"`
			TradePrice    float64 `json:"TradePrice" xml:"TradePrice"`
			Currency      string  `json:"Currency" xml:"Currency"`
		} `json:"Price" xml:"Price"`
		Rules struct {
			Rule []struct {
				RuleId      int    `json:"RuleId" xml:"RuleId"`
				Description string `json:"Description" xml:"Description"`
			} `json:"Rule" xml:"Rule"`
		} `json:"Rules" xml:"Rules"`
	} `json:"PriceInfo" xml:"PriceInfo"`
}

func CreateDescribeRenewalPriceRequest() (request *DescribeRenewalPriceRequest) {
	request = &DescribeRenewalPriceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeRenewalPrice", "", "")
	return
}

func CreateDescribeRenewalPriceResponse() (response *DescribeRenewalPriceResponse) {
	response = &DescribeRenewalPriceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
