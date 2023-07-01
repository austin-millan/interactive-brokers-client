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

// InlineObject3 struct for InlineObject3
type InlineObject3 struct {
	AcctIds *[]string `json:"acctIds,omitempty"`
}

// NewInlineObject3 instantiates a new InlineObject3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject3() *InlineObject3 {
	this := InlineObject3{}
	return &this
}

// NewInlineObject3WithDefaults instantiates a new InlineObject3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject3WithDefaults() *InlineObject3 {
	this := InlineObject3{}
	return &this
}

// GetAcctIds returns the AcctIds field value if set, zero value otherwise.
func (o *InlineObject3) GetAcctIds() []string {
	if o == nil || o.AcctIds == nil {
		var ret []string
		return ret
	}
	return *o.AcctIds
}

// GetAcctIdsOk returns a tuple with the AcctIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetAcctIdsOk() (*[]string, bool) {
	if o == nil || o.AcctIds == nil {
		return nil, false
	}
	return o.AcctIds, true
}

// HasAcctIds returns a boolean if a field has been set.
func (o *InlineObject3) HasAcctIds() bool {
	if o != nil && o.AcctIds != nil {
		return true
	}

	return false
}

// SetAcctIds gets a reference to the given []string and assigns it to the AcctIds field.
func (o *InlineObject3) SetAcctIds(v []string) {
	o.AcctIds = &v
}

func (o InlineObject3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AcctIds != nil {
		toSerialize["acctIds"] = o.AcctIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject3 struct {
	value *InlineObject3
	isSet bool
}

func (v NullableInlineObject3) Get() *InlineObject3 {
	return v.value
}

func (v *NullableInlineObject3) Set(val *InlineObject3) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject3) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject3(val *InlineObject3) *NullableInlineObject3 {
	return &NullableInlineObject3{value: val, isSet: true}
}

func (v NullableInlineObject3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

