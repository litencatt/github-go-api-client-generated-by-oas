# OrgsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingEmail** | Pointer to **string** | Billing email address. This address is not publicized. | [optional] 
**Company** | Pointer to **string** | The company name. | [optional] 
**Email** | Pointer to **string** | The publicly visible email address. | [optional] 
**TwitterUsername** | Pointer to **string** | The Twitter username of the company. | [optional] 
**Location** | Pointer to **string** | The location. | [optional] 
**Name** | Pointer to **string** | The shorthand name of the company. | [optional] 
**Description** | Pointer to **string** | The description of the company. | [optional] 
**HasOrganizationProjects** | Pointer to **bool** | Whether an organization can use organization projects. | [optional] 
**HasRepositoryProjects** | Pointer to **bool** | Whether repositories that belong to the organization can use repository projects. | [optional] 
**DefaultRepositoryPermission** | Pointer to **string** | Default permission level members have for organization repositories. | [optional] [default to "read"]
**MembersCanCreateRepositories** | Pointer to **bool** | Whether of non-admin organization members can create repositories. **Note:** A parameter can override this parameter. See &#x60;members_allowed_repository_creation_type&#x60; in this table for details. | [optional] [default to true]
**MembersCanCreateInternalRepositories** | Pointer to **bool** | Whether organization members can create internal repositories, which are visible to all enterprise members. You can only allow members to create internal repositories if your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+. For more information, see \&quot;[Restricting repository creation in your organization](https://docs.github.com/github/setting-up-and-managing-organizations-and-teams/restricting-repository-creation-in-your-organization)\&quot; in the GitHub Help documentation. | [optional] 
**MembersCanCreatePrivateRepositories** | Pointer to **bool** | Whether organization members can create private repositories, which are visible to organization members with permission. For more information, see \&quot;[Restricting repository creation in your organization](https://docs.github.com/github/setting-up-and-managing-organizations-and-teams/restricting-repository-creation-in-your-organization)\&quot; in the GitHub Help documentation. | [optional] 
**MembersCanCreatePublicRepositories** | Pointer to **bool** | Whether organization members can create public repositories, which are visible to anyone. For more information, see \&quot;[Restricting repository creation in your organization](https://docs.github.com/github/setting-up-and-managing-organizations-and-teams/restricting-repository-creation-in-your-organization)\&quot; in the GitHub Help documentation. | [optional] 
**MembersAllowedRepositoryCreationType** | Pointer to **string** | Specifies which types of repositories non-admin organization members can create. &#x60;private&#x60; is only available to repositories that are part of an organization on GitHub Enterprise Cloud.  **Note:** This parameter is deprecated and will be removed in the future. Its return value ignores internal repositories. Using this parameter overrides values set in &#x60;members_can_create_repositories&#x60;. See the parameter deprecation notice in the operation description for details. | [optional] 
**MembersCanCreatePages** | Pointer to **bool** | Whether organization members can create GitHub Pages sites. Existing published sites will not be impacted. | [optional] [default to true]
**MembersCanCreatePublicPages** | Pointer to **bool** | Whether organization members can create public GitHub Pages sites. Existing published sites will not be impacted. | [optional] [default to true]
**MembersCanCreatePrivatePages** | Pointer to **bool** | Whether organization members can create private GitHub Pages sites. Existing published sites will not be impacted. | [optional] [default to true]
**MembersCanForkPrivateRepositories** | Pointer to **bool** | Whether organization members can fork private organization repositories. | [optional] [default to false]
**Blog** | Pointer to **string** |  | [optional] 

## Methods

### NewOrgsUpdateRequest

`func NewOrgsUpdateRequest() *OrgsUpdateRequest`

NewOrgsUpdateRequest instantiates a new OrgsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsUpdateRequestWithDefaults

`func NewOrgsUpdateRequestWithDefaults() *OrgsUpdateRequest`

NewOrgsUpdateRequestWithDefaults instantiates a new OrgsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingEmail

`func (o *OrgsUpdateRequest) GetBillingEmail() string`

GetBillingEmail returns the BillingEmail field if non-nil, zero value otherwise.

### GetBillingEmailOk

`func (o *OrgsUpdateRequest) GetBillingEmailOk() (*string, bool)`

GetBillingEmailOk returns a tuple with the BillingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmail

`func (o *OrgsUpdateRequest) SetBillingEmail(v string)`

SetBillingEmail sets BillingEmail field to given value.

### HasBillingEmail

`func (o *OrgsUpdateRequest) HasBillingEmail() bool`

HasBillingEmail returns a boolean if a field has been set.

### GetCompany

`func (o *OrgsUpdateRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *OrgsUpdateRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *OrgsUpdateRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *OrgsUpdateRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetEmail

`func (o *OrgsUpdateRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrgsUpdateRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrgsUpdateRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrgsUpdateRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetTwitterUsername

`func (o *OrgsUpdateRequest) GetTwitterUsername() string`

GetTwitterUsername returns the TwitterUsername field if non-nil, zero value otherwise.

### GetTwitterUsernameOk

`func (o *OrgsUpdateRequest) GetTwitterUsernameOk() (*string, bool)`

GetTwitterUsernameOk returns a tuple with the TwitterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUsername

`func (o *OrgsUpdateRequest) SetTwitterUsername(v string)`

SetTwitterUsername sets TwitterUsername field to given value.

### HasTwitterUsername

`func (o *OrgsUpdateRequest) HasTwitterUsername() bool`

HasTwitterUsername returns a boolean if a field has been set.

### GetLocation

`func (o *OrgsUpdateRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OrgsUpdateRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OrgsUpdateRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OrgsUpdateRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *OrgsUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgsUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgsUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrgsUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *OrgsUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrgsUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrgsUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrgsUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHasOrganizationProjects

`func (o *OrgsUpdateRequest) GetHasOrganizationProjects() bool`

GetHasOrganizationProjects returns the HasOrganizationProjects field if non-nil, zero value otherwise.

### GetHasOrganizationProjectsOk

`func (o *OrgsUpdateRequest) GetHasOrganizationProjectsOk() (*bool, bool)`

GetHasOrganizationProjectsOk returns a tuple with the HasOrganizationProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOrganizationProjects

`func (o *OrgsUpdateRequest) SetHasOrganizationProjects(v bool)`

SetHasOrganizationProjects sets HasOrganizationProjects field to given value.

### HasHasOrganizationProjects

`func (o *OrgsUpdateRequest) HasHasOrganizationProjects() bool`

HasHasOrganizationProjects returns a boolean if a field has been set.

### GetHasRepositoryProjects

`func (o *OrgsUpdateRequest) GetHasRepositoryProjects() bool`

GetHasRepositoryProjects returns the HasRepositoryProjects field if non-nil, zero value otherwise.

### GetHasRepositoryProjectsOk

`func (o *OrgsUpdateRequest) GetHasRepositoryProjectsOk() (*bool, bool)`

GetHasRepositoryProjectsOk returns a tuple with the HasRepositoryProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRepositoryProjects

`func (o *OrgsUpdateRequest) SetHasRepositoryProjects(v bool)`

SetHasRepositoryProjects sets HasRepositoryProjects field to given value.

### HasHasRepositoryProjects

`func (o *OrgsUpdateRequest) HasHasRepositoryProjects() bool`

HasHasRepositoryProjects returns a boolean if a field has been set.

### GetDefaultRepositoryPermission

`func (o *OrgsUpdateRequest) GetDefaultRepositoryPermission() string`

GetDefaultRepositoryPermission returns the DefaultRepositoryPermission field if non-nil, zero value otherwise.

### GetDefaultRepositoryPermissionOk

`func (o *OrgsUpdateRequest) GetDefaultRepositoryPermissionOk() (*string, bool)`

GetDefaultRepositoryPermissionOk returns a tuple with the DefaultRepositoryPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRepositoryPermission

`func (o *OrgsUpdateRequest) SetDefaultRepositoryPermission(v string)`

SetDefaultRepositoryPermission sets DefaultRepositoryPermission field to given value.

### HasDefaultRepositoryPermission

`func (o *OrgsUpdateRequest) HasDefaultRepositoryPermission() bool`

HasDefaultRepositoryPermission returns a boolean if a field has been set.

### GetMembersCanCreateRepositories

`func (o *OrgsUpdateRequest) GetMembersCanCreateRepositories() bool`

GetMembersCanCreateRepositories returns the MembersCanCreateRepositories field if non-nil, zero value otherwise.

### GetMembersCanCreateRepositoriesOk

`func (o *OrgsUpdateRequest) GetMembersCanCreateRepositoriesOk() (*bool, bool)`

GetMembersCanCreateRepositoriesOk returns a tuple with the MembersCanCreateRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreateRepositories

`func (o *OrgsUpdateRequest) SetMembersCanCreateRepositories(v bool)`

SetMembersCanCreateRepositories sets MembersCanCreateRepositories field to given value.

### HasMembersCanCreateRepositories

`func (o *OrgsUpdateRequest) HasMembersCanCreateRepositories() bool`

HasMembersCanCreateRepositories returns a boolean if a field has been set.

### GetMembersCanCreateInternalRepositories

`func (o *OrgsUpdateRequest) GetMembersCanCreateInternalRepositories() bool`

GetMembersCanCreateInternalRepositories returns the MembersCanCreateInternalRepositories field if non-nil, zero value otherwise.

### GetMembersCanCreateInternalRepositoriesOk

`func (o *OrgsUpdateRequest) GetMembersCanCreateInternalRepositoriesOk() (*bool, bool)`

GetMembersCanCreateInternalRepositoriesOk returns a tuple with the MembersCanCreateInternalRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreateInternalRepositories

`func (o *OrgsUpdateRequest) SetMembersCanCreateInternalRepositories(v bool)`

SetMembersCanCreateInternalRepositories sets MembersCanCreateInternalRepositories field to given value.

### HasMembersCanCreateInternalRepositories

`func (o *OrgsUpdateRequest) HasMembersCanCreateInternalRepositories() bool`

HasMembersCanCreateInternalRepositories returns a boolean if a field has been set.

### GetMembersCanCreatePrivateRepositories

`func (o *OrgsUpdateRequest) GetMembersCanCreatePrivateRepositories() bool`

GetMembersCanCreatePrivateRepositories returns the MembersCanCreatePrivateRepositories field if non-nil, zero value otherwise.

### GetMembersCanCreatePrivateRepositoriesOk

`func (o *OrgsUpdateRequest) GetMembersCanCreatePrivateRepositoriesOk() (*bool, bool)`

GetMembersCanCreatePrivateRepositoriesOk returns a tuple with the MembersCanCreatePrivateRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreatePrivateRepositories

`func (o *OrgsUpdateRequest) SetMembersCanCreatePrivateRepositories(v bool)`

SetMembersCanCreatePrivateRepositories sets MembersCanCreatePrivateRepositories field to given value.

### HasMembersCanCreatePrivateRepositories

`func (o *OrgsUpdateRequest) HasMembersCanCreatePrivateRepositories() bool`

HasMembersCanCreatePrivateRepositories returns a boolean if a field has been set.

### GetMembersCanCreatePublicRepositories

`func (o *OrgsUpdateRequest) GetMembersCanCreatePublicRepositories() bool`

GetMembersCanCreatePublicRepositories returns the MembersCanCreatePublicRepositories field if non-nil, zero value otherwise.

### GetMembersCanCreatePublicRepositoriesOk

`func (o *OrgsUpdateRequest) GetMembersCanCreatePublicRepositoriesOk() (*bool, bool)`

GetMembersCanCreatePublicRepositoriesOk returns a tuple with the MembersCanCreatePublicRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreatePublicRepositories

`func (o *OrgsUpdateRequest) SetMembersCanCreatePublicRepositories(v bool)`

SetMembersCanCreatePublicRepositories sets MembersCanCreatePublicRepositories field to given value.

### HasMembersCanCreatePublicRepositories

`func (o *OrgsUpdateRequest) HasMembersCanCreatePublicRepositories() bool`

HasMembersCanCreatePublicRepositories returns a boolean if a field has been set.

### GetMembersAllowedRepositoryCreationType

`func (o *OrgsUpdateRequest) GetMembersAllowedRepositoryCreationType() string`

GetMembersAllowedRepositoryCreationType returns the MembersAllowedRepositoryCreationType field if non-nil, zero value otherwise.

### GetMembersAllowedRepositoryCreationTypeOk

`func (o *OrgsUpdateRequest) GetMembersAllowedRepositoryCreationTypeOk() (*string, bool)`

GetMembersAllowedRepositoryCreationTypeOk returns a tuple with the MembersAllowedRepositoryCreationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersAllowedRepositoryCreationType

`func (o *OrgsUpdateRequest) SetMembersAllowedRepositoryCreationType(v string)`

SetMembersAllowedRepositoryCreationType sets MembersAllowedRepositoryCreationType field to given value.

### HasMembersAllowedRepositoryCreationType

`func (o *OrgsUpdateRequest) HasMembersAllowedRepositoryCreationType() bool`

HasMembersAllowedRepositoryCreationType returns a boolean if a field has been set.

### GetMembersCanCreatePages

`func (o *OrgsUpdateRequest) GetMembersCanCreatePages() bool`

GetMembersCanCreatePages returns the MembersCanCreatePages field if non-nil, zero value otherwise.

### GetMembersCanCreatePagesOk

`func (o *OrgsUpdateRequest) GetMembersCanCreatePagesOk() (*bool, bool)`

GetMembersCanCreatePagesOk returns a tuple with the MembersCanCreatePages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreatePages

`func (o *OrgsUpdateRequest) SetMembersCanCreatePages(v bool)`

SetMembersCanCreatePages sets MembersCanCreatePages field to given value.

### HasMembersCanCreatePages

`func (o *OrgsUpdateRequest) HasMembersCanCreatePages() bool`

HasMembersCanCreatePages returns a boolean if a field has been set.

### GetMembersCanCreatePublicPages

`func (o *OrgsUpdateRequest) GetMembersCanCreatePublicPages() bool`

GetMembersCanCreatePublicPages returns the MembersCanCreatePublicPages field if non-nil, zero value otherwise.

### GetMembersCanCreatePublicPagesOk

`func (o *OrgsUpdateRequest) GetMembersCanCreatePublicPagesOk() (*bool, bool)`

GetMembersCanCreatePublicPagesOk returns a tuple with the MembersCanCreatePublicPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreatePublicPages

`func (o *OrgsUpdateRequest) SetMembersCanCreatePublicPages(v bool)`

SetMembersCanCreatePublicPages sets MembersCanCreatePublicPages field to given value.

### HasMembersCanCreatePublicPages

`func (o *OrgsUpdateRequest) HasMembersCanCreatePublicPages() bool`

HasMembersCanCreatePublicPages returns a boolean if a field has been set.

### GetMembersCanCreatePrivatePages

`func (o *OrgsUpdateRequest) GetMembersCanCreatePrivatePages() bool`

GetMembersCanCreatePrivatePages returns the MembersCanCreatePrivatePages field if non-nil, zero value otherwise.

### GetMembersCanCreatePrivatePagesOk

`func (o *OrgsUpdateRequest) GetMembersCanCreatePrivatePagesOk() (*bool, bool)`

GetMembersCanCreatePrivatePagesOk returns a tuple with the MembersCanCreatePrivatePages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanCreatePrivatePages

`func (o *OrgsUpdateRequest) SetMembersCanCreatePrivatePages(v bool)`

SetMembersCanCreatePrivatePages sets MembersCanCreatePrivatePages field to given value.

### HasMembersCanCreatePrivatePages

`func (o *OrgsUpdateRequest) HasMembersCanCreatePrivatePages() bool`

HasMembersCanCreatePrivatePages returns a boolean if a field has been set.

### GetMembersCanForkPrivateRepositories

`func (o *OrgsUpdateRequest) GetMembersCanForkPrivateRepositories() bool`

GetMembersCanForkPrivateRepositories returns the MembersCanForkPrivateRepositories field if non-nil, zero value otherwise.

### GetMembersCanForkPrivateRepositoriesOk

`func (o *OrgsUpdateRequest) GetMembersCanForkPrivateRepositoriesOk() (*bool, bool)`

GetMembersCanForkPrivateRepositoriesOk returns a tuple with the MembersCanForkPrivateRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCanForkPrivateRepositories

`func (o *OrgsUpdateRequest) SetMembersCanForkPrivateRepositories(v bool)`

SetMembersCanForkPrivateRepositories sets MembersCanForkPrivateRepositories field to given value.

### HasMembersCanForkPrivateRepositories

`func (o *OrgsUpdateRequest) HasMembersCanForkPrivateRepositories() bool`

HasMembersCanForkPrivateRepositories returns a boolean if a field has been set.

### GetBlog

`func (o *OrgsUpdateRequest) GetBlog() string`

GetBlog returns the Blog field if non-nil, zero value otherwise.

### GetBlogOk

`func (o *OrgsUpdateRequest) GetBlogOk() (*string, bool)`

GetBlogOk returns a tuple with the Blog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlog

`func (o *OrgsUpdateRequest) SetBlog(v string)`

SetBlog sets Blog field to given value.

### HasBlog

`func (o *OrgsUpdateRequest) HasBlog() bool`

HasBlog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


