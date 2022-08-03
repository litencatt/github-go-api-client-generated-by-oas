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

// ActionsBillingUsageMinutesUsedBreakdown struct for ActionsBillingUsageMinutesUsedBreakdown
type ActionsBillingUsageMinutesUsedBreakdown struct {
	// Total minutes used on Ubuntu runner machines.
	UBUNTU *int32 `json:"UBUNTU,omitempty"`
	// Total minutes used on macOS runner machines.
	MACOS *int32 `json:"MACOS,omitempty"`
	// Total minutes used on Windows runner machines.
	WINDOWS *int32 `json:"WINDOWS,omitempty"`
	// Total minutes used on Ubuntu 4 core runner machines.
	Ubuntu4Core *int32 `json:"ubuntu_4_core,omitempty"`
	// Total minutes used on Ubuntu 8 core runner machines.
	Ubuntu8Core *int32 `json:"ubuntu_8_core,omitempty"`
	// Total minutes used on Ubuntu 16 core runner machines.
	Ubuntu16Core *int32 `json:"ubuntu_16_core,omitempty"`
	// Total minutes used on Ubuntu 32 core runner machines.
	Ubuntu32Core *int32 `json:"ubuntu_32_core,omitempty"`
	// Total minutes used on Ubuntu 64 core runner machines.
	Ubuntu64Core *int32 `json:"ubuntu_64_core,omitempty"`
	// Total minutes used on Windows 4 core runner machines.
	Windows4Core *int32 `json:"windows_4_core,omitempty"`
	// Total minutes used on Windows 8 core runner machines.
	Windows8Core *int32 `json:"windows_8_core,omitempty"`
	// Total minutes used on Windows 16 core runner machines.
	Windows16Core *int32 `json:"windows_16_core,omitempty"`
	// Total minutes used on Windows 32 core runner machines.
	Windows32Core *int32 `json:"windows_32_core,omitempty"`
	// Total minutes used on Windows 64 core runner machines.
	Windows64Core *int32 `json:"windows_64_core,omitempty"`
	// Total minutes used on all runner machines.
	Total *int32 `json:"total,omitempty"`
}

// NewActionsBillingUsageMinutesUsedBreakdown instantiates a new ActionsBillingUsageMinutesUsedBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionsBillingUsageMinutesUsedBreakdown() *ActionsBillingUsageMinutesUsedBreakdown {
	this := ActionsBillingUsageMinutesUsedBreakdown{}
	return &this
}

// NewActionsBillingUsageMinutesUsedBreakdownWithDefaults instantiates a new ActionsBillingUsageMinutesUsedBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsBillingUsageMinutesUsedBreakdownWithDefaults() *ActionsBillingUsageMinutesUsedBreakdown {
	this := ActionsBillingUsageMinutesUsedBreakdown{}
	return &this
}

// GetUBUNTU returns the UBUNTU field value if set, zero value otherwise.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetUBUNTU() int32 {
	if o == nil || o.UBUNTU == nil {
		var ret int32
		return ret
	}
	return *o.UBUNTU
}

// GetUBUNTUOk returns a tuple with the UBUNTU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetUBUNTUOk() (*int32, bool) {
	if o == nil || o.UBUNTU == nil {
		return nil, false
	}
	return o.UBUNTU, true
}

// HasUBUNTU returns a boolean if a field has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) HasUBUNTU() bool {
	if o != nil && o.UBUNTU != nil {
		return true
	}

	return false
}

// SetUBUNTU gets a reference to the given int32 and assigns it to the UBUNTU field.
func (o *ActionsBillingUsageMinutesUsedBreakdown) SetUBUNTU(v int32) {
	o.UBUNTU = &v
}

// GetMACOS returns the MACOS field value if set, zero value otherwise.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetMACOS() int32 {
	if o == nil || o.MACOS == nil {
		var ret int32
		return ret
	}
	return *o.MACOS
}

// GetMACOSOk returns a tuple with the MACOS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetMACOSOk() (*int32, bool) {
	if o == nil || o.MACOS == nil {
		return nil, false
	}
	return o.MACOS, true
}

// HasMACOS returns a boolean if a field has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) HasMACOS() bool {
	if o != nil && o.MACOS != nil {
		return true
	}

	return false
}

