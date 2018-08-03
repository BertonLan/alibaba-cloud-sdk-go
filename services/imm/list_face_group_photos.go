package imm

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

// ListFaceGroupPhotos invokes the imm.ListFaceGroupPhotos API synchronously
// api document: https://help.aliyun.com/api/imm/listfacegroupphotos.html
func (client *Client) ListFaceGroupPhotos(request *ListFaceGroupPhotosRequest) (response *ListFaceGroupPhotosResponse, err error) {
	response = CreateListFaceGroupPhotosResponse()
	err = client.DoAction(request, response)
	return
}

// ListFaceGroupPhotosWithChan invokes the imm.ListFaceGroupPhotos API asynchronously
// api document: https://help.aliyun.com/api/imm/listfacegroupphotos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListFaceGroupPhotosWithChan(request *ListFaceGroupPhotosRequest) (<-chan *ListFaceGroupPhotosResponse, <-chan error) {
	responseChan := make(chan *ListFaceGroupPhotosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFaceGroupPhotos(request)
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

// ListFaceGroupPhotosWithCallback invokes the imm.ListFaceGroupPhotos API asynchronously
// api document: https://help.aliyun.com/api/imm/listfacegroupphotos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListFaceGroupPhotosWithCallback(request *ListFaceGroupPhotosRequest, callback func(response *ListFaceGroupPhotosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFaceGroupPhotosResponse
		var err error
		defer close(result)
		response, err = client.ListFaceGroupPhotos(request)
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

// ListFaceGroupPhotosRequest is the request struct for api ListFaceGroupPhotos
type ListFaceGroupPhotosRequest struct {
	*requests.RpcRequest
	MaxKeys requests.Integer `position:"Query" name:"MaxKeys"`
	Marker  string           `position:"Query" name:"Marker"`
	GroupId string           `position:"Query" name:"GroupId"`
	Project string           `position:"Query" name:"Project"`
	SetId   string           `position:"Query" name:"SetId"`
}

// ListFaceGroupPhotosResponse is the response struct for api ListFaceGroupPhotos
type ListFaceGroupPhotosResponse struct {
	*responses.BaseResponse
	RequestId  string       `json:"RequestId" xml:"RequestId"`
	NextMarker string       `json:"NextMarker" xml:"NextMarker"`
	Photos     []PhotosItem `json:"Photos" xml:"Photos"`
}

// CreateListFaceGroupPhotosRequest creates a request to invoke ListFaceGroupPhotos API
func CreateListFaceGroupPhotosRequest() (request *ListFaceGroupPhotosRequest) {
	request = &ListFaceGroupPhotosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "ListFaceGroupPhotos", "imm", "openAPI")
	return
}

// CreateListFaceGroupPhotosResponse creates a response to parse from ListFaceGroupPhotos response
func CreateListFaceGroupPhotosResponse() (response *ListFaceGroupPhotosResponse) {
	response = &ListFaceGroupPhotosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
