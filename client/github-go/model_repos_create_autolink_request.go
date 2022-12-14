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

// ReposCreateAutolinkRequest struct for ReposCreateAutolinkRequest
type ReposCreateAutolinkRequest struct {
	// The prefix appended by alphanumeric characters will generate a link any time it is found in an issue, pull request, or commit.
	KeyPrefix string `json:"key_prefix"`
	// The URL must contain `<num>` for the reference number. `<num>` matches alphanumeric characters `A-Z` (case insensitive), `0-9`, and `-`.
	UrlTemplate string `json:"url_template"`
}

// NewReposCreateAutolinkRequest instantiates a new ReposCreateAutolinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReposCreateAutolinkRequest(keyPrefix string, urlTemplate string) *ReposCreateAutolinkRequest {
	this := ReposCreateAutolinkRequest{}
	this.KeyPrefix = keyPrefix
	this.UrlTemplate = urlTemplate
	return &this
}

// NewReposCreateAutolinkRequestWithDefaults instantiates a new ReposCreateAutolinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReposCreateAutolinkRequestWithDefaults() *ReposCreateAutolinkRequest {
	this := ReposCreateAutolinkRequest{}
	return &this
}

// GetKeyPrefix returns the KeyPrefix field value
func (o *ReposCreateAutolinkRequest) GetKeyPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyPrefix
}

// GetKeyPrefixOk returns a tuple with the KeyPrefix field value
// and a boolean to check if the value has been set.
func (o *ReposCreateAutolinkRequest) GetKeyPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyPrefix, true
}

// SetKeyPrefix sets field value
func (o *ReposCreateAutolinkRequest) SetKeyPrefix(v string) {
	o.KeyPrefix = v
}

// GetUrlTemplate returns the UrlTemplate field value
func (o *ReposCreateAutolinkRequest) GetUrlTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UrlTemplate
}

// GetUrlTemplateOk returns a tuple with the UrlTemplate field value
// and a boolean to check if the value has been set.
func (o *ReposCreateAutolinkRequest) GetUrlTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UrlTemplate, true
}

// SetUrlTemplate sets field value
func (o *ReposCreateAutolinkRequest) SetUrlTemplate(v string) {
	o.UrlTemplate = v
}

func (o ReposCreateAutolinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key_prefix"] = o.KeyPrefix
	}
	if true {
		toSerialize["url_template"] = o.UrlTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableReposCreateAutolinkRequest struct {
	value *ReposCreateAutolinkRequest
	isSet bool
}

func (v NullableReposCreateAutolinkRequest) Get() *ReposCreateAutolinkRequest {
	return v.value
}

func (v *NullableReposCreateAutolinkRequest) Set(val *ReposCreateAutolinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReposCreateAutolinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReposCreateAutolinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReposCreateAutolinkRequest(val *ReposCreateAutolinkRequest) *NullableReposCreateAutolinkRequest {
	return &NullableReposCreateAutolinkRequest{value: val, isSet: true}
}

func (v NullableReposCreateAutolinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReposCreateAutolinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


