# UsersDeleteEmailForAuthenticatedUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | **[]string** | Email addresses associated with the GitHub user account. | 

## Methods

### NewUsersDeleteEmailForAuthenticatedUserRequest

`func NewUsersDeleteEmailForAuthenticatedUserRequest(emails []string, ) *UsersDeleteEmailForAuthenticatedUserRequest`

NewUsersDeleteEmailForAuthenticatedUserRequest instantiates a new UsersDeleteEmailForAuthenticatedUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersDeleteEmailForAuthenticatedUserRequestWithDefaults

`func NewUsersDeleteEmailForAuthenticatedUserRequestWithDefaults() *UsersDeleteEmailForAuthenticatedUserRequest`

NewUsersDeleteEmailForAuthenticatedUserRequestWithDefaults instantiates a new UsersDeleteEmailForAuthenticatedUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *UsersDeleteEmailForAuthenticatedUserRequest) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *UsersDeleteEmailForAuthenticatedUserRequest) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *UsersDeleteEmailForAuthenticatedUserRequest) SetEmails(v []string)`

SetEmails sets Emails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


