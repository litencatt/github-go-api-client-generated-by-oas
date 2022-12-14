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

// WorkflowRun An invocation of a workflow
type WorkflowRun struct {
	// The ID of the workflow run.
	Id int32 `json:"id"`
	// The name of the workflow run.
	Name NullableString `json:"name,omitempty"`
	NodeId string `json:"node_id"`
	// The ID of the associated check suite.
	CheckSuiteId *int32 `json:"check_suite_id,omitempty"`
	// The node ID of the associated check suite.
	CheckSuiteNodeId *string `json:"check_suite_node_id,omitempty"`
	HeadBranch NullableString `json:"head_branch"`
	// The SHA of the head commit that points to the version of the workflow being run.
	HeadSha string `json:"head_sha"`
	// The full path of the workflow
	Path string `json:"path"`
	// The auto incrementing run number for the workflow run.
	RunNumber int32 `json:"run_number"`
	// Attempt number of the run, 1 for first attempt and higher if the workflow was re-run.
	RunAttempt *int32 `json:"run_attempt,omitempty"`
	ReferencedWorkflows []ReferencedWorkflow `json:"referenced_workflows,omitempty"`
	Event string `json:"event"`
	Status NullableString `json:"status"`
	Conclusion NullableString `json:"conclusion"`
	// The ID of the parent workflow.
	WorkflowId int32 `json:"workflow_id"`
	// The URL to the workflow run.
	Url string `json:"url"`
	HtmlUrl string `json:"html_url"`
	PullRequests []PullRequestMinimal `json:"pull_requests"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Actor *SimpleUser `json:"actor,omitempty"`
	TriggeringActor *SimpleUser `json:"triggering_actor,omitempty"`
	// The start time of the latest run. Resets on re-run.
	RunStartedAt *time.Time `json:"run_started_at,omitempty"`
	// The URL to the jobs for the workflow run.
	JobsUrl string `json:"jobs_url"`
	// The URL to download the logs for the workflow run.
	LogsUrl string `json:"logs_url"`
	// The URL to the associated check suite.
	CheckSuiteUrl string `json:"check_suite_url"`
	// The URL to the artifacts for the workflow run.
	ArtifactsUrl string `json:"artifacts_url"`
	// The URL to cancel the workflow run.
	CancelUrl string `json:"cancel_url"`
	// The URL to rerun the workflow run.
	RerunUrl string `json:"rerun_url"`
	// The URL to the previous attempted run of this workflow, if one exists.
	PreviousAttemptUrl NullableString `json:"previous_attempt_url,omitempty"`
	// The URL to the workflow.
	WorkflowUrl string `json:"workflow_url"`
	HeadCommit NullableNullableSimpleCommit `json:"head_commit"`
	Repository MinimalRepository `json:"repository"`
	HeadRepository MinimalRepository `json:"head_repository"`
	HeadRepositoryId *int32 `json:"head_repository_id,omitempty"`
}

// NewWorkflowRun instantiates a new WorkflowRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRun(id int32, nodeId string, headBranch NullableString, headSha string, path string, runNumber int32, event string, status NullableString, conclusion NullableString, workflowId int32, url string, htmlUrl string, pullRequests []PullRequestMinimal, createdAt time.Time, updatedAt time.Time, jobsUrl string, logsUrl string, checkSuiteUrl string, artifactsUrl string, cancelUrl string, rerunUrl string, workflowUrl string, headCommit NullableNullableSimpleCommit, repository MinimalRepository, headRepository MinimalRepository) *WorkflowRun {
	this := WorkflowRun{}
	this.Id = id
	this.NodeId = nodeId
	this.HeadBranch = headBranch
	this.HeadSha = headSha
	this.Path = path
	this.RunNumber = runNumber
	this.Event = event
	this.Status = status
	this.Conclusion = conclusion
	this.WorkflowId = workflowId
	this.Url = url
	this.HtmlUrl = htmlUrl
	this.PullRequests = pullRequests
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.JobsUrl = jobsUrl
	this.LogsUrl = logsUrl
	this.CheckSuiteUrl = checkSuiteUrl
	this.ArtifactsUrl = artifactsUrl
	this.CancelUrl = cancelUrl
	this.RerunUrl = rerunUrl
	this.WorkflowUrl = workflowUrl
	this.HeadCommit = headCommit
	this.Repository = repository
	this.HeadRepository = headRepository
	return &this
}

// NewWorkflowRunWithDefaults instantiates a new WorkflowRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRunWithDefaults() *WorkflowRun {
	this := WorkflowRun{}
	return &this
}

// GetId returns the Id field value
func (o *WorkflowRun) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkflowRun) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowRun) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowRun) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowRun) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *WorkflowRun) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *WorkflowRun) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *WorkflowRun) UnsetName() {
	o.Name.Unset()
}

// GetNodeId returns the NodeId field value
func (o *WorkflowRun) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *WorkflowRun) SetNodeId(v string) {
	o.NodeId = v
}

// GetCheckSuiteId returns the CheckSuiteId field value if set, zero value otherwise.
func (o *WorkflowRun) GetCheckSuiteId() int32 {
	if o == nil || o.CheckSuiteId == nil {
		var ret int32
		return ret
	}
	return *o.CheckSuiteId
}

// GetCheckSuiteIdOk returns a tuple with the CheckSuiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetCheckSuiteIdOk() (*int32, bool) {
	if o == nil || o.CheckSuiteId == nil {
		return nil, false
	}
	return o.CheckSuiteId, true
}

// HasCheckSuiteId returns a boolean if a field has been set.
func (o *WorkflowRun) HasCheckSuiteId() bool {
	if o != nil && o.CheckSuiteId != nil {
		return true
	}

	return false
}

// SetCheckSuiteId gets a reference to the given int32 and assigns it to the CheckSuiteId field.
func (o *WorkflowRun) SetCheckSuiteId(v int32) {
	o.CheckSuiteId = &v
}

// GetCheckSuiteNodeId returns the CheckSuiteNodeId field value if set, zero value otherwise.
func (o *WorkflowRun) GetCheckSuiteNodeId() string {
	if o == nil || o.CheckSuiteNodeId == nil {
		var ret string
		return ret
	}
	return *o.CheckSuiteNodeId
}

// GetCheckSuiteNodeIdOk returns a tuple with the CheckSuiteNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetCheckSuiteNodeIdOk() (*string, bool) {
	if o == nil || o.CheckSuiteNodeId == nil {
		return nil, false
	}
	return o.CheckSuiteNodeId, true
}

// HasCheckSuiteNodeId returns a boolean if a field has been set.
func (o *WorkflowRun) HasCheckSuiteNodeId() bool {
	if o != nil && o.CheckSuiteNodeId != nil {
		return true
	}

	return false
}

// SetCheckSuiteNodeId gets a reference to the given string and assigns it to the CheckSuiteNodeId field.
func (o *WorkflowRun) SetCheckSuiteNodeId(v string) {
	o.CheckSuiteNodeId = &v
}

// GetHeadBranch returns the HeadBranch field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WorkflowRun) GetHeadBranch() string {
	if o == nil || o.HeadBranch.Get() == nil {
		var ret string
		return ret
	}

	return *o.HeadBranch.Get()
}

// GetHeadBranchOk returns a tuple with the HeadBranch field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowRun) GetHeadBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeadBranch.Get(), o.HeadBranch.IsSet()
}

// SetHeadBranch sets field value
func (o *WorkflowRun) SetHeadBranch(v string) {
	o.HeadBranch.Set(&v)
}

// GetHeadSha returns the HeadSha field value
func (o *WorkflowRun) GetHeadSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HeadSha
}

// GetHeadShaOk returns a tuple with the HeadSha field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetHeadShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeadSha, true
}

// SetHeadSha sets field value
func (o *WorkflowRun) SetHeadSha(v string) {
	o.HeadSha = v
}

// GetPath returns the Path field value
func (o *WorkflowRun) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *WorkflowRun) SetPath(v string) {
	o.Path = v
}

// GetRunNumber returns the RunNumber field value
func (o *WorkflowRun) GetRunNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RunNumber
}

// GetRunNumberOk returns a tuple with the RunNumber field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetRunNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunNumber, true
}

// SetRunNumber sets field value
func (o *WorkflowRun) SetRunNumber(v int32) {
	o.RunNumber = v
}

// GetRunAttempt returns the RunAttempt field value if set, zero value otherwise.
func (o *WorkflowRun) GetRunAttempt() int32 {
	if o == nil || o.RunAttempt == nil {
		var ret int32
		return ret
	}
	return *o.RunAttempt
}

// GetRunAttemptOk returns a tuple with the RunAttempt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetRunAttemptOk() (*int32, bool) {
	if o == nil || o.RunAttempt == nil {
		return nil, false
	}
	return o.RunAttempt, true
}

// HasRunAttempt returns a boolean if a field has been set.
func (o *WorkflowRun) HasRunAttempt() bool {
	if o != nil && o.RunAttempt != nil {
		return true
	}

	return false
}

// SetRunAttempt gets a reference to the given int32 and assigns it to the RunAttempt field.
func (o *WorkflowRun) SetRunAttempt(v int32) {
	o.RunAttempt = &v
}

// GetReferencedWorkflows returns the ReferencedWorkflows field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowRun) GetReferencedWorkflows() []ReferencedWorkflow {
	if o == nil {
		var ret []ReferencedWorkflow
		return ret
	}
	return o.ReferencedWorkflows
}

// GetReferencedWorkflowsOk returns a tuple with the ReferencedWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowRun) GetReferencedWorkflowsOk() ([]ReferencedWorkflow, bool) {
	if o == nil || o.ReferencedWorkflows == nil {
		return nil, false
	}
	return o.ReferencedWorkflows, true
}

// HasReferencedWorkflows returns a boolean if a field has been set.
func (o *WorkflowRun) HasReferencedWorkflows() bool {
	if o != nil && o.ReferencedWorkflows != nil {
		return true
	}

	return false
}

// SetReferencedWorkflows gets a reference to the given []ReferencedWorkflow and assigns it to the ReferencedWorkflows field.
func (o *WorkflowRun) SetReferencedWorkflows(v []ReferencedWorkflow) {
	o.ReferencedWorkflows = v
}

// GetEvent returns the Event field value
func (o *WorkflowRun) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *WorkflowRun) SetEvent(v string) {
	o.Event = v
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WorkflowRun) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}

	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowRun) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// SetStatus sets field value
func (o *WorkflowRun) SetStatus(v string) {
	o.Status.Set(&v)
}

// GetConclusion returns the Conclusion field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WorkflowRun) GetConclusion() string {
	if o == nil || o.Conclusion.Get() == nil {
		var ret string
		return ret
	}

	return *o.Conclusion.Get()
}

// GetConclusionOk returns a tuple with the Conclusion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowRun) GetConclusionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conclusion.Get(), o.Conclusion.IsSet()
}

// SetConclusion sets field value
func (o *WorkflowRun) SetConclusion(v string) {
	o.Conclusion.Set(&v)
}

// GetWorkflowId returns the WorkflowId field value
func (o *WorkflowRun) GetWorkflowId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetWorkflowIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowId, true
}

// SetWorkflowId sets field value
func (o *WorkflowRun) SetWorkflowId(v int32) {
	o.WorkflowId = v
}

// GetUrl returns the Url field value
func (o *WorkflowRun) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WorkflowRun) SetUrl(v string) {
	o.Url = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *WorkflowRun) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *WorkflowRun) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetPullRequests returns the PullRequests field value
// If the value is explicit nil, the zero value for []PullRequestMinimal will be returned
func (o *WorkflowRun) GetPullRequests() []PullRequestMinimal {
	if o == nil {
		var ret []PullRequestMinimal
		return ret
	}

	return o.PullRequests
}

// GetPullRequestsOk returns a tuple with the PullRequests field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowRun) GetPullRequestsOk() ([]PullRequestMinimal, bool) {
	if o == nil || o.PullRequests == nil {
		return nil, false
	}
	return o.PullRequests, true
}

// SetPullRequests sets field value
func (o *WorkflowRun) SetPullRequests(v []PullRequestMinimal) {
	o.PullRequests = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WorkflowRun) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WorkflowRun) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *WorkflowRun) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *WorkflowRun) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *WorkflowRun) GetActor() SimpleUser {
	if o == nil || o.Actor == nil {
		var ret SimpleUser
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetActorOk() (*SimpleUser, bool) {
	if o == nil || o.Actor == nil {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *WorkflowRun) HasActor() bool {
	if o != nil && o.Actor != nil {
		return true
	}

	return false
}

// SetActor gets a reference to the given SimpleUser and assigns it to the Actor field.
func (o *WorkflowRun) SetActor(v SimpleUser) {
	o.Actor = &v
}

// GetTriggeringActor returns the TriggeringActor field value if set, zero value otherwise.
func (o *WorkflowRun) GetTriggeringActor() SimpleUser {
	if o == nil || o.TriggeringActor == nil {
		var ret SimpleUser
		return ret
	}
	return *o.TriggeringActor
}

// GetTriggeringActorOk returns a tuple with the TriggeringActor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetTriggeringActorOk() (*SimpleUser, bool) {
	if o == nil || o.TriggeringActor == nil {
		return nil, false
	}
	return o.TriggeringActor, true
}

// HasTriggeringActor returns a boolean if a field has been set.
func (o *WorkflowRun) HasTriggeringActor() bool {
	if o != nil && o.TriggeringActor != nil {
		return true
	}

	return false
}

// SetTriggeringActor gets a reference to the given SimpleUser and assigns it to the TriggeringActor field.
func (o *WorkflowRun) SetTriggeringActor(v SimpleUser) {
	o.TriggeringActor = &v
}

// GetRunStartedAt returns the RunStartedAt field value if set, zero value otherwise.
func (o *WorkflowRun) GetRunStartedAt() time.Time {
	if o == nil || o.RunStartedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.RunStartedAt
}

// GetRunStartedAtOk returns a tuple with the RunStartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetRunStartedAtOk() (*time.Time, bool) {
	if o == nil || o.RunStartedAt == nil {
		return nil, false
	}
	return o.RunStartedAt, true
}

// HasRunStartedAt returns a boolean if a field has been set.
func (o *WorkflowRun) HasRunStartedAt() bool {
	if o != nil && o.RunStartedAt != nil {
		return true
	}

	return false
}

// SetRunStartedAt gets a reference to the given time.Time and assigns it to the RunStartedAt field.
func (o *WorkflowRun) SetRunStartedAt(v time.Time) {
	o.RunStartedAt = &v
}

// GetJobsUrl returns the JobsUrl field value
func (o *WorkflowRun) GetJobsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobsUrl
}

// GetJobsUrlOk returns a tuple with the JobsUrl field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetJobsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobsUrl, true
}

// SetJobsUrl sets field value
func (o *WorkflowRun) SetJobsUrl(v string) {
	o.JobsUrl = v
}

// GetLogsUrl returns the LogsUrl field value
func (o *WorkflowRun) GetLogsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogsUrl
}

// GetLogsUrlOk returns a tuple with the LogsUrl field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetLogsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogsUrl, true
}

// SetLogsUrl sets field value
func (o *WorkflowRun) SetLogsUrl(v string) {
	o.LogsUrl = v
}

// GetCheckSuiteUrl returns the CheckSuiteUrl field value
func (o *WorkflowRun) GetCheckSuiteUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckSuiteUrl
}

// GetCheckSuiteUrlOk returns a tuple with the CheckSuiteUrl field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetCheckSuiteUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckSuiteUrl, true
}

// SetCheckSuiteUrl sets field value
func (o *WorkflowRun) SetCheckSuiteUrl(v string) {
	o.CheckSuiteUrl = v
}

// GetArtifactsUrl returns the ArtifactsUrl field value
func (o *WorkflowRun) GetArtifactsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArtifactsUrl
}

// GetArtifactsUrlOk returns a tuple with the ArtifactsUrl field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetArtifactsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtifactsUrl, true
}

// SetArtifactsUrl sets field value
func (o *WorkflowRun) SetArtifactsUrl(v string) {
	o.ArtifactsUrl = v
}

// GetCancelUrl returns the CancelUrl field value
func (o *WorkflowRun) GetCancelUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CancelUrl
}

// GetCancelUrlOk returns a tuple with the CancelUrl field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetCancelUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CancelUrl, true
}

// SetCancelUrl sets field value
func (o *WorkflowRun) SetCancelUrl(v string) {
	o.CancelUrl = v
}

// GetRerunUrl returns the RerunUrl field value
func (o *WorkflowRun) GetRerunUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RerunUrl
}

// GetRerunUrlOk returns a tuple with the RerunUrl field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetRerunUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RerunUrl, true
}

// SetRerunUrl sets field value
func (o *WorkflowRun) SetRerunUrl(v string) {
	o.RerunUrl = v
}

// GetPreviousAttemptUrl returns the PreviousAttemptUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowRun) GetPreviousAttemptUrl() string {
	if o == nil || o.PreviousAttemptUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.PreviousAttemptUrl.Get()
}

// GetPreviousAttemptUrlOk returns a tuple with the PreviousAttemptUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowRun) GetPreviousAttemptUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousAttemptUrl.Get(), o.PreviousAttemptUrl.IsSet()
}

// HasPreviousAttemptUrl returns a boolean if a field has been set.
func (o *WorkflowRun) HasPreviousAttemptUrl() bool {
	if o != nil && o.PreviousAttemptUrl.IsSet() {
		return true
	}

	return false
}

// SetPreviousAttemptUrl gets a reference to the given NullableString and assigns it to the PreviousAttemptUrl field.
func (o *WorkflowRun) SetPreviousAttemptUrl(v string) {
	o.PreviousAttemptUrl.Set(&v)
}
// SetPreviousAttemptUrlNil sets the value for PreviousAttemptUrl to be an explicit nil
func (o *WorkflowRun) SetPreviousAttemptUrlNil() {
	o.PreviousAttemptUrl.Set(nil)
}

// UnsetPreviousAttemptUrl ensures that no value is present for PreviousAttemptUrl, not even an explicit nil
func (o *WorkflowRun) UnsetPreviousAttemptUrl() {
	o.PreviousAttemptUrl.Unset()
}

// GetWorkflowUrl returns the WorkflowUrl field value
func (o *WorkflowRun) GetWorkflowUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowUrl
}

// GetWorkflowUrlOk returns a tuple with the WorkflowUrl field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetWorkflowUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowUrl, true
}

// SetWorkflowUrl sets field value
func (o *WorkflowRun) SetWorkflowUrl(v string) {
	o.WorkflowUrl = v
}

// GetHeadCommit returns the HeadCommit field value
// If the value is explicit nil, the zero value for NullableSimpleCommit will be returned
func (o *WorkflowRun) GetHeadCommit() NullableSimpleCommit {
	if o == nil || o.HeadCommit.Get() == nil {
		var ret NullableSimpleCommit
		return ret
	}

	return *o.HeadCommit.Get()
}

// GetHeadCommitOk returns a tuple with the HeadCommit field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowRun) GetHeadCommitOk() (*NullableSimpleCommit, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeadCommit.Get(), o.HeadCommit.IsSet()
}

// SetHeadCommit sets field value
func (o *WorkflowRun) SetHeadCommit(v NullableSimpleCommit) {
	o.HeadCommit.Set(&v)
}

// GetRepository returns the Repository field value
func (o *WorkflowRun) GetRepository() MinimalRepository {
	if o == nil {
		var ret MinimalRepository
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetRepositoryOk() (*MinimalRepository, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *WorkflowRun) SetRepository(v MinimalRepository) {
	o.Repository = v
}

// GetHeadRepository returns the HeadRepository field value
func (o *WorkflowRun) GetHeadRepository() MinimalRepository {
	if o == nil {
		var ret MinimalRepository
		return ret
	}

	return o.HeadRepository
}

// GetHeadRepositoryOk returns a tuple with the HeadRepository field value
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetHeadRepositoryOk() (*MinimalRepository, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeadRepository, true
}

// SetHeadRepository sets field value
func (o *WorkflowRun) SetHeadRepository(v MinimalRepository) {
	o.HeadRepository = v
}

// GetHeadRepositoryId returns the HeadRepositoryId field value if set, zero value otherwise.
func (o *WorkflowRun) GetHeadRepositoryId() int32 {
	if o == nil || o.HeadRepositoryId == nil {
		var ret int32
		return ret
	}
	return *o.HeadRepositoryId
}

// GetHeadRepositoryIdOk returns a tuple with the HeadRepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRun) GetHeadRepositoryIdOk() (*int32, bool) {
	if o == nil || o.HeadRepositoryId == nil {
		return nil, false
	}
	return o.HeadRepositoryId, true
}

// HasHeadRepositoryId returns a boolean if a field has been set.
func (o *WorkflowRun) HasHeadRepositoryId() bool {
	if o != nil && o.HeadRepositoryId != nil {
		return true
	}

	return false
}

// SetHeadRepositoryId gets a reference to the given int32 and assigns it to the HeadRepositoryId field.
func (o *WorkflowRun) SetHeadRepositoryId(v int32) {
	o.HeadRepositoryId = &v
}

func (o WorkflowRun) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	if o.CheckSuiteId != nil {
		toSerialize["check_suite_id"] = o.CheckSuiteId
	}
	if o.CheckSuiteNodeId != nil {
		toSerialize["check_suite_node_id"] = o.CheckSuiteNodeId
	}
	if true {
		toSerialize["head_branch"] = o.HeadBranch.Get()
	}
	if true {
		toSerialize["head_sha"] = o.HeadSha
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["run_number"] = o.RunNumber
	}
	if o.RunAttempt != nil {
		toSerialize["run_attempt"] = o.RunAttempt
	}
	if o.ReferencedWorkflows != nil {
		toSerialize["referenced_workflows"] = o.ReferencedWorkflows
	}
	if true {
		toSerialize["event"] = o.Event
	}
	if true {
		toSerialize["status"] = o.Status.Get()
	}
	if true {
		toSerialize["conclusion"] = o.Conclusion.Get()
	}
	if true {
		toSerialize["workflow_id"] = o.WorkflowId
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if o.PullRequests != nil {
		toSerialize["pull_requests"] = o.PullRequests
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Actor != nil {
		toSerialize["actor"] = o.Actor
	}
	if o.TriggeringActor != nil {
		toSerialize["triggering_actor"] = o.TriggeringActor
	}
	if o.RunStartedAt != nil {
		toSerialize["run_started_at"] = o.RunStartedAt
	}
	if true {
		toSerialize["jobs_url"] = o.JobsUrl
	}
	if true {
		toSerialize["logs_url"] = o.LogsUrl
	}
	if true {
		toSerialize["check_suite_url"] = o.CheckSuiteUrl
	}
	if true {
		toSerialize["artifacts_url"] = o.ArtifactsUrl
	}
	if true {
		toSerialize["cancel_url"] = o.CancelUrl
	}
	if true {
		toSerialize["rerun_url"] = o.RerunUrl
	}
	if o.PreviousAttemptUrl.IsSet() {
		toSerialize["previous_attempt_url"] = o.PreviousAttemptUrl.Get()
	}
	if true {
		toSerialize["workflow_url"] = o.WorkflowUrl
	}
	if true {
		toSerialize["head_commit"] = o.HeadCommit.Get()
	}
	if true {
		toSerialize["repository"] = o.Repository
	}
	if true {
		toSerialize["head_repository"] = o.HeadRepository
	}
	if o.HeadRepositoryId != nil {
		toSerialize["head_repository_id"] = o.HeadRepositoryId
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowRun struct {
	value *WorkflowRun
	isSet bool
}

func (v NullableWorkflowRun) Get() *WorkflowRun {
	return v.value
}

func (v *NullableWorkflowRun) Set(val *WorkflowRun) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRun) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRun(val *WorkflowRun) *NullableWorkflowRun {
	return &NullableWorkflowRun{value: val, isSet: true}
}

func (v NullableWorkflowRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


