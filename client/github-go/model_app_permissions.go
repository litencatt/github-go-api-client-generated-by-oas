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

// AppPermissions The permissions granted to the user-to-server access token.
type AppPermissions struct {
	// The level of permission to grant the access token for GitHub Actions workflows, workflow runs, and artifacts.
	Actions *string `json:"actions,omitempty"`
	// The level of permission to grant the access token for repository creation, deletion, settings, teams, and collaborators creation.
	Administration *string `json:"administration,omitempty"`
	// The level of permission to grant the access token for checks on code.
	Checks *string `json:"checks,omitempty"`
	// The level of permission to grant the access token for repository contents, commits, branches, downloads, releases, and merges.
	Contents *string `json:"contents,omitempty"`
	// The level of permission to grant the access token for deployments and deployment statuses.
	Deployments *string `json:"deployments,omitempty"`
	// The level of permission to grant the access token for managing repository environments.
	Environments *string `json:"environments,omitempty"`
	// The level of permission to grant the access token for issues and related comments, assignees, labels, and milestones.
	Issues *string `json:"issues,omitempty"`
	// The level of permission to grant the access token to search repositories, list collaborators, and access repository metadata.
	Metadata *string `json:"metadata,omitempty"`
	// The level of permission to grant the access token for packages published to GitHub Packages.
	Packages *string `json:"packages,omitempty"`
	// The level of permission to grant the access token to retrieve Pages statuses, configuration, and builds, as well as create new builds.
	Pages *string `json:"pages,omitempty"`
	// The level of permission to grant the access token for pull requests and related comments, assignees, labels, milestones, and merges.
	PullRequests *string `json:"pull_requests,omitempty"`
	// The level of permission to grant the access token to manage the post-receive hooks for a repository.
	RepositoryHooks *string `json:"repository_hooks,omitempty"`
	// The level of permission to grant the access token to manage repository projects, columns, and cards.
	RepositoryProjects *string `json:"repository_projects,omitempty"`
	// The level of permission to grant the access token to view and manage secret scanning alerts.
	SecretScanningAlerts *string `json:"secret_scanning_alerts,omitempty"`
	// The level of permission to grant the access token to manage repository secrets.
	Secrets *string `json:"secrets,omitempty"`
	// The level of permission to grant the access token to view and manage security events like code scanning alerts.
	SecurityEvents *string `json:"security_events,omitempty"`
	// The level of permission to grant the access token to manage just a single file.
	SingleFile *string `json:"single_file,omitempty"`
	// The level of permission to grant the access token for commit statuses.
	Statuses *string `json:"statuses,omitempty"`
	// The level of permission to grant the access token to manage Dependabot alerts.
	VulnerabilityAlerts *string `json:"vulnerability_alerts,omitempty"`
	// The level of permission to grant the access token to update GitHub Actions workflow files.
	Workflows *string `json:"workflows,omitempty"`
	// The level of permission to grant the access token for organization teams and members.
	Members *string `json:"members,omitempty"`
	// The level of permission to grant the access token to manage access to an organization.
	OrganizationAdministration *string `json:"organization_administration,omitempty"`
	// The level of permission to grant the access token to manage the post-receive hooks for an organization.
	OrganizationHooks *string `json:"organization_hooks,omitempty"`
	// The level of permission to grant the access token for viewing an organization's plan.
	OrganizationPlan *string `json:"organization_plan,omitempty"`
	// The level of permission to grant the access token to manage organization projects and projects beta (where available).
	OrganizationProjects *string `json:"organization_projects,omitempty"`
	// The level of permission to grant the access token for organization packages published to GitHub Packages.
	OrganizationPackages *string `json:"organization_packages,omitempty"`
	// The level of permission to grant the access token to manage organization secrets.
	OrganizationSecrets *string `json:"organization_secrets,omitempty"`
	// The level of permission to grant the access token to view and manage GitHub Actions self-hosted runners available to an organization.
	OrganizationSelfHostedRunners *string `json:"organization_self_hosted_runners,omitempty"`
	// The level of permission to grant the access token to view and manage users blocked by the organization.
	OrganizationUserBlocking *string `json:"organization_user_blocking,omitempty"`
	// The level of permission to grant the access token to manage team discussions and related comments.
	TeamDiscussions *string `json:"team_discussions,omitempty"`
}

