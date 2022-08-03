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
	"fmt"
)

// InstallationAccount struct for InstallationAccount
type InstallationAccount struct {
	Enterprise *Enterprise
	SimpleUser *SimpleUser
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *InstallationAccount) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into Enterprise
	err = json.Unmarshal(data, &dst.Enterprise);
	if err == nil {
		jsonEnterprise, _ := json.Marshal(dst.Enterprise)
		if string(jsonEnterprise) == "{}" { // empty struct
			dst.Enterprise = nil
		} else {
			return nil // data stored in dst.Enterprise, return on the first match
		}
	} else {
		dst.Enterprise = nil
	}

	// try to unmarshal JSON data into SimpleUser
	err = json.Unmarshal(data, &dst.SimpleUser);
	if err == nil {
		jsonSimpleUser, _ := json.Marshal(dst.SimpleUser)
		if string(jsonSimpleUser) == "{}" { // empty struct
			dst.SimpleUser = nil
		} else {
			return nil // data stored in dst.SimpleUser, return on the first match
		}
	} else {
		dst.SimpleUser = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(InstallationAccount)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *InstallationAccount) MarshalJSON() ([]byte, error) {
	if src.Enterprise != nil {
		return json.Marshal(&src.Enterprise)
	}

	if src.SimpleUser != nil {
		return json.Marshal(&src.SimpleUser)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableInstallationAccount struct {
	value *InstallationAccount
	isSet bool
}

func (v NullableInstallationAccount) Get() *InstallationAccount {
	return v.value
}

func (v *NullableInstallationAccount) Set(val *InstallationAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallationAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallationAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallationAccount(val *InstallationAccount) *NullableInstallationAccount {
	return &NullableInstallationAccount{value: val, isSet: true}
}

func (v NullableInstallationAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallationAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


