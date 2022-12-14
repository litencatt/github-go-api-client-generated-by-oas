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

// ReactionsCreateForPullRequestReviewCommentRequest struct for ReactionsCreateForPullRequestReviewCommentRequest
type ReactionsCreateForPullRequestReviewCommentRequest struct {
	// The [reaction type](https://docs.github.com/rest/reference/reactions#reaction-types) to add to the pull request review comment.
	Content string `json:"content"`
}

// NewReactionsCreateForPullRequestReviewCommentRequest instantiates a new ReactionsCreateForPullRequestReviewCommentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactionsCreateForPullRequestReviewCommentRequest(content string) *ReactionsCreateForPullRequestReviewCommentRequest {
	this := ReactionsCreateForPullRequestReviewCommentRequest{}
	this.Content = content
	return &this
}

// NewReactionsCreateForPullRequestReviewCommentRequestWithDefaults instantiates a new ReactionsCreateForPullRequestReviewCommentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactionsCreateForPullRequestReviewCommentRequestWithDefaults() *ReactionsCreateForPullRequestReviewCommentRequest {
	this := ReactionsCreateForPullRequestReviewCommentRequest{}
	return &this
}

// GetContent returns the Content field value
func (o *ReactionsCreateForPullRequestReviewCommentRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ReactionsCreateForPullRequestReviewCommentRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *ReactionsCreateForPullRequestReviewCommentRequest) SetContent(v string) {
	o.Content = v
}

func (o ReactionsCreateForPullRequestReviewCommentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableReactionsCreateForPullRequestReviewCommentRequest struct {
	value *ReactionsCreateForPullRequestReviewCommentRequest
	isSet bool
}

func (v NullableReactionsCreateForPullRequestReviewCommentRequest) Get() *ReactionsCreateForPullRequestReviewCommentRequest {
	return v.value
}

func (v *NullableReactionsCreateForPullRequestReviewCommentRequest) Set(val *ReactionsCreateForPullRequestReviewCommentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReactionsCreateForPullRequestReviewCommentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReactionsCreateForPullRequestReviewCommentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactionsCreateForPullRequestReviewCommentRequest(val *ReactionsCreateForPullRequestReviewCommentRequest) *NullableReactionsCreateForPullRequestReviewCommentRequest {
	return &NullableReactionsCreateForPullRequestReviewCommentRequest{value: val, isSet: true}
}

func (v NullableReactionsCreateForPullRequestReviewCommentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactionsCreateForPullRequestReviewCommentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


