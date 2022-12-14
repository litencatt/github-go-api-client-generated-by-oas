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

// SelectedActions struct for SelectedActions
type SelectedActions struct {
	// Whether GitHub-owned actions are allowed. For example, this includes the actions in the `actions` organization.
	GithubOwnedAllowed *bool `json:"github_owned_allowed,omitempty"`
	// Whether actions from GitHub Marketplace verified creators are allowed. Set to `true` to allow all actions by GitHub Marketplace verified creators.
	VerifiedAllowed *bool `json:"verified_allowed,omitempty"`
	// Specifies a list of string-matching patterns to allow specific action(s) and reusable workflow(s). Wildcards, tags, and SHAs are allowed. For example, `monalisa/octocat@*`, `monalisa/octocat@v2`, `monalisa/_*`.\"
	PatternsAllowed []string `json:"patterns_allowed,omitempty"`
}

// NewSelectedActions instantiates a new SelectedActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectedActions() *SelectedActions {
	this := SelectedActions{}
	return &this
}

// NewSelectedActionsWithDefaults instantiates a new SelectedActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectedActionsWithDefaults() *SelectedActions {
	this := SelectedActions{}
	return &this
}

// GetGithubOwnedAllowed returns the GithubOwnedAllowed field value if set, zero value otherwise.
func (o *SelectedActions) GetGithubOwnedAllowed() bool {
	if o == nil || o.GithubOwnedAllowed == nil {
		var ret bool
		return ret
	}
	return *o.GithubOwnedAllowed
}

// GetGithubOwnedAllowedOk returns a tuple with the GithubOwnedAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectedActions) GetGithubOwnedAllowedOk() (*bool, bool) {
	if o == nil || o.GithubOwnedAllowed == nil {
		return nil, false
	}
	return o.GithubOwnedAllowed, true
}

// HasGithubOwnedAllowed returns a boolean if a field has been set.
func (o *SelectedActions) HasGithubOwnedAllowed() bool {
	if o != nil && o.GithubOwnedAllowed != nil {
		return true
	}

	return false
}

// SetGithubOwnedAllowed gets a reference to the given bool and assigns it to the GithubOwnedAllowed field.
func (o *SelectedActions) SetGithubOwnedAllowed(v bool) {
	o.GithubOwnedAllowed = &v
}

// GetVerifiedAllowed returns the VerifiedAllowed field value if set, zero value otherwise.
func (o *SelectedActions) GetVerifiedAllowed() bool {
	if o == nil || o.VerifiedAllowed == nil {
		var ret bool
		return ret
	}
	return *o.VerifiedAllowed
}

// GetVerifiedAllowedOk returns a tuple with the VerifiedAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectedActions) GetVerifiedAllowedOk() (*bool, bool) {
	if o == nil || o.VerifiedAllowed == nil {
		return nil, false
	}
	return o.VerifiedAllowed, true
}

// HasVerifiedAllowed returns a boolean if a field has been set.
func (o *SelectedActions) HasVerifiedAllowed() bool {
	if o != nil && o.VerifiedAllowed != nil {
		return true
	}

	return false
}

// SetVerifiedAllowed gets a reference to the given bool and assigns it to the VerifiedAllowed field.
func (o *SelectedActions) SetVerifiedAllowed(v bool) {
	o.VerifiedAllowed = &v
}

// GetPatternsAllowed returns the PatternsAllowed field value if set, zero value otherwise.
func (o *SelectedActions) GetPatternsAllowed() []string {
	if o == nil || o.PatternsAllowed == nil {
		var ret []string
		return ret
	}
	return o.PatternsAllowed
}

// GetPatternsAllowedOk returns a tuple with the PatternsAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectedActions) GetPatternsAllowedOk() ([]string, bool) {
	if o == nil || o.PatternsAllowed == nil {
		return nil, false
	}
	return o.PatternsAllowed, true
}

// HasPatternsAllowed returns a boolean if a field has been set.
func (o *SelectedActions) HasPatternsAllowed() bool {
	if o != nil && o.PatternsAllowed != nil {
		return true
	}

	return false
}

// SetPatternsAllowed gets a reference to the given []string and assigns it to the PatternsAllowed field.
func (o *SelectedActions) SetPatternsAllowed(v []string) {
	o.PatternsAllowed = v
}

func (o SelectedActions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GithubOwnedAllowed != nil {
		toSerialize["github_owned_allowed"] = o.GithubOwnedAllowed
	}
	if o.VerifiedAllowed != nil {
		toSerialize["verified_allowed"] = o.VerifiedAllowed
	}
	if o.PatternsAllowed != nil {
		toSerialize["patterns_allowed"] = o.PatternsAllowed
	}
	return json.Marshal(toSerialize)
}

type NullableSelectedActions struct {
	value *SelectedActions
	isSet bool
}

func (v NullableSelectedActions) Get() *SelectedActions {
	return v.value
}

func (v *NullableSelectedActions) Set(val *SelectedActions) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectedActions) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectedActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectedActions(val *SelectedActions) *NullableSelectedActions {
	return &NullableSelectedActions{value: val, isSet: true}
}

func (v NullableSelectedActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectedActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


