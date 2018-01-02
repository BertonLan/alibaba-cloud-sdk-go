package nas

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

func (client *Client) DescribeAccessGroups(request *DescribeAccessGroupsRequest) (response *DescribeAccessGroupsResponse, err error) {
	response = CreateDescribeAccessGroupsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeAccessGroupsWithChan(request *DescribeAccessGroupsRequest) (<-chan *DescribeAccessGroupsResponse, <-chan error) {
	responseChan := make(chan *DescribeAccessGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAccessGroups(request)
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

func (client *Client) DescribeAccessGroupsWithCallback(request *DescribeAccessGroupsRequest, callback func(response *DescribeAccessGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAccessGroupsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAccessGroups(request)
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

type DescribeAccessGroupsRequest struct {
	*requests.RpcRequest
	AccessGroupName string           `position:"Query" name:"AccessGroupName"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
}

type DescribeAccessGroupsResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	TotalCount   int    `json:"TotalCount" xml:"TotalCount"`
	PageSize     int    `json:"PageSize" xml:"PageSize"`
	PageNumber   int    `json:"PageNumber" xml:"PageNumber"`
	AccessGroups struct {
		AccessGroup []struct {
			AccessGroupName  string `json:"AccessGroupName" xml:"AccessGroupName"`
			AccessGroupType  string `json:"AccessGroupType" xml:"AccessGroupType"`
			RuleCount        int    `json:"RuleCount" xml:"RuleCount"`
			MountTargetCount int    `json:"MountTargetCount" xml:"MountTargetCount"`
			Description      string `json:"Description" xml:"Description"`
		} `json:"AccessGroup" xml:"AccessGroup"`
	} `json:"AccessGroups" xml:"AccessGroups"`
}

func CreateDescribeAccessGroupsRequest() (request *DescribeAccessGroupsRequest) {
	request = &DescribeAccessGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "DescribeAccessGroups", "", "")
	return
}

func CreateDescribeAccessGroupsResponse() (response *DescribeAccessGroupsResponse) {
	response = &DescribeAccessGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
