/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"encoding/json"
	"time"
)

// Codespace A codespace.
type Codespace struct {
	Id int32 `json:"id"`
	// Automatically generated name of this codespace.
	Name string `json:"name"`
	// Display name for this codespace.
	DisplayName NullableString `json:"display_name,omitempty"`
	// UUID identifying this codespace's environment.
	EnvironmentId NullableString `json:"environment_id"`
	Owner SimpleUser `json:"owner"`
	BillableOwner SimpleUser `json:"billable_owner"`
	Repository MinimalRepository `json:"repository"`
	Machine NullableNullableCodespaceMachine `json:"machine"`
	// Path to devcontainer.json from repo root used to create Codespace.
	DevcontainerPath NullableString `json:"devcontainer_path,omitempty"`
	// Whether the codespace was created from a prebuild.
	Prebuild NullableBool `json:"prebuild"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Last known time this codespace was started.
	LastUsedAt time.Time `json:"last_used_at"`
	// State of this codespace.
	State string `json:"state"`
	// API URL for this codespace.
	Url string `json:"url"`
	GitStatus CodespaceGitStatus `json:"git_status"`
	// The Azure region where this codespace is located.
	Location string `json:"location"`
	// The number of minutes of inactivity after which this codespace will be automatically stopped.
	IdleTimeoutMinutes NullableInt32 `json:"idle_timeout_minutes"`
	// URL to access this codespace on the web.
	WebUrl string `json:"web_url"`
	// API URL to access available alternate machine types for this codespace.
	MachinesUrl string `json:"machines_url"`
	// API URL to start this codespace.
	StartUrl string `json:"start_url"`
	// API URL to stop this codespace.
	StopUrl string `json:"stop_url"`
	// API URL for the Pull Request associated with this codespace, if any.
	PullsUrl NullableString `json:"pulls_url"`
	RecentFolders []string `json:"recent_folders"`
	RuntimeConstraints *CodespaceRuntimeConstraints `json:"runtime_constraints,omitempty"`
	// Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
	PendingOperation NullableBool `json:"pending_operation,omitempty"`
	// Text to show user when codespace is disabled by a pending operation
	PendingOperationDisabledReason NullableString `json:"pending_operation_disabled_reason,omitempty"`
	// Text to show user when codespace idle timeout minutes has been overriden by an organization policy
	IdleTimeoutNotice NullableString `json:"idle_timeout_notice,omitempty"`
}

// NewCodespace instantiates a new Codespace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodespace(id int32, name string, environmentId NullableString, owner SimpleUser, billableOwner SimpleUser, repository MinimalRepository, machine NullableNullableCodespaceMachine, prebuild NullableBool, createdAt time.Time, updatedAt time.Time, lastUsedAt time.Time, state string, url string, gitStatus CodespaceGitStatus, location string, idleTimeoutMinutes NullableInt32, webUrl string, machinesUrl string, startUrl string, stopUrl string, pullsUrl NullableString, recentFolders []string) *Codespace {
	this := Codespace{}
	this.Id = id
	this.Name = name
	this.EnvironmentId = environmentId
	this.Owner = owner
	this.BillableOwner = billableOwner
	this.Repository = repository
	this.Machine = machine
	this.Prebuild = prebuild
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.LastUsedAt = lastUsedAt
	this.State = state
	this.Url = url
	this.GitStatus = gitStatus
	this.Location = location
	this.IdleTimeoutMinutes = idleTimeoutMinutes
	this.WebUrl = webUrl
	this.MachinesUrl = machinesUrl
	this.StartUrl = startUrl
	this.StopUrl = stopUrl
	this.PullsUrl = pullsUrl
	this.RecentFolders = recentFolders
	return &this
}

// NewCodespaceWithDefaults instantiates a new Codespace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodespaceWithDefaults() *Codespace {
	this := Codespace{}
	return &this
}

// GetId returns the Id field value
func (o *Codespace) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Codespace) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Codespace) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Codespace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Codespace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Codespace) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Codespace) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Codespace) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Codespace) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *Codespace) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *Codespace) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *Codespace) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetEnvironmentId returns the EnvironmentId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Codespace) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId.Get() == nil {
		var ret string
		return ret
	}

	return *o.EnvironmentId.Get()
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Codespace) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvironmentId.Get(), o.EnvironmentId.IsSet()
}

// SetEnvironmentId sets field value
func (o *Codespace) SetEnvironmentId(v string) {
	o.EnvironmentId.Set(&v)
}

// GetOwner returns the Owner field value
func (o *Codespace) GetOwner() SimpleUser {
	if o == nil {
		var ret SimpleUser
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *Codespace) GetOwnerOk() (*SimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *Codespace) SetOwner(v SimpleUser) {
	o.Owner = v
}

// GetBillableOwner returns the BillableOwner field value
func (o *Codespace) GetBillableOwner() SimpleUser {
	if o == nil {
		var ret SimpleUser
		return ret
	}

	return o.BillableOwner
}

// GetBillableOwnerOk returns a tuple with the BillableOwner field value
// and a boolean to check if the value has been set.
func (o *Codespace) GetBillableOwnerOk() (*SimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillableOwner, true
}

// SetBillableOwner sets field value
func (o *Codespace) SetBillableOwner(v SimpleUser) {
	o.BillableOwner = v
}

// GetRepository returns the Repository field value
func (o *Codespace) GetRepository() MinimalRepository {
	if o == nil {
		var ret MinimalRepository
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *Codespace) GetRepositoryOk() (*MinimalRepository, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *Codespace) SetRepository(v MinimalRepository) {
	o.Repository = v
}

// GetMachine returns the Machine field value
// If the value is explicit nil, the zero value for NullableCodespaceMachine will be returned
func (o *Codespace) GetMachine() NullableCodespaceMachine {
	if o == nil || o.Machine.Get() == nil {
		var ret NullableCodespaceMachine
		return ret
	}

	return *o.Machine.Get()
}

// GetMachineOk returns a tuple with the Machine field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Codespace) GetMachineOk() (*NullableCodespaceMachine, bool) {
	if o == nil {
		return nil, false
	}
	return o.Machine.Get(), o.Machine.IsSet()
}

// SetMachine sets field value
func (o *Codespace) SetMachine(v NullableCodespaceMachine) {
	o.Machine.Set(&v)
}

// GetDevcontainerPath returns the DevcontainerPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Codespace) GetDevcontainerPath() string {
	if o == nil || o.DevcontainerPath.Get() == nil {
		var ret string
		return ret
	}
	return *o.DevcontainerPath.Get()
}

// GetDevcontainerPathOk returns a tuple with the DevcontainerPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Codespace) GetDevcontainerPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DevcontainerPath.Get(), o.DevcontainerPath.IsSet()
}

// HasDevcontainerPath returns a boolean if a field has been set.
func (o *Codespace) HasDevcontainerPath() bool {
	if o != nil && o.DevcontainerPath.IsSet() {
		return true
	}

	return false
}

// SetDevcontainerPath gets a reference to the given NullableString and assigns it to the DevcontainerPath field.
func (o *Codespace) SetDevcontainerPath(v string) {
	o.DevcontainerPath.Set(&v)
}
// SetDevcontainerPathNil sets the value for DevcontainerPath to be an explicit nil
func (o *Codespace) SetDevcontainerPathNil() {
	o.DevcontainerPath.Set(nil)
}

// UnsetDevcontainerPath ensures that no value is present for DevcontainerPath, not even an explicit nil
func (o *Codespace) UnsetDevcontainerPath() {
	o.DevcontainerPath.Unset()
}

// GetPrebuild returns the Prebuild field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *Codespace) GetPrebuild() bool {
	if o == nil || o.Prebuild.Get() == nil {
		var ret bool
		return ret
	}

	return *o.Prebuild.Get()
}

// GetPrebuildOk returns a tuple with the Prebuild field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Codespace) GetPrebuildOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prebuild.Get(), o.Prebuild.IsSet()
}

// SetPrebuild sets field value
func (o *Codespace) SetPrebuild(v bool) {
	o.Prebuild.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *Codespace) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Codespace) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Codespace) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Codespace) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Codespace) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Codespace) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetLastUsedAt returns the LastUsedAt field value
func (o *Codespace) GetLastUsedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value
// and a boolean to check if the value has been set.
func (o *Codespace) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUsedAt, true
}

// SetLastUsedAt sets field value
func (o *Codespace) SetLastUsedAt(v time.Time) {
	o.LastUsedAt = v
}

// GetState returns the State field value
func (o *Codespace) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Codespace) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Codespace) SetState(v string) {
	o.State = v
}

// GetUrl returns the Url field value
func (o *Codespace) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Codespace) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Codespace) SetUrl(v string) {
	o.Url = v
}

// GetGitStatus returns the GitStatus field value
func (o *Codespace) GetGitStatus() CodespaceGitStatus {
	if o == nil {
		var ret CodespaceGitStatus
		return ret
	}

	return o.GitStatus
}

// GetGitStatusOk returns a tuple with the GitStatus field value
// and a boolean to check if the value has been set.
func (o *Codespace) GetGitStatusOk() (*CodespaceGitStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GitStatus, true
}

// SetGitStatus sets field value
func (o *Codespace) SetGitStatus(v CodespaceGitStatus) {
	o.GitStatus = v
}

// GetLocation returns the Location field value
func (o *Codespace) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *Codespace) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *Codespace) SetLocation(v string) {
	o.Location = v
}

// GetIdleTimeoutMinutes returns the IdleTimeoutMinutes field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *Codespace) GetIdleTimeoutMinutes() int32 {
	if o == nil || o.IdleTimeoutMinutes.Get() == nil {
		var ret int32
		return ret
	}

	return *o.IdleTimeoutMinutes.Get()
}

// GetIdleTimeoutMinutesOk returns a tuple with the IdleTimeoutMinutes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Codespace) GetIdleTimeoutMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdleTimeoutMinutes.Get(), o.IdleTimeoutMinutes.IsSet()
}

// SetIdleTimeoutMinutes sets field value
func (o *Codespace) SetIdleTimeoutMinutes(v int32) {
	o.IdleTimeoutMinutes.Set(&v)
}

// GetWebUrl returns the WebUrl field value
func (o *Codespace) GetWebUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field value
// and a boolean to check if the value has been set.
func (o *Codespace) GetWebUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebUrl, true
}

// SetWebUrl sets field value
func (o *Codespace) SetWebUrl(v string) {
	o.WebUrl = v
}

// GetMachinesUrl returns the MachinesUrl field value
func (o *Codespace) GetMachinesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MachinesUrl
}

// GetMachinesUrlOk returns a tuple with the MachinesUrl field value
// and a boolean to check if the value has been set.
func (o *Codespace) GetMachinesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MachinesUrl, true
}

// SetMachinesUrl sets field value
func (o *Codespace) SetMachinesUrl(v string) {
	o.MachinesUrl = v
}

// GetStartUrl returns the StartUrl field value
func (o *Codespace) GetStartUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartUrl
}

// GetStartUrlOk returns a tuple with the StartUrl field value
// and a boolean to check if the value has been set.
func (o *Codespace) GetStartUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartUrl, true
}

// SetStartUrl sets field value
func (o *Codespace) SetStartUrl(v string) {
	o.StartUrl = v
}

// GetStopUrl returns the StopUrl field value
func (o *Codespace) GetStopUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StopUrl
}

// GetStopUrlOk returns a tuple with the StopUrl field value
// and a boolean to check if the value has been set.
func (o *Codespace) GetStopUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StopUrl, true
}

// SetStopUrl sets field value
func (o *Codespace) SetStopUrl(v string) {
	o.StopUrl = v
}

// GetPullsUrl returns the PullsUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Codespace) GetPullsUrl() string {
	if o == nil || o.PullsUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.PullsUrl.Get()
}

// GetPullsUrlOk returns a tuple with the PullsUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Codespace) GetPullsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PullsUrl.Get(), o.PullsUrl.IsSet()
}

// SetPullsUrl sets field value
func (o *Codespace) SetPullsUrl(v string) {
	o.PullsUrl.Set(&v)
}

// GetRecentFolders returns the RecentFolders field value
func (o *Codespace) GetRecentFolders() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RecentFolders
}

// GetRecentFoldersOk returns a tuple with the RecentFolders field value
// and a boolean to check if the value has been set.
func (o *Codespace) GetRecentFoldersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecentFolders, true
}

// SetRecentFolders sets field value
func (o *Codespace) SetRecentFolders(v []string) {
	o.RecentFolders = v
}

// GetRuntimeConstraints returns the RuntimeConstraints field value if set, zero value otherwise.
func (o *Codespace) GetRuntimeConstraints() CodespaceRuntimeConstraints {
	if o == nil || o.RuntimeConstraints == nil {
		var ret CodespaceRuntimeConstraints
		return ret
	}
	return *o.RuntimeConstraints
}

// GetRuntimeConstraintsOk returns a tuple with the RuntimeConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Codespace) GetRuntimeConstraintsOk() (*CodespaceRuntimeConstraints, bool) {
	if o == nil || o.RuntimeConstraints == nil {
		return nil, false
	}
	return o.RuntimeConstraints, true
}

// HasRuntimeConstraints returns a boolean if a field has been set.
func (o *Codespace) HasRuntimeConstraints() bool {
	if o != nil && o.RuntimeConstraints != nil {
		return true
	}

	return false
}

// SetRuntimeConstraints gets a reference to the given CodespaceRuntimeConstraints and assigns it to the RuntimeConstraints field.
func (o *Codespace) SetRuntimeConstraints(v CodespaceRuntimeConstraints) {
	o.RuntimeConstraints = &v
}

// GetPendingOperation returns the PendingOperation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Codespace) GetPendingOperation() bool {
	if o == nil || o.PendingOperation.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PendingOperation.Get()
}

// GetPendingOperationOk returns a tuple with the PendingOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Codespace) GetPendingOperationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PendingOperation.Get(), o.PendingOperation.IsSet()
}

// HasPendingOperation returns a boolean if a field has been set.
func (o *Codespace) HasPendingOperation() bool {
	if o != nil && o.PendingOperation.IsSet() {
		return true
	}

	return false
}

// SetPendingOperation gets a reference to the given NullableBool and assigns it to the PendingOperation field.
func (o *Codespace) SetPendingOperation(v bool) {
	o.PendingOperation.Set(&v)
}
// SetPendingOperationNil sets the value for PendingOperation to be an explicit nil
func (o *Codespace) SetPendingOperationNil() {
	o.PendingOperation.Set(nil)
}

// UnsetPendingOperation ensures that no value is present for PendingOperation, not even an explicit nil
func (o *Codespace) UnsetPendingOperation() {
	o.PendingOperation.Unset()
}

// GetPendingOperationDisabledReason returns the PendingOperationDisabledReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Codespace) GetPendingOperationDisabledReason() string {
	if o == nil || o.PendingOperationDisabledReason.Get() == nil {
		var ret string
		return ret
	}
	return *o.PendingOperationDisabledReason.Get()
}

// GetPendingOperationDisabledReasonOk returns a tuple with the PendingOperationDisabledReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Codespace) GetPendingOperationDisabledReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PendingOperationDisabledReason.Get(), o.PendingOperationDisabledReason.IsSet()
}

// HasPendingOperationDisabledReason returns a boolean if a field has been set.
func (o *Codespace) HasPendingOperationDisabledReason() bool {
	if o != nil && o.PendingOperationDisabledReason.IsSet() {
		return true
	}

	return false
}

// SetPendingOperationDisabledReason gets a reference to the given NullableString and assigns it to the PendingOperationDisabledReason field.
func (o *Codespace) SetPendingOperationDisabledReason(v string) {
	o.PendingOperationDisabledReason.Set(&v)
}
// SetPendingOperationDisabledReasonNil sets the value for PendingOperationDisabledReason to be an explicit nil
func (o *Codespace) SetPendingOperationDisabledReasonNil() {
	o.PendingOperationDisabledReason.Set(nil)
}

// UnsetPendingOperationDisabledReason ensures that no value is present for PendingOperationDisabledReason, not even an explicit nil
func (o *Codespace) UnsetPendingOperationDisabledReason() {
	o.PendingOperationDisabledReason.Unset()
}

// GetIdleTimeoutNotice returns the IdleTimeoutNotice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Codespace) GetIdleTimeoutNotice() string {
	if o == nil || o.IdleTimeoutNotice.Get() == nil {
		var ret string
		return ret
	}
	return *o.IdleTimeoutNotice.Get()
}

// GetIdleTimeoutNoticeOk returns a tuple with the IdleTimeoutNotice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Codespace) GetIdleTimeoutNoticeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdleTimeoutNotice.Get(), o.IdleTimeoutNotice.IsSet()
}

// HasIdleTimeoutNotice returns a boolean if a field has been set.
func (o *Codespace) HasIdleTimeoutNotice() bool {
	if o != nil && o.IdleTimeoutNotice.IsSet() {
		return true
	}

	return false
}

// SetIdleTimeoutNotice gets a reference to the given NullableString and assigns it to the IdleTimeoutNotice field.
func (o *Codespace) SetIdleTimeoutNotice(v string) {
	o.IdleTimeoutNotice.Set(&v)
}
// SetIdleTimeoutNoticeNil sets the value for IdleTimeoutNotice to be an explicit nil
func (o *Codespace) SetIdleTimeoutNoticeNil() {
	o.IdleTimeoutNotice.Set(nil)
}

// UnsetIdleTimeoutNotice ensures that no value is present for IdleTimeoutNotice, not even an explicit nil
func (o *Codespace) UnsetIdleTimeoutNotice() {
	o.IdleTimeoutNotice.Unset()
}

func (o Codespace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.DisplayName.IsSet() {
		toSerialize["display_name"] = o.DisplayName.Get()
	}
	if true {
		toSerialize["environment_id"] = o.EnvironmentId.Get()
	}
	if true {
		toSerialize["owner"] = o.Owner
	}
	if true {
		toSerialize["billable_owner"] = o.BillableOwner
	}
	if true {
		toSerialize["repository"] = o.Repository
	}
	if true {
		toSerialize["machine"] = o.Machine.Get()
	}
	if o.DevcontainerPath.IsSet() {
		toSerialize["devcontainer_path"] = o.DevcontainerPath.Get()
	}
	if true {
		toSerialize["prebuild"] = o.Prebuild.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["last_used_at"] = o.LastUsedAt
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["git_status"] = o.GitStatus
	}
	if true {
		toSerialize["location"] = o.Location
	}
	if true {
		toSerialize["idle_timeout_minutes"] = o.IdleTimeoutMinutes.Get()
	}
	if true {
		toSerialize["web_url"] = o.WebUrl
	}
	if true {
		toSerialize["machines_url"] = o.MachinesUrl
	}
	if true {
		toSerialize["start_url"] = o.StartUrl
	}
	if true {
		toSerialize["stop_url"] = o.StopUrl
	}
	if true {
		toSerialize["pulls_url"] = o.PullsUrl.Get()
	}
	if true {
		toSerialize["recent_folders"] = o.RecentFolders
	}
	if o.RuntimeConstraints != nil {
		toSerialize["runtime_constraints"] = o.RuntimeConstraints
	}
	if o.PendingOperation.IsSet() {
		toSerialize["pending_operation"] = o.PendingOperation.Get()
	}
	if o.PendingOperationDisabledReason.IsSet() {
		toSerialize["pending_operation_disabled_reason"] = o.PendingOperationDisabledReason.Get()
	}
	if o.IdleTimeoutNotice.IsSet() {
		toSerialize["idle_timeout_notice"] = o.IdleTimeoutNotice.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCodespace struct {
	value *Codespace
	isSet bool
}

func (v NullableCodespace) Get() *Codespace {
	return v.value
}

func (v *NullableCodespace) Set(val *Codespace) {
	v.value = val
	v.isSet = true
}

func (v NullableCodespace) IsSet() bool {
	return v.isSet
}

func (v *NullableCodespace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodespace(val *Codespace) *NullableCodespace {
	return &NullableCodespace{value: val, isSet: true}
}

func (v NullableCodespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodespace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

