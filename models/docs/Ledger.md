# Ledger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commoditymarketvalue** | Pointer to **float32** |  | [optional] 
**Futuremarketvalue** | Pointer to **float32** |  | [optional] 
**Settledcash** | Pointer to **float32** |  | [optional] 
**Exchangerate** | Pointer to **float32** |  | [optional] 
**Sessionid** | Pointer to **int32** |  | [optional] 
**Cashbalance** | Pointer to **float32** |  | [optional] 
**Corporatebondsmarketvalue** | Pointer to **float32** |  | [optional] 
**Warrantsmarketvalue** | Pointer to **float32** |  | [optional] 
**Netliquidationvalue** | Pointer to **float32** |  | [optional] 
**Interest** | Pointer to **float32** |  | [optional] 
**Unrealizedpnl** | Pointer to **float32** |  | [optional] 
**Stockmarketvalue** | Pointer to **float32** |  | [optional] 
**Moneyfunds** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Realizedpnl** | Pointer to **float32** |  | [optional] 
**Funds** | Pointer to **float32** |  | [optional] 
**Acctcode** | Pointer to **string** |  | [optional] 
**Issueroptionsmarketvalue** | Pointer to **float32** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int32** |  | [optional] 
**Severity** | Pointer to **int32** |  | [optional] 

## Methods

### NewLedger

`func NewLedger() *Ledger`

NewLedger instantiates a new Ledger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLedgerWithDefaults

`func NewLedgerWithDefaults() *Ledger`

NewLedgerWithDefaults instantiates a new Ledger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommoditymarketvalue

`func (o *Ledger) GetCommoditymarketvalue() float32`

GetCommoditymarketvalue returns the Commoditymarketvalue field if non-nil, zero value otherwise.

### GetCommoditymarketvalueOk

`func (o *Ledger) GetCommoditymarketvalueOk() (*float32, bool)`

GetCommoditymarketvalueOk returns a tuple with the Commoditymarketvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommoditymarketvalue

`func (o *Ledger) SetCommoditymarketvalue(v float32)`

SetCommoditymarketvalue sets Commoditymarketvalue field to given value.

### HasCommoditymarketvalue

`func (o *Ledger) HasCommoditymarketvalue() bool`

HasCommoditymarketvalue returns a boolean if a field has been set.

### GetFuturemarketvalue

`func (o *Ledger) GetFuturemarketvalue() float32`

GetFuturemarketvalue returns the Futuremarketvalue field if non-nil, zero value otherwise.

### GetFuturemarketvalueOk

`func (o *Ledger) GetFuturemarketvalueOk() (*float32, bool)`

GetFuturemarketvalueOk returns a tuple with the Futuremarketvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuturemarketvalue

`func (o *Ledger) SetFuturemarketvalue(v float32)`

SetFuturemarketvalue sets Futuremarketvalue field to given value.

### HasFuturemarketvalue

`func (o *Ledger) HasFuturemarketvalue() bool`

HasFuturemarketvalue returns a boolean if a field has been set.

### GetSettledcash

`func (o *Ledger) GetSettledcash() float32`

GetSettledcash returns the Settledcash field if non-nil, zero value otherwise.

### GetSettledcashOk

`func (o *Ledger) GetSettledcashOk() (*float32, bool)`

GetSettledcashOk returns a tuple with the Settledcash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledcash

`func (o *Ledger) SetSettledcash(v float32)`

SetSettledcash sets Settledcash field to given value.

### HasSettledcash

`func (o *Ledger) HasSettledcash() bool`

HasSettledcash returns a boolean if a field has been set.

### GetExchangerate

`func (o *Ledger) GetExchangerate() float32`

GetExchangerate returns the Exchangerate field if non-nil, zero value otherwise.

### GetExchangerateOk

`func (o *Ledger) GetExchangerateOk() (*float32, bool)`

GetExchangerateOk returns a tuple with the Exchangerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangerate

`func (o *Ledger) SetExchangerate(v float32)`

SetExchangerate sets Exchangerate field to given value.

### HasExchangerate

`func (o *Ledger) HasExchangerate() bool`

HasExchangerate returns a boolean if a field has been set.

### GetSessionid

`func (o *Ledger) GetSessionid() int32`

GetSessionid returns the Sessionid field if non-nil, zero value otherwise.

