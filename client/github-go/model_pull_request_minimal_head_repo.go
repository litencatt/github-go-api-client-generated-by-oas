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

// PullRequestMinimalHeadRepo struct for PullRequestMinimalHeadRepo
type PullRequestMinimalHeadRepo struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	Name string `json:"name"`
}

// NewPullRequestMinimalHeadRepo instantiates a new PullRequestMinimalHeadRepo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPullRequestMinimalHeadRepo(id int32, url string, name string) *PullRequestMinimalHeadRepo {
	this := PullRequestMinimalHeadRepo{}
	this.Id = id
	this.Url = url
	this.Name = name
	return &this
}

// NewPullRequestMinimalHeadRepoWithDefaults instantiates a new PullRequestMinimalHeadRepo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPullRequestMinimalHeadRepoWithDefaults() *PullRequestMinimalHeadRepo {
	this := PullRequestMinimalHeadRepo{}
	return &this
}

// GetId returns the Id field value
func (o *PullRequestMinimalHeadRepo) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PullRequestMinimalHeadRepo) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PullRequestMinimalHeadRepo) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *PullRequestMinimalHeadRepo) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PullRequestMinimalHeadRepo) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PullRequestMinimalHeadRepo) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value
func (o *PullRequestMinimalHeadRepo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PullRequestMinimalHeadRepo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PullRequestMinimalHeadRepo) SetName(v string) {
	o.Name = v
}

func (o PullRequestMinimalHeadRepo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePullRequestMinimalHeadRepo struct {
	value *PullRequestMinimalHeadRepo
	isSet bool
}

func (v NullablePullRequestMinimalHeadRepo) Get() *PullRequestMinimalHeadRepo {
	return v.value
}

func (v *NullablePullRequestMinimalHeadRepo) Set(val *PullRequestMinimalHeadRepo) {
	v.value = val
	v.isSet = true
}

func (v NullablePullRequestMinimalHeadRepo) IsSet() bool {
	return v.isSet
}

func (v *NullablePullRequestMinimalHeadRepo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullRequestMinimalHeadRepo(val *PullRequestMinimalHeadRepo) *NullablePullRequestMinimalHeadRepo {
	return &NullablePullRequestMinimalHeadRepo{value: val, isSet: true}
}

func (v NullablePullRequestMinimalHeadRepo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullRequestMinimalHeadRepo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

