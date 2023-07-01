# Contract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RTH** | Pointer to **bool** | true means you can trade outside RTH(regular trading hours) | [optional] 
**ConId** | Pointer to **string** | same as that in request | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**Exchange** | Pointer to **string** |  | [optional] 
**LocalSymbol** | Pointer to **string** | for exmple FB | [optional] 
**InstrumentType** | Pointer to **string** | for example STK | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Industry** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**ContractRules**](contract_rules.md) |  | [optional] 

## Methods

### NewContract

`func NewContract() *Contract`

NewContract instantiates a new Contract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractWithDefaults

`func NewContractWithDefaults() *Contract`

NewContractWithDefaults instantiates a new Contract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRTH

`func (o *Contract) GetRTH() bool`

GetRTH returns the RTH field if non-nil, zero value otherwise.

### GetRTHOk

`func (o *Contract) GetRTHOk() (*bool, bool)`

GetRTHOk returns a tuple with the RTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRTH

`func (o *Contract) SetRTH(v bool)`

SetRTH sets RTH field to given value.

### HasRTH

`func (o *Contract) HasRTH() bool`

HasRTH returns a boolean if a field has been set.

### GetConId

`func (o *Contract) GetConId() string`

GetConId returns the ConId field if non-nil, zero value otherwise.

### GetConIdOk

`func (o *Contract) GetConIdOk() (*string, bool)`

GetConIdOk returns a tuple with the ConId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConId

`func (o *Contract) SetConId(v string)`

SetConId sets ConId field to given value.

### HasConId

`func (o *Contract) HasConId() bool`

HasConId returns a boolean if a field has been set.

### GetCompanyName

`func (o *Contract) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Contract) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *Contract) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *Contract) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetExchange

`func (o *Contract) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *Contract) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *Contract) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *Contract) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetLocalSymbol

`func (o *Contract) GetLocalSymbol() string`

GetLocalSymbol returns the LocalSymbol field if non-nil, zero value otherwise.

### GetLocalSymbolOk

`func (o *Contract) GetLocalSymbolOk() (*string, bool)`

GetLocalSymbolOk returns a tuple with the LocalSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSymbol

`func (o *Contract) SetLocalSymbol(v string)`

SetLocalSymbol sets LocalSymbol field to given value.

### HasLocalSymbol

`func (o *Contract) HasLocalSymbol() bool`

HasLocalSymbol returns a boolean if a field has been set.

### GetInstrumentType

`func (o *Contract) GetInstrumentType() string`

GetInstrumentType returns the InstrumentType field if non-nil, zero value otherwise.

### GetInstrumentTypeOk

`func (o *Contract) GetInstrumentTypeOk() (*string, bool)`

GetInstrumentTypeOk returns a tuple with the InstrumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentType

`func (o *Contract) SetInstrumentType(v string)`

SetInstrumentType sets InstrumentType field to given value.

### HasInstrumentType

`func (o *Contract) HasInstrumentType() bool`

HasInstrumentType returns a boolean if a field has been set.

### GetCurrency

`func (o *Contract) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Contract) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Contract) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Contract) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCategory

`func (o *Contract) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Contract) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Contract) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Contract) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetIndustry

`func (o *Contract) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *Contract) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *Contract) SetIndustry(v string)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *Contract) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetRules

`func (o *Contract) GetRules() ContractRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *Contract) GetRulesOk() (*ContractRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *Contract) SetRules(v ContractRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *Contract) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


