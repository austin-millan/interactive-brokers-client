# CalendarRequestFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecentlyHeld** | Pointer to **string** | value can be &#39;true&#39; or &#39;false&#39;. | [optional] 
**CorporateEarnings** | Pointer to **string** | value can be &#39;true&#39; or &#39;false&#39;. | [optional] 
**DivExDates** | Pointer to **string** | value can be &#39;true&#39; or &#39;false&#39;. | [optional] 
**Ipo** | Pointer to **string** | value can be &#39;true&#39; or &#39;false&#39;. | [optional] 
**Splits** | Pointer to **string** | value can be &#39;true&#39; or &#39;false&#39;. | [optional] 
**CorporateEvents** | Pointer to **string** | value can be &#39;true&#39; or &#39;false&#39;. | [optional] 
**EconomicEvents** | Pointer to **string** | value can be &#39;true&#39; or &#39;false&#39;. | [optional] 
**OptionShowMonthly** | Pointer to **string** | value can be &#39;true&#39; or &#39;false&#39;. | [optional] 
**OptionShowWeekly** | Pointer to **string** | value can be &#39;true&#39; or &#39;false&#39;. | [optional] 
**Country** | Pointer to **string** | default is &#39;All&#39;. | [optional] 
**Limit** | Pointer to **string** | default is &#39;250&#39;. | [optional] 
**LimitRegion** | Pointer to **string** | default is &#39;50&#39;. | [optional] 

## Methods

### NewCalendarRequestFilters

`func NewCalendarRequestFilters() *CalendarRequestFilters`

NewCalendarRequestFilters instantiates a new CalendarRequestFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarRequestFiltersWithDefaults

`func NewCalendarRequestFiltersWithDefaults() *CalendarRequestFilters`

NewCalendarRequestFiltersWithDefaults instantiates a new CalendarRequestFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecentlyHeld

`func (o *CalendarRequestFilters) GetRecentlyHeld() string`

GetRecentlyHeld returns the RecentlyHeld field if non-nil, zero value otherwise.

### GetRecentlyHeldOk

`func (o *CalendarRequestFilters) GetRecentlyHeldOk() (*string, bool)`

GetRecentlyHeldOk returns a tuple with the RecentlyHeld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentlyHeld

`func (o *CalendarRequestFilters) SetRecentlyHeld(v string)`

SetRecentlyHeld sets RecentlyHeld field to given value.

### HasRecentlyHeld

`func (o *CalendarRequestFilters) HasRecentlyHeld() bool`

HasRecentlyHeld returns a boolean if a field has been set.

### GetCorporateEarnings

`func (o *CalendarRequestFilters) GetCorporateEarnings() string`

GetCorporateEarnings returns the CorporateEarnings field if non-nil, zero value otherwise.

### GetCorporateEarningsOk

`func (o *CalendarRequestFilters) GetCorporateEarningsOk() (*string, bool)`

GetCorporateEarningsOk returns a tuple with the CorporateEarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateEarnings

`func (o *CalendarRequestFilters) SetCorporateEarnings(v string)`

SetCorporateEarnings sets CorporateEarnings field to given value.

### HasCorporateEarnings

`func (o *CalendarRequestFilters) HasCorporateEarnings() bool`

HasCorporateEarnings returns a boolean if a field has been set.

### GetDivExDates

`func (o *CalendarRequestFilters) GetDivExDates() string`

GetDivExDates returns the DivExDates field if non-nil, zero value otherwise.

### GetDivExDatesOk

`func (o *CalendarRequestFilters) GetDivExDatesOk() (*string, bool)`

GetDivExDatesOk returns a tuple with the DivExDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivExDates

`func (o *CalendarRequestFilters) SetDivExDates(v string)`

SetDivExDates sets DivExDates field to given value.

### HasDivExDates

`func (o *CalendarRequestFilters) HasDivExDates() bool`

HasDivExDates returns a boolean if a field has been set.

### GetIpo

`func (o *CalendarRequestFilters) GetIpo() string`

GetIpo returns the Ipo field if non-nil, zero value otherwise.

### GetIpoOk

`func (o *CalendarRequestFilters) GetIpoOk() (*string, bool)`

GetIpoOk returns a tuple with the Ipo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpo

`func (o *CalendarRequestFilters) SetIpo(v string)`

SetIpo sets Ipo field to given value.

### HasIpo

`func (o *CalendarRequestFilters) HasIpo() bool`

