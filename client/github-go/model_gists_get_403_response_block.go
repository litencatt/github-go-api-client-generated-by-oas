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

// GistsGet403ResponseBlock struct for GistsGet403ResponseBlock
type GistsGet403ResponseBlock struct {
	Reason *string `json:"reason,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	HtmlUrl NullableString `json:"html_url,omitempty"`
}

// NewGistsGet403ResponseBlock instantiates a new GistsGet403ResponseBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGistsGet403ResponseBlock() *GistsGet403ResponseBlock {
	this := GistsGet403ResponseBlock{}
	return &this
}

// NewGistsGet403ResponseBlockWithDefaults instantiates a new GistsGet403ResponseBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGistsGet403ResponseBlockWithDefaults() *GistsGet403ResponseBlock {
	this := GistsGet403ResponseBlock{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *GistsGet403ResponseBlock) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GistsGet403ResponseBlock) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *GistsGet403ResponseBlock) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *GistsGet403ResponseBlock) SetReason(v string) {
	o.Reason = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GistsGet403ResponseBlock) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GistsGet403ResponseBlock) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GistsGet403ResponseBlock) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GistsGet403ResponseBlock) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetHtmlUrl returns the HtmlUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GistsGet403ResponseBlock) GetHtmlUrl() string {
	if o == nil || o.HtmlUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.HtmlUrl.Get()
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GistsGet403ResponseBlock) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HtmlUrl.Get(), o.HtmlUrl.IsSet()
}

// HasHtmlUrl returns a boolean if a field has been set.
func (o *GistsGet403ResponseBlock) HasHtmlUrl() bool {
	if o != nil && o.HtmlUrl.IsSet() {
		return true
	}

	return false
}

// SetHtmlUrl gets a reference to the given NullableString and assigns it to the HtmlUrl field.
func (o *GistsGet403ResponseBlock) SetHtmlUrl(v string) {
	o.HtmlUrl.Set(&v)
}
// SetHtmlUrlNil sets the value for HtmlUrl to be an explicit nil
func (o *GistsGet403ResponseBlock) SetHtmlUrlNil() {
	o.HtmlUrl.Set(nil)
}

// UnsetHtmlUrl ensures that no value is present for HtmlUrl, not even an explicit nil
func (o *GistsGet403ResponseBlock) UnsetHtmlUrl() {
	o.HtmlUrl.Unset()
}

func (o GistsGet403ResponseBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.HtmlUrl.IsSet() {
		toSerialize["html_url"] = o.HtmlUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGistsGet403ResponseBlock struct {
	value *GistsGet403ResponseBlock
	isSet bool
}

func (v NullableGistsGet403ResponseBlock) Get() *GistsGet403ResponseBlock {
	return v.value
}

func (v *NullableGistsGet403ResponseBlock) Set(val *GistsGet403ResponseBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableGistsGet403ResponseBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableGistsGet403ResponseBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGistsGet403ResponseBlock(val *GistsGet403ResponseBlock) *NullableGistsGet403ResponseBlock {
	return &NullableGistsGet403ResponseBlock{value: val, isSet: true}
}

func (v NullableGistsGet403ResponseBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGistsGet403ResponseBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

