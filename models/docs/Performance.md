# Performance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Cps** | Pointer to [**PerformanceCps**](performance_cps.md) |  | [optional] 
**Tpps** | Pointer to [**PerformanceTpps**](performance_tpps.md) |  | [optional] 
**Nav** | Pointer to [**PerformanceNav**](performance_nav.md) |  | [optional] 
**Pm** | Pointer to **string** |  | [optional] 
**Included** | Pointer to **[]string** |  | [optional] 
**CurrencyType** | Pointer to **string** |  | [optional] 
**Rc** | Pointer to **int32** |  | [optional] 

## Methods

### NewPerformance

`func NewPerformance() *Performance`

NewPerformance instantiates a new Performance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceWithDefaults

`func NewPerformanceWithDefaults() *Performance`

NewPerformanceWithDefaults instantiates a new Performance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Performance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Performance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Performance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Performance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCps

`func (o *Performance) GetCps() PerformanceCps`

GetCps returns the Cps field if non-nil, zero value otherwise.

### GetCpsOk

`func (o *Performance) GetCpsOk() (*PerformanceCps, bool)`

GetCpsOk returns a tuple with the Cps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCps

`func (o *Performance) SetCps(v PerformanceCps)`

SetCps sets Cps field to given value.

### HasCps

`func (o *Performance) HasCps() bool`

HasCps returns a boolean if a field has been set.

### GetTpps

`func (o *Performance) GetTpps() PerformanceTpps`

GetTpps returns the Tpps field if non-nil, zero value otherwise.

### GetTppsOk

`func (o *Performance) GetTppsOk() (*PerformanceTpps, bool)`

GetTppsOk returns a tuple with the Tpps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpps

`func (o *Performance) SetTpps(v PerformanceTpps)`

SetTpps sets Tpps field to given value.

### HasTpps

`func (o *Performance) HasTpps() bool`

HasTpps returns a boolean if a field has been set.

### GetNav

`func (o *Performance) GetNav() PerformanceNav`

GetNav returns the Nav field if non-nil, zero value otherwise.

### GetNavOk

`func (o *Performance) GetNavOk() (*PerformanceNav, bool)`

GetNavOk returns a tuple with the Nav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNav

`func (o *Performance) SetNav(v PerformanceNav)`

SetNav sets Nav field to given value.

### HasNav

`func (o *Performance) HasNav() bool`

HasNav returns a boolean if a field has been set.

### GetPm

`func (o *Performance) GetPm() string`

GetPm returns the Pm field if non-nil, zero value otherwise.

### GetPmOk

`func (o *Performance) GetPmOk() (*string, bool)`

GetPmOk returns a tuple with the Pm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm

`func (o *Performance) SetPm(v string)`

SetPm sets Pm field to given value.

### HasPm

`func (o *Performance) HasPm() bool`

HasPm returns a boolean if a field has been set.

### GetIncluded

`func (o *Performance) GetIncluded() []string`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *Performance) GetIncludedOk() (*[]string, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *Performance) SetIncluded(v []string)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *Performance) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetCurrencyType

`func (o *Performance) GetCurrencyType() string`

GetCurrencyType returns the CurrencyType field if non-nil, zero value otherwise.

### GetCurrencyTypeOk

`func (o *Performance) GetCurrencyTypeOk() (*string, bool)`

GetCurrencyTypeOk returns a tuple with the CurrencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyType

`func (o *Performance) SetCurrencyType(v string)`

SetCurrencyType sets CurrencyType field to given value.

### HasCurrencyType

`func (o *Performance) HasCurrencyType() bool`

HasCurrencyType returns a boolean if a field has been set.

### GetRc

`func (o *Performance) GetRc() int32`

GetRc returns the Rc field if non-nil, zero value otherwise.

### GetRcOk

`func (o *Performance) GetRcOk() (*int32, bool)`

GetRcOk returns a tuple with the Rc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRc

`func (o *Performance) SetRc(v int32)`

SetRc sets Rc field to given value.

### HasRc

`func (o *Performance) HasRc() bool`

HasRc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


