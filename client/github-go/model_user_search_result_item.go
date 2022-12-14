/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"encoding/json"
	"time"
)

// UserSearchResultItem User Search Result Item
type UserSearchResultItem struct {
	Login string `json:"login"`
	Id int32 `json:"id"`
	NodeId string `json:"node_id"`
	AvatarUrl string `json:"avatar_url"`
	GravatarId NullableString `json:"gravatar_id"`
	Url string `json:"url"`
	HtmlUrl string `json:"html_url"`
	FollowersUrl string `json:"followers_url"`
	SubscriptionsUrl string `json:"subscriptions_url"`
	OrganizationsUrl string `json:"organizations_url"`
	ReposUrl string `json:"repos_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type string `json:"type"`
	Score float32 `json:"score"`
	FollowingUrl string `json:"following_url"`
	GistsUrl string `json:"gists_url"`
	StarredUrl string `json:"starred_url"`
	EventsUrl string `json:"events_url"`
	PublicRepos *int32 `json:"public_repos,omitempty"`
	PublicGists *int32 `json:"public_gists,omitempty"`
	Followers *int32 `json:"followers,omitempty"`
	Following *int32 `json:"following,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Bio NullableString `json:"bio,omitempty"`
	Email NullableString `json:"email,omitempty"`
	Location NullableString `json:"location,omitempty"`
	SiteAdmin bool `json:"site_admin"`
	Hireable NullableBool `json:"hireable,omitempty"`
	TextMatches []SearchResultTextMatchesInner `json:"text_matches,omitempty"`
	Blog NullableString `json:"blog,omitempty"`
	Company NullableString `json:"company,omitempty"`
	SuspendedAt NullableTime `json:"suspended_at,omitempty"`
}

// NewUserSearchResultItem instantiates a new UserSearchResultItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSearchResultItem(login string, id int32, nodeId string, avatarUrl string, gravatarId NullableString, url string, htmlUrl string, followersUrl string, subscriptionsUrl string, organizationsUrl string, reposUrl string, receivedEventsUrl string, type_ string, score float32, followingUrl string, gistsUrl string, starredUrl string, eventsUrl string, siteAdmin bool) *UserSearchResultItem {
	this := UserSearchResultItem{}
	this.Login = login
	this.Id = id
	this.NodeId = nodeId
	this.AvatarUrl = avatarUrl
	this.GravatarId = gravatarId
	this.Url = url
	this.HtmlUrl = htmlUrl
	this.FollowersUrl = followersUrl
	this.SubscriptionsUrl = subscriptionsUrl
	this.OrganizationsUrl = organizationsUrl
	this.ReposUrl = reposUrl
	this.ReceivedEventsUrl = receivedEventsUrl
	this.Type = type_
	this.Score = score
	this.FollowingUrl = followingUrl
	this.GistsUrl = gistsUrl
	this.StarredUrl = starredUrl
	this.EventsUrl = eventsUrl
	this.SiteAdmin = siteAdmin
	return &this
}

// NewUserSearchResultItemWithDefaults instantiates a new UserSearchResultItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSearchResultItemWithDefaults() *UserSearchResultItem {
	this := UserSearchResultItem{}
	return &this
}

// GetLogin returns the Login field value
func (o *UserSearchResultItem) GetLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *UserSearchResultItem) SetLogin(v string) {
	o.Login = v
}

// GetId returns the Id field value
func (o *UserSearchResultItem) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserSearchResultItem) SetId(v int32) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *UserSearchResultItem) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *UserSearchResultItem) SetNodeId(v string) {
	o.NodeId = v
}

// GetAvatarUrl returns the AvatarUrl field value
func (o *UserSearchResultItem) GetAvatarUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetAvatarUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvatarUrl, true
}

// SetAvatarUrl sets field value
func (o *UserSearchResultItem) SetAvatarUrl(v string) {
	o.AvatarUrl = v
}

// GetGravatarId returns the GravatarId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserSearchResultItem) GetGravatarId() string {
	if o == nil || o.GravatarId.Get() == nil {
		var ret string
		return ret
	}

	return *o.GravatarId.Get()
}

// GetGravatarIdOk returns a tuple with the GravatarId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSearchResultItem) GetGravatarIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GravatarId.Get(), o.GravatarId.IsSet()
}

