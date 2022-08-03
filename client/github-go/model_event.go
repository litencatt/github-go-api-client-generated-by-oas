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

// Event Event
type Event struct {
	Id string `json:"id"`
	Type NullableString `json:"type"`
	Actor Actor `json:"actor"`
	Repo EventRepo `json:"repo"`
	Org *Actor `json:"org,omitempty"`
	Payload EventPayload `json:"payload"`
	Public bool `json:"public"`
	CreatedAt NullableTime `json:"created_at"`
}

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent(id string, type_ NullableString, actor Actor, repo EventRepo, payload EventPayload, public bool, createdAt NullableTime) *Event {
	this := Event{}
	this.Id = id
	this.Type = type_
	this.Actor = actor
	this.Repo = repo
	this.Payload = payload
	this.Public = public
	this.CreatedAt = createdAt
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetId returns the Id field value
func (o *Event) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Event) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Event) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Event) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}

	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Event) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// SetType sets field value
func (o *Event) SetType(v string) {
	o.Type.Set(&v)
}

// GetActor returns the Actor field value
func (o *Event) GetActor() Actor {
	if o == nil {
		var ret Actor
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *Event) GetActorOk() (*Actor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *Event) SetActor(v Actor) {
	o.Actor = v
}

// GetRepo returns the Repo field value
func (o *Event) GetRepo() EventRepo {
	if o == nil {
		var ret EventRepo
		return ret
	}

	return o.Repo
}

// GetRepoOk returns a tuple with the Repo field value
// and a boolean to check if the value has been set.
func (o *Event) GetRepoOk() (*EventRepo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repo, true
}

// SetRepo sets field value
func (o *Event) SetRepo(v EventRepo) {
	o.Repo = v
}

// GetOrg returns the Org field value if set, zero value otherwise.
func (o *Event) GetOrg() Actor {
	if o == nil || o.Org == nil {
		var ret Actor
		return ret
	}
	return *o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetOrgOk() (*Actor, bool) {
	if o == nil || o.Org == nil {
		return nil, false
	}
	return o.Org, true
}

// HasOrg returns a boolean if a field has been set.
func (o *Event) HasOrg() bool {
	if o != nil && o.Org != nil {
		return true
	}

	return false
}

// SetOrg gets a reference to the given Actor and assigns it to the Org field.
func (o *Event) SetOrg(v Actor) {
	o.Org = &v
}

// GetPayload returns the Payload field value
func (o *Event) GetPayload() EventPayload {
	if o == nil {
		var ret EventPayload
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *Event) GetPayloadOk() (*EventPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *Event) SetPayload(v EventPayload) {
	o.Payload = v
}

// GetPublic returns the Public field value
func (o *Event) GetPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Public
}

// GetPublicOk returns a tuple with the Public field value
// and a boolean to check if the value has been set.
func (o *Event) GetPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Public, true
}

// SetPublic sets field value
func (o *Event) SetPublic(v bool) {
	o.Public = v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Event) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Event) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *Event) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

func (o Event) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type.Get()
	}
	if true {
		toSerialize["actor"] = o.Actor
	}
	if true {
		toSerialize["repo"] = o.Repo
	}
	if o.Org != nil {
		toSerialize["org"] = o.Org
	}
	if true {
		toSerialize["payload"] = o.Payload
	}
	if true {
		toSerialize["public"] = o.Public
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


