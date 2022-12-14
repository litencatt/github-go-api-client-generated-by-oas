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

// ReleaseNotesContent Generated name and body describing a release
type ReleaseNotesContent struct {
	// The generated name of the release
	Name string `json:"name"`
	// The generated body describing the contents of the release supporting markdown formatting
	Body string `json:"body"`
}

// NewReleaseNotesContent instantiates a new ReleaseNotesContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseNotesContent(name string, body string) *ReleaseNotesContent {
	this := ReleaseNotesContent{}
	this.Name = name
	this.Body = body
	return &this
}

// NewReleaseNotesContentWithDefaults instantiates a new ReleaseNotesContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseNotesContentWithDefaults() *ReleaseNotesContent {
	this := ReleaseNotesContent{}
	return &this
}

// GetName returns the Name field value
func (o *ReleaseNotesContent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReleaseNotesContent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReleaseNotesContent) SetName(v string) {
	o.Name = v
}

// GetBody returns the Body field value
func (o *ReleaseNotesContent) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *ReleaseNotesContent) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *ReleaseNotesContent) SetBody(v string) {
	o.Body = v
}

func (o ReleaseNotesContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableReleaseNotesContent struct {
	value *ReleaseNotesContent
	isSet bool
}

func (v NullableReleaseNotesContent) Get() *ReleaseNotesContent {
	return v.value
}

func (v *NullableReleaseNotesContent) Set(val *ReleaseNotesContent) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseNotesContent) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseNotesContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseNotesContent(val *ReleaseNotesContent) *NullableReleaseNotesContent {
	return &NullableReleaseNotesContent{value: val, isSet: true}
}

func (v NullableReleaseNotesContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseNotesContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


