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

// ValidationErrorSimple Validation Error Simple
type ValidationErrorSimple struct {
	Message string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
	Errors []string `json:"errors,omitempty"`
}

// NewValidationErrorSimple instantiates a new ValidationErrorSimple object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationErrorSimple(message string, documentationUrl string) *ValidationErrorSimple {
	this := ValidationErrorSimple{}
	this.Message = message
	this.DocumentationUrl = documentationUrl
	return &this
}

// NewValidationErrorSimpleWithDefaults instantiates a new ValidationErrorSimple object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationErrorSimpleWithDefaults() *ValidationErrorSimple {
	this := ValidationErrorSimple{}
	return &this
}

// GetMessage returns the Message field value
func (o *ValidationErrorSimple) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ValidationErrorSimple) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ValidationErrorSimple) SetMessage(v string) {
	o.Message = v
}

// GetDocumentationUrl returns the DocumentationUrl field value
func (o *ValidationErrorSimple) GetDocumentationUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value
// and a boolean to check if the value has been set.
func (o *ValidationErrorSimple) GetDocumentationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentationUrl, true
}

// SetDocumentationUrl sets field value
func (o *ValidationErrorSimple) SetDocumentationUrl(v string) {
	o.DocumentationUrl = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ValidationErrorSimple) GetErrors() []string {
	if o == nil || o.Errors == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationErrorSimple) GetErrorsOk() ([]string, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ValidationErrorSimple) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *ValidationErrorSimple) SetErrors(v []string) {
	o.Errors = v
}

func (o ValidationErrorSimple) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["documentation_url"] = o.DocumentationUrl
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableValidationErrorSimple struct {
	value *ValidationErrorSimple
	isSet bool
}

func (v NullableValidationErrorSimple) Get() *ValidationErrorSimple {
	return v.value
}

func (v *NullableValidationErrorSimple) Set(val *ValidationErrorSimple) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationErrorSimple) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationErrorSimple) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationErrorSimple(val *ValidationErrorSimple) *NullableValidationErrorSimple {
	return &NullableValidationErrorSimple{value: val, isSet: true}
}

func (v NullableValidationErrorSimple) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationErrorSimple) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

