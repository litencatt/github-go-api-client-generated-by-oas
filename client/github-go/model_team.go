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

// Team Groups of organization members that gives permissions on specified repositories.
type Team struct {
	Id int32 `json:"id"`
	NodeId string `json:"node_id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Description NullableString `json:"description"`
	Privacy *string `json:"privacy,omitempty"`
	Permission string `json:"permission"`
	Permissions *TeamPermissions `json:"permissions,omitempty"`
	Url string `json:"url"`
	HtmlUrl string `json:"html_url"`
	MembersUrl string `json:"members_url"`
	RepositoriesUrl string `json:"repositories_url"`
	Parent NullableNullableTeamSimple `json:"parent"`
}

// NewTeam instantiates a new Team object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeam(id int32, nodeId string, name string, slug string, description NullableString, permission string, url string, htmlUrl string, membersUrl string, repositoriesUrl string, parent NullableNullableTeamSimple) *Team {
	this := Team{}
	this.Id = id
	this.NodeId = nodeId
	this.Name = name
	this.Slug = slug
	this.Description = description
	this.Permission = permission
	this.Url = url
	this.HtmlUrl = htmlUrl
	this.MembersUrl = membersUrl
	this.RepositoriesUrl = repositoriesUrl
	this.Parent = parent
	return &this
}

// NewTeamWithDefaults instantiates a new Team object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamWithDefaults() *Team {
	this := Team{}
	return &this
}

// GetId returns the Id field value
func (o *Team) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Team) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Team) SetId(v int32) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *Team) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *Team) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *Team) SetNodeId(v string) {
	o.NodeId = v
}

// GetName returns the Name field value
func (o *Team) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Team) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Team) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *Team) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Team) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Team) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Team) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Team) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *Team) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *Team) GetPrivacy() string {
	if o == nil || o.Privacy == nil {
		var ret string
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Team) GetPrivacyOk() (*string, bool) {
	if o == nil || o.Privacy == nil {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *Team) HasPrivacy() bool {
	if o != nil && o.Privacy != nil {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given string and assigns it to the Privacy field.
func (o *Team) SetPrivacy(v string) {
	o.Privacy = &v
}

// GetPermission returns the Permission field value
func (o *Team) GetPermission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *Team) GetPermissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *Team) SetPermission(v string) {
	o.Permission = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *Team) GetPermissions() TeamPermissions {
	if o == nil || o.Permissions == nil {
		var ret TeamPermissions
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Team) GetPermissionsOk() (*TeamPermissions, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *Team) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given TeamPermissions and assigns it to the Permissions field.
func (o *Team) SetPermissions(v TeamPermissions) {
	o.Permissions = &v
}

// GetUrl returns the Url field value
func (o *Team) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Team) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Team) SetUrl(v string) {
	o.Url = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *Team) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *Team) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *Team) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetMembersUrl returns the MembersUrl field value
func (o *Team) GetMembersUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MembersUrl
}

// GetMembersUrlOk returns a tuple with the MembersUrl field value
// and a boolean to check if the value has been set.
func (o *Team) GetMembersUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MembersUrl, true
}

// SetMembersUrl sets field value
func (o *Team) SetMembersUrl(v string) {
	o.MembersUrl = v
}

// GetRepositoriesUrl returns the RepositoriesUrl field value
func (o *Team) GetRepositoriesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepositoriesUrl
}

// GetRepositoriesUrlOk returns a tuple with the RepositoriesUrl field value
// and a boolean to check if the value has been set.
func (o *Team) GetRepositoriesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoriesUrl, true
}

// SetRepositoriesUrl sets field value
func (o *Team) SetRepositoriesUrl(v string) {
	o.RepositoriesUrl = v
}

// GetParent returns the Parent field value
// If the value is explicit nil, the zero value for NullableTeamSimple will be returned
func (o *Team) GetParent() NullableTeamSimple {
	if o == nil || o.Parent.Get() == nil {
		var ret NullableTeamSimple
		return ret
	}

	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Team) GetParentOk() (*NullableTeamSimple, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// SetParent sets field value
func (o *Team) SetParent(v NullableTeamSimple) {
	o.Parent.Set(&v)
}

func (o Team) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Privacy != nil {
		toSerialize["privacy"] = o.Privacy
	}
	if true {
		toSerialize["permission"] = o.Permission
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if true {
		toSerialize["members_url"] = o.MembersUrl
	}
	if true {
		toSerialize["repositories_url"] = o.RepositoriesUrl
	}
	if true {
		toSerialize["parent"] = o.Parent.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTeam struct {
	value *Team
	isSet bool
}

func (v NullableTeam) Get() *Team {
	return v.value
}

func (v *NullableTeam) Set(val *Team) {
	v.value = val
	v.isSet = true
}

func (v NullableTeam) IsSet() bool {
	return v.isSet
}

func (v *NullableTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeam(val *Team) *NullableTeam {
	return &NullableTeam{value: val, isSet: true}
}

func (v NullableTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


