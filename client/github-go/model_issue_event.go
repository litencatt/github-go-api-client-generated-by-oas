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

// IssueEvent Issue Event
type IssueEvent struct {
	Id int32 `json:"id"`
	NodeId string `json:"node_id"`
	Url string `json:"url"`
	Actor NullableNullableSimpleUser `json:"actor"`
	Event string `json:"event"`
	CommitId NullableString `json:"commit_id"`
	CommitUrl NullableString `json:"commit_url"`
	CreatedAt time.Time `json:"created_at"`
	Issue NullableNullableIssue `json:"issue,omitempty"`
	Label *IssueEventLabel `json:"label,omitempty"`
	Assignee NullableNullableSimpleUser `json:"assignee,omitempty"`
	Assigner NullableNullableSimpleUser `json:"assigner,omitempty"`
	ReviewRequester NullableNullableSimpleUser `json:"review_requester,omitempty"`
	RequestedReviewer NullableNullableSimpleUser `json:"requested_reviewer,omitempty"`
	RequestedTeam *Team `json:"requested_team,omitempty"`
	DismissedReview *IssueEventDismissedReview `json:"dismissed_review,omitempty"`
	Milestone *IssueEventMilestone `json:"milestone,omitempty"`
	ProjectCard *IssueEventProjectCard `json:"project_card,omitempty"`
	Rename *IssueEventRename `json:"rename,omitempty"`
	AuthorAssociation *AuthorAssociation `json:"author_association,omitempty"`
	LockReason NullableString `json:"lock_reason,omitempty"`
	PerformedViaGithubApp NullableNullableIntegration `json:"performed_via_github_app,omitempty"`
}

// NewIssueEvent instantiates a new IssueEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueEvent(id int32, nodeId string, url string, actor NullableNullableSimpleUser, event string, commitId NullableString, commitUrl NullableString, createdAt time.Time) *IssueEvent {
	this := IssueEvent{}
	this.Id = id
	this.NodeId = nodeId
	this.Url = url
	this.Actor = actor
	this.Event = event
	this.CommitId = commitId
	this.CommitUrl = commitUrl
	this.CreatedAt = createdAt
	return &this
}

// NewIssueEventWithDefaults instantiates a new IssueEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueEventWithDefaults() *IssueEvent {
	this := IssueEvent{}
	return &this
}

// GetId returns the Id field value
func (o *IssueEvent) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IssueEvent) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IssueEvent) SetId(v int32) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *IssueEvent) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *IssueEvent) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *IssueEvent) SetNodeId(v string) {
	o.NodeId = v
}

// GetUrl returns the Url field value
func (o *IssueEvent) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *IssueEvent) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *IssueEvent) SetUrl(v string) {
	o.Url = v
}

// GetActor returns the Actor field value
// If the value is explicit nil, the zero value for NullableSimpleUser will be returned
func (o *IssueEvent) GetActor() NullableSimpleUser {
	if o == nil || o.Actor.Get() == nil {
		var ret NullableSimpleUser
		return ret
	}

	return *o.Actor.Get()
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueEvent) GetActorOk() (*NullableSimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actor.Get(), o.Actor.IsSet()
}

// SetActor sets field value
func (o *IssueEvent) SetActor(v NullableSimpleUser) {
	o.Actor.Set(&v)
}

// GetEvent returns the Event field value
func (o *IssueEvent) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *IssueEvent) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *IssueEvent) SetEvent(v string) {
	o.Event = v
}

// GetCommitId returns the CommitId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IssueEvent) GetCommitId() string {
	if o == nil || o.CommitId.Get() == nil {
		var ret string
		return ret
	}

	return *o.CommitId.Get()
}

// GetCommitIdOk returns a tuple with the CommitId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueEvent) GetCommitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitId.Get(), o.CommitId.IsSet()
}

// SetCommitId sets field value
func (o *IssueEvent) SetCommitId(v string) {
	o.CommitId.Set(&v)
}

// GetCommitUrl returns the CommitUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *IssueEvent) GetCommitUrl() string {
	if o == nil || o.CommitUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.CommitUrl.Get()
}

// GetCommitUrlOk returns a tuple with the CommitUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueEvent) GetCommitUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitUrl.Get(), o.CommitUrl.IsSet()
}

