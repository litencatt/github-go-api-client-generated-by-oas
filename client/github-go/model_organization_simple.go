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

// OrganizationSimple Organization Simple
type OrganizationSimple struct {
	Login string `json:"login"`
	Id int32 `json:"id"`
	NodeId string `json:"node_id"`
	Url string `json:"url"`
	ReposUrl string `json:"repos_url"`
	EventsUrl string `json:"events_url"`
	HooksUrl string `json:"hooks_url"`
	IssuesUrl string `json:"issues_url"`
	MembersUrl string `json:"members_url"`
	PublicMembersUrl string `json:"public_members_url"`
	AvatarUrl string `json:"avatar_url"`
	Description NullableString `json:"description"`
}

// NewOrganizationSimple instantiates a new OrganizationSimple object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationSimple(login string, id int32, nodeId string, url string, reposUrl string, eventsUrl string, hooksUrl string, issuesUrl string, membersUrl string, publicMembersUrl string, avatarUrl string, description NullableString) *OrganizationSimple {
	this := OrganizationSimple{}
	this.Login = login
	this.Id = id
	this.NodeId = nodeId
	this.Url = url
	this.ReposUrl = reposUrl
	this.EventsUrl = eventsUrl
	this.HooksUrl = hooksUrl
	this.IssuesUrl = issuesUrl
	this.MembersUrl = membersUrl
	this.PublicMembersUrl = publicMembersUrl
	this.AvatarUrl = avatarUrl
	this.Description = description
	return &this
}

// NewOrganizationSimpleWithDefaults instantiates a new OrganizationSimple object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationSimpleWithDefaults() *OrganizationSimple {
	this := OrganizationSimple{}
	return &this
}

// GetLogin returns the Login field value
func (o *OrganizationSimple) GetLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
func (o *OrganizationSimple) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *OrganizationSimple) SetLogin(v string) {
	o.Login = v
}

// GetId returns the Id field value
func (o *OrganizationSimple) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationSimple) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationSimple) SetId(v int32) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *OrganizationSimple) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *OrganizationSimple) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *OrganizationSimple) SetNodeId(v string) {
	o.NodeId = v
}

// GetUrl returns the Url field value
func (o *OrganizationSimple) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *OrganizationSimple) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *OrganizationSimple) SetUrl(v string) {
	o.Url = v
}

// GetReposUrl returns the ReposUrl field value
func (o *OrganizationSimple) GetReposUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReposUrl
}

// GetReposUrlOk returns a tuple with the ReposUrl field value
// and a boolean to check if the value has been set.
func (o *OrganizationSimple) GetReposUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReposUrl, true
}

// SetReposUrl sets field value
func (o *OrganizationSimple) SetReposUrl(v string) {
	o.ReposUrl = v
}

// GetEventsUrl returns the EventsUrl field value
func (o *OrganizationSimple) GetEventsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventsUrl
}

// GetEventsUrlOk returns a tuple with the EventsUrl field value
// and a boolean to check if the value has been set.
func (o *OrganizationSimple) GetEventsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventsUrl, true
}

// SetEventsUrl sets field value
func (o *OrganizationSimple) SetEventsUrl(v string) {
	o.EventsUrl = v
}

// GetHooksUrl returns the HooksUrl field value
func (o *OrganizationSimple) GetHooksUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HooksUrl
}

// GetHooksUrlOk returns a tuple with the HooksUrl field value
// and a boolean to check if the value has been set.
func (o *OrganizationSimple) GetHooksUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HooksUrl, true
}

// SetHooksUrl sets field value
func (o *OrganizationSimple) SetHooksUrl(v string) {
	o.HooksUrl = v
}

// GetIssuesUrl returns the IssuesUrl field value
func (o *OrganizationSimple) GetIssuesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuesUrl
}

// GetIssuesUrlOk returns a tuple with the IssuesUrl field value
// and a boolean to check if the value has been set.
func (o *OrganizationSimple) GetIssuesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuesUrl, true
}

// SetIssuesUrl sets field value
func (o *OrganizationSimple) SetIssuesUrl(v string) {
	o.IssuesUrl = v
}

// GetMembersUrl returns the MembersUrl field value
func (o *OrganizationSimple) GetMembersUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MembersUrl
}

// GetMembersUrlOk returns a tuple with the MembersUrl field value
// and a boolean to check if the value has been set.
func (o *OrganizationSimple) GetMembersUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MembersUrl, true
}

// SetMembersUrl sets field value
func (o *OrganizationSimple) SetMembersUrl(v string) {
	o.MembersUrl = v
}

// GetPublicMembersUrl returns the PublicMembersUrl field value
func (o *OrganizationSimple) GetPublicMembersUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicMembersUrl
}

// GetPublicMembersUrlOk returns a tuple with the PublicMembersUrl field value
// and a boolean to check if the value has been set.
func (o *OrganizationSimple) GetPublicMembersUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicMembersUrl, true
}

// SetPublicMembersUrl sets field value
func (o *OrganizationSimple) SetPublicMembersUrl(v string) {
	o.PublicMembersUrl = v
}

// GetAvatarUrl returns the AvatarUrl field value
func (o *OrganizationSimple) GetAvatarUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value
// and a boolean to check if the value has been set.
func (o *OrganizationSimple) GetAvatarUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvatarUrl, true
}

// SetAvatarUrl sets field value
func (o *OrganizationSimple) SetAvatarUrl(v string) {
	o.AvatarUrl = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OrganizationSimple) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationSimple) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *OrganizationSimple) SetDescription(v string) {
	o.Description.Set(&v)
}

func (o OrganizationSimple) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["login"] = o.Login
	}
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
		toSerialize["repos_url"] = o.ReposUrl
	}
	if true {
		toSerialize["events_url"] = o.EventsUrl
	}
	if true {
		toSerialize["hooks_url"] = o.HooksUrl
	}
	if true {
		toSerialize["issues_url"] = o.IssuesUrl
	}
	if true {
		toSerialize["members_url"] = o.MembersUrl
	}
	if true {
		toSerialize["public_members_url"] = o.PublicMembersUrl
	}
	if true {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if true {
		toSerialize["description"] = o.Description.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationSimple struct {
	value *OrganizationSimple
	isSet bool
}

func (v NullableOrganizationSimple) Get() *OrganizationSimple {
	return v.value
}

func (v *NullableOrganizationSimple) Set(val *OrganizationSimple) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationSimple) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationSimple) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationSimple(val *OrganizationSimple) *NullableOrganizationSimple {
	return &NullableOrganizationSimple{value: val, isSet: true}
}

func (v NullableOrganizationSimple) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationSimple) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


