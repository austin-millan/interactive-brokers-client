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

// MarketData struct for MarketData
type MarketData struct {
	// IBKR Contract ID
	Conid *float32 `json:"Conid,omitempty"`
	// Exchange
	Exchange *string `json:"Exchange,omitempty"`
	MinTick *float32 `json:"minTick,omitempty"`
	Last *float32 `json:"Last,omitempty"`
	LastSize *float32 `json:"LastSize,omitempty"`
	Bid *float32 `json:"Bid,omitempty"`
	BidSize *float32 `json:"BidSize,omitempty"`
	Ask *float32 `json:"Ask,omitempty"`
	AskSize *float32 `json:"AskSize,omitempty"`
}

// NewMarketData instantiates a new MarketData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketData() *MarketData {
	this := MarketData{}
	return &this
}

// NewMarketDataWithDefaults instantiates a new MarketData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketDataWithDefaults() *MarketData {
	this := MarketData{}
	return &this
}

// GetConid returns the Conid field value if set, zero value otherwise.
func (o *MarketData) GetConid() float32 {
	if o == nil || o.Conid == nil {
		var ret float32
		return ret
	}
	return *o.Conid
}

// GetConidOk returns a tuple with the Conid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketData) GetConidOk() (*float32, bool) {
	if o == nil || o.Conid == nil {
		return nil, false
	}
	return o.Conid, true
}

// HasConid returns a boolean if a field has been set.
func (o *MarketData) HasConid() bool {
	if o != nil && o.Conid != nil {
		return true
	}

	return false
}

// SetConid gets a reference to the given float32 and assigns it to the Conid field.
func (o *MarketData) SetConid(v float32) {
	o.Conid = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *MarketData) GetExchange() string {
	if o == nil || o.Exchange == nil {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketData) GetExchangeOk() (*string, bool) {
	if o == nil || o.Exchange == nil {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *MarketData) HasExchange() bool {
	if o != nil && o.Exchange != nil {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *MarketData) SetExchange(v string) {
	o.Exchange = &v
}

// GetMinTick returns the MinTick field value if set, zero value otherwise.
func (o *MarketData) GetMinTick() float32 {
	if o == nil || o.MinTick == nil {
		var ret float32
		return ret
	}
	return *o.MinTick
}

// GetMinTickOk returns a tuple with the MinTick field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketData) GetMinTickOk() (*float32, bool) {
	if o == nil || o.MinTick == nil {
		return nil, false
	}
	return o.MinTick, true
}

// HasMinTick returns a boolean if a field has been set.
func (o *MarketData) HasMinTick() bool {
	if o != nil && o.MinTick != nil {
		return true
	}

	return false
}

// SetMinTick gets a reference to the given float32 and assigns it to the MinTick field.
func (o *MarketData) SetMinTick(v float32) {
	o.MinTick = &v
}

// GetLast returns the Last field value if set, zero value otherwise.
func (o *MarketData) GetLast() float32 {
	if o == nil || o.Last == nil {
		var ret float32
		return ret
	}
	return *o.Last
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketData) GetLastOk() (*float32, bool) {
	if o == nil || o.Last == nil {
		return nil, false
	}
	return o.Last, true
}

// HasLast returns a boolean if a field has been set.
func (o *MarketData) HasLast() bool {
	if o != nil && o.Last != nil {
		return true
	}

	return false
}

// SetLast gets a reference to the given float32 and assigns it to the Last field.
func (o *MarketData) SetLast(v float32) {
	o.Last = &v
}

// GetLastSize returns the LastSize field value if set, zero value otherwise.
func (o *MarketData) GetLastSize() float32 {
	if o == nil || o.LastSize == nil {
		var ret float32
		return ret
	}
	return *o.LastSize
}

// GetLastSizeOk returns a tuple with the LastSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketData) GetLastSizeOk() (*float32, bool) {
	if o == nil || o.LastSize == nil {
		return nil, false
	}
	return o.LastSize, true
}

// HasLastSize returns a boolean if a field has been set.
func (o *MarketData) HasLastSize() bool {
	if o != nil && o.LastSize != nil {
		return true
	}

	return false
}

// SetLastSize gets a reference to the given float32 and assigns it to the LastSize field.
func (o *MarketData) SetLastSize(v float32) {
	o.LastSize = &v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *MarketData) GetBid() float32 {
	if o == nil || o.Bid == nil {
		var ret float32
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketData) GetBidOk() (*float32, bool) {
	if o == nil || o.Bid == nil {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *MarketData) HasBid() bool {
	if o != nil && o.Bid != nil {
		return true
	}

	return false
}

// SetBid gets a reference to the given float32 and assigns it to the Bid field.
func (o *MarketData) SetBid(v float32) {
	o.Bid = &v
}

// GetBidSize returns the BidSize field value if set, zero value otherwise.
func (o *MarketData) GetBidSize() float32 {
	if o == nil || o.BidSize == nil {
		var ret float32
		return ret
	}
	return *o.BidSize
}

// GetBidSizeOk returns a tuple with the BidSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketData) GetBidSizeOk() (*float32, bool) {
	if o == nil || o.BidSize == nil {
		return nil, false
	}
	return o.BidSize, true
}

// HasBidSize returns a boolean if a field has been set.
func (o *MarketData) HasBidSize() bool {
	if o != nil && o.BidSize != nil {
		return true
	}

	return false
}

// SetBidSize gets a reference to the given float32 and assigns it to the BidSize field.
func (o *MarketData) SetBidSize(v float32) {
	o.BidSize = &v
}

// GetAsk returns the Ask field value if set, zero value otherwise.
func (o *MarketData) GetAsk() float32 {
	if o == nil || o.Ask == nil {
		var ret float32
		return ret
	}
	return *o.Ask
}

// GetAskOk returns a tuple with the Ask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketData) GetAskOk() (*float32, bool) {
	if o == nil || o.Ask == nil {
		return nil, false
	}
	return o.Ask, true
}

// HasAsk returns a boolean if a field has been set.
func (o *MarketData) HasAsk() bool {
	if o != nil && o.Ask != nil {
		return true
	}

	return false
}

// SetAsk gets a reference to the given float32 and assigns it to the Ask field.
func (o *MarketData) SetAsk(v float32) {
	o.Ask = &v
}

// GetAskSize returns the AskSize field value if set, zero value otherwise.
func (o *MarketData) GetAskSize() float32 {
	if o == nil || o.AskSize == nil {
		var ret float32
		return ret
	}
	return *o.AskSize
}

// GetAskSizeOk returns a tuple with the AskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketData) GetAskSizeOk() (*float32, bool) {
	if o == nil || o.AskSize == nil {
		return nil, false
	}
	return o.AskSize, true
}

