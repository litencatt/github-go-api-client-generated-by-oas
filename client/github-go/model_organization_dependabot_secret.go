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

// OrganizationDependabotSecret Secrets for GitHub Dependabot for an organization.
type OrganizationDependabotSecret struct {
	// The name of the secret.
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Visibility of a secret
	Visibility string `json:"visibility"`
	SelectedRepositoriesUrl *string `json:"selected_repositories_url,omitempty"`
}

// NewOrganizationDependabotSecret instantiates a new OrganizationDependabotSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationDependabotSecret(name string, createdAt time.Time, updatedAt time.Time, visibility string) *OrganizationDependabotSecret {
	this := OrganizationDependabotSecret{}
	this.Name = name
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Visibility = visibility
	return &this
}

// NewOrganizationDependabotSecretWithDefaults instantiates a new OrganizationDependabotSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationDependabotSecretWithDefaults() *OrganizationDependabotSecret {
	this := OrganizationDependabotSecret{}
	return &this
}

// GetName returns the Name field value
func (o *OrganizationDependabotSecret) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationDependabotSecret) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationDependabotSecret) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *OrganizationDependabotSecret) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationDependabotSecret) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OrganizationDependabotSecret) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *OrganizationDependabotSecret) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationDependabotSecret) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *OrganizationDependabotSecret) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetVisibility returns the Visibility field value
func (o *OrganizationDependabotSecret) GetVisibility() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *OrganizationDependabotSecret) GetVisibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *OrganizationDependabotSecret) SetVisibility(v string) {
	o.Visibility = v
}

// GetSelectedRepositoriesUrl returns the SelectedRepositoriesUrl field value if set, zero value otherwise.
func (o *OrganizationDependabotSecret) GetSelectedRepositoriesUrl() string {
	if o == nil || o.SelectedRepositoriesUrl == nil {
		var ret string
		return ret
	}
	return *o.SelectedRepositoriesUrl
}

// GetSelectedRepositoriesUrlOk returns a tuple with the SelectedRepositoriesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationDependabotSecret) GetSelectedRepositoriesUrlOk() (*string, bool) {
	if o == nil || o.SelectedRepositoriesUrl == nil {
		return nil, false
	}
	return o.SelectedRepositoriesUrl, true
}

// HasSelectedRepositoriesUrl returns a boolean if a field has been set.
func (o *OrganizationDependabotSecret) HasSelectedRepositoriesUrl() bool {
	if o != nil && o.SelectedRepositoriesUrl != nil {
		return true
	}

	return false
}

// SetSelectedRepositoriesUrl gets a reference to the given string and assigns it to the SelectedRepositoriesUrl field.
func (o *OrganizationDependabotSecret) SetSelectedRepositoriesUrl(v string) {
	o.SelectedRepositoriesUrl = &v
}

func (o OrganizationDependabotSecret) MarshalJSON() ([]byte, error) {
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
	if o.SelectedRepositoriesUrl != nil {
		toSerialize["selected_repositories_url"] = o.SelectedRepositoriesUrl
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationDependabotSecret struct {
	value *OrganizationDependabotSecret
	isSet bool
}

func (v NullableOrganizationDependabotSecret) Get() *OrganizationDependabotSecret {
	return v.value
}

func (v *NullableOrganizationDependabotSecret) Set(val *OrganizationDependabotSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationDependabotSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationDependabotSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationDependabotSecret(val *OrganizationDependabotSecret) *NullableOrganizationDependabotSecret {
	return &NullableOrganizationDependabotSecret{value: val, isSet: true}
}

func (v NullableOrganizationDependabotSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationDependabotSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

