# DeploymentBranchPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectedBranches** | **bool** | Whether only branches with branch protection rules can deploy to this environment. If &#x60;protected_branches&#x60; is &#x60;true&#x60;, &#x60;custom_branch_policies&#x60; must be &#x60;false&#x60;; if &#x60;protected_branches&#x60; is &#x60;false&#x60;, &#x60;custom_branch_policies&#x60; must be &#x60;true&#x60;. | 
**CustomBranchPolicies** | **bool** | Whether only branches that match the specified name patterns can deploy to this environment.  If &#x60;custom_branch_policies&#x60; is &#x60;true&#x60;, &#x60;protected_branches&#x60; must be &#x60;false&#x60;; if &#x60;custom_branch_policies&#x60; is &#x60;false&#x60;, &#x60;protected_branches&#x60; must be &#x60;true&#x60;. | 

## Methods

### NewDeploymentBranchPolicy

`func NewDeploymentBranchPolicy(protectedBranches bool, customBranchPolicies bool, ) *DeploymentBranchPolicy`

NewDeploymentBranchPolicy instantiates a new DeploymentBranchPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentBranchPolicyWithDefaults

`func NewDeploymentBranchPolicyWithDefaults() *DeploymentBranchPolicy`

NewDeploymentBranchPolicyWithDefaults instantiates a new DeploymentBranchPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectedBranches

`func (o *DeploymentBranchPolicy) GetProtectedBranches() bool`

GetProtectedBranches returns the ProtectedBranches field if non-nil, zero value otherwise.

### GetProtectedBranchesOk

`func (o *DeploymentBranchPolicy) GetProtectedBranchesOk() (*bool, bool)`

GetProtectedBranchesOk returns a tuple with the ProtectedBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedBranches

`func (o *DeploymentBranchPolicy) SetProtectedBranches(v bool)`

SetProtectedBranches sets ProtectedBranches field to given value.


### GetCustomBranchPolicies

`func (o *DeploymentBranchPolicy) GetCustomBranchPolicies() bool`

GetCustomBranchPolicies returns the CustomBranchPolicies field if non-nil, zero value otherwise.

### GetCustomBranchPoliciesOk

`func (o *DeploymentBranchPolicy) GetCustomBranchPoliciesOk() (*bool, bool)`

GetCustomBranchPoliciesOk returns a tuple with the CustomBranchPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBranchPolicies

`func (o *DeploymentBranchPolicy) SetCustomBranchPolicies(v bool)`

SetCustomBranchPolicies sets CustomBranchPolicies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


