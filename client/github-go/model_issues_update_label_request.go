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

// IssuesUpdateLabelRequest struct for IssuesUpdateLabelRequest
type IssuesUpdateLabelRequest struct {
	// The new name of the label. Emoji can be added to label names, using either native emoji or colon-style markup. For example, typing `:strawberry:` will render the emoji ![:strawberry:](https://github.githubassets.com/images/icons/emoji/unicode/1f353.png \":strawberry:\"). For a full list of available emoji and codes, see \"[Emoji cheat sheet](https://github.com/ikatyang/emoji-cheat-sheet).\"
	NewName *string `json:"new_name,omitempty"`
	// The [hexadecimal color code](http://www.color-hex.com/) for the label, without the leading `#`.
	Color *string `json:"color,omitempty"`
	// A short description of the label. Must be 100 characters or fewer.
	Description *string `json:"description,omitempty"`
}

// NewIssuesUpdateLabelRequest instantiates a new IssuesUpdateLabelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuesUpdateLabelRequest() *IssuesUpdateLabelRequest {
	this := IssuesUpdateLabelRequest{}
	return &this
}

// NewIssuesUpdateLabelRequestWithDefaults instantiates a new IssuesUpdateLabelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuesUpdateLabelRequestWithDefaults() *IssuesUpdateLabelRequest {
	this := IssuesUpdateLabelRequest{}
	return &this
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *IssuesUpdateLabelRequest) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesUpdateLabelRequest) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *IssuesUpdateLabelRequest) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *IssuesUpdateLabelRequest) SetNewName(v string) {
	o.NewName = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *IssuesUpdateLabelRequest) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesUpdateLabelRequest) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *IssuesUpdateLabelRequest) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *IssuesUpdateLabelRequest) SetColor(v string) {
	o.Color = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IssuesUpdateLabelRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesUpdateLabelRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IssuesUpdateLabelRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IssuesUpdateLabelRequest) SetDescription(v string) {
	o.Description = &v
}

func (o IssuesUpdateLabelRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NewName != nil {
		toSerialize["new_name"] = o.NewName
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableIssuesUpdateLabelRequest struct {
	value *IssuesUpdateLabelRequest
	isSet bool
}

func (v NullableIssuesUpdateLabelRequest) Get() *IssuesUpdateLabelRequest {
	return v.value
}

func (v *NullableIssuesUpdateLabelRequest) Set(val *IssuesUpdateLabelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuesUpdateLabelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuesUpdateLabelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuesUpdateLabelRequest(val *IssuesUpdateLabelRequest) *NullableIssuesUpdateLabelRequest {
	return &NullableIssuesUpdateLabelRequest{value: val, isSet: true}
}

func (v NullableIssuesUpdateLabelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuesUpdateLabelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

