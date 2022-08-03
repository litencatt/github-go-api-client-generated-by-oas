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

// ActivityMarkNotificationsAsReadRequest struct for ActivityMarkNotificationsAsReadRequest
type ActivityMarkNotificationsAsReadRequest struct {
	// Describes the last point that notifications were checked. Anything updated since this time will not be marked as read. If you omit this parameter, all notifications are marked as read. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`. Default: The current timestamp.
	LastReadAt *time.Time `json:"last_read_at,omitempty"`
	// Whether the notification has been read.
	Read *bool `json:"read,omitempty"`
}

// NewActivityMarkNotificationsAsReadRequest instantiates a new ActivityMarkNotificationsAsReadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityMarkNotificationsAsReadRequest() *ActivityMarkNotificationsAsReadRequest {
	this := ActivityMarkNotificationsAsReadRequest{}
	return &this
}

// NewActivityMarkNotificationsAsReadRequestWithDefaults instantiates a new ActivityMarkNotificationsAsReadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityMarkNotificationsAsReadRequestWithDefaults() *ActivityMarkNotificationsAsReadRequest {
	this := ActivityMarkNotificationsAsReadRequest{}
	return &this
}

// GetLastReadAt returns the LastReadAt field value if set, zero value otherwise.
func (o *ActivityMarkNotificationsAsReadRequest) GetLastReadAt() time.Time {
	if o == nil || o.LastReadAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastReadAt
}

// GetLastReadAtOk returns a tuple with the LastReadAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityMarkNotificationsAsReadRequest) GetLastReadAtOk() (*time.Time, bool) {
	if o == nil || o.LastReadAt == nil {
		return nil, false
	}
	return o.LastReadAt, true
}

// HasLastReadAt returns a boolean if a field has been set.
func (o *ActivityMarkNotificationsAsReadRequest) HasLastReadAt() bool {
	if o != nil && o.LastReadAt != nil {
		return true
	}

	return false
}

// SetLastReadAt gets a reference to the given time.Time and assigns it to the LastReadAt field.
func (o *ActivityMarkNotificationsAsReadRequest) SetLastReadAt(v time.Time) {
	o.LastReadAt = &v
}

// GetRead returns the Read field value if set, zero value otherwise.
func (o *ActivityMarkNotificationsAsReadRequest) GetRead() bool {
	if o == nil || o.Read == nil {
		var ret bool
		return ret
	}
	return *o.Read
}

// GetReadOk returns a tuple with the Read field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityMarkNotificationsAsReadRequest) GetReadOk() (*bool, bool) {
	if o == nil || o.Read == nil {
		return nil, false
	}
	return o.Read, true
}

// HasRead returns a boolean if a field has been set.
func (o *ActivityMarkNotificationsAsReadRequest) HasRead() bool {
	if o != nil && o.Read != nil {
		return true
	}

	return false
}

// SetRead gets a reference to the given bool and assigns it to the Read field.
func (o *ActivityMarkNotificationsAsReadRequest) SetRead(v bool) {
	o.Read = &v
}

func (o ActivityMarkNotificationsAsReadRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastReadAt != nil {
		toSerialize["last_read_at"] = o.LastReadAt
	}
	if o.Read != nil {
		toSerialize["read"] = o.Read
	}
	return json.Marshal(toSerialize)
}

type NullableActivityMarkNotificationsAsReadRequest struct {
	value *ActivityMarkNotificationsAsReadRequest
	isSet bool
}

func (v NullableActivityMarkNotificationsAsReadRequest) Get() *ActivityMarkNotificationsAsReadRequest {
	return v.value
}

func (v *NullableActivityMarkNotificationsAsReadRequest) Set(val *ActivityMarkNotificationsAsReadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityMarkNotificationsAsReadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityMarkNotificationsAsReadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityMarkNotificationsAsReadRequest(val *ActivityMarkNotificationsAsReadRequest) *NullableActivityMarkNotificationsAsReadRequest {
	return &NullableActivityMarkNotificationsAsReadRequest{value: val, isSet: true}
}

func (v NullableActivityMarkNotificationsAsReadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityMarkNotificationsAsReadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

