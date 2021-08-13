// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// NamedValueClient contains the methods for the NamedValue group.
// Don't use this type directly, use NewNamedValueClient() instead.
type NamedValueClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewNamedValueClient creates a new instance of NamedValueClient with the specified values.
func NewNamedValueClient(con *armcore.Connection, subscriptionID string) *NamedValueClient {
	return &NamedValueClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates named value.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NamedValueClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, parameters NamedValueCreateContract, options *NamedValueBeginCreateOrUpdateOptions) (NamedValueCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceName, namedValueID, parameters, options)
	if err != nil {
		return NamedValueCreateOrUpdatePollerResponse{}, err
	}
	result := NamedValueCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("NamedValueClient.CreateOrUpdate", "location", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return NamedValueCreateOrUpdatePollerResponse{}, err
	}
	poller := &namedValueCreateOrUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (NamedValueCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new NamedValueCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to NamedValueCreateOrUpdatePoller.ResumeToken().
func (client *NamedValueClient) ResumeCreateOrUpdate(ctx context.Context, token string) (NamedValueCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("NamedValueClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return NamedValueCreateOrUpdatePollerResponse{}, err
	}
	poller := &namedValueCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return NamedValueCreateOrUpdatePollerResponse{}, err
	}
	result := NamedValueCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (NamedValueCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates named value.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NamedValueClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, parameters NamedValueCreateContract, options *NamedValueBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, namedValueID, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NamedValueClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, parameters NamedValueCreateContract, options *NamedValueBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues/{namedValueId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if namedValueID == "" {
		return nil, errors.New("parameter namedValueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namedValueId}", url.PathEscape(namedValueID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *NamedValueClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Deletes specific named value from the API Management service instance.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NamedValueClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, ifMatch string, options *NamedValueDeleteOptions) (NamedValueDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, namedValueID, ifMatch, options)
	if err != nil {
		return NamedValueDeleteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return NamedValueDeleteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return NamedValueDeleteResponse{}, client.deleteHandleError(resp)
	}
	return NamedValueDeleteResponse{RawResponse: resp.Response}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NamedValueClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, ifMatch string, options *NamedValueDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues/{namedValueId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if namedValueID == "" {
		return nil, errors.New("parameter namedValueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namedValueId}", url.PathEscape(namedValueID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("If-Match", ifMatch)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *NamedValueClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets the details of the named value specified by its identifier.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NamedValueClient) Get(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, options *NamedValueGetOptions) (NamedValueGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, namedValueID, options)
	if err != nil {
		return NamedValueGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return NamedValueGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return NamedValueGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NamedValueClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, options *NamedValueGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues/{namedValueId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if namedValueID == "" {
		return nil, errors.New("parameter namedValueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namedValueId}", url.PathEscape(namedValueID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NamedValueClient) getHandleResponse(resp *azcore.Response) (NamedValueGetResponse, error) {
	result := NamedValueGetResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := resp.UnmarshalAsJSON(&result.NamedValueContract); err != nil {
		return NamedValueGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *NamedValueClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetEntityTag - Gets the entity state (Etag) version of the named value specified by its identifier.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NamedValueClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, options *NamedValueGetEntityTagOptions) (NamedValueGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, namedValueID, options)
	if err != nil {
		return NamedValueGetEntityTagResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return NamedValueGetEntityTagResponse{}, err
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *NamedValueClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, options *NamedValueGetEntityTagOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues/{namedValueId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if namedValueID == "" {
		return nil, errors.New("parameter namedValueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namedValueId}", url.PathEscape(namedValueID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *NamedValueClient) getEntityTagHandleResponse(resp *azcore.Response) (NamedValueGetEntityTagResponse, error) {
	result := NamedValueGetEntityTagResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// ListByService - Lists a collection of named values defined within a service instance.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NamedValueClient) ListByService(resourceGroupName string, serviceName string, options *NamedValueListByServiceOptions) NamedValueListByServicePager {
	return &namedValueListByServicePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
		},
		advancer: func(ctx context.Context, resp NamedValueListByServiceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NamedValueCollection.NextLink)
		},
	}
}

// listByServiceCreateRequest creates the ListByService request.
func (client *NamedValueClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *NamedValueListByServiceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.IsKeyVaultRefreshFailed != nil {
		reqQP.Set("isKeyVaultRefreshFailed", strconv.FormatBool(*options.IsKeyVaultRefreshFailed))
	}
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *NamedValueClient) listByServiceHandleResponse(resp *azcore.Response) (NamedValueListByServiceResponse, error) {
	result := NamedValueListByServiceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.NamedValueCollection); err != nil {
		return NamedValueListByServiceResponse{}, err
	}
	return result, nil
}

// listByServiceHandleError handles the ListByService error response.
func (client *NamedValueClient) listByServiceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListValue - Gets the secret of the named value specified by its identifier.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NamedValueClient) ListValue(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, options *NamedValueListValueOptions) (NamedValueListValueResponse, error) {
	req, err := client.listValueCreateRequest(ctx, resourceGroupName, serviceName, namedValueID, options)
	if err != nil {
		return NamedValueListValueResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return NamedValueListValueResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return NamedValueListValueResponse{}, client.listValueHandleError(resp)
	}
	return client.listValueHandleResponse(resp)
}

// listValueCreateRequest creates the ListValue request.
func (client *NamedValueClient) listValueCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, options *NamedValueListValueOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues/{namedValueId}/listValue"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if namedValueID == "" {
		return nil, errors.New("parameter namedValueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namedValueId}", url.PathEscape(namedValueID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listValueHandleResponse handles the ListValue response.
func (client *NamedValueClient) listValueHandleResponse(resp *azcore.Response) (NamedValueListValueResponse, error) {
	result := NamedValueListValueResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := resp.UnmarshalAsJSON(&result.NamedValueSecretContract); err != nil {
		return NamedValueListValueResponse{}, err
	}
	return result, nil
}

// listValueHandleError handles the ListValue error response.
func (client *NamedValueClient) listValueHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginRefreshSecret - Refresh the secret of the named value specified by its identifier.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NamedValueClient) BeginRefreshSecret(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, options *NamedValueBeginRefreshSecretOptions) (NamedValueRefreshSecretPollerResponse, error) {
	resp, err := client.refreshSecret(ctx, resourceGroupName, serviceName, namedValueID, options)
	if err != nil {
		return NamedValueRefreshSecretPollerResponse{}, err
	}
	result := NamedValueRefreshSecretPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("NamedValueClient.RefreshSecret", "location", resp, client.con.Pipeline(), client.refreshSecretHandleError)
	if err != nil {
		return NamedValueRefreshSecretPollerResponse{}, err
	}
	poller := &namedValueRefreshSecretPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (NamedValueRefreshSecretResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeRefreshSecret creates a new NamedValueRefreshSecretPoller from the specified resume token.
// token - The value must come from a previous call to NamedValueRefreshSecretPoller.ResumeToken().
func (client *NamedValueClient) ResumeRefreshSecret(ctx context.Context, token string) (NamedValueRefreshSecretPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("NamedValueClient.RefreshSecret", token, client.con.Pipeline(), client.refreshSecretHandleError)
	if err != nil {
		return NamedValueRefreshSecretPollerResponse{}, err
	}
	poller := &namedValueRefreshSecretPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return NamedValueRefreshSecretPollerResponse{}, err
	}
	result := NamedValueRefreshSecretPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (NamedValueRefreshSecretResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// RefreshSecret - Refresh the secret of the named value specified by its identifier.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NamedValueClient) refreshSecret(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, options *NamedValueBeginRefreshSecretOptions) (*azcore.Response, error) {
	req, err := client.refreshSecretCreateRequest(ctx, resourceGroupName, serviceName, namedValueID, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.refreshSecretHandleError(resp)
	}
	return resp, nil
}

// refreshSecretCreateRequest creates the RefreshSecret request.
func (client *NamedValueClient) refreshSecretCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, options *NamedValueBeginRefreshSecretOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues/{namedValueId}/refreshSecret"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if namedValueID == "" {
		return nil, errors.New("parameter namedValueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namedValueId}", url.PathEscape(namedValueID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// refreshSecretHandleError handles the RefreshSecret error response.
func (client *NamedValueClient) refreshSecretHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginUpdate - Updates the specific named value.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NamedValueClient) BeginUpdate(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, ifMatch string, parameters NamedValueUpdateParameters, options *NamedValueBeginUpdateOptions) (NamedValueUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, serviceName, namedValueID, ifMatch, parameters, options)
	if err != nil {
		return NamedValueUpdatePollerResponse{}, err
	}
	result := NamedValueUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("NamedValueClient.Update", "location", resp, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return NamedValueUpdatePollerResponse{}, err
	}
	poller := &namedValueUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (NamedValueUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new NamedValueUpdatePoller from the specified resume token.
// token - The value must come from a previous call to NamedValueUpdatePoller.ResumeToken().
func (client *NamedValueClient) ResumeUpdate(ctx context.Context, token string) (NamedValueUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("NamedValueClient.Update", token, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return NamedValueUpdatePollerResponse{}, err
	}
	poller := &namedValueUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return NamedValueUpdatePollerResponse{}, err
	}
	result := NamedValueUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (NamedValueUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Update - Updates the specific named value.
// If the operation fails it returns the *ErrorResponse error type.
func (client *NamedValueClient) update(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, ifMatch string, parameters NamedValueUpdateParameters, options *NamedValueBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, namedValueID, ifMatch, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *NamedValueClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, namedValueID string, ifMatch string, parameters NamedValueUpdateParameters, options *NamedValueBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/namedValues/{namedValueId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if namedValueID == "" {
		return nil, errors.New("parameter namedValueID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namedValueId}", url.PathEscape(namedValueID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("If-Match", ifMatch)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleError handles the Update error response.
func (client *NamedValueClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}