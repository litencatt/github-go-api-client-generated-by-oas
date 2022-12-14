/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"encoding/json"
	"time"
)

// GistSimpleForksInner struct for GistSimpleForksInner
type GistSimpleForksInner struct {
	Id *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	User *PublicUser `json:"user,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewGistSimpleForksInner instantiates a new GistSimpleForksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGistSimpleForksInner() *GistSimpleForksInner {
	this := GistSimpleForksInner{}
	return &this
}

// NewGistSimpleForksInnerWithDefaults instantiates a new GistSimpleForksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGistSimpleForksInnerWithDefaults() *GistSimpleForksInner {
	this := GistSimpleForksInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GistSimpleForksInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GistSimpleForksInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GistSimpleForksInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GistSimpleForksInner) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GistSimpleForksInner) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GistSimpleForksInner) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GistSimpleForksInner) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GistSimpleForksInner) SetUrl(v string) {
	o.Url = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *GistSimpleForksInner) GetUser() PublicUser {
	if o == nil || o.User == nil {
		var ret PublicUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GistSimpleForksInner) GetUserOk() (*PublicUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *GistSimpleForksInner) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given PublicUser and assigns it to the User field.
func (o *GistSimpleForksInner) SetUser(v PublicUser) {
	o.User = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GistSimpleForksInner) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GistSimpleForksInner) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GistSimpleForksInner) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GistSimpleForksInner) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GistSimpleForksInner) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GistSimpleForksInner) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GistSimpleForksInner) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *GistSimpleForksInner) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o GistSimpleForksInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableGistSimpleForksInner struct {
	value *GistSimpleForksInner
	isSet bool
}

func (v NullableGistSimpleForksInner) Get() *GistSimpleForksInner {
	return v.value
}

func (v *NullableGistSimpleForksInner) Set(val *GistSimpleForksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGistSimpleForksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGistSimpleForksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGistSimpleForksInner(val *GistSimpleForksInner) *NullableGistSimpleForksInner {
	return &NullableGistSimpleForksInner{value: val, isSet: true}
}

func (v NullableGistSimpleForksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGistSimpleForksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


