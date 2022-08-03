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

// PullRequestReviewCommentLinks struct for PullRequestReviewCommentLinks
type PullRequestReviewCommentLinks struct {
	Self PullRequestReviewCommentLinksSelf `json:"self"`
	Html PullRequestReviewCommentLinksHtml `json:"html"`
	PullRequest PullRequestReviewCommentLinksPullRequest `json:"pull_request"`
}

// NewPullRequestReviewCommentLinks instantiates a new PullRequestReviewCommentLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPullRequestReviewCommentLinks(self PullRequestReviewCommentLinksSelf, html PullRequestReviewCommentLinksHtml, pullRequest PullRequestReviewCommentLinksPullRequest) *PullRequestReviewCommentLinks {
	this := PullRequestReviewCommentLinks{}
	this.Self = self
	this.Html = html
	this.PullRequest = pullRequest
	return &this
}

// NewPullRequestReviewCommentLinksWithDefaults instantiates a new PullRequestReviewCommentLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPullRequestReviewCommentLinksWithDefaults() *PullRequestReviewCommentLinks {
	this := PullRequestReviewCommentLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *PullRequestReviewCommentLinks) GetSelf() PullRequestReviewCommentLinksSelf {
	if o == nil {
		var ret PullRequestReviewCommentLinksSelf
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *PullRequestReviewCommentLinks) GetSelfOk() (*PullRequestReviewCommentLinksSelf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *PullRequestReviewCommentLinks) SetSelf(v PullRequestReviewCommentLinksSelf) {
	o.Self = v
}

// GetHtml returns the Html field value
func (o *PullRequestReviewCommentLinks) GetHtml() PullRequestReviewCommentLinksHtml {
	if o == nil {
		var ret PullRequestReviewCommentLinksHtml
		return ret
	}

	return o.Html
}

// GetHtmlOk returns a tuple with the Html field value
// and a boolean to check if the value has been set.
func (o *PullRequestReviewCommentLinks) GetHtmlOk() (*PullRequestReviewCommentLinksHtml, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Html, true
}

// SetHtml sets field value
func (o *PullRequestReviewCommentLinks) SetHtml(v PullRequestReviewCommentLinksHtml) {
	o.Html = v
}

// GetPullRequest returns the PullRequest field value
func (o *PullRequestReviewCommentLinks) GetPullRequest() PullRequestReviewCommentLinksPullRequest {
	if o == nil {
		var ret PullRequestReviewCommentLinksPullRequest
		return ret
	}

	return o.PullRequest
}

// GetPullRequestOk returns a tuple with the PullRequest field value
// and a boolean to check if the value has been set.
func (o *PullRequestReviewCommentLinks) GetPullRequestOk() (*PullRequestReviewCommentLinksPullRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PullRequest, true
}

// SetPullRequest sets field value
func (o *PullRequestReviewCommentLinks) SetPullRequest(v PullRequestReviewCommentLinksPullRequest) {
	o.PullRequest = v
}

func (o PullRequestReviewCommentLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["self"] = o.Self
	}
	if true {
		toSerialize["html"] = o.Html
	}
	if true {
		toSerialize["pull_request"] = o.PullRequest
	}
	return json.Marshal(toSerialize)
}

type NullablePullRequestReviewCommentLinks struct {
	value *PullRequestReviewCommentLinks
	isSet bool
}

func (v NullablePullRequestReviewCommentLinks) Get() *PullRequestReviewCommentLinks {
	return v.value
}

func (v *NullablePullRequestReviewCommentLinks) Set(val *PullRequestReviewCommentLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePullRequestReviewCommentLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePullRequestReviewCommentLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullRequestReviewCommentLinks(val *PullRequestReviewCommentLinks) *NullablePullRequestReviewCommentLinks {
	return &NullablePullRequestReviewCommentLinks{value: val, isSet: true}
}

func (v NullablePullRequestReviewCommentLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullRequestReviewCommentLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


