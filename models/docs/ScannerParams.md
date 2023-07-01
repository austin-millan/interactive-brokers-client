# ScannerParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instrument** | Pointer to **string** | for example-STK | [optional] 
**Type** | Pointer to **string** | for example-TOP_PERC_GAIN | [optional] 
**Filter** | Pointer to [**[]ScannerParamsFilter**](ScannerParamsFilter.md) |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 

## Methods

### NewScannerParams

`func NewScannerParams() *ScannerParams`

NewScannerParams instantiates a new ScannerParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScannerParamsWithDefaults

`func NewScannerParamsWithDefaults() *ScannerParams`

NewScannerParamsWithDefaults instantiates a new ScannerParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrument

`func (o *ScannerParams) GetInstrument() string`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *ScannerParams) GetInstrumentOk() (*string, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *ScannerParams) SetInstrument(v string)`

SetInstrument sets Instrument field to given value.

### HasInstrument

`func (o *ScannerParams) HasInstrument() bool`

HasInstrument returns a boolean if a field has been set.

### GetType

`func (o *ScannerParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScannerParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScannerParams) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ScannerParams) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFilter

`func (o *ScannerParams) GetFilter() []ScannerParamsFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ScannerParams) GetFilterOk() (*[]ScannerParamsFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ScannerParams) SetFilter(v []ScannerParamsFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ScannerParams) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetLocation

`func (o *ScannerParams) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ScannerParams) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ScannerParams) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ScannerParams) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetSize

`func (o *ScannerParams) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ScannerParams) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ScannerParams) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ScannerParams) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


