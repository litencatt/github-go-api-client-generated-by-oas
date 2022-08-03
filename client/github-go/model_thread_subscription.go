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

// ThreadSubscription Thread Subscription
type ThreadSubscription struct {
	Subscribed bool `json:"subscribed"`
	Ignored bool `json:"ignored"`
	Reason NullableString `json:"reason"`
	CreatedAt NullableTime `json:"created_at"`
	Url string `json:"url"`
	ThreadUrl *string `json:"thread_url,omitempty"`
	RepositoryUrl *string `json:"repository_url,omitempty"`
}

// NewThreadSubscription instantiates a new ThreadSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreadSubscription(subscribed bool, ignored bool, reason NullableString, createdAt NullableTime, url string) *ThreadSubscription {
	this := ThreadSubscription{}
	this.Subscribed = subscribed
	this.Ignored = ignored
	this.Reason = reason
	this.CreatedAt = createdAt
	this.Url = url
	return &this
}

// NewThreadSubscriptionWithDefaults instantiates a new ThreadSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreadSubscriptionWithDefaults() *ThreadSubscription {
	this := ThreadSubscription{}
	return &this
}

// GetSubscribed returns the Subscribed field value
func (o *ThreadSubscription) GetSubscribed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Subscribed
}

// GetSubscribedOk returns a tuple with the Subscribed field value
// and a boolean to check if the value has been set.
func (o *ThreadSubscription) GetSubscribedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscribed, true
}

// SetSubscribed sets field value
func (o *ThreadSubscription) SetSubscribed(v bool) {
	o.Subscribed = v
}

// GetIgnored returns the Ignored field value
func (o *ThreadSubscription) GetIgnored() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ignored
}

// GetIgnoredOk returns a tuple with the Ignored field value
// and a boolean to check if the value has been set.
func (o *ThreadSubscription) GetIgnoredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ignored, true
}

// SetIgnored sets field value
func (o *ThreadSubscription) SetIgnored(v bool) {
	o.Ignored = v
}

// GetReason returns the Reason field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ThreadSubscription) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}

	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadSubscription) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// SetReason sets field value
func (o *ThreadSubscription) SetReason(v string) {
	o.Reason.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ThreadSubscription) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadSubscription) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *ThreadSubscription) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// GetUrl returns the Url field value
func (o *ThreadSubscription) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ThreadSubscription) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ThreadSubscription) SetUrl(v string) {
	o.Url = v
}

// GetThreadUrl returns the ThreadUrl field value if set, zero value otherwise.
func (o *ThreadSubscription) GetThreadUrl() string {
	if o == nil || o.ThreadUrl == nil {
		var ret string
		return ret
	}
	return *o.ThreadUrl
}

// GetThreadUrlOk returns a tuple with the ThreadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadSubscription) GetThreadUrlOk() (*string, bool) {
	if o == nil || o.ThreadUrl == nil {
		return nil, false
	}
	return o.ThreadUrl, true
}

// HasThreadUrl returns a boolean if a field has been set.
func (o *ThreadSubscription) HasThreadUrl() bool {
	if o != nil && o.ThreadUrl != nil {
		return true
	}

	return false
}

// SetThreadUrl gets a reference to the given string and assigns it to the ThreadUrl field.
func (o *ThreadSubscription) SetThreadUrl(v string) {
	o.ThreadUrl = &v
}

// GetRepositoryUrl returns the RepositoryUrl field value if set, zero value otherwise.
func (o *ThreadSubscription) GetRepositoryUrl() string {
	if o == nil || o.RepositoryUrl == nil {
		var ret string
		return ret
	}
	return *o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreadSubscription) GetRepositoryUrlOk() (*string, bool) {
	if o == nil || o.RepositoryUrl == nil {
		return nil, false
	}
	return o.RepositoryUrl, true
}

// HasRepositoryUrl returns a boolean if a field has been set.
func (o *ThreadSubscription) HasRepositoryUrl() bool {
	if o != nil && o.RepositoryUrl != nil {
		return true
	}

	return false
}

// SetRepositoryUrl gets a reference to the given string and assigns it to the RepositoryUrl field.
func (o *ThreadSubscription) SetRepositoryUrl(v string) {
	o.RepositoryUrl = &v
}

func (o ThreadSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subscribed"] = o.Subscribed
	}
	if true {
		toSerialize["ignored"] = o.Ignored
	}
	if true {
		toSerialize["reason"] = o.Reason.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.ThreadUrl != nil {
		toSerialize["thread_url"] = o.ThreadUrl
	}
	if o.RepositoryUrl != nil {
		toSerialize["repository_url"] = o.RepositoryUrl
	}
	return json.Marshal(toSerialize)
}

type NullableThreadSubscription struct {
	value *ThreadSubscription
	isSet bool
}

func (v NullableThreadSubscription) Get() *ThreadSubscription {
	return v.value
}

func (v *NullableThreadSubscription) Set(val *ThreadSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableThreadSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableThreadSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreadSubscription(val *ThreadSubscription) *NullableThreadSubscription {
	return &NullableThreadSubscription{value: val, isSet: true}
}

func (v NullableThreadSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreadSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