// SetGravatarId sets field value
func (o *UserSearchResultItem) SetGravatarId(v string) {
	o.GravatarId.Set(&v)
}

// GetUrl returns the Url field value
func (o *UserSearchResultItem) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *UserSearchResultItem) SetUrl(v string) {
	o.Url = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *UserSearchResultItem) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *UserSearchResultItem) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetFollowersUrl returns the FollowersUrl field value
func (o *UserSearchResultItem) GetFollowersUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FollowersUrl
}

// GetFollowersUrlOk returns a tuple with the FollowersUrl field value
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetFollowersUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowersUrl, true
}

// SetFollowersUrl sets field value
func (o *UserSearchResultItem) SetFollowersUrl(v string) {
	o.FollowersUrl = v
}

// GetSubscriptionsUrl returns the SubscriptionsUrl field value
func (o *UserSearchResultItem) GetSubscriptionsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionsUrl
}

// GetSubscriptionsUrlOk returns a tuple with the SubscriptionsUrl field value
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetSubscriptionsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionsUrl, true
}

// SetSubscriptionsUrl sets field value
func (o *UserSearchResultItem) SetSubscriptionsUrl(v string) {
	o.SubscriptionsUrl = v
}

// GetOrganizationsUrl returns the OrganizationsUrl field value
func (o *UserSearchResultItem) GetOrganizationsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationsUrl
}

// GetOrganizationsUrlOk returns a tuple with the OrganizationsUrl field value
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetOrganizationsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationsUrl, true
}

// SetOrganizationsUrl sets field value
func (o *UserSearchResultItem) SetOrganizationsUrl(v string) {
	o.OrganizationsUrl = v
}

// GetReposUrl returns the ReposUrl field value
func (o *UserSearchResultItem) GetReposUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReposUrl
}

// GetReposUrlOk returns a tuple with the ReposUrl field value
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetReposUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReposUrl, true
}

// SetReposUrl sets field value
func (o *UserSearchResultItem) SetReposUrl(v string) {
	o.ReposUrl = v
}

// GetReceivedEventsUrl returns the ReceivedEventsUrl field value
func (o *UserSearchResultItem) GetReceivedEventsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceivedEventsUrl
}

// GetReceivedEventsUrlOk returns a tuple with the ReceivedEventsUrl field value
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetReceivedEventsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceivedEventsUrl, true
}

// SetReceivedEventsUrl sets field value
func (o *UserSearchResultItem) SetReceivedEventsUrl(v string) {
	o.ReceivedEventsUrl = v
}

// GetType returns the Type field value
func (o *UserSearchResultItem) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserSearchResultItem) SetType(v string) {
	o.Type = v
}

// GetScore returns the Score field value
func (o *UserSearchResultItem) GetScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *UserSearchResultItem) SetScore(v float32) {
	o.Score = v
}

// GetFollowingUrl returns the FollowingUrl field value
func (o *UserSearchResultItem) GetFollowingUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FollowingUrl
}

// GetFollowingUrlOk returns a tuple with the FollowingUrl field value
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetFollowingUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowingUrl, true
}

// SetFollowingUrl sets field value
func (o *UserSearchResultItem) SetFollowingUrl(v string) {
	o.FollowingUrl = v
}

// GetGistsUrl returns the GistsUrl field value
func (o *UserSearchResultItem) GetGistsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GistsUrl
}

// GetGistsUrlOk returns a tuple with the GistsUrl field value
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetGistsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GistsUrl, true
}

// SetGistsUrl sets field value
func (o *UserSearchResultItem) SetGistsUrl(v string) {
	o.GistsUrl = v
}

// GetStarredUrl returns the StarredUrl field value
func (o *UserSearchResultItem) GetStarredUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StarredUrl
}

// GetStarredUrlOk returns a tuple with the StarredUrl field value
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetStarredUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StarredUrl, true
}

// SetStarredUrl sets field value
func (o *UserSearchResultItem) SetStarredUrl(v string) {
	o.StarredUrl = v
}

// GetEventsUrl returns the EventsUrl field value
func (o *UserSearchResultItem) GetEventsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventsUrl
}

// GetEventsUrlOk returns a tuple with the EventsUrl field value
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetEventsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventsUrl, true
}

