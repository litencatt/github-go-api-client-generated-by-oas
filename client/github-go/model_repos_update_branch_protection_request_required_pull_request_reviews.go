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

// ReposUpdateBranchProtectionRequestRequiredPullRequestReviews Require at least one approving review on a pull request, before merging. Set to `null` to disable.
type ReposUpdateBranchProtectionRequestRequiredPullRequestReviews struct {
	DismissalRestrictions *ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsDismissalRestrictions `json:"dismissal_restrictions,omitempty"`
	// Set to `true` if you want to automatically dismiss approving reviews when someone pushes a new commit.
	DismissStaleReviews *bool `json:"dismiss_stale_reviews,omitempty"`
	// Blocks merging pull requests until [code owners](https://docs.github.com/articles/about-code-owners/) review them.
	RequireCodeOwnerReviews *bool `json:"require_code_owner_reviews,omitempty"`
	// Specify the number of reviewers required to approve pull requests. Use a number between 1 and 6 or 0 to not require reviewers.
	RequiredApprovingReviewCount *int32 `json:"required_approving_review_count,omitempty"`
	BypassPullRequestAllowances *ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsBypassPullRequestAllowances `json:"bypass_pull_request_allowances,omitempty"`
}

// NewReposUpdateBranchProtectionRequestRequiredPullRequestReviews instantiates a new ReposUpdateBranchProtectionRequestRequiredPullRequestReviews object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReposUpdateBranchProtectionRequestRequiredPullRequestReviews() *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews {
	this := ReposUpdateBranchProtectionRequestRequiredPullRequestReviews{}
	return &this
}

// NewReposUpdateBranchProtectionRequestRequiredPullRequestReviewsWithDefaults instantiates a new ReposUpdateBranchProtectionRequestRequiredPullRequestReviews object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReposUpdateBranchProtectionRequestRequiredPullRequestReviewsWithDefaults() *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews {
	this := ReposUpdateBranchProtectionRequestRequiredPullRequestReviews{}
	return &this
}

// GetDismissalRestrictions returns the DismissalRestrictions field value if set, zero value otherwise.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) GetDismissalRestrictions() ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsDismissalRestrictions {
	if o == nil || o.DismissalRestrictions == nil {
		var ret ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsDismissalRestrictions
		return ret
	}
	return *o.DismissalRestrictions
}

// GetDismissalRestrictionsOk returns a tuple with the DismissalRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) GetDismissalRestrictionsOk() (*ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsDismissalRestrictions, bool) {
	if o == nil || o.DismissalRestrictions == nil {
		return nil, false
	}
	return o.DismissalRestrictions, true
}

// HasDismissalRestrictions returns a boolean if a field has been set.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) HasDismissalRestrictions() bool {
	if o != nil && o.DismissalRestrictions != nil {
		return true
	}

	return false
}

// SetDismissalRestrictions gets a reference to the given ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsDismissalRestrictions and assigns it to the DismissalRestrictions field.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) SetDismissalRestrictions(v ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsDismissalRestrictions) {
	o.DismissalRestrictions = &v
}

// GetDismissStaleReviews returns the DismissStaleReviews field value if set, zero value otherwise.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) GetDismissStaleReviews() bool {
	if o == nil || o.DismissStaleReviews == nil {
		var ret bool
		return ret
	}
	return *o.DismissStaleReviews
}

// GetDismissStaleReviewsOk returns a tuple with the DismissStaleReviews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) GetDismissStaleReviewsOk() (*bool, bool) {
	if o == nil || o.DismissStaleReviews == nil {
		return nil, false
	}
	return o.DismissStaleReviews, true
}

// HasDismissStaleReviews returns a boolean if a field has been set.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) HasDismissStaleReviews() bool {
	if o != nil && o.DismissStaleReviews != nil {
		return true
	}

	return false
}

// SetDismissStaleReviews gets a reference to the given bool and assigns it to the DismissStaleReviews field.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) SetDismissStaleReviews(v bool) {
	o.DismissStaleReviews = &v
}

// GetRequireCodeOwnerReviews returns the RequireCodeOwnerReviews field value if set, zero value otherwise.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) GetRequireCodeOwnerReviews() bool {
	if o == nil || o.RequireCodeOwnerReviews == nil {
		var ret bool
		return ret
	}
	return *o.RequireCodeOwnerReviews
}

// GetRequireCodeOwnerReviewsOk returns a tuple with the RequireCodeOwnerReviews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) GetRequireCodeOwnerReviewsOk() (*bool, bool) {
	if o == nil || o.RequireCodeOwnerReviews == nil {
		return nil, false
	}
	return o.RequireCodeOwnerReviews, true
}

