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

// BranchWithProtection Branch With Protection
type BranchWithProtection struct {
	Name string `json:"name"`
	Commit Commit `json:"commit"`
	Links BranchWithProtectionLinks `json:"_links"`
	Protected bool `json:"protected"`
	Protection BranchProtection `json:"protection"`
	ProtectionUrl string `json:"protection_url"`
	Pattern *string `json:"pattern,omitempty"`
	RequiredApprovingReviewCount *int32 `json:"required_approving_review_count,omitempty"`
}

// NewBranchWithProtection instantiates a new BranchWithProtection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBranchWithProtection(name string, commit Commit, links BranchWithProtectionLinks, protected bool, protection BranchProtection, protectionUrl string) *BranchWithProtection {
	this := BranchWithProtection{}
	this.Name = name
	this.Commit = commit
	this.Links = links
	this.Protected = protected
	this.Protection = protection
	this.ProtectionUrl = protectionUrl
	return &this
}

// NewBranchWithProtectionWithDefaults instantiates a new BranchWithProtection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBranchWithProtectionWithDefaults() *BranchWithProtection {
	this := BranchWithProtection{}
	return &this
}

// GetName returns the Name field value
func (o *BranchWithProtection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BranchWithProtection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BranchWithProtection) SetName(v string) {
	o.Name = v
}

// GetCommit returns the Commit field value
func (o *BranchWithProtection) GetCommit() Commit {
	if o == nil {
		var ret Commit
		return ret
	}

	return o.Commit
}

// GetCommitOk returns a tuple with the Commit field value
// and a boolean to check if the value has been set.
func (o *BranchWithProtection) GetCommitOk() (*Commit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commit, true
}

// SetCommit sets field value
func (o *BranchWithProtection) SetCommit(v Commit) {
	o.Commit = v
}

// GetLinks returns the Links field value
func (o *BranchWithProtection) GetLinks() BranchWithProtectionLinks {
	if o == nil {
		var ret BranchWithProtectionLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BranchWithProtection) GetLinksOk() (*BranchWithProtectionLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BranchWithProtection) SetLinks(v BranchWithProtectionLinks) {
	o.Links = v
}

// GetProtected returns the Protected field value
func (o *BranchWithProtection) GetProtected() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Protected
}

// GetProtectedOk returns a tuple with the Protected field value
// and a boolean to check if the value has been set.
func (o *BranchWithProtection) GetProtectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protected, true
}

// SetProtected sets field value
func (o *BranchWithProtection) SetProtected(v bool) {
	o.Protected = v
}

// GetProtection returns the Protection field value
func (o *BranchWithProtection) GetProtection() BranchProtection {
	if o == nil {
		var ret BranchProtection
		return ret
	}

	return o.Protection
}

// GetProtectionOk returns a tuple with the Protection field value
// and a boolean to check if the value has been set.
func (o *BranchWithProtection) GetProtectionOk() (*BranchProtection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protection, true
}

// SetProtection sets field value
func (o *BranchWithProtection) SetProtection(v BranchProtection) {
	o.Protection = v
}

// GetProtectionUrl returns the ProtectionUrl field value
func (o *BranchWithProtection) GetProtectionUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProtectionUrl
}

// GetProtectionUrlOk returns a tuple with the ProtectionUrl field value
// and a boolean to check if the value has been set.
func (o *BranchWithProtection) GetProtectionUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProtectionUrl, true
}

// SetProtectionUrl sets field value
func (o *BranchWithProtection) SetProtectionUrl(v string) {
	o.ProtectionUrl = v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *BranchWithProtection) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchWithProtection) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *BranchWithProtection) HasPattern() bool {
	if o != nil && o.Pattern != nil {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *BranchWithProtection) SetPattern(v string) {
	o.Pattern = &v
}

// GetRequiredApprovingReviewCount returns the RequiredApprovingReviewCount field value if set, zero value otherwise.
func (o *BranchWithProtection) GetRequiredApprovingReviewCount() int32 {
	if o == nil || o.RequiredApprovingReviewCount == nil {
		var ret int32
		return ret
	}
	return *o.RequiredApprovingReviewCount
}

// GetRequiredApprovingReviewCountOk returns a tuple with the RequiredApprovingReviewCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchWithProtection) GetRequiredApprovingReviewCountOk() (*int32, bool) {
	if o == nil || o.RequiredApprovingReviewCount == nil {
		return nil, false
	}
	return o.RequiredApprovingReviewCount, true
}

// HasRequiredApprovingReviewCount returns a boolean if a field has been set.
func (o *BranchWithProtection) HasRequiredApprovingReviewCount() bool {
	if o != nil && o.RequiredApprovingReviewCount != nil {
		return true
	}

	return false
}

// SetRequiredApprovingReviewCount gets a reference to the given int32 and assigns it to the RequiredApprovingReviewCount field.
func (o *BranchWithProtection) SetRequiredApprovingReviewCount(v int32) {
	o.RequiredApprovingReviewCount = &v
}

func (o BranchWithProtection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["commit"] = o.Commit
	}
	if true {
		toSerialize["_links"] = o.Links
	}
	if true {
		toSerialize["protected"] = o.Protected
	}
	if true {
		toSerialize["protection"] = o.Protection
	}
	if true {
		toSerialize["protection_url"] = o.ProtectionUrl
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.RequiredApprovingReviewCount != nil {
		toSerialize["required_approving_review_count"] = o.RequiredApprovingReviewCount
	}
	return json.Marshal(toSerialize)
}

type NullableBranchWithProtection struct {
	value *BranchWithProtection
	isSet bool
}

func (v NullableBranchWithProtection) Get() *BranchWithProtection {
	return v.value
}

func (v *NullableBranchWithProtection) Set(val *BranchWithProtection) {
	v.value = val
	v.isSet = true
}

func (v NullableBranchWithProtection) IsSet() bool {
	return v.isSet
}

func (v *NullableBranchWithProtection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBranchWithProtection(val *BranchWithProtection) *NullableBranchWithProtection {
	return &NullableBranchWithProtection{value: val, isSet: true}
}

func (v NullableBranchWithProtection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBranchWithProtection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