// SetCommitUrl sets field value
func (o *IssueEvent) SetCommitUrl(v string) {
	o.CommitUrl.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *IssueEvent) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IssueEvent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *IssueEvent) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetIssue returns the Issue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueEvent) GetIssue() NullableIssue {
	if o == nil || o.Issue.Get() == nil {
		var ret NullableIssue
		return ret
	}
	return *o.Issue.Get()
}

// GetIssueOk returns a tuple with the Issue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueEvent) GetIssueOk() (*NullableIssue, bool) {
	if o == nil {
		return nil, false
	}
	return o.Issue.Get(), o.Issue.IsSet()
}

// HasIssue returns a boolean if a field has been set.
func (o *IssueEvent) HasIssue() bool {
	if o != nil && o.Issue.IsSet() {
		return true
	}

	return false
}

// SetIssue gets a reference to the given NullableNullableIssue and assigns it to the Issue field.
func (o *IssueEvent) SetIssue(v NullableIssue) {
	o.Issue.Set(&v)
}
// SetIssueNil sets the value for Issue to be an explicit nil
func (o *IssueEvent) SetIssueNil() {
	o.Issue.Set(nil)
}

// UnsetIssue ensures that no value is present for Issue, not even an explicit nil
func (o *IssueEvent) UnsetIssue() {
	o.Issue.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *IssueEvent) GetLabel() IssueEventLabel {
	if o == nil || o.Label == nil {
		var ret IssueEventLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueEvent) GetLabelOk() (*IssueEventLabel, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *IssueEvent) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given IssueEventLabel and assigns it to the Label field.
func (o *IssueEvent) SetLabel(v IssueEventLabel) {
	o.Label = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueEvent) GetAssignee() NullableSimpleUser {
	if o == nil || o.Assignee.Get() == nil {
		var ret NullableSimpleUser
		return ret
	}
	return *o.Assignee.Get()
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueEvent) GetAssigneeOk() (*NullableSimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Assignee.Get(), o.Assignee.IsSet()
}

// HasAssignee returns a boolean if a field has been set.
func (o *IssueEvent) HasAssignee() bool {
	if o != nil && o.Assignee.IsSet() {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given NullableNullableSimpleUser and assigns it to the Assignee field.
func (o *IssueEvent) SetAssignee(v NullableSimpleUser) {
	o.Assignee.Set(&v)
}
// SetAssigneeNil sets the value for Assignee to be an explicit nil
func (o *IssueEvent) SetAssigneeNil() {
	o.Assignee.Set(nil)
}

// UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
func (o *IssueEvent) UnsetAssignee() {
	o.Assignee.Unset()
}

// GetAssigner returns the Assigner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueEvent) GetAssigner() NullableSimpleUser {
	if o == nil || o.Assigner.Get() == nil {
		var ret NullableSimpleUser
		return ret
	}
	return *o.Assigner.Get()
}

// GetAssignerOk returns a tuple with the Assigner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueEvent) GetAssignerOk() (*NullableSimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Assigner.Get(), o.Assigner.IsSet()
}

// HasAssigner returns a boolean if a field has been set.
func (o *IssueEvent) HasAssigner() bool {
	if o != nil && o.Assigner.IsSet() {
		return true
	}

	return false
}

// SetAssigner gets a reference to the given NullableNullableSimpleUser and assigns it to the Assigner field.
func (o *IssueEvent) SetAssigner(v NullableSimpleUser) {
	o.Assigner.Set(&v)
}
// SetAssignerNil sets the value for Assigner to be an explicit nil
func (o *IssueEvent) SetAssignerNil() {
	o.Assigner.Set(nil)
}

// UnsetAssigner ensures that no value is present for Assigner, not even an explicit nil
func (o *IssueEvent) UnsetAssigner() {
	o.Assigner.Unset()
}

// GetReviewRequester returns the ReviewRequester field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueEvent) GetReviewRequester() NullableSimpleUser {
	if o == nil || o.ReviewRequester.Get() == nil {
		var ret NullableSimpleUser
		return ret
	}
	return *o.ReviewRequester.Get()
}

// GetReviewRequesterOk returns a tuple with the ReviewRequester field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueEvent) GetReviewRequesterOk() (*NullableSimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReviewRequester.Get(), o.ReviewRequester.IsSet()
}

