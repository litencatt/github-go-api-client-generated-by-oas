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

// ExternalGroupTeamsInner struct for ExternalGroupTeamsInner
type ExternalGroupTeamsInner struct {
	// The id for a team
	TeamId int32 `json:"team_id"`
	// The name of the team
	TeamName string `json:"team_name"`
}

// NewExternalGroupTeamsInner instantiates a new ExternalGroupTeamsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalGroupTeamsInner(teamId int32, teamName string) *ExternalGroupTeamsInner {
	this := ExternalGroupTeamsInner{}
	this.TeamId = teamId
	this.TeamName = teamName
	return &this
}

// NewExternalGroupTeamsInnerWithDefaults instantiates a new ExternalGroupTeamsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalGroupTeamsInnerWithDefaults() *ExternalGroupTeamsInner {
	this := ExternalGroupTeamsInner{}
	return &this
}

// GetTeamId returns the TeamId field value
func (o *ExternalGroupTeamsInner) GetTeamId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *ExternalGroupTeamsInner) GetTeamIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *ExternalGroupTeamsInner) SetTeamId(v int32) {
	o.TeamId = v
}

// GetTeamName returns the TeamName field value
func (o *ExternalGroupTeamsInner) GetTeamName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value
// and a boolean to check if the value has been set.
func (o *ExternalGroupTeamsInner) GetTeamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamName, true
}

// SetTeamName sets field value
func (o *ExternalGroupTeamsInner) SetTeamName(v string) {
	o.TeamName = v
}

func (o ExternalGroupTeamsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["team_id"] = o.TeamId
	}
	if true {
		toSerialize["team_name"] = o.TeamName
	}
	return json.Marshal(toSerialize)
}

type NullableExternalGroupTeamsInner struct {
	value *ExternalGroupTeamsInner
	isSet bool
}

func (v NullableExternalGroupTeamsInner) Get() *ExternalGroupTeamsInner {
	return v.value
}

func (v *NullableExternalGroupTeamsInner) Set(val *ExternalGroupTeamsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalGroupTeamsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalGroupTeamsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalGroupTeamsInner(val *ExternalGroupTeamsInner) *NullableExternalGroupTeamsInner {
	return &NullableExternalGroupTeamsInner{value: val, isSet: true}
}

func (v NullableExternalGroupTeamsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalGroupTeamsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


