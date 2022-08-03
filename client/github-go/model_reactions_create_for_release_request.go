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

// ReactionsCreateForReleaseRequest struct for ReactionsCreateForReleaseRequest
type ReactionsCreateForReleaseRequest struct {
	// The [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types) to add to the release.
	Content string `json:"content"`
}

// NewReactionsCreateForReleaseRequest instantiates a new ReactionsCreateForReleaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactionsCreateForReleaseRequest(content string) *ReactionsCreateForReleaseRequest {
	this := ReactionsCreateForReleaseRequest{}
	this.Content = content
	return &this
}

// NewReactionsCreateForReleaseRequestWithDefaults instantiates a new ReactionsCreateForReleaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactionsCreateForReleaseRequestWithDefaults() *ReactionsCreateForReleaseRequest {
	this := ReactionsCreateForReleaseRequest{}
	return &this
}

// GetContent returns the Content field value
func (o *ReactionsCreateForReleaseRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ReactionsCreateForReleaseRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *ReactionsCreateForReleaseRequest) SetContent(v string) {
	o.Content = v
}

func (o ReactionsCreateForReleaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableReactionsCreateForReleaseRequest struct {
	value *ReactionsCreateForReleaseRequest
	isSet bool
}

func (v NullableReactionsCreateForReleaseRequest) Get() *ReactionsCreateForReleaseRequest {
	return v.value
}

func (v *NullableReactionsCreateForReleaseRequest) Set(val *ReactionsCreateForReleaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReactionsCreateForReleaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReactionsCreateForReleaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactionsCreateForReleaseRequest(val *ReactionsCreateForReleaseRequest) *NullableReactionsCreateForReleaseRequest {
	return &NullableReactionsCreateForReleaseRequest{value: val, isSet: true}
}

func (v NullableReactionsCreateForReleaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactionsCreateForReleaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


