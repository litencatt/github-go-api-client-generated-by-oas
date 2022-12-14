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

// OrgsUpdateWebhookRequestConfig Key/value pairs to provide settings for this webhook. [These are defined below](https://docs.github.com/rest/reference/orgs#update-hook-config-params).
type OrgsUpdateWebhookRequestConfig struct {
	// The URL to which the payloads will be delivered.
	Url string `json:"url"`
	// The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
	ContentType *string `json:"content_type,omitempty"`
	// If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
	Secret *string `json:"secret,omitempty"`
	InsecureSsl *WebhookConfigInsecureSsl `json:"insecure_ssl,omitempty"`
}

// NewOrgsUpdateWebhookRequestConfig instantiates a new OrgsUpdateWebhookRequestConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgsUpdateWebhookRequestConfig(url string) *OrgsUpdateWebhookRequestConfig {
	this := OrgsUpdateWebhookRequestConfig{}
	this.Url = url
	return &this
}

// NewOrgsUpdateWebhookRequestConfigWithDefaults instantiates a new OrgsUpdateWebhookRequestConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgsUpdateWebhookRequestConfigWithDefaults() *OrgsUpdateWebhookRequestConfig {
	this := OrgsUpdateWebhookRequestConfig{}
	return &this
}

// GetUrl returns the Url field value
func (o *OrgsUpdateWebhookRequestConfig) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *OrgsUpdateWebhookRequestConfig) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *OrgsUpdateWebhookRequestConfig) SetUrl(v string) {
	o.Url = v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *OrgsUpdateWebhookRequestConfig) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgsUpdateWebhookRequestConfig) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *OrgsUpdateWebhookRequestConfig) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *OrgsUpdateWebhookRequestConfig) SetContentType(v string) {
	o.ContentType = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *OrgsUpdateWebhookRequestConfig) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgsUpdateWebhookRequestConfig) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *OrgsUpdateWebhookRequestConfig) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *OrgsUpdateWebhookRequestConfig) SetSecret(v string) {
	o.Secret = &v
}

// GetInsecureSsl returns the InsecureSsl field value if set, zero value otherwise.
func (o *OrgsUpdateWebhookRequestConfig) GetInsecureSsl() WebhookConfigInsecureSsl {
	if o == nil || o.InsecureSsl == nil {
		var ret WebhookConfigInsecureSsl
		return ret
	}
	return *o.InsecureSsl
}

// GetInsecureSslOk returns a tuple with the InsecureSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgsUpdateWebhookRequestConfig) GetInsecureSslOk() (*WebhookConfigInsecureSsl, bool) {
	if o == nil || o.InsecureSsl == nil {
		return nil, false
	}
	return o.InsecureSsl, true
}

// HasInsecureSsl returns a boolean if a field has been set.
func (o *OrgsUpdateWebhookRequestConfig) HasInsecureSsl() bool {
	if o != nil && o.InsecureSsl != nil {
		return true
	}

	return false
}

// SetInsecureSsl gets a reference to the given WebhookConfigInsecureSsl and assigns it to the InsecureSsl field.
func (o *OrgsUpdateWebhookRequestConfig) SetInsecureSsl(v WebhookConfigInsecureSsl) {
	o.InsecureSsl = &v
}

func (o OrgsUpdateWebhookRequestConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.ContentType != nil {
		toSerialize["content_type"] = o.ContentType
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.InsecureSsl != nil {
		toSerialize["insecure_ssl"] = o.InsecureSsl
	}
	return json.Marshal(toSerialize)
}

type NullableOrgsUpdateWebhookRequestConfig struct {
	value *OrgsUpdateWebhookRequestConfig
	isSet bool
}

func (v NullableOrgsUpdateWebhookRequestConfig) Get() *OrgsUpdateWebhookRequestConfig {
	return v.value
}

func (v *NullableOrgsUpdateWebhookRequestConfig) Set(val *OrgsUpdateWebhookRequestConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgsUpdateWebhookRequestConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgsUpdateWebhookRequestConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgsUpdateWebhookRequestConfig(val *OrgsUpdateWebhookRequestConfig) *NullableOrgsUpdateWebhookRequestConfig {
	return &NullableOrgsUpdateWebhookRequestConfig{value: val, isSet: true}
}

func (v NullableOrgsUpdateWebhookRequestConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgsUpdateWebhookRequestConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


