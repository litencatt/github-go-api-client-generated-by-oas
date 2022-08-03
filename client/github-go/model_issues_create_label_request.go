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

// IssuesCreateLabelRequest struct for IssuesCreateLabelRequest
type IssuesCreateLabelRequest struct {
	// The name of the label. Emoji can be added to label names, using either native emoji or colon-style markup. For example, typing `:strawberry:` will render the emoji ![:strawberry:](https://github.githubassets.com/images/icons/emoji/unicode/1f353.png \":strawberry:\"). For a full list of available emoji and codes, see \"[Emoji cheat sheet](https://github.com/ikatyang/emoji-cheat-sheet).\"
	Name string `json:"name"`
	// The [hexadecimal color code](http://www.color-hex.com/) for the label, without the leading `#`.
	Color *string `json:"color,omitempty"`
	// A short description of the label. Must be 100 characters or fewer.
	Description *string `json:"description,omitempty"`
}

// NewIssuesCreateLabelRequest instantiates a new IssuesCreateLabelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssuesCreateLabelRequest(name string) *IssuesCreateLabelRequest {
	this := IssuesCreateLabelRequest{}
	this.Name = name
	return &this
}

// NewIssuesCreateLabelRequestWithDefaults instantiates a new IssuesCreateLabelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssuesCreateLabelRequestWithDefaults() *IssuesCreateLabelRequest {
	this := IssuesCreateLabelRequest{}
	return &this
}

// GetName returns the Name field value
func (o *IssuesCreateLabelRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IssuesCreateLabelRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IssuesCreateLabelRequest) SetName(v string) {
	o.Name = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *IssuesCreateLabelRequest) GetColor() string {
	if o == nil || o.Color == nil {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesCreateLabelRequest) GetColorOk() (*string, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *IssuesCreateLabelRequest) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *IssuesCreateLabelRequest) SetColor(v string) {
	o.Color = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IssuesCreateLabelRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesCreateLabelRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IssuesCreateLabelRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IssuesCreateLabelRequest) SetDescription(v string) {
	o.Description = &v
}

func (o IssuesCreateLabelRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableIssuesCreateLabelRequest struct {
	value *IssuesCreateLabelRequest
	isSet bool
}

func (v NullableIssuesCreateLabelRequest) Get() *IssuesCreateLabelRequest {
	return v.value
}

func (v *NullableIssuesCreateLabelRequest) Set(val *IssuesCreateLabelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIssuesCreateLabelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIssuesCreateLabelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssuesCreateLabelRequest(val *IssuesCreateLabelRequest) *NullableIssuesCreateLabelRequest {
	return &NullableIssuesCreateLabelRequest{value: val, isSet: true}
}

func (v NullableIssuesCreateLabelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssuesCreateLabelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


