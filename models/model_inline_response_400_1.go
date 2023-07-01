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

// InlineResponse4001 struct for InlineResponse4001
type InlineResponse4001 struct {
	Error *string `json:"error,omitempty"`
	StatusCode *int32 `json:"statusCode,omitempty"`
}

// NewInlineResponse4001 instantiates a new InlineResponse4001 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse4001() *InlineResponse4001 {
	this := InlineResponse4001{}
	return &this
}

// NewInlineResponse4001WithDefaults instantiates a new InlineResponse4001 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse4001WithDefaults() *InlineResponse4001 {
	this := InlineResponse4001{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *InlineResponse4001) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse4001) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *InlineResponse4001) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *InlineResponse4001) SetError(v string) {
	o.Error = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *InlineResponse4001) GetStatusCode() int32 {
	if o == nil || o.StatusCode == nil {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse4001) GetStatusCodeOk() (*int32, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *InlineResponse4001) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *InlineResponse4001) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o InlineResponse4001) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.StatusCode != nil {
		toSerialize["statusCode"] = o.StatusCode
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse4001 struct {
	value *InlineResponse4001
	isSet bool
}

func (v NullableInlineResponse4001) Get() *InlineResponse4001 {
	return v.value
}

func (v *NullableInlineResponse4001) Set(val *InlineResponse4001) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse4001) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse4001) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse4001(val *InlineResponse4001) *NullableInlineResponse4001 {
	return &NullableInlineResponse4001{value: val, isSet: true}
}

func (v NullableInlineResponse4001) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse4001) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


