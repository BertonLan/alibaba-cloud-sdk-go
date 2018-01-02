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

func (client *Client) SearchMedia(request *SearchMediaRequest) (response *SearchMediaResponse, err error) {
	response = CreateSearchMediaResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SearchMediaWithChan(request *SearchMediaRequest) (<-chan *SearchMediaResponse, <-chan error) {
	responseChan := make(chan *SearchMediaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchMedia(request)
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

func (client *Client) SearchMediaWithCallback(request *SearchMediaRequest, callback func(response *SearchMediaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchMediaResponse
		var err error
		defer close(result)
		response, err = client.SearchMedia(request)
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

type SearchMediaRequest struct {
	*requests.RpcRequest
	To                   string           `position:"Query" name:"To"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	CateId               string           `position:"Query" name:"CateId"`
	SortBy               string           `position:"Query" name:"SortBy"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	Tag                  string           `position:"Query" name:"Tag"`
	KeyWord              string           `position:"Query" name:"KeyWord"`
	From                 string           `position:"Query" name:"From"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Title                string           `position:"Query" name:"Title"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Description          string           `position:"Query" name:"Description"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type SearchMediaResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalNum   int    `json:"TotalNum" xml:"TotalNum"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	MediaList  struct {
		Media []struct {
			MediaId      string `json:"MediaId" xml:"MediaId"`
			Title        string `json:"Title" xml:"Title"`
			Description  string `json:"Description" xml:"Description"`
			CoverURL     string `json:"CoverURL" xml:"CoverURL"`
			CateId       int    `json:"CateId" xml:"CateId"`
			Duration     string `json:"Duration" xml:"Duration"`
			Format       string `json:"Format" xml:"Format"`
			Size         string `json:"Size" xml:"Size"`
			Bitrate      string `json:"Bitrate" xml:"Bitrate"`
			Width        string `json:"Width" xml:"Width"`
			Height       string `json:"Height" xml:"Height"`
			Fps          string `json:"Fps" xml:"Fps"`
			PublishState string `json:"PublishState" xml:"PublishState"`
			CreationTime string `json:"CreationTime" xml:"CreationTime"`
			Tags         struct {
				Tag []string `json:"Tag" xml:"Tag"`
			} `json:"Tags" xml:"Tags"`
			RunIdList struct {
				RunId []string `json:"RunId" xml:"RunId"`
			} `json:"RunIdList" xml:"RunIdList"`
			File struct {
				URL   string `json:"URL" xml:"URL"`
				State string `json:"State" xml:"State"`
			} `json:"File" xml:"File"`
		} `json:"Media" xml:"Media"`
	} `json:"MediaList" xml:"MediaList"`
}

func CreateSearchMediaRequest() (request *SearchMediaRequest) {
	request = &SearchMediaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SearchMedia", "", "")
	return
}

func CreateSearchMediaResponse() (response *SearchMediaResponse) {
	response = &SearchMediaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
