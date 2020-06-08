# HistoryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **string** | start date time | [optional] 
**MdAvailability** | **string** | Market Data Availability. The field may contain two chars. The first char is the primary code: R &#x3D; Realtime, D &#x3D; Delayed, Z &#x3D; Frozen, Y &#x3D; Frozen Delayed. The second char is the secondary code: P &#x3D; Snapshot Available, p &#x3D; Consolidated.  | [optional] 
**BarLength** | **int32** |  | [optional] 
**Delay** | **int32** |  | [optional] 
**High** | **string** | High value during this time series with format %h/%v/%t. %h is the high price (scaled by priceFactor), %v is volume (volume factor will always be 100 (reported volume &#x3D; actual volume/100)) and %t is minutes from start time of the chart  | [optional] 
**Low** | **string** | Low value during this time series with format %l/%v/%t. %l is the low price (scaled by priceFactor), %v is volume (volume factor will always be 100 (reported volume &#x3D; actual volume/100)) and %t is minutes from start time of the chart  | [optional] 
**Symbol** | **string** |  | [optional] 
**Text** | **string** |  | [optional] 
**TickNum** | **string** |  | [optional] 
**TimePeriod** | **string** |  | [optional] 
**PriceFactor** | **int32** | priceFactor is price increment obtained from display rule | [optional] 
**Data** | [**[]HistoryDataData**](history_data_data.md) |  | [optional] 
**Points** | **float32** | total number of points | [optional] 
**TravelTime** | **float32** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


