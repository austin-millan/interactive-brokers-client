# StocksContracts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conid** | Pointer to **int32** | conid of the stock contract | [optional] 
**Exchange** | Pointer to **string** |  | [optional] 

## Methods

### NewStocksContracts

`func NewStocksContracts() *StocksContracts`

NewStocksContracts instantiates a new StocksContracts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStocksContractsWithDefaults

`func NewStocksContractsWithDefaults() *StocksContracts`

NewStocksContractsWithDefaults instantiates a new StocksContracts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConid

`func (o *StocksContracts) GetConid() int32`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *StocksContracts) GetConidOk() (*int32, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *StocksContracts) SetConid(v int32)`

SetConid sets Conid field to given value.

### HasConid

`func (o *StocksContracts) HasConid() bool`

HasConid returns a boolean if a field has been set.

### GetExchange

`func (o *StocksContracts) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *StocksContracts) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *StocksContracts) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *StocksContracts) HasExchange() bool`

HasExchange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


