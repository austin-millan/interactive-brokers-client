# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acct** | Pointer to **string** | account id | [optional] 
**Conid** | Pointer to **int32** |  | [optional] 
**OrderDesc** | Pointer to **string** |  | [optional] 
**Description1** | Pointer to **string** |  | [optional] 
**Ticker** | Pointer to **string** | for exmple FB | [optional] 
**SecType** | Pointer to **string** | for example STK | [optional] 
**ListingExchange** | Pointer to **string** | for example NASDAQ.NMS | [optional] 
**RemainingQuantity** | Pointer to **string** |  | [optional] 
**FilledQuantity** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | PendingSubmit - Indicates the order was sent, but confirmation has not been received that it has been received by the destination.                  Occurs most commonly if an exchange is closed. PendingCancel - Indicates that a request has been sent to cancel an order but confirmation has not been received of its cancellation. PreSubmitted - Indicates that a simulated order type has been accepted by the IBKR system and that this order has yet to be elected.                 The order is held in the IBKR system until the election criteria are met. At that time the order is transmitted to the order destination as specified.  Submitted - Indicates that the order has been accepted at the order destination and is working. Cancelled - Indicates that the balance of the order has been confirmed cancelled by the IB system.              This could occur unexpectedly when IB or the destination has rejected the order.   Filled - Indicates that the order has been completely filled.  Inactive - Indicates the order is not working, for instance if the order was invalid and triggered an error message,            or if the order was to short a security and shares have not yet been located.   | [optional] 
**OrigOrderType** | Pointer to **string** | for example Limit | [optional] 
**Side** | Pointer to **string** | BUY or SELL | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**BgColor** | Pointer to **string** | back-ground color | [optional] 
**FgColor** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int32** |  | [optional] 
**ParentId** | Pointer to **string** | Only exists in child order of bracket | [optional] 
**OrderRef** | Pointer to **string** | User defined string used to identify the order. Value is set using \&quot;cOID\&quot; field while placing an order. | [optional] 

## Methods

### NewOrder

`func NewOrder() *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcct

`func (o *Order) GetAcct() string`

GetAcct returns the Acct field if non-nil, zero value otherwise.

### GetAcctOk

`func (o *Order) GetAcctOk() (*string, bool)`

GetAcctOk returns a tuple with the Acct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcct

`func (o *Order) SetAcct(v string)`

SetAcct sets Acct field to given value.

### HasAcct

`func (o *Order) HasAcct() bool`

HasAcct returns a boolean if a field has been set.

### GetConid

`func (o *Order) GetConid() int32`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *Order) GetConidOk() (*int32, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *Order) SetConid(v int32)`

SetConid sets Conid field to given value.

### HasConid

`func (o *Order) HasConid() bool`

HasConid returns a boolean if a field has been set.

### GetOrderDesc

`func (o *Order) GetOrderDesc() string`

GetOrderDesc returns the OrderDesc field if non-nil, zero value otherwise.

### GetOrderDescOk

`func (o *Order) GetOrderDescOk() (*string, bool)`

GetOrderDescOk returns a tuple with the OrderDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDesc

`func (o *Order) SetOrderDesc(v string)`

SetOrderDesc sets OrderDesc field to given value.

### HasOrderDesc

`func (o *Order) HasOrderDesc() bool`

HasOrderDesc returns a boolean if a field has been set.

### GetDescription1

`func (o *Order) GetDescription1() string`

GetDescription1 returns the Description1 field if non-nil, zero value otherwise.

### GetDescription1Ok

`func (o *Order) GetDescription1Ok() (*string, bool)`

GetDescription1Ok returns a tuple with the Description1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription1

`func (o *Order) SetDescription1(v string)`

SetDescription1 sets Description1 field to given value.

### HasDescription1

`func (o *Order) HasDescription1() bool`

HasDescription1 returns a boolean if a field has been set.

### GetTicker

