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

// InlineObject1 struct for InlineObject1
type InlineObject1 struct {
	Devicename *string `json:"devicename,omitempty"`
	DeviceId *string `json:"deviceId,omitempty"`
	UiName *string `json:"uiName,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineObject1 instantiates a new InlineObject1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1() *InlineObject1 {
	this := InlineObject1{}
	return &this
}

// NewInlineObject1WithDefaults instantiates a new InlineObject1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1WithDefaults() *InlineObject1 {
	this := InlineObject1{}
	return &this
}

// GetDevicename returns the Devicename field value if set, zero value otherwise.
func (o *InlineObject1) GetDevicename() string {
	if o == nil || o.Devicename == nil {
		var ret string
		return ret
	}
	return *o.Devicename
}

// GetDevicenameOk returns a tuple with the Devicename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetDevicenameOk() (*string, bool) {
	if o == nil || o.Devicename == nil {
		return nil, false
	}
	return o.Devicename, true
}

// HasDevicename returns a boolean if a field has been set.
func (o *InlineObject1) HasDevicename() bool {
	if o != nil && o.Devicename != nil {
		return true
	}

	return false
}

// SetDevicename gets a reference to the given string and assigns it to the Devicename field.
func (o *InlineObject1) SetDevicename(v string) {
	o.Devicename = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *InlineObject1) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *InlineObject1) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *InlineObject1) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetUiName returns the UiName field value if set, zero value otherwise.
func (o *InlineObject1) GetUiName() string {
	if o == nil || o.UiName == nil {
		var ret string
		return ret
	}
	return *o.UiName
}

// GetUiNameOk returns a tuple with the UiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetUiNameOk() (*string, bool) {
	if o == nil || o.UiName == nil {
		return nil, false
	}
	return o.UiName, true
}

// HasUiName returns a boolean if a field has been set.
func (o *InlineObject1) HasUiName() bool {
	if o != nil && o.UiName != nil {
		return true
	}

	return false
}

// SetUiName gets a reference to the given string and assigns it to the UiName field.
func (o *InlineObject1) SetUiName(v string) {
	o.UiName = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject1) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject1) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject1) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineObject1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Devicename != nil {
		toSerialize["devicename"] = o.Devicename
	}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.UiName != nil {
		toSerialize["uiName"] = o.UiName
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1 struct {
	value *InlineObject1
	isSet bool
}

func (v NullableInlineObject1) Get() *InlineObject1 {
	return v.value
}

func (v *NullableInlineObject1) Set(val *InlineObject1) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1(val *InlineObject1) *NullableInlineObject1 {
	return &NullableInlineObject1{value: val, isSet: true}
}

func (v NullableInlineObject1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

