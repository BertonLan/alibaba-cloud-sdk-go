package dms_enterprise

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

// GetLogicDatabase invokes the dms_enterprise.GetLogicDatabase API synchronously
// api document: https://help.aliyun.com/api/dms-enterprise/getlogicdatabase.html
func (client *Client) GetLogicDatabase(request *GetLogicDatabaseRequest) (response *GetLogicDatabaseResponse, err error) {
	response = CreateGetLogicDatabaseResponse()
	err = client.DoAction(request, response)
	return
}

// GetLogicDatabaseWithChan invokes the dms_enterprise.GetLogicDatabase API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/getlogicdatabase.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetLogicDatabaseWithChan(request *GetLogicDatabaseRequest) (<-chan *GetLogicDatabaseResponse, <-chan error) {
	responseChan := make(chan *GetLogicDatabaseResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetLogicDatabase(request)
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

// GetLogicDatabaseWithCallback invokes the dms_enterprise.GetLogicDatabase API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/getlogicdatabase.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetLogicDatabaseWithCallback(request *GetLogicDatabaseRequest, callback func(response *GetLogicDatabaseResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetLogicDatabaseResponse
		var err error
		defer close(result)
		response, err = client.GetLogicDatabase(request)
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

// GetLogicDatabaseRequest is the request struct for api GetLogicDatabase
type GetLogicDatabaseRequest struct {
	*requests.RpcRequest
	DbId string           `position:"Query" name:"DbId"`
	Tid  requests.Integer `position:"Query" name:"Tid"`
}

// GetLogicDatabaseResponse is the response struct for api GetLogicDatabase
type GetLogicDatabaseResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	Success       bool          `json:"Success" xml:"Success"`
	ErrorMessage  string        `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode     string        `json:"ErrorCode" xml:"ErrorCode"`
	LogicDatabase LogicDatabase `json:"LogicDatabase" xml:"LogicDatabase"`
}

// CreateGetLogicDatabaseRequest creates a request to invoke GetLogicDatabase API
func CreateGetLogicDatabaseRequest() (request *GetLogicDatabaseRequest) {
	request = &GetLogicDatabaseRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetLogicDatabase", "", "")
	return
}

// CreateGetLogicDatabaseResponse creates a response to parse from GetLogicDatabase response
func CreateGetLogicDatabaseResponse() (response *GetLogicDatabaseResponse) {
	response = &GetLogicDatabaseResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}