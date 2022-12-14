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

// ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response struct for ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response
type ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response struct {
	TotalCount float32 `json:"total_count"`
	Repositories []Repository `json:"repositories"`
}

// NewActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response instantiates a new ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response(totalCount float32, repositories []Repository) *ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response {
	this := ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response{}
	this.TotalCount = totalCount
	this.Repositories = repositories
	return &this
}

// NewActionsListSelectedRepositoriesEnabledGithubActionsOrganization200ResponseWithDefaults instantiates a new ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsListSelectedRepositoriesEnabledGithubActionsOrganization200ResponseWithDefaults() *ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response {
	this := ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response{}
	return &this
}

// GetTotalCount returns the TotalCount field value
func (o *ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response) GetTotalCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response) GetTotalCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response) SetTotalCount(v float32) {
	o.TotalCount = v
}

// GetRepositories returns the Repositories field value
func (o *ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response) GetRepositories() []Repository {
	if o == nil {
		var ret []Repository
		return ret
	}

	return o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value
// and a boolean to check if the value has been set.
func (o *ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response) GetRepositoriesOk() ([]Repository, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repositories, true
}

// SetRepositories sets field value
func (o *ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response) SetRepositories(v []Repository) {
	o.Repositories = v
}

func (o ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total_count"] = o.TotalCount
	}
	if true {
		toSerialize["repositories"] = o.Repositories
	}
	return json.Marshal(toSerialize)
}

type NullableActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response struct {
	value *ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response
	isSet bool
}

func (v NullableActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response) Get() *ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response {
	return v.value
}

func (v *NullableActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response) Set(val *ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response(val *ActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response) *NullableActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response {
	return &NullableActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response{value: val, isSet: true}
}

func (v NullableActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionsListSelectedRepositoriesEnabledGithubActionsOrganization200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


