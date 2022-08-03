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

// ActionsListJobsForWorkflowRunAttempt200Response struct for ActionsListJobsForWorkflowRunAttempt200Response
type ActionsListJobsForWorkflowRunAttempt200Response struct {
	TotalCount int32 `json:"total_count"`
	Jobs []Job `json:"jobs"`
}

// NewActionsListJobsForWorkflowRunAttempt200Response instantiates a new ActionsListJobsForWorkflowRunAttempt200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionsListJobsForWorkflowRunAttempt200Response(totalCount int32, jobs []Job) *ActionsListJobsForWorkflowRunAttempt200Response {
	this := ActionsListJobsForWorkflowRunAttempt200Response{}
	this.TotalCount = totalCount
	this.Jobs = jobs
	return &this
}

// NewActionsListJobsForWorkflowRunAttempt200ResponseWithDefaults instantiates a new ActionsListJobsForWorkflowRunAttempt200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsListJobsForWorkflowRunAttempt200ResponseWithDefaults() *ActionsListJobsForWorkflowRunAttempt200Response {
	this := ActionsListJobsForWorkflowRunAttempt200Response{}
	return &this
}

// GetTotalCount returns the TotalCount field value
func (o *ActionsListJobsForWorkflowRunAttempt200Response) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *ActionsListJobsForWorkflowRunAttempt200Response) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *ActionsListJobsForWorkflowRunAttempt200Response) SetTotalCount(v int32) {
	o.TotalCount = v
}

// GetJobs returns the Jobs field value
func (o *ActionsListJobsForWorkflowRunAttempt200Response) GetJobs() []Job {
	if o == nil {
		var ret []Job
		return ret
	}

	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value
// and a boolean to check if the value has been set.
func (o *ActionsListJobsForWorkflowRunAttempt200Response) GetJobsOk() ([]Job, bool) {
	if o == nil {
		return nil, false
	}
	return o.Jobs, true
}

// SetJobs sets field value
func (o *ActionsListJobsForWorkflowRunAttempt200Response) SetJobs(v []Job) {
	o.Jobs = v
}

func (o ActionsListJobsForWorkflowRunAttempt200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total_count"] = o.TotalCount
	}
	if true {
		toSerialize["jobs"] = o.Jobs
	}
	return json.Marshal(toSerialize)
}

type NullableActionsListJobsForWorkflowRunAttempt200Response struct {
	value *ActionsListJobsForWorkflowRunAttempt200Response
	isSet bool
}

func (v NullableActionsListJobsForWorkflowRunAttempt200Response) Get() *ActionsListJobsForWorkflowRunAttempt200Response {
	return v.value
}

func (v *NullableActionsListJobsForWorkflowRunAttempt200Response) Set(val *ActionsListJobsForWorkflowRunAttempt200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableActionsListJobsForWorkflowRunAttempt200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableActionsListJobsForWorkflowRunAttempt200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionsListJobsForWorkflowRunAttempt200Response(val *ActionsListJobsForWorkflowRunAttempt200Response) *NullableActionsListJobsForWorkflowRunAttempt200Response {
	return &NullableActionsListJobsForWorkflowRunAttempt200Response{value: val, isSet: true}
}

func (v NullableActionsListJobsForWorkflowRunAttempt200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionsListJobsForWorkflowRunAttempt200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


