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

// ExternalGroup Information about an external group's usage and its members
type ExternalGroup struct {
	// The internal ID of the group
	GroupId int32 `json:"group_id"`
	// The display name for the group
	GroupName string `json:"group_name"`
	// The date when the group was last updated_at
	UpdatedAt *string `json:"updated_at,omitempty"`
	// An array of teams linked to this group
	Teams []ExternalGroupTeamsInner `json:"teams"`
	// An array of external members linked to this group
	Members []ExternalGroupMembersInner `json:"members"`
}

// NewExternalGroup instantiates a new ExternalGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalGroup(groupId int32, groupName string, teams []ExternalGroupTeamsInner, members []ExternalGroupMembersInner) *ExternalGroup {
	this := ExternalGroup{}
	this.GroupId = groupId
	this.GroupName = groupName
	this.Teams = teams
	this.Members = members
	return &this
}

// NewExternalGroupWithDefaults instantiates a new ExternalGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalGroupWithDefaults() *ExternalGroup {
	this := ExternalGroup{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *ExternalGroup) GetGroupId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *ExternalGroup) GetGroupIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *ExternalGroup) SetGroupId(v int32) {
	o.GroupId = v
}

// GetGroupName returns the GroupName field value
func (o *ExternalGroup) GetGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value
// and a boolean to check if the value has been set.
func (o *ExternalGroup) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupName, true
}

// SetGroupName sets field value
func (o *ExternalGroup) SetGroupName(v string) {
	o.GroupName = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ExternalGroup) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalGroup) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ExternalGroup) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ExternalGroup) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetTeams returns the Teams field value
func (o *ExternalGroup) GetTeams() []ExternalGroupTeamsInner {
	if o == nil {
		var ret []ExternalGroupTeamsInner
		return ret
	}

	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value
// and a boolean to check if the value has been set.
func (o *ExternalGroup) GetTeamsOk() ([]ExternalGroupTeamsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Teams, true
}

// SetTeams sets field value
func (o *ExternalGroup) SetTeams(v []ExternalGroupTeamsInner) {
	o.Teams = v
}

// GetMembers returns the Members field value
func (o *ExternalGroup) GetMembers() []ExternalGroupMembersInner {
	if o == nil {
		var ret []ExternalGroupMembersInner
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *ExternalGroup) GetMembersOk() ([]ExternalGroupMembersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *ExternalGroup) SetMembers(v []ExternalGroupMembersInner) {
	o.Members = v
}

func (o ExternalGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["group_id"] = o.GroupId
	}
	if true {
		toSerialize["group_name"] = o.GroupName
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["teams"] = o.Teams
	}
	if true {
		toSerialize["members"] = o.Members
	}
	return json.Marshal(toSerialize)
}

type NullableExternalGroup struct {
	value *ExternalGroup
	isSet bool
}

func (v NullableExternalGroup) Get() *ExternalGroup {
	return v.value
}

func (v *NullableExternalGroup) Set(val *ExternalGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalGroup(val *ExternalGroup) *NullableExternalGroup {
	return &NullableExternalGroup{value: val, isSet: true}
}

func (v NullableExternalGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

