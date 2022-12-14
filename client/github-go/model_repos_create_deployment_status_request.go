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

// ReposCreateDeploymentStatusRequest struct for ReposCreateDeploymentStatusRequest
type ReposCreateDeploymentStatusRequest struct {
	// The state of the status. When you set a transient deployment to `inactive`, the deployment will be shown as `destroyed` in GitHub.
	State string `json:"state"`
	// The target URL to associate with this status. This URL should contain output to keep the user updated while the task is running or serve as historical information for what happened in the deployment. **Note:** It's recommended to use the `log_url` parameter, which replaces `target_url`.
	TargetUrl *string `json:"target_url,omitempty"`
	// The full URL of the deployment's output. This parameter replaces `target_url`. We will continue to accept `target_url` to support legacy uses, but we recommend replacing `target_url` with `log_url`. Setting `log_url` will automatically set `target_url` to the same value. Default: `\"\"`
	LogUrl *string `json:"log_url,omitempty"`
	// A short description of the status. The maximum description length is 140 characters.
	Description *string `json:"description,omitempty"`
	// Name for the target deployment environment, which can be changed when setting a deploy status. For example, `production`, `staging`, or `qa`.
	Environment *string `json:"environment,omitempty"`
	// Sets the URL for accessing your environment. Default: `\"\"`
	EnvironmentUrl *string `json:"environment_url,omitempty"`
	// Adds a new `inactive` status to all prior non-transient, non-production environment deployments with the same repository and `environment` name as the created status's deployment. An `inactive` status is only added to deployments that had a `success` state. Default: `true`
	AutoInactive *bool `json:"auto_inactive,omitempty"`
}

// NewReposCreateDeploymentStatusRequest instantiates a new ReposCreateDeploymentStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReposCreateDeploymentStatusRequest(state string) *ReposCreateDeploymentStatusRequest {
	this := ReposCreateDeploymentStatusRequest{}
	this.State = state
	var targetUrl string = ""
	this.TargetUrl = &targetUrl
	var logUrl string = ""
	this.LogUrl = &logUrl
	var description string = ""
	this.Description = &description
	var environmentUrl string = ""
	this.EnvironmentUrl = &environmentUrl
	return &this
}

// NewReposCreateDeploymentStatusRequestWithDefaults instantiates a new ReposCreateDeploymentStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReposCreateDeploymentStatusRequestWithDefaults() *ReposCreateDeploymentStatusRequest {
	this := ReposCreateDeploymentStatusRequest{}
	var targetUrl string = ""
	this.TargetUrl = &targetUrl
	var logUrl string = ""
	this.LogUrl = &logUrl
	var description string = ""
	this.Description = &description
	var environmentUrl string = ""
	this.EnvironmentUrl = &environmentUrl
	return &this
}

// GetState returns the State field value
func (o *ReposCreateDeploymentStatusRequest) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ReposCreateDeploymentStatusRequest) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ReposCreateDeploymentStatusRequest) SetState(v string) {
	o.State = v
}

// GetTargetUrl returns the TargetUrl field value if set, zero value otherwise.
func (o *ReposCreateDeploymentStatusRequest) GetTargetUrl() string {
	if o == nil || o.TargetUrl == nil {
		var ret string
		return ret
	}
	return *o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateDeploymentStatusRequest) GetTargetUrlOk() (*string, bool) {
	if o == nil || o.TargetUrl == nil {
		return nil, false
	}
	return o.TargetUrl, true
}

// HasTargetUrl returns a boolean if a field has been set.
func (o *ReposCreateDeploymentStatusRequest) HasTargetUrl() bool {
	if o != nil && o.TargetUrl != nil {
		return true
	}

	return false
}

// SetTargetUrl gets a reference to the given string and assigns it to the TargetUrl field.
func (o *ReposCreateDeploymentStatusRequest) SetTargetUrl(v string) {
	o.TargetUrl = &v
}

// GetLogUrl returns the LogUrl field value if set, zero value otherwise.
func (o *ReposCreateDeploymentStatusRequest) GetLogUrl() string {
	if o == nil || o.LogUrl == nil {
		var ret string
		return ret
	}
	return *o.LogUrl
}

// GetLogUrlOk returns a tuple with the LogUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateDeploymentStatusRequest) GetLogUrlOk() (*string, bool) {
	if o == nil || o.LogUrl == nil {
		return nil, false
	}
	return o.LogUrl, true
}

