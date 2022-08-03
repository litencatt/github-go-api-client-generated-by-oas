/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"encoding/json"
	"time"
)

// CodespacesSecret Secrets for a GitHub Codespace.
type CodespacesSecret struct {
	// The name of the secret
	Name string `json:"name"`
	// The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	CreatedAt time.Time `json:"created_at"`
	// The date and time at which the secret was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	UpdatedAt time.Time `json:"updated_at"`
	// The type of repositories in the organization that the secret is visible to
	Visibility string `json:"visibility"`
	// The API URL at which the list of repositories this secret is visible to can be retrieved
	SelectedRepositoriesUrl string `json:"selected_repositories_url"`
}

// NewCodespacesSecret instantiates a new CodespacesSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodespacesSecret(name string, createdAt time.Time, updatedAt time.Time, visibility string, selectedRepositoriesUrl string) *CodespacesSecret {
	this := CodespacesSecret{}
	this.Name = name
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Visibility = visibility
	this.SelectedRepositoriesUrl = selectedRepositoriesUrl
	return &this
}

// NewCodespacesSecretWithDefaults instantiates a new CodespacesSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodespacesSecretWithDefaults() *CodespacesSecret {
	this := CodespacesSecret{}
	return &this
}

// GetName returns the Name field value
func (o *CodespacesSecret) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CodespacesSecret) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CodespacesSecret) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CodespacesSecret) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CodespacesSecret) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CodespacesSecret) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CodespacesSecret) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CodespacesSecret) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CodespacesSecret) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetVisibility returns the Visibility field value
func (o *CodespacesSecret) GetVisibility() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *CodespacesSecret) GetVisibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *CodespacesSecret) SetVisibility(v string) {
	o.Visibility = v
}

// GetSelectedRepositoriesUrl returns the SelectedRepositoriesUrl field value
func (o *CodespacesSecret) GetSelectedRepositoriesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SelectedRepositoriesUrl
}

// GetSelectedRepositoriesUrlOk returns a tuple with the SelectedRepositoriesUrl field value
// and a boolean to check if the value has been set.
func (o *CodespacesSecret) GetSelectedRepositoriesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectedRepositoriesUrl, true
}

// SetSelectedRepositoriesUrl sets field value
func (o *CodespacesSecret) SetSelectedRepositoriesUrl(v string) {
	o.SelectedRepositoriesUrl = v
}

func (o CodespacesSecret) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["visibility"] = o.Visibility
	}
	if true {
		toSerialize["selected_repositories_url"] = o.SelectedRepositoriesUrl
	}
	return json.Marshal(toSerialize)
}

type NullableCodespacesSecret struct {
	value *CodespacesSecret
	isSet bool
}

func (v NullableCodespacesSecret) Get() *CodespacesSecret {
	return v.value
}

func (v *NullableCodespacesSecret) Set(val *CodespacesSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableCodespacesSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableCodespacesSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodespacesSecret(val *CodespacesSecret) *NullableCodespacesSecret {
	return &NullableCodespacesSecret{value: val, isSet: true}
}

func (v NullableCodespacesSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodespacesSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

