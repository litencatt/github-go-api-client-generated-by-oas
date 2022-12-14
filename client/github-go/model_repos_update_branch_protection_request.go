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

// ReposUpdateBranchProtectionRequest struct for ReposUpdateBranchProtectionRequest
type ReposUpdateBranchProtectionRequest struct {
	RequiredStatusChecks NullableReposUpdateBranchProtectionRequestRequiredStatusChecks `json:"required_status_checks"`
	// Enforce all configured restrictions for administrators. Set to `true` to enforce required status checks for repository administrators. Set to `null` to disable.
	EnforceAdmins NullableBool `json:"enforce_admins"`
	RequiredPullRequestReviews NullableReposUpdateBranchProtectionRequestRequiredPullRequestReviews `json:"required_pull_request_reviews"`
	Restrictions NullableReposUpdateBranchProtectionRequestRestrictions `json:"restrictions"`
	// Enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch. Set to `true` to enforce a linear commit history. Set to `false` to disable a linear commit Git history. Your repository must allow squash merging or rebase merging before you can enable a linear commit history. Default: `false`. For more information, see \"[Requiring a linear commit history](https://docs.github.com/github/administering-a-repository/requiring-a-linear-commit-history)\" in the GitHub Help documentation.
	RequiredLinearHistory *bool `json:"required_linear_history,omitempty"`
	// Permits force pushes to the protected branch by anyone with write access to the repository. Set to `true` to allow force pushes. Set to `false` or `null` to block force pushes. Default: `false`. For more information, see \"[Enabling force pushes to a protected branch](https://docs.github.com/en/github/administering-a-repository/enabling-force-pushes-to-a-protected-branch)\" in the GitHub Help documentation.\"
	AllowForcePushes NullableBool `json:"allow_force_pushes,omitempty"`
	// Allows deletion of the protected branch by anyone with write access to the repository. Set to `false` to prevent deletion of the protected branch. Default: `false`. For more information, see \"[Enabling force pushes to a protected branch](https://docs.github.com/en/github/administering-a-repository/enabling-force-pushes-to-a-protected-branch)\" in the GitHub Help documentation.
	AllowDeletions *bool `json:"allow_deletions,omitempty"`
	// If set to `true`, the `restrictions` branch protection settings which limits who can push will also block pushes which create new branches, unless the push is initiated by a user, team, or app which has the ability to push. Set to `true` to restrict new branch creation. Default: `false`.
	BlockCreations *bool `json:"block_creations,omitempty"`
	// Requires all conversations on code to be resolved before a pull request can be merged into a branch that matches this rule. Set to `false` to disable. Default: `false`.
	RequiredConversationResolution *bool `json:"required_conversation_resolution,omitempty"`
}

// NewReposUpdateBranchProtectionRequest instantiates a new ReposUpdateBranchProtectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReposUpdateBranchProtectionRequest(requiredStatusChecks NullableReposUpdateBranchProtectionRequestRequiredStatusChecks, enforceAdmins NullableBool, requiredPullRequestReviews NullableReposUpdateBranchProtectionRequestRequiredPullRequestReviews, restrictions NullableReposUpdateBranchProtectionRequestRestrictions) *ReposUpdateBranchProtectionRequest {
	this := ReposUpdateBranchProtectionRequest{}
	this.RequiredStatusChecks = requiredStatusChecks
	this.EnforceAdmins = enforceAdmins
	this.RequiredPullRequestReviews = requiredPullRequestReviews
	this.Restrictions = restrictions
	return &this
}

// NewReposUpdateBranchProtectionRequestWithDefaults instantiates a new ReposUpdateBranchProtectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReposUpdateBranchProtectionRequestWithDefaults() *ReposUpdateBranchProtectionRequest {
	this := ReposUpdateBranchProtectionRequest{}
	return &this
}

