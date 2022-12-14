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

// GpgKey A unique encryption key
type GpgKey struct {
	Id int32 `json:"id"`
	Name NullableString `json:"name,omitempty"`
	PrimaryKeyId NullableInt32 `json:"primary_key_id"`
	KeyId string `json:"key_id"`
	PublicKey string `json:"public_key"`
	Emails []GpgKeyEmailsInner `json:"emails"`
	Subkeys []GpgKeySubkeysInner `json:"subkeys"`
	CanSign bool `json:"can_sign"`
	CanEncryptComms bool `json:"can_encrypt_comms"`
	CanEncryptStorage bool `json:"can_encrypt_storage"`
	CanCertify bool `json:"can_certify"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt NullableTime `json:"expires_at"`
	Revoked bool `json:"revoked"`
	RawKey NullableString `json:"raw_key"`
}

// NewGpgKey instantiates a new GpgKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGpgKey(id int32, primaryKeyId NullableInt32, keyId string, publicKey string, emails []GpgKeyEmailsInner, subkeys []GpgKeySubkeysInner, canSign bool, canEncryptComms bool, canEncryptStorage bool, canCertify bool, createdAt time.Time, expiresAt NullableTime, revoked bool, rawKey NullableString) *GpgKey {
	this := GpgKey{}
	this.Id = id
	this.PrimaryKeyId = primaryKeyId
	this.KeyId = keyId
	this.PublicKey = publicKey
	this.Emails = emails
	this.Subkeys = subkeys
	this.CanSign = canSign
	this.CanEncryptComms = canEncryptComms
	this.CanEncryptStorage = canEncryptStorage
	this.CanCertify = canCertify
	this.CreatedAt = createdAt
	this.ExpiresAt = expiresAt
	this.Revoked = revoked
	this.RawKey = rawKey
	return &this
}

// NewGpgKeyWithDefaults instantiates a new GpgKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGpgKeyWithDefaults() *GpgKey {
	this := GpgKey{}
	return &this
}

// GetId returns the Id field value
func (o *GpgKey) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GpgKey) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GpgKey) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GpgKey) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GpgKey) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *GpgKey) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *GpgKey) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *GpgKey) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *GpgKey) UnsetName() {
	o.Name.Unset()
}

// GetPrimaryKeyId returns the PrimaryKeyId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *GpgKey) GetPrimaryKeyId() int32 {
	if o == nil || o.PrimaryKeyId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.PrimaryKeyId.Get()
}

// GetPrimaryKeyIdOk returns a tuple with the PrimaryKeyId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GpgKey) GetPrimaryKeyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryKeyId.Get(), o.PrimaryKeyId.IsSet()
}

// SetPrimaryKeyId sets field value
func (o *GpgKey) SetPrimaryKeyId(v int32) {
	o.PrimaryKeyId.Set(&v)
}

// GetKeyId returns the KeyId field value
func (o *GpgKey) GetKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value
// and a boolean to check if the value has been set.
func (o *GpgKey) GetKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyId, true
}

// SetKeyId sets field value
func (o *GpgKey) SetKeyId(v string) {
	o.KeyId = v
}

// GetPublicKey returns the PublicKey field value
func (o *GpgKey) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *GpgKey) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *GpgKey) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetEmails returns the Emails field value
func (o *GpgKey) GetEmails() []GpgKeyEmailsInner {
	if o == nil {
		var ret []GpgKeyEmailsInner
		return ret
	}

	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value
// and a boolean to check if the value has been set.
func (o *GpgKey) GetEmailsOk() ([]GpgKeyEmailsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Emails, true
}

// SetEmails sets field value
func (o *GpgKey) SetEmails(v []GpgKeyEmailsInner) {
	o.Emails = v
}

// GetSubkeys returns the Subkeys field value
func (o *GpgKey) GetSubkeys() []GpgKeySubkeysInner {
	if o == nil {
		var ret []GpgKeySubkeysInner
		return ret
	}

	return o.Subkeys
}

// GetSubkeysOk returns a tuple with the Subkeys field value
// and a boolean to check if the value has been set.
func (o *GpgKey) GetSubkeysOk() ([]GpgKeySubkeysInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subkeys, true
}

// SetSubkeys sets field value
func (o *GpgKey) SetSubkeys(v []GpgKeySubkeysInner) {
	o.Subkeys = v
}

// GetCanSign returns the CanSign field value
func (o *GpgKey) GetCanSign() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanSign
}

// GetCanSignOk returns a tuple with the CanSign field value
// and a boolean to check if the value has been set.
func (o *GpgKey) GetCanSignOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanSign, true
}

// SetCanSign sets field value
func (o *GpgKey) SetCanSign(v bool) {
	o.CanSign = v
}

// GetCanEncryptComms returns the CanEncryptComms field value
func (o *GpgKey) GetCanEncryptComms() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanEncryptComms
}

// GetCanEncryptCommsOk returns a tuple with the CanEncryptComms field value
// and a boolean to check if the value has been set.
func (o *GpgKey) GetCanEncryptCommsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanEncryptComms, true
}

// SetCanEncryptComms sets field value
func (o *GpgKey) SetCanEncryptComms(v bool) {
	o.CanEncryptComms = v
}

// GetCanEncryptStorage returns the CanEncryptStorage field value
func (o *GpgKey) GetCanEncryptStorage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanEncryptStorage
}

// GetCanEncryptStorageOk returns a tuple with the CanEncryptStorage field value
// and a boolean to check if the value has been set.
func (o *GpgKey) GetCanEncryptStorageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanEncryptStorage, true
}

// SetCanEncryptStorage sets field value
func (o *GpgKey) SetCanEncryptStorage(v bool) {
	o.CanEncryptStorage = v
}

// GetCanCertify returns the CanCertify field value
func (o *GpgKey) GetCanCertify() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanCertify
}

// GetCanCertifyOk returns a tuple with the CanCertify field value
// and a boolean to check if the value has been set.
func (o *GpgKey) GetCanCertifyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanCertify, true
}

// SetCanCertify sets field value
func (o *GpgKey) SetCanCertify(v bool) {
	o.CanCertify = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GpgKey) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GpgKey) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GpgKey) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetExpiresAt returns the ExpiresAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *GpgKey) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GpgKey) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// SetExpiresAt sets field value
func (o *GpgKey) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}

// GetRevoked returns the Revoked field value
func (o *GpgKey) GetRevoked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value
// and a boolean to check if the value has been set.
func (o *GpgKey) GetRevokedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revoked, true
}

// SetRevoked sets field value
func (o *GpgKey) SetRevoked(v bool) {
	o.Revoked = v
}

// GetRawKey returns the RawKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GpgKey) GetRawKey() string {
	if o == nil || o.RawKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.RawKey.Get()
}

// GetRawKeyOk returns a tuple with the RawKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GpgKey) GetRawKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RawKey.Get(), o.RawKey.IsSet()
}

// SetRawKey sets field value
func (o *GpgKey) SetRawKey(v string) {
	o.RawKey.Set(&v)
}

func (o GpgKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["primary_key_id"] = o.PrimaryKeyId.Get()
	}
	if true {
		toSerialize["key_id"] = o.KeyId
	}
	if true {
		toSerialize["public_key"] = o.PublicKey
	}
	if true {
		toSerialize["emails"] = o.Emails
	}
	if true {
		toSerialize["subkeys"] = o.Subkeys
	}
	if true {
		toSerialize["can_sign"] = o.CanSign
	}
	if true {
		toSerialize["can_encrypt_comms"] = o.CanEncryptComms
	}
	if true {
		toSerialize["can_encrypt_storage"] = o.CanEncryptStorage
	}
	if true {
		toSerialize["can_certify"] = o.CanCertify
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	if true {
		toSerialize["revoked"] = o.Revoked
	}
	if true {
		toSerialize["raw_key"] = o.RawKey.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGpgKey struct {
	value *GpgKey
	isSet bool
}

func (v NullableGpgKey) Get() *GpgKey {
	return v.value
}

func (v *NullableGpgKey) Set(val *GpgKey) {
	v.value = val
	v.isSet = true
}

func (v NullableGpgKey) IsSet() bool {
	return v.isSet
}

func (v *NullableGpgKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGpgKey(val *GpgKey) *NullableGpgKey {
	return &NullableGpgKey{value: val, isSet: true}
}

func (v NullableGpgKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGpgKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


