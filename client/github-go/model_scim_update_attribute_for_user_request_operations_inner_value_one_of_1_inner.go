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

// ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner struct for ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner
type ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner struct {
	Value *string `json:"value,omitempty"`
	Primary *bool `json:"primary,omitempty"`
}

// NewScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner instantiates a new ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner() *ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner {
	this := ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner{}
	return &this
}

// NewScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1InnerWithDefaults instantiates a new ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1InnerWithDefaults() *ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner {
	this := ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner) SetValue(v string) {
	o.Value = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner) GetPrimary() bool {
	if o == nil || o.Primary == nil {
		var ret bool
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner) GetPrimaryOk() (*bool, bool) {
	if o == nil || o.Primary == nil {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner) HasPrimary() bool {
	if o != nil && o.Primary != nil {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given bool and assigns it to the Primary field.
func (o *ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner) SetPrimary(v bool) {
	o.Primary = &v
}

func (o ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Primary != nil {
		toSerialize["primary"] = o.Primary
	}
	return json.Marshal(toSerialize)
}

type NullableScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner struct {
	value *ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner
	isSet bool
}

func (v NullableScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner) Get() *ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner {
	return v.value
}

func (v *NullableScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner) Set(val *ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner) {
	v.value = val
	v.isSet = true
}

func (v NullableScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner) IsSet() bool {
	return v.isSet
}

func (v *NullableScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner(val *ScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner) *NullableScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner {
	return &NullableScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner{value: val, isSet: true}
}

func (v NullableScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimUpdateAttributeForUserRequestOperationsInnerValueOneOf1Inner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

