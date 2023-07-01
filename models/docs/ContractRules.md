# ContractRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderTypes** | Pointer to **[]string** |  | [optional] 
**OrderTypesOutside** | Pointer to **[]string** |  | [optional] 
**DefaultSize** | Pointer to **float32** | default quantity you can use to place an order | [optional] 
**SizeIncrement** | Pointer to **float32** |  | [optional] 
**TifTypes** | Pointer to **[]string** |  | [optional] 
**LimitPrice** | Pointer to **float32** | default limit price you can use to prefill your order | [optional] 
**Stopprice** | Pointer to **float32** | default stop price you can use to prefill your order | [optional] 
**Preview** | Pointer to **bool** | if you can preview the order or not with the whatif end-point | [optional] 
**DisplaySize** | Pointer to **string** |  | [optional] 
**Increment** | Pointer to **string** |  | [optional] 

## Methods

### NewContractRules

`func NewContractRules() *ContractRules`

NewContractRules instantiates a new ContractRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractRulesWithDefaults

`func NewContractRulesWithDefaults() *ContractRules`

NewContractRulesWithDefaults instantiates a new ContractRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderTypes

`func (o *ContractRules) GetOrderTypes() []string`

GetOrderTypes returns the OrderTypes field if non-nil, zero value otherwise.

### GetOrderTypesOk

`func (o *ContractRules) GetOrderTypesOk() (*[]string, bool)`

GetOrderTypesOk returns a tuple with the OrderTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTypes

`func (o *ContractRules) SetOrderTypes(v []string)`

SetOrderTypes sets OrderTypes field to given value.

### HasOrderTypes

`func (o *ContractRules) HasOrderTypes() bool`

HasOrderTypes returns a boolean if a field has been set.

### GetOrderTypesOutside

`func (o *ContractRules) GetOrderTypesOutside() []string`

GetOrderTypesOutside returns the OrderTypesOutside field if non-nil, zero value otherwise.

### GetOrderTypesOutsideOk

`func (o *ContractRules) GetOrderTypesOutsideOk() (*[]string, bool)`

GetOrderTypesOutsideOk returns a tuple with the OrderTypesOutside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTypesOutside

`func (o *ContractRules) SetOrderTypesOutside(v []string)`

SetOrderTypesOutside sets OrderTypesOutside field to given value.

### HasOrderTypesOutside

`func (o *ContractRules) HasOrderTypesOutside() bool`

HasOrderTypesOutside returns a boolean if a field has been set.

### GetDefaultSize

`func (o *ContractRules) GetDefaultSize() float32`

GetDefaultSize returns the DefaultSize field if non-nil, zero value otherwise.

### GetDefaultSizeOk

`func (o *ContractRules) GetDefaultSizeOk() (*float32, bool)`

GetDefaultSizeOk returns a tuple with the DefaultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSize

`func (o *ContractRules) SetDefaultSize(v float32)`

SetDefaultSize sets DefaultSize field to given value.

### HasDefaultSize

`func (o *ContractRules) HasDefaultSize() bool`

HasDefaultSize returns a boolean if a field has been set.

### GetSizeIncrement

`func (o *ContractRules) GetSizeIncrement() float32`

GetSizeIncrement returns the SizeIncrement field if non-nil, zero value otherwise.

### GetSizeIncrementOk

`func (o *ContractRules) GetSizeIncrementOk() (*float32, bool)`

GetSizeIncrementOk returns a tuple with the SizeIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeIncrement

`func (o *ContractRules) SetSizeIncrement(v float32)`

SetSizeIncrement sets SizeIncrement field to given value.

### HasSizeIncrement

`func (o *ContractRules) HasSizeIncrement() bool`

HasSizeIncrement returns a boolean if a field has been set.

### GetTifTypes

`func (o *ContractRules) GetTifTypes() []string`

GetTifTypes returns the TifTypes field if non-nil, zero value otherwise.

### GetTifTypesOk

`func (o *ContractRules) GetTifTypesOk() (*[]string, bool)`

GetTifTypesOk returns a tuple with the TifTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTifTypes

`func (o *ContractRules) SetTifTypes(v []string)`

SetTifTypes sets TifTypes field to given value.

### HasTifTypes

`func (o *ContractRules) HasTifTypes() bool`

HasTifTypes returns a boolean if a field has been set.

### GetLimitPrice

`func (o *ContractRules) GetLimitPrice() float32`

GetLimitPrice returns the LimitPrice field if non-nil, zero value otherwise.

### GetLimitPriceOk

`func (o *ContractRules) GetLimitPriceOk() (*float32, bool)`

GetLimitPriceOk returns a tuple with the LimitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPrice

`func (o *ContractRules) SetLimitPrice(v float32)`

SetLimitPrice sets LimitPrice field to given value.

### HasLimitPrice

`func (o *ContractRules) HasLimitPrice() bool`

HasLimitPrice returns a boolean if a field has been set.

### GetStopprice

`func (o *ContractRules) GetStopprice() float32`

GetStopprice returns the Stopprice field if non-nil, zero value otherwise.

### GetStoppriceOk

`func (o *ContractRules) GetStoppriceOk() (*float32, bool)`

GetStoppriceOk returns a tuple with the Stopprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopprice

`func (o *ContractRules) SetStopprice(v float32)`

SetStopprice sets Stopprice field to given value.

### HasStopprice

`func (o *ContractRules) HasStopprice() bool`

HasStopprice returns a boolean if a field has been set.

### GetPreview

`func (o *ContractRules) GetPreview() bool`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *ContractRules) GetPreviewOk() (*bool, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *ContractRules) SetPreview(v bool)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *ContractRules) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetDisplaySize

`func (o *ContractRules) GetDisplaySize() string`

GetDisplaySize returns the DisplaySize field if non-nil, zero value otherwise.

### GetDisplaySizeOk

`func (o *ContractRules) GetDisplaySizeOk() (*string, bool)`

GetDisplaySizeOk returns a tuple with the DisplaySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySize

`func (o *ContractRules) SetDisplaySize(v string)`

SetDisplaySize sets DisplaySize field to given value.

### HasDisplaySize

`func (o *ContractRules) HasDisplaySize() bool`

HasDisplaySize returns a boolean if a field has been set.

### GetIncrement

`func (o *ContractRules) GetIncrement() string`

GetIncrement returns the Increment field if non-nil, zero value otherwise.

### GetIncrementOk

`func (o *ContractRules) GetIncrementOk() (*string, bool)`

GetIncrementOk returns a tuple with the Increment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrement

`func (o *ContractRules) SetIncrement(v string)`

SetIncrement sets Increment field to given value.

### HasIncrement

`func (o *ContractRules) HasIncrement() bool`

HasIncrement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


