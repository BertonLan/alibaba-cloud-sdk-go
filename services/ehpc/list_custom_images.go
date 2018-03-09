package ehpc

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

// ListCustomImages invokes the ehpc.ListCustomImages API synchronously
// api document: https://help.aliyun.com/api/ehpc/listcustomimages.html
func (client *Client) ListCustomImages(request *ListCustomImagesRequest) (response *ListCustomImagesResponse, err error) {
	response = CreateListCustomImagesResponse()
	err = client.DoAction(request, response)
	return
}

// ListCustomImagesWithChan invokes the ehpc.ListCustomImages API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listcustomimages.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCustomImagesWithChan(request *ListCustomImagesRequest) (<-chan *ListCustomImagesResponse, <-chan error) {
	responseChan := make(chan *ListCustomImagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCustomImages(request)
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

// ListCustomImagesWithCallback invokes the ehpc.ListCustomImages API asynchronously
// api document: https://help.aliyun.com/api/ehpc/listcustomimages.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCustomImagesWithCallback(request *ListCustomImagesRequest, callback func(response *ListCustomImagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCustomImagesResponse
		var err error
		defer close(result)
		response, err = client.ListCustomImages(request)
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

// ListCustomImagesRequest is the request struct for api ListCustomImages
type ListCustomImagesRequest struct {
	*requests.RpcRequest
	ImageOwnerAlias string `position:"Query" name:"ImageOwnerAlias"`
	BaseOsTag       string `position:"Query" name:"BaseOsTag"`
}

// ListCustomImagesResponse is the response struct for api ListCustomImages
type ListCustomImagesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Images    Images `json:"Images" xml:"Images"`
}

// CreateListCustomImagesRequest creates a request to invoke ListCustomImages API
func CreateListCustomImagesRequest(request *ListCustomImagesRequest) {
	request = &ListCustomImagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2017-07-14", "ListCustomImages", "ehs", "openAPI")
	return
}

// CreateListCustomImagesResponse creates a response to parse from ListCustomImages response
func CreateListCustomImagesResponse() (response *ListCustomImagesResponse) {
	response = &ListCustomImagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}