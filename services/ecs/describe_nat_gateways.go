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

func (client *Client) DescribeNatGateways(request *DescribeNatGatewaysRequest) (response *DescribeNatGatewaysResponse, err error) {
	response = CreateDescribeNatGatewaysResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeNatGatewaysWithChan(request *DescribeNatGatewaysRequest) (<-chan *DescribeNatGatewaysResponse, <-chan error) {
	responseChan := make(chan *DescribeNatGatewaysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeNatGateways(request)
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

func (client *Client) DescribeNatGatewaysWithCallback(request *DescribeNatGatewaysRequest, callback func(response *DescribeNatGatewaysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeNatGatewaysResponse
		var err error
		defer close(result)
		response, err = client.DescribeNatGateways(request)
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

type DescribeNatGatewaysRequest struct {
	*requests.RpcRequest
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	VpcId                string           `position:"Query" name:"VpcId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	NatGatewayId         string           `position:"Query" name:"NatGatewayId"`
}

type DescribeNatGatewaysResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	TotalCount  int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber  int    `json:"PageNumber" xml:"PageNumber"`
	PageSize    int    `json:"PageSize" xml:"PageSize"`
	NatGateways struct {
		NatGateway []struct {
			NatGatewayId       string `json:"NatGatewayId" xml:"NatGatewayId"`
			RegionId           string `json:"RegionId" xml:"RegionId"`
			Name               string `json:"Name" xml:"Name"`
			Description        string `json:"Description" xml:"Description"`
			VpcId              string `json:"VpcId" xml:"VpcId"`
			Spec               string `json:"Spec" xml:"Spec"`
			InstanceChargeType string `json:"InstanceChargeType" xml:"InstanceChargeType"`
			BusinessStatus     string `json:"BusinessStatus" xml:"BusinessStatus"`
			CreationTime       string `json:"CreationTime" xml:"CreationTime"`
			Status             string `json:"Status" xml:"Status"`
			ForwardTableIds    struct {
				ForwardTableId []string `json:"ForwardTableId" xml:"ForwardTableId"`
			} `json:"ForwardTableIds" xml:"ForwardTableIds"`
			BandwidthPackageIds struct {
				BandwidthPackageId []string `json:"BandwidthPackageId" xml:"BandwidthPackageId"`
			} `json:"BandwidthPackageIds" xml:"BandwidthPackageIds"`
		} `json:"NatGateway" xml:"NatGateway"`
	} `json:"NatGateways" xml:"NatGateways"`
}

func CreateDescribeNatGatewaysRequest() (request *DescribeNatGatewaysRequest) {
	request = &DescribeNatGatewaysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeNatGateways", "", "")
	return
}

func CreateDescribeNatGatewaysResponse() (response *DescribeNatGatewaysResponse) {
	response = &DescribeNatGatewaysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