### GetSessionidOk

`func (o *Ledger) GetSessionidOk() (*int32, bool)`

GetSessionidOk returns a tuple with the Sessionid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionid

`func (o *Ledger) SetSessionid(v int32)`

SetSessionid sets Sessionid field to given value.

### HasSessionid

`func (o *Ledger) HasSessionid() bool`

HasSessionid returns a boolean if a field has been set.

### GetCashbalance

`func (o *Ledger) GetCashbalance() float32`

GetCashbalance returns the Cashbalance field if non-nil, zero value otherwise.

### GetCashbalanceOk

`func (o *Ledger) GetCashbalanceOk() (*float32, bool)`

GetCashbalanceOk returns a tuple with the Cashbalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashbalance

`func (o *Ledger) SetCashbalance(v float32)`

SetCashbalance sets Cashbalance field to given value.

### HasCashbalance

`func (o *Ledger) HasCashbalance() bool`

HasCashbalance returns a boolean if a field has been set.

### GetCorporatebondsmarketvalue

`func (o *Ledger) GetCorporatebondsmarketvalue() float32`

GetCorporatebondsmarketvalue returns the Corporatebondsmarketvalue field if non-nil, zero value otherwise.

### GetCorporatebondsmarketvalueOk

`func (o *Ledger) GetCorporatebondsmarketvalueOk() (*float32, bool)`

GetCorporatebondsmarketvalueOk returns a tuple with the Corporatebondsmarketvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporatebondsmarketvalue

`func (o *Ledger) SetCorporatebondsmarketvalue(v float32)`

SetCorporatebondsmarketvalue sets Corporatebondsmarketvalue field to given value.

### HasCorporatebondsmarketvalue

`func (o *Ledger) HasCorporatebondsmarketvalue() bool`

HasCorporatebondsmarketvalue returns a boolean if a field has been set.

### GetWarrantsmarketvalue

`func (o *Ledger) GetWarrantsmarketvalue() float32`

GetWarrantsmarketvalue returns the Warrantsmarketvalue field if non-nil, zero value otherwise.

### GetWarrantsmarketvalueOk

`func (o *Ledger) GetWarrantsmarketvalueOk() (*float32, bool)`

GetWarrantsmarketvalueOk returns a tuple with the Warrantsmarketvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantsmarketvalue

`func (o *Ledger) SetWarrantsmarketvalue(v float32)`

SetWarrantsmarketvalue sets Warrantsmarketvalue field to given value.

### HasWarrantsmarketvalue

`func (o *Ledger) HasWarrantsmarketvalue() bool`

HasWarrantsmarketvalue returns a boolean if a field has been set.

### GetNetliquidationvalue

`func (o *Ledger) GetNetliquidationvalue() float32`

GetNetliquidationvalue returns the Netliquidationvalue field if non-nil, zero value otherwise.

### GetNetliquidationvalueOk

`func (o *Ledger) GetNetliquidationvalueOk() (*float32, bool)`

GetNetliquidationvalueOk returns a tuple with the Netliquidationvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetliquidationvalue

`func (o *Ledger) SetNetliquidationvalue(v float32)`

SetNetliquidationvalue sets Netliquidationvalue field to given value.

### HasNetliquidationvalue

`func (o *Ledger) HasNetliquidationvalue() bool`

HasNetliquidationvalue returns a boolean if a field has been set.

### GetInterest

