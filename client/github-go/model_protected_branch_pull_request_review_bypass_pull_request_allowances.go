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

// ProtectedBranchPullRequestReviewBypassPullRequestAllowances Allow specific users, teams, or apps to bypass pull request requirements.
type ProtectedBranchPullRequestReviewBypassPullRequestAllowances struct {
	// The list of users allowed to bypass pull request requirements.
	Users []SimpleUser `json:"users,omitempty"`
	// The list of teams allowed to bypass pull request requirements.
	Teams []Team `json:"teams,omitempty"`
	// The list of apps allowed to bypass pull request requirements.
	Apps []Integration `json:"apps,omitempty"`
}

// NewProtectedBranchPullRequestReviewBypassPullRequestAllowances instantiates a new ProtectedBranchPullRequestReviewBypassPullRequestAllowances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtectedBranchPullRequestReviewBypassPullRequestAllowances() *ProtectedBranchPullRequestReviewBypassPullRequestAllowances {
	this := ProtectedBranchPullRequestReviewBypassPullRequestAllowances{}
	return &this
}

// NewProtectedBranchPullRequestReviewBypassPullRequestAllowancesWithDefaults instantiates a new ProtectedBranchPullRequestReviewBypassPullRequestAllowances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtectedBranchPullRequestReviewBypassPullRequestAllowancesWithDefaults() *ProtectedBranchPullRequestReviewBypassPullRequestAllowances {
	this := ProtectedBranchPullRequestReviewBypassPullRequestAllowances{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) GetUsers() []SimpleUser {
	if o == nil || o.Users == nil {
		var ret []SimpleUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) GetUsersOk() ([]SimpleUser, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []SimpleUser and assigns it to the Users field.
func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) SetUsers(v []SimpleUser) {
	o.Users = v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) GetTeams() []Team {
	if o == nil || o.Teams == nil {
		var ret []Team
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) GetTeamsOk() ([]Team, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) HasTeams() bool {
	if o != nil && o.Teams != nil {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []Team and assigns it to the Teams field.
func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) SetTeams(v []Team) {
	o.Teams = v
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) GetApps() []Integration {
	if o == nil || o.Apps == nil {
		var ret []Integration
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) GetAppsOk() ([]Integration, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given []Integration and assigns it to the Apps field.
func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) SetApps(v []Integration) {
	o.Apps = v
}

func (o ProtectedBranchPullRequestReviewBypassPullRequestAllowances) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}
	return json.Marshal(toSerialize)
}

type NullableProtectedBranchPullRequestReviewBypassPullRequestAllowances struct {
	value *ProtectedBranchPullRequestReviewBypassPullRequestAllowances
	isSet bool
}

func (v NullableProtectedBranchPullRequestReviewBypassPullRequestAllowances) Get() *ProtectedBranchPullRequestReviewBypassPullRequestAllowances {
	return v.value
}

func (v *NullableProtectedBranchPullRequestReviewBypassPullRequestAllowances) Set(val *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) {
	v.value = val
	v.isSet = true
}

func (v NullableProtectedBranchPullRequestReviewBypassPullRequestAllowances) IsSet() bool {
	return v.isSet
}

func (v *NullableProtectedBranchPullRequestReviewBypassPullRequestAllowances) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtectedBranchPullRequestReviewBypassPullRequestAllowances(val *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) *NullableProtectedBranchPullRequestReviewBypassPullRequestAllowances {
	return &NullableProtectedBranchPullRequestReviewBypassPullRequestAllowances{value: val, isSet: true}
}

func (v NullableProtectedBranchPullRequestReviewBypassPullRequestAllowances) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtectedBranchPullRequestReviewBypassPullRequestAllowances) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


