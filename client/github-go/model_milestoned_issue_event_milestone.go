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

// MilestonedIssueEventMilestone struct for MilestonedIssueEventMilestone
type MilestonedIssueEventMilestone struct {
	Title string `json:"title"`
}

// NewMilestonedIssueEventMilestone instantiates a new MilestonedIssueEventMilestone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMilestonedIssueEventMilestone(title string) *MilestonedIssueEventMilestone {
	this := MilestonedIssueEventMilestone{}
	this.Title = title
	return &this
}

// NewMilestonedIssueEventMilestoneWithDefaults instantiates a new MilestonedIssueEventMilestone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMilestonedIssueEventMilestoneWithDefaults() *MilestonedIssueEventMilestone {
	this := MilestonedIssueEventMilestone{}
	return &this
}

// GetTitle returns the Title field value
func (o *MilestonedIssueEventMilestone) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *MilestonedIssueEventMilestone) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *MilestonedIssueEventMilestone) SetTitle(v string) {
	o.Title = v
}

func (o MilestonedIssueEventMilestone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableMilestonedIssueEventMilestone struct {
	value *MilestonedIssueEventMilestone
	isSet bool
}

func (v NullableMilestonedIssueEventMilestone) Get() *MilestonedIssueEventMilestone {
	return v.value
}

func (v *NullableMilestonedIssueEventMilestone) Set(val *MilestonedIssueEventMilestone) {
	v.value = val
	v.isSet = true
}

func (v NullableMilestonedIssueEventMilestone) IsSet() bool {
	return v.isSet
}

func (v *NullableMilestonedIssueEventMilestone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMilestonedIssueEventMilestone(val *MilestonedIssueEventMilestone) *NullableMilestonedIssueEventMilestone {
	return &NullableMilestonedIssueEventMilestone{value: val, isSet: true}
}

func (v NullableMilestonedIssueEventMilestone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMilestonedIssueEventMilestone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


