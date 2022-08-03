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

// ProjectsMoveCard403ResponseErrorsInner struct for ProjectsMoveCard403ResponseErrorsInner
type ProjectsMoveCard403ResponseErrorsInner struct {
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Resource *string `json:"resource,omitempty"`
	Field *string `json:"field,omitempty"`
}

// NewProjectsMoveCard403ResponseErrorsInner instantiates a new ProjectsMoveCard403ResponseErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsMoveCard403ResponseErrorsInner() *ProjectsMoveCard403ResponseErrorsInner {
	this := ProjectsMoveCard403ResponseErrorsInner{}
	return &this
}

// NewProjectsMoveCard403ResponseErrorsInnerWithDefaults instantiates a new ProjectsMoveCard403ResponseErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsMoveCard403ResponseErrorsInnerWithDefaults() *ProjectsMoveCard403ResponseErrorsInner {
	this := ProjectsMoveCard403ResponseErrorsInner{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ProjectsMoveCard403ResponseErrorsInner) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsMoveCard403ResponseErrorsInner) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ProjectsMoveCard403ResponseErrorsInner) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ProjectsMoveCard403ResponseErrorsInner) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ProjectsMoveCard403ResponseErrorsInner) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsMoveCard403ResponseErrorsInner) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ProjectsMoveCard403ResponseErrorsInner) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ProjectsMoveCard403ResponseErrorsInner) SetMessage(v string) {
	o.Message = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ProjectsMoveCard403ResponseErrorsInner) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsMoveCard403ResponseErrorsInner) GetResourceOk() (*string, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ProjectsMoveCard403ResponseErrorsInner) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *ProjectsMoveCard403ResponseErrorsInner) SetResource(v string) {
	o.Resource = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *ProjectsMoveCard403ResponseErrorsInner) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsMoveCard403ResponseErrorsInner) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *ProjectsMoveCard403ResponseErrorsInner) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *ProjectsMoveCard403ResponseErrorsInner) SetField(v string) {
	o.Field = &v
}

func (o ProjectsMoveCard403ResponseErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsMoveCard403ResponseErrorsInner struct {
	value *ProjectsMoveCard403ResponseErrorsInner
	isSet bool
}

func (v NullableProjectsMoveCard403ResponseErrorsInner) Get() *ProjectsMoveCard403ResponseErrorsInner {
	return v.value
}

func (v *NullableProjectsMoveCard403ResponseErrorsInner) Set(val *ProjectsMoveCard403ResponseErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsMoveCard403ResponseErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsMoveCard403ResponseErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsMoveCard403ResponseErrorsInner(val *ProjectsMoveCard403ResponseErrorsInner) *NullableProjectsMoveCard403ResponseErrorsInner {
	return &NullableProjectsMoveCard403ResponseErrorsInner{value: val, isSet: true}
}

func (v NullableProjectsMoveCard403ResponseErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsMoveCard403ResponseErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

