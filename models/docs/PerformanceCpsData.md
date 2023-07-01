# PerformanceCpsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**IdType** | Pointer to **string** | for example-- acctid | [optional] 
**Start** | Pointer to **string** | start date-- yyyyMMdd | [optional] 
**BaseCurrency** | Pointer to **string** |  | [optional] 
**Returns** | Pointer to **[]float32** | each value stands for price change percent of corresponding date in dates array | [optional] 
**End** | Pointer to **string** | end date-- yyyyMMdd | [optional] 

## Methods

### NewPerformanceCpsData

`func NewPerformanceCpsData() *PerformanceCpsData`

NewPerformanceCpsData instantiates a new PerformanceCpsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceCpsDataWithDefaults

`func NewPerformanceCpsDataWithDefaults() *PerformanceCpsData`

NewPerformanceCpsDataWithDefaults instantiates a new PerformanceCpsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PerformanceCpsData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PerformanceCpsData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PerformanceCpsData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PerformanceCpsData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdType

`func (o *PerformanceCpsData) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *PerformanceCpsData) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *PerformanceCpsData) SetIdType(v string)`

SetIdType sets IdType field to given value.

### HasIdType

`func (o *PerformanceCpsData) HasIdType() bool`

HasIdType returns a boolean if a field has been set.

### GetStart

`func (o *PerformanceCpsData) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *PerformanceCpsData) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *PerformanceCpsData) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *PerformanceCpsData) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetBaseCurrency

`func (o *PerformanceCpsData) GetBaseCurrency() string`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *PerformanceCpsData) GetBaseCurrencyOk() (*string, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *PerformanceCpsData) SetBaseCurrency(v string)`

SetBaseCurrency sets BaseCurrency field to given value.

### HasBaseCurrency

`func (o *PerformanceCpsData) HasBaseCurrency() bool`

HasBaseCurrency returns a boolean if a field has been set.

### GetReturns

`func (o *PerformanceCpsData) GetReturns() []float32`

GetReturns returns the Returns field if non-nil, zero value otherwise.

### GetReturnsOk

`func (o *PerformanceCpsData) GetReturnsOk() (*[]float32, bool)`

GetReturnsOk returns a tuple with the Returns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturns

`func (o *PerformanceCpsData) SetReturns(v []float32)`

SetReturns sets Returns field to given value.

### HasReturns

`func (o *PerformanceCpsData) HasReturns() bool`

HasReturns returns a boolean if a field has been set.

### GetEnd

`func (o *PerformanceCpsData) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *PerformanceCpsData) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *PerformanceCpsData) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *PerformanceCpsData) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


