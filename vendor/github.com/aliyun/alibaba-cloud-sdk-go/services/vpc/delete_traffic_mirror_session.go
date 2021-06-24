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

// DeleteTrafficMirrorSession invokes the vpc.DeleteTrafficMirrorSession API synchronously
func (client *Client) DeleteTrafficMirrorSession(request *DeleteTrafficMirrorSessionRequest) (response *DeleteTrafficMirrorSessionResponse, err error) {
	response = CreateDeleteTrafficMirrorSessionResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTrafficMirrorSessionWithChan invokes the vpc.DeleteTrafficMirrorSession API asynchronously
func (client *Client) DeleteTrafficMirrorSessionWithChan(request *DeleteTrafficMirrorSessionRequest) (<-chan *DeleteTrafficMirrorSessionResponse, <-chan error) {
	responseChan := make(chan *DeleteTrafficMirrorSessionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTrafficMirrorSession(request)
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

// DeleteTrafficMirrorSessionWithCallback invokes the vpc.DeleteTrafficMirrorSession API asynchronously
func (client *Client) DeleteTrafficMirrorSessionWithCallback(request *DeleteTrafficMirrorSessionRequest, callback func(response *DeleteTrafficMirrorSessionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTrafficMirrorSessionResponse
		var err error
		defer close(result)
		response, err = client.DeleteTrafficMirrorSession(request)
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

// DeleteTrafficMirrorSessionRequest is the request struct for api DeleteTrafficMirrorSession
type DeleteTrafficMirrorSessionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken            string           `position:"Query" name:"ClientToken"`
	DryRun                 requests.Boolean `position:"Query" name:"DryRun"`
	TrafficMirrorSessionId string           `position:"Query" name:"TrafficMirrorSessionId"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteTrafficMirrorSessionResponse is the response struct for api DeleteTrafficMirrorSession
type DeleteTrafficMirrorSessionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteTrafficMirrorSessionRequest creates a request to invoke DeleteTrafficMirrorSession API
func CreateDeleteTrafficMirrorSessionRequest() (request *DeleteTrafficMirrorSessionRequest) {
	request = &DeleteTrafficMirrorSessionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DeleteTrafficMirrorSession", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteTrafficMirrorSessionResponse creates a response to parse from DeleteTrafficMirrorSession response
func CreateDeleteTrafficMirrorSessionResponse() (response *DeleteTrafficMirrorSessionResponse) {
	response = &DeleteTrafficMirrorSessionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}