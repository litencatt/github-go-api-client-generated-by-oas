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

// PullRequestHeadRepoOwner struct for PullRequestHeadRepoOwner
type PullRequestHeadRepoOwner struct {
	AvatarUrl string `json:"avatar_url"`
	EventsUrl string `json:"events_url"`
	FollowersUrl string `json:"followers_url"`
	FollowingUrl string `json:"following_url"`
	GistsUrl string `json:"gists_url"`
	GravatarId NullableString `json:"gravatar_id"`
	HtmlUrl string `json:"html_url"`
	Id int32 `json:"id"`
	NodeId string `json:"node_id"`
	Login string `json:"login"`
	OrganizationsUrl string `json:"organizations_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	ReposUrl string `json:"repos_url"`
	SiteAdmin bool `json:"site_admin"`
	StarredUrl string `json:"starred_url"`
	SubscriptionsUrl string `json:"subscriptions_url"`
	Type string `json:"type"`
	Url string `json:"url"`
}

// NewPullRequestHeadRepoOwner instantiates a new PullRequestHeadRepoOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPullRequestHeadRepoOwner(avatarUrl string, eventsUrl string, followersUrl string, followingUrl string, gistsUrl string, gravatarId NullableString, htmlUrl string, id int32, nodeId string, login string, organizationsUrl string, receivedEventsUrl string, reposUrl string, siteAdmin bool, starredUrl string, subscriptionsUrl string, type_ string, url string) *PullRequestHeadRepoOwner {
	this := PullRequestHeadRepoOwner{}
	this.AvatarUrl = avatarUrl
	this.EventsUrl = eventsUrl
	this.FollowersUrl = followersUrl
	this.FollowingUrl = followingUrl
	this.GistsUrl = gistsUrl
	this.GravatarId = gravatarId
	this.HtmlUrl = htmlUrl
	this.Id = id
	this.NodeId = nodeId
	this.Login = login
	this.OrganizationsUrl = organizationsUrl
	this.ReceivedEventsUrl = receivedEventsUrl
	this.ReposUrl = reposUrl
	this.SiteAdmin = siteAdmin
	this.StarredUrl = starredUrl
	this.SubscriptionsUrl = subscriptionsUrl
	this.Type = type_
	this.Url = url
	return &this
}

// NewPullRequestHeadRepoOwnerWithDefaults instantiates a new PullRequestHeadRepoOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPullRequestHeadRepoOwnerWithDefaults() *PullRequestHeadRepoOwner {
	this := PullRequestHeadRepoOwner{}
	return &this
}

// GetAvatarUrl returns the AvatarUrl field value
func (o *PullRequestHeadRepoOwner) GetAvatarUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestHeadRepoOwner) GetAvatarUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvatarUrl, true
}

// SetAvatarUrl sets field value
func (o *PullRequestHeadRepoOwner) SetAvatarUrl(v string) {
	o.AvatarUrl = v
}

// GetEventsUrl returns the EventsUrl field value
func (o *PullRequestHeadRepoOwner) GetEventsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventsUrl
}

// GetEventsUrlOk returns a tuple with the EventsUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestHeadRepoOwner) GetEventsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventsUrl, true
}

// SetEventsUrl sets field value
func (o *PullRequestHeadRepoOwner) SetEventsUrl(v string) {
	o.EventsUrl = v
}

// GetFollowersUrl returns the FollowersUrl field value
func (o *PullRequestHeadRepoOwner) GetFollowersUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FollowersUrl
}

// GetFollowersUrlOk returns a tuple with the FollowersUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestHeadRepoOwner) GetFollowersUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowersUrl, true
}

// SetFollowersUrl sets field value
func (o *PullRequestHeadRepoOwner) SetFollowersUrl(v string) {
	o.FollowersUrl = v
}

// GetFollowingUrl returns the FollowingUrl field value
func (o *PullRequestHeadRepoOwner) GetFollowingUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FollowingUrl
}

// GetFollowingUrlOk returns a tuple with the FollowingUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestHeadRepoOwner) GetFollowingUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowingUrl, true
}

// SetFollowingUrl sets field value
func (o *PullRequestHeadRepoOwner) SetFollowingUrl(v string) {
	o.FollowingUrl = v
}

// GetGistsUrl returns the GistsUrl field value
func (o *PullRequestHeadRepoOwner) GetGistsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GistsUrl
}

// GetGistsUrlOk returns a tuple with the GistsUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestHeadRepoOwner) GetGistsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GistsUrl, true
}

// SetGistsUrl sets field value
func (o *PullRequestHeadRepoOwner) SetGistsUrl(v string) {
	o.GistsUrl = v
}

// GetGravatarId returns the GravatarId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PullRequestHeadRepoOwner) GetGravatarId() string {
	if o == nil || o.GravatarId.Get() == nil {
		var ret string
		return ret
	}

	return *o.GravatarId.Get()
}

// GetGravatarIdOk returns a tuple with the GravatarId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequestHeadRepoOwner) GetGravatarIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GravatarId.Get(), o.GravatarId.IsSet()
}

// SetGravatarId sets field value
func (o *PullRequestHeadRepoOwner) SetGravatarId(v string) {
	o.GravatarId.Set(&v)
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *PullRequestHeadRepoOwner) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestHeadRepoOwner) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *PullRequestHeadRepoOwner) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetId returns the Id field value
func (o *PullRequestHeadRepoOwner) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PullRequestHeadRepoOwner) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PullRequestHeadRepoOwner) SetId(v int32) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *PullRequestHeadRepoOwner) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *PullRequestHeadRepoOwner) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *PullRequestHeadRepoOwner) SetNodeId(v string) {
	o.NodeId = v
}

// GetLogin returns the Login field value
func (o *PullRequestHeadRepoOwner) GetLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
func (o *PullRequestHeadRepoOwner) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *PullRequestHeadRepoOwner) SetLogin(v string) {
	o.Login = v
}

// GetOrganizationsUrl returns the OrganizationsUrl field value
func (o *PullRequestHeadRepoOwner) GetOrganizationsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationsUrl
}

// GetOrganizationsUrlOk returns a tuple with the OrganizationsUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestHeadRepoOwner) GetOrganizationsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationsUrl, true
}

// SetOrganizationsUrl sets field value
func (o *PullRequestHeadRepoOwner) SetOrganizationsUrl(v string) {
	o.OrganizationsUrl = v
}

// GetReceivedEventsUrl returns the ReceivedEventsUrl field value
func (o *PullRequestHeadRepoOwner) GetReceivedEventsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceivedEventsUrl
}

// GetReceivedEventsUrlOk returns a tuple with the ReceivedEventsUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestHeadRepoOwner) GetReceivedEventsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceivedEventsUrl, true
}

// SetReceivedEventsUrl sets field value
func (o *PullRequestHeadRepoOwner) SetReceivedEventsUrl(v string) {
	o.ReceivedEventsUrl = v
}

// GetReposUrl returns the ReposUrl field value
func (o *PullRequestHeadRepoOwner) GetReposUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReposUrl
}

// GetReposUrlOk returns a tuple with the ReposUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestHeadRepoOwner) GetReposUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReposUrl, true
}

// SetReposUrl sets field value
func (o *PullRequestHeadRepoOwner) SetReposUrl(v string) {
	o.ReposUrl = v
}

// GetSiteAdmin returns the SiteAdmin field value
func (o *PullRequestHeadRepoOwner) GetSiteAdmin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SiteAdmin
}

// GetSiteAdminOk returns a tuple with the SiteAdmin field value
// and a boolean to check if the value has been set.
func (o *PullRequestHeadRepoOwner) GetSiteAdminOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteAdmin, true
}

// SetSiteAdmin sets field value
func (o *PullRequestHeadRepoOwner) SetSiteAdmin(v bool) {
	o.SiteAdmin = v
}

// GetStarredUrl returns the StarredUrl field value
func (o *PullRequestHeadRepoOwner) GetStarredUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StarredUrl
}

// GetStarredUrlOk returns a tuple with the StarredUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestHeadRepoOwner) GetStarredUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StarredUrl, true
}

// SetStarredUrl sets field value
func (o *PullRequestHeadRepoOwner) SetStarredUrl(v string) {
	o.StarredUrl = v
}

// GetSubscriptionsUrl returns the SubscriptionsUrl field value
func (o *PullRequestHeadRepoOwner) GetSubscriptionsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionsUrl
}

// GetSubscriptionsUrlOk returns a tuple with the SubscriptionsUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestHeadRepoOwner) GetSubscriptionsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionsUrl, true
}

// SetSubscriptionsUrl sets field value
func (o *PullRequestHeadRepoOwner) SetSubscriptionsUrl(v string) {
	o.SubscriptionsUrl = v
}

// GetType returns the Type field value
func (o *PullRequestHeadRepoOwner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PullRequestHeadRepoOwner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PullRequestHeadRepoOwner) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *PullRequestHeadRepoOwner) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PullRequestHeadRepoOwner) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PullRequestHeadRepoOwner) SetUrl(v string) {
	o.Url = v
}

func (o PullRequestHeadRepoOwner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if true {
		toSerialize["events_url"] = o.EventsUrl
	}
	if true {
		toSerialize["followers_url"] = o.FollowersUrl
	}
	if true {
		toSerialize["following_url"] = o.FollowingUrl
	}
	if true {
		toSerialize["gists_url"] = o.GistsUrl
	}
	if true {
		toSerialize["gravatar_id"] = o.GravatarId.Get()
	}
	if true {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	if true {
		toSerialize["login"] = o.Login
	}
	if true {
		toSerialize["organizations_url"] = o.OrganizationsUrl
	}
	if true {
		toSerialize["received_events_url"] = o.ReceivedEventsUrl
	}
	if true {
		toSerialize["repos_url"] = o.ReposUrl
	}
	if true {
		toSerialize["site_admin"] = o.SiteAdmin
	}
	if true {
		toSerialize["starred_url"] = o.StarredUrl
	}
	if true {
		toSerialize["subscriptions_url"] = o.SubscriptionsUrl
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullablePullRequestHeadRepoOwner struct {
	value *PullRequestHeadRepoOwner
	isSet bool
}

func (v NullablePullRequestHeadRepoOwner) Get() *PullRequestHeadRepoOwner {
	return v.value
}

func (v *NullablePullRequestHeadRepoOwner) Set(val *PullRequestHeadRepoOwner) {
	v.value = val
	v.isSet = true
}

func (v NullablePullRequestHeadRepoOwner) IsSet() bool {
	return v.isSet
}

func (v *NullablePullRequestHeadRepoOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullRequestHeadRepoOwner(val *PullRequestHeadRepoOwner) *NullablePullRequestHeadRepoOwner {
	return &NullablePullRequestHeadRepoOwner{value: val, isSet: true}
}

func (v NullablePullRequestHeadRepoOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullRequestHeadRepoOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

