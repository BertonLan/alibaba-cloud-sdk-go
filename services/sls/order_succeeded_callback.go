package sls

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

func (client *Client) OrderSucceededCallback(request *OrderSucceededCallbackRequest) (response *OrderSucceededCallbackResponse, err error) {
	response = CreateOrderSucceededCallbackResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) OrderSucceededCallbackWithChan(request *OrderSucceededCallbackRequest) (<-chan *OrderSucceededCallbackResponse, <-chan error) {
	responseChan := make(chan *OrderSucceededCallbackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OrderSucceededCallback(request)
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

func (client *Client) OrderSucceededCallbackWithCallback(request *OrderSucceededCallbackRequest, callback func(response *OrderSucceededCallbackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OrderSucceededCallbackResponse
		var err error
		defer close(result)
		response, err = client.OrderSucceededCallback(request)
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

type OrderSucceededCallbackRequest struct {
	*requests.RpcRequest
	Data string `position:"Body" name:"data"`
}

type OrderSucceededCallbackResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Data      string `json:"data" xml:"data"`
	Success   bool   `json:"success" xml:"success"`
	Synchro   bool   `json:"synchro" xml:"synchro"`
}

func CreateOrderSucceededCallbackRequest() (request *OrderSucceededCallbackRequest) {
	request = &OrderSucceededCallbackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sls", "2016-08-01", "OrderSucceededCallback", "", "")
	return
}

func CreateOrderSucceededCallbackResponse() (response *OrderSucceededCallbackResponse) {
	response = &OrderSucceededCallbackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
