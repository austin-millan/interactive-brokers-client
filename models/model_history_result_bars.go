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

// HistoryResultBars struct for HistoryResultBars
type HistoryResultBars struct {
	Open *float32 `json:"open,omitempty"`
	High *float32 `json:"high,omitempty"`
	Low *float32 `json:"low,omitempty"`
	Close *float32 `json:"close,omitempty"`
	Volume *float32 `json:"volume,omitempty"`
	Time *string `json:"time,omitempty"`
	EndTime *string `json:"endTime,omitempty"`
	WeightedAvg *float32 `json:"weightedAvg,omitempty"`
	Count *float32 `json:"count,omitempty"`
}

// NewHistoryResultBars instantiates a new HistoryResultBars object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoryResultBars() *HistoryResultBars {
	this := HistoryResultBars{}
	return &this
}

// NewHistoryResultBarsWithDefaults instantiates a new HistoryResultBars object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoryResultBarsWithDefaults() *HistoryResultBars {
	this := HistoryResultBars{}
	return &this
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *HistoryResultBars) GetOpen() float32 {
	if o == nil || o.Open == nil {
		var ret float32
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResultBars) GetOpenOk() (*float32, bool) {
	if o == nil || o.Open == nil {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *HistoryResultBars) HasOpen() bool {
	if o != nil && o.Open != nil {
		return true
	}

	return false
}

// SetOpen gets a reference to the given float32 and assigns it to the Open field.
func (o *HistoryResultBars) SetOpen(v float32) {
	o.Open = &v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *HistoryResultBars) GetHigh() float32 {
	if o == nil || o.High == nil {
		var ret float32
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResultBars) GetHighOk() (*float32, bool) {
	if o == nil || o.High == nil {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *HistoryResultBars) HasHigh() bool {
	if o != nil && o.High != nil {
		return true
	}

	return false
}

// SetHigh gets a reference to the given float32 and assigns it to the High field.
func (o *HistoryResultBars) SetHigh(v float32) {
	o.High = &v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *HistoryResultBars) GetLow() float32 {
	if o == nil || o.Low == nil {
		var ret float32
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResultBars) GetLowOk() (*float32, bool) {
	if o == nil || o.Low == nil {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *HistoryResultBars) HasLow() bool {
	if o != nil && o.Low != nil {
		return true
	}

	return false
}

// SetLow gets a reference to the given float32 and assigns it to the Low field.
func (o *HistoryResultBars) SetLow(v float32) {
	o.Low = &v
}

// GetClose returns the Close field value if set, zero value otherwise.
func (o *HistoryResultBars) GetClose() float32 {
	if o == nil || o.Close == nil {
		var ret float32
		return ret
	}
	return *o.Close
}

// GetCloseOk returns a tuple with the Close field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResultBars) GetCloseOk() (*float32, bool) {
	if o == nil || o.Close == nil {
		return nil, false
	}
	return o.Close, true
}

// HasClose returns a boolean if a field has been set.
func (o *HistoryResultBars) HasClose() bool {
	if o != nil && o.Close != nil {
		return true
	}

	return false
}

// SetClose gets a reference to the given float32 and assigns it to the Close field.
func (o *HistoryResultBars) SetClose(v float32) {
	o.Close = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *HistoryResultBars) GetVolume() float32 {
	if o == nil || o.Volume == nil {
		var ret float32
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResultBars) GetVolumeOk() (*float32, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *HistoryResultBars) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given float32 and assigns it to the Volume field.
func (o *HistoryResultBars) SetVolume(v float32) {
	o.Volume = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *HistoryResultBars) GetTime() string {
	if o == nil || o.Time == nil {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResultBars) GetTimeOk() (*string, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *HistoryResultBars) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *HistoryResultBars) SetTime(v string) {
	o.Time = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *HistoryResultBars) GetEndTime() string {
	if o == nil || o.EndTime == nil {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResultBars) GetEndTimeOk() (*string, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *HistoryResultBars) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *HistoryResultBars) SetEndTime(v string) {
	o.EndTime = &v
}

// GetWeightedAvg returns the WeightedAvg field value if set, zero value otherwise.
func (o *HistoryResultBars) GetWeightedAvg() float32 {
	if o == nil || o.WeightedAvg == nil {
		var ret float32
		return ret
	}
	return *o.WeightedAvg
}

// GetWeightedAvgOk returns a tuple with the WeightedAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResultBars) GetWeightedAvgOk() (*float32, bool) {
	if o == nil || o.WeightedAvg == nil {
		return nil, false
	}
	return o.WeightedAvg, true
}

// HasWeightedAvg returns a boolean if a field has been set.
func (o *HistoryResultBars) HasWeightedAvg() bool {
	if o != nil && o.WeightedAvg != nil {
		return true
	}

	return false
}

// SetWeightedAvg gets a reference to the given float32 and assigns it to the WeightedAvg field.
func (o *HistoryResultBars) SetWeightedAvg(v float32) {
	o.WeightedAvg = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *HistoryResultBars) GetCount() float32 {
	if o == nil || o.Count == nil {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoryResultBars) GetCountOk() (*float32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *HistoryResultBars) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *HistoryResultBars) SetCount(v float32) {
	o.Count = &v
}

func (o HistoryResultBars) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Open != nil {
		toSerialize["open"] = o.Open
	}
	if o.High != nil {
		toSerialize["high"] = o.High
	}
	if o.Low != nil {
		toSerialize["low"] = o.Low
	}
	if o.Close != nil {
		toSerialize["close"] = o.Close
	}
	if o.Volume != nil {
		toSerialize["volume"] = o.Volume
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.EndTime != nil {
		toSerialize["endTime"] = o.EndTime
	}
	if o.WeightedAvg != nil {
		toSerialize["weightedAvg"] = o.WeightedAvg
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableHistoryResultBars struct {
	value *HistoryResultBars
	isSet bool
}

func (v NullableHistoryResultBars) Get() *HistoryResultBars {
	return v.value
}

func (v *NullableHistoryResultBars) Set(val *HistoryResultBars) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoryResultBars) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoryResultBars) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoryResultBars(val *HistoryResultBars) *NullableHistoryResultBars {
	return &NullableHistoryResultBars{value: val, isSet: true}
}

func (v NullableHistoryResultBars) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoryResultBars) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


