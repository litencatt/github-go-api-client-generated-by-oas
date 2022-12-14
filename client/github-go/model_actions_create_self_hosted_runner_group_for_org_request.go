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

// ActionsCreateSelfHostedRunnerGroupForOrgRequest struct for ActionsCreateSelfHostedRunnerGroupForOrgRequest
type ActionsCreateSelfHostedRunnerGroupForOrgRequest struct {
	// Name of the runner group.
	Name string `json:"name"`
	// Visibility of a runner group. You can select all repositories, select individual repositories, or limit access to private repositories.
	Visibility *string `json:"visibility,omitempty"`
	// List of repository IDs that can access the runner group.
	SelectedRepositoryIds []int32 `json:"selected_repository_ids,omitempty"`
	// List of runner IDs to add to the runner group.
	Runners []int32 `json:"runners,omitempty"`
	// Whether the runner group can be used by `public` repositories.
	AllowsPublicRepositories *bool `json:"allows_public_repositories,omitempty"`
	// If `true`, the runner group will be restricted to running only the workflows specified in the `selected_workflows` array.
	RestrictedToWorkflows *bool `json:"restricted_to_workflows,omitempty"`
	// List of workflows the runner group should be allowed to run. This setting will be ignored unless `restricted_to_workflows` is set to `true`.
	SelectedWorkflows []string `json:"selected_workflows,omitempty"`
}

// NewActionsCreateSelfHostedRunnerGroupForOrgRequest instantiates a new ActionsCreateSelfHostedRunnerGroupForOrgRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionsCreateSelfHostedRunnerGroupForOrgRequest(name string) *ActionsCreateSelfHostedRunnerGroupForOrgRequest {
	this := ActionsCreateSelfHostedRunnerGroupForOrgRequest{}
	this.Name = name
	var visibility string = "all"
	this.Visibility = &visibility
	var allowsPublicRepositories bool = false
	this.AllowsPublicRepositories = &allowsPublicRepositories
	var restrictedToWorkflows bool = false
	this.RestrictedToWorkflows = &restrictedToWorkflows
	return &this
}

// NewActionsCreateSelfHostedRunnerGroupForOrgRequestWithDefaults instantiates a new ActionsCreateSelfHostedRunnerGroupForOrgRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsCreateSelfHostedRunnerGroupForOrgRequestWithDefaults() *ActionsCreateSelfHostedRunnerGroupForOrgRequest {
	this := ActionsCreateSelfHostedRunnerGroupForOrgRequest{}
	var visibility string = "all"
	this.Visibility = &visibility
	var allowsPublicRepositories bool = false
	this.AllowsPublicRepositories = &allowsPublicRepositories
	var restrictedToWorkflows bool = false
	this.RestrictedToWorkflows = &restrictedToWorkflows
	return &this
}

// GetName returns the Name field value
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) SetName(v string) {
	o.Name = v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) SetVisibility(v string) {
	o.Visibility = &v
}

// GetSelectedRepositoryIds returns the SelectedRepositoryIds field value if set, zero value otherwise.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetSelectedRepositoryIds() []int32 {
	if o == nil || o.SelectedRepositoryIds == nil {
		var ret []int32
		return ret
	}
	return o.SelectedRepositoryIds
}

// GetSelectedRepositoryIdsOk returns a tuple with the SelectedRepositoryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetSelectedRepositoryIdsOk() ([]int32, bool) {
	if o == nil || o.SelectedRepositoryIds == nil {
		return nil, false
	}
	return o.SelectedRepositoryIds, true
}

// HasSelectedRepositoryIds returns a boolean if a field has been set.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) HasSelectedRepositoryIds() bool {
	if o != nil && o.SelectedRepositoryIds != nil {
		return true
	}

	return false
}

// SetSelectedRepositoryIds gets a reference to the given []int32 and assigns it to the SelectedRepositoryIds field.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) SetSelectedRepositoryIds(v []int32) {
	o.SelectedRepositoryIds = v
}

// GetRunners returns the Runners field value if set, zero value otherwise.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetRunners() []int32 {
	if o == nil || o.Runners == nil {
		var ret []int32
		return ret
	}
	return o.Runners
}

// GetRunnersOk returns a tuple with the Runners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetRunnersOk() ([]int32, bool) {
	if o == nil || o.Runners == nil {
		return nil, false
	}
	return o.Runners, true
}

// HasRunners returns a boolean if a field has been set.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) HasRunners() bool {
	if o != nil && o.Runners != nil {
		return true
	}

	return false
}

// SetRunners gets a reference to the given []int32 and assigns it to the Runners field.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) SetRunners(v []int32) {
	o.Runners = v
}

