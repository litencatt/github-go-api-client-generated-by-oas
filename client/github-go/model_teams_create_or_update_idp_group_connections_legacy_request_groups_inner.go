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

// TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner struct for TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner
type TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner struct {
	// ID of the IdP group.
	GroupId string `json:"group_id"`
	// Name of the IdP group.
	GroupName string `json:"group_name"`
	// Description of the IdP group.
	GroupDescription string `json:"group_description"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner instantiates a new TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner(groupId string, groupName string, groupDescription string) *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner {
	this := TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner{}
	this.GroupId = groupId
	this.GroupName = groupName
	this.GroupDescription = groupDescription
	return &this
}

// NewTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInnerWithDefaults instantiates a new TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInnerWithDefaults() *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner {
	this := TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) SetGroupId(v string) {
	o.GroupId = v
}

// GetGroupName returns the GroupName field value
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) GetGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value
// and a boolean to check if the value has been set.
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupName, true
}

// SetGroupName sets field value
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) SetGroupName(v string) {
	o.GroupName = v
}

// GetGroupDescription returns the GroupDescription field value
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) GetGroupDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupDescription
}

// GetGroupDescriptionOk returns a tuple with the GroupDescription field value
// and a boolean to check if the value has been set.
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) GetGroupDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupDescription, true
}

// SetGroupDescription sets field value
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) SetGroupDescription(v string) {
	o.GroupDescription = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) SetDescription(v string) {
	o.Description = &v
}

func (o TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["group_id"] = o.GroupId
	}
	if true {
		toSerialize["group_name"] = o.GroupName
	}
	if true {
		toSerialize["group_description"] = o.GroupDescription
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner struct {
	value *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner
	isSet bool
}

func (v NullableTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) Get() *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner {
	return v.value
}

func (v *NullableTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) Set(val *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner(val *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) *NullableTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner {
	return &NullableTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner{value: val, isSet: true}
}

func (v NullableTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

