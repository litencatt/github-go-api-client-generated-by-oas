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


// LicensesApiService LicensesApi service
type LicensesApiService service

type ApiLicensesGetRequest struct {
	ctx context.Context
	ApiService *LicensesApiService
	license string
}

func (r ApiLicensesGetRequest) Execute() (*License, *http.Response, error) {
	return r.ApiService.LicensesGetExecute(r)
}

/*
LicensesGet Get a license



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param license
 @return ApiLicensesGetRequest
*/
func (a *LicensesApiService) LicensesGet(ctx context.Context, license string) ApiLicensesGetRequest {
	return ApiLicensesGetRequest{
		ApiService: a,
		ctx: ctx,
		license: license,
	}
}

// Execute executes the request
//  @return License
func (a *LicensesApiService) LicensesGetExecute(r ApiLicensesGetRequest) (*License, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *License
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LicensesApiService.LicensesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/licenses/{license}"
	localVarPath = strings.Replace(localVarPath, "{"+"license"+"}", url.PathEscape(parameterToString(r.license, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 403 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v BasicError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiLicensesGetAllCommonlyUsedRequest struct {
	ctx context.Context
	ApiService *LicensesApiService
	featured *bool
	perPage *int32
	page *int32
}

func (r ApiLicensesGetAllCommonlyUsedRequest) Featured(featured bool) ApiLicensesGetAllCommonlyUsedRequest {
	r.featured = &featured
	return r
}

// The number of results per page (max 100).
func (r ApiLicensesGetAllCommonlyUsedRequest) PerPage(perPage int32) ApiLicensesGetAllCommonlyUsedRequest {
	r.perPage = &perPage
	return r
}

// Page number of the results to fetch.
func (r ApiLicensesGetAllCommonlyUsedRequest) Page(page int32) ApiLicensesGetAllCommonlyUsedRequest {
	r.page = &page
	return r
}

func (r ApiLicensesGetAllCommonlyUsedRequest) Execute() ([]LicenseSimple, *http.Response, error) {
	return r.ApiService.LicensesGetAllCommonlyUsedExecute(r)
}

/*
LicensesGetAllCommonlyUsed Get all commonly used licenses



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLicensesGetAllCommonlyUsedRequest
*/
func (a *LicensesApiService) LicensesGetAllCommonlyUsed(ctx context.Context) ApiLicensesGetAllCommonlyUsedRequest {
	return ApiLicensesGetAllCommonlyUsedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []LicenseSimple
func (a *LicensesApiService) LicensesGetAllCommonlyUsedExecute(r ApiLicensesGetAllCommonlyUsedRequest) ([]LicenseSimple, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []LicenseSimple
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LicensesApiService.LicensesGetAllCommonlyUsed")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/licenses"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.featured != nil {
		localVarQueryParams.Add("featured", parameterToString(*r.featured, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("per_page", parameterToString(*r.perPage, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
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

type ApiLicensesGetForRepoRequest struct {
	ctx context.Context
	ApiService *LicensesApiService
	owner string
	repo string
}

func (r ApiLicensesGetForRepoRequest) Execute() (*LicenseContent, *http.Response, error) {
	return r.ApiService.LicensesGetForRepoExecute(r)
}

/*
LicensesGetForRepo Get the license for a repository

This method returns the contents of the repository's license file, if one is detected.

Similar to [Get repository content](https://docs.github.com/rest/reference/repos#get-repository-content), this method also supports [custom media types](https://docs.github.com/rest/overview/media-types) for retrieving the raw license content or rendered license HTML.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param owner The account owner of the repository. The name is not case sensitive.
 @param repo The name of the repository. The name is not case sensitive.
 @return ApiLicensesGetForRepoRequest
*/
func (a *LicensesApiService) LicensesGetForRepo(ctx context.Context, owner string, repo string) ApiLicensesGetForRepoRequest {
	return ApiLicensesGetForRepoRequest{
		ApiService: a,
		ctx: ctx,
		owner: owner,
		repo: repo,
	}
}

// Execute executes the request
//  @return LicenseContent
func (a *LicensesApiService) LicensesGetForRepoExecute(r ApiLicensesGetForRepoRequest) (*LicenseContent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LicenseContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LicensesApiService.LicensesGetForRepo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/repos/{owner}/{repo}/license"
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
