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

// SecurityAndAnalysisAdvancedSecurity struct for SecurityAndAnalysisAdvancedSecurity
type SecurityAndAnalysisAdvancedSecurity struct {
	Status *string `json:"status,omitempty"`
}

// NewSecurityAndAnalysisAdvancedSecurity instantiates a new SecurityAndAnalysisAdvancedSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAndAnalysisAdvancedSecurity() *SecurityAndAnalysisAdvancedSecurity {
	this := SecurityAndAnalysisAdvancedSecurity{}
	return &this
}

// NewSecurityAndAnalysisAdvancedSecurityWithDefaults instantiates a new SecurityAndAnalysisAdvancedSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAndAnalysisAdvancedSecurityWithDefaults() *SecurityAndAnalysisAdvancedSecurity {
	this := SecurityAndAnalysisAdvancedSecurity{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SecurityAndAnalysisAdvancedSecurity) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAndAnalysisAdvancedSecurity) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SecurityAndAnalysisAdvancedSecurity) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SecurityAndAnalysisAdvancedSecurity) SetStatus(v string) {
	o.Status = &v
}

func (o SecurityAndAnalysisAdvancedSecurity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityAndAnalysisAdvancedSecurity struct {
	value *SecurityAndAnalysisAdvancedSecurity
	isSet bool
}

func (v NullableSecurityAndAnalysisAdvancedSecurity) Get() *SecurityAndAnalysisAdvancedSecurity {
	return v.value
}

func (v *NullableSecurityAndAnalysisAdvancedSecurity) Set(val *SecurityAndAnalysisAdvancedSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAndAnalysisAdvancedSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAndAnalysisAdvancedSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAndAnalysisAdvancedSecurity(val *SecurityAndAnalysisAdvancedSecurity) *NullableSecurityAndAnalysisAdvancedSecurity {
	return &NullableSecurityAndAnalysisAdvancedSecurity{value: val, isSet: true}
}

func (v NullableSecurityAndAnalysisAdvancedSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAndAnalysisAdvancedSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

