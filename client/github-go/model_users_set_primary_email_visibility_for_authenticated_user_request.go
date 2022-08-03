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

// UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest struct for UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest
type UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest struct {
	// Denotes whether an email is publicly visible.
	Visibility string `json:"visibility"`
}

// NewUsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest instantiates a new UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest(visibility string) *UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest {
	this := UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest{}
	this.Visibility = visibility
	return &this
}

// NewUsersSetPrimaryEmailVisibilityForAuthenticatedUserRequestWithDefaults instantiates a new UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersSetPrimaryEmailVisibilityForAuthenticatedUserRequestWithDefaults() *UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest {
	this := UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest{}
	return &this
}

// GetVisibility returns the Visibility field value
func (o *UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest) GetVisibility() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest) GetVisibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest) SetVisibility(v string) {
	o.Visibility = v
}

func (o UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["visibility"] = o.Visibility
	}
	return json.Marshal(toSerialize)
}

type NullableUsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest struct {
	value *UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest
	isSet bool
}

func (v NullableUsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest) Get() *UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest {
	return v.value
}

func (v *NullableUsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest) Set(val *UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest(val *UsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest) *NullableUsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest {
	return &NullableUsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest{value: val, isSet: true}
}

func (v NullableUsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersSetPrimaryEmailVisibilityForAuthenticatedUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


