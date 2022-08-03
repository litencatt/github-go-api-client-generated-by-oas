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

// AllowedActions The permissions policy that controls the actions and reusable workflows that are allowed to run.
type AllowedActions string

// List of allowed-actions
const (
	ALL AllowedActions = "all"
	LOCAL_ONLY AllowedActions = "local_only"
	SELECTED AllowedActions = "selected"
)

// All allowed values of AllowedActions enum
var AllowedAllowedActionsEnumValues = []AllowedActions{
	"all",
	"local_only",
	"selected",
}

func (v *AllowedActions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AllowedActions(value)
	for _, existing := range AllowedAllowedActionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AllowedActions", value)
}

// NewAllowedActionsFromValue returns a pointer to a valid AllowedActions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAllowedActionsFromValue(v string) (*AllowedActions, error) {
	ev := AllowedActions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AllowedActions: valid values are %v", v, AllowedAllowedActionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AllowedActions) IsValid() bool {
	for _, existing := range AllowedAllowedActionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to allowed-actions value
func (v AllowedActions) Ptr() *AllowedActions {
	return &v
}

type NullableAllowedActions struct {
	value *AllowedActions
	isSet bool
}

func (v NullableAllowedActions) Get() *AllowedActions {
	return v.value
}

func (v *NullableAllowedActions) Set(val *AllowedActions) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowedActions) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowedActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowedActions(val *AllowedActions) *NullableAllowedActions {
	return &NullableAllowedActions{value: val, isSet: true}
}

func (v NullableAllowedActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowedActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