// HasRequireCodeOwnerReviews returns a boolean if a field has been set.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) HasRequireCodeOwnerReviews() bool {
	if o != nil && o.RequireCodeOwnerReviews != nil {
		return true
	}

	return false
}

// SetRequireCodeOwnerReviews gets a reference to the given bool and assigns it to the RequireCodeOwnerReviews field.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) SetRequireCodeOwnerReviews(v bool) {
	o.RequireCodeOwnerReviews = &v
}

// GetRequiredApprovingReviewCount returns the RequiredApprovingReviewCount field value if set, zero value otherwise.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) GetRequiredApprovingReviewCount() int32 {
	if o == nil || o.RequiredApprovingReviewCount == nil {
		var ret int32
		return ret
	}
	return *o.RequiredApprovingReviewCount
}

// GetRequiredApprovingReviewCountOk returns a tuple with the RequiredApprovingReviewCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) GetRequiredApprovingReviewCountOk() (*int32, bool) {
	if o == nil || o.RequiredApprovingReviewCount == nil {
		return nil, false
	}
	return o.RequiredApprovingReviewCount, true
}

// HasRequiredApprovingReviewCount returns a boolean if a field has been set.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) HasRequiredApprovingReviewCount() bool {
	if o != nil && o.RequiredApprovingReviewCount != nil {
		return true
	}

	return false
}

// SetRequiredApprovingReviewCount gets a reference to the given int32 and assigns it to the RequiredApprovingReviewCount field.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) SetRequiredApprovingReviewCount(v int32) {
	o.RequiredApprovingReviewCount = &v
}

// GetBypassPullRequestAllowances returns the BypassPullRequestAllowances field value if set, zero value otherwise.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) GetBypassPullRequestAllowances() ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsBypassPullRequestAllowances {
	if o == nil || o.BypassPullRequestAllowances == nil {
		var ret ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsBypassPullRequestAllowances
		return ret
	}
	return *o.BypassPullRequestAllowances
}

// GetBypassPullRequestAllowancesOk returns a tuple with the BypassPullRequestAllowances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) GetBypassPullRequestAllowancesOk() (*ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsBypassPullRequestAllowances, bool) {
	if o == nil || o.BypassPullRequestAllowances == nil {
		return nil, false
	}
	return o.BypassPullRequestAllowances, true
}

// HasBypassPullRequestAllowances returns a boolean if a field has been set.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) HasBypassPullRequestAllowances() bool {
	if o != nil && o.BypassPullRequestAllowances != nil {
		return true
	}

	return false
}

// SetBypassPullRequestAllowances gets a reference to the given ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsBypassPullRequestAllowances and assigns it to the BypassPullRequestAllowances field.
func (o *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) SetBypassPullRequestAllowances(v ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsBypassPullRequestAllowances) {
	o.BypassPullRequestAllowances = &v
}

func (o ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DismissalRestrictions != nil {
		toSerialize["dismissal_restrictions"] = o.DismissalRestrictions
	}
	if o.DismissStaleReviews != nil {
		toSerialize["dismiss_stale_reviews"] = o.DismissStaleReviews
	}
	if o.RequireCodeOwnerReviews != nil {
		toSerialize["require_code_owner_reviews"] = o.RequireCodeOwnerReviews
	}
	if o.RequiredApprovingReviewCount != nil {
		toSerialize["required_approving_review_count"] = o.RequiredApprovingReviewCount
	}
	if o.BypassPullRequestAllowances != nil {
		toSerialize["bypass_pull_request_allowances"] = o.BypassPullRequestAllowances
	}
	return json.Marshal(toSerialize)
}

type NullableReposUpdateBranchProtectionRequestRequiredPullRequestReviews struct {
	value *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews
	isSet bool
}

func (v NullableReposUpdateBranchProtectionRequestRequiredPullRequestReviews) Get() *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews {
	return v.value
}

func (v *NullableReposUpdateBranchProtectionRequestRequiredPullRequestReviews) Set(val *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) {
	v.value = val
	v.isSet = true
}

func (v NullableReposUpdateBranchProtectionRequestRequiredPullRequestReviews) IsSet() bool {
	return v.isSet
}

func (v *NullableReposUpdateBranchProtectionRequestRequiredPullRequestReviews) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReposUpdateBranchProtectionRequestRequiredPullRequestReviews(val *ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) *NullableReposUpdateBranchProtectionRequestRequiredPullRequestReviews {
	return &NullableReposUpdateBranchProtectionRequestRequiredPullRequestReviews{value: val, isSet: true}
}

func (v NullableReposUpdateBranchProtectionRequestRequiredPullRequestReviews) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReposUpdateBranchProtectionRequestRequiredPullRequestReviews) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


