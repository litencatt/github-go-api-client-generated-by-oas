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

// RenamedIssueEventRename struct for RenamedIssueEventRename
type RenamedIssueEventRename struct {
	From string `json:"from"`
	To string `json:"to"`
}

// NewRenamedIssueEventRename instantiates a new RenamedIssueEventRename object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenamedIssueEventRename(from string, to string) *RenamedIssueEventRename {
	this := RenamedIssueEventRename{}
	this.From = from
	this.To = to
	return &this
}

// NewRenamedIssueEventRenameWithDefaults instantiates a new RenamedIssueEventRename object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenamedIssueEventRenameWithDefaults() *RenamedIssueEventRename {
	this := RenamedIssueEventRename{}
	return &this
}

// GetFrom returns the From field value
func (o *RenamedIssueEventRename) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *RenamedIssueEventRename) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *RenamedIssueEventRename) SetFrom(v string) {
	o.From = v
}

// GetTo returns the To field value
func (o *RenamedIssueEventRename) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *RenamedIssueEventRename) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *RenamedIssueEventRename) SetTo(v string) {
	o.To = v
}

func (o RenamedIssueEventRename) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableRenamedIssueEventRename struct {
	value *RenamedIssueEventRename
	isSet bool
}

func (v NullableRenamedIssueEventRename) Get() *RenamedIssueEventRename {
	return v.value
}

func (v *NullableRenamedIssueEventRename) Set(val *RenamedIssueEventRename) {
	v.value = val
	v.isSet = true
}

func (v NullableRenamedIssueEventRename) IsSet() bool {
	return v.isSet
}

func (v *NullableRenamedIssueEventRename) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenamedIssueEventRename(val *RenamedIssueEventRename) *NullableRenamedIssueEventRename {
	return &NullableRenamedIssueEventRename{value: val, isSet: true}
}

func (v NullableRenamedIssueEventRename) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenamedIssueEventRename) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

