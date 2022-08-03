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

// CodespaceGitStatus Details about the codespace's git repository.
type CodespaceGitStatus struct {
	// The number of commits the local repository is ahead of the remote.
	Ahead *int32 `json:"ahead,omitempty"`
	// The number of commits the local repository is behind the remote.
	Behind *int32 `json:"behind,omitempty"`
	// Whether the local repository has unpushed changes.
	HasUnpushedChanges *bool `json:"has_unpushed_changes,omitempty"`
	// Whether the local repository has uncommitted changes.
	HasUncommittedChanges *bool `json:"has_uncommitted_changes,omitempty"`
	// The current branch (or SHA if in detached HEAD state) of the local repository.
	Ref *string `json:"ref,omitempty"`
}

// NewCodespaceGitStatus instantiates a new CodespaceGitStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodespaceGitStatus() *CodespaceGitStatus {
	this := CodespaceGitStatus{}
	return &this
}

// NewCodespaceGitStatusWithDefaults instantiates a new CodespaceGitStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodespaceGitStatusWithDefaults() *CodespaceGitStatus {
	this := CodespaceGitStatus{}
	return &this
}

// GetAhead returns the Ahead field value if set, zero value otherwise.
func (o *CodespaceGitStatus) GetAhead() int32 {
	if o == nil || o.Ahead == nil {
		var ret int32
		return ret
	}
	return *o.Ahead
}

// GetAheadOk returns a tuple with the Ahead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodespaceGitStatus) GetAheadOk() (*int32, bool) {
	if o == nil || o.Ahead == nil {
		return nil, false
	}
	return o.Ahead, true
}

// HasAhead returns a boolean if a field has been set.
func (o *CodespaceGitStatus) HasAhead() bool {
	if o != nil && o.Ahead != nil {
		return true
	}

	return false
}

// SetAhead gets a reference to the given int32 and assigns it to the Ahead field.
func (o *CodespaceGitStatus) SetAhead(v int32) {
	o.Ahead = &v
}

// GetBehind returns the Behind field value if set, zero value otherwise.
func (o *CodespaceGitStatus) GetBehind() int32 {
	if o == nil || o.Behind == nil {
		var ret int32
		return ret
	}
	return *o.Behind
}

// GetBehindOk returns a tuple with the Behind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodespaceGitStatus) GetBehindOk() (*int32, bool) {
	if o == nil || o.Behind == nil {
		return nil, false
	}
	return o.Behind, true
}

// HasBehind returns a boolean if a field has been set.
func (o *CodespaceGitStatus) HasBehind() bool {
	if o != nil && o.Behind != nil {
		return true
	}

	return false
}

// SetBehind gets a reference to the given int32 and assigns it to the Behind field.
func (o *CodespaceGitStatus) SetBehind(v int32) {
	o.Behind = &v
}

// GetHasUnpushedChanges returns the HasUnpushedChanges field value if set, zero value otherwise.
func (o *CodespaceGitStatus) GetHasUnpushedChanges() bool {
	if o == nil || o.HasUnpushedChanges == nil {
		var ret bool
		return ret
	}
	return *o.HasUnpushedChanges
}

// GetHasUnpushedChangesOk returns a tuple with the HasUnpushedChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodespaceGitStatus) GetHasUnpushedChangesOk() (*bool, bool) {
	if o == nil || o.HasUnpushedChanges == nil {
		return nil, false
	}
	return o.HasUnpushedChanges, true
}

// HasHasUnpushedChanges returns a boolean if a field has been set.
func (o *CodespaceGitStatus) HasHasUnpushedChanges() bool {
	if o != nil && o.HasUnpushedChanges != nil {
		return true
	}

	return false
}

// SetHasUnpushedChanges gets a reference to the given bool and assigns it to the HasUnpushedChanges field.
func (o *CodespaceGitStatus) SetHasUnpushedChanges(v bool) {
	o.HasUnpushedChanges = &v
}

// GetHasUncommittedChanges returns the HasUncommittedChanges field value if set, zero value otherwise.
func (o *CodespaceGitStatus) GetHasUncommittedChanges() bool {
	if o == nil || o.HasUncommittedChanges == nil {
		var ret bool
		return ret
	}
	return *o.HasUncommittedChanges
}

// GetHasUncommittedChangesOk returns a tuple with the HasUncommittedChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodespaceGitStatus) GetHasUncommittedChangesOk() (*bool, bool) {
	if o == nil || o.HasUncommittedChanges == nil {
		return nil, false
	}
	return o.HasUncommittedChanges, true
}

// HasHasUncommittedChanges returns a boolean if a field has been set.
func (o *CodespaceGitStatus) HasHasUncommittedChanges() bool {
	if o != nil && o.HasUncommittedChanges != nil {
		return true
	}

	return false
}

// SetHasUncommittedChanges gets a reference to the given bool and assigns it to the HasUncommittedChanges field.
func (o *CodespaceGitStatus) SetHasUncommittedChanges(v bool) {
	o.HasUncommittedChanges = &v
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *CodespaceGitStatus) GetRef() string {
	if o == nil || o.Ref == nil {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodespaceGitStatus) GetRefOk() (*string, bool) {
	if o == nil || o.Ref == nil {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *CodespaceGitStatus) HasRef() bool {
	if o != nil && o.Ref != nil {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *CodespaceGitStatus) SetRef(v string) {
	o.Ref = &v
}

func (o CodespaceGitStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ahead != nil {
		toSerialize["ahead"] = o.Ahead
	}
	if o.Behind != nil {
		toSerialize["behind"] = o.Behind
	}
	if o.HasUnpushedChanges != nil {
		toSerialize["has_unpushed_changes"] = o.HasUnpushedChanges
	}
	if o.HasUncommittedChanges != nil {
		toSerialize["has_uncommitted_changes"] = o.HasUncommittedChanges
	}
	if o.Ref != nil {
		toSerialize["ref"] = o.Ref
	}
	return json.Marshal(toSerialize)
}

type NullableCodespaceGitStatus struct {
	value *CodespaceGitStatus
	isSet bool
}

func (v NullableCodespaceGitStatus) Get() *CodespaceGitStatus {
	return v.value
}

func (v *NullableCodespaceGitStatus) Set(val *CodespaceGitStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCodespaceGitStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCodespaceGitStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodespaceGitStatus(val *CodespaceGitStatus) *NullableCodespaceGitStatus {
	return &NullableCodespaceGitStatus{value: val, isSet: true}
}

func (v NullableCodespaceGitStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodespaceGitStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


