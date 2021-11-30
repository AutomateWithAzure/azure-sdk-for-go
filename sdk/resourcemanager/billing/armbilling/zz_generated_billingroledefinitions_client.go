//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

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

// BillingRoleDefinitionsClient contains the methods for the BillingRoleDefinitions group.
// Don't use this type directly, use NewBillingRoleDefinitionsClient() instead.
type BillingRoleDefinitionsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewBillingRoleDefinitionsClient creates a new instance of BillingRoleDefinitionsClient with the specified values.
func NewBillingRoleDefinitionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *BillingRoleDefinitionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &BillingRoleDefinitionsClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// GetByBillingAccount - Gets the definition for a role on a billing account. The operation is supported for billing accounts with agreement type Microsoft
// Partner Agreement or Microsoft Customer Agreement.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BillingRoleDefinitionsClient) GetByBillingAccount(ctx context.Context, billingAccountName string, billingRoleDefinitionName string, options *BillingRoleDefinitionsGetByBillingAccountOptions) (BillingRoleDefinitionsGetByBillingAccountResponse, error) {
	req, err := client.getByBillingAccountCreateRequest(ctx, billingAccountName, billingRoleDefinitionName, options)
	if err != nil {
		return BillingRoleDefinitionsGetByBillingAccountResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BillingRoleDefinitionsGetByBillingAccountResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BillingRoleDefinitionsGetByBillingAccountResponse{}, client.getByBillingAccountHandleError(resp)
	}
	return client.getByBillingAccountHandleResponse(resp)
}