// SetMACOS gets a reference to the given int32 and assigns it to the MACOS field.
func (o *ActionsBillingUsageMinutesUsedBreakdown) SetMACOS(v int32) {
	o.MACOS = &v
}

// GetWINDOWS returns the WINDOWS field value if set, zero value otherwise.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetWINDOWS() int32 {
	if o == nil || o.WINDOWS == nil {
		var ret int32
		return ret
	}
	return *o.WINDOWS
}

// GetWINDOWSOk returns a tuple with the WINDOWS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetWINDOWSOk() (*int32, bool) {
	if o == nil || o.WINDOWS == nil {
		return nil, false
	}
	return o.WINDOWS, true
}

// HasWINDOWS returns a boolean if a field has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) HasWINDOWS() bool {
	if o != nil && o.WINDOWS != nil {
		return true
	}

	return false
}

// SetWINDOWS gets a reference to the given int32 and assigns it to the WINDOWS field.
func (o *ActionsBillingUsageMinutesUsedBreakdown) SetWINDOWS(v int32) {
	o.WINDOWS = &v
}

// GetUbuntu4Core returns the Ubuntu4Core field value if set, zero value otherwise.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetUbuntu4Core() int32 {
	if o == nil || o.Ubuntu4Core == nil {
		var ret int32
		return ret
	}
	return *o.Ubuntu4Core
}

// GetUbuntu4CoreOk returns a tuple with the Ubuntu4Core field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetUbuntu4CoreOk() (*int32, bool) {
	if o == nil || o.Ubuntu4Core == nil {
		return nil, false
	}
	return o.Ubuntu4Core, true
}

// HasUbuntu4Core returns a boolean if a field has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) HasUbuntu4Core() bool {
	if o != nil && o.Ubuntu4Core != nil {
		return true
	}

	return false
}

// SetUbuntu4Core gets a reference to the given int32 and assigns it to the Ubuntu4Core field.
func (o *ActionsBillingUsageMinutesUsedBreakdown) SetUbuntu4Core(v int32) {
	o.Ubuntu4Core = &v
}

// GetUbuntu8Core returns the Ubuntu8Core field value if set, zero value otherwise.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetUbuntu8Core() int32 {
	if o == nil || o.Ubuntu8Core == nil {
		var ret int32
		return ret
	}
	return *o.Ubuntu8Core
}

// GetUbuntu8CoreOk returns a tuple with the Ubuntu8Core field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetUbuntu8CoreOk() (*int32, bool) {
	if o == nil || o.Ubuntu8Core == nil {
		return nil, false
	}
	return o.Ubuntu8Core, true
}

// HasUbuntu8Core returns a boolean if a field has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) HasUbuntu8Core() bool {
	if o != nil && o.Ubuntu8Core != nil {
		return true
	}

	return false
}

// SetUbuntu8Core gets a reference to the given int32 and assigns it to the Ubuntu8Core field.
func (o *ActionsBillingUsageMinutesUsedBreakdown) SetUbuntu8Core(v int32) {
	o.Ubuntu8Core = &v
}

// GetUbuntu16Core returns the Ubuntu16Core field value if set, zero value otherwise.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetUbuntu16Core() int32 {
	if o == nil || o.Ubuntu16Core == nil {
		var ret int32
		return ret
	}
	return *o.Ubuntu16Core
}

// GetUbuntu16CoreOk returns a tuple with the Ubuntu16Core field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetUbuntu16CoreOk() (*int32, bool) {
	if o == nil || o.Ubuntu16Core == nil {
		return nil, false
	}
	return o.Ubuntu16Core, true
}

// HasUbuntu16Core returns a boolean if a field has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) HasUbuntu16Core() bool {
	if o != nil && o.Ubuntu16Core != nil {
		return true
	}

	return false
}

// SetUbuntu16Core gets a reference to the given int32 and assigns it to the Ubuntu16Core field.
func (o *ActionsBillingUsageMinutesUsedBreakdown) SetUbuntu16Core(v int32) {
	o.Ubuntu16Core = &v
}

// GetUbuntu32Core returns the Ubuntu32Core field value if set, zero value otherwise.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetUbuntu32Core() int32 {
	if o == nil || o.Ubuntu32Core == nil {
		var ret int32
		return ret
	}
	return *o.Ubuntu32Core
}

