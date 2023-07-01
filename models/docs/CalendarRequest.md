# CalendarRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to [**CalendarRequestDate**](calendar_request_date.md) |  | [optional] 
**Filters** | Pointer to [**CalendarRequestFilters**](calendar_request_filters.md) |  | [optional] 

## Methods

### NewCalendarRequest

`func NewCalendarRequest() *CalendarRequest`

NewCalendarRequest instantiates a new CalendarRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarRequestWithDefaults

`func NewCalendarRequestWithDefaults() *CalendarRequest`

NewCalendarRequestWithDefaults instantiates a new CalendarRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CalendarRequest) GetDate() CalendarRequestDate`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CalendarRequest) GetDateOk() (*CalendarRequestDate, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CalendarRequest) SetDate(v CalendarRequestDate)`

SetDate sets Date field to given value.

### HasDate

`func (o *CalendarRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetFilters

`func (o *CalendarRequest) GetFilters() CalendarRequestFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CalendarRequest) GetFiltersOk() (*CalendarRequestFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CalendarRequest) SetFilters(v CalendarRequestFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *CalendarRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


