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

// ProtectedBranchRequiredStatusCheck Protected Branch Required Status Check
type ProtectedBranchRequiredStatusCheck struct {
	Url *string `json:"url,omitempty"`
	EnforcementLevel *string `json:"enforcement_level,omitempty"`
	Contexts []string `json:"contexts"`
	Checks []ProtectedBranchRequiredStatusCheckChecksInner `json:"checks"`
	ContextsUrl *string `json:"contexts_url,omitempty"`
	Strict *bool `json:"strict,omitempty"`
}

// NewProtectedBranchRequiredStatusCheck instantiates a new ProtectedBranchRequiredStatusCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtectedBranchRequiredStatusCheck(contexts []string, checks []ProtectedBranchRequiredStatusCheckChecksInner) *ProtectedBranchRequiredStatusCheck {
	this := ProtectedBranchRequiredStatusCheck{}
	this.Contexts = contexts
	this.Checks = checks
	return &this
}

// NewProtectedBranchRequiredStatusCheckWithDefaults instantiates a new ProtectedBranchRequiredStatusCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtectedBranchRequiredStatusCheckWithDefaults() *ProtectedBranchRequiredStatusCheck {
	this := ProtectedBranchRequiredStatusCheck{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ProtectedBranchRequiredStatusCheck) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectedBranchRequiredStatusCheck) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ProtectedBranchRequiredStatusCheck) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ProtectedBranchRequiredStatusCheck) SetUrl(v string) {
	o.Url = &v
}

// GetEnforcementLevel returns the EnforcementLevel field value if set, zero value otherwise.
func (o *ProtectedBranchRequiredStatusCheck) GetEnforcementLevel() string {
	if o == nil || o.EnforcementLevel == nil {
		var ret string
		return ret
	}
	return *o.EnforcementLevel
}

// GetEnforcementLevelOk returns a tuple with the EnforcementLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectedBranchRequiredStatusCheck) GetEnforcementLevelOk() (*string, bool) {
	if o == nil || o.EnforcementLevel == nil {
		return nil, false
	}
	return o.EnforcementLevel, true
}

// HasEnforcementLevel returns a boolean if a field has been set.
func (o *ProtectedBranchRequiredStatusCheck) HasEnforcementLevel() bool {
	if o != nil && o.EnforcementLevel != nil {
		return true
	}

	return false
}

// SetEnforcementLevel gets a reference to the given string and assigns it to the EnforcementLevel field.
func (o *ProtectedBranchRequiredStatusCheck) SetEnforcementLevel(v string) {
	o.EnforcementLevel = &v
}

// GetContexts returns the Contexts field value
func (o *ProtectedBranchRequiredStatusCheck) GetContexts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Contexts
}

// GetContextsOk returns a tuple with the Contexts field value
// and a boolean to check if the value has been set.
func (o *ProtectedBranchRequiredStatusCheck) GetContextsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contexts, true
}

// SetContexts sets field value
func (o *ProtectedBranchRequiredStatusCheck) SetContexts(v []string) {
	o.Contexts = v
}

// GetChecks returns the Checks field value
func (o *ProtectedBranchRequiredStatusCheck) GetChecks() []ProtectedBranchRequiredStatusCheckChecksInner {
	if o == nil {
		var ret []ProtectedBranchRequiredStatusCheckChecksInner
		return ret
	}

	return o.Checks
}

// GetChecksOk returns a tuple with the Checks field value
// and a boolean to check if the value has been set.
func (o *ProtectedBranchRequiredStatusCheck) GetChecksOk() ([]ProtectedBranchRequiredStatusCheckChecksInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Checks, true
}

// SetChecks sets field value
func (o *ProtectedBranchRequiredStatusCheck) SetChecks(v []ProtectedBranchRequiredStatusCheckChecksInner) {
	o.Checks = v
}

// GetContextsUrl returns the ContextsUrl field value if set, zero value otherwise.
func (o *ProtectedBranchRequiredStatusCheck) GetContextsUrl() string {
	if o == nil || o.ContextsUrl == nil {
		var ret string
		return ret
	}
	return *o.ContextsUrl
}

// GetContextsUrlOk returns a tuple with the ContextsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectedBranchRequiredStatusCheck) GetContextsUrlOk() (*string, bool) {
	if o == nil || o.ContextsUrl == nil {
		return nil, false
	}
	return o.ContextsUrl, true
}

// HasContextsUrl returns a boolean if a field has been set.
func (o *ProtectedBranchRequiredStatusCheck) HasContextsUrl() bool {
	if o != nil && o.ContextsUrl != nil {
		return true
	}

	return false
}

// SetContextsUrl gets a reference to the given string and assigns it to the ContextsUrl field.
func (o *ProtectedBranchRequiredStatusCheck) SetContextsUrl(v string) {
	o.ContextsUrl = &v
}

// GetStrict returns the Strict field value if set, zero value otherwise.
func (o *ProtectedBranchRequiredStatusCheck) GetStrict() bool {
	if o == nil || o.Strict == nil {
		var ret bool
		return ret
	}
	return *o.Strict
}

// GetStrictOk returns a tuple with the Strict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtectedBranchRequiredStatusCheck) GetStrictOk() (*bool, bool) {
	if o == nil || o.Strict == nil {
		return nil, false
	}
	return o.Strict, true
}

// HasStrict returns a boolean if a field has been set.
func (o *ProtectedBranchRequiredStatusCheck) HasStrict() bool {
	if o != nil && o.Strict != nil {
		return true
	}

	return false
}

// SetStrict gets a reference to the given bool and assigns it to the Strict field.
func (o *ProtectedBranchRequiredStatusCheck) SetStrict(v bool) {
	o.Strict = &v
}

func (o ProtectedBranchRequiredStatusCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.EnforcementLevel != nil {
		toSerialize["enforcement_level"] = o.EnforcementLevel
	}
	if true {
		toSerialize["contexts"] = o.Contexts
	}
	if true {
		toSerialize["checks"] = o.Checks
	}
	if o.ContextsUrl != nil {
		toSerialize["contexts_url"] = o.ContextsUrl
	}
	if o.Strict != nil {
		toSerialize["strict"] = o.Strict
	}
	return json.Marshal(toSerialize)
}

type NullableProtectedBranchRequiredStatusCheck struct {
	value *ProtectedBranchRequiredStatusCheck
	isSet bool
}

func (v NullableProtectedBranchRequiredStatusCheck) Get() *ProtectedBranchRequiredStatusCheck {
	return v.value
}

func (v *NullableProtectedBranchRequiredStatusCheck) Set(val *ProtectedBranchRequiredStatusCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableProtectedBranchRequiredStatusCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableProtectedBranchRequiredStatusCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtectedBranchRequiredStatusCheck(val *ProtectedBranchRequiredStatusCheck) *NullableProtectedBranchRequiredStatusCheck {
	return &NullableProtectedBranchRequiredStatusCheck{value: val, isSet: true}
}

func (v NullableProtectedBranchRequiredStatusCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtectedBranchRequiredStatusCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