// HasReviewRequester returns a boolean if a field has been set.
func (o *IssueEvent) HasReviewRequester() bool {
	if o != nil && o.ReviewRequester.IsSet() {
		return true
	}

	return false
}

// SetReviewRequester gets a reference to the given NullableNullableSimpleUser and assigns it to the ReviewRequester field.
func (o *IssueEvent) SetReviewRequester(v NullableSimpleUser) {
	o.ReviewRequester.Set(&v)
}
// SetReviewRequesterNil sets the value for ReviewRequester to be an explicit nil
func (o *IssueEvent) SetReviewRequesterNil() {
	o.ReviewRequester.Set(nil)
}

// UnsetReviewRequester ensures that no value is present for ReviewRequester, not even an explicit nil
func (o *IssueEvent) UnsetReviewRequester() {
	o.ReviewRequester.Unset()
}

// GetRequestedReviewer returns the RequestedReviewer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueEvent) GetRequestedReviewer() NullableSimpleUser {
	if o == nil || o.RequestedReviewer.Get() == nil {
		var ret NullableSimpleUser
		return ret
	}
	return *o.RequestedReviewer.Get()
}

// GetRequestedReviewerOk returns a tuple with the RequestedReviewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueEvent) GetRequestedReviewerOk() (*NullableSimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestedReviewer.Get(), o.RequestedReviewer.IsSet()
}

// HasRequestedReviewer returns a boolean if a field has been set.
func (o *IssueEvent) HasRequestedReviewer() bool {
	if o != nil && o.RequestedReviewer.IsSet() {
		return true
	}

	return false
}

// SetRequestedReviewer gets a reference to the given NullableNullableSimpleUser and assigns it to the RequestedReviewer field.
func (o *IssueEvent) SetRequestedReviewer(v NullableSimpleUser) {
	o.RequestedReviewer.Set(&v)
}
// SetRequestedReviewerNil sets the value for RequestedReviewer to be an explicit nil
func (o *IssueEvent) SetRequestedReviewerNil() {
	o.RequestedReviewer.Set(nil)
}

// UnsetRequestedReviewer ensures that no value is present for RequestedReviewer, not even an explicit nil
func (o *IssueEvent) UnsetRequestedReviewer() {
	o.RequestedReviewer.Unset()
}

// GetRequestedTeam returns the RequestedTeam field value if set, zero value otherwise.
func (o *IssueEvent) GetRequestedTeam() Team {
	if o == nil || o.RequestedTeam == nil {
		var ret Team
		return ret
	}
	return *o.RequestedTeam
}

// GetRequestedTeamOk returns a tuple with the RequestedTeam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueEvent) GetRequestedTeamOk() (*Team, bool) {
	if o == nil || o.RequestedTeam == nil {
		return nil, false
	}
	return o.RequestedTeam, true
}

// HasRequestedTeam returns a boolean if a field has been set.
func (o *IssueEvent) HasRequestedTeam() bool {
	if o != nil && o.RequestedTeam != nil {
		return true
	}

	return false
}

// SetRequestedTeam gets a reference to the given Team and assigns it to the RequestedTeam field.
func (o *IssueEvent) SetRequestedTeam(v Team) {
	o.RequestedTeam = &v
}

// GetDismissedReview returns the DismissedReview field value if set, zero value otherwise.
func (o *IssueEvent) GetDismissedReview() IssueEventDismissedReview {
	if o == nil || o.DismissedReview == nil {
		var ret IssueEventDismissedReview
		return ret
	}
	return *o.DismissedReview
}

// GetDismissedReviewOk returns a tuple with the DismissedReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueEvent) GetDismissedReviewOk() (*IssueEventDismissedReview, bool) {
	if o == nil || o.DismissedReview == nil {
		return nil, false
	}
	return o.DismissedReview, true
}

// HasDismissedReview returns a boolean if a field has been set.
func (o *IssueEvent) HasDismissedReview() bool {
	if o != nil && o.DismissedReview != nil {
		return true
	}

	return false
}

// SetDismissedReview gets a reference to the given IssueEventDismissedReview and assigns it to the DismissedReview field.
func (o *IssueEvent) SetDismissedReview(v IssueEventDismissedReview) {
	o.DismissedReview = &v
}

