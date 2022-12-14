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

// ReactionsCreateForCommitCommentRequest struct for ReactionsCreateForCommitCommentRequest
type ReactionsCreateForCommitCommentRequest struct {
	// The [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types) to add to the commit comment.
	Content string `json:"content"`
}

// NewReactionsCreateForCommitCommentRequest instantiates a new ReactionsCreateForCommitCommentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactionsCreateForCommitCommentRequest(content string) *ReactionsCreateForCommitCommentRequest {
	this := ReactionsCreateForCommitCommentRequest{}
	this.Content = content
	return &this
}

// NewReactionsCreateForCommitCommentRequestWithDefaults instantiates a new ReactionsCreateForCommitCommentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactionsCreateForCommitCommentRequestWithDefaults() *ReactionsCreateForCommitCommentRequest {
	this := ReactionsCreateForCommitCommentRequest{}
	return &this
}

// GetContent returns the Content field value
func (o *ReactionsCreateForCommitCommentRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ReactionsCreateForCommitCommentRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *ReactionsCreateForCommitCommentRequest) SetContent(v string) {
	o.Content = v
}

func (o ReactionsCreateForCommitCommentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableReactionsCreateForCommitCommentRequest struct {
	value *ReactionsCreateForCommitCommentRequest
	isSet bool
}

func (v NullableReactionsCreateForCommitCommentRequest) Get() *ReactionsCreateForCommitCommentRequest {
	return v.value
}

func (v *NullableReactionsCreateForCommitCommentRequest) Set(val *ReactionsCreateForCommitCommentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReactionsCreateForCommitCommentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReactionsCreateForCommitCommentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactionsCreateForCommitCommentRequest(val *ReactionsCreateForCommitCommentRequest) *NullableReactionsCreateForCommitCommentRequest {
	return &NullableReactionsCreateForCommitCommentRequest{value: val, isSet: true}
}

func (v NullableReactionsCreateForCommitCommentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactionsCreateForCommitCommentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


