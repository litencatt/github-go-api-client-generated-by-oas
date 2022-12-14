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

// DeployKey An SSH key granting access to a single repository.
type DeployKey struct {
	Id int32 `json:"id"`
	Key string `json:"key"`
	Url string `json:"url"`
	Title string `json:"title"`
	Verified bool `json:"verified"`
	CreatedAt string `json:"created_at"`
	ReadOnly bool `json:"read_only"`
}

// NewDeployKey instantiates a new DeployKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployKey(id int32, key string, url string, title string, verified bool, createdAt string, readOnly bool) *DeployKey {
	this := DeployKey{}
	this.Id = id
	this.Key = key
	this.Url = url
	this.Title = title
	this.Verified = verified
	this.CreatedAt = createdAt
	this.ReadOnly = readOnly
	return &this
}

// NewDeployKeyWithDefaults instantiates a new DeployKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeployKeyWithDefaults() *DeployKey {
	this := DeployKey{}
	return &this
}

// GetId returns the Id field value
func (o *DeployKey) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeployKey) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeployKey) SetId(v int32) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *DeployKey) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *DeployKey) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *DeployKey) SetKey(v string) {
	o.Key = v
}

// GetUrl returns the Url field value
func (o *DeployKey) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DeployKey) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DeployKey) SetUrl(v string) {
	o.Url = v
}

// GetTitle returns the Title field value
func (o *DeployKey) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DeployKey) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *DeployKey) SetTitle(v string) {
	o.Title = v
}

// GetVerified returns the Verified field value
func (o *DeployKey) GetVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value
// and a boolean to check if the value has been set.
func (o *DeployKey) GetVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verified, true
}

// SetVerified sets field value
func (o *DeployKey) SetVerified(v bool) {
	o.Verified = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DeployKey) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DeployKey) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DeployKey) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetReadOnly returns the ReadOnly field value
func (o *DeployKey) GetReadOnly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value
// and a boolean to check if the value has been set.
func (o *DeployKey) GetReadOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadOnly, true
}

// SetReadOnly sets field value
func (o *DeployKey) SetReadOnly(v bool) {
	o.ReadOnly = v
}

func (o DeployKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["verified"] = o.Verified
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["read_only"] = o.ReadOnly
	}
	return json.Marshal(toSerialize)
}

type NullableDeployKey struct {
	value *DeployKey
	isSet bool
}

func (v NullableDeployKey) Get() *DeployKey {
	return v.value
}

func (v *NullableDeployKey) Set(val *DeployKey) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployKey) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployKey(val *DeployKey) *NullableDeployKey {
	return &NullableDeployKey{value: val, isSet: true}
}

func (v NullableDeployKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


