# MarketData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conid** | Pointer to **float32** | IBKR Contract ID | [optional] 
**Exchange** | Pointer to **string** | Exchange | [optional] 
**MinTick** | Pointer to **float32** |  | [optional] 
**Last** | Pointer to **float32** |  | [optional] 
**LastSize** | Pointer to **float32** |  | [optional] 
**Bid** | Pointer to **float32** |  | [optional] 
**BidSize** | Pointer to **float32** |  | [optional] 
**Ask** | Pointer to **float32** |  | [optional] 
**AskSize** | Pointer to **float32** |  | [optional] 

## Methods

### NewMarketData

`func NewMarketData() *MarketData`

NewMarketData instantiates a new MarketData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketDataWithDefaults

`func NewMarketDataWithDefaults() *MarketData`

NewMarketDataWithDefaults instantiates a new MarketData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConid

`func (o *MarketData) GetConid() float32`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *MarketData) GetConidOk() (*float32, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *MarketData) SetConid(v float32)`

SetConid sets Conid field to given value.

### HasConid

`func (o *MarketData) HasConid() bool`

HasConid returns a boolean if a field has been set.

### GetExchange

`func (o *MarketData) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *MarketData) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *MarketData) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *MarketData) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetMinTick

`func (o *MarketData) GetMinTick() float32`

GetMinTick returns the MinTick field if non-nil, zero value otherwise.

### GetMinTickOk

`func (o *MarketData) GetMinTickOk() (*float32, bool)`

GetMinTickOk returns a tuple with the MinTick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTick

`func (o *MarketData) SetMinTick(v float32)`

SetMinTick sets MinTick field to given value.

### HasMinTick

`func (o *MarketData) HasMinTick() bool`

HasMinTick returns a boolean if a field has been set.

### GetLast

`func (o *MarketData) GetLast() float32`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *MarketData) GetLastOk() (*float32, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *MarketData) SetLast(v float32)`

SetLast sets Last field to given value.

### HasLast

`func (o *MarketData) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetLastSize

`func (o *MarketData) GetLastSize() float32`

GetLastSize returns the LastSize field if non-nil, zero value otherwise.

### GetLastSizeOk

`func (o *MarketData) GetLastSizeOk() (*float32, bool)`

GetLastSizeOk returns a tuple with the LastSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSize

`func (o *MarketData) SetLastSize(v float32)`

SetLastSize sets LastSize field to given value.

### HasLastSize

`func (o *MarketData) HasLastSize() bool`

HasLastSize returns a boolean if a field has been set.

### GetBid

`func (o *MarketData) GetBid() float32`

GetBid returns the Bid field if non-nil, zero value otherwise.

### GetBidOk

`func (o *MarketData) GetBidOk() (*float32, bool)`

GetBidOk returns a tuple with the Bid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBid

`func (o *MarketData) SetBid(v float32)`

SetBid sets Bid field to given value.

### HasBid

`func (o *MarketData) HasBid() bool`

HasBid returns a boolean if a field has been set.

### GetBidSize

`func (o *MarketData) GetBidSize() float32`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *MarketData) GetBidSizeOk() (*float32, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *MarketData) SetBidSize(v float32)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *MarketData) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetAsk

`func (o *MarketData) GetAsk() float32`

GetAsk returns the Ask field if non-nil, zero value otherwise.

### GetAskOk

`func (o *MarketData) GetAskOk() (*float32, bool)`

GetAskOk returns a tuple with the Ask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsk

`func (o *MarketData) SetAsk(v float32)`

SetAsk sets Ask field to given value.

### HasAsk

`func (o *MarketData) HasAsk() bool`

HasAsk returns a boolean if a field has been set.

### GetAskSize

`func (o *MarketData) GetAskSize() float32`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *MarketData) GetAskSizeOk() (*float32, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *MarketData) SetAskSize(v float32)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *MarketData) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


