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

// DependabotSetSelectedReposForOrgSecretRequest struct for DependabotSetSelectedReposForOrgSecretRequest
type DependabotSetSelectedReposForOrgSecretRequest struct {
	// An array of repository ids that can access the organization secret. You can only provide a list of repository ids when the `visibility` is set to `selected`. You can add and remove individual repositories using the [Set selected repositories for an organization secret](https://docs.github.com/rest/reference/dependabot#set-selected-repositories-for-an-organization-secret) and [Remove selected repository from an organization secret](https://docs.github.com/rest/reference/dependabot#remove-selected-repository-from-an-organization-secret) endpoints.
	SelectedRepositoryIds []int32 `json:"selected_repository_ids"`
}

// NewDependabotSetSelectedReposForOrgSecretRequest instantiates a new DependabotSetSelectedReposForOrgSecretRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDependabotSetSelectedReposForOrgSecretRequest(selectedRepositoryIds []int32) *DependabotSetSelectedReposForOrgSecretRequest {
	this := DependabotSetSelectedReposForOrgSecretRequest{}
	this.SelectedRepositoryIds = selectedRepositoryIds
	return &this
}

// NewDependabotSetSelectedReposForOrgSecretRequestWithDefaults instantiates a new DependabotSetSelectedReposForOrgSecretRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDependabotSetSelectedReposForOrgSecretRequestWithDefaults() *DependabotSetSelectedReposForOrgSecretRequest {
	this := DependabotSetSelectedReposForOrgSecretRequest{}
	return &this
}

// GetSelectedRepositoryIds returns the SelectedRepositoryIds field value
func (o *DependabotSetSelectedReposForOrgSecretRequest) GetSelectedRepositoryIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.SelectedRepositoryIds
}

// GetSelectedRepositoryIdsOk returns a tuple with the SelectedRepositoryIds field value
// and a boolean to check if the value has been set.
func (o *DependabotSetSelectedReposForOrgSecretRequest) GetSelectedRepositoryIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelectedRepositoryIds, true
}

// SetSelectedRepositoryIds sets field value
func (o *DependabotSetSelectedReposForOrgSecretRequest) SetSelectedRepositoryIds(v []int32) {
	o.SelectedRepositoryIds = v
}

func (o DependabotSetSelectedReposForOrgSecretRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["selected_repository_ids"] = o.SelectedRepositoryIds
	}
	return json.Marshal(toSerialize)
}

type NullableDependabotSetSelectedReposForOrgSecretRequest struct {
	value *DependabotSetSelectedReposForOrgSecretRequest
	isSet bool
}

func (v NullableDependabotSetSelectedReposForOrgSecretRequest) Get() *DependabotSetSelectedReposForOrgSecretRequest {
	return v.value
}

func (v *NullableDependabotSetSelectedReposForOrgSecretRequest) Set(val *DependabotSetSelectedReposForOrgSecretRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDependabotSetSelectedReposForOrgSecretRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDependabotSetSelectedReposForOrgSecretRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDependabotSetSelectedReposForOrgSecretRequest(val *DependabotSetSelectedReposForOrgSecretRequest) *NullableDependabotSetSelectedReposForOrgSecretRequest {
	return &NullableDependabotSetSelectedReposForOrgSecretRequest{value: val, isSet: true}
}

func (v NullableDependabotSetSelectedReposForOrgSecretRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDependabotSetSelectedReposForOrgSecretRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


