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

// WorkflowRunUsage Workflow Run Usage
type WorkflowRunUsage struct {
	Billable WorkflowRunUsageBillable `json:"billable"`
	RunDurationMs *int32 `json:"run_duration_ms,omitempty"`
}

// NewWorkflowRunUsage instantiates a new WorkflowRunUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRunUsage(billable WorkflowRunUsageBillable) *WorkflowRunUsage {
	this := WorkflowRunUsage{}
	this.Billable = billable
	return &this
}

// NewWorkflowRunUsageWithDefaults instantiates a new WorkflowRunUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRunUsageWithDefaults() *WorkflowRunUsage {
	this := WorkflowRunUsage{}
	return &this
}

// GetBillable returns the Billable field value
func (o *WorkflowRunUsage) GetBillable() WorkflowRunUsageBillable {
	if o == nil {
		var ret WorkflowRunUsageBillable
		return ret
	}

	return o.Billable
}

// GetBillableOk returns a tuple with the Billable field value
// and a boolean to check if the value has been set.
func (o *WorkflowRunUsage) GetBillableOk() (*WorkflowRunUsageBillable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Billable, true
}

// SetBillable sets field value
func (o *WorkflowRunUsage) SetBillable(v WorkflowRunUsageBillable) {
	o.Billable = v
}

// GetRunDurationMs returns the RunDurationMs field value if set, zero value otherwise.
func (o *WorkflowRunUsage) GetRunDurationMs() int32 {
	if o == nil || o.RunDurationMs == nil {
		var ret int32
		return ret
	}
	return *o.RunDurationMs
}

// GetRunDurationMsOk returns a tuple with the RunDurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunUsage) GetRunDurationMsOk() (*int32, bool) {
	if o == nil || o.RunDurationMs == nil {
		return nil, false
	}
	return o.RunDurationMs, true
}

// HasRunDurationMs returns a boolean if a field has been set.
func (o *WorkflowRunUsage) HasRunDurationMs() bool {
	if o != nil && o.RunDurationMs != nil {
		return true
	}

	return false
}

// SetRunDurationMs gets a reference to the given int32 and assigns it to the RunDurationMs field.
func (o *WorkflowRunUsage) SetRunDurationMs(v int32) {
	o.RunDurationMs = &v
}

func (o WorkflowRunUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["billable"] = o.Billable
	}
	if o.RunDurationMs != nil {
		toSerialize["run_duration_ms"] = o.RunDurationMs
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowRunUsage struct {
	value *WorkflowRunUsage
	isSet bool
}

func (v NullableWorkflowRunUsage) Get() *WorkflowRunUsage {
	return v.value
}

func (v *NullableWorkflowRunUsage) Set(val *WorkflowRunUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRunUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRunUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRunUsage(val *WorkflowRunUsage) *NullableWorkflowRunUsage {
	return &NullableWorkflowRunUsage{value: val, isSet: true}
}

func (v NullableWorkflowRunUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRunUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


