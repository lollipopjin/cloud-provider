package rosapi

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

// Update invokes the ros.Update API synchronously
// api document: https://help.aliyun.com/api/ros/updatestack.html
func (client *Client) UpdateStack(request *UpdateStackRequest) (response *UpdateStackResponse, err error) {
	response = CreateUpdateStackResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateStackWithChan invokes the ros.Update API asynchronously
// api document: https://help.aliyun.com/api/ros/updatestack.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateStackWithChan(request *UpdateStackRequest) (<-chan *UpdateStackResponse, <-chan error) {
	responseChan := make(chan *UpdateStackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateStack(request)
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

// UpdateStackWithCallback invokes the ros.Update API asynchronously
// api document: https://help.aliyun.com/api/ros/updatestack.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateStackWithCallback(request *UpdateStackRequest, callback func(response *UpdateStackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateStackResponse
		var err error
		defer close(result)
		response, err = client.UpdateStack(request)
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

// UpdateStackRequest is the request struct for api Update
type UpdateStackRequest struct {
	*requests.RpcRequest
	EnableRecover               requests.Boolean         `position:"Query" name:"EnableRecover"`
	TimeoutInMinutes            requests.Integer         `position:"Query" name:"TimeoutInMinutes"`
	StackPolicyDuringUpdateBody string                   `position:"Query" name:"StackPolicyDuringUpdateBody"`
	DisableRollback             requests.Boolean         `position:"Query" name:"DisableRollback"`
	Parameters                  *[]UpdateStackParameters `position:"Query" name:"Parameters"  type:"Repeated"`
	ClientToken                 string                   `position:"Query" name:"ClientToken"`
	TemplateBody                string                   `position:"Query" name:"TemplateBody"`
	StackId                     string                   `position:"Query" name:"StackId"`
	TemplateURL                 string                   `position:"Query" name:"TemplateURL"`
	StackPolicyBody             string                   `position:"Query" name:"StackPolicyBody"`
	StackPolicyDuringUpdateURL  string                   `position:"Query" name:"StackPolicyDuringUpdateURL"`
	UpdateAllowPolicy           string                   `position:"Query" name:"UpdateAllowPolicy"`
	UsePreviousParameters       requests.Boolean         `position:"Query" name:"UsePreviousParameters"`
	StackPolicyURL              string                   `position:"Query" name:"StackPolicyURL"`
}

// UpdateStackParameters is a repeated param struct in UpdateStackRequest
type UpdateStackParameters struct {
	ParameterValue string `name:"ParameterValue"`
	ParameterKey   string `name:"ParameterKey"`
}

// UpdateStackResponse is the response struct for api Update
type UpdateStackResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	StackId   string `json:"StackId" xml:"StackId"`
}

// CreateUpdateStackRequest creates a request to invoke Update API
func CreateUpdateStackRequest() (request *UpdateStackRequest) {
	request = &UpdateStackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ROS", "2019-09-10", "UpdateStack", "ros", "openAPI")
	return
}

// CreateUpdateStackResponse creates a response to parse from Update response
func CreateUpdateStackResponse() (response *UpdateStackResponse) {
	response = &UpdateStackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}