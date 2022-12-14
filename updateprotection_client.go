//go:build go1.18
// +build go1.18

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.9.0, generator: @autorest/go@4.0.0-preview.42)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package dataprotectiondatasourceplugin

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

type UpdateProtectionClient struct {
	pl runtime.Pipeline
}

// NewUpdateProtectionClient creates a new instance of UpdateProtectionClient with the specified values.
// pl - the pipeline used for sending requests and handling responses.
func NewUpdateProtectionClient(pl runtime.Pipeline) *UpdateProtectionClient {
	client := &UpdateProtectionClient{
		pl: pl,
	}
	return client
}

// Cancel - Cancel the operation. Poll the LRO to get the final status.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// operationID - Unique Id of this LRO operation
// subscriptionID - SubscriptionId of the resource
// resourceID - unique id of the resource
// taskID - unique id of the current task
// xmsClientRequestID - correlation request Id for tracking a particular call.
// parameters - Request body for operation
// options - UpdateProtectionClientCancelOptions contains the optional parameters for the UpdateProtectionClient.Cancel method.
func (client *UpdateProtectionClient) Cancel(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters CancelRequest, options *UpdateProtectionClientCancelOptions) (UpdateProtectionClientCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, operationID, subscriptionID, resourceID, taskID, xmsClientRequestID, parameters, options)
	if err != nil {
		return UpdateProtectionClientCancelResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return UpdateProtectionClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return UpdateProtectionClientCancelResponse{}, runtime.NewResponseError(resp)
	}
	return client.cancelHandleResponse(resp)
}

// cancelCreateRequest creates the Cancel request.
func (client *UpdateProtectionClient) cancelCreateRequest(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters CancelRequest, options *UpdateProtectionClientCancelOptions) (*policy.Request, error) {
	urlPath := "/plugin/updateProtectionOperations/{operationId}:cancel"
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["subscription-id"] = []string{subscriptionID}
	req.Raw().Header["resource-id"] = []string{resourceID}
	req.Raw().Header["task-id"] = []string{taskID}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// cancelHandleResponse handles the Cancel response.
func (client *UpdateProtectionClient) cancelHandleResponse(resp *http.Response) (UpdateProtectionClientCancelResponse, error) {
	result := UpdateProtectionClientCancelResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return UpdateProtectionClientCancelResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("x-ms-error-code"); val != "" {
		result.XMSErrorCode = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.UpdateProtectionResponse); err != nil {
		return UpdateProtectionClientCancelResponse{}, err
	}
	return result, nil
}

// Get - Gets the status of a updateProtection LRO.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// operationID - Unique Id of this LRO operation
// subscriptionID - SubscriptionId of the resource
// resourceID - Unique id of the resource
// taskID - Unique id of the current task
// xmsClientRequestID - Correlation request Id for tracking a particular request.
// options - UpdateProtectionClientGetOptions contains the optional parameters for the UpdateProtectionClient.Get method.
func (client *UpdateProtectionClient) Get(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, options *UpdateProtectionClientGetOptions) (UpdateProtectionClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, operationID, subscriptionID, resourceID, taskID, xmsClientRequestID, options)
	if err != nil {
		return UpdateProtectionClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return UpdateProtectionClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return UpdateProtectionClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *UpdateProtectionClient) getCreateRequest(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, options *UpdateProtectionClientGetOptions) (*policy.Request, error) {
	urlPath := "/plugin/updateProtectionOperations/{operationId}"
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["subscription-id"] = []string{subscriptionID}
	req.Raw().Header["resource-id"] = []string{resourceID}
	req.Raw().Header["task-id"] = []string{taskID}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *UpdateProtectionClient) getHandleResponse(resp *http.Response) (UpdateProtectionClientGetResponse, error) {
	result := UpdateProtectionClientGetResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return UpdateProtectionClientGetResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("x-ms-error-code"); val != "" {
		result.XMSErrorCode = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.UpdateProtectionResponse); err != nil {
		return UpdateProtectionClientGetResponse{}, err
	}
	return result, nil
}

// RefreshTokens - Refresh tokens for a given operation.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-03-01-preview
// operationID - Unique Id of this LRO operation
// subscriptionID - SubscriptionId of the resource
// resourceID - unique id of the resource
// taskID - unique id of the current task
// xmsClientRequestID - correlation request Id for tracking a particular call.
// parameters - Request body for operation
// options - UpdateProtectionClientRefreshTokensOptions contains the optional parameters for the UpdateProtectionClient.RefreshTokens
// method.
func (client *UpdateProtectionClient) RefreshTokens(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters RefreshTokensRequest, options *UpdateProtectionClientRefreshTokensOptions) (UpdateProtectionClientRefreshTokensResponse, error) {
	req, err := client.refreshTokensCreateRequest(ctx, operationID, subscriptionID, resourceID, taskID, xmsClientRequestID, parameters, options)
	if err != nil {
		return UpdateProtectionClientRefreshTokensResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return UpdateProtectionClientRefreshTokensResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return UpdateProtectionClientRefreshTokensResponse{}, runtime.NewResponseError(resp)
	}
	return client.refreshTokensHandleResponse(resp)
}

// refreshTokensCreateRequest creates the RefreshTokens request.
func (client *UpdateProtectionClient) refreshTokensCreateRequest(ctx context.Context, operationID string, subscriptionID string, resourceID string, taskID string, xmsClientRequestID string, parameters RefreshTokensRequest, options *UpdateProtectionClientRefreshTokensOptions) (*policy.Request, error) {
	urlPath := "/plugin/updateProtectionOperations/{operationId}:refreshTokens"
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["subscription-id"] = []string{subscriptionID}
	req.Raw().Header["resource-id"] = []string{resourceID}
	req.Raw().Header["task-id"] = []string{taskID}
	req.Raw().Header["x-ms-client-request-id"] = []string{xmsClientRequestID}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// refreshTokensHandleResponse handles the RefreshTokens response.
func (client *UpdateProtectionClient) refreshTokensHandleResponse(resp *http.Response) (UpdateProtectionClientRefreshTokensResponse, error) {
	result := UpdateProtectionClientRefreshTokensResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return UpdateProtectionClientRefreshTokensResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if val := resp.Header.Get("x-ms-error-code"); val != "" {
		result.XMSErrorCode = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.XMSRequestID = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.UpdateProtectionResponse); err != nil {
		return UpdateProtectionClientRefreshTokensResponse{}, err
	}
	return result, nil
}