// getByBillingAccountCreateRequest creates the GetByBillingAccount request.
func (client *BillingRoleDefinitionsClient) getByBillingAccountCreateRequest(ctx context.Context, billingAccountName string, billingRoleDefinitionName string, options *BillingRoleDefinitionsGetByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingRoleDefinitions/{billingRoleDefinitionName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingRoleDefinitionName == "" {
		return nil, errors.New("parameter billingRoleDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingRoleDefinitionName}", url.PathEscape(billingRoleDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByBillingAccountHandleResponse handles the GetByBillingAccount response.
func (client *BillingRoleDefinitionsClient) getByBillingAccountHandleResponse(resp *http.Response) (BillingRoleDefinitionsGetByBillingAccountResponse, error) {
	result := BillingRoleDefinitionsGetByBillingAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BillingRoleDefinition); err != nil {
		return BillingRoleDefinitionsGetByBillingAccountResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getByBillingAccountHandleError handles the GetByBillingAccount error response.
func (client *BillingRoleDefinitionsClient) getByBillingAccountHandleError(resp *http.Response) error {
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

// GetByBillingProfile - Gets the definition for a role on a billing profile. The operation is supported for billing accounts with agreement type Microsoft
// Partner Agreement or Microsoft Customer Agreement.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BillingRoleDefinitionsClient) GetByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleDefinitionName string, options *BillingRoleDefinitionsGetByBillingProfileOptions) (BillingRoleDefinitionsGetByBillingProfileResponse, error) {
	req, err := client.getByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, billingRoleDefinitionName, options)
	if err != nil {
		return BillingRoleDefinitionsGetByBillingProfileResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BillingRoleDefinitionsGetByBillingProfileResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BillingRoleDefinitionsGetByBillingProfileResponse{}, client.getByBillingProfileHandleError(resp)
	}
	return client.getByBillingProfileHandleResponse(resp)
}

// getByBillingProfileCreateRequest creates the GetByBillingProfile request.
func (client *BillingRoleDefinitionsClient) getByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, billingRoleDefinitionName string, options *BillingRoleDefinitionsGetByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingRoleDefinitions/{billingRoleDefinitionName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if billingRoleDefinitionName == "" {
		return nil, errors.New("parameter billingRoleDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingRoleDefinitionName}", url.PathEscape(billingRoleDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByBillingProfileHandleResponse handles the GetByBillingProfile response.
func (client *BillingRoleDefinitionsClient) getByBillingProfileHandleResponse(resp *http.Response) (BillingRoleDefinitionsGetByBillingProfileResponse, error) {
	result := BillingRoleDefinitionsGetByBillingProfileResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BillingRoleDefinition); err != nil {
		return BillingRoleDefinitionsGetByBillingProfileResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getByBillingProfileHandleError handles the GetByBillingProfile error response.
func (client *BillingRoleDefinitionsClient) getByBillingProfileHandleError(resp *http.Response) error {
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

// GetByInvoiceSection - Gets the definition for a role on an invoice section. The operation is supported only for billing accounts with agreement type
// Microsoft Customer Agreement.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BillingRoleDefinitionsClient) GetByInvoiceSection(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleDefinitionName string, options *BillingRoleDefinitionsGetByInvoiceSectionOptions) (BillingRoleDefinitionsGetByInvoiceSectionResponse, error) {
	req, err := client.getByInvoiceSectionCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, billingRoleDefinitionName, options)
	if err != nil {
		return BillingRoleDefinitionsGetByInvoiceSectionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BillingRoleDefinitionsGetByInvoiceSectionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BillingRoleDefinitionsGetByInvoiceSectionResponse{}, client.getByInvoiceSectionHandleError(resp)
	}
	return client.getByInvoiceSectionHandleResponse(resp)
}

// getByInvoiceSectionCreateRequest creates the GetByInvoiceSection request.
func (client *BillingRoleDefinitionsClient) getByInvoiceSectionCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleDefinitionName string, options *BillingRoleDefinitionsGetByInvoiceSectionOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingRoleDefinitions/{billingRoleDefinitionName}"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if invoiceSectionName == "" {
		return nil, errors.New("parameter invoiceSectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceSectionName}", url.PathEscape(invoiceSectionName))
	if billingRoleDefinitionName == "" {
		return nil, errors.New("parameter billingRoleDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingRoleDefinitionName}", url.PathEscape(billingRoleDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByInvoiceSectionHandleResponse handles the GetByInvoiceSection response.
func (client *BillingRoleDefinitionsClient) getByInvoiceSectionHandleResponse(resp *http.Response) (BillingRoleDefinitionsGetByInvoiceSectionResponse, error) {
	result := BillingRoleDefinitionsGetByInvoiceSectionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BillingRoleDefinition); err != nil {
		return BillingRoleDefinitionsGetByInvoiceSectionResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getByInvoiceSectionHandleError handles the GetByInvoiceSection error response.
func (client *BillingRoleDefinitionsClient) getByInvoiceSectionHandleError(resp *http.Response) error {
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

// ListByBillingAccount - Lists the role definitions for a billing account. The operation is supported for billing accounts with agreement type Microsoft
// Partner Agreement or Microsoft Customer Agreement.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BillingRoleDefinitionsClient) ListByBillingAccount(billingAccountName string, options *BillingRoleDefinitionsListByBillingAccountOptions) *BillingRoleDefinitionsListByBillingAccountPager {
	return &BillingRoleDefinitionsListByBillingAccountPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByBillingAccountCreateRequest(ctx, billingAccountName, options)
		},
		advancer: func(ctx context.Context, resp BillingRoleDefinitionsListByBillingAccountResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.BillingRoleDefinitionListResult.NextLink)
		},
	}
}

// listByBillingAccountCreateRequest creates the ListByBillingAccount request.
func (client *BillingRoleDefinitionsClient) listByBillingAccountCreateRequest(ctx context.Context, billingAccountName string, options *BillingRoleDefinitionsListByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingRoleDefinitions"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByBillingAccountHandleResponse handles the ListByBillingAccount response.
func (client *BillingRoleDefinitionsClient) listByBillingAccountHandleResponse(resp *http.Response) (BillingRoleDefinitionsListByBillingAccountResponse, error) {
	result := BillingRoleDefinitionsListByBillingAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BillingRoleDefinitionListResult); err != nil {
		return BillingRoleDefinitionsListByBillingAccountResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByBillingAccountHandleError handles the ListByBillingAccount error response.
func (client *BillingRoleDefinitionsClient) listByBillingAccountHandleError(resp *http.Response) error {
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

// ListByBillingProfile - Lists the role definitions for a billing profile. The operation is supported for billing accounts with agreement type Microsoft
// Partner Agreement or Microsoft Customer Agreement.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BillingRoleDefinitionsClient) ListByBillingProfile(billingAccountName string, billingProfileName string, options *BillingRoleDefinitionsListByBillingProfileOptions) *BillingRoleDefinitionsListByBillingProfilePager {
	return &BillingRoleDefinitionsListByBillingProfilePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, options)
		},
		advancer: func(ctx context.Context, resp BillingRoleDefinitionsListByBillingProfileResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.BillingRoleDefinitionListResult.NextLink)
		},
	}
}

// listByBillingProfileCreateRequest creates the ListByBillingProfile request.
func (client *BillingRoleDefinitionsClient) listByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, options *BillingRoleDefinitionsListByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/billingRoleDefinitions"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByBillingProfileHandleResponse handles the ListByBillingProfile response.
func (client *BillingRoleDefinitionsClient) listByBillingProfileHandleResponse(resp *http.Response) (BillingRoleDefinitionsListByBillingProfileResponse, error) {
	result := BillingRoleDefinitionsListByBillingProfileResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BillingRoleDefinitionListResult); err != nil {
		return BillingRoleDefinitionsListByBillingProfileResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByBillingProfileHandleError handles the ListByBillingProfile error response.
func (client *BillingRoleDefinitionsClient) listByBillingProfileHandleError(resp *http.Response) error {
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

// ListByInvoiceSection - Lists the role definitions for an invoice section. The operation is supported for billing accounts with agreement type Microsoft
// Partner Agreement or Microsoft Customer Agreement.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BillingRoleDefinitionsClient) ListByInvoiceSection(billingAccountName string, billingProfileName string, invoiceSectionName string, options *BillingRoleDefinitionsListByInvoiceSectionOptions) *BillingRoleDefinitionsListByInvoiceSectionPager {
	return &BillingRoleDefinitionsListByInvoiceSectionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByInvoiceSectionCreateRequest(ctx, billingAccountName, billingProfileName, invoiceSectionName, options)
		},
		advancer: func(ctx context.Context, resp BillingRoleDefinitionsListByInvoiceSectionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.BillingRoleDefinitionListResult.NextLink)
		},
	}
}

// listByInvoiceSectionCreateRequest creates the ListByInvoiceSection request.
func (client *BillingRoleDefinitionsClient) listByInvoiceSectionCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, invoiceSectionName string, options *BillingRoleDefinitionsListByInvoiceSectionOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/billingRoleDefinitions"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	if invoiceSectionName == "" {
		return nil, errors.New("parameter invoiceSectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{invoiceSectionName}", url.PathEscape(invoiceSectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByInvoiceSectionHandleResponse handles the ListByInvoiceSection response.
func (client *BillingRoleDefinitionsClient) listByInvoiceSectionHandleResponse(resp *http.Response) (BillingRoleDefinitionsListByInvoiceSectionResponse, error) {
	result := BillingRoleDefinitionsListByInvoiceSectionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BillingRoleDefinitionListResult); err != nil {
		return BillingRoleDefinitionsListByInvoiceSectionResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByInvoiceSectionHandleError handles the ListByInvoiceSection error response.
func (client *BillingRoleDefinitionsClient) listByInvoiceSectionHandleError(resp *http.Response) error {
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