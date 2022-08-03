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

// ReposSetStatusCheckContextsRequestOneOf struct for ReposSetStatusCheckContextsRequestOneOf
type ReposSetStatusCheckContextsRequestOneOf struct {
	// contexts parameter
	Contexts []string `json:"contexts"`
}

// NewReposSetStatusCheckContextsRequestOneOf instantiates a new ReposSetStatusCheckContextsRequestOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReposSetStatusCheckContextsRequestOneOf(contexts []string) *ReposSetStatusCheckContextsRequestOneOf {
	this := ReposSetStatusCheckContextsRequestOneOf{}
	this.Contexts = contexts
	return &this
}

// NewReposSetStatusCheckContextsRequestOneOfWithDefaults instantiates a new ReposSetStatusCheckContextsRequestOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReposSetStatusCheckContextsRequestOneOfWithDefaults() *ReposSetStatusCheckContextsRequestOneOf {
	this := ReposSetStatusCheckContextsRequestOneOf{}
	return &this
}

// GetContexts returns the Contexts field value
func (o *ReposSetStatusCheckContextsRequestOneOf) GetContexts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Contexts
}

// GetContextsOk returns a tuple with the Contexts field value
// and a boolean to check if the value has been set.
func (o *ReposSetStatusCheckContextsRequestOneOf) GetContextsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contexts, true
}

// SetContexts sets field value
func (o *ReposSetStatusCheckContextsRequestOneOf) SetContexts(v []string) {
	o.Contexts = v
}

func (o ReposSetStatusCheckContextsRequestOneOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contexts"] = o.Contexts
	}
	return json.Marshal(toSerialize)
}

type NullableReposSetStatusCheckContextsRequestOneOf struct {
	value *ReposSetStatusCheckContextsRequestOneOf
	isSet bool
}

func (v NullableReposSetStatusCheckContextsRequestOneOf) Get() *ReposSetStatusCheckContextsRequestOneOf {
	return v.value
}

func (v *NullableReposSetStatusCheckContextsRequestOneOf) Set(val *ReposSetStatusCheckContextsRequestOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReposSetStatusCheckContextsRequestOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReposSetStatusCheckContextsRequestOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReposSetStatusCheckContextsRequestOneOf(val *ReposSetStatusCheckContextsRequestOneOf) *NullableReposSetStatusCheckContextsRequestOneOf {
	return &NullableReposSetStatusCheckContextsRequestOneOf{value: val, isSet: true}
}

func (v NullableReposSetStatusCheckContextsRequestOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReposSetStatusCheckContextsRequestOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


