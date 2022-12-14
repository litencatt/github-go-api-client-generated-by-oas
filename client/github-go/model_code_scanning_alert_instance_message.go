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

// CodeScanningAlertInstanceMessage struct for CodeScanningAlertInstanceMessage
type CodeScanningAlertInstanceMessage struct {
	Text *string `json:"text,omitempty"`
}

// NewCodeScanningAlertInstanceMessage instantiates a new CodeScanningAlertInstanceMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodeScanningAlertInstanceMessage() *CodeScanningAlertInstanceMessage {
	this := CodeScanningAlertInstanceMessage{}
	return &this
}

// NewCodeScanningAlertInstanceMessageWithDefaults instantiates a new CodeScanningAlertInstanceMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodeScanningAlertInstanceMessageWithDefaults() *CodeScanningAlertInstanceMessage {
	this := CodeScanningAlertInstanceMessage{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *CodeScanningAlertInstanceMessage) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeScanningAlertInstanceMessage) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *CodeScanningAlertInstanceMessage) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *CodeScanningAlertInstanceMessage) SetText(v string) {
	o.Text = &v
}

func (o CodeScanningAlertInstanceMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableCodeScanningAlertInstanceMessage struct {
	value *CodeScanningAlertInstanceMessage
	isSet bool
}

func (v NullableCodeScanningAlertInstanceMessage) Get() *CodeScanningAlertInstanceMessage {
	return v.value
}

func (v *NullableCodeScanningAlertInstanceMessage) Set(val *CodeScanningAlertInstanceMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeScanningAlertInstanceMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeScanningAlertInstanceMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeScanningAlertInstanceMessage(val *CodeScanningAlertInstanceMessage) *NullableCodeScanningAlertInstanceMessage {
	return &NullableCodeScanningAlertInstanceMessage{value: val, isSet: true}
}

func (v NullableCodeScanningAlertInstanceMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeScanningAlertInstanceMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


