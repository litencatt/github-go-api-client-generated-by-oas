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

// CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest struct for CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest
type CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest struct {
	// Value for your secret, encrypted with [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages) using the public key retrieved from the [Get the public key for the authenticated user](https://docs.github.com/rest/reference/codespaces#get-the-public-key-for-the-authenticated-user) endpoint.
	EncryptedValue *string `json:"encrypted_value,omitempty"`
	// ID of the key you used to encrypt the secret.
	KeyId string `json:"key_id"`
	// An array of repository ids that can access the user secret. You can manage the list of selected repositories using the [List selected repositories for a user secret](https://docs.github.com/rest/reference/codespaces#list-selected-repositories-for-a-user-secret), [Set selected repositories for a user secret](https://docs.github.com/rest/reference/codespaces#set-selected-repositories-for-a-user-secret), and [Remove a selected repository from a user secret](https://docs.github.com/rest/reference/codespaces#remove-a-selected-repository-from-a-user-secret) endpoints.
	SelectedRepositoryIds []string `json:"selected_repository_ids,omitempty"`
}

// NewCodespacesCreateOrUpdateSecretForAuthenticatedUserRequest instantiates a new CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodespacesCreateOrUpdateSecretForAuthenticatedUserRequest(keyId string) *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest {
	this := CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest{}
	this.KeyId = keyId
	return &this
}

// NewCodespacesCreateOrUpdateSecretForAuthenticatedUserRequestWithDefaults instantiates a new CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodespacesCreateOrUpdateSecretForAuthenticatedUserRequestWithDefaults() *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest {
	this := CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest{}
	return &this
}

// GetEncryptedValue returns the EncryptedValue field value if set, zero value otherwise.
func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) GetEncryptedValue() string {
	if o == nil || o.EncryptedValue == nil {
		var ret string
		return ret
	}
	return *o.EncryptedValue
}

// GetEncryptedValueOk returns a tuple with the EncryptedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) GetEncryptedValueOk() (*string, bool) {
	if o == nil || o.EncryptedValue == nil {
		return nil, false
	}
	return o.EncryptedValue, true
}

// HasEncryptedValue returns a boolean if a field has been set.
func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) HasEncryptedValue() bool {
	if o != nil && o.EncryptedValue != nil {
		return true
	}

	return false
}

// SetEncryptedValue gets a reference to the given string and assigns it to the EncryptedValue field.
func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) SetEncryptedValue(v string) {
	o.EncryptedValue = &v
}

// GetKeyId returns the KeyId field value
func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) GetKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value
// and a boolean to check if the value has been set.
func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) GetKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyId, true
}

// SetKeyId sets field value
func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) SetKeyId(v string) {
	o.KeyId = v
}

// GetSelectedRepositoryIds returns the SelectedRepositoryIds field value if set, zero value otherwise.
func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) GetSelectedRepositoryIds() []string {
	if o == nil || o.SelectedRepositoryIds == nil {
		var ret []string
		return ret
	}
	return o.SelectedRepositoryIds
}

// GetSelectedRepositoryIdsOk returns a tuple with the SelectedRepositoryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) GetSelectedRepositoryIdsOk() ([]string, bool) {
	if o == nil || o.SelectedRepositoryIds == nil {
		return nil, false
	}
	return o.SelectedRepositoryIds, true
}

// HasSelectedRepositoryIds returns a boolean if a field has been set.
func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) HasSelectedRepositoryIds() bool {
	if o != nil && o.SelectedRepositoryIds != nil {
		return true
	}

	return false
}

// SetSelectedRepositoryIds gets a reference to the given []string and assigns it to the SelectedRepositoryIds field.
func (o *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) SetSelectedRepositoryIds(v []string) {
	o.SelectedRepositoryIds = v
}

func (o CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EncryptedValue != nil {
		toSerialize["encrypted_value"] = o.EncryptedValue
	}
	if true {
		toSerialize["key_id"] = o.KeyId
	}
	if o.SelectedRepositoryIds != nil {
		toSerialize["selected_repository_ids"] = o.SelectedRepositoryIds
	}
	return json.Marshal(toSerialize)
}

type NullableCodespacesCreateOrUpdateSecretForAuthenticatedUserRequest struct {
	value *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest
	isSet bool
}

func (v NullableCodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) Get() *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest {
	return v.value
}

func (v *NullableCodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) Set(val *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodespacesCreateOrUpdateSecretForAuthenticatedUserRequest(val *CodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) *NullableCodespacesCreateOrUpdateSecretForAuthenticatedUserRequest {
	return &NullableCodespacesCreateOrUpdateSecretForAuthenticatedUserRequest{value: val, isSet: true}
}

func (v NullableCodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodespacesCreateOrUpdateSecretForAuthenticatedUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


