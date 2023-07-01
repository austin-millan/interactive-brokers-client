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

// InlineResponse20021 struct for InlineResponse20021
type InlineResponse20021 struct {
	ServerId *string `json:"server_id,omitempty"`
	ColumnName *string `json:"column_name,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Conidex *string `json:"conidex,omitempty"`
	ConId *float32 `json:"con_id,omitempty"`
	AvailableChartPeriods *string `json:"available_chart_periods,omitempty"`
	CompanyName *string `json:"company_name,omitempty"`
	ContractDescription1 *string `json:"contract_description_1,omitempty"`
	ListingExchange *string `json:"listing_exchange,omitempty"`
	SecType *string `json:"sec_type,omitempty"`
}

// NewInlineResponse20021 instantiates a new InlineResponse20021 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20021() *InlineResponse20021 {
	this := InlineResponse20021{}
	return &this
}

// NewInlineResponse20021WithDefaults instantiates a new InlineResponse20021 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20021WithDefaults() *InlineResponse20021 {
	this := InlineResponse20021{}
	return &this
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *InlineResponse20021) GetServerId() string {
	if o == nil || o.ServerId == nil {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetServerIdOk() (*string, bool) {
	if o == nil || o.ServerId == nil {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *InlineResponse20021) HasServerId() bool {
	if o != nil && o.ServerId != nil {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *InlineResponse20021) SetServerId(v string) {
	o.ServerId = &v
}

// GetColumnName returns the ColumnName field value if set, zero value otherwise.
func (o *InlineResponse20021) GetColumnName() string {
	if o == nil || o.ColumnName == nil {
		var ret string
		return ret
	}
	return *o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetColumnNameOk() (*string, bool) {
	if o == nil || o.ColumnName == nil {
		return nil, false
	}
	return o.ColumnName, true
}

// HasColumnName returns a boolean if a field has been set.
func (o *InlineResponse20021) HasColumnName() bool {
	if o != nil && o.ColumnName != nil {
		return true
	}

	return false
}

// SetColumnName gets a reference to the given string and assigns it to the ColumnName field.
func (o *InlineResponse20021) SetColumnName(v string) {
	o.ColumnName = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *InlineResponse20021) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *InlineResponse20021) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *InlineResponse20021) SetSymbol(v string) {
	o.Symbol = &v
}

// GetConidex returns the Conidex field value if set, zero value otherwise.
func (o *InlineResponse20021) GetConidex() string {
	if o == nil || o.Conidex == nil {
		var ret string
		return ret
	}
	return *o.Conidex
}

// GetConidexOk returns a tuple with the Conidex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetConidexOk() (*string, bool) {
	if o == nil || o.Conidex == nil {
		return nil, false
	}
	return o.Conidex, true
}

// HasConidex returns a boolean if a field has been set.
func (o *InlineResponse20021) HasConidex() bool {
	if o != nil && o.Conidex != nil {
		return true
	}

	return false
}

// SetConidex gets a reference to the given string and assigns it to the Conidex field.
func (o *InlineResponse20021) SetConidex(v string) {
	o.Conidex = &v
}

// GetConId returns the ConId field value if set, zero value otherwise.
func (o *InlineResponse20021) GetConId() float32 {
	if o == nil || o.ConId == nil {
		var ret float32
		return ret
	}
	return *o.ConId
}

// GetConIdOk returns a tuple with the ConId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetConIdOk() (*float32, bool) {
	if o == nil || o.ConId == nil {
		return nil, false
	}
	return o.ConId, true
}

// HasConId returns a boolean if a field has been set.
func (o *InlineResponse20021) HasConId() bool {
	if o != nil && o.ConId != nil {
		return true
	}

	return false
}

// SetConId gets a reference to the given float32 and assigns it to the ConId field.
func (o *InlineResponse20021) SetConId(v float32) {
	o.ConId = &v
}

// GetAvailableChartPeriods returns the AvailableChartPeriods field value if set, zero value otherwise.
func (o *InlineResponse20021) GetAvailableChartPeriods() string {
	if o == nil || o.AvailableChartPeriods == nil {
		var ret string
		return ret
	}
	return *o.AvailableChartPeriods
}

// GetAvailableChartPeriodsOk returns a tuple with the AvailableChartPeriods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetAvailableChartPeriodsOk() (*string, bool) {
	if o == nil || o.AvailableChartPeriods == nil {
		return nil, false
	}
	return o.AvailableChartPeriods, true
}

// HasAvailableChartPeriods returns a boolean if a field has been set.
func (o *InlineResponse20021) HasAvailableChartPeriods() bool {
	if o != nil && o.AvailableChartPeriods != nil {
		return true
	}

	return false
}

// SetAvailableChartPeriods gets a reference to the given string and assigns it to the AvailableChartPeriods field.
func (o *InlineResponse20021) SetAvailableChartPeriods(v string) {
	o.AvailableChartPeriods = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *InlineResponse20021) GetCompanyName() string {
	if o == nil || o.CompanyName == nil {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetCompanyNameOk() (*string, bool) {
	if o == nil || o.CompanyName == nil {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *InlineResponse20021) HasCompanyName() bool {
	if o != nil && o.CompanyName != nil {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *InlineResponse20021) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetContractDescription1 returns the ContractDescription1 field value if set, zero value otherwise.
func (o *InlineResponse20021) GetContractDescription1() string {
	if o == nil || o.ContractDescription1 == nil {
		var ret string
		return ret
	}
	return *o.ContractDescription1
}

// GetContractDescription1Ok returns a tuple with the ContractDescription1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetContractDescription1Ok() (*string, bool) {
	if o == nil || o.ContractDescription1 == nil {
		return nil, false
	}
	return o.ContractDescription1, true
}

// HasContractDescription1 returns a boolean if a field has been set.
func (o *InlineResponse20021) HasContractDescription1() bool {
	if o != nil && o.ContractDescription1 != nil {
		return true
	}

	return false
}

// SetContractDescription1 gets a reference to the given string and assigns it to the ContractDescription1 field.
func (o *InlineResponse20021) SetContractDescription1(v string) {
	o.ContractDescription1 = &v
}

// GetListingExchange returns the ListingExchange field value if set, zero value otherwise.
func (o *InlineResponse20021) GetListingExchange() string {
	if o == nil || o.ListingExchange == nil {
		var ret string
		return ret
	}
	return *o.ListingExchange
}

// GetListingExchangeOk returns a tuple with the ListingExchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetListingExchangeOk() (*string, bool) {
	if o == nil || o.ListingExchange == nil {
		return nil, false
	}
	return o.ListingExchange, true
}

// HasListingExchange returns a boolean if a field has been set.
func (o *InlineResponse20021) HasListingExchange() bool {
	if o != nil && o.ListingExchange != nil {
		return true
	}

	return false
}

// SetListingExchange gets a reference to the given string and assigns it to the ListingExchange field.
func (o *InlineResponse20021) SetListingExchange(v string) {
	o.ListingExchange = &v
}

// GetSecType returns the SecType field value if set, zero value otherwise.
func (o *InlineResponse20021) GetSecType() string {
	if o == nil || o.SecType == nil {
		var ret string
		return ret
	}
	return *o.SecType
}

// GetSecTypeOk returns a tuple with the SecType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20021) GetSecTypeOk() (*string, bool) {
	if o == nil || o.SecType == nil {
		return nil, false
	}
	return o.SecType, true
}

// HasSecType returns a boolean if a field has been set.
func (o *InlineResponse20021) HasSecType() bool {
	if o != nil && o.SecType != nil {
		return true
	}

	return false
}

// SetSecType gets a reference to the given string and assigns it to the SecType field.
func (o *InlineResponse20021) SetSecType(v string) {
	o.SecType = &v
}

func (o InlineResponse20021) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServerId != nil {
		toSerialize["server_id"] = o.ServerId
	}
	if o.ColumnName != nil {
		toSerialize["column_name"] = o.ColumnName
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Conidex != nil {
		toSerialize["conidex"] = o.Conidex
	}
	if o.ConId != nil {
		toSerialize["con_id"] = o.ConId
	}
	if o.AvailableChartPeriods != nil {
		toSerialize["available_chart_periods"] = o.AvailableChartPeriods
	}
	if o.CompanyName != nil {
		toSerialize["company_name"] = o.CompanyName
	}
	if o.ContractDescription1 != nil {
		toSerialize["contract_description_1"] = o.ContractDescription1
	}
	if o.ListingExchange != nil {
		toSerialize["listing_exchange"] = o.ListingExchange
	}
	if o.SecType != nil {
		toSerialize["sec_type"] = o.SecType
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20021 struct {
	value *InlineResponse20021
	isSet bool
}

func (v NullableInlineResponse20021) Get() *InlineResponse20021 {
	return v.value
}

func (v *NullableInlineResponse20021) Set(val *InlineResponse20021) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20021) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20021) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20021(val *InlineResponse20021) *NullableInlineResponse20021 {
	return &NullableInlineResponse20021{value: val, isSet: true}
}

func (v NullableInlineResponse20021) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20021) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