// GetUbuntu32CoreOk returns a tuple with the Ubuntu32Core field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetUbuntu32CoreOk() (*int32, bool) {
	if o == nil || o.Ubuntu32Core == nil {
		return nil, false
	}
	return o.Ubuntu32Core, true
}

// HasUbuntu32Core returns a boolean if a field has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) HasUbuntu32Core() bool {
	if o != nil && o.Ubuntu32Core != nil {
		return true
	}

	return false
}

// SetUbuntu32Core gets a reference to the given int32 and assigns it to the Ubuntu32Core field.
func (o *ActionsBillingUsageMinutesUsedBreakdown) SetUbuntu32Core(v int32) {
	o.Ubuntu32Core = &v
}

// GetUbuntu64Core returns the Ubuntu64Core field value if set, zero value otherwise.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetUbuntu64Core() int32 {
	if o == nil || o.Ubuntu64Core == nil {
		var ret int32
		return ret
	}
	return *o.Ubuntu64Core
}

// GetUbuntu64CoreOk returns a tuple with the Ubuntu64Core field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetUbuntu64CoreOk() (*int32, bool) {
	if o == nil || o.Ubuntu64Core == nil {
		return nil, false
	}
	return o.Ubuntu64Core, true
}

// HasUbuntu64Core returns a boolean if a field has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) HasUbuntu64Core() bool {
	if o != nil && o.Ubuntu64Core != nil {
		return true
	}

	return false
}

// SetUbuntu64Core gets a reference to the given int32 and assigns it to the Ubuntu64Core field.
func (o *ActionsBillingUsageMinutesUsedBreakdown) SetUbuntu64Core(v int32) {
	o.Ubuntu64Core = &v
}

// GetWindows4Core returns the Windows4Core field value if set, zero value otherwise.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetWindows4Core() int32 {
	if o == nil || o.Windows4Core == nil {
		var ret int32
		return ret
	}
	return *o.Windows4Core
}

// GetWindows4CoreOk returns a tuple with the Windows4Core field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetWindows4CoreOk() (*int32, bool) {
	if o == nil || o.Windows4Core == nil {
		return nil, false
	}
	return o.Windows4Core, true
}

// HasWindows4Core returns a boolean if a field has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) HasWindows4Core() bool {
	if o != nil && o.Windows4Core != nil {
		return true
	}

	return false
}

// SetWindows4Core gets a reference to the given int32 and assigns it to the Windows4Core field.
func (o *ActionsBillingUsageMinutesUsedBreakdown) SetWindows4Core(v int32) {
	o.Windows4Core = &v
}

// GetWindows8Core returns the Windows8Core field value if set, zero value otherwise.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetWindows8Core() int32 {
	if o == nil || o.Windows8Core == nil {
		var ret int32
		return ret
	}
	return *o.Windows8Core
}

// GetWindows8CoreOk returns a tuple with the Windows8Core field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetWindows8CoreOk() (*int32, bool) {
	if o == nil || o.Windows8Core == nil {
		return nil, false
	}
	return o.Windows8Core, true
}

// HasWindows8Core returns a boolean if a field has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) HasWindows8Core() bool {
	if o != nil && o.Windows8Core != nil {
		return true
	}

	return false
}

// SetWindows8Core gets a reference to the given int32 and assigns it to the Windows8Core field.
func (o *ActionsBillingUsageMinutesUsedBreakdown) SetWindows8Core(v int32) {
	o.Windows8Core = &v
}

// GetWindows16Core returns the Windows16Core field value if set, zero value otherwise.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetWindows16Core() int32 {
	if o == nil || o.Windows16Core == nil {
		var ret int32
		return ret
	}
	return *o.Windows16Core
}

// GetWindows16CoreOk returns a tuple with the Windows16Core field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetWindows16CoreOk() (*int32, bool) {
	if o == nil || o.Windows16Core == nil {
		return nil, false
	}
	return o.Windows16Core, true
}

// HasWindows16Core returns a boolean if a field has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) HasWindows16Core() bool {
	if o != nil && o.Windows16Core != nil {
		return true
	}

	return false
}

// SetWindows16Core gets a reference to the given int32 and assigns it to the Windows16Core field.
func (o *ActionsBillingUsageMinutesUsedBreakdown) SetWindows16Core(v int32) {
	o.Windows16Core = &v
}

