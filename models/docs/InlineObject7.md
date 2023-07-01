# InlineObject7

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | symbol or name to be searched | 
**Name** | Pointer to **bool** | should be true if the search is to be performed by name. false by default. | [optional] 
**SecType** | Pointer to **string** | If search is done by name, only the assets provided in this field will be returned. Currently, only STK is supported. | [optional] 

## Methods

### NewInlineObject7

`func NewInlineObject7(symbol string, ) *InlineObject7`

NewInlineObject7 instantiates a new InlineObject7 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject7WithDefaults

`func NewInlineObject7WithDefaults() *InlineObject7`

NewInlineObject7WithDefaults instantiates a new InlineObject7 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *InlineObject7) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *InlineObject7) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *InlineObject7) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *InlineObject7) GetName() bool`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject7) GetNameOk() (*bool, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject7) SetName(v bool)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject7) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecType

`func (o *InlineObject7) GetSecType() string`

GetSecType returns the SecType field if non-nil, zero value otherwise.

### GetSecTypeOk

`func (o *InlineObject7) GetSecTypeOk() (*string, bool)`

GetSecTypeOk returns a tuple with the SecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecType

`func (o *InlineObject7) SetSecType(v string)`

SetSecType sets SecType field to given value.

### HasSecType

`func (o *InlineObject7) HasSecType() bool`

HasSecType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


