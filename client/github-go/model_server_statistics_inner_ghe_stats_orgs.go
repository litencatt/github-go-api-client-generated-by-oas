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

// ServerStatisticsInnerGheStatsOrgs struct for ServerStatisticsInnerGheStatsOrgs
type ServerStatisticsInnerGheStatsOrgs struct {
	TotalOrgs *int32 `json:"total_orgs,omitempty"`
	DisabledOrgs *int32 `json:"disabled_orgs,omitempty"`
	TotalTeams *int32 `json:"total_teams,omitempty"`
	TotalTeamMembers *int32 `json:"total_team_members,omitempty"`
}

// NewServerStatisticsInnerGheStatsOrgs instantiates a new ServerStatisticsInnerGheStatsOrgs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerStatisticsInnerGheStatsOrgs() *ServerStatisticsInnerGheStatsOrgs {
	this := ServerStatisticsInnerGheStatsOrgs{}
	return &this
}

// NewServerStatisticsInnerGheStatsOrgsWithDefaults instantiates a new ServerStatisticsInnerGheStatsOrgs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerStatisticsInnerGheStatsOrgsWithDefaults() *ServerStatisticsInnerGheStatsOrgs {
	this := ServerStatisticsInnerGheStatsOrgs{}
	return &this
}

// GetTotalOrgs returns the TotalOrgs field value if set, zero value otherwise.
func (o *ServerStatisticsInnerGheStatsOrgs) GetTotalOrgs() int32 {
	if o == nil || o.TotalOrgs == nil {
		var ret int32
		return ret
	}
	return *o.TotalOrgs
}

// GetTotalOrgsOk returns a tuple with the TotalOrgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStatisticsInnerGheStatsOrgs) GetTotalOrgsOk() (*int32, bool) {
	if o == nil || o.TotalOrgs == nil {
		return nil, false
	}
	return o.TotalOrgs, true
}

// HasTotalOrgs returns a boolean if a field has been set.
func (o *ServerStatisticsInnerGheStatsOrgs) HasTotalOrgs() bool {
	if o != nil && o.TotalOrgs != nil {
		return true
	}

	return false
}

// SetTotalOrgs gets a reference to the given int32 and assigns it to the TotalOrgs field.
func (o *ServerStatisticsInnerGheStatsOrgs) SetTotalOrgs(v int32) {
	o.TotalOrgs = &v
}

// GetDisabledOrgs returns the DisabledOrgs field value if set, zero value otherwise.
func (o *ServerStatisticsInnerGheStatsOrgs) GetDisabledOrgs() int32 {
	if o == nil || o.DisabledOrgs == nil {
		var ret int32
		return ret
	}
	return *o.DisabledOrgs
}

// GetDisabledOrgsOk returns a tuple with the DisabledOrgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStatisticsInnerGheStatsOrgs) GetDisabledOrgsOk() (*int32, bool) {
	if o == nil || o.DisabledOrgs == nil {
		return nil, false
	}
	return o.DisabledOrgs, true
}

// HasDisabledOrgs returns a boolean if a field has been set.
func (o *ServerStatisticsInnerGheStatsOrgs) HasDisabledOrgs() bool {
	if o != nil && o.DisabledOrgs != nil {
		return true
	}

	return false
}

// SetDisabledOrgs gets a reference to the given int32 and assigns it to the DisabledOrgs field.
func (o *ServerStatisticsInnerGheStatsOrgs) SetDisabledOrgs(v int32) {
	o.DisabledOrgs = &v
}

// GetTotalTeams returns the TotalTeams field value if set, zero value otherwise.
func (o *ServerStatisticsInnerGheStatsOrgs) GetTotalTeams() int32 {
	if o == nil || o.TotalTeams == nil {
		var ret int32
		return ret
	}
	return *o.TotalTeams
}

// GetTotalTeamsOk returns a tuple with the TotalTeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStatisticsInnerGheStatsOrgs) GetTotalTeamsOk() (*int32, bool) {
	if o == nil || o.TotalTeams == nil {
		return nil, false
	}
	return o.TotalTeams, true
}

// HasTotalTeams returns a boolean if a field has been set.
func (o *ServerStatisticsInnerGheStatsOrgs) HasTotalTeams() bool {
	if o != nil && o.TotalTeams != nil {
		return true
	}

	return false
}

// SetTotalTeams gets a reference to the given int32 and assigns it to the TotalTeams field.
func (o *ServerStatisticsInnerGheStatsOrgs) SetTotalTeams(v int32) {
	o.TotalTeams = &v
}

// GetTotalTeamMembers returns the TotalTeamMembers field value if set, zero value otherwise.
func (o *ServerStatisticsInnerGheStatsOrgs) GetTotalTeamMembers() int32 {
	if o == nil || o.TotalTeamMembers == nil {
		var ret int32
		return ret
	}
	return *o.TotalTeamMembers
}

// GetTotalTeamMembersOk returns a tuple with the TotalTeamMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStatisticsInnerGheStatsOrgs) GetTotalTeamMembersOk() (*int32, bool) {
	if o == nil || o.TotalTeamMembers == nil {
		return nil, false
	}
	return o.TotalTeamMembers, true
}

// HasTotalTeamMembers returns a boolean if a field has been set.
func (o *ServerStatisticsInnerGheStatsOrgs) HasTotalTeamMembers() bool {
	if o != nil && o.TotalTeamMembers != nil {
		return true
	}

	return false
}

// SetTotalTeamMembers gets a reference to the given int32 and assigns it to the TotalTeamMembers field.
func (o *ServerStatisticsInnerGheStatsOrgs) SetTotalTeamMembers(v int32) {
	o.TotalTeamMembers = &v
}

func (o ServerStatisticsInnerGheStatsOrgs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TotalOrgs != nil {
		toSerialize["total_orgs"] = o.TotalOrgs
	}
	if o.DisabledOrgs != nil {
		toSerialize["disabled_orgs"] = o.DisabledOrgs
	}
	if o.TotalTeams != nil {
		toSerialize["total_teams"] = o.TotalTeams
	}
	if o.TotalTeamMembers != nil {
		toSerialize["total_team_members"] = o.TotalTeamMembers
	}
	return json.Marshal(toSerialize)
}

type NullableServerStatisticsInnerGheStatsOrgs struct {
	value *ServerStatisticsInnerGheStatsOrgs
	isSet bool
}

func (v NullableServerStatisticsInnerGheStatsOrgs) Get() *ServerStatisticsInnerGheStatsOrgs {
	return v.value
}

func (v *NullableServerStatisticsInnerGheStatsOrgs) Set(val *ServerStatisticsInnerGheStatsOrgs) {
	v.value = val
	v.isSet = true
}

func (v NullableServerStatisticsInnerGheStatsOrgs) IsSet() bool {
	return v.isSet
}

func (v *NullableServerStatisticsInnerGheStatsOrgs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerStatisticsInnerGheStatsOrgs(val *ServerStatisticsInnerGheStatsOrgs) *NullableServerStatisticsInnerGheStatsOrgs {
	return &NullableServerStatisticsInnerGheStatsOrgs{value: val, isSet: true}
}

func (v NullableServerStatisticsInnerGheStatsOrgs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerStatisticsInnerGheStatsOrgs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


