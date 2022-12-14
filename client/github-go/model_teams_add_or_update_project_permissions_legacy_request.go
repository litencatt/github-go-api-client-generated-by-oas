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

// TeamsAddOrUpdateProjectPermissionsLegacyRequest struct for TeamsAddOrUpdateProjectPermissionsLegacyRequest
type TeamsAddOrUpdateProjectPermissionsLegacyRequest struct {
	// The permission to grant to the team for this project. Default: the team's `permission` attribute will be used to determine what permission to grant the team on this project. Note that, if you choose not to pass any parameters, you'll need to set `Content-Length` to zero when calling this endpoint. For more information, see \"[HTTP verbs](https://docs.github.com/rest/overview/resources-in-the-rest-api#http-verbs).\"
	Permission *string `json:"permission,omitempty"`
}

// NewTeamsAddOrUpdateProjectPermissionsLegacyRequest instantiates a new TeamsAddOrUpdateProjectPermissionsLegacyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamsAddOrUpdateProjectPermissionsLegacyRequest() *TeamsAddOrUpdateProjectPermissionsLegacyRequest {
	this := TeamsAddOrUpdateProjectPermissionsLegacyRequest{}
	return &this
}

// NewTeamsAddOrUpdateProjectPermissionsLegacyRequestWithDefaults instantiates a new TeamsAddOrUpdateProjectPermissionsLegacyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamsAddOrUpdateProjectPermissionsLegacyRequestWithDefaults() *TeamsAddOrUpdateProjectPermissionsLegacyRequest {
	this := TeamsAddOrUpdateProjectPermissionsLegacyRequest{}
	return &this
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *TeamsAddOrUpdateProjectPermissionsLegacyRequest) GetPermission() string {
	if o == nil || o.Permission == nil {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsAddOrUpdateProjectPermissionsLegacyRequest) GetPermissionOk() (*string, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *TeamsAddOrUpdateProjectPermissionsLegacyRequest) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *TeamsAddOrUpdateProjectPermissionsLegacyRequest) SetPermission(v string) {
	o.Permission = &v
}

func (o TeamsAddOrUpdateProjectPermissionsLegacyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Permission != nil {
		toSerialize["permission"] = o.Permission
	}
	return json.Marshal(toSerialize)
}

type NullableTeamsAddOrUpdateProjectPermissionsLegacyRequest struct {
	value *TeamsAddOrUpdateProjectPermissionsLegacyRequest
	isSet bool
}

func (v NullableTeamsAddOrUpdateProjectPermissionsLegacyRequest) Get() *TeamsAddOrUpdateProjectPermissionsLegacyRequest {
	return v.value
}

func (v *NullableTeamsAddOrUpdateProjectPermissionsLegacyRequest) Set(val *TeamsAddOrUpdateProjectPermissionsLegacyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamsAddOrUpdateProjectPermissionsLegacyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamsAddOrUpdateProjectPermissionsLegacyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamsAddOrUpdateProjectPermissionsLegacyRequest(val *TeamsAddOrUpdateProjectPermissionsLegacyRequest) *NullableTeamsAddOrUpdateProjectPermissionsLegacyRequest {
	return &NullableTeamsAddOrUpdateProjectPermissionsLegacyRequest{value: val, isSet: true}
}

func (v NullableTeamsAddOrUpdateProjectPermissionsLegacyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamsAddOrUpdateProjectPermissionsLegacyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