// GetMilestone returns the Milestone field value if set, zero value otherwise.
func (o *IssueEvent) GetMilestone() IssueEventMilestone {
	if o == nil || o.Milestone == nil {
		var ret IssueEventMilestone
		return ret
	}
	return *o.Milestone
}

// GetMilestoneOk returns a tuple with the Milestone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueEvent) GetMilestoneOk() (*IssueEventMilestone, bool) {
	if o == nil || o.Milestone == nil {
		return nil, false
	}
	return o.Milestone, true
}

// HasMilestone returns a boolean if a field has been set.
func (o *IssueEvent) HasMilestone() bool {
	if o != nil && o.Milestone != nil {
		return true
	}

	return false
}

// SetMilestone gets a reference to the given IssueEventMilestone and assigns it to the Milestone field.
func (o *IssueEvent) SetMilestone(v IssueEventMilestone) {
	o.Milestone = &v
}

// GetProjectCard returns the ProjectCard field value if set, zero value otherwise.
func (o *IssueEvent) GetProjectCard() IssueEventProjectCard {
	if o == nil || o.ProjectCard == nil {
		var ret IssueEventProjectCard
		return ret
	}
	return *o.ProjectCard
}

// GetProjectCardOk returns a tuple with the ProjectCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueEvent) GetProjectCardOk() (*IssueEventProjectCard, bool) {
	if o == nil || o.ProjectCard == nil {
		return nil, false
	}
	return o.ProjectCard, true
}

// HasProjectCard returns a boolean if a field has been set.
func (o *IssueEvent) HasProjectCard() bool {
	if o != nil && o.ProjectCard != nil {
		return true
	}

	return false
}

// SetProjectCard gets a reference to the given IssueEventProjectCard and assigns it to the ProjectCard field.
func (o *IssueEvent) SetProjectCard(v IssueEventProjectCard) {
	o.ProjectCard = &v
}

// GetRename returns the Rename field value if set, zero value otherwise.
func (o *IssueEvent) GetRename() IssueEventRename {
	if o == nil || o.Rename == nil {
		var ret IssueEventRename
		return ret
	}
	return *o.Rename
}

// GetRenameOk returns a tuple with the Rename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueEvent) GetRenameOk() (*IssueEventRename, bool) {
	if o == nil || o.Rename == nil {
		return nil, false
	}
	return o.Rename, true
}

// HasRename returns a boolean if a field has been set.
func (o *IssueEvent) HasRename() bool {
	if o != nil && o.Rename != nil {
		return true
	}

	return false
}

// SetRename gets a reference to the given IssueEventRename and assigns it to the Rename field.
func (o *IssueEvent) SetRename(v IssueEventRename) {
	o.Rename = &v
}

// GetAuthorAssociation returns the AuthorAssociation field value if set, zero value otherwise.
func (o *IssueEvent) GetAuthorAssociation() AuthorAssociation {
	if o == nil || o.AuthorAssociation == nil {
		var ret AuthorAssociation
		return ret
	}
	return *o.AuthorAssociation
}

// GetAuthorAssociationOk returns a tuple with the AuthorAssociation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueEvent) GetAuthorAssociationOk() (*AuthorAssociation, bool) {
	if o == nil || o.AuthorAssociation == nil {
		return nil, false
	}
	return o.AuthorAssociation, true
}

// HasAuthorAssociation returns a boolean if a field has been set.
func (o *IssueEvent) HasAuthorAssociation() bool {
	if o != nil && o.AuthorAssociation != nil {
		return true
	}

	return false
}

// SetAuthorAssociation gets a reference to the given AuthorAssociation and assigns it to the AuthorAssociation field.
func (o *IssueEvent) SetAuthorAssociation(v AuthorAssociation) {
	o.AuthorAssociation = &v
}

// GetLockReason returns the LockReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueEvent) GetLockReason() string {
	if o == nil || o.LockReason.Get() == nil {
		var ret string
		return ret
	}
	return *o.LockReason.Get()
}

// GetLockReasonOk returns a tuple with the LockReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueEvent) GetLockReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LockReason.Get(), o.LockReason.IsSet()
}

