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

// CombinedBillingUsage struct for CombinedBillingUsage
type CombinedBillingUsage struct {
	// Numbers of days left in billing cycle.
	DaysLeftInBillingCycle int32 `json:"days_left_in_billing_cycle"`
	// Estimated storage space (GB) used in billing cycle.
	EstimatedPaidStorageForMonth int32 `json:"estimated_paid_storage_for_month"`
	// Estimated sum of free and paid storage space (GB) used in billing cycle.
	EstimatedStorageForMonth int32 `json:"estimated_storage_for_month"`
}

// NewCombinedBillingUsage instantiates a new CombinedBillingUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCombinedBillingUsage(daysLeftInBillingCycle int32, estimatedPaidStorageForMonth int32, estimatedStorageForMonth int32) *CombinedBillingUsage {
	this := CombinedBillingUsage{}
	this.DaysLeftInBillingCycle = daysLeftInBillingCycle
	this.EstimatedPaidStorageForMonth = estimatedPaidStorageForMonth
	this.EstimatedStorageForMonth = estimatedStorageForMonth
	return &this
}

// NewCombinedBillingUsageWithDefaults instantiates a new CombinedBillingUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCombinedBillingUsageWithDefaults() *CombinedBillingUsage {
	this := CombinedBillingUsage{}
	return &this
}

// GetDaysLeftInBillingCycle returns the DaysLeftInBillingCycle field value
func (o *CombinedBillingUsage) GetDaysLeftInBillingCycle() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DaysLeftInBillingCycle
}

// GetDaysLeftInBillingCycleOk returns a tuple with the DaysLeftInBillingCycle field value
// and a boolean to check if the value has been set.
func (o *CombinedBillingUsage) GetDaysLeftInBillingCycleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DaysLeftInBillingCycle, true
}

// SetDaysLeftInBillingCycle sets field value
func (o *CombinedBillingUsage) SetDaysLeftInBillingCycle(v int32) {
	o.DaysLeftInBillingCycle = v
}

// GetEstimatedPaidStorageForMonth returns the EstimatedPaidStorageForMonth field value
func (o *CombinedBillingUsage) GetEstimatedPaidStorageForMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EstimatedPaidStorageForMonth
}

// GetEstimatedPaidStorageForMonthOk returns a tuple with the EstimatedPaidStorageForMonth field value
// and a boolean to check if the value has been set.
func (o *CombinedBillingUsage) GetEstimatedPaidStorageForMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimatedPaidStorageForMonth, true
}

// SetEstimatedPaidStorageForMonth sets field value
func (o *CombinedBillingUsage) SetEstimatedPaidStorageForMonth(v int32) {
	o.EstimatedPaidStorageForMonth = v
}

// GetEstimatedStorageForMonth returns the EstimatedStorageForMonth field value
func (o *CombinedBillingUsage) GetEstimatedStorageForMonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EstimatedStorageForMonth
}

// GetEstimatedStorageForMonthOk returns a tuple with the EstimatedStorageForMonth field value
// and a boolean to check if the value has been set.
func (o *CombinedBillingUsage) GetEstimatedStorageForMonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EstimatedStorageForMonth, true
}

// SetEstimatedStorageForMonth sets field value
func (o *CombinedBillingUsage) SetEstimatedStorageForMonth(v int32) {
	o.EstimatedStorageForMonth = v
}

func (o CombinedBillingUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["days_left_in_billing_cycle"] = o.DaysLeftInBillingCycle
	}
	if true {
		toSerialize["estimated_paid_storage_for_month"] = o.EstimatedPaidStorageForMonth
	}
	if true {
		toSerialize["estimated_storage_for_month"] = o.EstimatedStorageForMonth
	}
	return json.Marshal(toSerialize)
}

type NullableCombinedBillingUsage struct {
	value *CombinedBillingUsage
	isSet bool
}

func (v NullableCombinedBillingUsage) Get() *CombinedBillingUsage {
	return v.value
}

func (v *NullableCombinedBillingUsage) Set(val *CombinedBillingUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableCombinedBillingUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableCombinedBillingUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCombinedBillingUsage(val *CombinedBillingUsage) *NullableCombinedBillingUsage {
	return &NullableCombinedBillingUsage{value: val, isSet: true}
}

func (v NullableCombinedBillingUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCombinedBillingUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

