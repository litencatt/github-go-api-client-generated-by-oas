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

// ProjectsUpdateRequest struct for ProjectsUpdateRequest
type ProjectsUpdateRequest struct {
	// Name of the project
	Name *string `json:"name,omitempty"`
	// Body of the project
	Body NullableString `json:"body,omitempty"`
	// State of the project; either 'open' or 'closed'
	State *string `json:"state,omitempty"`
	// The baseline permission that all organization members have on this project
	OrganizationPermission *string `json:"organization_permission,omitempty"`
	// Whether or not this project can be seen by everyone.
	Private *bool `json:"private,omitempty"`
}

// NewProjectsUpdateRequest instantiates a new ProjectsUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsUpdateRequest() *ProjectsUpdateRequest {
	this := ProjectsUpdateRequest{}
	return &this
}

// NewProjectsUpdateRequestWithDefaults instantiates a new ProjectsUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsUpdateRequestWithDefaults() *ProjectsUpdateRequest {
	this := ProjectsUpdateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectsUpdateRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectsUpdateRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectsUpdateRequest) SetName(v string) {
	o.Name = &v
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectsUpdateRequest) GetBody() string {
	if o == nil || o.Body.Get() == nil {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectsUpdateRequest) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *ProjectsUpdateRequest) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *ProjectsUpdateRequest) SetBody(v string) {
	o.Body.Set(&v)
}
// SetBodyNil sets the value for Body to be an explicit nil
func (o *ProjectsUpdateRequest) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *ProjectsUpdateRequest) UnsetBody() {
	o.Body.Unset()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ProjectsUpdateRequest) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsUpdateRequest) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ProjectsUpdateRequest) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ProjectsUpdateRequest) SetState(v string) {
	o.State = &v
}

// GetOrganizationPermission returns the OrganizationPermission field value if set, zero value otherwise.
func (o *ProjectsUpdateRequest) GetOrganizationPermission() string {
	if o == nil || o.OrganizationPermission == nil {
		var ret string
		return ret
	}
	return *o.OrganizationPermission
}

// GetOrganizationPermissionOk returns a tuple with the OrganizationPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsUpdateRequest) GetOrganizationPermissionOk() (*string, bool) {
	if o == nil || o.OrganizationPermission == nil {
		return nil, false
	}
	return o.OrganizationPermission, true
}

// HasOrganizationPermission returns a boolean if a field has been set.
func (o *ProjectsUpdateRequest) HasOrganizationPermission() bool {
	if o != nil && o.OrganizationPermission != nil {
		return true
	}

	return false
}

// SetOrganizationPermission gets a reference to the given string and assigns it to the OrganizationPermission field.
func (o *ProjectsUpdateRequest) SetOrganizationPermission(v string) {
	o.OrganizationPermission = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *ProjectsUpdateRequest) GetPrivate() bool {
	if o == nil || o.Private == nil {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsUpdateRequest) GetPrivateOk() (*bool, bool) {
	if o == nil || o.Private == nil {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *ProjectsUpdateRequest) HasPrivate() bool {
	if o != nil && o.Private != nil {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *ProjectsUpdateRequest) SetPrivate(v bool) {
	o.Private = &v
}

func (o ProjectsUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Body.IsSet() {
		toSerialize["body"] = o.Body.Get()
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.OrganizationPermission != nil {
		toSerialize["organization_permission"] = o.OrganizationPermission
	}
	if o.Private != nil {
		toSerialize["private"] = o.Private
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsUpdateRequest struct {
	value *ProjectsUpdateRequest
	isSet bool
}

func (v NullableProjectsUpdateRequest) Get() *ProjectsUpdateRequest {
	return v.value
}

func (v *NullableProjectsUpdateRequest) Set(val *ProjectsUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsUpdateRequest(val *ProjectsUpdateRequest) *NullableProjectsUpdateRequest {
	return &NullableProjectsUpdateRequest{value: val, isSet: true}
}

func (v NullableProjectsUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


