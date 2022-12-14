/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"encoding/json"
	"fmt"
)

// DeploymentReviewerType The type of reviewer.
type DeploymentReviewerType string

// List of deployment-reviewer-type
const (
	USER DeploymentReviewerType = "User"
	TEAM DeploymentReviewerType = "Team"
)

// All allowed values of DeploymentReviewerType enum
var AllowedDeploymentReviewerTypeEnumValues = []DeploymentReviewerType{
	"User",
	"Team",
}

func (v *DeploymentReviewerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeploymentReviewerType(value)
	for _, existing := range AllowedDeploymentReviewerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeploymentReviewerType", value)
}

// NewDeploymentReviewerTypeFromValue returns a pointer to a valid DeploymentReviewerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeploymentReviewerTypeFromValue(v string) (*DeploymentReviewerType, error) {
	ev := DeploymentReviewerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeploymentReviewerType: valid values are %v", v, AllowedDeploymentReviewerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeploymentReviewerType) IsValid() bool {
	for _, existing := range AllowedDeploymentReviewerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to deployment-reviewer-type value
func (v DeploymentReviewerType) Ptr() *DeploymentReviewerType {
	return &v
}

type NullableDeploymentReviewerType struct {
	value *DeploymentReviewerType
	isSet bool
}

func (v NullableDeploymentReviewerType) Get() *DeploymentReviewerType {
	return v.value
}

func (v *NullableDeploymentReviewerType) Set(val *DeploymentReviewerType) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentReviewerType) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentReviewerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentReviewerType(val *DeploymentReviewerType) *NullableDeploymentReviewerType {
	return &NullableDeploymentReviewerType{value: val, isSet: true}
}

func (v NullableDeploymentReviewerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentReviewerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

