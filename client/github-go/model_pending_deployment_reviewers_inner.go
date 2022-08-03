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

// PendingDeploymentReviewersInner struct for PendingDeploymentReviewersInner
type PendingDeploymentReviewersInner struct {
	Type *DeploymentReviewerType `json:"type,omitempty"`
	Reviewer *PendingDeploymentReviewersInnerReviewer `json:"reviewer,omitempty"`
}

// NewPendingDeploymentReviewersInner instantiates a new PendingDeploymentReviewersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingDeploymentReviewersInner() *PendingDeploymentReviewersInner {
	this := PendingDeploymentReviewersInner{}
	return &this
}

// NewPendingDeploymentReviewersInnerWithDefaults instantiates a new PendingDeploymentReviewersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingDeploymentReviewersInnerWithDefaults() *PendingDeploymentReviewersInner {
	this := PendingDeploymentReviewersInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PendingDeploymentReviewersInner) GetType() DeploymentReviewerType {
	if o == nil || o.Type == nil {
		var ret DeploymentReviewerType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingDeploymentReviewersInner) GetTypeOk() (*DeploymentReviewerType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PendingDeploymentReviewersInner) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given DeploymentReviewerType and assigns it to the Type field.
func (o *PendingDeploymentReviewersInner) SetType(v DeploymentReviewerType) {
	o.Type = &v
}

// GetReviewer returns the Reviewer field value if set, zero value otherwise.
func (o *PendingDeploymentReviewersInner) GetReviewer() PendingDeploymentReviewersInnerReviewer {
	if o == nil || o.Reviewer == nil {
		var ret PendingDeploymentReviewersInnerReviewer
		return ret
	}
	return *o.Reviewer
}

// GetReviewerOk returns a tuple with the Reviewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PendingDeploymentReviewersInner) GetReviewerOk() (*PendingDeploymentReviewersInnerReviewer, bool) {
	if o == nil || o.Reviewer == nil {
		return nil, false
	}
	return o.Reviewer, true
}

// HasReviewer returns a boolean if a field has been set.
func (o *PendingDeploymentReviewersInner) HasReviewer() bool {
	if o != nil && o.Reviewer != nil {
		return true
	}

	return false
}

// SetReviewer gets a reference to the given PendingDeploymentReviewersInnerReviewer and assigns it to the Reviewer field.
func (o *PendingDeploymentReviewersInner) SetReviewer(v PendingDeploymentReviewersInnerReviewer) {
	o.Reviewer = &v
}

func (o PendingDeploymentReviewersInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Reviewer != nil {
		toSerialize["reviewer"] = o.Reviewer
	}
	return json.Marshal(toSerialize)
}

type NullablePendingDeploymentReviewersInner struct {
	value *PendingDeploymentReviewersInner
	isSet bool
}

func (v NullablePendingDeploymentReviewersInner) Get() *PendingDeploymentReviewersInner {
	return v.value
}

func (v *NullablePendingDeploymentReviewersInner) Set(val *PendingDeploymentReviewersInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingDeploymentReviewersInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingDeploymentReviewersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingDeploymentReviewersInner(val *PendingDeploymentReviewersInner) *NullablePendingDeploymentReviewersInner {
	return &NullablePendingDeploymentReviewersInner{value: val, isSet: true}
}

func (v NullablePendingDeploymentReviewersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingDeploymentReviewersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


