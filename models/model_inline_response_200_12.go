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

// InlineResponse20012 struct for InlineResponse20012
type InlineResponse20012 struct {
	Orders *[]Order `json:"orders,omitempty"`
	Notifications *[]map[string]interface{} `json:"notifications,omitempty"`
}

// NewInlineResponse20012 instantiates a new InlineResponse20012 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20012() *InlineResponse20012 {
	this := InlineResponse20012{}
	return &this
}

// NewInlineResponse20012WithDefaults instantiates a new InlineResponse20012 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20012WithDefaults() *InlineResponse20012 {
	this := InlineResponse20012{}
	return &this
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *InlineResponse20012) GetOrders() []Order {
	if o == nil || o.Orders == nil {
		var ret []Order
		return ret
	}
	return *o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20012) GetOrdersOk() (*[]Order, bool) {
	if o == nil || o.Orders == nil {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *InlineResponse20012) HasOrders() bool {
	if o != nil && o.Orders != nil {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []Order and assigns it to the Orders field.
func (o *InlineResponse20012) SetOrders(v []Order) {
	o.Orders = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *InlineResponse20012) GetNotifications() []map[string]interface{} {
	if o == nil || o.Notifications == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20012) GetNotificationsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Notifications == nil {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *InlineResponse20012) HasNotifications() bool {
	if o != nil && o.Notifications != nil {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []map[string]interface{} and assigns it to the Notifications field.
func (o *InlineResponse20012) SetNotifications(v []map[string]interface{}) {
	o.Notifications = &v
}

func (o InlineResponse20012) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Orders != nil {
		toSerialize["orders"] = o.Orders
	}
	if o.Notifications != nil {
		toSerialize["notifications"] = o.Notifications
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20012 struct {
	value *InlineResponse20012
	isSet bool
}

func (v NullableInlineResponse20012) Get() *InlineResponse20012 {
	return v.value
}

func (v *NullableInlineResponse20012) Set(val *InlineResponse20012) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20012) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20012) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20012(val *InlineResponse20012) *NullableInlineResponse20012 {
	return &NullableInlineResponse20012{value: val, isSet: true}
}

func (v NullableInlineResponse20012) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20012) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