// GetRequiredStatusChecks returns the RequiredStatusChecks field value
// If the value is explicit nil, the zero value for ReposUpdateBranchProtectionRequestRequiredStatusChecks will be returned
func (o *ReposUpdateBranchProtectionRequest) GetRequiredStatusChecks() ReposUpdateBranchProtectionRequestRequiredStatusChecks {
	if o == nil || o.RequiredStatusChecks.Get() == nil {
		var ret ReposUpdateBranchProtectionRequestRequiredStatusChecks
		return ret
	}

	return *o.RequiredStatusChecks.Get()
}

// GetRequiredStatusChecksOk returns a tuple with the RequiredStatusChecks field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReposUpdateBranchProtectionRequest) GetRequiredStatusChecksOk() (*ReposUpdateBranchProtectionRequestRequiredStatusChecks, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiredStatusChecks.Get(), o.RequiredStatusChecks.IsSet()
}

// SetRequiredStatusChecks sets field value
func (o *ReposUpdateBranchProtectionRequest) SetRequiredStatusChecks(v ReposUpdateBranchProtectionRequestRequiredStatusChecks) {
	o.RequiredStatusChecks.Set(&v)
}

// GetEnforceAdmins returns the EnforceAdmins field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *ReposUpdateBranchProtectionRequest) GetEnforceAdmins() bool {
	if o == nil || o.EnforceAdmins.Get() == nil {
		var ret bool
		return ret
	}

	return *o.EnforceAdmins.Get()
}

// GetEnforceAdminsOk returns a tuple with the EnforceAdmins field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReposUpdateBranchProtectionRequest) GetEnforceAdminsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnforceAdmins.Get(), o.EnforceAdmins.IsSet()
}

// SetEnforceAdmins sets field value
func (o *ReposUpdateBranchProtectionRequest) SetEnforceAdmins(v bool) {
	o.EnforceAdmins.Set(&v)
}

// GetRequiredPullRequestReviews returns the RequiredPullRequestReviews field value
// If the value is explicit nil, the zero value for ReposUpdateBranchProtectionRequestRequiredPullRequestReviews will be returned
func (o *ReposUpdateBranchProtectionRequest) GetRequiredPullRequestReviews() ReposUpdateBranchProtectionRequestRequiredPullRequestReviews {
	if o == nil || o.RequiredPullRequestReviews.Get() == nil {
		var ret ReposUpdateBranchProtectionRequestRequiredPullRequestReviews
		return ret
	}

	return *o.RequiredPullRequestReviews.Get()
}

// GetRequiredPullRequestReviewsOk returns a tuple with the RequiredPullRequestReviews field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReposUpdateBranchProtectionRequest) GetRequiredPullRequestReviewsOk() (*ReposUpdateBranchProtectionRequestRequiredPullRequestReviews, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiredPullRequestReviews.Get(), o.RequiredPullRequestReviews.IsSet()
}

// SetRequiredPullRequestReviews sets field value
func (o *ReposUpdateBranchProtectionRequest) SetRequiredPullRequestReviews(v ReposUpdateBranchProtectionRequestRequiredPullRequestReviews) {
	o.RequiredPullRequestReviews.Set(&v)
}

// GetRestrictions returns the Restrictions field value
// If the value is explicit nil, the zero value for ReposUpdateBranchProtectionRequestRestrictions will be returned
func (o *ReposUpdateBranchProtectionRequest) GetRestrictions() ReposUpdateBranchProtectionRequestRestrictions {
	if o == nil || o.Restrictions.Get() == nil {
		var ret ReposUpdateBranchProtectionRequestRestrictions
		return ret
	}

	return *o.Restrictions.Get()
}

// GetRestrictionsOk returns a tuple with the Restrictions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReposUpdateBranchProtectionRequest) GetRestrictionsOk() (*ReposUpdateBranchProtectionRequestRestrictions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Restrictions.Get(), o.Restrictions.IsSet()
}

// SetRestrictions sets field value
func (o *ReposUpdateBranchProtectionRequest) SetRestrictions(v ReposUpdateBranchProtectionRequestRestrictions) {
	o.Restrictions.Set(&v)
}

