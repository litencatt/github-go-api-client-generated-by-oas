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

// ViewTraffic View Traffic
type ViewTraffic struct {
	Count int32 `json:"count"`
	Uniques int32 `json:"uniques"`
	Views []Traffic `json:"views"`
}

// NewViewTraffic instantiates a new ViewTraffic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewTraffic(count int32, uniques int32, views []Traffic) *ViewTraffic {
	this := ViewTraffic{}
	this.Count = count
	this.Uniques = uniques
	this.Views = views
	return &this
}

// NewViewTrafficWithDefaults instantiates a new ViewTraffic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewTrafficWithDefaults() *ViewTraffic {
	this := ViewTraffic{}
	return &this
}

// GetCount returns the Count field value
func (o *ViewTraffic) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ViewTraffic) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ViewTraffic) SetCount(v int32) {
	o.Count = v
}

// GetUniques returns the Uniques field value
func (o *ViewTraffic) GetUniques() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Uniques
}

// GetUniquesOk returns a tuple with the Uniques field value
// and a boolean to check if the value has been set.
func (o *ViewTraffic) GetUniquesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uniques, true
}

// SetUniques sets field value
func (o *ViewTraffic) SetUniques(v int32) {
	o.Uniques = v
}

// GetViews returns the Views field value
func (o *ViewTraffic) GetViews() []Traffic {
	if o == nil {
		var ret []Traffic
		return ret
	}

	return o.Views
}

// GetViewsOk returns a tuple with the Views field value
// and a boolean to check if the value has been set.
func (o *ViewTraffic) GetViewsOk() ([]Traffic, bool) {
	if o == nil {
		return nil, false
	}
	return o.Views, true
}

// SetViews sets field value
func (o *ViewTraffic) SetViews(v []Traffic) {
	o.Views = v
}

func (o ViewTraffic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["uniques"] = o.Uniques
	}
	if true {
		toSerialize["views"] = o.Views
	}
	return json.Marshal(toSerialize)
}

type NullableViewTraffic struct {
	value *ViewTraffic
	isSet bool
}

func (v NullableViewTraffic) Get() *ViewTraffic {
	return v.value
}

func (v *NullableViewTraffic) Set(val *ViewTraffic) {
	v.value = val
	v.isSet = true
}

func (v NullableViewTraffic) IsSet() bool {
	return v.isSet
}

func (v *NullableViewTraffic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewTraffic(val *ViewTraffic) *NullableViewTraffic {
	return &NullableViewTraffic{value: val, isSet: true}
}

func (v NullableViewTraffic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewTraffic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


