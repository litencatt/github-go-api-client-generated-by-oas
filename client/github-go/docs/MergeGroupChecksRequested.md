# MergeGroupChecksRequested

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Installation** | Pointer to [**SimpleInstallation**](SimpleInstallation.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationSimple**](OrganizationSimple.md) |  | [optional] 
**Repository** | Pointer to [**Repository**](Repository.md) |  | [optional] 
**Sender** | Pointer to [**SimpleUser**](SimpleUser.md) |  | [optional] 
**MergeGroup** | [**MergeGroupChecksRequestedMergeGroup**](MergeGroupChecksRequestedMergeGroup.md) |  | 

## Methods

### NewMergeGroupChecksRequested

`func NewMergeGroupChecksRequested(action string, mergeGroup MergeGroupChecksRequestedMergeGroup, ) *MergeGroupChecksRequested`

NewMergeGroupChecksRequested instantiates a new MergeGroupChecksRequested object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeGroupChecksRequestedWithDefaults

`func NewMergeGroupChecksRequestedWithDefaults() *MergeGroupChecksRequested`

NewMergeGroupChecksRequestedWithDefaults instantiates a new MergeGroupChecksRequested object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *MergeGroupChecksRequested) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *MergeGroupChecksRequested) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *MergeGroupChecksRequested) SetAction(v string)`

SetAction sets Action field to given value.


### GetInstallation

`func (o *MergeGroupChecksRequested) GetInstallation() SimpleInstallation`

GetInstallation returns the Installation field if non-nil, zero value otherwise.

### GetInstallationOk

`func (o *MergeGroupChecksRequested) GetInstallationOk() (*SimpleInstallation, bool)`

GetInstallationOk returns a tuple with the Installation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallation

`func (o *MergeGroupChecksRequested) SetInstallation(v SimpleInstallation)`

SetInstallation sets Installation field to given value.

### HasInstallation

`func (o *MergeGroupChecksRequested) HasInstallation() bool`

HasInstallation returns a boolean if a field has been set.

### GetOrganization

`func (o *MergeGroupChecksRequested) GetOrganization() OrganizationSimple`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MergeGroupChecksRequested) GetOrganizationOk() (*OrganizationSimple, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MergeGroupChecksRequested) SetOrganization(v OrganizationSimple)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MergeGroupChecksRequested) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRepository

`func (o *MergeGroupChecksRequested) GetRepository() Repository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *MergeGroupChecksRequested) GetRepositoryOk() (*Repository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *MergeGroupChecksRequested) SetRepository(v Repository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *MergeGroupChecksRequested) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetSender

`func (o *MergeGroupChecksRequested) GetSender() SimpleUser`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *MergeGroupChecksRequested) GetSenderOk() (*SimpleUser, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *MergeGroupChecksRequested) SetSender(v SimpleUser)`

SetSender sets Sender field to given value.

### HasSender

`func (o *MergeGroupChecksRequested) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetMergeGroup

`func (o *MergeGroupChecksRequested) GetMergeGroup() MergeGroupChecksRequestedMergeGroup`

GetMergeGroup returns the MergeGroup field if non-nil, zero value otherwise.

### GetMergeGroupOk

`func (o *MergeGroupChecksRequested) GetMergeGroupOk() (*MergeGroupChecksRequestedMergeGroup, bool)`

GetMergeGroupOk returns a tuple with the MergeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeGroup

`func (o *MergeGroupChecksRequested) SetMergeGroup(v MergeGroupChecksRequestedMergeGroup)`

SetMergeGroup sets MergeGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