// NewAppPermissions instantiates a new AppPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPermissions() *AppPermissions {
	this := AppPermissions{}
	return &this
}

// NewAppPermissionsWithDefaults instantiates a new AppPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPermissionsWithDefaults() *AppPermissions {
	this := AppPermissions{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *AppPermissions) GetActions() string {
	if o == nil || o.Actions == nil {
		var ret string
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetActionsOk() (*string, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *AppPermissions) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given string and assigns it to the Actions field.
func (o *AppPermissions) SetActions(v string) {
	o.Actions = &v
}

// GetAdministration returns the Administration field value if set, zero value otherwise.
func (o *AppPermissions) GetAdministration() string {
	if o == nil || o.Administration == nil {
		var ret string
		return ret
	}
	return *o.Administration
}

// GetAdministrationOk returns a tuple with the Administration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetAdministrationOk() (*string, bool) {
	if o == nil || o.Administration == nil {
		return nil, false
	}
	return o.Administration, true
}

// HasAdministration returns a boolean if a field has been set.
func (o *AppPermissions) HasAdministration() bool {
	if o != nil && o.Administration != nil {
		return true
	}

	return false
}

// SetAdministration gets a reference to the given string and assigns it to the Administration field.
func (o *AppPermissions) SetAdministration(v string) {
	o.Administration = &v
}

// GetChecks returns the Checks field value if set, zero value otherwise.
func (o *AppPermissions) GetChecks() string {
	if o == nil || o.Checks == nil {
		var ret string
		return ret
	}
	return *o.Checks
}

// GetChecksOk returns a tuple with the Checks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetChecksOk() (*string, bool) {
	if o == nil || o.Checks == nil {
		return nil, false
	}
	return o.Checks, true
}

// HasChecks returns a boolean if a field has been set.
func (o *AppPermissions) HasChecks() bool {
	if o != nil && o.Checks != nil {
		return true
	}

	return false
}

