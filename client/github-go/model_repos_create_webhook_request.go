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

// ReposCreateWebhookRequest struct for ReposCreateWebhookRequest
type ReposCreateWebhookRequest struct {
	// Use `web` to create a webhook. Default: `web`. This parameter only accepts the value `web`.
	Name *string `json:"name,omitempty"`
	Config *ReposCreateWebhookRequestConfig `json:"config,omitempty"`
	// Determines what [events](https://docs.github.com/webhooks/event-payloads) the hook is triggered for.
	Events []string `json:"events,omitempty"`
	// Determines if notifications are sent when the webhook is triggered. Set to `true` to send notifications.
	Active *bool `json:"active,omitempty"`
}

// NewReposCreateWebhookRequest instantiates a new ReposCreateWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReposCreateWebhookRequest() *ReposCreateWebhookRequest {
	this := ReposCreateWebhookRequest{}
	var active bool = true
	this.Active = &active
	return &this
}

// NewReposCreateWebhookRequestWithDefaults instantiates a new ReposCreateWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReposCreateWebhookRequestWithDefaults() *ReposCreateWebhookRequest {
	this := ReposCreateWebhookRequest{}
	var active bool = true
	this.Active = &active
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReposCreateWebhookRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateWebhookRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReposCreateWebhookRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReposCreateWebhookRequest) SetName(v string) {
	o.Name = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ReposCreateWebhookRequest) GetConfig() ReposCreateWebhookRequestConfig {
	if o == nil || o.Config == nil {
		var ret ReposCreateWebhookRequestConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateWebhookRequest) GetConfigOk() (*ReposCreateWebhookRequestConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ReposCreateWebhookRequest) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ReposCreateWebhookRequestConfig and assigns it to the Config field.
func (o *ReposCreateWebhookRequest) SetConfig(v ReposCreateWebhookRequestConfig) {
	o.Config = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ReposCreateWebhookRequest) GetEvents() []string {
	if o == nil || o.Events == nil {
		var ret []string
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateWebhookRequest) GetEventsOk() ([]string, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ReposCreateWebhookRequest) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []string and assigns it to the Events field.
func (o *ReposCreateWebhookRequest) SetEvents(v []string) {
	o.Events = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ReposCreateWebhookRequest) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateWebhookRequest) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ReposCreateWebhookRequest) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ReposCreateWebhookRequest) SetActive(v bool) {
	o.Active = &v
}

func (o ReposCreateWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	return json.Marshal(toSerialize)
}

type NullableReposCreateWebhookRequest struct {
	value *ReposCreateWebhookRequest
	isSet bool
}

func (v NullableReposCreateWebhookRequest) Get() *ReposCreateWebhookRequest {
	return v.value
}

func (v *NullableReposCreateWebhookRequest) Set(val *ReposCreateWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReposCreateWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReposCreateWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReposCreateWebhookRequest(val *ReposCreateWebhookRequest) *NullableReposCreateWebhookRequest {
	return &NullableReposCreateWebhookRequest{value: val, isSet: true}
}

func (v NullableReposCreateWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReposCreateWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