// GetWindows32Core returns the Windows32Core field value if set, zero value otherwise.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetWindows32Core() int32 {
	if o == nil || o.Windows32Core == nil {
		var ret int32
		return ret
	}
	return *o.Windows32Core
}

// GetWindows32CoreOk returns a tuple with the Windows32Core field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetWindows32CoreOk() (*int32, bool) {
	if o == nil || o.Windows32Core == nil {
		return nil, false
	}
	return o.Windows32Core, true
}

// HasWindows32Core returns a boolean if a field has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) HasWindows32Core() bool {
	if o != nil && o.Windows32Core != nil {
		return true
	}

	return false
}

// SetWindows32Core gets a reference to the given int32 and assigns it to the Windows32Core field.
func (o *ActionsBillingUsageMinutesUsedBreakdown) SetWindows32Core(v int32) {
	o.Windows32Core = &v
}

// GetWindows64Core returns the Windows64Core field value if set, zero value otherwise.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetWindows64Core() int32 {
	if o == nil || o.Windows64Core == nil {
		var ret int32
		return ret
	}
	return *o.Windows64Core
}

// GetWindows64CoreOk returns a tuple with the Windows64Core field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetWindows64CoreOk() (*int32, bool) {
	if o == nil || o.Windows64Core == nil {
		return nil, false
	}
	return o.Windows64Core, true
}

// HasWindows64Core returns a boolean if a field has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) HasWindows64Core() bool {
	if o != nil && o.Windows64Core != nil {
		return true
	}

	return false
}

// SetWindows64Core gets a reference to the given int32 and assigns it to the Windows64Core field.
func (o *ActionsBillingUsageMinutesUsedBreakdown) SetWindows64Core(v int32) {
	o.Windows64Core = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ActionsBillingUsageMinutesUsedBreakdown) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ActionsBillingUsageMinutesUsedBreakdown) SetTotal(v int32) {
	o.Total = &v
}

func (o ActionsBillingUsageMinutesUsedBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UBUNTU != nil {
		toSerialize["UBUNTU"] = o.UBUNTU
	}
	if o.MACOS != nil {
		toSerialize["MACOS"] = o.MACOS
	}
	if o.WINDOWS != nil {
		toSerialize["WINDOWS"] = o.WINDOWS
	}
	if o.Ubuntu4Core != nil {
		toSerialize["ubuntu_4_core"] = o.Ubuntu4Core
	}
	if o.Ubuntu8Core != nil {
		toSerialize["ubuntu_8_core"] = o.Ubuntu8Core
	}
	if o.Ubuntu16Core != nil {
		toSerialize["ubuntu_16_core"] = o.Ubuntu16Core
	}
	if o.Ubuntu32Core != nil {
		toSerialize["ubuntu_32_core"] = o.Ubuntu32Core
	}
	if o.Ubuntu64Core != nil {
		toSerialize["ubuntu_64_core"] = o.Ubuntu64Core
	}
	if o.Windows4Core != nil {
		toSerialize["windows_4_core"] = o.Windows4Core
	}
	if o.Windows8Core != nil {
		toSerialize["windows_8_core"] = o.Windows8Core
	}
	if o.Windows16Core != nil {
		toSerialize["windows_16_core"] = o.Windows16Core
	}
	if o.Windows32Core != nil {
		toSerialize["windows_32_core"] = o.Windows32Core
	}
	if o.Windows64Core != nil {
		toSerialize["windows_64_core"] = o.Windows64Core
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableActionsBillingUsageMinutesUsedBreakdown struct {
	value *ActionsBillingUsageMinutesUsedBreakdown
	isSet bool
}

func (v NullableActionsBillingUsageMinutesUsedBreakdown) Get() *ActionsBillingUsageMinutesUsedBreakdown {
	return v.value
}

func (v *NullableActionsBillingUsageMinutesUsedBreakdown) Set(val *ActionsBillingUsageMinutesUsedBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableActionsBillingUsageMinutesUsedBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableActionsBillingUsageMinutesUsedBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionsBillingUsageMinutesUsedBreakdown(val *ActionsBillingUsageMinutesUsedBreakdown) *NullableActionsBillingUsageMinutesUsedBreakdown {
	return &NullableActionsBillingUsageMinutesUsedBreakdown{value: val, isSet: true}
}

func (v NullableActionsBillingUsageMinutesUsedBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionsBillingUsageMinutesUsedBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


