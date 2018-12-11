package baas

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

// DescribeOrganizationUserCerts invokes the baas.DescribeOrganizationUserCerts API synchronously
// api document: https://help.aliyun.com/api/baas/describeorganizationusercerts.html
func (client *Client) DescribeOrganizationUserCerts(request *DescribeOrganizationUserCertsRequest) (response *DescribeOrganizationUserCertsResponse, err error) {
	response = CreateDescribeOrganizationUserCertsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOrganizationUserCertsWithChan invokes the baas.DescribeOrganizationUserCerts API asynchronously
// api document: https://help.aliyun.com/api/baas/describeorganizationusercerts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOrganizationUserCertsWithChan(request *DescribeOrganizationUserCertsRequest) (<-chan *DescribeOrganizationUserCertsResponse, <-chan error) {
	responseChan := make(chan *DescribeOrganizationUserCertsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOrganizationUserCerts(request)
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

// DescribeOrganizationUserCertsWithCallback invokes the baas.DescribeOrganizationUserCerts API asynchronously
// api document: https://help.aliyun.com/api/baas/describeorganizationusercerts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeOrganizationUserCertsWithCallback(request *DescribeOrganizationUserCertsRequest, callback func(response *DescribeOrganizationUserCertsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOrganizationUserCertsResponse
		var err error
		defer close(result)
		response, err = client.DescribeOrganizationUserCerts(request)
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

// DescribeOrganizationUserCertsRequest is the request struct for api DescribeOrganizationUserCerts
type DescribeOrganizationUserCertsRequest struct {
	*requests.RpcRequest
	OrganizationId string `position:"Body" name:"OrganizationId"`
	Username       string `position:"Body" name:"Username"`
}

// DescribeOrganizationUserCertsResponse is the response struct for api DescribeOrganizationUserCerts
type DescribeOrganizationUserCertsResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      int    `json:"ErrorCode" xml:"ErrorCode"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	Result         Result `json:"Result" xml:"Result"`
}

// CreateDescribeOrganizationUserCertsRequest creates a request to invoke DescribeOrganizationUserCerts API
func CreateDescribeOrganizationUserCertsRequest() (request *DescribeOrganizationUserCertsRequest) {
	request = &DescribeOrganizationUserCertsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-07-31", "DescribeOrganizationUserCerts", "", "")
	return
}

// CreateDescribeOrganizationUserCertsResponse creates a response to parse from DescribeOrganizationUserCerts response
func CreateDescribeOrganizationUserCertsResponse() (response *DescribeOrganizationUserCertsResponse) {
	response = &DescribeOrganizationUserCertsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}