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

// HookConfig struct for HookConfig
type HookConfig struct {
	Email *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
	Room *string `json:"room,omitempty"`
	Subdomain *string `json:"subdomain,omitempty"`
	// The URL to which the payloads will be delivered.
	Url *string `json:"url,omitempty"`
	InsecureSsl *WebhookConfigInsecureSsl `json:"insecure_ssl,omitempty"`
	// The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
	ContentType *string `json:"content_type,omitempty"`
	Digest *string `json:"digest,omitempty"`
	// If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
	Secret *string `json:"secret,omitempty"`
	Token *string `json:"token,omitempty"`
}

// NewHookConfig instantiates a new HookConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHookConfig() *HookConfig {
	this := HookConfig{}
	return &this
}

// NewHookConfigWithDefaults instantiates a new HookConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHookConfigWithDefaults() *HookConfig {
	this := HookConfig{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *HookConfig) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookConfig) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *HookConfig) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *HookConfig) SetEmail(v string) {
	o.Email = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *HookConfig) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookConfig) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *HookConfig) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *HookConfig) SetPassword(v string) {
	o.Password = &v
}

// GetRoom returns the Room field value if set, zero value otherwise.
func (o *HookConfig) GetRoom() string {
	if o == nil || o.Room == nil {
		var ret string
		return ret
	}
	return *o.Room
}

// GetRoomOk returns a tuple with the Room field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookConfig) GetRoomOk() (*string, bool) {
	if o == nil || o.Room == nil {
		return nil, false
	}
	return o.Room, true
}

// HasRoom returns a boolean if a field has been set.
func (o *HookConfig) HasRoom() bool {
	if o != nil && o.Room != nil {
		return true
	}

	return false
}

// SetRoom gets a reference to the given string and assigns it to the Room field.
func (o *HookConfig) SetRoom(v string) {
	o.Room = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *HookConfig) GetSubdomain() string {
	if o == nil || o.Subdomain == nil {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookConfig) GetSubdomainOk() (*string, bool) {
	if o == nil || o.Subdomain == nil {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *HookConfig) HasSubdomain() bool {
	if o != nil && o.Subdomain != nil {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *HookConfig) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *HookConfig) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookConfig) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *HookConfig) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *HookConfig) SetUrl(v string) {
	o.Url = &v
}

// GetInsecureSsl returns the InsecureSsl field value if set, zero value otherwise.
func (o *HookConfig) GetInsecureSsl() WebhookConfigInsecureSsl {
	if o == nil || o.InsecureSsl == nil {
		var ret WebhookConfigInsecureSsl
		return ret
	}
	return *o.InsecureSsl
}

// GetInsecureSslOk returns a tuple with the InsecureSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookConfig) GetInsecureSslOk() (*WebhookConfigInsecureSsl, bool) {
	if o == nil || o.InsecureSsl == nil {
		return nil, false
	}
	return o.InsecureSsl, true
}

// HasInsecureSsl returns a boolean if a field has been set.
func (o *HookConfig) HasInsecureSsl() bool {
	if o != nil && o.InsecureSsl != nil {
		return true
	}

	return false
}

// SetInsecureSsl gets a reference to the given WebhookConfigInsecureSsl and assigns it to the InsecureSsl field.
func (o *HookConfig) SetInsecureSsl(v WebhookConfigInsecureSsl) {
	o.InsecureSsl = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *HookConfig) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookConfig) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *HookConfig) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *HookConfig) SetContentType(v string) {
	o.ContentType = &v
}

// GetDigest returns the Digest field value if set, zero value otherwise.
func (o *HookConfig) GetDigest() string {
	if o == nil || o.Digest == nil {
		var ret string
		return ret
	}
	return *o.Digest
}

// GetDigestOk returns a tuple with the Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookConfig) GetDigestOk() (*string, bool) {
	if o == nil || o.Digest == nil {
		return nil, false
	}
	return o.Digest, true
}

// HasDigest returns a boolean if a field has been set.
func (o *HookConfig) HasDigest() bool {
	if o != nil && o.Digest != nil {
		return true
	}

	return false
}

// SetDigest gets a reference to the given string and assigns it to the Digest field.
func (o *HookConfig) SetDigest(v string) {
	o.Digest = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *HookConfig) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookConfig) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *HookConfig) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *HookConfig) SetSecret(v string) {
	o.Secret = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *HookConfig) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HookConfig) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *HookConfig) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *HookConfig) SetToken(v string) {
	o.Token = &v
}

func (o HookConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Room != nil {
		toSerialize["room"] = o.Room
	}
	if o.Subdomain != nil {
		toSerialize["subdomain"] = o.Subdomain
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.InsecureSsl != nil {
		toSerialize["insecure_ssl"] = o.InsecureSsl
	}
	if o.ContentType != nil {
		toSerialize["content_type"] = o.ContentType
	}
	if o.Digest != nil {
		toSerialize["digest"] = o.Digest
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableHookConfig struct {
	value *HookConfig
	isSet bool
}

func (v NullableHookConfig) Get() *HookConfig {
	return v.value
}

func (v *NullableHookConfig) Set(val *HookConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableHookConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableHookConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHookConfig(val *HookConfig) *NullableHookConfig {
	return &NullableHookConfig{value: val, isSet: true}
}

func (v NullableHookConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHookConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


