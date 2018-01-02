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

func (client *Client) QueryCensorJobList(request *QueryCensorJobListRequest) (response *QueryCensorJobListResponse, err error) {
	response = CreateQueryCensorJobListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryCensorJobListWithChan(request *QueryCensorJobListRequest) (<-chan *QueryCensorJobListResponse, <-chan error) {
	responseChan := make(chan *QueryCensorJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCensorJobList(request)
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

func (client *Client) QueryCensorJobListWithCallback(request *QueryCensorJobListRequest, callback func(response *QueryCensorJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCensorJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryCensorJobList(request)
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

type QueryCensorJobListRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string           `position:"Query" name:"JobIds"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type QueryCensorJobListResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	NonExistIds struct {
		String []string `json:"String" xml:"String"`
	} `json:"NonExistIds" xml:"NonExistIds"`
	CensorJobList struct {
		CensorJob []struct {
			Id           string `json:"Id" xml:"Id"`
			UserData     string `json:"UserData" xml:"UserData"`
			PipelineId   string `json:"PipelineId" xml:"PipelineId"`
			State        string `json:"State" xml:"State"`
			Code         string `json:"Code" xml:"Code"`
			Message      string `json:"Message" xml:"Message"`
			CreationTime string `json:"CreationTime" xml:"CreationTime"`
			Input        struct {
				Bucket   string `json:"Bucket" xml:"Bucket"`
				Location string `json:"Location" xml:"Location"`
				Object   string `json:"Object" xml:"Object"`
			} `json:"Input" xml:"Input"`
			CensorConfig struct {
				Interval   string `json:"Interval" xml:"Interval"`
				BizType    string `json:"BizType" xml:"BizType"`
				OutputFile struct {
					Bucket   string `json:"Bucket" xml:"Bucket"`
					Location string `json:"Location" xml:"Location"`
					Object   string `json:"Object" xml:"Object"`
				} `json:"OutputFile" xml:"OutputFile"`
			} `json:"CensorConfig" xml:"CensorConfig"`
			CensorPornResult struct {
				Label           string `json:"Label" xml:"Label"`
				Suggestion      string `json:"Suggestion" xml:"Suggestion"`
				MaxScore        string `json:"MaxScore" xml:"MaxScore"`
				AverageScore    string `json:"AverageScore" xml:"AverageScore"`
				PornCounterList struct {
					Counter []struct {
						Count int    `json:"Count" xml:"Count"`
						Label string `json:"Label" xml:"Label"`
					} `json:"Counter" xml:"Counter"`
				} `json:"PornCounterList" xml:"PornCounterList"`
				PornTopList struct {
					Top []struct {
						Label     string `json:"Label" xml:"Label"`
						Score     string `json:"Score" xml:"Score"`
						Timestamp string `json:"Timestamp" xml:"Timestamp"`
						Index     string `json:"Index" xml:"Index"`
						Object    string `json:"Object" xml:"Object"`
					} `json:"Top" xml:"Top"`
				} `json:"PornTopList" xml:"PornTopList"`
			} `json:"CensorPornResult" xml:"CensorPornResult"`
			CensorTerrorismResult struct {
				Label                string `json:"Label" xml:"Label"`
				Suggestion           string `json:"Suggestion" xml:"Suggestion"`
				MaxScore             string `json:"MaxScore" xml:"MaxScore"`
				AverageScore         string `json:"AverageScore" xml:"AverageScore"`
				TerrorismCounterList struct {
					Counter []struct {
						Count int    `json:"Count" xml:"Count"`
						Label string `json:"Label" xml:"Label"`
					} `json:"Counter" xml:"Counter"`
				} `json:"TerrorismCounterList" xml:"TerrorismCounterList"`
				TerrorismTopList struct {
					Top []struct {
						Label     string `json:"Label" xml:"Label"`
						Score     string `json:"Score" xml:"Score"`
						Timestamp string `json:"Timestamp" xml:"Timestamp"`
						Index     string `json:"Index" xml:"Index"`
						Object    string `json:"Object" xml:"Object"`
					} `json:"Top" xml:"Top"`
				} `json:"TerrorismTopList" xml:"TerrorismTopList"`
			} `json:"CensorTerrorismResult" xml:"CensorTerrorismResult"`
		} `json:"CensorJob" xml:"CensorJob"`
	} `json:"CensorJobList" xml:"CensorJobList"`
}

func CreateQueryCensorJobListRequest() (request *QueryCensorJobListRequest) {
	request = &QueryCensorJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryCensorJobList", "", "")
	return
}

func CreateQueryCensorJobListResponse() (response *QueryCensorJobListResponse) {
	response = &QueryCensorJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
