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

// UnassignedIssueEvent Unassigned Issue Event
type UnassignedIssueEvent struct {
	Id int32 `json:"id"`
	NodeId string `json:"node_id"`
	Url string `json:"url"`
	Actor SimpleUser `json:"actor"`
	Event string `json:"event"`
	CommitId NullableString `json:"commit_id"`
	CommitUrl NullableString `json:"commit_url"`
	CreatedAt string `json:"created_at"`
	PerformedViaGithubApp NullableNullableIntegration `json:"performed_via_github_app"`
	Assignee SimpleUser `json:"assignee"`
	Assigner SimpleUser `json:"assigner"`
}

// NewUnassignedIssueEvent instantiates a new UnassignedIssueEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnassignedIssueEvent(id int32, nodeId string, url string, actor SimpleUser, event string, commitId NullableString, commitUrl NullableString, createdAt string, performedViaGithubApp NullableNullableIntegration, assignee SimpleUser, assigner SimpleUser) *UnassignedIssueEvent {
	this := UnassignedIssueEvent{}
	this.Id = id
	this.NodeId = nodeId
	this.Url = url
	this.Actor = actor
	this.Event = event
	this.CommitId = commitId
	this.CommitUrl = commitUrl
	this.CreatedAt = createdAt
	this.PerformedViaGithubApp = performedViaGithubApp
	this.Assignee = assignee
	this.Assigner = assigner
	return &this
}

// NewUnassignedIssueEventWithDefaults instantiates a new UnassignedIssueEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnassignedIssueEventWithDefaults() *UnassignedIssueEvent {
	this := UnassignedIssueEvent{}
	return &this
}

// GetId returns the Id field value
func (o *UnassignedIssueEvent) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UnassignedIssueEvent) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UnassignedIssueEvent) SetId(v int32) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *UnassignedIssueEvent) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *UnassignedIssueEvent) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *UnassignedIssueEvent) SetNodeId(v string) {
	o.NodeId = v
}

// GetUrl returns the Url field value
func (o *UnassignedIssueEvent) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *UnassignedIssueEvent) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *UnassignedIssueEvent) SetUrl(v string) {
	o.Url = v
}

// GetActor returns the Actor field value
func (o *UnassignedIssueEvent) GetActor() SimpleUser {
	if o == nil {
		var ret SimpleUser
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *UnassignedIssueEvent) GetActorOk() (*SimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *UnassignedIssueEvent) SetActor(v SimpleUser) {
	o.Actor = v
}

// GetEvent returns the Event field value
func (o *UnassignedIssueEvent) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *UnassignedIssueEvent) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *UnassignedIssueEvent) SetEvent(v string) {
	o.Event = v
}

// GetCommitId returns the CommitId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UnassignedIssueEvent) GetCommitId() string {
	if o == nil || o.CommitId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CommitId.Get()
}

// GetCommitIdOk returns a tuple with the CommitId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnassignedIssueEvent) GetCommitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitId.Get(), o.CommitId.IsSet()
}

// SetCommitId sets field value
func (o *UnassignedIssueEvent) SetCommitId(v string) {
	o.CommitId.Set(&v)
}

// GetCommitUrl returns the CommitUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UnassignedIssueEvent) GetCommitUrl() string {
	if o == nil || o.CommitUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.CommitUrl.Get()
}

// GetCommitUrlOk returns a tuple with the CommitUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnassignedIssueEvent) GetCommitUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitUrl.Get(), o.CommitUrl.IsSet()
}

// SetCommitUrl sets field value
func (o *UnassignedIssueEvent) SetCommitUrl(v string) {
	o.CommitUrl.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *UnassignedIssueEvent) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *UnassignedIssueEvent) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *UnassignedIssueEvent) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetPerformedViaGithubApp returns the PerformedViaGithubApp field value
// If the value is explicit nil, the zero value for NullableIntegration will be returned
func (o *UnassignedIssueEvent) GetPerformedViaGithubApp() NullableIntegration {
	if o == nil || o.PerformedViaGithubApp.Get() == nil {
		var ret NullableIntegration
		return ret
	}

	return *o.PerformedViaGithubApp.Get()
}

// GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UnassignedIssueEvent) GetPerformedViaGithubAppOk() (*NullableIntegration, bool) {
	if o == nil {
		return nil, false
	}
	return o.PerformedViaGithubApp.Get(), o.PerformedViaGithubApp.IsSet()
}

// SetPerformedViaGithubApp sets field value
func (o *UnassignedIssueEvent) SetPerformedViaGithubApp(v NullableIntegration) {
	o.PerformedViaGithubApp.Set(&v)
}

// GetAssignee returns the Assignee field value
func (o *UnassignedIssueEvent) GetAssignee() SimpleUser {
	if o == nil {
		var ret SimpleUser
		return ret
	}

	return o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value
// and a boolean to check if the value has been set.
func (o *UnassignedIssueEvent) GetAssigneeOk() (*SimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assignee, true
}

// SetAssignee sets field value
func (o *UnassignedIssueEvent) SetAssignee(v SimpleUser) {
	o.Assignee = v
}

// GetAssigner returns the Assigner field value
func (o *UnassignedIssueEvent) GetAssigner() SimpleUser {
	if o == nil {
		var ret SimpleUser
		return ret
	}

	return o.Assigner
}

// GetAssignerOk returns a tuple with the Assigner field value
// and a boolean to check if the value has been set.
func (o *UnassignedIssueEvent) GetAssignerOk() (*SimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Assigner, true
}

// SetAssigner sets field value
func (o *UnassignedIssueEvent) SetAssigner(v SimpleUser) {
	o.Assigner = v
}

func (o UnassignedIssueEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["actor"] = o.Actor
	}
	if true {
		toSerialize["event"] = o.Event
	}
	if true {
		toSerialize["commit_id"] = o.CommitId.Get()
	}
	if true {
		toSerialize["commit_url"] = o.CommitUrl.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["performed_via_github_app"] = o.PerformedViaGithubApp.Get()
	}
	if true {
		toSerialize["assignee"] = o.Assignee
	}
	if true {
		toSerialize["assigner"] = o.Assigner
	}
	return json.Marshal(toSerialize)
}

type NullableUnassignedIssueEvent struct {
	value *UnassignedIssueEvent
	isSet bool
}

func (v NullableUnassignedIssueEvent) Get() *UnassignedIssueEvent {
	return v.value
}

func (v *NullableUnassignedIssueEvent) Set(val *UnassignedIssueEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableUnassignedIssueEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableUnassignedIssueEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnassignedIssueEvent(val *UnassignedIssueEvent) *NullableUnassignedIssueEvent {
	return &NullableUnassignedIssueEvent{value: val, isSet: true}
}

func (v NullableUnassignedIssueEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnassignedIssueEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


