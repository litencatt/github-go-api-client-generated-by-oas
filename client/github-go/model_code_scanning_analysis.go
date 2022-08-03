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

// CodeScanningAnalysis struct for CodeScanningAnalysis
type CodeScanningAnalysis struct {
	// The full Git reference, formatted as `refs/heads/<branch name>`, `refs/pull/<number>/merge`, or `refs/pull/<number>/head`.
	Ref string `json:"ref"`
	// The SHA of the commit to which the analysis you are uploading relates.
	CommitSha string `json:"commit_sha"`
	// Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
	AnalysisKey string `json:"analysis_key"`
	// Identifies the variable values associated with the environment in which this analysis was performed.
	Environment string `json:"environment"`
	// Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
	Category *string `json:"category,omitempty"`
	Error string `json:"error"`
	// The time that the analysis was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	CreatedAt time.Time `json:"created_at"`
	// The total number of results in the analysis.
	ResultsCount int32 `json:"results_count"`
	// The total number of rules used in the analysis.
	RulesCount int32 `json:"rules_count"`
	// Unique identifier for this analysis.
	Id int32 `json:"id"`
	// The REST API URL of the analysis resource.
	Url string `json:"url"`
	// An identifier for the upload.
	SarifId string `json:"sarif_id"`
	Tool CodeScanningAnalysisTool `json:"tool"`
	Deletable bool `json:"deletable"`
	// Warning generated when processing the analysis
	Warning string `json:"warning"`
}

// NewCodeScanningAnalysis instantiates a new CodeScanningAnalysis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodeScanningAnalysis(ref string, commitSha string, analysisKey string, environment string, error_ string, createdAt time.Time, resultsCount int32, rulesCount int32, id int32, url string, sarifId string, tool CodeScanningAnalysisTool, deletable bool, warning string) *CodeScanningAnalysis {
	this := CodeScanningAnalysis{}
	this.Ref = ref
	this.CommitSha = commitSha
	this.AnalysisKey = analysisKey
	this.Environment = environment
	this.Error = error_
	this.CreatedAt = createdAt
	this.ResultsCount = resultsCount
	this.RulesCount = rulesCount
	this.Id = id
	this.Url = url
	this.SarifId = sarifId
	this.Tool = tool
	this.Deletable = deletable
	this.Warning = warning
	return &this
}

// NewCodeScanningAnalysisWithDefaults instantiates a new CodeScanningAnalysis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodeScanningAnalysisWithDefaults() *CodeScanningAnalysis {
	this := CodeScanningAnalysis{}
	return &this
}

// GetRef returns the Ref field value
func (o *CodeScanningAnalysis) GetRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ref
}

// GetRefOk returns a tuple with the Ref field value
// and a boolean to check if the value has been set.
func (o *CodeScanningAnalysis) GetRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ref, true
}

// SetRef sets field value
func (o *CodeScanningAnalysis) SetRef(v string) {
	o.Ref = v
}

// GetCommitSha returns the CommitSha field value
func (o *CodeScanningAnalysis) GetCommitSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitSha
}

// GetCommitShaOk returns a tuple with the CommitSha field value
// and a boolean to check if the value has been set.
func (o *CodeScanningAnalysis) GetCommitShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitSha, true
}

// SetCommitSha sets field value
func (o *CodeScanningAnalysis) SetCommitSha(v string) {
	o.CommitSha = v
}

// GetAnalysisKey returns the AnalysisKey field value
func (o *CodeScanningAnalysis) GetAnalysisKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnalysisKey
}

// GetAnalysisKeyOk returns a tuple with the AnalysisKey field value
// and a boolean to check if the value has been set.
func (o *CodeScanningAnalysis) GetAnalysisKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnalysisKey, true
}

// SetAnalysisKey sets field value
func (o *CodeScanningAnalysis) SetAnalysisKey(v string) {
	o.AnalysisKey = v
}

// GetEnvironment returns the Environment field value
func (o *CodeScanningAnalysis) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *CodeScanningAnalysis) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *CodeScanningAnalysis) SetEnvironment(v string) {
	o.Environment = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CodeScanningAnalysis) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeScanningAnalysis) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CodeScanningAnalysis) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *CodeScanningAnalysis) SetCategory(v string) {
	o.Category = &v
}

// GetError returns the Error field value
func (o *CodeScanningAnalysis) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *CodeScanningAnalysis) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *CodeScanningAnalysis) SetError(v string) {
	o.Error = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CodeScanningAnalysis) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CodeScanningAnalysis) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CodeScanningAnalysis) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetResultsCount returns the ResultsCount field value
