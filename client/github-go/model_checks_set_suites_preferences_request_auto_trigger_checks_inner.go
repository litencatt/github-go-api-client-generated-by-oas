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

// ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner struct for ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner
type ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner struct {
	// The `id` of the GitHub App.
	AppId int32 `json:"app_id"`
	// Set to `true` to enable automatic creation of CheckSuite events upon pushes to the repository, or `false` to disable them.
	Setting bool `json:"setting"`
}

// NewChecksSetSuitesPreferencesRequestAutoTriggerChecksInner instantiates a new ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChecksSetSuitesPreferencesRequestAutoTriggerChecksInner(appId int32, setting bool) *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner {
	this := ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner{}
	this.AppId = appId
	this.Setting = setting
	return &this
}

// NewChecksSetSuitesPreferencesRequestAutoTriggerChecksInnerWithDefaults instantiates a new ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChecksSetSuitesPreferencesRequestAutoTriggerChecksInnerWithDefaults() *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner {
	this := ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner{}
	var setting bool = true
	this.Setting = setting
	return &this
}

// GetAppId returns the AppId field value
func (o *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) GetAppId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) GetAppIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) SetAppId(v int32) {
	o.AppId = v
}

// GetSetting returns the Setting field value
func (o *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) GetSetting() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Setting
}

// GetSettingOk returns a tuple with the Setting field value
// and a boolean to check if the value has been set.
func (o *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) GetSettingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Setting, true
}

// SetSetting sets field value
func (o *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) SetSetting(v bool) {
	o.Setting = v
}

func (o ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["app_id"] = o.AppId
	}
	if true {
		toSerialize["setting"] = o.Setting
	}
	return json.Marshal(toSerialize)
}

type NullableChecksSetSuitesPreferencesRequestAutoTriggerChecksInner struct {
	value *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner
	isSet bool
}

func (v NullableChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) Get() *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner {
	return v.value
}

func (v *NullableChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) Set(val *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChecksSetSuitesPreferencesRequestAutoTriggerChecksInner(val *ChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) *NullableChecksSetSuitesPreferencesRequestAutoTriggerChecksInner {
	return &NullableChecksSetSuitesPreferencesRequestAutoTriggerChecksInner{value: val, isSet: true}
}

func (v NullableChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChecksSetSuitesPreferencesRequestAutoTriggerChecksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