// SetEventsUrl sets field value
func (o *UserSearchResultItem) SetEventsUrl(v string) {
	o.EventsUrl = v
}

// GetPublicRepos returns the PublicRepos field value if set, zero value otherwise.
func (o *UserSearchResultItem) GetPublicRepos() int32 {
	if o == nil || o.PublicRepos == nil {
		var ret int32
		return ret
	}
	return *o.PublicRepos
}

// GetPublicReposOk returns a tuple with the PublicRepos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetPublicReposOk() (*int32, bool) {
	if o == nil || o.PublicRepos == nil {
		return nil, false
	}
	return o.PublicRepos, true
}

// HasPublicRepos returns a boolean if a field has been set.
func (o *UserSearchResultItem) HasPublicRepos() bool {
	if o != nil && o.PublicRepos != nil {
		return true
	}

	return false
}

// SetPublicRepos gets a reference to the given int32 and assigns it to the PublicRepos field.
func (o *UserSearchResultItem) SetPublicRepos(v int32) {
	o.PublicRepos = &v
}

// GetPublicGists returns the PublicGists field value if set, zero value otherwise.
func (o *UserSearchResultItem) GetPublicGists() int32 {
	if o == nil || o.PublicGists == nil {
		var ret int32
		return ret
	}
	return *o.PublicGists
}

// GetPublicGistsOk returns a tuple with the PublicGists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetPublicGistsOk() (*int32, bool) {
	if o == nil || o.PublicGists == nil {
		return nil, false
	}
	return o.PublicGists, true
}

// HasPublicGists returns a boolean if a field has been set.
func (o *UserSearchResultItem) HasPublicGists() bool {
	if o != nil && o.PublicGists != nil {
		return true
	}

	return false
}

// SetPublicGists gets a reference to the given int32 and assigns it to the PublicGists field.
func (o *UserSearchResultItem) SetPublicGists(v int32) {
	o.PublicGists = &v
}

// GetFollowers returns the Followers field value if set, zero value otherwise.
func (o *UserSearchResultItem) GetFollowers() int32 {
	if o == nil || o.Followers == nil {
		var ret int32
		return ret
	}
	return *o.Followers
}

// GetFollowersOk returns a tuple with the Followers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetFollowersOk() (*int32, bool) {
	if o == nil || o.Followers == nil {
		return nil, false
	}
	return o.Followers, true
}

// HasFollowers returns a boolean if a field has been set.
func (o *UserSearchResultItem) HasFollowers() bool {
	if o != nil && o.Followers != nil {
		return true
	}

	return false
}

// SetFollowers gets a reference to the given int32 and assigns it to the Followers field.
func (o *UserSearchResultItem) SetFollowers(v int32) {
	o.Followers = &v
}

// GetFollowing returns the Following field value if set, zero value otherwise.
func (o *UserSearchResultItem) GetFollowing() int32 {
	if o == nil || o.Following == nil {
		var ret int32
		return ret
	}
	return *o.Following
}

// GetFollowingOk returns a tuple with the Following field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetFollowingOk() (*int32, bool) {
	if o == nil || o.Following == nil {
		return nil, false
	}
	return o.Following, true
}

// HasFollowing returns a boolean if a field has been set.
func (o *UserSearchResultItem) HasFollowing() bool {
	if o != nil && o.Following != nil {
		return true
	}

	return false
}

// SetFollowing gets a reference to the given int32 and assigns it to the Following field.
func (o *UserSearchResultItem) SetFollowing(v int32) {
	o.Following = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UserSearchResultItem) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserSearchResultItem) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *UserSearchResultItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UserSearchResultItem) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UserSearchResultItem) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *UserSearchResultItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSearchResultItem) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSearchResultItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UserSearchResultItem) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UserSearchResultItem) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UserSearchResultItem) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UserSearchResultItem) UnsetName() {
	o.Name.Unset()
}

// GetBio returns the Bio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSearchResultItem) GetBio() string {
	if o == nil || o.Bio.Get() == nil {
		var ret string
		return ret
	}
	return *o.Bio.Get()
}

// GetBioOk returns a tuple with the Bio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSearchResultItem) GetBioOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bio.Get(), o.Bio.IsSet()
}