`func (o *Ledger) GetInterest() float32`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *Ledger) GetInterestOk() (*float32, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *Ledger) SetInterest(v float32)`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *Ledger) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### GetUnrealizedpnl

`func (o *Ledger) GetUnrealizedpnl() float32`

GetUnrealizedpnl returns the Unrealizedpnl field if non-nil, zero value otherwise.

### GetUnrealizedpnlOk

`func (o *Ledger) GetUnrealizedpnlOk() (*float32, bool)`

GetUnrealizedpnlOk returns a tuple with the Unrealizedpnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedpnl

`func (o *Ledger) SetUnrealizedpnl(v float32)`

SetUnrealizedpnl sets Unrealizedpnl field to given value.

### HasUnrealizedpnl

`func (o *Ledger) HasUnrealizedpnl() bool`

HasUnrealizedpnl returns a boolean if a field has been set.

### GetStockmarketvalue

`func (o *Ledger) GetStockmarketvalue() float32`

GetStockmarketvalue returns the Stockmarketvalue field if non-nil, zero value otherwise.

### GetStockmarketvalueOk

`func (o *Ledger) GetStockmarketvalueOk() (*float32, bool)`

GetStockmarketvalueOk returns a tuple with the Stockmarketvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockmarketvalue

`func (o *Ledger) SetStockmarketvalue(v float32)`

SetStockmarketvalue sets Stockmarketvalue field to given value.

### HasStockmarketvalue

`func (o *Ledger) HasStockmarketvalue() bool`

HasStockmarketvalue returns a boolean if a field has been set.

### GetMoneyfunds

`func (o *Ledger) GetMoneyfunds() float32`

GetMoneyfunds returns the Moneyfunds field if non-nil, zero value otherwise.

### GetMoneyfundsOk

`func (o *Ledger) GetMoneyfundsOk() (*float32, bool)`

GetMoneyfundsOk returns a tuple with the Moneyfunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyfunds

`func (o *Ledger) SetMoneyfunds(v float32)`

SetMoneyfunds sets Moneyfunds field to given value.

### HasMoneyfunds

`func (o *Ledger) HasMoneyfunds() bool`

HasMoneyfunds returns a boolean if a field has been set.

### GetCurrency

`func (o *Ledger) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Ledger) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Ledger) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Ledger) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetRealizedpnl

`func (o *Ledger) GetRealizedpnl() float32`

GetRealizedpnl returns the Realizedpnl field if non-nil, zero value otherwise.

### GetRealizedpnlOk

`func (o *Ledger) GetRealizedpnlOk() (*float32, bool)`

GetRealizedpnlOk returns a tuple with the Realizedpnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedpnl

`func (o *Ledger) SetRealizedpnl(v float32)`

SetRealizedpnl sets Realizedpnl field to given value.

### HasRealizedpnl

`func (o *Ledger) HasRealizedpnl() bool`

HasRealizedpnl returns a boolean if a field has been set.

### GetFunds

`func (o *Ledger) GetFunds() float32`

GetFunds returns the Funds field if non-nil, zero value otherwise.

### GetFundsOk

`func (o *Ledger) GetFundsOk() (*float32, bool)`

GetFundsOk returns a tuple with the Funds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunds

`func (o *Ledger) SetFunds(v float32)`

SetFunds sets Funds field to given value.

### HasFunds

`func (o *Ledger) HasFunds() bool`

HasFunds returns a boolean if a field has been set.

### GetAcctcode

`func (o *Ledger) GetAcctcode() string`

GetAcctcode returns the Acctcode field if non-nil, zero value otherwise.

### GetAcctcodeOk

`func (o *Ledger) GetAcctcodeOk() (*string, bool)`

GetAcctcodeOk returns a tuple with the Acctcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctcode

`func (o *Ledger) SetAcctcode(v string)`

SetAcctcode sets Acctcode field to given value.

### HasAcctcode

`func (o *Ledger) HasAcctcode() bool`

HasAcctcode returns a boolean if a field has been set.

### GetIssueroptionsmarketvalue

`func (o *Ledger) GetIssueroptionsmarketvalue() float32`

GetIssueroptionsmarketvalue returns the Issueroptionsmarketvalue field if non-nil, zero value otherwise.

### GetIssueroptionsmarketvalueOk

`func (o *Ledger) GetIssueroptionsmarketvalueOk() (*float32, bool)`

GetIssueroptionsmarketvalueOk returns a tuple with the Issueroptionsmarketvalue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueroptionsmarketvalue

`func (o *Ledger) SetIssueroptionsmarketvalue(v float32)`

SetIssueroptionsmarketvalue sets Issueroptionsmarketvalue field to given value.

### HasIssueroptionsmarketvalue

`func (o *Ledger) HasIssueroptionsmarketvalue() bool`

HasIssueroptionsmarketvalue returns a boolean if a field has been set.

### GetKey

`func (o *Ledger) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Ledger) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Ledger) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Ledger) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetTimestamp

`func (o *Ledger) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Ledger) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Ledger) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Ledger) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetSeverity

`func (o *Ledger) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Ledger) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Ledger) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Ledger) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