func (o *CodeScanningAnalysis) GetResultsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResultsCount
}

// GetResultsCountOk returns a tuple with the ResultsCount field value
// and a boolean to check if the value has been set.
func (o *CodeScanningAnalysis) GetResultsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultsCount, true
}

// SetResultsCount sets field value
func (o *CodeScanningAnalysis) SetResultsCount(v int32) {
	o.ResultsCount = v
}

// GetRulesCount returns the RulesCount field value
func (o *CodeScanningAnalysis) GetRulesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RulesCount
}

// GetRulesCountOk returns a tuple with the RulesCount field value
// and a boolean to check if the value has been set.
func (o *CodeScanningAnalysis) GetRulesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RulesCount, true
}

// SetRulesCount sets field value
func (o *CodeScanningAnalysis) SetRulesCount(v int32) {
	o.RulesCount = v
}

// GetId returns the Id field value
func (o *CodeScanningAnalysis) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CodeScanningAnalysis) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CodeScanningAnalysis) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *CodeScanningAnalysis) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CodeScanningAnalysis) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CodeScanningAnalysis) SetUrl(v string) {
	o.Url = v
}

// GetSarifId returns the SarifId field value
func (o *CodeScanningAnalysis) GetSarifId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SarifId
}

// GetSarifIdOk returns a tuple with the SarifId field value
// and a boolean to check if the value has been set.
func (o *CodeScanningAnalysis) GetSarifIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SarifId, true
}

// SetSarifId sets field value
func (o *CodeScanningAnalysis) SetSarifId(v string) {
	o.SarifId = v
}

// GetTool returns the Tool field value
func (o *CodeScanningAnalysis) GetTool() CodeScanningAnalysisTool {
	if o == nil {
		var ret CodeScanningAnalysisTool
		return ret
	}

	return o.Tool
}

// GetToolOk returns a tuple with the Tool field value
// and a boolean to check if the value has been set.
func (o *CodeScanningAnalysis) GetToolOk() (*CodeScanningAnalysisTool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tool, true
}

// SetTool sets field value
func (o *CodeScanningAnalysis) SetTool(v CodeScanningAnalysisTool) {
	o.Tool = v
}

// GetDeletable returns the Deletable field value
func (o *CodeScanningAnalysis) GetDeletable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deletable
}

// GetDeletableOk returns a tuple with the Deletable field value
// and a boolean to check if the value has been set.
func (o *CodeScanningAnalysis) GetDeletableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deletable, true
}

// SetDeletable sets field value
func (o *CodeScanningAnalysis) SetDeletable(v bool) {
	o.Deletable = v
}

// GetWarning returns the Warning field value
func (o *CodeScanningAnalysis) GetWarning() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Warning
}

// GetWarningOk returns a tuple with the Warning field value
// and a boolean to check if the value has been set.
func (o *CodeScanningAnalysis) GetWarningOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Warning, true
}

// SetWarning sets field value
func (o *CodeScanningAnalysis) SetWarning(v string) {
	o.Warning = v
}

func (o CodeScanningAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ref"] = o.Ref
	}
	if true {
		toSerialize["commit_sha"] = o.CommitSha
	}
	if true {
		toSerialize["analysis_key"] = o.AnalysisKey
	}
	if true {
		toSerialize["environment"] = o.Environment
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["error"] = o.Error
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["results_count"] = o.ResultsCount
	}
	if true {
		toSerialize["rules_count"] = o.RulesCount
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["sarif_id"] = o.SarifId
	}
	if true {
		toSerialize["tool"] = o.Tool
	}
	if true {
		toSerialize["deletable"] = o.Deletable
	}
	if true {
		toSerialize["warning"] = o.Warning
	}
	return json.Marshal(toSerialize)
}

type NullableCodeScanningAnalysis struct {
	value *CodeScanningAnalysis
	isSet bool
}

func (v NullableCodeScanningAnalysis) Get() *CodeScanningAnalysis {
	return v.value
}

func (v *NullableCodeScanningAnalysis) Set(val *CodeScanningAnalysis) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeScanningAnalysis) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeScanningAnalysis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeScanningAnalysis(val *CodeScanningAnalysis) *NullableCodeScanningAnalysis {
	return &NullableCodeScanningAnalysis{value: val, isSet: true}
}

func (v NullableCodeScanningAnalysis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeScanningAnalysis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

