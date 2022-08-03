# IssuesAddLabelsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to [**[]IssuesSetLabelsRequestOneOf1LabelsInner**](IssuesSetLabelsRequestOneOf1LabelsInner.md) |  | [optional] 

## Methods

### NewIssuesAddLabelsRequest

`func NewIssuesAddLabelsRequest() *IssuesAddLabelsRequest`

NewIssuesAddLabelsRequest instantiates a new IssuesAddLabelsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuesAddLabelsRequestWithDefaults

`func NewIssuesAddLabelsRequestWithDefaults() *IssuesAddLabelsRequest`

NewIssuesAddLabelsRequestWithDefaults instantiates a new IssuesAddLabelsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *IssuesAddLabelsRequest) GetLabels() []IssuesSetLabelsRequestOneOf1LabelsInner`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *IssuesAddLabelsRequest) GetLabelsOk() (*[]IssuesSetLabelsRequestOneOf1LabelsInner, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *IssuesAddLabelsRequest) SetLabels(v []IssuesSetLabelsRequestOneOf1LabelsInner)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *IssuesAddLabelsRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


