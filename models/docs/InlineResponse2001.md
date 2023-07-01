# InlineResponse2001

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**A** | Pointer to **int32** | optional, if A doesn&#39;t exist, it means user can&#39;t toggle this option. 0-off, 1-on. | [optional] 
**FC** | Pointer to **string** | fyi code | [optional] 
**H** | Pointer to **int32** | disclaimer read, 1 &#x3D; yes, &#x3D; 0 no. | [optional] 
**FD** | Pointer to **string** | detailed description | [optional] 
**FN** | Pointer to **string** | title | [optional] 

## Methods

### NewInlineResponse2001

`func NewInlineResponse2001() *InlineResponse2001`

NewInlineResponse2001 instantiates a new InlineResponse2001 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2001WithDefaults

`func NewInlineResponse2001WithDefaults() *InlineResponse2001`

NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetA

`func (o *InlineResponse2001) GetA() int32`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *InlineResponse2001) GetAOk() (*int32, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *InlineResponse2001) SetA(v int32)`

SetA sets A field to given value.

### HasA

`func (o *InlineResponse2001) HasA() bool`

HasA returns a boolean if a field has been set.

### GetFC

`func (o *InlineResponse2001) GetFC() string`

GetFC returns the FC field if non-nil, zero value otherwise.

### GetFCOk

`func (o *InlineResponse2001) GetFCOk() (*string, bool)`

GetFCOk returns a tuple with the FC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFC

`func (o *InlineResponse2001) SetFC(v string)`

SetFC sets FC field to given value.

### HasFC

`func (o *InlineResponse2001) HasFC() bool`

HasFC returns a boolean if a field has been set.

### GetH

`func (o *InlineResponse2001) GetH() int32`

GetH returns the H field if non-nil, zero value otherwise.

### GetHOk

`func (o *InlineResponse2001) GetHOk() (*int32, bool)`

GetHOk returns a tuple with the H field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH

`func (o *InlineResponse2001) SetH(v int32)`

SetH sets H field to given value.

### HasH

`func (o *InlineResponse2001) HasH() bool`

HasH returns a boolean if a field has been set.

### GetFD

`func (o *InlineResponse2001) GetFD() string`

GetFD returns the FD field if non-nil, zero value otherwise.

### GetFDOk

`func (o *InlineResponse2001) GetFDOk() (*string, bool)`

GetFDOk returns a tuple with the FD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFD

`func (o *InlineResponse2001) SetFD(v string)`

SetFD sets FD field to given value.

### HasFD

`func (o *InlineResponse2001) HasFD() bool`

HasFD returns a boolean if a field has been set.

### GetFN

`func (o *InlineResponse2001) GetFN() string`

GetFN returns the FN field if non-nil, zero value otherwise.

### GetFNOk

`func (o *InlineResponse2001) GetFNOk() (*string, bool)`

GetFNOk returns a tuple with the FN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFN

`func (o *InlineResponse2001) SetFN(v string)`

SetFN sets FN field to given value.

### HasFN

`func (o *InlineResponse2001) HasFN() bool`

HasFN returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


