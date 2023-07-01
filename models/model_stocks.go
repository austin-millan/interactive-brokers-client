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

// Stocks This is an object of object(s), with the symbol(for example, FB) as key and contract details as the object 
type Stocks struct {
	// company name
	Name *string `json:"name,omitempty"`
	// company name in Chinese
	ChineseName *string `json:"chineseName,omitempty"`
	AssetClass *string `json:"assetClass,omitempty"`
	// array of contracts from different exchanges
	Contracts *[]StocksContracts `json:"contracts,omitempty"`
}

// NewStocks instantiates a new Stocks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStocks() *Stocks {
	this := Stocks{}
	return &this
}

// NewStocksWithDefaults instantiates a new Stocks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStocksWithDefaults() *Stocks {
	this := Stocks{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Stocks) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stocks) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Stocks) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Stocks) SetName(v string) {
	o.Name = &v
}

// GetChineseName returns the ChineseName field value if set, zero value otherwise.
func (o *Stocks) GetChineseName() string {
	if o == nil || o.ChineseName == nil {
		var ret string
		return ret
	}
	return *o.ChineseName
}

// GetChineseNameOk returns a tuple with the ChineseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stocks) GetChineseNameOk() (*string, bool) {
	if o == nil || o.ChineseName == nil {
		return nil, false
	}
	return o.ChineseName, true
}

// HasChineseName returns a boolean if a field has been set.
func (o *Stocks) HasChineseName() bool {
	if o != nil && o.ChineseName != nil {
		return true
	}

	return false
}

// SetChineseName gets a reference to the given string and assigns it to the ChineseName field.
func (o *Stocks) SetChineseName(v string) {
	o.ChineseName = &v
}

// GetAssetClass returns the AssetClass field value if set, zero value otherwise.
func (o *Stocks) GetAssetClass() string {
	if o == nil || o.AssetClass == nil {
		var ret string
		return ret
	}
	return *o.AssetClass
}

// GetAssetClassOk returns a tuple with the AssetClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stocks) GetAssetClassOk() (*string, bool) {
	if o == nil || o.AssetClass == nil {
		return nil, false
	}
	return o.AssetClass, true
}

// HasAssetClass returns a boolean if a field has been set.
func (o *Stocks) HasAssetClass() bool {
	if o != nil && o.AssetClass != nil {
		return true
	}

	return false
}

// SetAssetClass gets a reference to the given string and assigns it to the AssetClass field.
func (o *Stocks) SetAssetClass(v string) {
	o.AssetClass = &v
}

// GetContracts returns the Contracts field value if set, zero value otherwise.
func (o *Stocks) GetContracts() []StocksContracts {
	if o == nil || o.Contracts == nil {
		var ret []StocksContracts
		return ret
	}
	return *o.Contracts
}

// GetContractsOk returns a tuple with the Contracts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stocks) GetContractsOk() (*[]StocksContracts, bool) {
	if o == nil || o.Contracts == nil {
		return nil, false
	}
	return o.Contracts, true
}

// HasContracts returns a boolean if a field has been set.
func (o *Stocks) HasContracts() bool {
	if o != nil && o.Contracts != nil {
		return true
	}

	return false
}

// SetContracts gets a reference to the given []StocksContracts and assigns it to the Contracts field.
func (o *Stocks) SetContracts(v []StocksContracts) {
	o.Contracts = &v
}

func (o Stocks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ChineseName != nil {
		toSerialize["chineseName"] = o.ChineseName
	}
	if o.AssetClass != nil {
		toSerialize["assetClass"] = o.AssetClass
	}
	if o.Contracts != nil {
		toSerialize["contracts"] = o.Contracts
	}
	return json.Marshal(toSerialize)
}

type NullableStocks struct {
	value *Stocks
	isSet bool
}

func (v NullableStocks) Get() *Stocks {
	return v.value
}

func (v *NullableStocks) Set(val *Stocks) {
	v.value = val
	v.isSet = true
}

func (v NullableStocks) IsSet() bool {
	return v.isSet
}

func (v *NullableStocks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStocks(val *Stocks) *NullableStocks {
	return &NullableStocks{value: val, isSet: true}
}

func (v NullableStocks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStocks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


