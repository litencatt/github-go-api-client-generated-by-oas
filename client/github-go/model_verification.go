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

// Verification struct for Verification
type Verification struct {
	Verified bool `json:"verified"`
	Reason string `json:"reason"`
	Payload NullableString `json:"payload"`
	Signature NullableString `json:"signature"`
}

// NewVerification instantiates a new Verification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerification(verified bool, reason string, payload NullableString, signature NullableString) *Verification {
	this := Verification{}
	this.Verified = verified
	this.Reason = reason
	this.Payload = payload
	this.Signature = signature
	return &this
}

// NewVerificationWithDefaults instantiates a new Verification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationWithDefaults() *Verification {
	this := Verification{}
	return &this
}

// GetVerified returns the Verified field value
func (o *Verification) GetVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value
// and a boolean to check if the value has been set.
func (o *Verification) GetVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verified, true
}

// SetVerified sets field value
func (o *Verification) SetVerified(v bool) {
	o.Verified = v
}

// GetReason returns the Reason field value
func (o *Verification) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *Verification) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *Verification) SetReason(v string) {
	o.Reason = v
}

// GetPayload returns the Payload field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Verification) GetPayload() string {
	if o == nil || o.Payload.Get() == nil {
		var ret string
		return ret
	}

	return *o.Payload.Get()
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Verification) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payload.Get(), o.Payload.IsSet()
}

// SetPayload sets field value
func (o *Verification) SetPayload(v string) {
	o.Payload.Set(&v)
}

// GetSignature returns the Signature field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Verification) GetSignature() string {
	if o == nil || o.Signature.Get() == nil {
		var ret string
		return ret
	}

	return *o.Signature.Get()
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Verification) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signature.Get(), o.Signature.IsSet()
}

// SetSignature sets field value
func (o *Verification) SetSignature(v string) {
	o.Signature.Set(&v)
}

func (o Verification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["verified"] = o.Verified
	}
	if true {
		toSerialize["reason"] = o.Reason
	}
	if true {
		toSerialize["payload"] = o.Payload.Get()
	}
	if true {
		toSerialize["signature"] = o.Signature.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableVerification struct {
	value *Verification
	isSet bool
}

func (v NullableVerification) Get() *Verification {
	return v.value
}

func (v *NullableVerification) Set(val *Verification) {
	v.value = val
	v.isSet = true
}

func (v NullableVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerification(val *Verification) *NullableVerification {
	return &NullableVerification{value: val, isSet: true}
}

func (v NullableVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


