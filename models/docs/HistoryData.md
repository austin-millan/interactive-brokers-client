# HistoryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **string** | start date time | [optional] 
**MdAvailability** | Pointer to **string** | Market Data Availability. The field may contain two chars. The first char is the primary code: R &#x3D; Realtime, D &#x3D; Delayed, Z &#x3D; Frozen, Y &#x3D; Frozen Delayed. The second char is the secondary code: P &#x3D; Snapshot Available, p &#x3D; Consolidated.  | [optional] 
**BarLength** | Pointer to **int32** |  | [optional] 
**Delay** | Pointer to **int32** |  | [optional] 
**High** | Pointer to **string** | High value during this time series with format %h/%v/%t. %h is the high price (scaled by priceFactor), %v is volume (volume factor will always be 100 (reported volume &#x3D; actual volume/100)) and %t is minutes from start time of the chart  | [optional] 
**Low** | Pointer to **string** | Low value during this time series with format %l/%v/%t. %l is the low price (scaled by priceFactor), %v is volume (volume factor will always be 100 (reported volume &#x3D; actual volume/100)) and %t is minutes from start time of the chart  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**TickNum** | Pointer to **string** |  | [optional] 
**TimePeriod** | Pointer to **string** |  | [optional] 
**PriceFactor** | Pointer to **int32** | priceFactor is price increment obtained from display rule | [optional] 
**Data** | Pointer to [**[]HistoryDataData**](HistoryDataData.md) |  | [optional] 
**Points** | Pointer to **float32** | total number of points | [optional] 
**TravelTime** | Pointer to **float32** |  | [optional] 

## Methods

### NewHistoryData

`func NewHistoryData() *HistoryData`

NewHistoryData instantiates a new HistoryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryDataWithDefaults

`func NewHistoryDataWithDefaults() *HistoryData`

NewHistoryDataWithDefaults instantiates a new HistoryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *HistoryData) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *HistoryData) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *HistoryData) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *HistoryData) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetMdAvailability

`func (o *HistoryData) GetMdAvailability() string`

GetMdAvailability returns the MdAvailability field if non-nil, zero value otherwise.

### GetMdAvailabilityOk

`func (o *HistoryData) GetMdAvailabilityOk() (*string, bool)`

GetMdAvailabilityOk returns a tuple with the MdAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdAvailability

`func (o *HistoryData) SetMdAvailability(v string)`

SetMdAvailability sets MdAvailability field to given value.

### HasMdAvailability

`func (o *HistoryData) HasMdAvailability() bool`

HasMdAvailability returns a boolean if a field has been set.

### GetBarLength

`func (o *HistoryData) GetBarLength() int32`

GetBarLength returns the BarLength field if non-nil, zero value otherwise.

### GetBarLengthOk

`func (o *HistoryData) GetBarLengthOk() (*int32, bool)`

GetBarLengthOk returns a tuple with the BarLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarLength

`func (o *HistoryData) SetBarLength(v int32)`

SetBarLength sets BarLength field to given value.

### HasBarLength

`func (o *HistoryData) HasBarLength() bool`

HasBarLength returns a boolean if a field has been set.

### GetDelay

`func (o *HistoryData) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *HistoryData) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *HistoryData) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *HistoryData) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetHigh

`func (o *HistoryData) GetHigh() string`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *HistoryData) GetHighOk() (*string, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *HistoryData) SetHigh(v string)`

SetHigh sets High field to given value.

### HasHigh

`func (o *HistoryData) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetLow

`func (o *HistoryData) GetLow() string`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *HistoryData) GetLowOk() (*string, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *HistoryData) SetLow(v string)`

SetLow sets Low field to given value.

### HasLow

`func (o *HistoryData) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetSymbol

`func (o *HistoryData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *HistoryData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *HistoryData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *HistoryData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetText

`func (o *HistoryData) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *HistoryData) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *HistoryData) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *HistoryData) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTickNum

`func (o *HistoryData) GetTickNum() string`

GetTickNum returns the TickNum field if non-nil, zero value otherwise.

### GetTickNumOk

`func (o *HistoryData) GetTickNumOk() (*string, bool)`

GetTickNumOk returns a tuple with the TickNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickNum

`func (o *HistoryData) SetTickNum(v string)`

SetTickNum sets TickNum field to given value.

### HasTickNum

`func (o *HistoryData) HasTickNum() bool`

HasTickNum returns a boolean if a field has been set.

### GetTimePeriod

`func (o *HistoryData) GetTimePeriod() string`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *HistoryData) GetTimePeriodOk() (*string, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *HistoryData) SetTimePeriod(v string)`

SetTimePeriod sets TimePeriod field to given value.

### HasTimePeriod

`func (o *HistoryData) HasTimePeriod() bool`

HasTimePeriod returns a boolean if a field has been set.

### GetPriceFactor

`func (o *HistoryData) GetPriceFactor() int32`

GetPriceFactor returns the PriceFactor field if non-nil, zero value otherwise.

### GetPriceFactorOk

`func (o *HistoryData) GetPriceFactorOk() (*int32, bool)`

GetPriceFactorOk returns a tuple with the PriceFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceFactor

`func (o *HistoryData) SetPriceFactor(v int32)`

SetPriceFactor sets PriceFactor field to given value.

### HasPriceFactor

`func (o *HistoryData) HasPriceFactor() bool`

HasPriceFactor returns a boolean if a field has been set.

### GetData

`func (o *HistoryData) GetData() []HistoryDataData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoryData) GetDataOk() (*[]HistoryDataData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HistoryData) SetData(v []HistoryDataData)`

SetData sets Data field to given value.

### HasData

`func (o *HistoryData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPoints

`func (o *HistoryData) GetPoints() float32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *HistoryData) GetPointsOk() (*float32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *HistoryData) SetPoints(v float32)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *HistoryData) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetTravelTime

`func (o *HistoryData) GetTravelTime() float32`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *HistoryData) GetTravelTimeOk() (*float32, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *HistoryData) SetTravelTime(v float32)`

SetTravelTime sets TravelTime field to given value.

### HasTravelTime

`func (o *HistoryData) HasTravelTime() bool`

HasTravelTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