// HasBio returns a boolean if a field has been set.
func (o *UserSearchResultItem) HasBio() bool {
	if o != nil && o.Bio.IsSet() {
		return true
	}

	return false
}

// SetBio gets a reference to the given NullableString and assigns it to the Bio field.
func (o *UserSearchResultItem) SetBio(v string) {
	o.Bio.Set(&v)
}
// SetBioNil sets the value for Bio to be an explicit nil
func (o *UserSearchResultItem) SetBioNil() {
	o.Bio.Set(nil)
}

// UnsetBio ensures that no value is present for Bio, not even an explicit nil
func (o *UserSearchResultItem) UnsetBio() {
	o.Bio.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSearchResultItem) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSearchResultItem) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *UserSearchResultItem) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *UserSearchResultItem) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *UserSearchResultItem) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *UserSearchResultItem) UnsetEmail() {
	o.Email.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSearchResultItem) GetLocation() string {
	if o == nil || o.Location.Get() == nil {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSearchResultItem) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *UserSearchResultItem) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *UserSearchResultItem) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *UserSearchResultItem) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *UserSearchResultItem) UnsetLocation() {
	o.Location.Unset()
}

// GetSiteAdmin returns the SiteAdmin field value
func (o *UserSearchResultItem) GetSiteAdmin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SiteAdmin
}

// GetSiteAdminOk returns a tuple with the SiteAdmin field value
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetSiteAdminOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteAdmin, true
}

// SetSiteAdmin sets field value
func (o *UserSearchResultItem) SetSiteAdmin(v bool) {
	o.SiteAdmin = v
}

// GetHireable returns the Hireable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSearchResultItem) GetHireable() bool {
	if o == nil || o.Hireable.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Hireable.Get()
}

// GetHireableOk returns a tuple with the Hireable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSearchResultItem) GetHireableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hireable.Get(), o.Hireable.IsSet()
}

// HasHireable returns a boolean if a field has been set.
func (o *UserSearchResultItem) HasHireable() bool {
	if o != nil && o.Hireable.IsSet() {
		return true
	}

	return false
}

// SetHireable gets a reference to the given NullableBool and assigns it to the Hireable field.
func (o *UserSearchResultItem) SetHireable(v bool) {
	o.Hireable.Set(&v)
}
// SetHireableNil sets the value for Hireable to be an explicit nil
func (o *UserSearchResultItem) SetHireableNil() {
	o.Hireable.Set(nil)
}

// UnsetHireable ensures that no value is present for Hireable, not even an explicit nil
func (o *UserSearchResultItem) UnsetHireable() {
	o.Hireable.Unset()
}

// GetTextMatches returns the TextMatches field value if set, zero value otherwise.
func (o *UserSearchResultItem) GetTextMatches() []SearchResultTextMatchesInner {
	if o == nil || o.TextMatches == nil {
		var ret []SearchResultTextMatchesInner
		return ret
	}
	return o.TextMatches
}

// GetTextMatchesOk returns a tuple with the TextMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSearchResultItem) GetTextMatchesOk() ([]SearchResultTextMatchesInner, bool) {
	if o == nil || o.TextMatches == nil {
		return nil, false
	}
	return o.TextMatches, true
}

// HasTextMatches returns a boolean if a field has been set.
func (o *UserSearchResultItem) HasTextMatches() bool {
	if o != nil && o.TextMatches != nil {
		return true
	}

	return false
}

// SetTextMatches gets a reference to the given []SearchResultTextMatchesInner and assigns it to the TextMatches field.
func (o *UserSearchResultItem) SetTextMatches(v []SearchResultTextMatchesInner) {
	o.TextMatches = v
}

// GetBlog returns the Blog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSearchResultItem) GetBlog() string {
	if o == nil || o.Blog.Get() == nil {
		var ret string
		return ret
	}
	return *o.Blog.Get()
}

// GetBlogOk returns a tuple with the Blog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSearchResultItem) GetBlogOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Blog.Get(), o.Blog.IsSet()
}

// HasBlog returns a boolean if a field has been set.
func (o *UserSearchResultItem) HasBlog() bool {
	if o != nil && o.Blog.IsSet() {
		return true
	}

	return false
}

