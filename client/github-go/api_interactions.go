/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// InteractionsApiService InteractionsApi service
type InteractionsApiService service

type ApiInteractionsGetRestrictionsForAuthenticatedUserRequest struct {
	ctx context.Context
	ApiService *InteractionsApiService
}

func (r ApiInteractionsGetRestrictionsForAuthenticatedUserRequest) Execute() (*InteractionsGetRestrictionsForOrg200Response, *http.Response, error) {
	return r.ApiService.InteractionsGetRestrictionsForAuthenticatedUserExecute(r)
}

/*
InteractionsGetRestrictionsForAuthenticatedUser Get interaction restrictions for your public repositories

Shows which type of GitHub user can interact with your public repositories and when the restriction expires.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiInteractionsGetRestrictionsForAuthenticatedUserRequest
*/
func (a *InteractionsApiService) InteractionsGetRestrictionsForAuthenticatedUser(ctx context.Context) ApiInteractionsGetRestrictionsForAuthenticatedUserRequest {
	return ApiInteractionsGetRestrictionsForAuthenticatedUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InteractionsGetRestrictionsForOrg200Response
func (a *InteractionsApiService) InteractionsGetRestrictionsForAuthenticatedUserExecute(r ApiInteractionsGetRestrictionsForAuthenticatedUserRequest) (*InteractionsGetRestrictionsForOrg200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InteractionsGetRestrictionsForOrg200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InteractionsApiService.InteractionsGetRestrictionsForAuthenticatedUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/interaction-limits"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiInteractionsGetRestrictionsForOrgRequest struct {
	ctx context.Context
	ApiService *InteractionsApiService
	org string
}

func (r ApiInteractionsGetRestrictionsForOrgRequest) Execute() (*InteractionsGetRestrictionsForOrg200Response, *http.Response, error) {
	return r.ApiService.InteractionsGetRestrictionsForOrgExecute(r)
}

/*
InteractionsGetRestrictionsForOrg Get interaction restrictions for an organization

Shows which type of GitHub user can interact with this organization and when the restriction expires. If there is no restrictions, you will see an empty response.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param org The organization name. The name is not case sensitive.
 @return ApiInteractionsGetRestrictionsForOrgRequest
*/
func (a *InteractionsApiService) InteractionsGetRestrictionsForOrg(ctx context.Context, org string) ApiInteractionsGetRestrictionsForOrgRequest {
	return ApiInteractionsGetRestrictionsForOrgRequest{
		ApiService: a,
		ctx: ctx,
		org: org,
	}
}

// Execute executes the request
//  @return InteractionsGetRestrictionsForOrg200Response
func (a *InteractionsApiService) InteractionsGetRestrictionsForOrgExecute(r ApiInteractionsGetRestrictionsForOrgRequest) (*InteractionsGetRestrictionsForOrg200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InteractionsGetRestrictionsForOrg200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InteractionsApiService.InteractionsGetRestrictionsForOrg")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org}/interaction-limits"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", url.PathEscape(parameterToString(r.org, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiInteractionsGetRestrictionsForRepoRequest struct {
	ctx context.Context
	ApiService *InteractionsApiService
	owner string
	repo string
}

func (r ApiInteractionsGetRestrictionsForRepoRequest) Execute() (*InteractionsGetRestrictionsForOrg200Response, *http.Response, error) {
	return r.ApiService.InteractionsGetRestrictionsForRepoExecute(r)
}

/*
InteractionsGetRestrictionsForRepo Get interaction restrictions for a repository

Shows which type of GitHub user can interact with this repository and when the restriction expires. If there are no restrictions, you will see an empty response.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param owner The account owner of the repository. The name is not case sensitive.
 @param repo The name of the repository. The name is not case sensitive.
 @return ApiInteractionsGetRestrictionsForRepoRequest
*/
func (a *InteractionsApiService) InteractionsGetRestrictionsForRepo(ctx context.Context, owner string, repo string) ApiInteractionsGetRestrictionsForRepoRequest {
	return ApiInteractionsGetRestrictionsForRepoRequest{
		ApiService: a,
		ctx: ctx,
		owner: owner,
		repo: repo,
	}
}

// Execute executes the request
//  @return InteractionsGetRestrictionsForOrg200Response
func (a *InteractionsApiService) InteractionsGetRestrictionsForRepoExecute(r ApiInteractionsGetRestrictionsForRepoRequest) (*InteractionsGetRestrictionsForOrg200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InteractionsGetRestrictionsForOrg200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InteractionsApiService.InteractionsGetRestrictionsForRepo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/repos/{owner}/{repo}/interaction-limits"
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", url.PathEscape(parameterToString(r.owner, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo"+"}", url.PathEscape(parameterToString(r.repo, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiInteractionsRemoveRestrictionsForAuthenticatedUserRequest struct {
	ctx context.Context
	ApiService *InteractionsApiService
}

func (r ApiInteractionsRemoveRestrictionsForAuthenticatedUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.InteractionsRemoveRestrictionsForAuthenticatedUserExecute(r)
}

/*
InteractionsRemoveRestrictionsForAuthenticatedUser Remove interaction restrictions from your public repositories

Removes any interaction restrictions from your public repositories.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiInteractionsRemoveRestrictionsForAuthenticatedUserRequest
*/
func (a *InteractionsApiService) InteractionsRemoveRestrictionsForAuthenticatedUser(ctx context.Context) ApiInteractionsRemoveRestrictionsForAuthenticatedUserRequest {
	return ApiInteractionsRemoveRestrictionsForAuthenticatedUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *InteractionsApiService) InteractionsRemoveRestrictionsForAuthenticatedUserExecute(r ApiInteractionsRemoveRestrictionsForAuthenticatedUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InteractionsApiService.InteractionsRemoveRestrictionsForAuthenticatedUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/interaction-limits"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiInteractionsRemoveRestrictionsForOrgRequest struct {
	ctx context.Context
	ApiService *InteractionsApiService
	org string
}

func (r ApiInteractionsRemoveRestrictionsForOrgRequest) Execute() (*http.Response, error) {
	return r.ApiService.InteractionsRemoveRestrictionsForOrgExecute(r)
}

/*
InteractionsRemoveRestrictionsForOrg Remove interaction restrictions for an organization

Removes all interaction restrictions from public repositories in the given organization. You must be an organization owner to remove restrictions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param org The organization name. The name is not case sensitive.
 @return ApiInteractionsRemoveRestrictionsForOrgRequest
*/
func (a *InteractionsApiService) InteractionsRemoveRestrictionsForOrg(ctx context.Context, org string) ApiInteractionsRemoveRestrictionsForOrgRequest {
	return ApiInteractionsRemoveRestrictionsForOrgRequest{
		ApiService: a,
		ctx: ctx,
		org: org,
	}
}

// Execute executes the request
func (a *InteractionsApiService) InteractionsRemoveRestrictionsForOrgExecute(r ApiInteractionsRemoveRestrictionsForOrgRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InteractionsApiService.InteractionsRemoveRestrictionsForOrg")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org}/interaction-limits"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", url.PathEscape(parameterToString(r.org, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiInteractionsRemoveRestrictionsForRepoRequest struct {
	ctx context.Context
	ApiService *InteractionsApiService
	owner string
	repo string
}

func (r ApiInteractionsRemoveRestrictionsForRepoRequest) Execute() (*http.Response, error) {
	return r.ApiService.InteractionsRemoveRestrictionsForRepoExecute(r)
}

/*
InteractionsRemoveRestrictionsForRepo Remove interaction restrictions for a repository

Removes all interaction restrictions from the given repository. You must have owner or admin access to remove restrictions. If the interaction limit is set for the user or organization that owns this repository, you will receive a `409 Conflict` response and will not be able to use this endpoint to change the interaction limit for a single repository.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param owner The account owner of the repository. The name is not case sensitive.
 @param repo The name of the repository. The name is not case sensitive.
 @return ApiInteractionsRemoveRestrictionsForRepoRequest
*/
func (a *InteractionsApiService) InteractionsRemoveRestrictionsForRepo(ctx context.Context, owner string, repo string) ApiInteractionsRemoveRestrictionsForRepoRequest {
	return ApiInteractionsRemoveRestrictionsForRepoRequest{
		ApiService: a,
		ctx: ctx,
		owner: owner,
		repo: repo,
	}
}

// Execute executes the request
func (a *InteractionsApiService) InteractionsRemoveRestrictionsForRepoExecute(r ApiInteractionsRemoveRestrictionsForRepoRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InteractionsApiService.InteractionsRemoveRestrictionsForRepo")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/repos/{owner}/{repo}/interaction-limits"
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", url.PathEscape(parameterToString(r.owner, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo"+"}", url.PathEscape(parameterToString(r.repo, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiInteractionsSetRestrictionsForAuthenticatedUserRequest struct {
	ctx context.Context
	ApiService *InteractionsApiService
	interactionLimit *InteractionLimit
}

func (r ApiInteractionsSetRestrictionsForAuthenticatedUserRequest) InteractionLimit(interactionLimit InteractionLimit) ApiInteractionsSetRestrictionsForAuthenticatedUserRequest {
	r.interactionLimit = &interactionLimit
	return r
}

func (r ApiInteractionsSetRestrictionsForAuthenticatedUserRequest) Execute() (*InteractionLimitResponse, *http.Response, error) {
	return r.ApiService.InteractionsSetRestrictionsForAuthenticatedUserExecute(r)
}

/*
InteractionsSetRestrictionsForAuthenticatedUser Set interaction restrictions for your public repositories

Temporarily restricts which type of GitHub user can interact with your public repositories. Setting the interaction limit at the user level will overwrite any interaction limits that are set for individual repositories owned by the user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiInteractionsSetRestrictionsForAuthenticatedUserRequest
*/
func (a *InteractionsApiService) InteractionsSetRestrictionsForAuthenticatedUser(ctx context.Context) ApiInteractionsSetRestrictionsForAuthenticatedUserRequest {
	return ApiInteractionsSetRestrictionsForAuthenticatedUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InteractionLimitResponse
func (a *InteractionsApiService) InteractionsSetRestrictionsForAuthenticatedUserExecute(r ApiInteractionsSetRestrictionsForAuthenticatedUserRequest) (*InteractionLimitResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InteractionLimitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InteractionsApiService.InteractionsSetRestrictionsForAuthenticatedUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/interaction-limits"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.interactionLimit == nil {
		return localVarReturnValue, nil, reportError("interactionLimit is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.interactionLimit
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiInteractionsSetRestrictionsForOrgRequest struct {
	ctx context.Context
	ApiService *InteractionsApiService
	org string
	interactionLimit *InteractionLimit
}

func (r ApiInteractionsSetRestrictionsForOrgRequest) InteractionLimit(interactionLimit InteractionLimit) ApiInteractionsSetRestrictionsForOrgRequest {
	r.interactionLimit = &interactionLimit
	return r
}

func (r ApiInteractionsSetRestrictionsForOrgRequest) Execute() (*InteractionLimitResponse, *http.Response, error) {
	return r.ApiService.InteractionsSetRestrictionsForOrgExecute(r)
}

/*
InteractionsSetRestrictionsForOrg Set interaction restrictions for an organization

Temporarily restricts interactions to a certain type of GitHub user in any public repository in the given organization. You must be an organization owner to set these restrictions. Setting the interaction limit at the organization level will overwrite any interaction limits that are set for individual repositories owned by the organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param org The organization name. The name is not case sensitive.
 @return ApiInteractionsSetRestrictionsForOrgRequest
*/
func (a *InteractionsApiService) InteractionsSetRestrictionsForOrg(ctx context.Context, org string) ApiInteractionsSetRestrictionsForOrgRequest {
	return ApiInteractionsSetRestrictionsForOrgRequest{
		ApiService: a,
		ctx: ctx,
		org: org,
	}
}

// Execute executes the request
//  @return InteractionLimitResponse
func (a *InteractionsApiService) InteractionsSetRestrictionsForOrgExecute(r ApiInteractionsSetRestrictionsForOrgRequest) (*InteractionLimitResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InteractionLimitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InteractionsApiService.InteractionsSetRestrictionsForOrg")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orgs/{org}/interaction-limits"
	localVarPath = strings.Replace(localVarPath, "{"+"org"+"}", url.PathEscape(parameterToString(r.org, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.interactionLimit == nil {
		return localVarReturnValue, nil, reportError("interactionLimit is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.interactionLimit
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiInteractionsSetRestrictionsForRepoRequest struct {
	ctx context.Context
	ApiService *InteractionsApiService
	owner string
	repo string
	interactionLimit *InteractionLimit
}

func (r ApiInteractionsSetRestrictionsForRepoRequest) InteractionLimit(interactionLimit InteractionLimit) ApiInteractionsSetRestrictionsForRepoRequest {
	r.interactionLimit = &interactionLimit
	return r
}

func (r ApiInteractionsSetRestrictionsForRepoRequest) Execute() (*InteractionLimitResponse, *http.Response, error) {
	return r.ApiService.InteractionsSetRestrictionsForRepoExecute(r)
}

/*
InteractionsSetRestrictionsForRepo Set interaction restrictions for a repository

Temporarily restricts interactions to a certain type of GitHub user within the given repository. You must have owner or admin access to set these restrictions. If an interaction limit is set for the user or organization that owns this repository, you will receive a `409 Conflict` response and will not be able to use this endpoint to change the interaction limit for a single repository.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param owner The account owner of the repository. The name is not case sensitive.
 @param repo The name of the repository. The name is not case sensitive.
 @return ApiInteractionsSetRestrictionsForRepoRequest
*/
func (a *InteractionsApiService) InteractionsSetRestrictionsForRepo(ctx context.Context, owner string, repo string) ApiInteractionsSetRestrictionsForRepoRequest {
	return ApiInteractionsSetRestrictionsForRepoRequest{
		ApiService: a,
		ctx: ctx,
		owner: owner,
		repo: repo,
	}
}

// Execute executes the request
//  @return InteractionLimitResponse
func (a *InteractionsApiService) InteractionsSetRestrictionsForRepoExecute(r ApiInteractionsSetRestrictionsForRepoRequest) (*InteractionLimitResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InteractionLimitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InteractionsApiService.InteractionsSetRestrictionsForRepo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/repos/{owner}/{repo}/interaction-limits"
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", url.PathEscape(parameterToString(r.owner, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo"+"}", url.PathEscape(parameterToString(r.repo, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.interactionLimit == nil {
		return localVarReturnValue, nil, reportError("interactionLimit is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.interactionLimit
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
