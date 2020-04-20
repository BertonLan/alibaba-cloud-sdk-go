package mts

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

// AddMedia invokes the mts.AddMedia API synchronously
// api document: https://help.aliyun.com/api/mts/addmedia.html
func (client *Client) AddMedia(request *AddMediaRequest) (response *AddMediaResponse, err error) {
	response = CreateAddMediaResponse()
	err = client.DoAction(request, response)
	return
}

// AddMediaWithChan invokes the mts.AddMedia API asynchronously
// api document: https://help.aliyun.com/api/mts/addmedia.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddMediaWithChan(request *AddMediaRequest) (<-chan *AddMediaResponse, <-chan error) {
	responseChan := make(chan *AddMediaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddMedia(request)
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

// AddMediaWithCallback invokes the mts.AddMedia API asynchronously
// api document: https://help.aliyun.com/api/mts/addmedia.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddMediaWithCallback(request *AddMediaRequest, callback func(response *AddMediaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddMediaResponse
		var err error
		defer close(result)
		response, err = client.AddMedia(request)
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

// AddMediaRequest is the request struct for api AddMedia
type AddMediaRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Description           string           `position:"Query" name:"Description"`
	Title                 string           `position:"Query" name:"Title"`
	InputUnbind           requests.Boolean `position:"Query" name:"InputUnbind"`
	CoverURL              string           `position:"Query" name:"CoverURL"`
	CateId                requests.Integer `position:"Query" name:"CateId"`
	MediaWorkflowId       string           `position:"Query" name:"MediaWorkflowId"`
	MediaWorkflowUserData string           `position:"Query" name:"MediaWorkflowUserData"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OverrideParams        string           `position:"Query" name:"OverrideParams"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	Tags                  string           `position:"Query" name:"Tags"`
	FileURL               string           `position:"Query" name:"FileURL"`
}

// AddMediaResponse is the response struct for api AddMedia
type AddMediaResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Media     Media  `json:"Media" xml:"Media"`
}

// CreateAddMediaRequest creates a request to invoke AddMedia API
func CreateAddMediaRequest() (request *AddMediaRequest) {
	request = &AddMediaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "AddMedia", "", "")
	return
}

// CreateAddMediaResponse creates a response to parse from AddMedia response
func CreateAddMediaResponse() (response *AddMediaResponse) {
	response = &AddMediaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
