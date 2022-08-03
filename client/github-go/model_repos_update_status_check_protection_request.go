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

// ReposUpdateStatusCheckProtectionRequest struct for ReposUpdateStatusCheckProtectionRequest
type ReposUpdateStatusCheckProtectionRequest struct {
	// Require branches to be up to date before merging.
	Strict *bool `json:"strict,omitempty"`
	// **Deprecated**: The list of status checks to require in order to merge into this branch. If any of these checks have recently been set by a particular GitHub App, they will be required to come from that app in future for the branch to merge. Use `checks` instead of `contexts` for more fine-grained control. 
	// Deprecated
	Contexts []string `json:"contexts,omitempty"`
	// The list of status checks to require in order to merge into this branch.
	Checks []ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner `json:"checks,omitempty"`
}

// NewReposUpdateStatusCheckProtectionRequest instantiates a new ReposUpdateStatusCheckProtectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReposUpdateStatusCheckProtectionRequest() *ReposUpdateStatusCheckProtectionRequest {
	this := ReposUpdateStatusCheckProtectionRequest{}
	return &this
}

// NewReposUpdateStatusCheckProtectionRequestWithDefaults instantiates a new ReposUpdateStatusCheckProtectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReposUpdateStatusCheckProtectionRequestWithDefaults() *ReposUpdateStatusCheckProtectionRequest {
	this := ReposUpdateStatusCheckProtectionRequest{}
	return &this
}

// GetStrict returns the Strict field value if set, zero value otherwise.
func (o *ReposUpdateStatusCheckProtectionRequest) GetStrict() bool {
	if o == nil || o.Strict == nil {
		var ret bool
		return ret
	}
	return *o.Strict
}

// GetStrictOk returns a tuple with the Strict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposUpdateStatusCheckProtectionRequest) GetStrictOk() (*bool, bool) {
	if o == nil || o.Strict == nil {
		return nil, false
	}
	return o.Strict, true
}

// HasStrict returns a boolean if a field has been set.
func (o *ReposUpdateStatusCheckProtectionRequest) HasStrict() bool {
	if o != nil && o.Strict != nil {
		return true
	}

	return false
}

// SetStrict gets a reference to the given bool and assigns it to the Strict field.
func (o *ReposUpdateStatusCheckProtectionRequest) SetStrict(v bool) {
	o.Strict = &v
}

// GetContexts returns the Contexts field value if set, zero value otherwise.
// Deprecated
func (o *ReposUpdateStatusCheckProtectionRequest) GetContexts() []string {
	if o == nil || o.Contexts == nil {
		var ret []string
		return ret
	}
	return o.Contexts
}

// GetContextsOk returns a tuple with the Contexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ReposUpdateStatusCheckProtectionRequest) GetContextsOk() ([]string, bool) {
	if o == nil || o.Contexts == nil {
		return nil, false
	}
	return o.Contexts, true
}

// HasContexts returns a boolean if a field has been set.
func (o *ReposUpdateStatusCheckProtectionRequest) HasContexts() bool {
	if o != nil && o.Contexts != nil {
		return true
	}

	return false
}

// SetContexts gets a reference to the given []string and assigns it to the Contexts field.
// Deprecated
func (o *ReposUpdateStatusCheckProtectionRequest) SetContexts(v []string) {
	o.Contexts = v
}

// GetChecks returns the Checks field value if set, zero value otherwise.
func (o *ReposUpdateStatusCheckProtectionRequest) GetChecks() []ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner {
	if o == nil || o.Checks == nil {
		var ret []ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner
		return ret
	}
	return o.Checks
}

// GetChecksOk returns a tuple with the Checks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposUpdateStatusCheckProtectionRequest) GetChecksOk() ([]ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner, bool) {
	if o == nil || o.Checks == nil {
		return nil, false
	}
	return o.Checks, true
}

// HasChecks returns a boolean if a field has been set.
func (o *ReposUpdateStatusCheckProtectionRequest) HasChecks() bool {
	if o != nil && o.Checks != nil {
		return true
	}

	return false
}

// SetChecks gets a reference to the given []ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner and assigns it to the Checks field.
func (o *ReposUpdateStatusCheckProtectionRequest) SetChecks(v []ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner) {
	o.Checks = v
}

func (o ReposUpdateStatusCheckProtectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Strict != nil {
		toSerialize["strict"] = o.Strict
	}
	if o.Contexts != nil {
		toSerialize["contexts"] = o.Contexts
	}
	if o.Checks != nil {
		toSerialize["checks"] = o.Checks
	}
	return json.Marshal(toSerialize)
}

type NullableReposUpdateStatusCheckProtectionRequest struct {
	value *ReposUpdateStatusCheckProtectionRequest
	isSet bool
}

func (v NullableReposUpdateStatusCheckProtectionRequest) Get() *ReposUpdateStatusCheckProtectionRequest {
	return v.value
}

func (v *NullableReposUpdateStatusCheckProtectionRequest) Set(val *ReposUpdateStatusCheckProtectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReposUpdateStatusCheckProtectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReposUpdateStatusCheckProtectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReposUpdateStatusCheckProtectionRequest(val *ReposUpdateStatusCheckProtectionRequest) *NullableReposUpdateStatusCheckProtectionRequest {
	return &NullableReposUpdateStatusCheckProtectionRequest{value: val, isSet: true}
}

func (v NullableReposUpdateStatusCheckProtectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReposUpdateStatusCheckProtectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