// SetChecks gets a reference to the given string and assigns it to the Checks field.
func (o *AppPermissions) SetChecks(v string) {
	o.Checks = &v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *AppPermissions) GetContents() string {
	if o == nil || o.Contents == nil {
		var ret string
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetContentsOk() (*string, bool) {
	if o == nil || o.Contents == nil {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *AppPermissions) HasContents() bool {
	if o != nil && o.Contents != nil {
		return true
	}

	return false
}

// SetContents gets a reference to the given string and assigns it to the Contents field.
func (o *AppPermissions) SetContents(v string) {
	o.Contents = &v
}

// GetDeployments returns the Deployments field value if set, zero value otherwise.
func (o *AppPermissions) GetDeployments() string {
	if o == nil || o.Deployments == nil {
		var ret string
		return ret
	}
	return *o.Deployments
}

// GetDeploymentsOk returns a tuple with the Deployments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetDeploymentsOk() (*string, bool) {
	if o == nil || o.Deployments == nil {
		return nil, false
	}
	return o.Deployments, true
}

// HasDeployments returns a boolean if a field has been set.
func (o *AppPermissions) HasDeployments() bool {
	if o != nil && o.Deployments != nil {
		return true
	}

	return false
}

// SetDeployments gets a reference to the given string and assigns it to the Deployments field.
func (o *AppPermissions) SetDeployments(v string) {
	o.Deployments = &v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *AppPermissions) GetEnvironments() string {
	if o == nil || o.Environments == nil {
		var ret string
		return ret
	}
	return *o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetEnvironmentsOk() (*string, bool) {
	if o == nil || o.Environments == nil {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *AppPermissions) HasEnvironments() bool {
	if o != nil && o.Environments != nil {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given string and assigns it to the Environments field.
func (o *AppPermissions) SetEnvironments(v string) {
	o.Environments = &v
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *AppPermissions) GetIssues() string {
	if o == nil || o.Issues == nil {
		var ret string
		return ret
	}
	return *o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetIssuesOk() (*string, bool) {
	if o == nil || o.Issues == nil {
		return nil, false
	}
	return o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *AppPermissions) HasIssues() bool {
	if o != nil && o.Issues != nil {
		return true
	}

	return false
}

// SetIssues gets a reference to the given string and assigns it to the Issues field.
func (o *AppPermissions) SetIssues(v string) {
	o.Issues = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AppPermissions) GetMetadata() string {
	if o == nil || o.Metadata == nil {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetMetadataOk() (*string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AppPermissions) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *AppPermissions) SetMetadata(v string) {
	o.Metadata = &v
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *AppPermissions) GetPackages() string {
	if o == nil || o.Packages == nil {
		var ret string
		return ret
	}
	return *o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetPackagesOk() (*string, bool) {
	if o == nil || o.Packages == nil {
		return nil, false
	}
	return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *AppPermissions) HasPackages() bool {
	if o != nil && o.Packages != nil {
		return true
	}

	return false
}

// SetPackages gets a reference to the given string and assigns it to the Packages field.
func (o *AppPermissions) SetPackages(v string) {
	o.Packages = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *AppPermissions) GetPages() string {
	if o == nil || o.Pages == nil {
		var ret string
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetPagesOk() (*string, bool) {
	if o == nil || o.Pages == nil {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *AppPermissions) HasPages() bool {
	if o != nil && o.Pages != nil {
		return true
	}

	return false
}

// SetPages gets a reference to the given string and assigns it to the Pages field.
func (o *AppPermissions) SetPages(v string) {
	o.Pages = &v
}

// GetPullRequests returns the PullRequests field value if set, zero value otherwise.
func (o *AppPermissions) GetPullRequests() string {
	if o == nil || o.PullRequests == nil {
		var ret string
		return ret
	}
	return *o.PullRequests
}

// GetPullRequestsOk returns a tuple with the PullRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetPullRequestsOk() (*string, bool) {
	if o == nil || o.PullRequests == nil {
		return nil, false
	}
	return o.PullRequests, true
}

// HasPullRequests returns a boolean if a field has been set.
func (o *AppPermissions) HasPullRequests() bool {
	if o != nil && o.PullRequests != nil {
		return true
	}

	return false
}

// SetPullRequests gets a reference to the given string and assigns it to the PullRequests field.
func (o *AppPermissions) SetPullRequests(v string) {
	o.PullRequests = &v
}

// GetRepositoryHooks returns the RepositoryHooks field value if set, zero value otherwise.
func (o *AppPermissions) GetRepositoryHooks() string {
	if o == nil || o.RepositoryHooks == nil {
		var ret string
		return ret
	}
	return *o.RepositoryHooks
}

// GetRepositoryHooksOk returns a tuple with the RepositoryHooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetRepositoryHooksOk() (*string, bool) {
	if o == nil || o.RepositoryHooks == nil {
		return nil, false
	}
	return o.RepositoryHooks, true
}

// HasRepositoryHooks returns a boolean if a field has been set.
func (o *AppPermissions) HasRepositoryHooks() bool {
	if o != nil && o.RepositoryHooks != nil {
		return true
	}

	return false
}

// SetRepositoryHooks gets a reference to the given string and assigns it to the RepositoryHooks field.
func (o *AppPermissions) SetRepositoryHooks(v string) {
	o.RepositoryHooks = &v
}

// GetRepositoryProjects returns the RepositoryProjects field value if set, zero value otherwise.
func (o *AppPermissions) GetRepositoryProjects() string {
	if o == nil || o.RepositoryProjects == nil {
		var ret string
		return ret
	}
	return *o.RepositoryProjects
}

// GetRepositoryProjectsOk returns a tuple with the RepositoryProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetRepositoryProjectsOk() (*string, bool) {
	if o == nil || o.RepositoryProjects == nil {
		return nil, false
	}
	return o.RepositoryProjects, true
}

// HasRepositoryProjects returns a boolean if a field has been set.
func (o *AppPermissions) HasRepositoryProjects() bool {
	if o != nil && o.RepositoryProjects != nil {
		return true
	}

	return false
}

// SetRepositoryProjects gets a reference to the given string and assigns it to the RepositoryProjects field.
func (o *AppPermissions) SetRepositoryProjects(v string) {
	o.RepositoryProjects = &v
}

// GetSecretScanningAlerts returns the SecretScanningAlerts field value if set, zero value otherwise.
func (o *AppPermissions) GetSecretScanningAlerts() string {
	if o == nil || o.SecretScanningAlerts == nil {
		var ret string
		return ret
	}
	return *o.SecretScanningAlerts
}

// GetSecretScanningAlertsOk returns a tuple with the SecretScanningAlerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetSecretScanningAlertsOk() (*string, bool) {
	if o == nil || o.SecretScanningAlerts == nil {
		return nil, false
	}
	return o.SecretScanningAlerts, true
}

// HasSecretScanningAlerts returns a boolean if a field has been set.
func (o *AppPermissions) HasSecretScanningAlerts() bool {
	if o != nil && o.SecretScanningAlerts != nil {
		return true
	}

	return false
}

// SetSecretScanningAlerts gets a reference to the given string and assigns it to the SecretScanningAlerts field.
func (o *AppPermissions) SetSecretScanningAlerts(v string) {
	o.SecretScanningAlerts = &v
}

// GetSecrets returns the Secrets field value if set, zero value otherwise.
func (o *AppPermissions) GetSecrets() string {
	if o == nil || o.Secrets == nil {
		var ret string
		return ret
	}
	return *o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetSecretsOk() (*string, bool) {
	if o == nil || o.Secrets == nil {
		return nil, false
	}
	return o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *AppPermissions) HasSecrets() bool {
	if o != nil && o.Secrets != nil {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given string and assigns it to the Secrets field.
func (o *AppPermissions) SetSecrets(v string) {
	o.Secrets = &v
}

// GetSecurityEvents returns the SecurityEvents field value if set, zero value otherwise.
func (o *AppPermissions) GetSecurityEvents() string {
	if o == nil || o.SecurityEvents == nil {
		var ret string
		return ret
	}
	return *o.SecurityEvents
}

// GetSecurityEventsOk returns a tuple with the SecurityEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetSecurityEventsOk() (*string, bool) {
	if o == nil || o.SecurityEvents == nil {
		return nil, false
	}
	return o.SecurityEvents, true
}

// HasSecurityEvents returns a boolean if a field has been set.
func (o *AppPermissions) HasSecurityEvents() bool {
	if o != nil && o.SecurityEvents != nil {
		return true
	}

	return false
}

// SetSecurityEvents gets a reference to the given string and assigns it to the SecurityEvents field.
func (o *AppPermissions) SetSecurityEvents(v string) {
	o.SecurityEvents = &v
}

// GetSingleFile returns the SingleFile field value if set, zero value otherwise.
func (o *AppPermissions) GetSingleFile() string {
	if o == nil || o.SingleFile == nil {
		var ret string
		return ret
	}
	return *o.SingleFile
}

// GetSingleFileOk returns a tuple with the SingleFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetSingleFileOk() (*string, bool) {
	if o == nil || o.SingleFile == nil {
		return nil, false
	}
	return o.SingleFile, true
}

// HasSingleFile returns a boolean if a field has been set.
func (o *AppPermissions) HasSingleFile() bool {
	if o != nil && o.SingleFile != nil {
		return true
	}

	return false
}

// SetSingleFile gets a reference to the given string and assigns it to the SingleFile field.
func (o *AppPermissions) SetSingleFile(v string) {
	o.SingleFile = &v
}

// GetStatuses returns the Statuses field value if set, zero value otherwise.
func (o *AppPermissions) GetStatuses() string {
	if o == nil || o.Statuses == nil {
		var ret string
		return ret
	}
	return *o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetStatusesOk() (*string, bool) {
	if o == nil || o.Statuses == nil {
		return nil, false
	}
	return o.Statuses, true
}

// HasStatuses returns a boolean if a field has been set.
func (o *AppPermissions) HasStatuses() bool {
	if o != nil && o.Statuses != nil {
		return true
	}

	return false
}

// SetStatuses gets a reference to the given string and assigns it to the Statuses field.
func (o *AppPermissions) SetStatuses(v string) {
	o.Statuses = &v
}

// GetVulnerabilityAlerts returns the VulnerabilityAlerts field value if set, zero value otherwise.
func (o *AppPermissions) GetVulnerabilityAlerts() string {
	if o == nil || o.VulnerabilityAlerts == nil {
		var ret string
		return ret
	}
	return *o.VulnerabilityAlerts
}

// GetVulnerabilityAlertsOk returns a tuple with the VulnerabilityAlerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetVulnerabilityAlertsOk() (*string, bool) {
	if o == nil || o.VulnerabilityAlerts == nil {
		return nil, false
	}
	return o.VulnerabilityAlerts, true
}

// HasVulnerabilityAlerts returns a boolean if a field has been set.
func (o *AppPermissions) HasVulnerabilityAlerts() bool {
	if o != nil && o.VulnerabilityAlerts != nil {
		return true
	}

	return false
}

// SetVulnerabilityAlerts gets a reference to the given string and assigns it to the VulnerabilityAlerts field.
func (o *AppPermissions) SetVulnerabilityAlerts(v string) {
	o.VulnerabilityAlerts = &v
}

// GetWorkflows returns the Workflows field value if set, zero value otherwise.
func (o *AppPermissions) GetWorkflows() string {
	if o == nil || o.Workflows == nil {
		var ret string
		return ret
	}
	return *o.Workflows
}

// GetWorkflowsOk returns a tuple with the Workflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetWorkflowsOk() (*string, bool) {
	if o == nil || o.Workflows == nil {
		return nil, false
	}
	return o.Workflows, true
}

// HasWorkflows returns a boolean if a field has been set.
func (o *AppPermissions) HasWorkflows() bool {
	if o != nil && o.Workflows != nil {
		return true
	}

	return false
}

// SetWorkflows gets a reference to the given string and assigns it to the Workflows field.
func (o *AppPermissions) SetWorkflows(v string) {
	o.Workflows = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *AppPermissions) GetMembers() string {
	if o == nil || o.Members == nil {
		var ret string
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetMembersOk() (*string, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *AppPermissions) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given string and assigns it to the Members field.
func (o *AppPermissions) SetMembers(v string) {
	o.Members = &v
}

// GetOrganizationAdministration returns the OrganizationAdministration field value if set, zero value otherwise.
func (o *AppPermissions) GetOrganizationAdministration() string {
	if o == nil || o.OrganizationAdministration == nil {
		var ret string
		return ret
	}
	return *o.OrganizationAdministration
}

// GetOrganizationAdministrationOk returns a tuple with the OrganizationAdministration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetOrganizationAdministrationOk() (*string, bool) {
	if o == nil || o.OrganizationAdministration == nil {
		return nil, false
	}
	return o.OrganizationAdministration, true
}

// HasOrganizationAdministration returns a boolean if a field has been set.
func (o *AppPermissions) HasOrganizationAdministration() bool {
	if o != nil && o.OrganizationAdministration != nil {
		return true
	}

	return false
}

// SetOrganizationAdministration gets a reference to the given string and assigns it to the OrganizationAdministration field.
func (o *AppPermissions) SetOrganizationAdministration(v string) {
	o.OrganizationAdministration = &v
}

// GetOrganizationHooks returns the OrganizationHooks field value if set, zero value otherwise.
func (o *AppPermissions) GetOrganizationHooks() string {
	if o == nil || o.OrganizationHooks == nil {
		var ret string
		return ret
	}
	return *o.OrganizationHooks
}

// GetOrganizationHooksOk returns a tuple with the OrganizationHooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetOrganizationHooksOk() (*string, bool) {
	if o == nil || o.OrganizationHooks == nil {
		return nil, false
	}
	return o.OrganizationHooks, true
}

// HasOrganizationHooks returns a boolean if a field has been set.
func (o *AppPermissions) HasOrganizationHooks() bool {
	if o != nil && o.OrganizationHooks != nil {
		return true
	}

	return false
}

// SetOrganizationHooks gets a reference to the given string and assigns it to the OrganizationHooks field.
func (o *AppPermissions) SetOrganizationHooks(v string) {
	o.OrganizationHooks = &v
}

// GetOrganizationPlan returns the OrganizationPlan field value if set, zero value otherwise.
func (o *AppPermissions) GetOrganizationPlan() string {
	if o == nil || o.OrganizationPlan == nil {
		var ret string
		return ret
	}
	return *o.OrganizationPlan
}

// GetOrganizationPlanOk returns a tuple with the OrganizationPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetOrganizationPlanOk() (*string, bool) {
	if o == nil || o.OrganizationPlan == nil {
		return nil, false
	}
	return o.OrganizationPlan, true
}

// HasOrganizationPlan returns a boolean if a field has been set.
func (o *AppPermissions) HasOrganizationPlan() bool {
	if o != nil && o.OrganizationPlan != nil {
		return true
	}

	return false
}

// SetOrganizationPlan gets a reference to the given string and assigns it to the OrganizationPlan field.
func (o *AppPermissions) SetOrganizationPlan(v string) {
	o.OrganizationPlan = &v
}

// GetOrganizationProjects returns the OrganizationProjects field value if set, zero value otherwise.
func (o *AppPermissions) GetOrganizationProjects() string {
	if o == nil || o.OrganizationProjects == nil {
		var ret string
		return ret
	}
	return *o.OrganizationProjects
}

// GetOrganizationProjectsOk returns a tuple with the OrganizationProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetOrganizationProjectsOk() (*string, bool) {
	if o == nil || o.OrganizationProjects == nil {
		return nil, false
	}
	return o.OrganizationProjects, true
}

// HasOrganizationProjects returns a boolean if a field has been set.
func (o *AppPermissions) HasOrganizationProjects() bool {
	if o != nil && o.OrganizationProjects != nil {
		return true
	}

	return false
}

// SetOrganizationProjects gets a reference to the given string and assigns it to the OrganizationProjects field.
func (o *AppPermissions) SetOrganizationProjects(v string) {
	o.OrganizationProjects = &v
}

// GetOrganizationPackages returns the OrganizationPackages field value if set, zero value otherwise.
func (o *AppPermissions) GetOrganizationPackages() string {
	if o == nil || o.OrganizationPackages == nil {
		var ret string
		return ret
	}
	return *o.OrganizationPackages
}

// GetOrganizationPackagesOk returns a tuple with the OrganizationPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetOrganizationPackagesOk() (*string, bool) {
	if o == nil || o.OrganizationPackages == nil {
		return nil, false
	}
	return o.OrganizationPackages, true
}

// HasOrganizationPackages returns a boolean if a field has been set.
func (o *AppPermissions) HasOrganizationPackages() bool {
	if o != nil && o.OrganizationPackages != nil {
		return true
	}

	return false
}

// SetOrganizationPackages gets a reference to the given string and assigns it to the OrganizationPackages field.
func (o *AppPermissions) SetOrganizationPackages(v string) {
	o.OrganizationPackages = &v
}

// GetOrganizationSecrets returns the OrganizationSecrets field value if set, zero value otherwise.
func (o *AppPermissions) GetOrganizationSecrets() string {
	if o == nil || o.OrganizationSecrets == nil {
		var ret string
		return ret
	}
	return *o.OrganizationSecrets
}

// GetOrganizationSecretsOk returns a tuple with the OrganizationSecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetOrganizationSecretsOk() (*string, bool) {
	if o == nil || o.OrganizationSecrets == nil {
		return nil, false
	}
	return o.OrganizationSecrets, true
}

// HasOrganizationSecrets returns a boolean if a field has been set.
func (o *AppPermissions) HasOrganizationSecrets() bool {
	if o != nil && o.OrganizationSecrets != nil {
		return true
	}

	return false
}

// SetOrganizationSecrets gets a reference to the given string and assigns it to the OrganizationSecrets field.
func (o *AppPermissions) SetOrganizationSecrets(v string) {
	o.OrganizationSecrets = &v
}

// GetOrganizationSelfHostedRunners returns the OrganizationSelfHostedRunners field value if set, zero value otherwise.
func (o *AppPermissions) GetOrganizationSelfHostedRunners() string {
	if o == nil || o.OrganizationSelfHostedRunners == nil {
		var ret string
		return ret
	}
	return *o.OrganizationSelfHostedRunners
}

// GetOrganizationSelfHostedRunnersOk returns a tuple with the OrganizationSelfHostedRunners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetOrganizationSelfHostedRunnersOk() (*string, bool) {
	if o == nil || o.OrganizationSelfHostedRunners == nil {
		return nil, false
	}
	return o.OrganizationSelfHostedRunners, true
}

// HasOrganizationSelfHostedRunners returns a boolean if a field has been set.
func (o *AppPermissions) HasOrganizationSelfHostedRunners() bool {
	if o != nil && o.OrganizationSelfHostedRunners != nil {
		return true
	}

	return false
}

// SetOrganizationSelfHostedRunners gets a reference to the given string and assigns it to the OrganizationSelfHostedRunners field.
func (o *AppPermissions) SetOrganizationSelfHostedRunners(v string) {
	o.OrganizationSelfHostedRunners = &v
}

// GetOrganizationUserBlocking returns the OrganizationUserBlocking field value if set, zero value otherwise.
func (o *AppPermissions) GetOrganizationUserBlocking() string {
	if o == nil || o.OrganizationUserBlocking == nil {
		var ret string
		return ret
	}
	return *o.OrganizationUserBlocking
}

// GetOrganizationUserBlockingOk returns a tuple with the OrganizationUserBlocking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetOrganizationUserBlockingOk() (*string, bool) {
	if o == nil || o.OrganizationUserBlocking == nil {
		return nil, false
	}
	return o.OrganizationUserBlocking, true
}

// HasOrganizationUserBlocking returns a boolean if a field has been set.
func (o *AppPermissions) HasOrganizationUserBlocking() bool {
	if o != nil && o.OrganizationUserBlocking != nil {
		return true
	}

	return false
}

// SetOrganizationUserBlocking gets a reference to the given string and assigns it to the OrganizationUserBlocking field.
func (o *AppPermissions) SetOrganizationUserBlocking(v string) {
	o.OrganizationUserBlocking = &v
}

// GetTeamDiscussions returns the TeamDiscussions field value if set, zero value otherwise.
func (o *AppPermissions) GetTeamDiscussions() string {
	if o == nil || o.TeamDiscussions == nil {
		var ret string
		return ret
	}
	return *o.TeamDiscussions
}

// GetTeamDiscussionsOk returns a tuple with the TeamDiscussions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPermissions) GetTeamDiscussionsOk() (*string, bool) {
	if o == nil || o.TeamDiscussions == nil {
		return nil, false
	}
	return o.TeamDiscussions, true
}

// HasTeamDiscussions returns a boolean if a field has been set.
func (o *AppPermissions) HasTeamDiscussions() bool {
	if o != nil && o.TeamDiscussions != nil {
		return true
	}

	return false
}

// SetTeamDiscussions gets a reference to the given string and assigns it to the TeamDiscussions field.
func (o *AppPermissions) SetTeamDiscussions(v string) {
	o.TeamDiscussions = &v
}

func (o AppPermissions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Administration != nil {
		toSerialize["administration"] = o.Administration
	}
	if o.Checks != nil {
		toSerialize["checks"] = o.Checks
	}
	if o.Contents != nil {
		toSerialize["contents"] = o.Contents
	}
	if o.Deployments != nil {
		toSerialize["deployments"] = o.Deployments
	}
	if o.Environments != nil {
		toSerialize["environments"] = o.Environments
	}
	if o.Issues != nil {
		toSerialize["issues"] = o.Issues
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Packages != nil {
		toSerialize["packages"] = o.Packages
	}
	if o.Pages != nil {
		toSerialize["pages"] = o.Pages
	}
	if o.PullRequests != nil {
		toSerialize["pull_requests"] = o.PullRequests
	}
	if o.RepositoryHooks != nil {
		toSerialize["repository_hooks"] = o.RepositoryHooks
	}
	if o.RepositoryProjects != nil {
		toSerialize["repository_projects"] = o.RepositoryProjects
	}
	if o.SecretScanningAlerts != nil {
		toSerialize["secret_scanning_alerts"] = o.SecretScanningAlerts
	}
	if o.Secrets != nil {
		toSerialize["secrets"] = o.Secrets
	}
	if o.SecurityEvents != nil {
		toSerialize["security_events"] = o.SecurityEvents
	}
	if o.SingleFile != nil {
		toSerialize["single_file"] = o.SingleFile
	}
	if o.Statuses != nil {
		toSerialize["statuses"] = o.Statuses
	}
	if o.VulnerabilityAlerts != nil {
		toSerialize["vulnerability_alerts"] = o.VulnerabilityAlerts
	}
	if o.Workflows != nil {
		toSerialize["workflows"] = o.Workflows
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if o.OrganizationAdministration != nil {
		toSerialize["organization_administration"] = o.OrganizationAdministration
	}
	if o.OrganizationHooks != nil {
		toSerialize["organization_hooks"] = o.OrganizationHooks
	}
	if o.OrganizationPlan != nil {
		toSerialize["organization_plan"] = o.OrganizationPlan
	}
	if o.OrganizationProjects != nil {
		toSerialize["organization_projects"] = o.OrganizationProjects
	}
	if o.OrganizationPackages != nil {
		toSerialize["organization_packages"] = o.OrganizationPackages
	}
	if o.OrganizationSecrets != nil {
		toSerialize["organization_secrets"] = o.OrganizationSecrets
	}
	if o.OrganizationSelfHostedRunners != nil {
		toSerialize["organization_self_hosted_runners"] = o.OrganizationSelfHostedRunners
	}
	if o.OrganizationUserBlocking != nil {
		toSerialize["organization_user_blocking"] = o.OrganizationUserBlocking
	}
	if o.TeamDiscussions != nil {
		toSerialize["team_discussions"] = o.TeamDiscussions
	}
	return json.Marshal(toSerialize)
}

type NullableAppPermissions struct {
	value *AppPermissions
	isSet bool
}

func (v NullableAppPermissions) Get() *AppPermissions {
	return v.value
}

func (v *NullableAppPermissions) Set(val *AppPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPermissions(val *AppPermissions) *NullableAppPermissions {
	return &NullableAppPermissions{value: val, isSet: true}
}

func (v NullableAppPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


