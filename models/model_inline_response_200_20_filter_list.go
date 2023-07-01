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

// InlineResponse20020FilterList struct for InlineResponse20020FilterList
type InlineResponse20020FilterList struct {
	Group *string `json:"group,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Code *string `json:"code,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewInlineResponse20020FilterList instantiates a new InlineResponse20020FilterList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20020FilterList() *InlineResponse20020FilterList {
	this := InlineResponse20020FilterList{}
	return &this
}

// NewInlineResponse20020FilterListWithDefaults instantiates a new InlineResponse20020FilterList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20020FilterListWithDefaults() *InlineResponse20020FilterList {
	this := InlineResponse20020FilterList{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *InlineResponse20020FilterList) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020FilterList) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *InlineResponse20020FilterList) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *InlineResponse20020FilterList) SetGroup(v string) {
	o.Group = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *InlineResponse20020FilterList) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020FilterList) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *InlineResponse20020FilterList) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *InlineResponse20020FilterList) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InlineResponse20020FilterList) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020FilterList) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InlineResponse20020FilterList) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InlineResponse20020FilterList) SetCode(v string) {
	o.Code = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse20020FilterList) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20020FilterList) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse20020FilterList) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse20020FilterList) SetType(v string) {
	o.Type = &v
}

func (o InlineResponse20020FilterList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20020FilterList struct {
	value *InlineResponse20020FilterList
	isSet bool
}

func (v NullableInlineResponse20020FilterList) Get() *InlineResponse20020FilterList {
	return v.value
}

func (v *NullableInlineResponse20020FilterList) Set(val *InlineResponse20020FilterList) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20020FilterList) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20020FilterList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20020FilterList(val *InlineResponse20020FilterList) *NullableInlineResponse20020FilterList {
	return &NullableInlineResponse20020FilterList{value: val, isSet: true}
}

func (v NullableInlineResponse20020FilterList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20020FilterList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

