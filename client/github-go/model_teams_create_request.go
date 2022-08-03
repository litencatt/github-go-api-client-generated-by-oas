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

// TeamsCreateRequest struct for TeamsCreateRequest
type TeamsCreateRequest struct {
	// The name of the team.
	Name string `json:"name"`
	// The description of the team.
	Description *string `json:"description,omitempty"`
	// List GitHub IDs for organization members who will become team maintainers.
	Maintainers []string `json:"maintainers,omitempty"`
	// The full name (e.g., \"organization-name/repository-name\") of repositories to add the team to.
	RepoNames []string `json:"repo_names,omitempty"`
	// The level of privacy this team should have. The options are:   **For a non-nested team:**   \\* `secret` - only visible to organization owners and members of this team.   \\* `closed` - visible to all members of this organization.   Default: `secret`   **For a parent or child team:**   \\* `closed` - visible to all members of this organization.   Default for child team: `closed`
	Privacy *string `json:"privacy,omitempty"`
	// **Deprecated**. The permission that new repositories will be added to the team with when none is specified.
	Permission *string `json:"permission,omitempty"`
	// The ID of a team to set as the parent team.
	ParentTeamId *int32 `json:"parent_team_id,omitempty"`
}

// NewTeamsCreateRequest instantiates a new TeamsCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamsCreateRequest(name string) *TeamsCreateRequest {
	this := TeamsCreateRequest{}
	this.Name = name
	var permission string = "pull"
	this.Permission = &permission
	return &this
}

// NewTeamsCreateRequestWithDefaults instantiates a new TeamsCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamsCreateRequestWithDefaults() *TeamsCreateRequest {
	this := TeamsCreateRequest{}
	var permission string = "pull"
	this.Permission = &permission
	return &this
}

// GetName returns the Name field value
func (o *TeamsCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TeamsCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TeamsCreateRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TeamsCreateRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TeamsCreateRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TeamsCreateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMaintainers returns the Maintainers field value if set, zero value otherwise.
func (o *TeamsCreateRequest) GetMaintainers() []string {
	if o == nil || o.Maintainers == nil {
		var ret []string
		return ret
	}
	return o.Maintainers
}

// GetMaintainersOk returns a tuple with the Maintainers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsCreateRequest) GetMaintainersOk() ([]string, bool) {
	if o == nil || o.Maintainers == nil {
		return nil, false
	}
	return o.Maintainers, true
}

// HasMaintainers returns a boolean if a field has been set.
func (o *TeamsCreateRequest) HasMaintainers() bool {
	if o != nil && o.Maintainers != nil {
		return true
	}

	return false
}

// SetMaintainers gets a reference to the given []string and assigns it to the Maintainers field.
func (o *TeamsCreateRequest) SetMaintainers(v []string) {
	o.Maintainers = v
}

// GetRepoNames returns the RepoNames field value if set, zero value otherwise.
func (o *TeamsCreateRequest) GetRepoNames() []string {
	if o == nil || o.RepoNames == nil {
		var ret []string
		return ret
	}
	return o.RepoNames
}

// GetRepoNamesOk returns a tuple with the RepoNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsCreateRequest) GetRepoNamesOk() ([]string, bool) {
	if o == nil || o.RepoNames == nil {
		return nil, false
	}
	return o.RepoNames, true
}

// HasRepoNames returns a boolean if a field has been set.
func (o *TeamsCreateRequest) HasRepoNames() bool {
	if o != nil && o.RepoNames != nil {
		return true
	}

	return false
}

// SetRepoNames gets a reference to the given []string and assigns it to the RepoNames field.
func (o *TeamsCreateRequest) SetRepoNames(v []string) {
	o.RepoNames = v
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *TeamsCreateRequest) GetPrivacy() string {
	if o == nil || o.Privacy == nil {
		var ret string
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsCreateRequest) GetPrivacyOk() (*string, bool) {
	if o == nil || o.Privacy == nil {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *TeamsCreateRequest) HasPrivacy() bool {
	if o != nil && o.Privacy != nil {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given string and assigns it to the Privacy field.
func (o *TeamsCreateRequest) SetPrivacy(v string) {
	o.Privacy = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *TeamsCreateRequest) GetPermission() string {
	if o == nil || o.Permission == nil {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsCreateRequest) GetPermissionOk() (*string, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *TeamsCreateRequest) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *TeamsCreateRequest) SetPermission(v string) {
	o.Permission = &v
}

// GetParentTeamId returns the ParentTeamId field value if set, zero value otherwise.
func (o *TeamsCreateRequest) GetParentTeamId() int32 {
	if o == nil || o.ParentTeamId == nil {
		var ret int32
		return ret
	}
	return *o.ParentTeamId
}

// GetParentTeamIdOk returns a tuple with the ParentTeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamsCreateRequest) GetParentTeamIdOk() (*int32, bool) {
	if o == nil || o.ParentTeamId == nil {
		return nil, false
	}
	return o.ParentTeamId, true
}

// HasParentTeamId returns a boolean if a field has been set.
func (o *TeamsCreateRequest) HasParentTeamId() bool {
	if o != nil && o.ParentTeamId != nil {
		return true
	}

	return false
}

// SetParentTeamId gets a reference to the given int32 and assigns it to the ParentTeamId field.
func (o *TeamsCreateRequest) SetParentTeamId(v int32) {
	o.ParentTeamId = &v
}

func (o TeamsCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Maintainers != nil {
		toSerialize["maintainers"] = o.Maintainers
	}
	if o.RepoNames != nil {
		toSerialize["repo_names"] = o.RepoNames
	}
	if o.Privacy != nil {
		toSerialize["privacy"] = o.Privacy
	}
	if o.Permission != nil {
		toSerialize["permission"] = o.Permission
	}
	if o.ParentTeamId != nil {
		toSerialize["parent_team_id"] = o.ParentTeamId
	}
	return json.Marshal(toSerialize)
}

type NullableTeamsCreateRequest struct {
	value *TeamsCreateRequest
	isSet bool
}

func (v NullableTeamsCreateRequest) Get() *TeamsCreateRequest {
	return v.value
}

func (v *NullableTeamsCreateRequest) Set(val *TeamsCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamsCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamsCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamsCreateRequest(val *TeamsCreateRequest) *NullableTeamsCreateRequest {
	return &NullableTeamsCreateRequest{value: val, isSet: true}
}

func (v NullableTeamsCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamsCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


