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

// GitCommit Low-level Git commit operations within a repository
type GitCommit struct {
	// SHA for the commit
	Sha string `json:"sha"`
	NodeId string `json:"node_id"`
	Url string `json:"url"`
	Author GitCommitAuthor `json:"author"`
	Committer GitCommitAuthor `json:"committer"`
	// Message describing the purpose of the commit
	Message string `json:"message"`
	Tree GitCommitTree `json:"tree"`
	Parents []GitCommitParentsInner `json:"parents"`
	Verification GitCommitVerification `json:"verification"`
	HtmlUrl string `json:"html_url"`
}

// NewGitCommit instantiates a new GitCommit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitCommit(sha string, nodeId string, url string, author GitCommitAuthor, committer GitCommitAuthor, message string, tree GitCommitTree, parents []GitCommitParentsInner, verification GitCommitVerification, htmlUrl string) *GitCommit {
	this := GitCommit{}
	this.Sha = sha
	this.NodeId = nodeId
	this.Url = url
	this.Author = author
	this.Committer = committer
	this.Message = message
	this.Tree = tree
	this.Parents = parents
	this.Verification = verification
	this.HtmlUrl = htmlUrl
	return &this
}

// NewGitCommitWithDefaults instantiates a new GitCommit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitCommitWithDefaults() *GitCommit {
	this := GitCommit{}
	return &this
}

// GetSha returns the Sha field value
func (o *GitCommit) GetSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha
}

// GetShaOk returns a tuple with the Sha field value
// and a boolean to check if the value has been set.
func (o *GitCommit) GetShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha, true
}

// SetSha sets field value
func (o *GitCommit) SetSha(v string) {
	o.Sha = v
}

// GetNodeId returns the NodeId field value
func (o *GitCommit) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *GitCommit) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *GitCommit) SetNodeId(v string) {
	o.NodeId = v
}

// GetUrl returns the Url field value
func (o *GitCommit) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *GitCommit) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *GitCommit) SetUrl(v string) {
	o.Url = v
}

// GetAuthor returns the Author field value
func (o *GitCommit) GetAuthor() GitCommitAuthor {
	if o == nil {
		var ret GitCommitAuthor
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *GitCommit) GetAuthorOk() (*GitCommitAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *GitCommit) SetAuthor(v GitCommitAuthor) {
	o.Author = v
}

// GetCommitter returns the Committer field value
func (o *GitCommit) GetCommitter() GitCommitAuthor {
	if o == nil {
		var ret GitCommitAuthor
		return ret
	}

	return o.Committer
}

// GetCommitterOk returns a tuple with the Committer field value
// and a boolean to check if the value has been set.
func (o *GitCommit) GetCommitterOk() (*GitCommitAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Committer, true
}

// SetCommitter sets field value
func (o *GitCommit) SetCommitter(v GitCommitAuthor) {
	o.Committer = v
}

// GetMessage returns the Message field value
func (o *GitCommit) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *GitCommit) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *GitCommit) SetMessage(v string) {
	o.Message = v
}

// GetTree returns the Tree field value
func (o *GitCommit) GetTree() GitCommitTree {
	if o == nil {
		var ret GitCommitTree
		return ret
	}

	return o.Tree
}

// GetTreeOk returns a tuple with the Tree field value
// and a boolean to check if the value has been set.
func (o *GitCommit) GetTreeOk() (*GitCommitTree, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tree, true
}

// SetTree sets field value
func (o *GitCommit) SetTree(v GitCommitTree) {
	o.Tree = v
}

// GetParents returns the Parents field value
func (o *GitCommit) GetParents() []GitCommitParentsInner {
	if o == nil {
		var ret []GitCommitParentsInner
		return ret
	}

	return o.Parents
}

// GetParentsOk returns a tuple with the Parents field value
// and a boolean to check if the value has been set.
func (o *GitCommit) GetParentsOk() ([]GitCommitParentsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parents, true
}

// SetParents sets field value
func (o *GitCommit) SetParents(v []GitCommitParentsInner) {
	o.Parents = v
}

// GetVerification returns the Verification field value
func (o *GitCommit) GetVerification() GitCommitVerification {
	if o == nil {
		var ret GitCommitVerification
		return ret
	}

	return o.Verification
}

// GetVerificationOk returns a tuple with the Verification field value
// and a boolean to check if the value has been set.
func (o *GitCommit) GetVerificationOk() (*GitCommitVerification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verification, true
}

// SetVerification sets field value
func (o *GitCommit) SetVerification(v GitCommitVerification) {
	o.Verification = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *GitCommit) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *GitCommit) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *GitCommit) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

func (o GitCommit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sha"] = o.Sha
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["author"] = o.Author
	}
	if true {
		toSerialize["committer"] = o.Committer
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["tree"] = o.Tree
	}
	if true {
		toSerialize["parents"] = o.Parents
	}
	if true {
		toSerialize["verification"] = o.Verification
	}
	if true {
		toSerialize["html_url"] = o.HtmlUrl
	}
	return json.Marshal(toSerialize)
}

type NullableGitCommit struct {
	value *GitCommit
	isSet bool
}

func (v NullableGitCommit) Get() *GitCommit {
	return v.value
}

func (v *NullableGitCommit) Set(val *GitCommit) {
	v.value = val
	v.isSet = true
}

func (v NullableGitCommit) IsSet() bool {
	return v.isSet
}

func (v *NullableGitCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitCommit(val *GitCommit) *NullableGitCommit {
	return &NullableGitCommit{value: val, isSet: true}
}

func (v NullableGitCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


