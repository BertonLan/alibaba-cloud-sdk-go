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

func (client *Client) DownloadVpnConnectionConfig(request *DownloadVpnConnectionConfigRequest) (response *DownloadVpnConnectionConfigResponse, err error) {
	response = CreateDownloadVpnConnectionConfigResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DownloadVpnConnectionConfigWithChan(request *DownloadVpnConnectionConfigRequest) (<-chan *DownloadVpnConnectionConfigResponse, <-chan error) {
	responseChan := make(chan *DownloadVpnConnectionConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DownloadVpnConnectionConfig(request)
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

func (client *Client) DownloadVpnConnectionConfigWithCallback(request *DownloadVpnConnectionConfigRequest, callback func(response *DownloadVpnConnectionConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DownloadVpnConnectionConfigResponse
		var err error
		defer close(result)
		response, err = client.DownloadVpnConnectionConfig(request)
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

type DownloadVpnConnectionConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	VpnConnectionId      string           `position:"Query" name:"VpnConnectionId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type DownloadVpnConnectionConfigResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	VpnConnectionConfig struct {
		LocalSubnet  string `json:"LocalSubnet" xml:"LocalSubnet"`
		RemoteSubnet string `json:"RemoteSubnet" xml:"RemoteSubnet"`
		Local        string `json:"Local" xml:"Local"`
		Remote       string `json:"Remote" xml:"Remote"`
		IkeConfig    struct {
			Psk         string `json:"Psk" xml:"Psk"`
			IkeVersion  string `json:"IkeVersion" xml:"IkeVersion"`
			IkeMode     string `json:"IkeMode" xml:"IkeMode"`
			IkeEncAlg   string `json:"IkeEncAlg" xml:"IkeEncAlg"`
			IkeAuthAlg  string `json:"IkeAuthAlg" xml:"IkeAuthAlg"`
			IkePfs      string `json:"IkePfs" xml:"IkePfs"`
			IkeLifetime int    `json:"IkeLifetime" xml:"IkeLifetime"`
			LocalId     string `json:"LocalId" xml:"LocalId"`
			RemoteId    string `json:"RemoteId" xml:"RemoteId"`
		} `json:"IkeConfig" xml:"IkeConfig"`
		IpsecConfig struct {
			IpsecEncAlg   string `json:"IpsecEncAlg" xml:"IpsecEncAlg"`
			IpsecAuthAlg  string `json:"IpsecAuthAlg" xml:"IpsecAuthAlg"`
			IpsecPfs      string `json:"IpsecPfs" xml:"IpsecPfs"`
			IpsecLifetime int    `json:"IpsecLifetime" xml:"IpsecLifetime"`
		} `json:"IpsecConfig" xml:"IpsecConfig"`
	} `json:"VpnConnectionConfig" xml:"VpnConnectionConfig"`
}

func CreateDownloadVpnConnectionConfigRequest() (request *DownloadVpnConnectionConfigRequest) {
	request = &DownloadVpnConnectionConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DownloadVpnConnectionConfig", "", "")
	return
}

func CreateDownloadVpnConnectionConfigResponse() (response *DownloadVpnConnectionConfigResponse) {
	response = &DownloadVpnConnectionConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
