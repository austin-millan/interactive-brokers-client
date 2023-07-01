# AuthStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authenticated** | Pointer to **bool** | Brokerage session is authenticated | [optional] 
**Connected** | Pointer to **bool** | Connected to backend | [optional] 
**Competing** | Pointer to **bool** | Brokerage session is competing, e.g. user is logged in to IBKR Mobile, WebTrader, TWS or other trading platforms. | [optional] 
**Fail** | Pointer to **string** | Authentication failed, why. | [optional] 
**Message** | Pointer to **string** | System messages that may affect trading | [optional] 
**Prompts** | Pointer to **[]string** | Prompt messages that may affect trading or the account | [optional] 

## Methods

### NewAuthStatus

`func NewAuthStatus() *AuthStatus`

NewAuthStatus instantiates a new AuthStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthStatusWithDefaults

`func NewAuthStatusWithDefaults() *AuthStatus`

NewAuthStatusWithDefaults instantiates a new AuthStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticated

`func (o *AuthStatus) GetAuthenticated() bool`

GetAuthenticated returns the Authenticated field if non-nil, zero value otherwise.

### GetAuthenticatedOk

`func (o *AuthStatus) GetAuthenticatedOk() (*bool, bool)`

GetAuthenticatedOk returns a tuple with the Authenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticated

`func (o *AuthStatus) SetAuthenticated(v bool)`

SetAuthenticated sets Authenticated field to given value.

### HasAuthenticated

`func (o *AuthStatus) HasAuthenticated() bool`

HasAuthenticated returns a boolean if a field has been set.

### GetConnected

`func (o *AuthStatus) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *AuthStatus) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *AuthStatus) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *AuthStatus) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetCompeting

`func (o *AuthStatus) GetCompeting() bool`

GetCompeting returns the Competing field if non-nil, zero value otherwise.

### GetCompetingOk

`func (o *AuthStatus) GetCompetingOk() (*bool, bool)`

GetCompetingOk returns a tuple with the Competing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompeting

`func (o *AuthStatus) SetCompeting(v bool)`

SetCompeting sets Competing field to given value.

### HasCompeting

`func (o *AuthStatus) HasCompeting() bool`

HasCompeting returns a boolean if a field has been set.

### GetFail

`func (o *AuthStatus) GetFail() string`

GetFail returns the Fail field if non-nil, zero value otherwise.

### GetFailOk

`func (o *AuthStatus) GetFailOk() (*string, bool)`

GetFailOk returns a tuple with the Fail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFail

`func (o *AuthStatus) SetFail(v string)`

SetFail sets Fail field to given value.

### HasFail

`func (o *AuthStatus) HasFail() bool`

HasFail returns a boolean if a field has been set.

### GetMessage

`func (o *AuthStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuthStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuthStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuthStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPrompts

`func (o *AuthStatus) GetPrompts() []string`

GetPrompts returns the Prompts field if non-nil, zero value otherwise.

### GetPromptsOk

`func (o *AuthStatus) GetPromptsOk() (*[]string, bool)`

GetPromptsOk returns a tuple with the Prompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompts

`func (o *AuthStatus) SetPrompts(v []string)`

SetPrompts sets Prompts field to given value.

### HasPrompts

`func (o *AuthStatus) HasPrompts() bool`

HasPrompts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


