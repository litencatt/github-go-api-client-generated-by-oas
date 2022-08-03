/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"encoding/json"
)

// ActionsGetActionsCacheUsageByRepoForOrg200Response struct for ActionsGetActionsCacheUsageByRepoForOrg200Response
type ActionsGetActionsCacheUsageByRepoForOrg200Response struct {
	TotalCount int32 `json:"total_count"`
	RepositoryCacheUsages []ActionsCacheUsageByRepository `json:"repository_cache_usages"`
}

// NewActionsGetActionsCacheUsageByRepoForOrg200Response instantiates a new ActionsGetActionsCacheUsageByRepoForOrg200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionsGetActionsCacheUsageByRepoForOrg200Response(totalCount int32, repositoryCacheUsages []ActionsCacheUsageByRepository) *ActionsGetActionsCacheUsageByRepoForOrg200Response {
	this := ActionsGetActionsCacheUsageByRepoForOrg200Response{}
	this.TotalCount = totalCount
	this.RepositoryCacheUsages = repositoryCacheUsages
	return &this
}

// NewActionsGetActionsCacheUsageByRepoForOrg200ResponseWithDefaults instantiates a new ActionsGetActionsCacheUsageByRepoForOrg200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsGetActionsCacheUsageByRepoForOrg200ResponseWithDefaults() *ActionsGetActionsCacheUsageByRepoForOrg200Response {
	this := ActionsGetActionsCacheUsageByRepoForOrg200Response{}
	return &this
}

// GetTotalCount returns the TotalCount field value
func (o *ActionsGetActionsCacheUsageByRepoForOrg200Response) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *ActionsGetActionsCacheUsageByRepoForOrg200Response) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *ActionsGetActionsCacheUsageByRepoForOrg200Response) SetTotalCount(v int32) {
	o.TotalCount = v
}

// GetRepositoryCacheUsages returns the RepositoryCacheUsages field value
func (o *ActionsGetActionsCacheUsageByRepoForOrg200Response) GetRepositoryCacheUsages() []ActionsCacheUsageByRepository {
	if o == nil {
		var ret []ActionsCacheUsageByRepository
		return ret
	}

	return o.RepositoryCacheUsages
}

// GetRepositoryCacheUsagesOk returns a tuple with the RepositoryCacheUsages field value
// and a boolean to check if the value has been set.
func (o *ActionsGetActionsCacheUsageByRepoForOrg200Response) GetRepositoryCacheUsagesOk() ([]ActionsCacheUsageByRepository, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepositoryCacheUsages, true
}

// SetRepositoryCacheUsages sets field value
func (o *ActionsGetActionsCacheUsageByRepoForOrg200Response) SetRepositoryCacheUsages(v []ActionsCacheUsageByRepository) {
	o.RepositoryCacheUsages = v
}

func (o ActionsGetActionsCacheUsageByRepoForOrg200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total_count"] = o.TotalCount
	}
	if true {
		toSerialize["repository_cache_usages"] = o.RepositoryCacheUsages
	}
	return json.Marshal(toSerialize)
}

type NullableActionsGetActionsCacheUsageByRepoForOrg200Response struct {
	value *ActionsGetActionsCacheUsageByRepoForOrg200Response
	isSet bool
}

func (v NullableActionsGetActionsCacheUsageByRepoForOrg200Response) Get() *ActionsGetActionsCacheUsageByRepoForOrg200Response {
	return v.value
}

func (v *NullableActionsGetActionsCacheUsageByRepoForOrg200Response) Set(val *ActionsGetActionsCacheUsageByRepoForOrg200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableActionsGetActionsCacheUsageByRepoForOrg200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableActionsGetActionsCacheUsageByRepoForOrg200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionsGetActionsCacheUsageByRepoForOrg200Response(val *ActionsGetActionsCacheUsageByRepoForOrg200Response) *NullableActionsGetActionsCacheUsageByRepoForOrg200Response {
	return &NullableActionsGetActionsCacheUsageByRepoForOrg200Response{value: val, isSet: true}
}

func (v NullableActionsGetActionsCacheUsageByRepoForOrg200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionsGetActionsCacheUsageByRepoForOrg200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

