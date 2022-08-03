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

// ScimUserMeta struct for ScimUserMeta
type ScimUserMeta struct {
	ResourceType *string `json:"resourceType,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	LastModified *time.Time `json:"lastModified,omitempty"`
	Location *string `json:"location,omitempty"`
}

// NewScimUserMeta instantiates a new ScimUserMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimUserMeta() *ScimUserMeta {
	this := ScimUserMeta{}
	return &this
}

// NewScimUserMetaWithDefaults instantiates a new ScimUserMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimUserMetaWithDefaults() *ScimUserMeta {
	this := ScimUserMeta{}
	return &this
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *ScimUserMeta) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUserMeta) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *ScimUserMeta) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *ScimUserMeta) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ScimUserMeta) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUserMeta) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ScimUserMeta) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ScimUserMeta) SetCreated(v time.Time) {
	o.Created = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *ScimUserMeta) GetLastModified() time.Time {
	if o == nil || o.LastModified == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUserMeta) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || o.LastModified == nil {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *ScimUserMeta) HasLastModified() bool {
	if o != nil && o.LastModified != nil {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *ScimUserMeta) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ScimUserMeta) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUserMeta) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ScimUserMeta) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *ScimUserMeta) SetLocation(v string) {
	o.Location = &v
}

func (o ScimUserMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.LastModified != nil {
		toSerialize["lastModified"] = o.LastModified
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	return json.Marshal(toSerialize)
}

type NullableScimUserMeta struct {
	value *ScimUserMeta
	isSet bool
}

func (v NullableScimUserMeta) Get() *ScimUserMeta {
	return v.value
}

func (v *NullableScimUserMeta) Set(val *ScimUserMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableScimUserMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableScimUserMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimUserMeta(val *ScimUserMeta) *NullableScimUserMeta {
	return &NullableScimUserMeta{value: val, isSet: true}
}

func (v NullableScimUserMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimUserMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


