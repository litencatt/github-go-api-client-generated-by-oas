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

// MigrationsStartImportRequest struct for MigrationsStartImportRequest
type MigrationsStartImportRequest struct {
	// The URL of the originating repository.
	VcsUrl string `json:"vcs_url"`
	// The originating VCS type. Without this parameter, the import job will take additional time to detect the VCS type before beginning the import. This detection step will be reflected in the response.
	Vcs *string `json:"vcs,omitempty"`
	// If authentication is required, the username to provide to `vcs_url`.
	VcsUsername *string `json:"vcs_username,omitempty"`
	// If authentication is required, the password to provide to `vcs_url`.
	VcsPassword *string `json:"vcs_password,omitempty"`
	// For a tfvc import, the name of the project that is being imported.
	TfvcProject *string `json:"tfvc_project,omitempty"`
}

// NewMigrationsStartImportRequest instantiates a new MigrationsStartImportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrationsStartImportRequest(vcsUrl string) *MigrationsStartImportRequest {
	this := MigrationsStartImportRequest{}
	this.VcsUrl = vcsUrl
	return &this
}

// NewMigrationsStartImportRequestWithDefaults instantiates a new MigrationsStartImportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrationsStartImportRequestWithDefaults() *MigrationsStartImportRequest {
	this := MigrationsStartImportRequest{}
	return &this
}

// GetVcsUrl returns the VcsUrl field value
func (o *MigrationsStartImportRequest) GetVcsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VcsUrl
}

// GetVcsUrlOk returns a tuple with the VcsUrl field value
// and a boolean to check if the value has been set.
func (o *MigrationsStartImportRequest) GetVcsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VcsUrl, true
}

// SetVcsUrl sets field value
func (o *MigrationsStartImportRequest) SetVcsUrl(v string) {
	o.VcsUrl = v
}

// GetVcs returns the Vcs field value if set, zero value otherwise.
func (o *MigrationsStartImportRequest) GetVcs() string {
	if o == nil || o.Vcs == nil {
		var ret string
		return ret
	}
	return *o.Vcs
}

// GetVcsOk returns a tuple with the Vcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationsStartImportRequest) GetVcsOk() (*string, bool) {
	if o == nil || o.Vcs == nil {
		return nil, false
	}
	return o.Vcs, true
}

// HasVcs returns a boolean if a field has been set.
func (o *MigrationsStartImportRequest) HasVcs() bool {
	if o != nil && o.Vcs != nil {
		return true
	}

	return false
}

// SetVcs gets a reference to the given string and assigns it to the Vcs field.
func (o *MigrationsStartImportRequest) SetVcs(v string) {
	o.Vcs = &v
}

// GetVcsUsername returns the VcsUsername field value if set, zero value otherwise.
func (o *MigrationsStartImportRequest) GetVcsUsername() string {
	if o == nil || o.VcsUsername == nil {
		var ret string
		return ret
	}
	return *o.VcsUsername
}

// GetVcsUsernameOk returns a tuple with the VcsUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationsStartImportRequest) GetVcsUsernameOk() (*string, bool) {
	if o == nil || o.VcsUsername == nil {
		return nil, false
	}
	return o.VcsUsername, true
}

// HasVcsUsername returns a boolean if a field has been set.
func (o *MigrationsStartImportRequest) HasVcsUsername() bool {
	if o != nil && o.VcsUsername != nil {
		return true
	}

	return false
}

// SetVcsUsername gets a reference to the given string and assigns it to the VcsUsername field.
func (o *MigrationsStartImportRequest) SetVcsUsername(v string) {
	o.VcsUsername = &v
}

// GetVcsPassword returns the VcsPassword field value if set, zero value otherwise.
func (o *MigrationsStartImportRequest) GetVcsPassword() string {
	if o == nil || o.VcsPassword == nil {
		var ret string
		return ret
	}
	return *o.VcsPassword
}

// GetVcsPasswordOk returns a tuple with the VcsPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationsStartImportRequest) GetVcsPasswordOk() (*string, bool) {
	if o == nil || o.VcsPassword == nil {
		return nil, false
	}
	return o.VcsPassword, true
}

// HasVcsPassword returns a boolean if a field has been set.
func (o *MigrationsStartImportRequest) HasVcsPassword() bool {
	if o != nil && o.VcsPassword != nil {
		return true
	}

	return false
}

// SetVcsPassword gets a reference to the given string and assigns it to the VcsPassword field.
func (o *MigrationsStartImportRequest) SetVcsPassword(v string) {
	o.VcsPassword = &v
}

// GetTfvcProject returns the TfvcProject field value if set, zero value otherwise.
func (o *MigrationsStartImportRequest) GetTfvcProject() string {
	if o == nil || o.TfvcProject == nil {
		var ret string
		return ret
	}
	return *o.TfvcProject
}

// GetTfvcProjectOk returns a tuple with the TfvcProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrationsStartImportRequest) GetTfvcProjectOk() (*string, bool) {
	if o == nil || o.TfvcProject == nil {
		return nil, false
	}
	return o.TfvcProject, true
}

// HasTfvcProject returns a boolean if a field has been set.
func (o *MigrationsStartImportRequest) HasTfvcProject() bool {
	if o != nil && o.TfvcProject != nil {
		return true
	}

	return false
}

// SetTfvcProject gets a reference to the given string and assigns it to the TfvcProject field.
func (o *MigrationsStartImportRequest) SetTfvcProject(v string) {
	o.TfvcProject = &v
}

func (o MigrationsStartImportRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["vcs_url"] = o.VcsUrl
	}
	if o.Vcs != nil {
		toSerialize["vcs"] = o.Vcs
	}
	if o.VcsUsername != nil {
		toSerialize["vcs_username"] = o.VcsUsername
	}
	if o.VcsPassword != nil {
		toSerialize["vcs_password"] = o.VcsPassword
	}
	if o.TfvcProject != nil {
		toSerialize["tfvc_project"] = o.TfvcProject
	}
	return json.Marshal(toSerialize)
}

type NullableMigrationsStartImportRequest struct {
	value *MigrationsStartImportRequest
	isSet bool
}

func (v NullableMigrationsStartImportRequest) Get() *MigrationsStartImportRequest {
	return v.value
}

func (v *NullableMigrationsStartImportRequest) Set(val *MigrationsStartImportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrationsStartImportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrationsStartImportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrationsStartImportRequest(val *MigrationsStartImportRequest) *NullableMigrationsStartImportRequest {
	return &NullableMigrationsStartImportRequest{value: val, isSet: true}
}

func (v NullableMigrationsStartImportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrationsStartImportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