// HasLogUrl returns a boolean if a field has been set.
func (o *ReposCreateDeploymentStatusRequest) HasLogUrl() bool {
	if o != nil && o.LogUrl != nil {
		return true
	}

	return false
}

// SetLogUrl gets a reference to the given string and assigns it to the LogUrl field.
func (o *ReposCreateDeploymentStatusRequest) SetLogUrl(v string) {
	o.LogUrl = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ReposCreateDeploymentStatusRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateDeploymentStatusRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ReposCreateDeploymentStatusRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ReposCreateDeploymentStatusRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ReposCreateDeploymentStatusRequest) GetEnvironment() string {
	if o == nil || o.Environment == nil {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateDeploymentStatusRequest) GetEnvironmentOk() (*string, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ReposCreateDeploymentStatusRequest) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *ReposCreateDeploymentStatusRequest) SetEnvironment(v string) {
	o.Environment = &v
}

// GetEnvironmentUrl returns the EnvironmentUrl field value if set, zero value otherwise.
func (o *ReposCreateDeploymentStatusRequest) GetEnvironmentUrl() string {
	if o == nil || o.EnvironmentUrl == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentUrl
}

// GetEnvironmentUrlOk returns a tuple with the EnvironmentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateDeploymentStatusRequest) GetEnvironmentUrlOk() (*string, bool) {
	if o == nil || o.EnvironmentUrl == nil {
		return nil, false
	}
	return o.EnvironmentUrl, true
}

// HasEnvironmentUrl returns a boolean if a field has been set.
func (o *ReposCreateDeploymentStatusRequest) HasEnvironmentUrl() bool {
	if o != nil && o.EnvironmentUrl != nil {
		return true
	}

	return false
}

// SetEnvironmentUrl gets a reference to the given string and assigns it to the EnvironmentUrl field.
func (o *ReposCreateDeploymentStatusRequest) SetEnvironmentUrl(v string) {
	o.EnvironmentUrl = &v
}

// GetAutoInactive returns the AutoInactive field value if set, zero value otherwise.
func (o *ReposCreateDeploymentStatusRequest) GetAutoInactive() bool {
	if o == nil || o.AutoInactive == nil {
		var ret bool
		return ret
	}
	return *o.AutoInactive
}

// GetAutoInactiveOk returns a tuple with the AutoInactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposCreateDeploymentStatusRequest) GetAutoInactiveOk() (*bool, bool) {
	if o == nil || o.AutoInactive == nil {
		return nil, false
	}
	return o.AutoInactive, true
}

// HasAutoInactive returns a boolean if a field has been set.
func (o *ReposCreateDeploymentStatusRequest) HasAutoInactive() bool {
	if o != nil && o.AutoInactive != nil {
		return true
	}

	return false
}

// SetAutoInactive gets a reference to the given bool and assigns it to the AutoInactive field.
func (o *ReposCreateDeploymentStatusRequest) SetAutoInactive(v bool) {
	o.AutoInactive = &v
}

func (o ReposCreateDeploymentStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["state"] = o.State
	}
	if o.TargetUrl != nil {
		toSerialize["target_url"] = o.TargetUrl
	}
	if o.LogUrl != nil {
		toSerialize["log_url"] = o.LogUrl
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.EnvironmentUrl != nil {
		toSerialize["environment_url"] = o.EnvironmentUrl
	}
	if o.AutoInactive != nil {
		toSerialize["auto_inactive"] = o.AutoInactive
	}
	return json.Marshal(toSerialize)
}

type NullableReposCreateDeploymentStatusRequest struct {
	value *ReposCreateDeploymentStatusRequest
	isSet bool
}

func (v NullableReposCreateDeploymentStatusRequest) Get() *ReposCreateDeploymentStatusRequest {
	return v.value
}

func (v *NullableReposCreateDeploymentStatusRequest) Set(val *ReposCreateDeploymentStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReposCreateDeploymentStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReposCreateDeploymentStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReposCreateDeploymentStatusRequest(val *ReposCreateDeploymentStatusRequest) *NullableReposCreateDeploymentStatusRequest {
	return &NullableReposCreateDeploymentStatusRequest{value: val, isSet: true}
}

func (v NullableReposCreateDeploymentStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReposCreateDeploymentStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