// GetRequiredLinearHistory returns the RequiredLinearHistory field value if set, zero value otherwise.
func (o *ReposUpdateBranchProtectionRequest) GetRequiredLinearHistory() bool {
	if o == nil || o.RequiredLinearHistory == nil {
		var ret bool
		return ret
	}
	return *o.RequiredLinearHistory
}

// GetRequiredLinearHistoryOk returns a tuple with the RequiredLinearHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposUpdateBranchProtectionRequest) GetRequiredLinearHistoryOk() (*bool, bool) {
	if o == nil || o.RequiredLinearHistory == nil {
		return nil, false
	}
	return o.RequiredLinearHistory, true
}

// HasRequiredLinearHistory returns a boolean if a field has been set.
func (o *ReposUpdateBranchProtectionRequest) HasRequiredLinearHistory() bool {
	if o != nil && o.RequiredLinearHistory != nil {
		return true
	}

	return false
}

// SetRequiredLinearHistory gets a reference to the given bool and assigns it to the RequiredLinearHistory field.
func (o *ReposUpdateBranchProtectionRequest) SetRequiredLinearHistory(v bool) {
	o.RequiredLinearHistory = &v
}

// GetAllowForcePushes returns the AllowForcePushes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReposUpdateBranchProtectionRequest) GetAllowForcePushes() bool {
	if o == nil || o.AllowForcePushes.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AllowForcePushes.Get()
}

// GetAllowForcePushesOk returns a tuple with the AllowForcePushes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReposUpdateBranchProtectionRequest) GetAllowForcePushesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowForcePushes.Get(), o.AllowForcePushes.IsSet()
}

// HasAllowForcePushes returns a boolean if a field has been set.
func (o *ReposUpdateBranchProtectionRequest) HasAllowForcePushes() bool {
	if o != nil && o.AllowForcePushes.IsSet() {
		return true
	}

	return false
}

// SetAllowForcePushes gets a reference to the given NullableBool and assigns it to the AllowForcePushes field.
func (o *ReposUpdateBranchProtectionRequest) SetAllowForcePushes(v bool) {
	o.AllowForcePushes.Set(&v)
}
// SetAllowForcePushesNil sets the value for AllowForcePushes to be an explicit nil
func (o *ReposUpdateBranchProtectionRequest) SetAllowForcePushesNil() {
	o.AllowForcePushes.Set(nil)
}

// UnsetAllowForcePushes ensures that no value is present for AllowForcePushes, not even an explicit nil
func (o *ReposUpdateBranchProtectionRequest) UnsetAllowForcePushes() {
	o.AllowForcePushes.Unset()
}

// GetAllowDeletions returns the AllowDeletions field value if set, zero value otherwise.
func (o *ReposUpdateBranchProtectionRequest) GetAllowDeletions() bool {
	if o == nil || o.AllowDeletions == nil {
		var ret bool
		return ret
	}
	return *o.AllowDeletions
}

// GetAllowDeletionsOk returns a tuple with the AllowDeletions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposUpdateBranchProtectionRequest) GetAllowDeletionsOk() (*bool, bool) {
	if o == nil || o.AllowDeletions == nil {
		return nil, false
	}
	return o.AllowDeletions, true
}

// HasAllowDeletions returns a boolean if a field has been set.
func (o *ReposUpdateBranchProtectionRequest) HasAllowDeletions() bool {
	if o != nil && o.AllowDeletions != nil {
		return true
	}

	return false
}

// SetAllowDeletions gets a reference to the given bool and assigns it to the AllowDeletions field.
func (o *ReposUpdateBranchProtectionRequest) SetAllowDeletions(v bool) {
	o.AllowDeletions = &v
}

