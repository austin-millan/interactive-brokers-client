# PerformanceCps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dates** | Pointer to **[]string** | array of dates, the length should be same as the length of returns inside data. | [optional] 
**Freq** | Pointer to **string** | D means Day | [optional] 
**Data** | Pointer to [**[]PerformanceCpsData**](PerformanceCpsData.md) |  | [optional] 

## Methods

### NewPerformanceCps

`func NewPerformanceCps() *PerformanceCps`

NewPerformanceCps instantiates a new PerformanceCps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceCpsWithDefaults

`func NewPerformanceCpsWithDefaults() *PerformanceCps`

NewPerformanceCpsWithDefaults instantiates a new PerformanceCps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDates

`func (o *PerformanceCps) GetDates() []string`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *PerformanceCps) GetDatesOk() (*[]string, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *PerformanceCps) SetDates(v []string)`

SetDates sets Dates field to given value.

### HasDates

`func (o *PerformanceCps) HasDates() bool`

HasDates returns a boolean if a field has been set.

### GetFreq

`func (o *PerformanceCps) GetFreq() string`

GetFreq returns the Freq field if non-nil, zero value otherwise.

### GetFreqOk

`func (o *PerformanceCps) GetFreqOk() (*string, bool)`

GetFreqOk returns a tuple with the Freq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreq

`func (o *PerformanceCps) SetFreq(v string)`

SetFreq sets Freq field to given value.

### HasFreq

`func (o *PerformanceCps) HasFreq() bool`

HasFreq returns a boolean if a field has been set.

### GetData

`func (o *PerformanceCps) GetData() []PerformanceCpsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PerformanceCps) GetDataOk() (*[]PerformanceCpsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PerformanceCps) SetData(v []PerformanceCpsData)`

SetData sets Data field to given value.

### HasData

`func (o *PerformanceCps) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