// HasLockReason returns a boolean if a field has been set.
func (o *IssueEvent) HasLockReason() bool {
	if o != nil && o.LockReason.IsSet() {
		return true
	}

	return false
}

// SetLockReason gets a reference to the given NullableString and assigns it to the LockReason field.
func (o *IssueEvent) SetLockReason(v string) {
	o.LockReason.Set(&v)
}
// SetLockReasonNil sets the value for LockReason to be an explicit nil
func (o *IssueEvent) SetLockReasonNil() {
	o.LockReason.Set(nil)
}

// UnsetLockReason ensures that no value is present for LockReason, not even an explicit nil
func (o *IssueEvent) UnsetLockReason() {
	o.LockReason.Unset()
}

// GetPerformedViaGithubApp returns the PerformedViaGithubApp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueEvent) GetPerformedViaGithubApp() NullableIntegration {
	if o == nil || o.PerformedViaGithubApp.Get() == nil {
		var ret NullableIntegration
		return ret
	}
	return *o.PerformedViaGithubApp.Get()
}

// GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueEvent) GetPerformedViaGithubAppOk() (*NullableIntegration, bool) {
	if o == nil {
		return nil, false
	}
	return o.PerformedViaGithubApp.Get(), o.PerformedViaGithubApp.IsSet()
}

// HasPerformedViaGithubApp returns a boolean if a field has been set.
func (o *IssueEvent) HasPerformedViaGithubApp() bool {
	if o != nil && o.PerformedViaGithubApp.IsSet() {
		return true
	}

	return false
}

// SetPerformedViaGithubApp gets a reference to the given NullableNullableIntegration and assigns it to the PerformedViaGithubApp field.
func (o *IssueEvent) SetPerformedViaGithubApp(v NullableIntegration) {
	o.PerformedViaGithubApp.Set(&v)
}
// SetPerformedViaGithubAppNil sets the value for PerformedViaGithubApp to be an explicit nil
func (o *IssueEvent) SetPerformedViaGithubAppNil() {
	o.PerformedViaGithubApp.Set(nil)
}

// UnsetPerformedViaGithubApp ensures that no value is present for PerformedViaGithubApp, not even an explicit nil
func (o *IssueEvent) UnsetPerformedViaGithubApp() {
	o.PerformedViaGithubApp.Unset()
}

func (o IssueEvent) MarshalJSON() ([]byte, error) {
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
		toSerialize["actor"] = o.Actor.Get()
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
	if o.Issue.IsSet() {
		toSerialize["issue"] = o.Issue.Get()
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Assignee.IsSet() {
		toSerialize["assignee"] = o.Assignee.Get()
	}
	if o.Assigner.IsSet() {
		toSerialize["assigner"] = o.Assigner.Get()
	}
	if o.ReviewRequester.IsSet() {
		toSerialize["review_requester"] = o.ReviewRequester.Get()
	}
	if o.RequestedReviewer.IsSet() {
		toSerialize["requested_reviewer"] = o.RequestedReviewer.Get()
	}
	if o.RequestedTeam != nil {
		toSerialize["requested_team"] = o.RequestedTeam
	}
	if o.DismissedReview != nil {
		toSerialize["dismissed_review"] = o.DismissedReview
	}
	if o.Milestone != nil {
		toSerialize["milestone"] = o.Milestone
	}
	if o.ProjectCard != nil {
		toSerialize["project_card"] = o.ProjectCard
	}
	if o.Rename != nil {
		toSerialize["rename"] = o.Rename
	}
	if o.AuthorAssociation != nil {
		toSerialize["author_association"] = o.AuthorAssociation
	}
	if o.LockReason.IsSet() {
		toSerialize["lock_reason"] = o.LockReason.Get()
	}
	if o.PerformedViaGithubApp.IsSet() {
		toSerialize["performed_via_github_app"] = o.PerformedViaGithubApp.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIssueEvent struct {
	value *IssueEvent
	isSet bool
}

func (v NullableIssueEvent) Get() *IssueEvent {
	return v.value
}

func (v *NullableIssueEvent) Set(val *IssueEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueEvent(val *IssueEvent) *NullableIssueEvent {
	return &NullableIssueEvent{value: val, isSet: true}
}

func (v NullableIssueEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


