/*
 * Client Portal Web API
 *
 * Production version of the Client Portal Web API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SetAccount struct for SetAccount
type SetAccount struct {
	// Account ID
	AcctId *string `json:"acctId,omitempty"`
}

// NewSetAccount instantiates a new SetAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetAccount() *SetAccount {
	this := SetAccount{}
	return &this
}

// NewSetAccountWithDefaults instantiates a new SetAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetAccountWithDefaults() *SetAccount {
	this := SetAccount{}
	return &this
}

// GetAcctId returns the AcctId field value if set, zero value otherwise.
func (o *SetAccount) GetAcctId() string {
	if o == nil || o.AcctId == nil {
		var ret string
		return ret
	}
	return *o.AcctId
}

// GetAcctIdOk returns a tuple with the AcctId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetAccount) GetAcctIdOk() (*string, bool) {
	if o == nil || o.AcctId == nil {
		return nil, false
	}
	return o.AcctId, true
}

// HasAcctId returns a boolean if a field has been set.
func (o *SetAccount) HasAcctId() bool {
	if o != nil && o.AcctId != nil {
		return true
	}

	return false
}

// SetAcctId gets a reference to the given string and assigns it to the AcctId field.
func (o *SetAccount) SetAcctId(v string) {
	o.AcctId = &v
}

func (o SetAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AcctId != nil {
		toSerialize["acctId"] = o.AcctId
	}
	return json.Marshal(toSerialize)
}

type NullableSetAccount struct {
	value *SetAccount
	isSet bool
}

func (v NullableSetAccount) Get() *SetAccount {
	return v.value
}

func (v *NullableSetAccount) Set(val *SetAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableSetAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableSetAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetAccount(val *SetAccount) *NullableSetAccount {
	return &NullableSetAccount{value: val, isSet: true}
}

func (v NullableSetAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

