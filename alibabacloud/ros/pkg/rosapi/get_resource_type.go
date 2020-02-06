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

// GetResourceType invokes the ros.GetResourceType API synchronously
// api document: https://help.aliyun.com/api/ros/getresourcetype.html
func (client *Client) GetResourceType(request *GetResourceTypeRequest) (response *GetResourceTypeResponse, err error) {
	response = CreateGetResourceTypeResponse()
	err = client.DoAction(request, response)
	return
}

// GetResourceTypeWithChan invokes the ros.GetResourceType API asynchronously
// api document: https://help.aliyun.com/api/ros/getresourcetype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetResourceTypeWithChan(request *GetResourceTypeRequest) (<-chan *GetResourceTypeResponse, <-chan error) {
	responseChan := make(chan *GetResourceTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetResourceType(request)
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

// GetResourceTypeWithCallback invokes the ros.GetResourceType API asynchronously
// api document: https://help.aliyun.com/api/ros/getresourcetype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetResourceTypeWithCallback(request *GetResourceTypeRequest, callback func(response *GetResourceTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetResourceTypeResponse
		var err error
		defer close(result)
		response, err = client.GetResourceType(request)
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

// GetResourceTypeRequest is the request struct for api GetResourceType
type GetResourceTypeRequest struct {
	*requests.RpcRequest
	ResourceType string `position:"Query" name:"ResourceType"`
}

// GetResourceTypeResponse is the response struct for api GetResourceType
type GetResourceTypeResponse struct {
	*responses.BaseResponse
	Attributes   map[string]interface{} `json:"Attributes" xml:"Attributes"`
	Properties   map[string]interface{} `json:"Properties" xml:"Properties"`
	RequestId    string                 `json:"RequestId" xml:"RequestId"`
	ResourceType string                 `json:"ResourceType" xml:"ResourceType"`
}

// CreateGetResourceTypeRequest creates a request to invoke GetResourceType API
func CreateGetResourceTypeRequest() (request *GetResourceTypeRequest) {
	request = &GetResourceTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ROS", "2019-09-10", "GetResourceType", "ros", "openAPI")
	return
}

// CreateGetResourceTypeResponse creates a response to parse from GetResourceType response
func CreateGetResourceTypeResponse() (response *GetResourceTypeResponse) {
	response = &GetResourceTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
