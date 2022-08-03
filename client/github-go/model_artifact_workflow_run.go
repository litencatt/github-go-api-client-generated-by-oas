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

// ArtifactWorkflowRun struct for ArtifactWorkflowRun
type ArtifactWorkflowRun struct {
	Id *int32 `json:"id,omitempty"`
	RepositoryId *int32 `json:"repository_id,omitempty"`
	HeadRepositoryId *int32 `json:"head_repository_id,omitempty"`
	HeadBranch *string `json:"head_branch,omitempty"`
	HeadSha *string `json:"head_sha,omitempty"`
}

// NewArtifactWorkflowRun instantiates a new ArtifactWorkflowRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactWorkflowRun() *ArtifactWorkflowRun {
	this := ArtifactWorkflowRun{}
	return &this
}

// NewArtifactWorkflowRunWithDefaults instantiates a new ArtifactWorkflowRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactWorkflowRunWithDefaults() *ArtifactWorkflowRun {
	this := ArtifactWorkflowRun{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ArtifactWorkflowRun) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactWorkflowRun) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ArtifactWorkflowRun) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ArtifactWorkflowRun) SetId(v int32) {
	o.Id = &v
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise.
func (o *ArtifactWorkflowRun) GetRepositoryId() int32 {
	if o == nil || o.RepositoryId == nil {
		var ret int32
		return ret
	}
	return *o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactWorkflowRun) GetRepositoryIdOk() (*int32, bool) {
	if o == nil || o.RepositoryId == nil {
		return nil, false
	}
	return o.RepositoryId, true
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *ArtifactWorkflowRun) HasRepositoryId() bool {
	if o != nil && o.RepositoryId != nil {
		return true
	}

	return false
}

// SetRepositoryId gets a reference to the given int32 and assigns it to the RepositoryId field.
func (o *ArtifactWorkflowRun) SetRepositoryId(v int32) {
	o.RepositoryId = &v
}

// GetHeadRepositoryId returns the HeadRepositoryId field value if set, zero value otherwise.
func (o *ArtifactWorkflowRun) GetHeadRepositoryId() int32 {
	if o == nil || o.HeadRepositoryId == nil {
		var ret int32
		return ret
	}
	return *o.HeadRepositoryId
}

// GetHeadRepositoryIdOk returns a tuple with the HeadRepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactWorkflowRun) GetHeadRepositoryIdOk() (*int32, bool) {
	if o == nil || o.HeadRepositoryId == nil {
		return nil, false
	}
	return o.HeadRepositoryId, true
}

// HasHeadRepositoryId returns a boolean if a field has been set.
func (o *ArtifactWorkflowRun) HasHeadRepositoryId() bool {
	if o != nil && o.HeadRepositoryId != nil {
		return true
	}

	return false
}

// SetHeadRepositoryId gets a reference to the given int32 and assigns it to the HeadRepositoryId field.
func (o *ArtifactWorkflowRun) SetHeadRepositoryId(v int32) {
	o.HeadRepositoryId = &v
}

// GetHeadBranch returns the HeadBranch field value if set, zero value otherwise.
func (o *ArtifactWorkflowRun) GetHeadBranch() string {
	if o == nil || o.HeadBranch == nil {
		var ret string
		return ret
	}
	return *o.HeadBranch
}

// GetHeadBranchOk returns a tuple with the HeadBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactWorkflowRun) GetHeadBranchOk() (*string, bool) {
	if o == nil || o.HeadBranch == nil {
		return nil, false
	}
	return o.HeadBranch, true
}

// HasHeadBranch returns a boolean if a field has been set.
func (o *ArtifactWorkflowRun) HasHeadBranch() bool {
	if o != nil && o.HeadBranch != nil {
		return true
	}

	return false
}

// SetHeadBranch gets a reference to the given string and assigns it to the HeadBranch field.
func (o *ArtifactWorkflowRun) SetHeadBranch(v string) {
	o.HeadBranch = &v
}

// GetHeadSha returns the HeadSha field value if set, zero value otherwise.
func (o *ArtifactWorkflowRun) GetHeadSha() string {
	if o == nil || o.HeadSha == nil {
		var ret string
		return ret
	}
	return *o.HeadSha
}

// GetHeadShaOk returns a tuple with the HeadSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactWorkflowRun) GetHeadShaOk() (*string, bool) {
	if o == nil || o.HeadSha == nil {
		return nil, false
	}
	return o.HeadSha, true
}

// HasHeadSha returns a boolean if a field has been set.
func (o *ArtifactWorkflowRun) HasHeadSha() bool {
	if o != nil && o.HeadSha != nil {
		return true
	}

	return false
}

// SetHeadSha gets a reference to the given string and assigns it to the HeadSha field.
func (o *ArtifactWorkflowRun) SetHeadSha(v string) {
	o.HeadSha = &v
}

func (o ArtifactWorkflowRun) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RepositoryId != nil {
		toSerialize["repository_id"] = o.RepositoryId
	}
	if o.HeadRepositoryId != nil {
		toSerialize["head_repository_id"] = o.HeadRepositoryId
	}
	if o.HeadBranch != nil {
		toSerialize["head_branch"] = o.HeadBranch
	}
	if o.HeadSha != nil {
		toSerialize["head_sha"] = o.HeadSha
	}
	return json.Marshal(toSerialize)
}

type NullableArtifactWorkflowRun struct {
	value *ArtifactWorkflowRun
	isSet bool
}

func (v NullableArtifactWorkflowRun) Get() *ArtifactWorkflowRun {
	return v.value
}

func (v *NullableArtifactWorkflowRun) Set(val *ArtifactWorkflowRun) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactWorkflowRun) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactWorkflowRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactWorkflowRun(val *ArtifactWorkflowRun) *NullableArtifactWorkflowRun {
	return &NullableArtifactWorkflowRun{value: val, isSet: true}
}

func (v NullableArtifactWorkflowRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactWorkflowRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

