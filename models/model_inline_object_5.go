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

// InlineObject5 struct for InlineObject5
type InlineObject5 struct {
	// Notes for bracket orders: 1. Children orders will not have its own \"cOID\", so please donot pass \"cOID\" parameter in child order.Instead, they will have a \"parentId\" which must be equal to \"cOID\" of parent. 2. When you cancel a parent order, it will cancel all bracket orders, when you cancel one child order, it will also cancel its sibling order. 
	Orders *[]OrderRequest `json:"orders,omitempty"`
}

// NewInlineObject5 instantiates a new InlineObject5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject5() *InlineObject5 {
	this := InlineObject5{}
	return &this
}

// NewInlineObject5WithDefaults instantiates a new InlineObject5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject5WithDefaults() *InlineObject5 {
	this := InlineObject5{}
	return &this
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *InlineObject5) GetOrders() []OrderRequest {
	if o == nil || o.Orders == nil {
		var ret []OrderRequest
		return ret
	}
	return *o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject5) GetOrdersOk() (*[]OrderRequest, bool) {
	if o == nil || o.Orders == nil {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *InlineObject5) HasOrders() bool {
	if o != nil && o.Orders != nil {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []OrderRequest and assigns it to the Orders field.
func (o *InlineObject5) SetOrders(v []OrderRequest) {
	o.Orders = &v
}

func (o InlineObject5) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Orders != nil {
		toSerialize["orders"] = o.Orders
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject5 struct {
	value *InlineObject5
	isSet bool
}

func (v NullableInlineObject5) Get() *InlineObject5 {
	return v.value
}

func (v *NullableInlineObject5) Set(val *InlineObject5) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject5) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject5(val *InlineObject5) *NullableInlineObject5 {
	return &NullableInlineObject5{value: val, isSet: true}
}

func (v NullableInlineObject5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