HasIpo returns a boolean if a field has been set.

### GetSplits

`func (o *CalendarRequestFilters) GetSplits() string`

GetSplits returns the Splits field if non-nil, zero value otherwise.

### GetSplitsOk

`func (o *CalendarRequestFilters) GetSplitsOk() (*string, bool)`

GetSplitsOk returns a tuple with the Splits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplits

`func (o *CalendarRequestFilters) SetSplits(v string)`

SetSplits sets Splits field to given value.

### HasSplits

`func (o *CalendarRequestFilters) HasSplits() bool`

HasSplits returns a boolean if a field has been set.

### GetCorporateEvents

`func (o *CalendarRequestFilters) GetCorporateEvents() string`

GetCorporateEvents returns the CorporateEvents field if non-nil, zero value otherwise.

### GetCorporateEventsOk

`func (o *CalendarRequestFilters) GetCorporateEventsOk() (*string, bool)`

GetCorporateEventsOk returns a tuple with the CorporateEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateEvents

`func (o *CalendarRequestFilters) SetCorporateEvents(v string)`

SetCorporateEvents sets CorporateEvents field to given value.

### HasCorporateEvents

`func (o *CalendarRequestFilters) HasCorporateEvents() bool`

HasCorporateEvents returns a boolean if a field has been set.

### GetEconomicEvents

`func (o *CalendarRequestFilters) GetEconomicEvents() string`

GetEconomicEvents returns the EconomicEvents field if non-nil, zero value otherwise.

### GetEconomicEventsOk

`func (o *CalendarRequestFilters) GetEconomicEventsOk() (*string, bool)`

GetEconomicEventsOk returns a tuple with the EconomicEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEconomicEvents

`func (o *CalendarRequestFilters) SetEconomicEvents(v string)`

SetEconomicEvents sets EconomicEvents field to given value.

### HasEconomicEvents

`func (o *CalendarRequestFilters) HasEconomicEvents() bool`

HasEconomicEvents returns a boolean if a field has been set.

### GetOptionShowMonthly

`func (o *CalendarRequestFilters) GetOptionShowMonthly() string`

GetOptionShowMonthly returns the OptionShowMonthly field if non-nil, zero value otherwise.

### GetOptionShowMonthlyOk

`func (o *CalendarRequestFilters) GetOptionShowMonthlyOk() (*string, bool)`

GetOptionShowMonthlyOk returns a tuple with the OptionShowMonthly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionShowMonthly

`func (o *CalendarRequestFilters) SetOptionShowMonthly(v string)`

SetOptionShowMonthly sets OptionShowMonthly field to given value.

### HasOptionShowMonthly

`func (o *CalendarRequestFilters) HasOptionShowMonthly() bool`

HasOptionShowMonthly returns a boolean if a field has been set.

### GetOptionShowWeekly

`func (o *CalendarRequestFilters) GetOptionShowWeekly() string`

GetOptionShowWeekly returns the OptionShowWeekly field if non-nil, zero value otherwise.

### GetOptionShowWeeklyOk

`func (o *CalendarRequestFilters) GetOptionShowWeeklyOk() (*string, bool)`

GetOptionShowWeeklyOk returns a tuple with the OptionShowWeekly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionShowWeekly

`func (o *CalendarRequestFilters) SetOptionShowWeekly(v string)`

SetOptionShowWeekly sets OptionShowWeekly field to given value.

### HasOptionShowWeekly

`func (o *CalendarRequestFilters) HasOptionShowWeekly() bool`

HasOptionShowWeekly returns a boolean if a field has been set.

### GetCountry

`func (o *CalendarRequestFilters) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CalendarRequestFilters) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CalendarRequestFilters) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CalendarRequestFilters) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetLimit

`func (o *CalendarRequestFilters) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CalendarRequestFilters) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CalendarRequestFilters) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CalendarRequestFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLimitRegion

`func (o *CalendarRequestFilters) GetLimitRegion() string`

GetLimitRegion returns the LimitRegion field if non-nil, zero value otherwise.

### GetLimitRegionOk

`func (o *CalendarRequestFilters) GetLimitRegionOk() (*string, bool)`

GetLimitRegionOk returns a tuple with the LimitRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitRegion

`func (o *CalendarRequestFilters) SetLimitRegion(v string)`

SetLimitRegion sets LimitRegion field to given value.

### HasLimitRegion

`func (o *CalendarRequestFilters) HasLimitRegion() bool`

HasLimitRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


