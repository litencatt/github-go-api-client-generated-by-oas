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

// ActionsCreateOrUpdateOrgSecretRequest struct for ActionsCreateOrUpdateOrgSecretRequest
type ActionsCreateOrUpdateOrgSecretRequest struct {
	// Value for your secret, encrypted with [LibSodium](https://libsodium.gitbook.io/doc/bindings_for_other_languages) using the public key retrieved from the [Get an organization public key](https://docs.github.com/rest/reference/actions#get-an-organization-public-key) endpoint.
	EncryptedValue *string `json:"encrypted_value,omitempty"`
	// ID of the key you used to encrypt the secret.
	KeyId *string `json:"key_id,omitempty"`
	// Which type of organization repositories have access to the organization secret. `selected` means only the repositories specified by `selected_repository_ids` can access the secret.
	Visibility string `json:"visibility"`
	// An array of repository ids that can access the organization secret. You can only provide a list of repository ids when the `visibility` is set to `selected`. You can manage the list of selected repositories using the [List selected repositories for an organization secret](https://docs.github.com/rest/reference/actions#list-selected-repositories-for-an-organization-secret), [Set selected repositories for an organization secret](https://docs.github.com/rest/reference/actions#set-selected-repositories-for-an-organization-secret), and [Remove selected repository from an organization secret](https://docs.github.com/rest/reference/actions#remove-selected-repository-from-an-organization-secret) endpoints.
	SelectedRepositoryIds []int32 `json:"selected_repository_ids,omitempty"`
}

// NewActionsCreateOrUpdateOrgSecretRequest instantiates a new ActionsCreateOrUpdateOrgSecretRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionsCreateOrUpdateOrgSecretRequest(visibility string) *ActionsCreateOrUpdateOrgSecretRequest {
	this := ActionsCreateOrUpdateOrgSecretRequest{}
	this.Visibility = visibility
	return &this
}

// NewActionsCreateOrUpdateOrgSecretRequestWithDefaults instantiates a new ActionsCreateOrUpdateOrgSecretRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsCreateOrUpdateOrgSecretRequestWithDefaults() *ActionsCreateOrUpdateOrgSecretRequest {
	this := ActionsCreateOrUpdateOrgSecretRequest{}
	return &this
}

// GetEncryptedValue returns the EncryptedValue field value if set, zero value otherwise.
func (o *ActionsCreateOrUpdateOrgSecretRequest) GetEncryptedValue() string {
	if o == nil || o.EncryptedValue == nil {
		var ret string
		return ret
	}
	return *o.EncryptedValue
}

// GetEncryptedValueOk returns a tuple with the EncryptedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsCreateOrUpdateOrgSecretRequest) GetEncryptedValueOk() (*string, bool) {
	if o == nil || o.EncryptedValue == nil {
		return nil, false
	}
	return o.EncryptedValue, true
}

// HasEncryptedValue returns a boolean if a field has been set.
func (o *ActionsCreateOrUpdateOrgSecretRequest) HasEncryptedValue() bool {
	if o != nil && o.EncryptedValue != nil {
		return true
	}

	return false
}

// SetEncryptedValue gets a reference to the given string and assigns it to the EncryptedValue field.
func (o *ActionsCreateOrUpdateOrgSecretRequest) SetEncryptedValue(v string) {
	o.EncryptedValue = &v
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *ActionsCreateOrUpdateOrgSecretRequest) GetKeyId() string {
	if o == nil || o.KeyId == nil {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsCreateOrUpdateOrgSecretRequest) GetKeyIdOk() (*string, bool) {
	if o == nil || o.KeyId == nil {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *ActionsCreateOrUpdateOrgSecretRequest) HasKeyId() bool {
	if o != nil && o.KeyId != nil {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *ActionsCreateOrUpdateOrgSecretRequest) SetKeyId(v string) {
	o.KeyId = &v
}

// GetVisibility returns the Visibility field value
func (o *ActionsCreateOrUpdateOrgSecretRequest) GetVisibility() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *ActionsCreateOrUpdateOrgSecretRequest) GetVisibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *ActionsCreateOrUpdateOrgSecretRequest) SetVisibility(v string) {
	o.Visibility = v
}

// GetSelectedRepositoryIds returns the SelectedRepositoryIds field value if set, zero value otherwise.
func (o *ActionsCreateOrUpdateOrgSecretRequest) GetSelectedRepositoryIds() []int32 {
	if o == nil || o.SelectedRepositoryIds == nil {
		var ret []int32
		return ret
	}
	return o.SelectedRepositoryIds
}

// GetSelectedRepositoryIdsOk returns a tuple with the SelectedRepositoryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsCreateOrUpdateOrgSecretRequest) GetSelectedRepositoryIdsOk() ([]int32, bool) {
	if o == nil || o.SelectedRepositoryIds == nil {
		return nil, false
	}
	return o.SelectedRepositoryIds, true
}

// HasSelectedRepositoryIds returns a boolean if a field has been set.
func (o *ActionsCreateOrUpdateOrgSecretRequest) HasSelectedRepositoryIds() bool {
	if o != nil && o.SelectedRepositoryIds != nil {
		return true
	}

	return false
}

// SetSelectedRepositoryIds gets a reference to the given []int32 and assigns it to the SelectedRepositoryIds field.
func (o *ActionsCreateOrUpdateOrgSecretRequest) SetSelectedRepositoryIds(v []int32) {
	o.SelectedRepositoryIds = v
}

func (o ActionsCreateOrUpdateOrgSecretRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EncryptedValue != nil {
		toSerialize["encrypted_value"] = o.EncryptedValue
	}
	if o.KeyId != nil {
		toSerialize["key_id"] = o.KeyId
	}
	if true {
		toSerialize["visibility"] = o.Visibility
	}
	if o.SelectedRepositoryIds != nil {
		toSerialize["selected_repository_ids"] = o.SelectedRepositoryIds
	}
	return json.Marshal(toSerialize)
}

type NullableActionsCreateOrUpdateOrgSecretRequest struct {
	value *ActionsCreateOrUpdateOrgSecretRequest
	isSet bool
}

func (v NullableActionsCreateOrUpdateOrgSecretRequest) Get() *ActionsCreateOrUpdateOrgSecretRequest {
	return v.value
}

func (v *NullableActionsCreateOrUpdateOrgSecretRequest) Set(val *ActionsCreateOrUpdateOrgSecretRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableActionsCreateOrUpdateOrgSecretRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableActionsCreateOrUpdateOrgSecretRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionsCreateOrUpdateOrgSecretRequest(val *ActionsCreateOrUpdateOrgSecretRequest) *NullableActionsCreateOrUpdateOrgSecretRequest {
	return &NullableActionsCreateOrUpdateOrgSecretRequest{value: val, isSet: true}
}

func (v NullableActionsCreateOrUpdateOrgSecretRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionsCreateOrUpdateOrgSecretRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

