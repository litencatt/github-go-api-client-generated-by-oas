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

// PullRequestSimple Pull Request Simple
type PullRequestSimple struct {
	Url string `json:"url"`
	Id int32 `json:"id"`
	NodeId string `json:"node_id"`
	HtmlUrl string `json:"html_url"`
	DiffUrl string `json:"diff_url"`
	PatchUrl string `json:"patch_url"`
	IssueUrl string `json:"issue_url"`
	CommitsUrl string `json:"commits_url"`
	ReviewCommentsUrl string `json:"review_comments_url"`
	ReviewCommentUrl string `json:"review_comment_url"`
	CommentsUrl string `json:"comments_url"`
	StatusesUrl string `json:"statuses_url"`
	Number int32 `json:"number"`
	State string `json:"state"`
	Locked bool `json:"locked"`
	Title string `json:"title"`
	User NullableNullableSimpleUser `json:"user"`
	Body NullableString `json:"body"`
	Labels []PullRequestSimpleLabelsInner `json:"labels"`
	Milestone NullableNullableMilestone `json:"milestone"`
	ActiveLockReason NullableString `json:"active_lock_reason,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ClosedAt NullableTime `json:"closed_at"`
	MergedAt NullableTime `json:"merged_at"`
	MergeCommitSha NullableString `json:"merge_commit_sha"`
	Assignee NullableNullableSimpleUser `json:"assignee"`
	Assignees []SimpleUser `json:"assignees,omitempty"`
	RequestedReviewers []SimpleUser `json:"requested_reviewers,omitempty"`
	RequestedTeams []Team `json:"requested_teams,omitempty"`
	Head PullRequestSimpleHead `json:"head"`
	Base PullRequestSimpleHead `json:"base"`
	Links PullRequestSimpleLinks `json:"_links"`
	AuthorAssociation AuthorAssociation `json:"author_association"`
	AutoMerge NullableAutoMerge `json:"auto_merge"`
	// Indicates whether or not the pull request is a draft.
	Draft *bool `json:"draft,omitempty"`
}

// NewPullRequestSimple instantiates a new PullRequestSimple object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPullRequestSimple(url string, id int32, nodeId string, htmlUrl string, diffUrl string, patchUrl string, issueUrl string, commitsUrl string, reviewCommentsUrl string, reviewCommentUrl string, commentsUrl string, statusesUrl string, number int32, state string, locked bool, title string, user NullableNullableSimpleUser, body NullableString, labels []PullRequestSimpleLabelsInner, milestone NullableNullableMilestone, createdAt time.Time, updatedAt time.Time, closedAt NullableTime, mergedAt NullableTime, mergeCommitSha NullableString, assignee NullableNullableSimpleUser, head PullRequestSimpleHead, base PullRequestSimpleHead, links PullRequestSimpleLinks, authorAssociation AuthorAssociation, autoMerge NullableAutoMerge) *PullRequestSimple {
	this := PullRequestSimple{}
	this.Url = url
	this.Id = id
	this.NodeId = nodeId
	this.HtmlUrl = htmlUrl
	this.DiffUrl = diffUrl
	this.PatchUrl = patchUrl
	this.IssueUrl = issueUrl
	this.CommitsUrl = commitsUrl
	this.ReviewCommentsUrl = reviewCommentsUrl
	this.ReviewCommentUrl = reviewCommentUrl
	this.CommentsUrl = commentsUrl
	this.StatusesUrl = statusesUrl
	this.Number = number
	this.State = state
	this.Locked = locked
	this.Title = title
	this.User = user
	this.Body = body
	this.Labels = labels
	this.Milestone = milestone
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.ClosedAt = closedAt
	this.MergedAt = mergedAt
	this.MergeCommitSha = mergeCommitSha
	this.Assignee = assignee
	this.Head = head
	this.Base = base
	this.Links = links
	this.AuthorAssociation = authorAssociation
	this.AutoMerge = autoMerge
	return &this
}

// NewPullRequestSimpleWithDefaults instantiates a new PullRequestSimple object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPullRequestSimpleWithDefaults() *PullRequestSimple {
	this := PullRequestSimple{}
	return &this
}

// GetUrl returns the Url field value
func (o *PullRequestSimple) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PullRequestSimple) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *PullRequestSimple) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PullRequestSimple) SetId(v int32) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *PullRequestSimple) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *PullRequestSimple) SetNodeId(v string) {
	o.NodeId = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *PullRequestSimple) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *PullRequestSimple) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetDiffUrl returns the DiffUrl field value
func (o *PullRequestSimple) GetDiffUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiffUrl
}

// GetDiffUrlOk returns a tuple with the DiffUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetDiffUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiffUrl, true
}

// SetDiffUrl sets field value
func (o *PullRequestSimple) SetDiffUrl(v string) {
	o.DiffUrl = v
}

// GetPatchUrl returns the PatchUrl field value
func (o *PullRequestSimple) GetPatchUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PatchUrl
}

// GetPatchUrlOk returns a tuple with the PatchUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetPatchUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PatchUrl, true
}

// SetPatchUrl sets field value
func (o *PullRequestSimple) SetPatchUrl(v string) {
	o.PatchUrl = v
}

// GetIssueUrl returns the IssueUrl field value
func (o *PullRequestSimple) GetIssueUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssueUrl
}

// GetIssueUrlOk returns a tuple with the IssueUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetIssueUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueUrl, true
}

// SetIssueUrl sets field value
func (o *PullRequestSimple) SetIssueUrl(v string) {
	o.IssueUrl = v
}

// GetCommitsUrl returns the CommitsUrl field value
func (o *PullRequestSimple) GetCommitsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitsUrl
}

// GetCommitsUrlOk returns a tuple with the CommitsUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetCommitsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitsUrl, true
}

// SetCommitsUrl sets field value
func (o *PullRequestSimple) SetCommitsUrl(v string) {
	o.CommitsUrl = v
}

// GetReviewCommentsUrl returns the ReviewCommentsUrl field value
func (o *PullRequestSimple) GetReviewCommentsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReviewCommentsUrl
}

// GetReviewCommentsUrlOk returns a tuple with the ReviewCommentsUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetReviewCommentsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewCommentsUrl, true
}

// SetReviewCommentsUrl sets field value
func (o *PullRequestSimple) SetReviewCommentsUrl(v string) {
	o.ReviewCommentsUrl = v
}

// GetReviewCommentUrl returns the ReviewCommentUrl field value
func (o *PullRequestSimple) GetReviewCommentUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReviewCommentUrl
}

// GetReviewCommentUrlOk returns a tuple with the ReviewCommentUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetReviewCommentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewCommentUrl, true
}

// SetReviewCommentUrl sets field value
func (o *PullRequestSimple) SetReviewCommentUrl(v string) {
	o.ReviewCommentUrl = v
}

// GetCommentsUrl returns the CommentsUrl field value
func (o *PullRequestSimple) GetCommentsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommentsUrl
}

// GetCommentsUrlOk returns a tuple with the CommentsUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetCommentsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommentsUrl, true
}

// SetCommentsUrl sets field value
func (o *PullRequestSimple) SetCommentsUrl(v string) {
	o.CommentsUrl = v
}

// GetStatusesUrl returns the StatusesUrl field value
func (o *PullRequestSimple) GetStatusesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusesUrl
}

// GetStatusesUrlOk returns a tuple with the StatusesUrl field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetStatusesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusesUrl, true
}

// SetStatusesUrl sets field value
func (o *PullRequestSimple) SetStatusesUrl(v string) {
	o.StatusesUrl = v
}

// GetNumber returns the Number field value
func (o *PullRequestSimple) GetNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *PullRequestSimple) SetNumber(v int32) {
	o.Number = v
}

// GetState returns the State field value
func (o *PullRequestSimple) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *PullRequestSimple) SetState(v string) {
	o.State = v
}

// GetLocked returns the Locked field value
func (o *PullRequestSimple) GetLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Locked
}

// GetLockedOk returns a tuple with the Locked field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locked, true
}

// SetLocked sets field value
func (o *PullRequestSimple) SetLocked(v bool) {
	o.Locked = v
}

// GetTitle returns the Title field value
func (o *PullRequestSimple) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PullRequestSimple) SetTitle(v string) {
	o.Title = v
}

// GetUser returns the User field value
// If the value is explicit nil, the zero value for NullableSimpleUser will be returned
func (o *PullRequestSimple) GetUser() NullableSimpleUser {
	if o == nil || o.User.Get() == nil {
		var ret NullableSimpleUser
		return ret
	}

	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequestSimple) GetUserOk() (*NullableSimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// SetUser sets field value
func (o *PullRequestSimple) SetUser(v NullableSimpleUser) {
	o.User.Set(&v)
}

// GetBody returns the Body field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PullRequestSimple) GetBody() string {
	if o == nil || o.Body.Get() == nil {
		var ret string
		return ret
	}

	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequestSimple) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// SetBody sets field value
func (o *PullRequestSimple) SetBody(v string) {
	o.Body.Set(&v)
}

// GetLabels returns the Labels field value
func (o *PullRequestSimple) GetLabels() []PullRequestSimpleLabelsInner {
	if o == nil {
		var ret []PullRequestSimpleLabelsInner
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetLabelsOk() ([]PullRequestSimpleLabelsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *PullRequestSimple) SetLabels(v []PullRequestSimpleLabelsInner) {
	o.Labels = v
}

// GetMilestone returns the Milestone field value
// If the value is explicit nil, the zero value for NullableMilestone will be returned
func (o *PullRequestSimple) GetMilestone() NullableMilestone {
	if o == nil || o.Milestone.Get() == nil {
		var ret NullableMilestone
		return ret
	}

	return *o.Milestone.Get()
}

// GetMilestoneOk returns a tuple with the Milestone field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequestSimple) GetMilestoneOk() (*NullableMilestone, bool) {
	if o == nil {
		return nil, false
	}
	return o.Milestone.Get(), o.Milestone.IsSet()
}

// SetMilestone sets field value
func (o *PullRequestSimple) SetMilestone(v NullableMilestone) {
	o.Milestone.Set(&v)
}

// GetActiveLockReason returns the ActiveLockReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PullRequestSimple) GetActiveLockReason() string {
	if o == nil || o.ActiveLockReason.Get() == nil {
		var ret string
		return ret
	}
	return *o.ActiveLockReason.Get()
}

// GetActiveLockReasonOk returns a tuple with the ActiveLockReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequestSimple) GetActiveLockReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveLockReason.Get(), o.ActiveLockReason.IsSet()
}

// HasActiveLockReason returns a boolean if a field has been set.
func (o *PullRequestSimple) HasActiveLockReason() bool {
	if o != nil && o.ActiveLockReason.IsSet() {
		return true
	}

	return false
}

// SetActiveLockReason gets a reference to the given NullableString and assigns it to the ActiveLockReason field.
func (o *PullRequestSimple) SetActiveLockReason(v string) {
	o.ActiveLockReason.Set(&v)
}
// SetActiveLockReasonNil sets the value for ActiveLockReason to be an explicit nil
func (o *PullRequestSimple) SetActiveLockReasonNil() {
	o.ActiveLockReason.Set(nil)
}

// UnsetActiveLockReason ensures that no value is present for ActiveLockReason, not even an explicit nil
func (o *PullRequestSimple) UnsetActiveLockReason() {
	o.ActiveLockReason.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *PullRequestSimple) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PullRequestSimple) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PullRequestSimple) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PullRequestSimple) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetClosedAt returns the ClosedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *PullRequestSimple) GetClosedAt() time.Time {
	if o == nil || o.ClosedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ClosedAt.Get()
}

// GetClosedAtOk returns a tuple with the ClosedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequestSimple) GetClosedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClosedAt.Get(), o.ClosedAt.IsSet()
}

// SetClosedAt sets field value
func (o *PullRequestSimple) SetClosedAt(v time.Time) {
	o.ClosedAt.Set(&v)
}

// GetMergedAt returns the MergedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *PullRequestSimple) GetMergedAt() time.Time {
	if o == nil || o.MergedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.MergedAt.Get()
}

// GetMergedAtOk returns a tuple with the MergedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequestSimple) GetMergedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.MergedAt.Get(), o.MergedAt.IsSet()
}

// SetMergedAt sets field value
func (o *PullRequestSimple) SetMergedAt(v time.Time) {
	o.MergedAt.Set(&v)
}

// GetMergeCommitSha returns the MergeCommitSha field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PullRequestSimple) GetMergeCommitSha() string {
	if o == nil || o.MergeCommitSha.Get() == nil {
		var ret string
		return ret
	}

	return *o.MergeCommitSha.Get()
}

// GetMergeCommitShaOk returns a tuple with the MergeCommitSha field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequestSimple) GetMergeCommitShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MergeCommitSha.Get(), o.MergeCommitSha.IsSet()
}

// SetMergeCommitSha sets field value
func (o *PullRequestSimple) SetMergeCommitSha(v string) {
	o.MergeCommitSha.Set(&v)
}

// GetAssignee returns the Assignee field value
// If the value is explicit nil, the zero value for NullableSimpleUser will be returned
func (o *PullRequestSimple) GetAssignee() NullableSimpleUser {
	if o == nil || o.Assignee.Get() == nil {
		var ret NullableSimpleUser
		return ret
	}

	return *o.Assignee.Get()
}

// GetAssigneeOk returns a tuple with the Assignee field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequestSimple) GetAssigneeOk() (*NullableSimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Assignee.Get(), o.Assignee.IsSet()
}

// SetAssignee sets field value
func (o *PullRequestSimple) SetAssignee(v NullableSimpleUser) {
	o.Assignee.Set(&v)
}

// GetAssignees returns the Assignees field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PullRequestSimple) GetAssignees() []SimpleUser {
	if o == nil {
		var ret []SimpleUser
		return ret
	}
	return o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequestSimple) GetAssigneesOk() ([]SimpleUser, bool) {
	if o == nil || o.Assignees == nil {
		return nil, false
	}
	return o.Assignees, true
}

// HasAssignees returns a boolean if a field has been set.
func (o *PullRequestSimple) HasAssignees() bool {
	if o != nil && o.Assignees != nil {
		return true
	}

	return false
}

// SetAssignees gets a reference to the given []SimpleUser and assigns it to the Assignees field.
func (o *PullRequestSimple) SetAssignees(v []SimpleUser) {
	o.Assignees = v
}

// GetRequestedReviewers returns the RequestedReviewers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PullRequestSimple) GetRequestedReviewers() []SimpleUser {
	if o == nil {
		var ret []SimpleUser
		return ret
	}
	return o.RequestedReviewers
}

// GetRequestedReviewersOk returns a tuple with the RequestedReviewers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequestSimple) GetRequestedReviewersOk() ([]SimpleUser, bool) {
	if o == nil || o.RequestedReviewers == nil {
		return nil, false
	}
	return o.RequestedReviewers, true
}

// HasRequestedReviewers returns a boolean if a field has been set.
func (o *PullRequestSimple) HasRequestedReviewers() bool {
	if o != nil && o.RequestedReviewers != nil {
		return true
	}

	return false
}

// SetRequestedReviewers gets a reference to the given []SimpleUser and assigns it to the RequestedReviewers field.
func (o *PullRequestSimple) SetRequestedReviewers(v []SimpleUser) {
	o.RequestedReviewers = v
}

// GetRequestedTeams returns the RequestedTeams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PullRequestSimple) GetRequestedTeams() []Team {
	if o == nil {
		var ret []Team
		return ret
	}
	return o.RequestedTeams
}

// GetRequestedTeamsOk returns a tuple with the RequestedTeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequestSimple) GetRequestedTeamsOk() ([]Team, bool) {
	if o == nil || o.RequestedTeams == nil {
		return nil, false
	}
	return o.RequestedTeams, true
}

// HasRequestedTeams returns a boolean if a field has been set.
func (o *PullRequestSimple) HasRequestedTeams() bool {
	if o != nil && o.RequestedTeams != nil {
		return true
	}

	return false
}

// SetRequestedTeams gets a reference to the given []Team and assigns it to the RequestedTeams field.
func (o *PullRequestSimple) SetRequestedTeams(v []Team) {
	o.RequestedTeams = v
}

// GetHead returns the Head field value
func (o *PullRequestSimple) GetHead() PullRequestSimpleHead {
	if o == nil {
		var ret PullRequestSimpleHead
		return ret
	}

	return o.Head
}

// GetHeadOk returns a tuple with the Head field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetHeadOk() (*PullRequestSimpleHead, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Head, true
}

// SetHead sets field value
func (o *PullRequestSimple) SetHead(v PullRequestSimpleHead) {
	o.Head = v
}

// GetBase returns the Base field value
func (o *PullRequestSimple) GetBase() PullRequestSimpleHead {
	if o == nil {
		var ret PullRequestSimpleHead
		return ret
	}

	return o.Base
}

// GetBaseOk returns a tuple with the Base field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetBaseOk() (*PullRequestSimpleHead, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Base, true
}

// SetBase sets field value
func (o *PullRequestSimple) SetBase(v PullRequestSimpleHead) {
	o.Base = v
}

// GetLinks returns the Links field value
func (o *PullRequestSimple) GetLinks() PullRequestSimpleLinks {
	if o == nil {
		var ret PullRequestSimpleLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetLinksOk() (*PullRequestSimpleLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PullRequestSimple) SetLinks(v PullRequestSimpleLinks) {
	o.Links = v
}

// GetAuthorAssociation returns the AuthorAssociation field value
func (o *PullRequestSimple) GetAuthorAssociation() AuthorAssociation {
	if o == nil {
		var ret AuthorAssociation
		return ret
	}

	return o.AuthorAssociation
}

// GetAuthorAssociationOk returns a tuple with the AuthorAssociation field value
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetAuthorAssociationOk() (*AuthorAssociation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorAssociation, true
}

// SetAuthorAssociation sets field value
func (o *PullRequestSimple) SetAuthorAssociation(v AuthorAssociation) {
	o.AuthorAssociation = v
}

// GetAutoMerge returns the AutoMerge field value
// If the value is explicit nil, the zero value for AutoMerge will be returned
func (o *PullRequestSimple) GetAutoMerge() AutoMerge {
	if o == nil || o.AutoMerge.Get() == nil {
		var ret AutoMerge
		return ret
	}

	return *o.AutoMerge.Get()
}

// GetAutoMergeOk returns a tuple with the AutoMerge field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PullRequestSimple) GetAutoMergeOk() (*AutoMerge, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoMerge.Get(), o.AutoMerge.IsSet()
}

// SetAutoMerge sets field value
func (o *PullRequestSimple) SetAutoMerge(v AutoMerge) {
	o.AutoMerge.Set(&v)
}

// GetDraft returns the Draft field value if set, zero value otherwise.
func (o *PullRequestSimple) GetDraft() bool {
	if o == nil || o.Draft == nil {
		var ret bool
		return ret
	}
	return *o.Draft
}

// GetDraftOk returns a tuple with the Draft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PullRequestSimple) GetDraftOk() (*bool, bool) {
	if o == nil || o.Draft == nil {
		return nil, false
	}
	return o.Draft, true
}

// HasDraft returns a boolean if a field has been set.
func (o *PullRequestSimple) HasDraft() bool {
	if o != nil && o.Draft != nil {
		return true
	}

	return false
}

// SetDraft gets a reference to the given bool and assigns it to the Draft field.
func (o *PullRequestSimple) SetDraft(v bool) {
	o.Draft = &v
}

func (o PullRequestSimple) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	if true {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if true {
		toSerialize["diff_url"] = o.DiffUrl
	}
	if true {
		toSerialize["patch_url"] = o.PatchUrl
	}
	if true {
		toSerialize["issue_url"] = o.IssueUrl
	}
	if true {
		toSerialize["commits_url"] = o.CommitsUrl
	}
	if true {
		toSerialize["review_comments_url"] = o.ReviewCommentsUrl
	}
	if true {
		toSerialize["review_comment_url"] = o.ReviewCommentUrl
	}
	if true {
		toSerialize["comments_url"] = o.CommentsUrl
	}
	if true {
		toSerialize["statuses_url"] = o.StatusesUrl
	}
	if true {
		toSerialize["number"] = o.Number
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["locked"] = o.Locked
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["user"] = o.User.Get()
	}
	if true {
		toSerialize["body"] = o.Body.Get()
	}
	if true {
		toSerialize["labels"] = o.Labels
	}
	if true {
		toSerialize["milestone"] = o.Milestone.Get()
	}
	if o.ActiveLockReason.IsSet() {
		toSerialize["active_lock_reason"] = o.ActiveLockReason.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["closed_at"] = o.ClosedAt.Get()
	}
	if true {
		toSerialize["merged_at"] = o.MergedAt.Get()
	}
	if true {
		toSerialize["merge_commit_sha"] = o.MergeCommitSha.Get()
	}
	if true {
		toSerialize["assignee"] = o.Assignee.Get()
	}
	if o.Assignees != nil {
		toSerialize["assignees"] = o.Assignees
	}
	if o.RequestedReviewers != nil {
		toSerialize["requested_reviewers"] = o.RequestedReviewers
	}
	if o.RequestedTeams != nil {
		toSerialize["requested_teams"] = o.RequestedTeams
	}
	if true {
		toSerialize["head"] = o.Head
	}
	if true {
		toSerialize["base"] = o.Base
	}
	if true {
		toSerialize["_links"] = o.Links
	}
	if true {
		toSerialize["author_association"] = o.AuthorAssociation
	}
	if true {
		toSerialize["auto_merge"] = o.AutoMerge.Get()
	}
	if o.Draft != nil {
		toSerialize["draft"] = o.Draft
	}
	return json.Marshal(toSerialize)
}

type NullablePullRequestSimple struct {
	value *PullRequestSimple
	isSet bool
}

func (v NullablePullRequestSimple) Get() *PullRequestSimple {
	return v.value
}

func (v *NullablePullRequestSimple) Set(val *PullRequestSimple) {
	v.value = val
	v.isSet = true
}

func (v NullablePullRequestSimple) IsSet() bool {
	return v.isSet
}

func (v *NullablePullRequestSimple) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullRequestSimple(val *PullRequestSimple) *NullablePullRequestSimple {
	return &NullablePullRequestSimple{value: val, isSet: true}
}

func (v NullablePullRequestSimple) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullRequestSimple) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

