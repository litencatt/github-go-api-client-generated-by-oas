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

// FeedLinks struct for FeedLinks
type FeedLinks struct {
	Timeline LinkWithType `json:"timeline"`
	User LinkWithType `json:"user"`
	SecurityAdvisories *LinkWithType `json:"security_advisories,omitempty"`
	CurrentUser *LinkWithType `json:"current_user,omitempty"`
	CurrentUserPublic *LinkWithType `json:"current_user_public,omitempty"`
	CurrentUserActor *LinkWithType `json:"current_user_actor,omitempty"`
	CurrentUserOrganization *LinkWithType `json:"current_user_organization,omitempty"`
	CurrentUserOrganizations []LinkWithType `json:"current_user_organizations,omitempty"`
}

// NewFeedLinks instantiates a new FeedLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedLinks(timeline LinkWithType, user LinkWithType) *FeedLinks {
	this := FeedLinks{}
	this.Timeline = timeline
	this.User = user
	return &this
}

// NewFeedLinksWithDefaults instantiates a new FeedLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedLinksWithDefaults() *FeedLinks {
	this := FeedLinks{}
	return &this
}

// GetTimeline returns the Timeline field value
func (o *FeedLinks) GetTimeline() LinkWithType {
	if o == nil {
		var ret LinkWithType
		return ret
	}

	return o.Timeline
}

// GetTimelineOk returns a tuple with the Timeline field value
// and a boolean to check if the value has been set.
func (o *FeedLinks) GetTimelineOk() (*LinkWithType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeline, true
}

// SetTimeline sets field value
func (o *FeedLinks) SetTimeline(v LinkWithType) {
	o.Timeline = v
}

// GetUser returns the User field value
func (o *FeedLinks) GetUser() LinkWithType {
	if o == nil {
		var ret LinkWithType
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *FeedLinks) GetUserOk() (*LinkWithType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *FeedLinks) SetUser(v LinkWithType) {
	o.User = v
}

// GetSecurityAdvisories returns the SecurityAdvisories field value if set, zero value otherwise.
func (o *FeedLinks) GetSecurityAdvisories() LinkWithType {
	if o == nil || o.SecurityAdvisories == nil {
		var ret LinkWithType
		return ret
	}
	return *o.SecurityAdvisories
}

// GetSecurityAdvisoriesOk returns a tuple with the SecurityAdvisories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedLinks) GetSecurityAdvisoriesOk() (*LinkWithType, bool) {
	if o == nil || o.SecurityAdvisories == nil {
		return nil, false
	}
	return o.SecurityAdvisories, true
}

// HasSecurityAdvisories returns a boolean if a field has been set.
func (o *FeedLinks) HasSecurityAdvisories() bool {
	if o != nil && o.SecurityAdvisories != nil {
		return true
	}

	return false
}

// SetSecurityAdvisories gets a reference to the given LinkWithType and assigns it to the SecurityAdvisories field.
func (o *FeedLinks) SetSecurityAdvisories(v LinkWithType) {
	o.SecurityAdvisories = &v
}

// GetCurrentUser returns the CurrentUser field value if set, zero value otherwise.
func (o *FeedLinks) GetCurrentUser() LinkWithType {
	if o == nil || o.CurrentUser == nil {
		var ret LinkWithType
		return ret
	}
	return *o.CurrentUser
}

// GetCurrentUserOk returns a tuple with the CurrentUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedLinks) GetCurrentUserOk() (*LinkWithType, bool) {
	if o == nil || o.CurrentUser == nil {
		return nil, false
	}
	return o.CurrentUser, true
}

// HasCurrentUser returns a boolean if a field has been set.
func (o *FeedLinks) HasCurrentUser() bool {
	if o != nil && o.CurrentUser != nil {
		return true
	}

	return false
}

// SetCurrentUser gets a reference to the given LinkWithType and assigns it to the CurrentUser field.
func (o *FeedLinks) SetCurrentUser(v LinkWithType) {
	o.CurrentUser = &v
}

// GetCurrentUserPublic returns the CurrentUserPublic field value if set, zero value otherwise.
func (o *FeedLinks) GetCurrentUserPublic() LinkWithType {
	if o == nil || o.CurrentUserPublic == nil {
		var ret LinkWithType
		return ret
	}
	return *o.CurrentUserPublic
}

// GetCurrentUserPublicOk returns a tuple with the CurrentUserPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedLinks) GetCurrentUserPublicOk() (*LinkWithType, bool) {
	if o == nil || o.CurrentUserPublic == nil {
		return nil, false
	}
	return o.CurrentUserPublic, true
}

// HasCurrentUserPublic returns a boolean if a field has been set.
func (o *FeedLinks) HasCurrentUserPublic() bool {
	if o != nil && o.CurrentUserPublic != nil {
		return true
	}

	return false
}

