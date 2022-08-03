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

// PrivateUser Private User
type PrivateUser struct {
	Login string `json:"login"`
	Id int32 `json:"id"`
	NodeId string `json:"node_id"`
	AvatarUrl string `json:"avatar_url"`
	GravatarId NullableString `json:"gravatar_id"`
	Url string `json:"url"`
	HtmlUrl string `json:"html_url"`
	FollowersUrl string `json:"followers_url"`
	FollowingUrl string `json:"following_url"`
	GistsUrl string `json:"gists_url"`
	StarredUrl string `json:"starred_url"`
	SubscriptionsUrl string `json:"subscriptions_url"`
	OrganizationsUrl string `json:"organizations_url"`
	ReposUrl string `json:"repos_url"`
	EventsUrl string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type string `json:"type"`
	SiteAdmin bool `json:"site_admin"`
	Name NullableString `json:"name"`
	Company NullableString `json:"company"`
	Blog NullableString `json:"blog"`
	Location NullableString `json:"location"`
	Email NullableString `json:"email"`
	Hireable NullableBool `json:"hireable"`
	Bio NullableString `json:"bio"`
	TwitterUsername NullableString `json:"twitter_username,omitempty"`
	PublicRepos int32 `json:"public_repos"`
	PublicGists int32 `json:"public_gists"`
	Followers int32 `json:"followers"`
	Following int32 `json:"following"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	PrivateGists int32 `json:"private_gists"`
	TotalPrivateRepos int32 `json:"total_private_repos"`
	OwnedPrivateRepos int32 `json:"owned_private_repos"`
	DiskUsage int32 `json:"disk_usage"`
	Collaborators int32 `json:"collaborators"`
	TwoFactorAuthentication bool `json:"two_factor_authentication"`
	Plan *PublicUserPlan `json:"plan,omitempty"`
	SuspendedAt NullableTime `json:"suspended_at,omitempty"`
	BusinessPlus *bool `json:"business_plus,omitempty"`
	LdapDn *string `json:"ldap_dn,omitempty"`
}

// NewPrivateUser instantiates a new PrivateUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateUser(login string, id int32, nodeId string, avatarUrl string, gravatarId NullableString, url string, htmlUrl string, followersUrl string, followingUrl string, gistsUrl string, starredUrl string, subscriptionsUrl string, organizationsUrl string, reposUrl string, eventsUrl string, receivedEventsUrl string, type_ string, siteAdmin bool, name NullableString, company NullableString, blog NullableString, location NullableString, email NullableString, hireable NullableBool, bio NullableString, publicRepos int32, publicGists int32, followers int32, following int32, createdAt time.Time, updatedAt time.Time, privateGists int32, totalPrivateRepos int32, ownedPrivateRepos int32, diskUsage int32, collaborators int32, twoFactorAuthentication bool) *PrivateUser {
	this := PrivateUser{}
	this.Login = login
	this.Id = id
	this.NodeId = nodeId
	this.AvatarUrl = avatarUrl
	this.GravatarId = gravatarId
	this.Url = url
	this.HtmlUrl = htmlUrl
	this.FollowersUrl = followersUrl
	this.FollowingUrl = followingUrl
	this.GistsUrl = gistsUrl
	this.StarredUrl = starredUrl
	this.SubscriptionsUrl = subscriptionsUrl
	this.OrganizationsUrl = organizationsUrl
	this.ReposUrl = reposUrl
	this.EventsUrl = eventsUrl
	this.ReceivedEventsUrl = receivedEventsUrl
	this.Type = type_
	this.SiteAdmin = siteAdmin
	this.Name = name
	this.Company = company
	this.Blog = blog
	this.Location = location
	this.Email = email
	this.Hireable = hireable
	this.Bio = bio
	this.PublicRepos = publicRepos
	this.PublicGists = publicGists
	this.Followers = followers
	this.Following = following
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.PrivateGists = privateGists
	this.TotalPrivateRepos = totalPrivateRepos
	this.OwnedPrivateRepos = ownedPrivateRepos
	this.DiskUsage = diskUsage
	this.Collaborators = collaborators
	this.TwoFactorAuthentication = twoFactorAuthentication
	return &this
}

// NewPrivateUserWithDefaults instantiates a new PrivateUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateUserWithDefaults() *PrivateUser {
	this := PrivateUser{}
	return &this
}

// GetLogin returns the Login field value
func (o *PrivateUser) GetLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *PrivateUser) SetLogin(v string) {
	o.Login = v
}

// GetId returns the Id field value
func (o *PrivateUser) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PrivateUser) SetId(v int32) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *PrivateUser) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *PrivateUser) SetNodeId(v string) {
	o.NodeId = v
}

// GetAvatarUrl returns the AvatarUrl field value
func (o *PrivateUser) GetAvatarUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetAvatarUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvatarUrl, true
}

// SetAvatarUrl sets field value
func (o *PrivateUser) SetAvatarUrl(v string) {
	o.AvatarUrl = v
}

// GetGravatarId returns the GravatarId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PrivateUser) GetGravatarId() string {
	if o == nil || o.GravatarId.Get() == nil {
		var ret string
		return ret
	}

	return *o.GravatarId.Get()
}

// GetGravatarIdOk returns a tuple with the GravatarId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateUser) GetGravatarIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GravatarId.Get(), o.GravatarId.IsSet()
}

// SetGravatarId sets field value
func (o *PrivateUser) SetGravatarId(v string) {
	o.GravatarId.Set(&v)
}

// GetUrl returns the Url field value
func (o *PrivateUser) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PrivateUser) SetUrl(v string) {
	o.Url = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *PrivateUser) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *PrivateUser) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetFollowersUrl returns the FollowersUrl field value
func (o *PrivateUser) GetFollowersUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FollowersUrl
}

// GetFollowersUrlOk returns a tuple with the FollowersUrl field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetFollowersUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowersUrl, true
}

// SetFollowersUrl sets field value
func (o *PrivateUser) SetFollowersUrl(v string) {
	o.FollowersUrl = v
}

// GetFollowingUrl returns the FollowingUrl field value
func (o *PrivateUser) GetFollowingUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FollowingUrl
}

// GetFollowingUrlOk returns a tuple with the FollowingUrl field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetFollowingUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowingUrl, true
}

// SetFollowingUrl sets field value
func (o *PrivateUser) SetFollowingUrl(v string) {
	o.FollowingUrl = v
}

// GetGistsUrl returns the GistsUrl field value
func (o *PrivateUser) GetGistsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GistsUrl
}

// GetGistsUrlOk returns a tuple with the GistsUrl field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetGistsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GistsUrl, true
}

// SetGistsUrl sets field value
func (o *PrivateUser) SetGistsUrl(v string) {
	o.GistsUrl = v
}

// GetStarredUrl returns the StarredUrl field value
func (o *PrivateUser) GetStarredUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StarredUrl
}

// GetStarredUrlOk returns a tuple with the StarredUrl field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetStarredUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StarredUrl, true
}

// SetStarredUrl sets field value
func (o *PrivateUser) SetStarredUrl(v string) {
	o.StarredUrl = v
}

// GetSubscriptionsUrl returns the SubscriptionsUrl field value
func (o *PrivateUser) GetSubscriptionsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionsUrl
}

// GetSubscriptionsUrlOk returns a tuple with the SubscriptionsUrl field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetSubscriptionsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionsUrl, true
}

// SetSubscriptionsUrl sets field value
func (o *PrivateUser) SetSubscriptionsUrl(v string) {
	o.SubscriptionsUrl = v
}

// GetOrganizationsUrl returns the OrganizationsUrl field value
func (o *PrivateUser) GetOrganizationsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationsUrl
}

// GetOrganizationsUrlOk returns a tuple with the OrganizationsUrl field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetOrganizationsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationsUrl, true
}

// SetOrganizationsUrl sets field value
func (o *PrivateUser) SetOrganizationsUrl(v string) {
	o.OrganizationsUrl = v
}

// GetReposUrl returns the ReposUrl field value
func (o *PrivateUser) GetReposUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReposUrl
}

// GetReposUrlOk returns a tuple with the ReposUrl field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetReposUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReposUrl, true
}

// SetReposUrl sets field value
func (o *PrivateUser) SetReposUrl(v string) {
	o.ReposUrl = v
}

// GetEventsUrl returns the EventsUrl field value
func (o *PrivateUser) GetEventsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventsUrl
}

// GetEventsUrlOk returns a tuple with the EventsUrl field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetEventsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventsUrl, true
}

// SetEventsUrl sets field value
func (o *PrivateUser) SetEventsUrl(v string) {
	o.EventsUrl = v
}

// GetReceivedEventsUrl returns the ReceivedEventsUrl field value
func (o *PrivateUser) GetReceivedEventsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceivedEventsUrl
}

// GetReceivedEventsUrlOk returns a tuple with the ReceivedEventsUrl field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetReceivedEventsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceivedEventsUrl, true
}

// SetReceivedEventsUrl sets field value
func (o *PrivateUser) SetReceivedEventsUrl(v string) {
	o.ReceivedEventsUrl = v
}

// GetType returns the Type field value
func (o *PrivateUser) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PrivateUser) SetType(v string) {
	o.Type = v
}

// GetSiteAdmin returns the SiteAdmin field value
func (o *PrivateUser) GetSiteAdmin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SiteAdmin
}

// GetSiteAdminOk returns a tuple with the SiteAdmin field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetSiteAdminOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteAdmin, true
}

// SetSiteAdmin sets field value
func (o *PrivateUser) SetSiteAdmin(v bool) {
	o.SiteAdmin = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PrivateUser) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateUser) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *PrivateUser) SetName(v string) {
	o.Name.Set(&v)
}

// GetCompany returns the Company field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PrivateUser) GetCompany() string {
	if o == nil || o.Company.Get() == nil {
		var ret string
		return ret
	}

	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateUser) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// SetCompany sets field value
func (o *PrivateUser) SetCompany(v string) {
	o.Company.Set(&v)
}

// GetBlog returns the Blog field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PrivateUser) GetBlog() string {
	if o == nil || o.Blog.Get() == nil {
		var ret string
		return ret
	}

	return *o.Blog.Get()
}

// GetBlogOk returns a tuple with the Blog field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateUser) GetBlogOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Blog.Get(), o.Blog.IsSet()
}

// SetBlog sets field value
func (o *PrivateUser) SetBlog(v string) {
	o.Blog.Set(&v)
}

// GetLocation returns the Location field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PrivateUser) GetLocation() string {
	if o == nil || o.Location.Get() == nil {
		var ret string
		return ret
	}

	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateUser) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// SetLocation sets field value
func (o *PrivateUser) SetLocation(v string) {
	o.Location.Set(&v)
}

// GetEmail returns the Email field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PrivateUser) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}

	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateUser) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// SetEmail sets field value
func (o *PrivateUser) SetEmail(v string) {
	o.Email.Set(&v)
}

// GetHireable returns the Hireable field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *PrivateUser) GetHireable() bool {
	if o == nil || o.Hireable.Get() == nil {
		var ret bool
		return ret
	}

	return *o.Hireable.Get()
}

// GetHireableOk returns a tuple with the Hireable field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateUser) GetHireableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hireable.Get(), o.Hireable.IsSet()
}

// SetHireable sets field value
func (o *PrivateUser) SetHireable(v bool) {
	o.Hireable.Set(&v)
}

// GetBio returns the Bio field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PrivateUser) GetBio() string {
	if o == nil || o.Bio.Get() == nil {
		var ret string
		return ret
	}

	return *o.Bio.Get()
}

// GetBioOk returns a tuple with the Bio field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateUser) GetBioOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bio.Get(), o.Bio.IsSet()
}

// SetBio sets field value
func (o *PrivateUser) SetBio(v string) {
	o.Bio.Set(&v)
}

// GetTwitterUsername returns the TwitterUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateUser) GetTwitterUsername() string {
	if o == nil || o.TwitterUsername.Get() == nil {
		var ret string
		return ret
	}
	return *o.TwitterUsername.Get()
}

// GetTwitterUsernameOk returns a tuple with the TwitterUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateUser) GetTwitterUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TwitterUsername.Get(), o.TwitterUsername.IsSet()
}

// HasTwitterUsername returns a boolean if a field has been set.
func (o *PrivateUser) HasTwitterUsername() bool {
	if o != nil && o.TwitterUsername.IsSet() {
		return true
	}

	return false
}

// SetTwitterUsername gets a reference to the given NullableString and assigns it to the TwitterUsername field.
func (o *PrivateUser) SetTwitterUsername(v string) {
	o.TwitterUsername.Set(&v)
}
// SetTwitterUsernameNil sets the value for TwitterUsername to be an explicit nil
func (o *PrivateUser) SetTwitterUsernameNil() {
	o.TwitterUsername.Set(nil)
}

// UnsetTwitterUsername ensures that no value is present for TwitterUsername, not even an explicit nil
func (o *PrivateUser) UnsetTwitterUsername() {
	o.TwitterUsername.Unset()
}

// GetPublicRepos returns the PublicRepos field value
func (o *PrivateUser) GetPublicRepos() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PublicRepos
}

// GetPublicReposOk returns a tuple with the PublicRepos field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetPublicReposOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicRepos, true
}

// SetPublicRepos sets field value
func (o *PrivateUser) SetPublicRepos(v int32) {
	o.PublicRepos = v
}

// GetPublicGists returns the PublicGists field value
func (o *PrivateUser) GetPublicGists() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PublicGists
}

// GetPublicGistsOk returns a tuple with the PublicGists field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetPublicGistsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicGists, true
}

// SetPublicGists sets field value
func (o *PrivateUser) SetPublicGists(v int32) {
	o.PublicGists = v
}

// GetFollowers returns the Followers field value
func (o *PrivateUser) GetFollowers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Followers
}

// GetFollowersOk returns a tuple with the Followers field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetFollowersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Followers, true
}

// SetFollowers sets field value
func (o *PrivateUser) SetFollowers(v int32) {
	o.Followers = v
}

// GetFollowing returns the Following field value
func (o *PrivateUser) GetFollowing() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Following
}

// GetFollowingOk returns a tuple with the Following field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetFollowingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Following, true
}

// SetFollowing sets field value
func (o *PrivateUser) SetFollowing(v int32) {
	o.Following = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PrivateUser) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PrivateUser) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PrivateUser) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PrivateUser) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetPrivateGists returns the PrivateGists field value
func (o *PrivateUser) GetPrivateGists() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrivateGists
}

// GetPrivateGistsOk returns a tuple with the PrivateGists field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetPrivateGistsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateGists, true
}

// SetPrivateGists sets field value
func (o *PrivateUser) SetPrivateGists(v int32) {
	o.PrivateGists = v
}

// GetTotalPrivateRepos returns the TotalPrivateRepos field value
func (o *PrivateUser) GetTotalPrivateRepos() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalPrivateRepos
}

// GetTotalPrivateReposOk returns a tuple with the TotalPrivateRepos field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetTotalPrivateReposOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPrivateRepos, true
}

// SetTotalPrivateRepos sets field value
func (o *PrivateUser) SetTotalPrivateRepos(v int32) {
	o.TotalPrivateRepos = v
}

// GetOwnedPrivateRepos returns the OwnedPrivateRepos field value
func (o *PrivateUser) GetOwnedPrivateRepos() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OwnedPrivateRepos
}

// GetOwnedPrivateReposOk returns a tuple with the OwnedPrivateRepos field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetOwnedPrivateReposOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnedPrivateRepos, true
}

// SetOwnedPrivateRepos sets field value
func (o *PrivateUser) SetOwnedPrivateRepos(v int32) {
	o.OwnedPrivateRepos = v
}

// GetDiskUsage returns the DiskUsage field value
func (o *PrivateUser) GetDiskUsage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DiskUsage
}

// GetDiskUsageOk returns a tuple with the DiskUsage field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetDiskUsageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskUsage, true
}

// SetDiskUsage sets field value
func (o *PrivateUser) SetDiskUsage(v int32) {
	o.DiskUsage = v
}

// GetCollaborators returns the Collaborators field value
func (o *PrivateUser) GetCollaborators() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Collaborators
}

// GetCollaboratorsOk returns a tuple with the Collaborators field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetCollaboratorsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collaborators, true
}

// SetCollaborators sets field value
func (o *PrivateUser) SetCollaborators(v int32) {
	o.Collaborators = v
}

// GetTwoFactorAuthentication returns the TwoFactorAuthentication field value
func (o *PrivateUser) GetTwoFactorAuthentication() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TwoFactorAuthentication
}

// GetTwoFactorAuthenticationOk returns a tuple with the TwoFactorAuthentication field value
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetTwoFactorAuthenticationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TwoFactorAuthentication, true
}

// SetTwoFactorAuthentication sets field value
func (o *PrivateUser) SetTwoFactorAuthentication(v bool) {
	o.TwoFactorAuthentication = v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *PrivateUser) GetPlan() PublicUserPlan {
	if o == nil || o.Plan == nil {
		var ret PublicUserPlan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetPlanOk() (*PublicUserPlan, bool) {
	if o == nil || o.Plan == nil {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *PrivateUser) HasPlan() bool {
	if o != nil && o.Plan != nil {
		return true
	}

	return false
}

// SetPlan gets a reference to the given PublicUserPlan and assigns it to the Plan field.
func (o *PrivateUser) SetPlan(v PublicUserPlan) {
	o.Plan = &v
}

// GetSuspendedAt returns the SuspendedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateUser) GetSuspendedAt() time.Time {
	if o == nil || o.SuspendedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SuspendedAt.Get()
}

// GetSuspendedAtOk returns a tuple with the SuspendedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateUser) GetSuspendedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SuspendedAt.Get(), o.SuspendedAt.IsSet()
}

// HasSuspendedAt returns a boolean if a field has been set.
func (o *PrivateUser) HasSuspendedAt() bool {
	if o != nil && o.SuspendedAt.IsSet() {
		return true
	}

	return false
}

// SetSuspendedAt gets a reference to the given NullableTime and assigns it to the SuspendedAt field.
func (o *PrivateUser) SetSuspendedAt(v time.Time) {
	o.SuspendedAt.Set(&v)
}
// SetSuspendedAtNil sets the value for SuspendedAt to be an explicit nil
func (o *PrivateUser) SetSuspendedAtNil() {
	o.SuspendedAt.Set(nil)
}

// UnsetSuspendedAt ensures that no value is present for SuspendedAt, not even an explicit nil
func (o *PrivateUser) UnsetSuspendedAt() {
	o.SuspendedAt.Unset()
}

// GetBusinessPlus returns the BusinessPlus field value if set, zero value otherwise.
func (o *PrivateUser) GetBusinessPlus() bool {
	if o == nil || o.BusinessPlus == nil {
		var ret bool
		return ret
	}
	return *o.BusinessPlus
}

// GetBusinessPlusOk returns a tuple with the BusinessPlus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetBusinessPlusOk() (*bool, bool) {
	if o == nil || o.BusinessPlus == nil {
		return nil, false
	}
	return o.BusinessPlus, true
}

// HasBusinessPlus returns a boolean if a field has been set.
func (o *PrivateUser) HasBusinessPlus() bool {
	if o != nil && o.BusinessPlus != nil {
		return true
	}

	return false
}

// SetBusinessPlus gets a reference to the given bool and assigns it to the BusinessPlus field.
func (o *PrivateUser) SetBusinessPlus(v bool) {
	o.BusinessPlus = &v
}

// GetLdapDn returns the LdapDn field value if set, zero value otherwise.
func (o *PrivateUser) GetLdapDn() string {
	if o == nil || o.LdapDn == nil {
		var ret string
		return ret
	}
	return *o.LdapDn
}

// GetLdapDnOk returns a tuple with the LdapDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateUser) GetLdapDnOk() (*string, bool) {
	if o == nil || o.LdapDn == nil {
		return nil, false
	}
	return o.LdapDn, true
}

// HasLdapDn returns a boolean if a field has been set.
func (o *PrivateUser) HasLdapDn() bool {
	if o != nil && o.LdapDn != nil {
		return true
	}

	return false
}

// SetLdapDn gets a reference to the given string and assigns it to the LdapDn field.
func (o *PrivateUser) SetLdapDn(v string) {
	o.LdapDn = &v
}

func (o PrivateUser) MarshalJSON() ([]byte, error) {
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
		toSerialize["following_url"] = o.FollowingUrl
	}
	if true {
		toSerialize["gists_url"] = o.GistsUrl
	}
	if true {
		toSerialize["starred_url"] = o.StarredUrl
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
		toSerialize["events_url"] = o.EventsUrl
	}
	if true {
		toSerialize["received_events_url"] = o.ReceivedEventsUrl
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["site_admin"] = o.SiteAdmin
	}
	if true {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["company"] = o.Company.Get()
	}
	if true {
		toSerialize["blog"] = o.Blog.Get()
	}
	if true {
		toSerialize["location"] = o.Location.Get()
	}
	if true {
		toSerialize["email"] = o.Email.Get()
	}
	if true {
		toSerialize["hireable"] = o.Hireable.Get()
	}
	if true {
		toSerialize["bio"] = o.Bio.Get()
	}
	if o.TwitterUsername.IsSet() {
		toSerialize["twitter_username"] = o.TwitterUsername.Get()
	}
	if true {
		toSerialize["public_repos"] = o.PublicRepos
	}
	if true {
		toSerialize["public_gists"] = o.PublicGists
	}
	if true {
		toSerialize["followers"] = o.Followers
	}
	if true {
		toSerialize["following"] = o.Following
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["private_gists"] = o.PrivateGists
	}
	if true {
		toSerialize["total_private_repos"] = o.TotalPrivateRepos
	}
	if true {
		toSerialize["owned_private_repos"] = o.OwnedPrivateRepos
	}
	if true {
		toSerialize["disk_usage"] = o.DiskUsage
	}
	if true {
		toSerialize["collaborators"] = o.Collaborators
	}
	if true {
		toSerialize["two_factor_authentication"] = o.TwoFactorAuthentication
	}
	if o.Plan != nil {
		toSerialize["plan"] = o.Plan
	}
	if o.SuspendedAt.IsSet() {
		toSerialize["suspended_at"] = o.SuspendedAt.Get()
	}
	if o.BusinessPlus != nil {
		toSerialize["business_plus"] = o.BusinessPlus
	}
	if o.LdapDn != nil {
		toSerialize["ldap_dn"] = o.LdapDn
	}
	return json.Marshal(toSerialize)
}

type NullablePrivateUser struct {
	value *PrivateUser
	isSet bool
}

func (v NullablePrivateUser) Get() *PrivateUser {
	return v.value
}

func (v *NullablePrivateUser) Set(val *PrivateUser) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateUser) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateUser(val *PrivateUser) *NullablePrivateUser {
	return &NullablePrivateUser{value: val, isSet: true}
}

func (v NullablePrivateUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


