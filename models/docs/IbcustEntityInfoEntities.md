# IbcustEntityInfoEntities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanTrade** | Pointer to **bool** |  | [optional] 
**CanSign** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Name** | Pointer to [**IbcustEntityInfoName**](_ibcust_entity_info_name.md) |  | [optional] 
**Address** | Pointer to [**IbcustEntityInfoAddress**](_ibcust_entity_info_address.md) |  | [optional] 
**IdentDocs** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewIbcustEntityInfoEntities

`func NewIbcustEntityInfoEntities() *IbcustEntityInfoEntities`

NewIbcustEntityInfoEntities instantiates a new IbcustEntityInfoEntities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIbcustEntityInfoEntitiesWithDefaults

`func NewIbcustEntityInfoEntitiesWithDefaults() *IbcustEntityInfoEntities`

NewIbcustEntityInfoEntitiesWithDefaults instantiates a new IbcustEntityInfoEntities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanTrade

`func (o *IbcustEntityInfoEntities) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *IbcustEntityInfoEntities) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *IbcustEntityInfoEntities) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *IbcustEntityInfoEntities) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanSign

`func (o *IbcustEntityInfoEntities) GetCanSign() bool`

GetCanSign returns the CanSign field if non-nil, zero value otherwise.

### GetCanSignOk

`func (o *IbcustEntityInfoEntities) GetCanSignOk() (*bool, bool)`

GetCanSignOk returns a tuple with the CanSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSign

`func (o *IbcustEntityInfoEntities) SetCanSign(v bool)`

SetCanSign sets CanSign field to given value.

### HasCanSign

`func (o *IbcustEntityInfoEntities) HasCanSign() bool`

HasCanSign returns a boolean if a field has been set.

### GetType

`func (o *IbcustEntityInfoEntities) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IbcustEntityInfoEntities) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IbcustEntityInfoEntities) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IbcustEntityInfoEntities) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *IbcustEntityInfoEntities) GetName() IbcustEntityInfoName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IbcustEntityInfoEntities) GetNameOk() (*IbcustEntityInfoName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IbcustEntityInfoEntities) SetName(v IbcustEntityInfoName)`

SetName sets Name field to given value.

### HasName

`func (o *IbcustEntityInfoEntities) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *IbcustEntityInfoEntities) GetAddress() IbcustEntityInfoAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IbcustEntityInfoEntities) GetAddressOk() (*IbcustEntityInfoAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IbcustEntityInfoEntities) SetAddress(v IbcustEntityInfoAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IbcustEntityInfoEntities) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIdentDocs

`func (o *IbcustEntityInfoEntities) GetIdentDocs() []map[string]interface{}`

GetIdentDocs returns the IdentDocs field if non-nil, zero value otherwise.

### GetIdentDocsOk

`func (o *IbcustEntityInfoEntities) GetIdentDocsOk() (*[]map[string]interface{}, bool)`

GetIdentDocsOk returns a tuple with the IdentDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentDocs

`func (o *IbcustEntityInfoEntities) SetIdentDocs(v []map[string]interface{})`

SetIdentDocs sets IdentDocs field to given value.

### HasIdentDocs

`func (o *IbcustEntityInfoEntities) HasIdentDocs() bool`

HasIdentDocs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


