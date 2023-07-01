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

// StocksContracts struct for StocksContracts
type StocksContracts struct {
	// conid of the stock contract
	Conid *int32 `json:"conid,omitempty"`
	Exchange *string `json:"exchange,omitempty"`
}

// NewStocksContracts instantiates a new StocksContracts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStocksContracts() *StocksContracts {
	this := StocksContracts{}
	return &this
}

// NewStocksContractsWithDefaults instantiates a new StocksContracts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStocksContractsWithDefaults() *StocksContracts {
	this := StocksContracts{}
	return &this
}

// GetConid returns the Conid field value if set, zero value otherwise.
func (o *StocksContracts) GetConid() int32 {
	if o == nil || o.Conid == nil {
		var ret int32
		return ret
	}
	return *o.Conid
}

// GetConidOk returns a tuple with the Conid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StocksContracts) GetConidOk() (*int32, bool) {
	if o == nil || o.Conid == nil {
		return nil, false
	}
	return o.Conid, true
}

// HasConid returns a boolean if a field has been set.
func (o *StocksContracts) HasConid() bool {
	if o != nil && o.Conid != nil {
		return true
	}

	return false
}

// SetConid gets a reference to the given int32 and assigns it to the Conid field.
func (o *StocksContracts) SetConid(v int32) {
	o.Conid = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *StocksContracts) GetExchange() string {
	if o == nil || o.Exchange == nil {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StocksContracts) GetExchangeOk() (*string, bool) {
	if o == nil || o.Exchange == nil {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *StocksContracts) HasExchange() bool {
	if o != nil && o.Exchange != nil {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *StocksContracts) SetExchange(v string) {
	o.Exchange = &v
}

func (o StocksContracts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Conid != nil {
		toSerialize["conid"] = o.Conid
	}
	if o.Exchange != nil {
		toSerialize["exchange"] = o.Exchange
	}
	return json.Marshal(toSerialize)
}

type NullableStocksContracts struct {
	value *StocksContracts
	isSet bool
}

func (v NullableStocksContracts) Get() *StocksContracts {
	return v.value
}

func (v *NullableStocksContracts) Set(val *StocksContracts) {
	v.value = val
	v.isSet = true
}

func (v NullableStocksContracts) IsSet() bool {
	return v.isSet
}

func (v *NullableStocksContracts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStocksContracts(val *StocksContracts) *NullableStocksContracts {
	return &NullableStocksContracts{value: val, isSet: true}
}

func (v NullableStocksContracts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStocksContracts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

