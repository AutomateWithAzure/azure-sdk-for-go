//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcostmanagement

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ForecastClient contains the methods for the Forecast group.
// Don't use this type directly, use NewForecastClient() instead.
type ForecastClient struct {
	ep string
	pl runtime.Pipeline
}

// NewForecastClient creates a new instance of ForecastClient with the specified values.
func NewForecastClient(credential azcore.TokenCredential, options *arm.ClientOptions) *ForecastClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ForecastClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// ExternalCloudProviderUsage - Lists the forecast charges for external cloud provider type defined.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ForecastClient) ExternalCloudProviderUsage(ctx context.Context, externalCloudProviderType ExternalCloudProviderType, externalCloudProviderID string, parameters ForecastDefinition, options *ForecastExternalCloudProviderUsageOptions) (ForecastExternalCloudProviderUsageResponse, error) {
	req, err := client.externalCloudProviderUsageCreateRequest(ctx, externalCloudProviderType, externalCloudProviderID, parameters, options)
	if err != nil {
		return ForecastExternalCloudProviderUsageResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ForecastExternalCloudProviderUsageResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ForecastExternalCloudProviderUsageResponse{}, client.externalCloudProviderUsageHandleError(resp)
	}
	return client.externalCloudProviderUsageHandleResponse(resp)
}

// externalCloudProviderUsageCreateRequest creates the ExternalCloudProviderUsage request.
func (client *ForecastClient) externalCloudProviderUsageCreateRequest(ctx context.Context, externalCloudProviderType ExternalCloudProviderType, externalCloudProviderID string, parameters ForecastDefinition, options *ForecastExternalCloudProviderUsageOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.CostManagement/{externalCloudProviderType}/{externalCloudProviderId}/forecast"
	if externalCloudProviderType == "" {
		return nil, errors.New("parameter externalCloudProviderType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalCloudProviderType}", url.PathEscape(string(externalCloudProviderType)))
	if externalCloudProviderID == "" {
		return nil, errors.New("parameter externalCloudProviderID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalCloudProviderId}", url.PathEscape(externalCloudProviderID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// externalCloudProviderUsageHandleResponse handles the ExternalCloudProviderUsage response.
func (client *ForecastClient) externalCloudProviderUsageHandleResponse(resp *http.Response) (ForecastExternalCloudProviderUsageResponse, error) {
	result := ForecastExternalCloudProviderUsageResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueryResult); err != nil {
		return ForecastExternalCloudProviderUsageResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// externalCloudProviderUsageHandleError handles the ExternalCloudProviderUsage error response.
func (client *ForecastClient) externalCloudProviderUsageHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Usage - Lists the forecast charges for scope defined.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ForecastClient) Usage(ctx context.Context, scope string, parameters ForecastDefinition, options *ForecastUsageOptions) (ForecastUsageResponse, error) {
	req, err := client.usageCreateRequest(ctx, scope, parameters, options)
	if err != nil {
		return ForecastUsageResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ForecastUsageResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ForecastUsageResponse{}, client.usageHandleError(resp)
	}
	return client.usageHandleResponse(resp)
}

// usageCreateRequest creates the Usage request.
func (client *ForecastClient) usageCreateRequest(ctx context.Context, scope string, parameters ForecastDefinition, options *ForecastUsageOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CostManagement/forecast"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// usageHandleResponse handles the Usage response.
func (client *ForecastClient) usageHandleResponse(resp *http.Response) (ForecastUsageResponse, error) {
	result := ForecastUsageResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueryResult); err != nil {
		return ForecastUsageResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// usageHandleError handles the Usage error response.
func (client *ForecastClient) usageHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}