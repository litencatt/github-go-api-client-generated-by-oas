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

// LabeledIssueEvent Labeled Issue Event
type LabeledIssueEvent struct {
	Id int32 `json:"id"`
	NodeId string `json:"node_id"`
	Url string `json:"url"`
	Actor SimpleUser `json:"actor"`
	Event string `json:"event"`
	CommitId NullableString `json:"commit_id"`
	CommitUrl NullableString `json:"commit_url"`
	CreatedAt string `json:"created_at"`
	PerformedViaGithubApp NullableNullableIntegration `json:"performed_via_github_app"`
	Label LabeledIssueEventLabel `json:"label"`
}

// NewLabeledIssueEvent instantiates a new LabeledIssueEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabeledIssueEvent(id int32, nodeId string, url string, actor SimpleUser, event string, commitId NullableString, commitUrl NullableString, createdAt string, performedViaGithubApp NullableNullableIntegration, label LabeledIssueEventLabel) *LabeledIssueEvent {
	this := LabeledIssueEvent{}
	this.Id = id
	this.NodeId = nodeId
	this.Url = url
	this.Actor = actor
	this.Event = event
	this.CommitId = commitId
	this.CommitUrl = commitUrl
	this.CreatedAt = createdAt
	this.PerformedViaGithubApp = performedViaGithubApp
	this.Label = label
	return &this
}

// NewLabeledIssueEventWithDefaults instantiates a new LabeledIssueEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabeledIssueEventWithDefaults() *LabeledIssueEvent {
	this := LabeledIssueEvent{}
	return &this
}

// GetId returns the Id field value
func (o *LabeledIssueEvent) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LabeledIssueEvent) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LabeledIssueEvent) SetId(v int32) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *LabeledIssueEvent) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *LabeledIssueEvent) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *LabeledIssueEvent) SetNodeId(v string) {
	o.NodeId = v
}

// GetUrl returns the Url field value
func (o *LabeledIssueEvent) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *LabeledIssueEvent) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *LabeledIssueEvent) SetUrl(v string) {
	o.Url = v
}

// GetActor returns the Actor field value
func (o *LabeledIssueEvent) GetActor() SimpleUser {
	if o == nil {
		var ret SimpleUser
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *LabeledIssueEvent) GetActorOk() (*SimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *LabeledIssueEvent) SetActor(v SimpleUser) {
	o.Actor = v
}

// GetEvent returns the Event field value
func (o *LabeledIssueEvent) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *LabeledIssueEvent) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *LabeledIssueEvent) SetEvent(v string) {
	o.Event = v
}

// GetCommitId returns the CommitId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LabeledIssueEvent) GetCommitId() string {
	if o == nil || o.CommitId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CommitId.Get()
}

// GetCommitIdOk returns a tuple with the CommitId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LabeledIssueEvent) GetCommitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitId.Get(), o.CommitId.IsSet()
}

// SetCommitId sets field value
func (o *LabeledIssueEvent) SetCommitId(v string) {
	o.CommitId.Set(&v)
}

// GetCommitUrl returns the CommitUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LabeledIssueEvent) GetCommitUrl() string {
	if o == nil || o.CommitUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.CommitUrl.Get()
}

// GetCommitUrlOk returns a tuple with the CommitUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LabeledIssueEvent) GetCommitUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitUrl.Get(), o.CommitUrl.IsSet()
}

// SetCommitUrl sets field value
func (o *LabeledIssueEvent) SetCommitUrl(v string) {
	o.CommitUrl.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *LabeledIssueEvent) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LabeledIssueEvent) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *LabeledIssueEvent) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetPerformedViaGithubApp returns the PerformedViaGithubApp field value
// If the value is explicit nil, the zero value for NullableIntegration will be returned
func (o *LabeledIssueEvent) GetPerformedViaGithubApp() NullableIntegration {
	if o == nil || o.PerformedViaGithubApp.Get() == nil {
		var ret NullableIntegration
		return ret
	}

	return *o.PerformedViaGithubApp.Get()
}

// GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LabeledIssueEvent) GetPerformedViaGithubAppOk() (*NullableIntegration, bool) {
	if o == nil {
		return nil, false
	}
	return o.PerformedViaGithubApp.Get(), o.PerformedViaGithubApp.IsSet()
}

// SetPerformedViaGithubApp sets field value
func (o *LabeledIssueEvent) SetPerformedViaGithubApp(v NullableIntegration) {
	o.PerformedViaGithubApp.Set(&v)
}

// GetLabel returns the Label field value
func (o *LabeledIssueEvent) GetLabel() LabeledIssueEventLabel {
	if o == nil {
		var ret LabeledIssueEventLabel
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *LabeledIssueEvent) GetLabelOk() (*LabeledIssueEventLabel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *LabeledIssueEvent) SetLabel(v LabeledIssueEventLabel) {
	o.Label = v
}

func (o LabeledIssueEvent) MarshalJSON() ([]byte, error) {
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
		toSerialize["label"] = o.Label
	}
	return json.Marshal(toSerialize)
}

type NullableLabeledIssueEvent struct {
	value *LabeledIssueEvent
	isSet bool
}

func (v NullableLabeledIssueEvent) Get() *LabeledIssueEvent {
	return v.value
}

func (v *NullableLabeledIssueEvent) Set(val *LabeledIssueEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableLabeledIssueEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableLabeledIssueEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabeledIssueEvent(val *LabeledIssueEvent) *NullableLabeledIssueEvent {
	return &NullableLabeledIssueEvent{value: val, isSet: true}
}

func (v NullableLabeledIssueEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabeledIssueEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


