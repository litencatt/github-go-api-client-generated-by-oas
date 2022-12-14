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

// GitRef Git references within a repository
type GitRef struct {
	Ref string `json:"ref"`
	NodeId string `json:"node_id"`
	Url string `json:"url"`
	Object GitRefObject `json:"object"`
}

// NewGitRef instantiates a new GitRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitRef(ref string, nodeId string, url string, object GitRefObject) *GitRef {
	this := GitRef{}
	this.Ref = ref
	this.NodeId = nodeId
	this.Url = url
	this.Object = object
	return &this
}

// NewGitRefWithDefaults instantiates a new GitRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitRefWithDefaults() *GitRef {
	this := GitRef{}
	return &this
}

// GetRef returns the Ref field value
func (o *GitRef) GetRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ref
}

// GetRefOk returns a tuple with the Ref field value
// and a boolean to check if the value has been set.
func (o *GitRef) GetRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ref, true
}

// SetRef sets field value
func (o *GitRef) SetRef(v string) {
	o.Ref = v
}

// GetNodeId returns the NodeId field value
func (o *GitRef) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *GitRef) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *GitRef) SetNodeId(v string) {
	o.NodeId = v
}

// GetUrl returns the Url field value
func (o *GitRef) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *GitRef) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *GitRef) SetUrl(v string) {
	o.Url = v
}

// GetObject returns the Object field value
func (o *GitRef) GetObject() GitRefObject {
	if o == nil {
		var ret GitRefObject
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *GitRef) GetObjectOk() (*GitRefObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *GitRef) SetObject(v GitRefObject) {
	o.Object = v
}

func (o GitRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ref"] = o.Ref
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["object"] = o.Object
	}
	return json.Marshal(toSerialize)
}

type NullableGitRef struct {
	value *GitRef
	isSet bool
}

func (v NullableGitRef) Get() *GitRef {
	return v.value
}

func (v *NullableGitRef) Set(val *GitRef) {
	v.value = val
	v.isSet = true
}

func (v NullableGitRef) IsSet() bool {
	return v.isSet
}

func (v *NullableGitRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitRef(val *GitRef) *NullableGitRef {
	return &NullableGitRef{value: val, isSet: true}
}

func (v NullableGitRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


