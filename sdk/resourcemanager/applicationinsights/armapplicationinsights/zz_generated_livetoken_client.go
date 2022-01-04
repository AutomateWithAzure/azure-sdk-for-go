//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// LiveTokenClient contains the methods for the LiveToken group.
// Don't use this type directly, use NewLiveTokenClient() instead.
type LiveTokenClient struct {
	ep string
	pl runtime.Pipeline
}

// NewLiveTokenClient creates a new instance of LiveTokenClient with the specified values.
func NewLiveTokenClient(credential azcore.TokenCredential, options *arm.ClientOptions) *LiveTokenClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &LiveTokenClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Gets an access token for live metrics stream data.
// If the operation fails it returns the *ErrorResponseLinkedStorage error type.
func (client *LiveTokenClient) Get(ctx context.Context, resourceURI string, options *LiveTokenGetOptions) (LiveTokenGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return LiveTokenGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LiveTokenGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LiveTokenGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LiveTokenClient) getCreateRequest(ctx context.Context, resourceURI string, options *LiveTokenGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/microsoft.insights/generatelivetoken"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LiveTokenClient) getHandleResponse(resp *http.Response) (LiveTokenGetResponse, error) {
	result := LiveTokenGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LiveTokenResponse); err != nil {
		return LiveTokenGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *LiveTokenClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponseLinkedStorage{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}