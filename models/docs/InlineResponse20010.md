# InlineResponse20010

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **[]string** | Unique account id | [optional] 
**Aliases** | Pointer to **map[string]interface{}** | Account Id and its alias | [optional] 
**SelectedAccount** | Pointer to **string** |  | [optional] 

## Methods

### NewInlineResponse20010

`func NewInlineResponse20010() *InlineResponse20010`

NewInlineResponse20010 instantiates a new InlineResponse20010 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20010WithDefaults

`func NewInlineResponse20010WithDefaults() *InlineResponse20010`

NewInlineResponse20010WithDefaults instantiates a new InlineResponse20010 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *InlineResponse20010) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *InlineResponse20010) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *InlineResponse20010) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *InlineResponse20010) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetAliases

`func (o *InlineResponse20010) GetAliases() map[string]interface{}`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *InlineResponse20010) GetAliasesOk() (*map[string]interface{}, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *InlineResponse20010) SetAliases(v map[string]interface{})`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *InlineResponse20010) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetSelectedAccount

`func (o *InlineResponse20010) GetSelectedAccount() string`

GetSelectedAccount returns the SelectedAccount field if non-nil, zero value otherwise.

### GetSelectedAccountOk

`func (o *InlineResponse20010) GetSelectedAccountOk() (*string, bool)`

GetSelectedAccountOk returns a tuple with the SelectedAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedAccount

`func (o *InlineResponse20010) SetSelectedAccount(v string)`

SetSelectedAccount sets SelectedAccount field to given value.

### HasSelectedAccount

`func (o *InlineResponse20010) HasSelectedAccount() bool`

HasSelectedAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


