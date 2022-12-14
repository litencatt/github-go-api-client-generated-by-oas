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

// OptOutOidcCustomSub OIDC Customer Subject
type OptOutOidcCustomSub struct {
	UseDefault bool `json:"use_default"`
}

// NewOptOutOidcCustomSub instantiates a new OptOutOidcCustomSub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptOutOidcCustomSub(useDefault bool) *OptOutOidcCustomSub {
	this := OptOutOidcCustomSub{}
	this.UseDefault = useDefault
	return &this
}

// NewOptOutOidcCustomSubWithDefaults instantiates a new OptOutOidcCustomSub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptOutOidcCustomSubWithDefaults() *OptOutOidcCustomSub {
	this := OptOutOidcCustomSub{}
	return &this
}

// GetUseDefault returns the UseDefault field value
func (o *OptOutOidcCustomSub) GetUseDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseDefault
}

// GetUseDefaultOk returns a tuple with the UseDefault field value
// and a boolean to check if the value has been set.
func (o *OptOutOidcCustomSub) GetUseDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseDefault, true
}

// SetUseDefault sets field value
func (o *OptOutOidcCustomSub) SetUseDefault(v bool) {
	o.UseDefault = v
}

func (o OptOutOidcCustomSub) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["use_default"] = o.UseDefault
	}
	return json.Marshal(toSerialize)
}

type NullableOptOutOidcCustomSub struct {
	value *OptOutOidcCustomSub
	isSet bool
}

func (v NullableOptOutOidcCustomSub) Get() *OptOutOidcCustomSub {
	return v.value
}

func (v *NullableOptOutOidcCustomSub) Set(val *OptOutOidcCustomSub) {
	v.value = val
	v.isSet = true
}

func (v NullableOptOutOidcCustomSub) IsSet() bool {
	return v.isSet
}

func (v *NullableOptOutOidcCustomSub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptOutOidcCustomSub(val *OptOutOidcCustomSub) *NullableOptOutOidcCustomSub {
	return &NullableOptOutOidcCustomSub{value: val, isSet: true}
}

func (v NullableOptOutOidcCustomSub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptOutOidcCustomSub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


