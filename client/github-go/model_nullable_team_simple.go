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

// NullableTeamSimple Groups of organization members that gives permissions on specified repositories.
type NullableTeamSimple struct {
	// Unique identifier of the team
	Id int32 `json:"id"`
	NodeId string `json:"node_id"`
	// URL for the team
	Url string `json:"url"`
	MembersUrl string `json:"members_url"`
	// Name of the team
	Name string `json:"name"`
	// Description of the team
	Description NullableString `json:"description"`
	// Permission that the team will have for its repositories
	Permission string `json:"permission"`
	// The level of privacy this team should have
	Privacy *string `json:"privacy,omitempty"`
	HtmlUrl string `json:"html_url"`
	RepositoriesUrl string `json:"repositories_url"`
	Slug string `json:"slug"`
	// Distinguished Name (DN) that team maps to within LDAP environment
	LdapDn *string `json:"ldap_dn,omitempty"`
}

// NewNullableTeamSimple instantiates a new NullableTeamSimple object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNullableTeamSimple(id int32, nodeId string, url string, membersUrl string, name string, description NullableString, permission string, htmlUrl string, repositoriesUrl string, slug string) *NullableTeamSimple {
	this := NullableTeamSimple{}
	this.Id = id
	this.NodeId = nodeId
	this.Url = url
	this.MembersUrl = membersUrl
	this.Name = name
	this.Description = description
	this.Permission = permission
	this.HtmlUrl = htmlUrl
	this.RepositoriesUrl = repositoriesUrl
	this.Slug = slug
	return &this
}

// NewNullableTeamSimpleWithDefaults instantiates a new NullableTeamSimple object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNullableTeamSimpleWithDefaults() *NullableTeamSimple {
	this := NullableTeamSimple{}
	return &this
}

// GetId returns the Id field value
func (o *NullableTeamSimple) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NullableTeamSimple) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NullableTeamSimple) SetId(v int32) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *NullableTeamSimple) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *NullableTeamSimple) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *NullableTeamSimple) SetNodeId(v string) {
	o.NodeId = v
}

// GetUrl returns the Url field value
func (o *NullableTeamSimple) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NullableTeamSimple) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NullableTeamSimple) SetUrl(v string) {
	o.Url = v
}

// GetMembersUrl returns the MembersUrl field value
func (o *NullableTeamSimple) GetMembersUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MembersUrl
}

// GetMembersUrlOk returns a tuple with the MembersUrl field value
// and a boolean to check if the value has been set.
func (o *NullableTeamSimple) GetMembersUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MembersUrl, true
}

// SetMembersUrl sets field value
func (o *NullableTeamSimple) SetMembersUrl(v string) {
	o.MembersUrl = v
}

// GetName returns the Name field value
func (o *NullableTeamSimple) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NullableTeamSimple) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NullableTeamSimple) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NullableTeamSimple) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NullableTeamSimple) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *NullableTeamSimple) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetPermission returns the Permission field value
func (o *NullableTeamSimple) GetPermission() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value
// and a boolean to check if the value has been set.
func (o *NullableTeamSimple) GetPermissionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permission, true
}

// SetPermission sets field value
func (o *NullableTeamSimple) SetPermission(v string) {
	o.Permission = v
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *NullableTeamSimple) GetPrivacy() string {
	if o == nil || o.Privacy == nil {
		var ret string
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullableTeamSimple) GetPrivacyOk() (*string, bool) {
	if o == nil || o.Privacy == nil {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *NullableTeamSimple) HasPrivacy() bool {
	if o != nil && o.Privacy != nil {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given string and assigns it to the Privacy field.
func (o *NullableTeamSimple) SetPrivacy(v string) {
	o.Privacy = &v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *NullableTeamSimple) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *NullableTeamSimple) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *NullableTeamSimple) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetRepositoriesUrl returns the RepositoriesUrl field value
func (o *NullableTeamSimple) GetRepositoriesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepositoriesUrl
}

// GetRepositoriesUrlOk returns a tuple with the RepositoriesUrl field value
// and a boolean to check if the value has been set.
func (o *NullableTeamSimple) GetRepositoriesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoriesUrl, true
}

// SetRepositoriesUrl sets field value
func (o *NullableTeamSimple) SetRepositoriesUrl(v string) {
	o.RepositoriesUrl = v
}

// GetSlug returns the Slug field value
func (o *NullableTeamSimple) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *NullableTeamSimple) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *NullableTeamSimple) SetSlug(v string) {
	o.Slug = v
}

// GetLdapDn returns the LdapDn field value if set, zero value otherwise.
func (o *NullableTeamSimple) GetLdapDn() string {
	if o == nil || o.LdapDn == nil {
		var ret string
		return ret
	}
	return *o.LdapDn
}

// GetLdapDnOk returns a tuple with the LdapDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NullableTeamSimple) GetLdapDnOk() (*string, bool) {
	if o == nil || o.LdapDn == nil {
		return nil, false
	}
	return o.LdapDn, true
}

// HasLdapDn returns a boolean if a field has been set.
func (o *NullableTeamSimple) HasLdapDn() bool {
	if o != nil && o.LdapDn != nil {
		return true
	}

	return false
}

// SetLdapDn gets a reference to the given string and assigns it to the LdapDn field.
func (o *NullableTeamSimple) SetLdapDn(v string) {
	o.LdapDn = &v
}

func (o NullableTeamSimple) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["members_url"] = o.MembersUrl
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["permission"] = o.Permission
	}
	if o.Privacy != nil {
		toSerialize["privacy"] = o.Privacy
	}
	if true {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if true {
		toSerialize["repositories_url"] = o.RepositoriesUrl
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if o.LdapDn != nil {
		toSerialize["ldap_dn"] = o.LdapDn
	}
	return json.Marshal(toSerialize)
}

type NullableNullableTeamSimple struct {
	value *NullableTeamSimple
	isSet bool
}

func (v NullableNullableTeamSimple) Get() *NullableTeamSimple {
	return v.value
}

func (v *NullableNullableTeamSimple) Set(val *NullableTeamSimple) {
	v.value = val
	v.isSet = true
}

func (v NullableNullableTeamSimple) IsSet() bool {
	return v.isSet
}

func (v *NullableNullableTeamSimple) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNullableTeamSimple(val *NullableTeamSimple) *NullableNullableTeamSimple {
	return &NullableNullableTeamSimple{value: val, isSet: true}
}

func (v NullableNullableTeamSimple) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNullableTeamSimple) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


