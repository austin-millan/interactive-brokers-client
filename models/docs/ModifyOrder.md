# ModifyOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctId** | Pointer to **string** |  | [optional] 
**Conid** | Pointer to **int32** |  | [optional] 
**OrderId** | Pointer to **int32** | customer orderid | [optional] 
**OrderType** | Pointer to **string** | for example LMT | [optional] 
**OutsideRTH** | Pointer to **bool** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**AuxPrice** | Pointer to **float32** |  | [optional] 
**Side** | Pointer to **string** | SELL or BUY | [optional] 
**ListingExchange** | Pointer to **string** | optional, not required | [optional] 
**Ticker** | Pointer to **string** |  | [optional] 
**Tif** | Pointer to **string** | for example DAY | [optional] 
**Quantity** | Pointer to **float32** | usually integer, for some special cases can be float numbers | [optional] 

## Methods

### NewModifyOrder

`func NewModifyOrder() *ModifyOrder`

NewModifyOrder instantiates a new ModifyOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyOrderWithDefaults

`func NewModifyOrderWithDefaults() *ModifyOrder`

NewModifyOrderWithDefaults instantiates a new ModifyOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctId

`func (o *ModifyOrder) GetAcctId() string`

GetAcctId returns the AcctId field if non-nil, zero value otherwise.

### GetAcctIdOk

`func (o *ModifyOrder) GetAcctIdOk() (*string, bool)`

GetAcctIdOk returns a tuple with the AcctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctId

`func (o *ModifyOrder) SetAcctId(v string)`

SetAcctId sets AcctId field to given value.

### HasAcctId

`func (o *ModifyOrder) HasAcctId() bool`

HasAcctId returns a boolean if a field has been set.

### GetConid

`func (o *ModifyOrder) GetConid() int32`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *ModifyOrder) GetConidOk() (*int32, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *ModifyOrder) SetConid(v int32)`

SetConid sets Conid field to given value.

### HasConid

`func (o *ModifyOrder) HasConid() bool`

HasConid returns a boolean if a field has been set.

### GetOrderId

`func (o *ModifyOrder) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ModifyOrder) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ModifyOrder) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ModifyOrder) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderType

`func (o *ModifyOrder) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *ModifyOrder) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *ModifyOrder) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *ModifyOrder) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetOutsideRTH

`func (o *ModifyOrder) GetOutsideRTH() bool`

GetOutsideRTH returns the OutsideRTH field if non-nil, zero value otherwise.

### GetOutsideRTHOk

`func (o *ModifyOrder) GetOutsideRTHOk() (*bool, bool)`

GetOutsideRTHOk returns a tuple with the OutsideRTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideRTH

`func (o *ModifyOrder) SetOutsideRTH(v bool)`

SetOutsideRTH sets OutsideRTH field to given value.

### HasOutsideRTH

`func (o *ModifyOrder) HasOutsideRTH() bool`

HasOutsideRTH returns a boolean if a field has been set.

### GetPrice

`func (o *ModifyOrder) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ModifyOrder) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ModifyOrder) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ModifyOrder) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetAuxPrice

`func (o *ModifyOrder) GetAuxPrice() float32`

GetAuxPrice returns the AuxPrice field if non-nil, zero value otherwise.

### GetAuxPriceOk

`func (o *ModifyOrder) GetAuxPriceOk() (*float32, bool)`

GetAuxPriceOk returns a tuple with the AuxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxPrice

`func (o *ModifyOrder) SetAuxPrice(v float32)`

SetAuxPrice sets AuxPrice field to given value.

### HasAuxPrice

`func (o *ModifyOrder) HasAuxPrice() bool`

HasAuxPrice returns a boolean if a field has been set.

### GetSide

`func (o *ModifyOrder) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *ModifyOrder) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *ModifyOrder) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *ModifyOrder) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetListingExchange

`func (o *ModifyOrder) GetListingExchange() string`

GetListingExchange returns the ListingExchange field if non-nil, zero value otherwise.

### GetListingExchangeOk

`func (o *ModifyOrder) GetListingExchangeOk() (*string, bool)`

GetListingExchangeOk returns a tuple with the ListingExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingExchange

`func (o *ModifyOrder) SetListingExchange(v string)`

SetListingExchange sets ListingExchange field to given value.

### HasListingExchange

`func (o *ModifyOrder) HasListingExchange() bool`

HasListingExchange returns a boolean if a field has been set.

### GetTicker

`func (o *ModifyOrder) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *ModifyOrder) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *ModifyOrder) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *ModifyOrder) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetTif

`func (o *ModifyOrder) GetTif() string`

GetTif returns the Tif field if non-nil, zero value otherwise.

### GetTifOk

`func (o *ModifyOrder) GetTifOk() (*string, bool)`

GetTifOk returns a tuple with the Tif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTif

`func (o *ModifyOrder) SetTif(v string)`

SetTif sets Tif field to given value.

### HasTif

`func (o *ModifyOrder) HasTif() bool`

HasTif returns a boolean if a field has been set.

### GetQuantity

`func (o *ModifyOrder) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ModifyOrder) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ModifyOrder) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ModifyOrder) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


