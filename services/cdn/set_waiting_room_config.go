package cdn

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

func (client *Client) SetWaitingRoomConfig(request *SetWaitingRoomConfigRequest) (response *SetWaitingRoomConfigResponse, err error) {
	response = CreateSetWaitingRoomConfigResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SetWaitingRoomConfigWithChan(request *SetWaitingRoomConfigRequest) (<-chan *SetWaitingRoomConfigResponse, <-chan error) {
	responseChan := make(chan *SetWaitingRoomConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetWaitingRoomConfig(request)
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

func (client *Client) SetWaitingRoomConfigWithCallback(request *SetWaitingRoomConfigRequest, callback func(response *SetWaitingRoomConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetWaitingRoomConfigResponse
		var err error
		defer close(result)
		response, err = client.SetWaitingRoomConfig(request)
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

type SetWaitingRoomConfigRequest struct {
	*requests.RpcRequest
	WaitUrl       string           `position:"Query" name:"WaitUrl"`
	WaitUri       string           `position:"Query" name:"WaitUri"`
	Version       string           `position:"Query" name:"Version"`
	DomainName    string           `position:"Query" name:"DomainName"`
	AllowPct      requests.Integer `position:"Query" name:"AllowPct"`
	MaxQps        requests.Integer `position:"Query" name:"MaxQps"`
	MaxTimeWait   requests.Integer `position:"Query" name:"MaxTimeWait"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	GapTime       requests.Integer `position:"Query" name:"GapTime"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type SetWaitingRoomConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateSetWaitingRoomConfigRequest() (request *SetWaitingRoomConfigRequest) {
	request = &SetWaitingRoomConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "SetWaitingRoomConfig", "", "")
	return
}

func CreateSetWaitingRoomConfigResponse() (response *SetWaitingRoomConfigResponse) {
	response = &SetWaitingRoomConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
