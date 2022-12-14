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

// AdvancedSecurityActiveCommittersUser struct for AdvancedSecurityActiveCommittersUser
type AdvancedSecurityActiveCommittersUser struct {
	UserLogin string `json:"user_login"`
	LastPushedDate string `json:"last_pushed_date"`
}

// NewAdvancedSecurityActiveCommittersUser instantiates a new AdvancedSecurityActiveCommittersUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvancedSecurityActiveCommittersUser(userLogin string, lastPushedDate string) *AdvancedSecurityActiveCommittersUser {
	this := AdvancedSecurityActiveCommittersUser{}
	this.UserLogin = userLogin
	this.LastPushedDate = lastPushedDate
	return &this
}

// NewAdvancedSecurityActiveCommittersUserWithDefaults instantiates a new AdvancedSecurityActiveCommittersUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvancedSecurityActiveCommittersUserWithDefaults() *AdvancedSecurityActiveCommittersUser {
	this := AdvancedSecurityActiveCommittersUser{}
	return &this
}

// GetUserLogin returns the UserLogin field value
func (o *AdvancedSecurityActiveCommittersUser) GetUserLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserLogin
}

// GetUserLoginOk returns a tuple with the UserLogin field value
// and a boolean to check if the value has been set.
func (o *AdvancedSecurityActiveCommittersUser) GetUserLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserLogin, true
}

// SetUserLogin sets field value
func (o *AdvancedSecurityActiveCommittersUser) SetUserLogin(v string) {
	o.UserLogin = v
}

// GetLastPushedDate returns the LastPushedDate field value
func (o *AdvancedSecurityActiveCommittersUser) GetLastPushedDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastPushedDate
}

// GetLastPushedDateOk returns a tuple with the LastPushedDate field value
// and a boolean to check if the value has been set.
func (o *AdvancedSecurityActiveCommittersUser) GetLastPushedDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastPushedDate, true
}

// SetLastPushedDate sets field value
func (o *AdvancedSecurityActiveCommittersUser) SetLastPushedDate(v string) {
	o.LastPushedDate = v
}

func (o AdvancedSecurityActiveCommittersUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["user_login"] = o.UserLogin
	}
	if true {
		toSerialize["last_pushed_date"] = o.LastPushedDate
	}
	return json.Marshal(toSerialize)
}

type NullableAdvancedSecurityActiveCommittersUser struct {
	value *AdvancedSecurityActiveCommittersUser
	isSet bool
}

func (v NullableAdvancedSecurityActiveCommittersUser) Get() *AdvancedSecurityActiveCommittersUser {
	return v.value
}

func (v *NullableAdvancedSecurityActiveCommittersUser) Set(val *AdvancedSecurityActiveCommittersUser) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvancedSecurityActiveCommittersUser) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvancedSecurityActiveCommittersUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvancedSecurityActiveCommittersUser(val *AdvancedSecurityActiveCommittersUser) *NullableAdvancedSecurityActiveCommittersUser {
	return &NullableAdvancedSecurityActiveCommittersUser{value: val, isSet: true}
}

func (v NullableAdvancedSecurityActiveCommittersUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvancedSecurityActiveCommittersUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