// GetBlockCreations returns the BlockCreations field value if set, zero value otherwise.
func (o *ReposUpdateBranchProtectionRequest) GetBlockCreations() bool {
	if o == nil || o.BlockCreations == nil {
		var ret bool
		return ret
	}
	return *o.BlockCreations
}

// GetBlockCreationsOk returns a tuple with the BlockCreations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposUpdateBranchProtectionRequest) GetBlockCreationsOk() (*bool, bool) {
	if o == nil || o.BlockCreations == nil {
		return nil, false
	}
	return o.BlockCreations, true
}

// HasBlockCreations returns a boolean if a field has been set.
func (o *ReposUpdateBranchProtectionRequest) HasBlockCreations() bool {
	if o != nil && o.BlockCreations != nil {
		return true
	}

	return false
}

// SetBlockCreations gets a reference to the given bool and assigns it to the BlockCreations field.
func (o *ReposUpdateBranchProtectionRequest) SetBlockCreations(v bool) {
	o.BlockCreations = &v
}

// GetRequiredConversationResolution returns the RequiredConversationResolution field value if set, zero value otherwise.
func (o *ReposUpdateBranchProtectionRequest) GetRequiredConversationResolution() bool {
	if o == nil || o.RequiredConversationResolution == nil {
		var ret bool
		return ret
	}
	return *o.RequiredConversationResolution
}

// GetRequiredConversationResolutionOk returns a tuple with the RequiredConversationResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposUpdateBranchProtectionRequest) GetRequiredConversationResolutionOk() (*bool, bool) {
	if o == nil || o.RequiredConversationResolution == nil {
		return nil, false
	}
	return o.RequiredConversationResolution, true
}

// HasRequiredConversationResolution returns a boolean if a field has been set.
func (o *ReposUpdateBranchProtectionRequest) HasRequiredConversationResolution() bool {
	if o != nil && o.RequiredConversationResolution != nil {
		return true
	}

	return false
}

// SetRequiredConversationResolution gets a reference to the given bool and assigns it to the RequiredConversationResolution field.
func (o *ReposUpdateBranchProtectionRequest) SetRequiredConversationResolution(v bool) {
	o.RequiredConversationResolution = &v
}

func (o ReposUpdateBranchProtectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["required_status_checks"] = o.RequiredStatusChecks.Get()
	}
	if true {
		toSerialize["enforce_admins"] = o.EnforceAdmins.Get()
	}
	if true {
		toSerialize["required_pull_request_reviews"] = o.RequiredPullRequestReviews.Get()
	}
	if true {
		toSerialize["restrictions"] = o.Restrictions.Get()
	}
	if o.RequiredLinearHistory != nil {
		toSerialize["required_linear_history"] = o.RequiredLinearHistory
	}
	if o.AllowForcePushes.IsSet() {
		toSerialize["allow_force_pushes"] = o.AllowForcePushes.Get()
	}
	if o.AllowDeletions != nil {
		toSerialize["allow_deletions"] = o.AllowDeletions
	}
	if o.BlockCreations != nil {
		toSerialize["block_creations"] = o.BlockCreations
	}
	if o.RequiredConversationResolution != nil {
		toSerialize["required_conversation_resolution"] = o.RequiredConversationResolution
	}
	return json.Marshal(toSerialize)
}

type NullableReposUpdateBranchProtectionRequest struct {
	value *ReposUpdateBranchProtectionRequest
	isSet bool
}

func (v NullableReposUpdateBranchProtectionRequest) Get() *ReposUpdateBranchProtectionRequest {
	return v.value
}

func (v *NullableReposUpdateBranchProtectionRequest) Set(val *ReposUpdateBranchProtectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReposUpdateBranchProtectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReposUpdateBranchProtectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReposUpdateBranchProtectionRequest(val *ReposUpdateBranchProtectionRequest) *NullableReposUpdateBranchProtectionRequest {
	return &NullableReposUpdateBranchProtectionRequest{value: val, isSet: true}
}

func (v NullableReposUpdateBranchProtectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReposUpdateBranchProtectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