// SetCurrentUserPublic gets a reference to the given LinkWithType and assigns it to the CurrentUserPublic field.
func (o *FeedLinks) SetCurrentUserPublic(v LinkWithType) {
	o.CurrentUserPublic = &v
}

// GetCurrentUserActor returns the CurrentUserActor field value if set, zero value otherwise.
func (o *FeedLinks) GetCurrentUserActor() LinkWithType {
	if o == nil || o.CurrentUserActor == nil {
		var ret LinkWithType
		return ret
	}
	return *o.CurrentUserActor
}

// GetCurrentUserActorOk returns a tuple with the CurrentUserActor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedLinks) GetCurrentUserActorOk() (*LinkWithType, bool) {
	if o == nil || o.CurrentUserActor == nil {
		return nil, false
	}
	return o.CurrentUserActor, true
}

// HasCurrentUserActor returns a boolean if a field has been set.
func (o *FeedLinks) HasCurrentUserActor() bool {
	if o != nil && o.CurrentUserActor != nil {
		return true
	}

	return false
}

// SetCurrentUserActor gets a reference to the given LinkWithType and assigns it to the CurrentUserActor field.
func (o *FeedLinks) SetCurrentUserActor(v LinkWithType) {
	o.CurrentUserActor = &v
}

// GetCurrentUserOrganization returns the CurrentUserOrganization field value if set, zero value otherwise.
func (o *FeedLinks) GetCurrentUserOrganization() LinkWithType {
	if o == nil || o.CurrentUserOrganization == nil {
		var ret LinkWithType
		return ret
	}
	return *o.CurrentUserOrganization
}

// GetCurrentUserOrganizationOk returns a tuple with the CurrentUserOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedLinks) GetCurrentUserOrganizationOk() (*LinkWithType, bool) {
	if o == nil || o.CurrentUserOrganization == nil {
		return nil, false
	}
	return o.CurrentUserOrganization, true
}

// HasCurrentUserOrganization returns a boolean if a field has been set.
func (o *FeedLinks) HasCurrentUserOrganization() bool {
	if o != nil && o.CurrentUserOrganization != nil {
		return true
	}

	return false
}

// SetCurrentUserOrganization gets a reference to the given LinkWithType and assigns it to the CurrentUserOrganization field.
func (o *FeedLinks) SetCurrentUserOrganization(v LinkWithType) {
	o.CurrentUserOrganization = &v
}

// GetCurrentUserOrganizations returns the CurrentUserOrganizations field value if set, zero value otherwise.
func (o *FeedLinks) GetCurrentUserOrganizations() []LinkWithType {
	if o == nil || o.CurrentUserOrganizations == nil {
		var ret []LinkWithType
		return ret
	}
	return o.CurrentUserOrganizations
}

// GetCurrentUserOrganizationsOk returns a tuple with the CurrentUserOrganizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedLinks) GetCurrentUserOrganizationsOk() ([]LinkWithType, bool) {
	if o == nil || o.CurrentUserOrganizations == nil {
		return nil, false
	}
	return o.CurrentUserOrganizations, true
}

// HasCurrentUserOrganizations returns a boolean if a field has been set.
func (o *FeedLinks) HasCurrentUserOrganizations() bool {
	if o != nil && o.CurrentUserOrganizations != nil {
		return true
	}

	return false
}

// SetCurrentUserOrganizations gets a reference to the given []LinkWithType and assigns it to the CurrentUserOrganizations field.
func (o *FeedLinks) SetCurrentUserOrganizations(v []LinkWithType) {
	o.CurrentUserOrganizations = v
}

func (o FeedLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["timeline"] = o.Timeline
	}
	if true {
		toSerialize["user"] = o.User
	}
	if o.SecurityAdvisories != nil {
		toSerialize["security_advisories"] = o.SecurityAdvisories
	}
	if o.CurrentUser != nil {
		toSerialize["current_user"] = o.CurrentUser
	}
	if o.CurrentUserPublic != nil {
		toSerialize["current_user_public"] = o.CurrentUserPublic
	}
	if o.CurrentUserActor != nil {
		toSerialize["current_user_actor"] = o.CurrentUserActor
	}
	if o.CurrentUserOrganization != nil {
		toSerialize["current_user_organization"] = o.CurrentUserOrganization
	}
	if o.CurrentUserOrganizations != nil {
		toSerialize["current_user_organizations"] = o.CurrentUserOrganizations
	}
	return json.Marshal(toSerialize)
}

type NullableFeedLinks struct {
	value *FeedLinks
	isSet bool
}

func (v NullableFeedLinks) Get() *FeedLinks {
	return v.value
}

func (v *NullableFeedLinks) Set(val *FeedLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedLinks(val *FeedLinks) *NullableFeedLinks {
	return &NullableFeedLinks{value: val, isSet: true}
}

func (v NullableFeedLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