// SetBlog gets a reference to the given NullableString and assigns it to the Blog field.
func (o *UserSearchResultItem) SetBlog(v string) {
	o.Blog.Set(&v)
}
// SetBlogNil sets the value for Blog to be an explicit nil
func (o *UserSearchResultItem) SetBlogNil() {
	o.Blog.Set(nil)
}

// UnsetBlog ensures that no value is present for Blog, not even an explicit nil
func (o *UserSearchResultItem) UnsetBlog() {
	o.Blog.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSearchResultItem) GetCompany() string {
	if o == nil || o.Company.Get() == nil {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSearchResultItem) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *UserSearchResultItem) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *UserSearchResultItem) SetCompany(v string) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *UserSearchResultItem) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *UserSearchResultItem) UnsetCompany() {
	o.Company.Unset()
}

// GetSuspendedAt returns the SuspendedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserSearchResultItem) GetSuspendedAt() time.Time {
	if o == nil || o.SuspendedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SuspendedAt.Get()
}

// GetSuspendedAtOk returns a tuple with the SuspendedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserSearchResultItem) GetSuspendedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SuspendedAt.Get(), o.SuspendedAt.IsSet()
}

// HasSuspendedAt returns a boolean if a field has been set.
func (o *UserSearchResultItem) HasSuspendedAt() bool {
	if o != nil && o.SuspendedAt.IsSet() {
		return true
	}

	return false
}

// SetSuspendedAt gets a reference to the given NullableTime and assigns it to the SuspendedAt field.
func (o *UserSearchResultItem) SetSuspendedAt(v time.Time) {
	o.SuspendedAt.Set(&v)
}
// SetSuspendedAtNil sets the value for SuspendedAt to be an explicit nil
func (o *UserSearchResultItem) SetSuspendedAtNil() {
	o.SuspendedAt.Set(nil)
}

// UnsetSuspendedAt ensures that no value is present for SuspendedAt, not even an explicit nil
func (o *UserSearchResultItem) UnsetSuspendedAt() {
	o.SuspendedAt.Unset()
}

func (o UserSearchResultItem) MarshalJSON() ([]byte, error) {
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
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if true {
		toSerialize["gravatar_id"] = o.GravatarId.Get()
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if true {
		toSerialize["followers_url"] = o.FollowersUrl
	}
	if true {
		toSerialize["subscriptions_url"] = o.SubscriptionsUrl
	}
	if true {
		toSerialize["organizations_url"] = o.OrganizationsUrl
	}
	if true {
		toSerialize["repos_url"] = o.ReposUrl
	}
	if true {
		toSerialize["received_events_url"] = o.ReceivedEventsUrl
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["score"] = o.Score
	}
	if true {
		toSerialize["following_url"] = o.FollowingUrl
	}
	if true {
		toSerialize["gists_url"] = o.GistsUrl
	}
	if true {
		toSerialize["starred_url"] = o.StarredUrl
	}
	if true {
		toSerialize["events_url"] = o.EventsUrl
	}
	if o.PublicRepos != nil {
		toSerialize["public_repos"] = o.PublicRepos
	}
	if o.PublicGists != nil {
		toSerialize["public_gists"] = o.PublicGists
	}
	if o.Followers != nil {
		toSerialize["followers"] = o.Followers
	}
	if o.Following != nil {
		toSerialize["following"] = o.Following
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Bio.IsSet() {
		toSerialize["bio"] = o.Bio.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if true {
		toSerialize["site_admin"] = o.SiteAdmin
	}
	if o.Hireable.IsSet() {
		toSerialize["hireable"] = o.Hireable.Get()
	}
	if o.TextMatches != nil {
		toSerialize["text_matches"] = o.TextMatches
	}
	if o.Blog.IsSet() {
		toSerialize["blog"] = o.Blog.Get()
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if o.SuspendedAt.IsSet() {
		toSerialize["suspended_at"] = o.SuspendedAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableUserSearchResultItem struct {
	value *UserSearchResultItem
	isSet bool
}

func (v NullableUserSearchResultItem) Get() *UserSearchResultItem {
	return v.value
}

func (v *NullableUserSearchResultItem) Set(val *UserSearchResultItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSearchResultItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSearchResultItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSearchResultItem(val *UserSearchResultItem) *NullableUserSearchResultItem {
	return &NullableUserSearchResultItem{value: val, isSet: true}
}

func (v NullableUserSearchResultItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSearchResultItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


