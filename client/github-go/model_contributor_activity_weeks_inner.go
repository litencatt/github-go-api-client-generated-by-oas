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

// ContributorActivityWeeksInner struct for ContributorActivityWeeksInner
type ContributorActivityWeeksInner struct {
	W *int32 `json:"w,omitempty"`
	A *int32 `json:"a,omitempty"`
	D *int32 `json:"d,omitempty"`
	C *int32 `json:"c,omitempty"`
}

// NewContributorActivityWeeksInner instantiates a new ContributorActivityWeeksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContributorActivityWeeksInner() *ContributorActivityWeeksInner {
	this := ContributorActivityWeeksInner{}
	return &this
}

// NewContributorActivityWeeksInnerWithDefaults instantiates a new ContributorActivityWeeksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContributorActivityWeeksInnerWithDefaults() *ContributorActivityWeeksInner {
	this := ContributorActivityWeeksInner{}
	return &this
}

// GetW returns the W field value if set, zero value otherwise.
func (o *ContributorActivityWeeksInner) GetW() int32 {
	if o == nil || o.W == nil {
		var ret int32
		return ret
	}
	return *o.W
}

// GetWOk returns a tuple with the W field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContributorActivityWeeksInner) GetWOk() (*int32, bool) {
	if o == nil || o.W == nil {
		return nil, false
	}
	return o.W, true
}

// HasW returns a boolean if a field has been set.
func (o *ContributorActivityWeeksInner) HasW() bool {
	if o != nil && o.W != nil {
		return true
	}

	return false
}

// SetW gets a reference to the given int32 and assigns it to the W field.
func (o *ContributorActivityWeeksInner) SetW(v int32) {
	o.W = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *ContributorActivityWeeksInner) GetA() int32 {
	if o == nil || o.A == nil {
		var ret int32
		return ret
	}
	return *o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContributorActivityWeeksInner) GetAOk() (*int32, bool) {
	if o == nil || o.A == nil {
		return nil, false
	}
	return o.A, true
}

// HasA returns a boolean if a field has been set.
func (o *ContributorActivityWeeksInner) HasA() bool {
	if o != nil && o.A != nil {
		return true
	}

	return false
}

// SetA gets a reference to the given int32 and assigns it to the A field.
func (o *ContributorActivityWeeksInner) SetA(v int32) {
	o.A = &v
}

// GetD returns the D field value if set, zero value otherwise.
func (o *ContributorActivityWeeksInner) GetD() int32 {
	if o == nil || o.D == nil {
		var ret int32
		return ret
	}
	return *o.D
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContributorActivityWeeksInner) GetDOk() (*int32, bool) {
	if o == nil || o.D == nil {
		return nil, false
	}
	return o.D, true
}

// HasD returns a boolean if a field has been set.
func (o *ContributorActivityWeeksInner) HasD() bool {
	if o != nil && o.D != nil {
		return true
	}

	return false
}

// SetD gets a reference to the given int32 and assigns it to the D field.
func (o *ContributorActivityWeeksInner) SetD(v int32) {
	o.D = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *ContributorActivityWeeksInner) GetC() int32 {
	if o == nil || o.C == nil {
		var ret int32
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContributorActivityWeeksInner) GetCOk() (*int32, bool) {
	if o == nil || o.C == nil {
		return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *ContributorActivityWeeksInner) HasC() bool {
	if o != nil && o.C != nil {
		return true
	}

	return false
}

// SetC gets a reference to the given int32 and assigns it to the C field.
func (o *ContributorActivityWeeksInner) SetC(v int32) {
	o.C = &v
}

func (o ContributorActivityWeeksInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.W != nil {
		toSerialize["w"] = o.W
	}
	if o.A != nil {
		toSerialize["a"] = o.A
	}
	if o.D != nil {
		toSerialize["d"] = o.D
	}
	if o.C != nil {
		toSerialize["c"] = o.C
	}
	return json.Marshal(toSerialize)
}

type NullableContributorActivityWeeksInner struct {
	value *ContributorActivityWeeksInner
	isSet bool
}

func (v NullableContributorActivityWeeksInner) Get() *ContributorActivityWeeksInner {
	return v.value
}

func (v *NullableContributorActivityWeeksInner) Set(val *ContributorActivityWeeksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableContributorActivityWeeksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableContributorActivityWeeksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContributorActivityWeeksInner(val *ContributorActivityWeeksInner) *NullableContributorActivityWeeksInner {
	return &NullableContributorActivityWeeksInner{value: val, isSet: true}
}

func (v NullableContributorActivityWeeksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContributorActivityWeeksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

