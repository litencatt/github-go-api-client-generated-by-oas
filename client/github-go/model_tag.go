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

// Tag Tag
type Tag struct {
	Name string `json:"name"`
	Commit ShortBranchCommit `json:"commit"`
	ZipballUrl string `json:"zipball_url"`
	TarballUrl string `json:"tarball_url"`
	NodeId string `json:"node_id"`
}

// NewTag instantiates a new Tag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTag(name string, commit ShortBranchCommit, zipballUrl string, tarballUrl string, nodeId string) *Tag {
	this := Tag{}
	this.Name = name
	this.Commit = commit
	this.ZipballUrl = zipballUrl
	this.TarballUrl = tarballUrl
	this.NodeId = nodeId
	return &this
}

// NewTagWithDefaults instantiates a new Tag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagWithDefaults() *Tag {
	this := Tag{}
	return &this
}

// GetName returns the Name field value
func (o *Tag) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Tag) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Tag) SetName(v string) {
	o.Name = v
}

// GetCommit returns the Commit field value
func (o *Tag) GetCommit() ShortBranchCommit {
	if o == nil {
		var ret ShortBranchCommit
		return ret
	}

	return o.Commit
}

// GetCommitOk returns a tuple with the Commit field value
// and a boolean to check if the value has been set.
func (o *Tag) GetCommitOk() (*ShortBranchCommit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commit, true
}

// SetCommit sets field value
func (o *Tag) SetCommit(v ShortBranchCommit) {
	o.Commit = v
}

// GetZipballUrl returns the ZipballUrl field value
func (o *Tag) GetZipballUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZipballUrl
}

// GetZipballUrlOk returns a tuple with the ZipballUrl field value
// and a boolean to check if the value has been set.
func (o *Tag) GetZipballUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZipballUrl, true
}

// SetZipballUrl sets field value
func (o *Tag) SetZipballUrl(v string) {
	o.ZipballUrl = v
}

// GetTarballUrl returns the TarballUrl field value
func (o *Tag) GetTarballUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TarballUrl
}

// GetTarballUrlOk returns a tuple with the TarballUrl field value
// and a boolean to check if the value has been set.
func (o *Tag) GetTarballUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TarballUrl, true
}

// SetTarballUrl sets field value
func (o *Tag) SetTarballUrl(v string) {
	o.TarballUrl = v
}

// GetNodeId returns the NodeId field value
func (o *Tag) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *Tag) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *Tag) SetNodeId(v string) {
	o.NodeId = v
}

func (o Tag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["commit"] = o.Commit
	}
	if true {
		toSerialize["zipball_url"] = o.ZipballUrl
	}
	if true {
		toSerialize["tarball_url"] = o.TarballUrl
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	return json.Marshal(toSerialize)
}

type NullableTag struct {
	value *Tag
	isSet bool
}

func (v NullableTag) Get() *Tag {
	return v.value
}

func (v *NullableTag) Set(val *Tag) {
	v.value = val
	v.isSet = true
}

func (v NullableTag) IsSet() bool {
	return v.isSet
}

func (v *NullableTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTag(val *Tag) *NullableTag {
	return &NullableTag{value: val, isSet: true}
}

func (v NullableTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