// GetAllowsPublicRepositories returns the AllowsPublicRepositories field value if set, zero value otherwise.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetAllowsPublicRepositories() bool {
	if o == nil || o.AllowsPublicRepositories == nil {
		var ret bool
		return ret
	}
	return *o.AllowsPublicRepositories
}

// GetAllowsPublicRepositoriesOk returns a tuple with the AllowsPublicRepositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetAllowsPublicRepositoriesOk() (*bool, bool) {
	if o == nil || o.AllowsPublicRepositories == nil {
		return nil, false
	}
	return o.AllowsPublicRepositories, true
}

// HasAllowsPublicRepositories returns a boolean if a field has been set.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) HasAllowsPublicRepositories() bool {
	if o != nil && o.AllowsPublicRepositories != nil {
		return true
	}

	return false
}

// SetAllowsPublicRepositories gets a reference to the given bool and assigns it to the AllowsPublicRepositories field.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) SetAllowsPublicRepositories(v bool) {
	o.AllowsPublicRepositories = &v
}

// GetRestrictedToWorkflows returns the RestrictedToWorkflows field value if set, zero value otherwise.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetRestrictedToWorkflows() bool {
	if o == nil || o.RestrictedToWorkflows == nil {
		var ret bool
		return ret
	}
	return *o.RestrictedToWorkflows
}

// GetRestrictedToWorkflowsOk returns a tuple with the RestrictedToWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetRestrictedToWorkflowsOk() (*bool, bool) {
	if o == nil || o.RestrictedToWorkflows == nil {
		return nil, false
	}
	return o.RestrictedToWorkflows, true
}

// HasRestrictedToWorkflows returns a boolean if a field has been set.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) HasRestrictedToWorkflows() bool {
	if o != nil && o.RestrictedToWorkflows != nil {
		return true
	}

	return false
}

// SetRestrictedToWorkflows gets a reference to the given bool and assigns it to the RestrictedToWorkflows field.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) SetRestrictedToWorkflows(v bool) {
	o.RestrictedToWorkflows = &v
}

// GetSelectedWorkflows returns the SelectedWorkflows field value if set, zero value otherwise.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetSelectedWorkflows() []string {
	if o == nil || o.SelectedWorkflows == nil {
		var ret []string
		return ret
	}
	return o.SelectedWorkflows
}

// GetSelectedWorkflowsOk returns a tuple with the SelectedWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) GetSelectedWorkflowsOk() ([]string, bool) {
	if o == nil || o.SelectedWorkflows == nil {
		return nil, false
	}
	return o.SelectedWorkflows, true
}

// HasSelectedWorkflows returns a boolean if a field has been set.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) HasSelectedWorkflows() bool {
	if o != nil && o.SelectedWorkflows != nil {
		return true
	}

	return false
}

// SetSelectedWorkflows gets a reference to the given []string and assigns it to the SelectedWorkflows field.
func (o *ActionsCreateSelfHostedRunnerGroupForOrgRequest) SetSelectedWorkflows(v []string) {
	o.SelectedWorkflows = v
}

func (o ActionsCreateSelfHostedRunnerGroupForOrgRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.SelectedRepositoryIds != nil {
		toSerialize["selected_repository_ids"] = o.SelectedRepositoryIds
	}
	if o.Runners != nil {
		toSerialize["runners"] = o.Runners
	}
	if o.AllowsPublicRepositories != nil {
		toSerialize["allows_public_repositories"] = o.AllowsPublicRepositories
	}
	if o.RestrictedToWorkflows != nil {
		toSerialize["restricted_to_workflows"] = o.RestrictedToWorkflows
	}
	if o.SelectedWorkflows != nil {
		toSerialize["selected_workflows"] = o.SelectedWorkflows
	}
	return json.Marshal(toSerialize)
}

type NullableActionsCreateSelfHostedRunnerGroupForOrgRequest struct {
	value *ActionsCreateSelfHostedRunnerGroupForOrgRequest
	isSet bool
}

func (v NullableActionsCreateSelfHostedRunnerGroupForOrgRequest) Get() *ActionsCreateSelfHostedRunnerGroupForOrgRequest {
	return v.value
}

func (v *NullableActionsCreateSelfHostedRunnerGroupForOrgRequest) Set(val *ActionsCreateSelfHostedRunnerGroupForOrgRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableActionsCreateSelfHostedRunnerGroupForOrgRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableActionsCreateSelfHostedRunnerGroupForOrgRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionsCreateSelfHostedRunnerGroupForOrgRequest(val *ActionsCreateSelfHostedRunnerGroupForOrgRequest) *NullableActionsCreateSelfHostedRunnerGroupForOrgRequest {
	return &NullableActionsCreateSelfHostedRunnerGroupForOrgRequest{value: val, isSet: true}
}

func (v NullableActionsCreateSelfHostedRunnerGroupForOrgRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionsCreateSelfHostedRunnerGroupForOrgRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