// HasAskSize returns a boolean if a field has been set.
func (o *MarketData) HasAskSize() bool {
	if o != nil && o.AskSize != nil {
		return true
	}

	return false
}

// SetAskSize gets a reference to the given float32 and assigns it to the AskSize field.
func (o *MarketData) SetAskSize(v float32) {
	o.AskSize = &v
}

func (o MarketData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Conid != nil {
		toSerialize["Conid"] = o.Conid
	}
	if o.Exchange != nil {
		toSerialize["Exchange"] = o.Exchange
	}
	if o.MinTick != nil {
		toSerialize["minTick"] = o.MinTick
	}
	if o.Last != nil {
		toSerialize["Last"] = o.Last
	}
	if o.LastSize != nil {
		toSerialize["LastSize"] = o.LastSize
	}
	if o.Bid != nil {
		toSerialize["Bid"] = o.Bid
	}
	if o.BidSize != nil {
		toSerialize["BidSize"] = o.BidSize
	}
	if o.Ask != nil {
		toSerialize["Ask"] = o.Ask
	}
	if o.AskSize != nil {
		toSerialize["AskSize"] = o.AskSize
	}
	return json.Marshal(toSerialize)
}

type NullableMarketData struct {
	value *MarketData
	isSet bool
}

func (v NullableMarketData) Get() *MarketData {
	return v.value
}

func (v *NullableMarketData) Set(val *MarketData) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketData) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketData(val *MarketData) *NullableMarketData {
	return &NullableMarketData{value: val, isSet: true}
}

func (v NullableMarketData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


