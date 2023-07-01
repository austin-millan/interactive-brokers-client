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

// InlineResponse2005 struct for InlineResponse2005
type InlineResponse2005 struct {
	// true means username is still logged in, false means it is not
	Confirmed *bool `json:"confirmed,omitempty"`
}

// NewInlineResponse2005 instantiates a new InlineResponse2005 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2005() *InlineResponse2005 {
	this := InlineResponse2005{}
	return &this
}

// NewInlineResponse2005WithDefaults instantiates a new InlineResponse2005 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2005WithDefaults() *InlineResponse2005 {
	this := InlineResponse2005{}
	return &this
}

// GetConfirmed returns the Confirmed field value if set, zero value otherwise.
func (o *InlineResponse2005) GetConfirmed() bool {
	if o == nil || o.Confirmed == nil {
		var ret bool
		return ret
	}
	return *o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005) GetConfirmedOk() (*bool, bool) {
	if o == nil || o.Confirmed == nil {
		return nil, false
	}
	return o.Confirmed, true
}

// HasConfirmed returns a boolean if a field has been set.
func (o *InlineResponse2005) HasConfirmed() bool {
	if o != nil && o.Confirmed != nil {
		return true
	}

	return false
}

// SetConfirmed gets a reference to the given bool and assigns it to the Confirmed field.
func (o *InlineResponse2005) SetConfirmed(v bool) {
	o.Confirmed = &v
}

func (o InlineResponse2005) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Confirmed != nil {
		toSerialize["confirmed"] = o.Confirmed
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2005 struct {
	value *InlineResponse2005
	isSet bool
}

func (v NullableInlineResponse2005) Get() *InlineResponse2005 {
	return v.value
}

func (v *NullableInlineResponse2005) Set(val *InlineResponse2005) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2005) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2005) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2005(val *InlineResponse2005) *NullableInlineResponse2005 {
	return &NullableInlineResponse2005{value: val, isSet: true}
}

func (v NullableInlineResponse2005) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2005) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


