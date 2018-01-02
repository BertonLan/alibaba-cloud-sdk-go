package vpc

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

func (client *Client) EnablePhysicalConnection(request *EnablePhysicalConnectionRequest) (response *EnablePhysicalConnectionResponse, err error) {
	response = CreateEnablePhysicalConnectionResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) EnablePhysicalConnectionWithChan(request *EnablePhysicalConnectionRequest) (<-chan *EnablePhysicalConnectionResponse, <-chan error) {
	responseChan := make(chan *EnablePhysicalConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnablePhysicalConnection(request)
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

func (client *Client) EnablePhysicalConnectionWithCallback(request *EnablePhysicalConnectionRequest, callback func(response *EnablePhysicalConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnablePhysicalConnectionResponse
		var err error
		defer close(result)
		response, err = client.EnablePhysicalConnection(request)
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

type EnablePhysicalConnectionRequest struct {
	*requests.RpcRequest
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PhysicalConnectionId string           `position:"Query" name:"PhysicalConnectionId"`
}

type EnablePhysicalConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateEnablePhysicalConnectionRequest() (request *EnablePhysicalConnectionRequest) {
	request = &EnablePhysicalConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "EnablePhysicalConnection", "", "")
	return
}

func CreateEnablePhysicalConnectionResponse() (response *EnablePhysicalConnectionResponse) {
	response = &EnablePhysicalConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
