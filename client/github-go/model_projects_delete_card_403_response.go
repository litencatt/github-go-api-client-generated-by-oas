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

// ProjectsDeleteCard403Response struct for ProjectsDeleteCard403Response
type ProjectsDeleteCard403Response struct {
	Message *string `json:"message,omitempty"`
	DocumentationUrl *string `json:"documentation_url,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

// NewProjectsDeleteCard403Response instantiates a new ProjectsDeleteCard403Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectsDeleteCard403Response() *ProjectsDeleteCard403Response {
	this := ProjectsDeleteCard403Response{}
	return &this
}

// NewProjectsDeleteCard403ResponseWithDefaults instantiates a new ProjectsDeleteCard403Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectsDeleteCard403ResponseWithDefaults() *ProjectsDeleteCard403Response {
	this := ProjectsDeleteCard403Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ProjectsDeleteCard403Response) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsDeleteCard403Response) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ProjectsDeleteCard403Response) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ProjectsDeleteCard403Response) SetMessage(v string) {
	o.Message = &v
}

// GetDocumentationUrl returns the DocumentationUrl field value if set, zero value otherwise.
func (o *ProjectsDeleteCard403Response) GetDocumentationUrl() string {
	if o == nil || o.DocumentationUrl == nil {
		var ret string
		return ret
	}
	return *o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsDeleteCard403Response) GetDocumentationUrlOk() (*string, bool) {
	if o == nil || o.DocumentationUrl == nil {
		return nil, false
	}
	return o.DocumentationUrl, true
}

// HasDocumentationUrl returns a boolean if a field has been set.
func (o *ProjectsDeleteCard403Response) HasDocumentationUrl() bool {
	if o != nil && o.DocumentationUrl != nil {
		return true
	}

	return false
}

// SetDocumentationUrl gets a reference to the given string and assigns it to the DocumentationUrl field.
func (o *ProjectsDeleteCard403Response) SetDocumentationUrl(v string) {
	o.DocumentationUrl = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ProjectsDeleteCard403Response) GetErrors() []string {
	if o == nil || o.Errors == nil {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectsDeleteCard403Response) GetErrorsOk() ([]string, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ProjectsDeleteCard403Response) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *ProjectsDeleteCard403Response) SetErrors(v []string) {
	o.Errors = v
}

func (o ProjectsDeleteCard403Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.DocumentationUrl != nil {
		toSerialize["documentation_url"] = o.DocumentationUrl
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableProjectsDeleteCard403Response struct {
	value *ProjectsDeleteCard403Response
	isSet bool
}

func (v NullableProjectsDeleteCard403Response) Get() *ProjectsDeleteCard403Response {
	return v.value
}

func (v *NullableProjectsDeleteCard403Response) Set(val *ProjectsDeleteCard403Response) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectsDeleteCard403Response) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectsDeleteCard403Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectsDeleteCard403Response(val *ProjectsDeleteCard403Response) *NullableProjectsDeleteCard403Response {
	return &NullableProjectsDeleteCard403Response{value: val, isSet: true}
}

func (v NullableProjectsDeleteCard403Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectsDeleteCard403Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

