# OrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctId** | Pointer to **string** | acctId is optional. It should be one of the accounts returned by /iserver/accounts. If not passed, the first one in the list is selected.  | [optional] 
**Conid** | Pointer to **int32** | conid is the identifier of the security you want to trade, you can find the conid with /iserver/secdef/search.  | [optional] 
**SecType** | Pointer to **string** | conid:type for example 265598:STK | [optional] 
**COID** | Pointer to **string** | Customer Order ID. An arbitraty string that can be used to identify the order, e.g \&quot;my-fb-order\&quot;. The value must be unique for a 24h span. Please do not set this value for child orders when placing a bracket order.  | [optional] 
**ParentId** | Pointer to **string** | When placing bracket orders, the child parentId must be equal to the cOId (customer order id) of the parent.  | [optional] 
**OrderType** | Pointer to **string** | orderType can be one of MKT (Market), LMT (Limit), STP (Stop) or STP_LIMIT (stop limit)  | [optional] 
**ListingExchange** | Pointer to **string** | listingExchange is optional. By default we use \&quot;SMART\&quot; routing. Possible values are available via this end point: /v1/portal/iserver/contract/{{conid}}/info, see valid_exchange: e.g: SMART,AMEX,NYSE, CBOE,ISE,CHX,ARCA,ISLAND,DRCTEDGE,BEX,BATS,EDGEA,CSFBALGO,JE FFALGO,BYX,IEX,FOXRIVER,TPLUS1,NYSENAT,PSX  | [optional] 
**OutsideRTH** | Pointer to **bool** | set to true if the order can be executed outside regular trading hours.  | [optional] 
**Price** | Pointer to **float32** | optional if order is MKT, for LMT, this is the limit price. For STP this is the stop price.  | [optional] 
**Side** | Pointer to **string** | SELL or BUY | [optional] 
**Ticker** | Pointer to **string** |  | [optional] 
**Tif** | Pointer to **string** | GTC (Good Till Cancel) or DAY. DAY orders are automatically cancelled at the end of the Day or Trading hours.  | [optional] 
**Referrer** | Pointer to **string** | for example QuickTrade | [optional] 
**Quantity** | Pointer to **float32** | usually integer, for some special cases can be float numbers | [optional] 
**UseAdaptive** | Pointer to **bool** | If true, the system will use the Adaptive Algo to submit the order https://www.interactivebrokers.com/en/index.php?f&#x3D;19091  | [optional] 

## Methods

### NewOrderRequest

`func NewOrderRequest() *OrderRequest`

NewOrderRequest instantiates a new OrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderRequestWithDefaults

`func NewOrderRequestWithDefaults() *OrderRequest`

NewOrderRequestWithDefaults instantiates a new OrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctId

`func (o *OrderRequest) GetAcctId() string`

GetAcctId returns the AcctId field if non-nil, zero value otherwise.

### GetAcctIdOk

`func (o *OrderRequest) GetAcctIdOk() (*string, bool)`

GetAcctIdOk returns a tuple with the AcctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctId

`func (o *OrderRequest) SetAcctId(v string)`

SetAcctId sets AcctId field to given value.

### HasAcctId

`func (o *OrderRequest) HasAcctId() bool`

HasAcctId returns a boolean if a field has been set.

### GetConid

`func (o *OrderRequest) GetConid() int32`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *OrderRequest) GetConidOk() (*int32, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *OrderRequest) SetConid(v int32)`

SetConid sets Conid field to given value.

### HasConid

`func (o *OrderRequest) HasConid() bool`

HasConid returns a boolean if a field has been set.

### GetSecType

`func (o *OrderRequest) GetSecType() string`

GetSecType returns the SecType field if non-nil, zero value otherwise.

### GetSecTypeOk

`func (o *OrderRequest) GetSecTypeOk() (*string, bool)`

GetSecTypeOk returns a tuple with the SecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecType

`func (o *OrderRequest) SetSecType(v string)`

SetSecType sets SecType field to given value.

### HasSecType

`func (o *OrderRequest) HasSecType() bool`

HasSecType returns a boolean if a field has been set.

### GetCOID

`func (o *OrderRequest) GetCOID() string`

GetCOID returns the COID field if non-nil, zero value otherwise.

### GetCOIDOk

`func (o *OrderRequest) GetCOIDOk() (*string, bool)`

GetCOIDOk returns a tuple with the COID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCOID

`func (o *OrderRequest) SetCOID(v string)`

SetCOID sets COID field to given value.

### HasCOID

`func (o *OrderRequest) HasCOID() bool`

HasCOID returns a boolean if a field has been set.

### GetParentId

`func (o *OrderRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *OrderRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *OrderRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *OrderRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetOrderType

`func (o *OrderRequest) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *OrderRequest) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *OrderRequest) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *OrderRequest) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetListingExchange

`func (o *OrderRequest) GetListingExchange() string`

GetListingExchange returns the ListingExchange field if non-nil, zero value otherwise.

### GetListingExchangeOk

`func (o *OrderRequest) GetListingExchangeOk() (*string, bool)`

GetListingExchangeOk returns a tuple with the ListingExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingExchange

`func (o *OrderRequest) SetListingExchange(v string)`

SetListingExchange sets ListingExchange field to given value.

### HasListingExchange

`func (o *OrderRequest) HasListingExchange() bool`

HasListingExchange returns a boolean if a field has been set.

### GetOutsideRTH

`func (o *OrderRequest) GetOutsideRTH() bool`

GetOutsideRTH returns the OutsideRTH field if non-nil, zero value otherwise.

### GetOutsideRTHOk

`func (o *OrderRequest) GetOutsideRTHOk() (*bool, bool)`

GetOutsideRTHOk returns a tuple with the OutsideRTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideRTH

`func (o *OrderRequest) SetOutsideRTH(v bool)`

SetOutsideRTH sets OutsideRTH field to given value.

### HasOutsideRTH

`func (o *OrderRequest) HasOutsideRTH() bool`

HasOutsideRTH returns a boolean if a field has been set.

### GetPrice

`func (o *OrderRequest) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderRequest) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderRequest) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderRequest) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSide

`func (o *OrderRequest) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OrderRequest) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OrderRequest) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *OrderRequest) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetTicker

`func (o *OrderRequest) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *OrderRequest) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *OrderRequest) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *OrderRequest) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetTif

`func (o *OrderRequest) GetTif() string`

GetTif returns the Tif field if non-nil, zero value otherwise.

### GetTifOk

`func (o *OrderRequest) GetTifOk() (*string, bool)`

GetTifOk returns a tuple with the Tif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTif

`func (o *OrderRequest) SetTif(v string)`

SetTif sets Tif field to given value.

### HasTif

`func (o *OrderRequest) HasTif() bool`

HasTif returns a boolean if a field has been set.

### GetReferrer

`func (o *OrderRequest) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *OrderRequest) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *OrderRequest) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *OrderRequest) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderRequest) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderRequest) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderRequest) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUseAdaptive

`func (o *OrderRequest) GetUseAdaptive() bool`

GetUseAdaptive returns the UseAdaptive field if non-nil, zero value otherwise.

### GetUseAdaptiveOk

`func (o *OrderRequest) GetUseAdaptiveOk() (*bool, bool)`

GetUseAdaptiveOk returns a tuple with the UseAdaptive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAdaptive

`func (o *OrderRequest) SetUseAdaptive(v bool)`

SetUseAdaptive sets UseAdaptive field to given value.

### HasUseAdaptive

`func (o *OrderRequest) HasUseAdaptive() bool`

HasUseAdaptive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


