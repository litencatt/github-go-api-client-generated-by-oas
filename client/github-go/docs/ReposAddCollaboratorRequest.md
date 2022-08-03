# ReposAddCollaboratorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | Pointer to **string** | The permission to grant the collaborator. **Only valid on organization-owned repositories.** In addition to the enumerated values, you can also specify a custom repository role name, if the owning organization has defined any. | [optional] [default to "push"]

## Methods

### NewReposAddCollaboratorRequest

`func NewReposAddCollaboratorRequest() *ReposAddCollaboratorRequest`

NewReposAddCollaboratorRequest instantiates a new ReposAddCollaboratorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposAddCollaboratorRequestWithDefaults

`func NewReposAddCollaboratorRequestWithDefaults() *ReposAddCollaboratorRequest`

NewReposAddCollaboratorRequestWithDefaults instantiates a new ReposAddCollaboratorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *ReposAddCollaboratorRequest) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ReposAddCollaboratorRequest) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ReposAddCollaboratorRequest) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *ReposAddCollaboratorRequest) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


