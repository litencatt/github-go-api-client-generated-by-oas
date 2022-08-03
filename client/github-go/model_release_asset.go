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

// ReleaseAsset Data related to a release.
type ReleaseAsset struct {
	Url string `json:"url"`
	BrowserDownloadUrl string `json:"browser_download_url"`
	Id int32 `json:"id"`
	NodeId string `json:"node_id"`
	// The file name of the asset.
	Name string `json:"name"`
	Label NullableString `json:"label"`
	// State of the release asset.
	State string `json:"state"`
	ContentType string `json:"content_type"`
	Size int32 `json:"size"`
	DownloadCount int32 `json:"download_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Uploader NullableNullableSimpleUser `json:"uploader"`
}

// NewReleaseAsset instantiates a new ReleaseAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseAsset(url string, browserDownloadUrl string, id int32, nodeId string, name string, label NullableString, state string, contentType string, size int32, downloadCount int32, createdAt time.Time, updatedAt time.Time, uploader NullableNullableSimpleUser) *ReleaseAsset {
	this := ReleaseAsset{}
	this.Url = url
	this.BrowserDownloadUrl = browserDownloadUrl
	this.Id = id
	this.NodeId = nodeId
	this.Name = name
	this.Label = label
	this.State = state
	this.ContentType = contentType
	this.Size = size
	this.DownloadCount = downloadCount
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Uploader = uploader
	return &this
}

// NewReleaseAssetWithDefaults instantiates a new ReleaseAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseAssetWithDefaults() *ReleaseAsset {
	this := ReleaseAsset{}
	return &this
}

// GetUrl returns the Url field value
func (o *ReleaseAsset) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ReleaseAsset) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ReleaseAsset) SetUrl(v string) {
	o.Url = v
}

// GetBrowserDownloadUrl returns the BrowserDownloadUrl field value
func (o *ReleaseAsset) GetBrowserDownloadUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrowserDownloadUrl
}

// GetBrowserDownloadUrlOk returns a tuple with the BrowserDownloadUrl field value
// and a boolean to check if the value has been set.
func (o *ReleaseAsset) GetBrowserDownloadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrowserDownloadUrl, true
}

// SetBrowserDownloadUrl sets field value
func (o *ReleaseAsset) SetBrowserDownloadUrl(v string) {
	o.BrowserDownloadUrl = v
}

// GetId returns the Id field value
func (o *ReleaseAsset) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReleaseAsset) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReleaseAsset) SetId(v int32) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *ReleaseAsset) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ReleaseAsset) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ReleaseAsset) SetNodeId(v string) {
	o.NodeId = v
}

// GetName returns the Name field value
func (o *ReleaseAsset) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReleaseAsset) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReleaseAsset) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ReleaseAsset) GetLabel() string {
	if o == nil || o.Label.Get() == nil {
		var ret string
		return ret
	}

	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReleaseAsset) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// SetLabel sets field value
func (o *ReleaseAsset) SetLabel(v string) {
	o.Label.Set(&v)
}

// GetState returns the State field value
func (o *ReleaseAsset) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ReleaseAsset) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ReleaseAsset) SetState(v string) {
	o.State = v
}

// GetContentType returns the ContentType field value
func (o *ReleaseAsset) GetContentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *ReleaseAsset) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *ReleaseAsset) SetContentType(v string) {
	o.ContentType = v
}

// GetSize returns the Size field value
func (o *ReleaseAsset) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ReleaseAsset) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ReleaseAsset) SetSize(v int32) {
	o.Size = v
}

// GetDownloadCount returns the DownloadCount field value
func (o *ReleaseAsset) GetDownloadCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DownloadCount
}

// GetDownloadCountOk returns a tuple with the DownloadCount field value
// and a boolean to check if the value has been set.
func (o *ReleaseAsset) GetDownloadCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DownloadCount, true
}

// SetDownloadCount sets field value
func (o *ReleaseAsset) SetDownloadCount(v int32) {
	o.DownloadCount = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ReleaseAsset) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ReleaseAsset) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ReleaseAsset) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ReleaseAsset) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ReleaseAsset) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ReleaseAsset) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUploader returns the Uploader field value
// If the value is explicit nil, the zero value for NullableSimpleUser will be returned
func (o *ReleaseAsset) GetUploader() NullableSimpleUser {
	if o == nil || o.Uploader.Get() == nil {
		var ret NullableSimpleUser
		return ret
	}

	return *o.Uploader.Get()
}

// GetUploaderOk returns a tuple with the Uploader field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReleaseAsset) GetUploaderOk() (*NullableSimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uploader.Get(), o.Uploader.IsSet()
}

// SetUploader sets field value
func (o *ReleaseAsset) SetUploader(v NullableSimpleUser) {
	o.Uploader.Set(&v)
}

func (o ReleaseAsset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["browser_download_url"] = o.BrowserDownloadUrl
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["label"] = o.Label.Get()
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["content_type"] = o.ContentType
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["download_count"] = o.DownloadCount
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["uploader"] = o.Uploader.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableReleaseAsset struct {
	value *ReleaseAsset
	isSet bool
}

func (v NullableReleaseAsset) Get() *ReleaseAsset {
	return v.value
}

func (v *NullableReleaseAsset) Set(val *ReleaseAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseAsset(val *ReleaseAsset) *NullableReleaseAsset {
	return &NullableReleaseAsset{value: val, isSet: true}
}

func (v NullableReleaseAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

