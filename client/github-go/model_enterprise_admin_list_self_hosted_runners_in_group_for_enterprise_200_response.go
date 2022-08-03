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

// EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response struct for EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response
type EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response struct {
	TotalCount float32 `json:"total_count"`
	Runners []Runner `json:"runners"`
}

// NewEnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response instantiates a new EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response(totalCount float32, runners []Runner) *EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response {
	this := EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response{}
	this.TotalCount = totalCount
	this.Runners = runners
	return &this
}

// NewEnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200ResponseWithDefaults instantiates a new EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200ResponseWithDefaults() *EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response {
	this := EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response{}
	return &this
}

// GetTotalCount returns the TotalCount field value
func (o *EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response) GetTotalCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response) GetTotalCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response) SetTotalCount(v float32) {
	o.TotalCount = v
}

// GetRunners returns the Runners field value
func (o *EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response) GetRunners() []Runner {
	if o == nil {
		var ret []Runner
		return ret
	}

	return o.Runners
}

// GetRunnersOk returns a tuple with the Runners field value
// and a boolean to check if the value has been set.
func (o *EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response) GetRunnersOk() ([]Runner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Runners, true
}

// SetRunners sets field value
func (o *EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response) SetRunners(v []Runner) {
	o.Runners = v
}

func (o EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total_count"] = o.TotalCount
	}
	if true {
		toSerialize["runners"] = o.Runners
	}
	return json.Marshal(toSerialize)
}

type NullableEnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response struct {
	value *EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response
	isSet bool
}

func (v NullableEnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response) Get() *EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response {
	return v.value
}

func (v *NullableEnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response) Set(val *EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response(val *EnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response) *NullableEnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response {
	return &NullableEnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response{value: val, isSet: true}
}

func (v NullableEnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseAdminListSelfHostedRunnersInGroupForEnterprise200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

