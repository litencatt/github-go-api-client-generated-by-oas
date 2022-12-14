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

// ActionsOidcCustomIssuerPolicyForEnterprise struct for ActionsOidcCustomIssuerPolicyForEnterprise
type ActionsOidcCustomIssuerPolicyForEnterprise struct {
	// Whether the enterprise customer requested a custom issuer URL.
	IncludeEnterpriseSlug *bool `json:"include_enterprise_slug,omitempty"`
}

// NewActionsOidcCustomIssuerPolicyForEnterprise instantiates a new ActionsOidcCustomIssuerPolicyForEnterprise object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionsOidcCustomIssuerPolicyForEnterprise() *ActionsOidcCustomIssuerPolicyForEnterprise {
	this := ActionsOidcCustomIssuerPolicyForEnterprise{}
	return &this
}

// NewActionsOidcCustomIssuerPolicyForEnterpriseWithDefaults instantiates a new ActionsOidcCustomIssuerPolicyForEnterprise object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsOidcCustomIssuerPolicyForEnterpriseWithDefaults() *ActionsOidcCustomIssuerPolicyForEnterprise {
	this := ActionsOidcCustomIssuerPolicyForEnterprise{}
	return &this
}

// GetIncludeEnterpriseSlug returns the IncludeEnterpriseSlug field value if set, zero value otherwise.
func (o *ActionsOidcCustomIssuerPolicyForEnterprise) GetIncludeEnterpriseSlug() bool {
	if o == nil || o.IncludeEnterpriseSlug == nil {
		var ret bool
		return ret
	}
	return *o.IncludeEnterpriseSlug
}

// GetIncludeEnterpriseSlugOk returns a tuple with the IncludeEnterpriseSlug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsOidcCustomIssuerPolicyForEnterprise) GetIncludeEnterpriseSlugOk() (*bool, bool) {
	if o == nil || o.IncludeEnterpriseSlug == nil {
		return nil, false
	}
	return o.IncludeEnterpriseSlug, true
}

// HasIncludeEnterpriseSlug returns a boolean if a field has been set.
func (o *ActionsOidcCustomIssuerPolicyForEnterprise) HasIncludeEnterpriseSlug() bool {
	if o != nil && o.IncludeEnterpriseSlug != nil {
		return true
	}

	return false
}

// SetIncludeEnterpriseSlug gets a reference to the given bool and assigns it to the IncludeEnterpriseSlug field.
func (o *ActionsOidcCustomIssuerPolicyForEnterprise) SetIncludeEnterpriseSlug(v bool) {
	o.IncludeEnterpriseSlug = &v
}

func (o ActionsOidcCustomIssuerPolicyForEnterprise) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IncludeEnterpriseSlug != nil {
		toSerialize["include_enterprise_slug"] = o.IncludeEnterpriseSlug
	}
	return json.Marshal(toSerialize)
}

type NullableActionsOidcCustomIssuerPolicyForEnterprise struct {
	value *ActionsOidcCustomIssuerPolicyForEnterprise
	isSet bool
}

func (v NullableActionsOidcCustomIssuerPolicyForEnterprise) Get() *ActionsOidcCustomIssuerPolicyForEnterprise {
	return v.value
}

func (v *NullableActionsOidcCustomIssuerPolicyForEnterprise) Set(val *ActionsOidcCustomIssuerPolicyForEnterprise) {
	v.value = val
	v.isSet = true
}

func (v NullableActionsOidcCustomIssuerPolicyForEnterprise) IsSet() bool {
	return v.isSet
}

func (v *NullableActionsOidcCustomIssuerPolicyForEnterprise) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionsOidcCustomIssuerPolicyForEnterprise(val *ActionsOidcCustomIssuerPolicyForEnterprise) *NullableActionsOidcCustomIssuerPolicyForEnterprise {
	return &NullableActionsOidcCustomIssuerPolicyForEnterprise{value: val, isSet: true}
}

func (v NullableActionsOidcCustomIssuerPolicyForEnterprise) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionsOidcCustomIssuerPolicyForEnterprise) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