`func (o *Order) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *Order) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *Order) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *Order) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetSecType

`func (o *Order) GetSecType() string`

GetSecType returns the SecType field if non-nil, zero value otherwise.

### GetSecTypeOk

`func (o *Order) GetSecTypeOk() (*string, bool)`

GetSecTypeOk returns a tuple with the SecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecType

`func (o *Order) SetSecType(v string)`

SetSecType sets SecType field to given value.

### HasSecType

`func (o *Order) HasSecType() bool`

HasSecType returns a boolean if a field has been set.

### GetListingExchange

`func (o *Order) GetListingExchange() string`

GetListingExchange returns the ListingExchange field if non-nil, zero value otherwise.

### GetListingExchangeOk

`func (o *Order) GetListingExchangeOk() (*string, bool)`

GetListingExchangeOk returns a tuple with the ListingExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingExchange

`func (o *Order) SetListingExchange(v string)`

SetListingExchange sets ListingExchange field to given value.

### HasListingExchange

`func (o *Order) HasListingExchange() bool`

HasListingExchange returns a boolean if a field has been set.

### GetRemainingQuantity

`func (o *Order) GetRemainingQuantity() string`

GetRemainingQuantity returns the RemainingQuantity field if non-nil, zero value otherwise.

### GetRemainingQuantityOk

`func (o *Order) GetRemainingQuantityOk() (*string, bool)`

GetRemainingQuantityOk returns a tuple with the RemainingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingQuantity

`func (o *Order) SetRemainingQuantity(v string)`

SetRemainingQuantity sets RemainingQuantity field to given value.

### HasRemainingQuantity

`func (o *Order) HasRemainingQuantity() bool`

HasRemainingQuantity returns a boolean if a field has been set.

### GetFilledQuantity

`func (o *Order) GetFilledQuantity() string`

GetFilledQuantity returns the FilledQuantity field if non-nil, zero value otherwise.

### GetFilledQuantityOk

`func (o *Order) GetFilledQuantityOk() (*string, bool)`

GetFilledQuantityOk returns a tuple with the FilledQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledQuantity

`func (o *Order) SetFilledQuantity(v string)`

SetFilledQuantity sets FilledQuantity field to given value.

### HasFilledQuantity

`func (o *Order) HasFilledQuantity() bool`

HasFilledQuantity returns a boolean if a field has been set.

### GetCompanyName

`func (o *Order) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Order) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *Order) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *Order) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetStatus

`func (o *Order) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Order) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Order) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Order) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOrigOrderType

`func (o *Order) GetOrigOrderType() string`

GetOrigOrderType returns the OrigOrderType field if non-nil, zero value otherwise.

### GetOrigOrderTypeOk

`func (o *Order) GetOrigOrderTypeOk() (*string, bool)`

GetOrigOrderTypeOk returns a tuple with the OrigOrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigOrderType

`func (o *Order) SetOrigOrderType(v string)`

SetOrigOrderType sets OrigOrderType field to given value.

### HasOrigOrderType

`func (o *Order) HasOrigOrderType() bool`

HasOrigOrderType returns a boolean if a field has been set.

### GetSide

`func (o *Order) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *Order) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *Order) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *Order) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetPrice

`func (o *Order) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Order) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Order) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Order) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetBgColor

`func (o *Order) GetBgColor() string`

GetBgColor returns the BgColor field if non-nil, zero value otherwise.

### GetBgColorOk

`func (o *Order) GetBgColorOk() (*string, bool)`

GetBgColorOk returns a tuple with the BgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgColor

`func (o *Order) SetBgColor(v string)`

SetBgColor sets BgColor field to given value.

### HasBgColor

`func (o *Order) HasBgColor() bool`

HasBgColor returns a boolean if a field has been set.

### GetFgColor

`func (o *Order) GetFgColor() string`

GetFgColor returns the FgColor field if non-nil, zero value otherwise.

### GetFgColorOk

`func (o *Order) GetFgColorOk() (*string, bool)`

GetFgColorOk returns a tuple with the FgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFgColor

`func (o *Order) SetFgColor(v string)`

SetFgColor sets FgColor field to given value.

### HasFgColor

`func (o *Order) HasFgColor() bool`

HasFgColor returns a boolean if a field has been set.

### GetOrderId

`func (o *Order) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Order) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Order) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *Order) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetParentId

`func (o *Order) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Order) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Order) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Order) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetOrderRef

`func (o *Order) GetOrderRef() string`

GetOrderRef returns the OrderRef field if non-nil, zero value otherwise.

### GetOrderRefOk

`func (o *Order) GetOrderRefOk() (*string, bool)`

GetOrderRefOk returns a tuple with the OrderRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderRef

`func (o *Order) SetOrderRef(v string)`

SetOrderRef sets OrderRef field to given value.

### HasOrderRef

`func (o *Order) HasOrderRef() bool`

HasOrderRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


