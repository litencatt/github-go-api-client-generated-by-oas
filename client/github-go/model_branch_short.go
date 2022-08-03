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

// BranchShort Branch Short
type BranchShort struct {
	Name string `json:"name"`
	Commit BranchShortCommit `json:"commit"`
	Protected bool `json:"protected"`
}

// NewBranchShort instantiates a new BranchShort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBranchShort(name string, commit BranchShortCommit, protected bool) *BranchShort {
	this := BranchShort{}
	this.Name = name
	this.Commit = commit
	this.Protected = protected
	return &this
}

// NewBranchShortWithDefaults instantiates a new BranchShort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBranchShortWithDefaults() *BranchShort {
	this := BranchShort{}
	return &this
}

// GetName returns the Name field value
func (o *BranchShort) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BranchShort) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BranchShort) SetName(v string) {
	o.Name = v
}

// GetCommit returns the Commit field value
func (o *BranchShort) GetCommit() BranchShortCommit {
	if o == nil {
		var ret BranchShortCommit
		return ret
	}

	return o.Commit
}

// GetCommitOk returns a tuple with the Commit field value
// and a boolean to check if the value has been set.
func (o *BranchShort) GetCommitOk() (*BranchShortCommit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commit, true
}

// SetCommit sets field value
func (o *BranchShort) SetCommit(v BranchShortCommit) {
	o.Commit = v
}

// GetProtected returns the Protected field value
func (o *BranchShort) GetProtected() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Protected
}

// GetProtectedOk returns a tuple with the Protected field value
// and a boolean to check if the value has been set.
func (o *BranchShort) GetProtectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protected, true
}

// SetProtected sets field value
func (o *BranchShort) SetProtected(v bool) {
	o.Protected = v
}

func (o BranchShort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["commit"] = o.Commit
	}
	if true {
		toSerialize["protected"] = o.Protected
	}
	return json.Marshal(toSerialize)
}

type NullableBranchShort struct {
	value *BranchShort
	isSet bool
}

func (v NullableBranchShort) Get() *BranchShort {
	return v.value
}

func (v *NullableBranchShort) Set(val *BranchShort) {
	v.value = val
	v.isSet = true
}

func (v NullableBranchShort) IsSet() bool {
	return v.isSet
}

func (v *NullableBranchShort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBranchShort(val *BranchShort) *NullableBranchShort {
	return &NullableBranchShort{value: val, isSet: true}
}

func (v NullableBranchShort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBranchShort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


