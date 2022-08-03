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

// PullsUpdateReviewCommentRequest struct for PullsUpdateReviewCommentRequest
type PullsUpdateReviewCommentRequest struct {
	// The text of the reply to the review comment.
	Body string `json:"body"`
}

// NewPullsUpdateReviewCommentRequest instantiates a new PullsUpdateReviewCommentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPullsUpdateReviewCommentRequest(body string) *PullsUpdateReviewCommentRequest {
	this := PullsUpdateReviewCommentRequest{}
	this.Body = body
	return &this
}

// NewPullsUpdateReviewCommentRequestWithDefaults instantiates a new PullsUpdateReviewCommentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPullsUpdateReviewCommentRequestWithDefaults() *PullsUpdateReviewCommentRequest {
	this := PullsUpdateReviewCommentRequest{}
	return &this
}

// GetBody returns the Body field value
func (o *PullsUpdateReviewCommentRequest) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *PullsUpdateReviewCommentRequest) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *PullsUpdateReviewCommentRequest) SetBody(v string) {
	o.Body = v
}

func (o PullsUpdateReviewCommentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullablePullsUpdateReviewCommentRequest struct {
	value *PullsUpdateReviewCommentRequest
	isSet bool
}

func (v NullablePullsUpdateReviewCommentRequest) Get() *PullsUpdateReviewCommentRequest {
	return v.value
}

func (v *NullablePullsUpdateReviewCommentRequest) Set(val *PullsUpdateReviewCommentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePullsUpdateReviewCommentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePullsUpdateReviewCommentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullsUpdateReviewCommentRequest(val *PullsUpdateReviewCommentRequest) *NullablePullsUpdateReviewCommentRequest {
	return &NullablePullsUpdateReviewCommentRequest{value: val, isSet: true}
}

func (v NullablePullsUpdateReviewCommentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullsUpdateReviewCommentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

