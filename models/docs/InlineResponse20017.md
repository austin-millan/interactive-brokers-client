# InlineResponse20017

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **string** |  | [optional] 
**Conid** | Pointer to **int32** |  | [optional] 
**Updated** | Pointer to **int32** |  | [optional] 
**Var31** | Pointer to **string** | Last Price | [optional] 
**Var55** | Pointer to **string** | Symbol | [optional] 
**Var58** | Pointer to **string** | Text | [optional] 
**Var70** | Pointer to **string** | High | [optional] 
**Var71** | Pointer to **string** | Low | [optional] 
**Var72** | Pointer to **string** | Position | [optional] 
**Var73** | Pointer to **string** | Market Value | [optional] 
**Var74** | Pointer to **string** | Average Price | [optional] 
**Var75** | Pointer to **string** | Unrealized PnL | [optional] 
**Var76** | Pointer to **string** | Formatted position | [optional] 
**Var77** | Pointer to **string** | Formatted Unrealized PnL | [optional] 
**Var78** | Pointer to **string** | Daily PnL | [optional] 
**Var82** | Pointer to **string** | Change Price | [optional] 
**Var83** | Pointer to **string** | Change Percent | [optional] 
**Var84** | Pointer to **string** | Bid Price | [optional] 
**Var85** | Pointer to **string** | Ask Size | [optional] 
**Var86** | Pointer to **string** | Ask Price | [optional] 
**Var87** | Pointer to **string** | Volume | [optional] 
**Var88** | Pointer to **string** | Bid Size | [optional] 
**Var6004** | Pointer to **string** | Exchange | [optional] 
**Var6008** | Pointer to **string** | Conid | [optional] 
**Var6070** | Pointer to **string** | Security Type | [optional] 
**Var6072** | Pointer to **string** | Months | [optional] 
**Var6073** | Pointer to **string** | Regular Expiry | [optional] 
**Var6119** | Pointer to **string** | Marker for market data delivery method (similar to request id) | [optional] 
**Var6457** | Pointer to **string** | Underlying Conid. Use /trsrv/secdef to get more information about the security | [optional] 
**Var6509** | Pointer to **string** | Market Data Availability. The field may contain two chars. The first char is the primary code: R &#x3D; Realtime, D &#x3D; Delayed, Z &#x3D; Frozen, Y &#x3D; Frozen Delayed. The second char is the secondary code: P &#x3D; Snapshot Available, p &#x3D; Consolidated.  | [optional] 
**Var7051** | Pointer to **string** | Company name | [optional] 
**Var7094** | Pointer to **string** | Conid + Exchange | [optional] 
**Var7219** | Pointer to **string** | Contract Description | [optional] 
**Var7220** | Pointer to **string** | Contract Description | [optional] 
**Var7221** | Pointer to **string** | Listing Exchange | [optional] 
**Var7280** | Pointer to **string** | Industry | [optional] 
**Var7281** | Pointer to **string** | Category | [optional] 
**Var7282** | Pointer to **string** | Average Daily Volume | [optional] 
**Var7633** | Pointer to **string** | Implied volatility of the option | [optional] 
**Var7284** | Pointer to **string** | Historic Volume (30d) | [optional] 
**Var7285** | Pointer to **string** | Put/Call Ratio | [optional] 
**Var7286** | Pointer to **string** | Dividend Amount | [optional] 
**Var7287** | Pointer to **string** | Dividend Yield % | [optional] 
**Var7288** | Pointer to **string** | Ex-date of the dividend | [optional] 
**Var7289** | Pointer to **string** | Market Cap | [optional] 
**Var7290** | Pointer to **string** | P/E | [optional] 
**Var7291** | Pointer to **string** | EPS | [optional] 
**Var7292** | Pointer to **string** | Cost Basis | [optional] 
**Var7293** | Pointer to **string** | 52 Week High | [optional] 
**Var7294** | Pointer to **string** | 52 Week Low | [optional] 
**Var7295** | Pointer to **string** | Open Price | [optional] 
**Var7296** | Pointer to **string** | Close Price | [optional] 

## Methods

### NewInlineResponse20017

`func NewInlineResponse20017() *InlineResponse20017`

NewInlineResponse20017 instantiates a new InlineResponse20017 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20017WithDefaults

`func NewInlineResponse20017WithDefaults() *InlineResponse20017`

NewInlineResponse20017WithDefaults instantiates a new InlineResponse20017 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *InlineResponse20017) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *InlineResponse20017) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *InlineResponse20017) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *InlineResponse20017) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetConid

`func (o *InlineResponse20017) GetConid() int32`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *InlineResponse20017) GetConidOk() (*int32, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *InlineResponse20017) SetConid(v int32)`

SetConid sets Conid field to given value.

### HasConid

`func (o *InlineResponse20017) HasConid() bool`

HasConid returns a boolean if a field has been set.

### GetUpdated

`func (o *InlineResponse20017) GetUpdated() int32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *InlineResponse20017) GetUpdatedOk() (*int32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *InlineResponse20017) SetUpdated(v int32)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *InlineResponse20017) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetVar31

`func (o *InlineResponse20017) GetVar31() string`

GetVar31 returns the Var31 field if non-nil, zero value otherwise.

### GetVar31Ok

`func (o *InlineResponse20017) GetVar31Ok() (*string, bool)`

GetVar31Ok returns a tuple with the Var31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar31

`func (o *InlineResponse20017) SetVar31(v string)`

SetVar31 sets Var31 field to given value.

### HasVar31

`func (o *InlineResponse20017) HasVar31() bool`

HasVar31 returns a boolean if a field has been set.

### GetVar55

`func (o *InlineResponse20017) GetVar55() string`

GetVar55 returns the Var55 field if non-nil, zero value otherwise.

### GetVar55Ok

`func (o *InlineResponse20017) GetVar55Ok() (*string, bool)`

GetVar55Ok returns a tuple with the Var55 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar55

`func (o *InlineResponse20017) SetVar55(v string)`

SetVar55 sets Var55 field to given value.

### HasVar55

`func (o *InlineResponse20017) HasVar55() bool`

HasVar55 returns a boolean if a field has been set.

### GetVar58

`func (o *InlineResponse20017) GetVar58() string`

GetVar58 returns the Var58 field if non-nil, zero value otherwise.

### GetVar58Ok

`func (o *InlineResponse20017) GetVar58Ok() (*string, bool)`

GetVar58Ok returns a tuple with the Var58 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar58

`func (o *InlineResponse20017) SetVar58(v string)`

SetVar58 sets Var58 field to given value.

### HasVar58

`func (o *InlineResponse20017) HasVar58() bool`

HasVar58 returns a boolean if a field has been set.

### GetVar70

`func (o *InlineResponse20017) GetVar70() string`

GetVar70 returns the Var70 field if non-nil, zero value otherwise.

### GetVar70Ok

`func (o *InlineResponse20017) GetVar70Ok() (*string, bool)`

GetVar70Ok returns a tuple with the Var70 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar70

`func (o *InlineResponse20017) SetVar70(v string)`

SetVar70 sets Var70 field to given value.

### HasVar70

`func (o *InlineResponse20017) HasVar70() bool`

HasVar70 returns a boolean if a field has been set.

### GetVar71

`func (o *InlineResponse20017) GetVar71() string`

GetVar71 returns the Var71 field if non-nil, zero value otherwise.

### GetVar71Ok

`func (o *InlineResponse20017) GetVar71Ok() (*string, bool)`

GetVar71Ok returns a tuple with the Var71 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar71

`func (o *InlineResponse20017) SetVar71(v string)`

SetVar71 sets Var71 field to given value.

### HasVar71

`func (o *InlineResponse20017) HasVar71() bool`

HasVar71 returns a boolean if a field has been set.

### GetVar72

`func (o *InlineResponse20017) GetVar72() string`

GetVar72 returns the Var72 field if non-nil, zero value otherwise.

### GetVar72Ok

`func (o *InlineResponse20017) GetVar72Ok() (*string, bool)`

GetVar72Ok returns a tuple with the Var72 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar72

`func (o *InlineResponse20017) SetVar72(v string)`

SetVar72 sets Var72 field to given value.

### HasVar72

`func (o *InlineResponse20017) HasVar72() bool`

HasVar72 returns a boolean if a field has been set.

### GetVar73

`func (o *InlineResponse20017) GetVar73() string`

GetVar73 returns the Var73 field if non-nil, zero value otherwise.

### GetVar73Ok

`func (o *InlineResponse20017) GetVar73Ok() (*string, bool)`

GetVar73Ok returns a tuple with the Var73 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar73

`func (o *InlineResponse20017) SetVar73(v string)`

SetVar73 sets Var73 field to given value.

### HasVar73

`func (o *InlineResponse20017) HasVar73() bool`

HasVar73 returns a boolean if a field has been set.

### GetVar74

`func (o *InlineResponse20017) GetVar74() string`

GetVar74 returns the Var74 field if non-nil, zero value otherwise.

### GetVar74Ok

`func (o *InlineResponse20017) GetVar74Ok() (*string, bool)`

GetVar74Ok returns a tuple with the Var74 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar74

`func (o *InlineResponse20017) SetVar74(v string)`

SetVar74 sets Var74 field to given value.

### HasVar74

`func (o *InlineResponse20017) HasVar74() bool`

HasVar74 returns a boolean if a field has been set.

### GetVar75

`func (o *InlineResponse20017) GetVar75() string`

GetVar75 returns the Var75 field if non-nil, zero value otherwise.

### GetVar75Ok

`func (o *InlineResponse20017) GetVar75Ok() (*string, bool)`

GetVar75Ok returns a tuple with the Var75 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar75

`func (o *InlineResponse20017) SetVar75(v string)`

SetVar75 sets Var75 field to given value.

### HasVar75

`func (o *InlineResponse20017) HasVar75() bool`

HasVar75 returns a boolean if a field has been set.

### GetVar76

`func (o *InlineResponse20017) GetVar76() string`

GetVar76 returns the Var76 field if non-nil, zero value otherwise.

### GetVar76Ok

`func (o *InlineResponse20017) GetVar76Ok() (*string, bool)`

GetVar76Ok returns a tuple with the Var76 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar76

`func (o *InlineResponse20017) SetVar76(v string)`

SetVar76 sets Var76 field to given value.

### HasVar76

`func (o *InlineResponse20017) HasVar76() bool`

HasVar76 returns a boolean if a field has been set.

### GetVar77

`func (o *InlineResponse20017) GetVar77() string`

GetVar77 returns the Var77 field if non-nil, zero value otherwise.

### GetVar77Ok

`func (o *InlineResponse20017) GetVar77Ok() (*string, bool)`

GetVar77Ok returns a tuple with the Var77 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar77

`func (o *InlineResponse20017) SetVar77(v string)`

SetVar77 sets Var77 field to given value.

### HasVar77

`func (o *InlineResponse20017) HasVar77() bool`

HasVar77 returns a boolean if a field has been set.

### GetVar78

`func (o *InlineResponse20017) GetVar78() string`

GetVar78 returns the Var78 field if non-nil, zero value otherwise.

### GetVar78Ok

`func (o *InlineResponse20017) GetVar78Ok() (*string, bool)`

GetVar78Ok returns a tuple with the Var78 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar78

`func (o *InlineResponse20017) SetVar78(v string)`

SetVar78 sets Var78 field to given value.

### HasVar78

`func (o *InlineResponse20017) HasVar78() bool`

HasVar78 returns a boolean if a field has been set.

### GetVar82

`func (o *InlineResponse20017) GetVar82() string`

GetVar82 returns the Var82 field if non-nil, zero value otherwise.

### GetVar82Ok

`func (o *InlineResponse20017) GetVar82Ok() (*string, bool)`

GetVar82Ok returns a tuple with the Var82 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar82

`func (o *InlineResponse20017) SetVar82(v string)`

SetVar82 sets Var82 field to given value.

### HasVar82

`func (o *InlineResponse20017) HasVar82() bool`

HasVar82 returns a boolean if a field has been set.

### GetVar83

`func (o *InlineResponse20017) GetVar83() string`

GetVar83 returns the Var83 field if non-nil, zero value otherwise.

### GetVar83Ok

`func (o *InlineResponse20017) GetVar83Ok() (*string, bool)`

GetVar83Ok returns a tuple with the Var83 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar83

`func (o *InlineResponse20017) SetVar83(v string)`

SetVar83 sets Var83 field to given value.

### HasVar83

`func (o *InlineResponse20017) HasVar83() bool`

HasVar83 returns a boolean if a field has been set.

### GetVar84

`func (o *InlineResponse20017) GetVar84() string`

GetVar84 returns the Var84 field if non-nil, zero value otherwise.

### GetVar84Ok

`func (o *InlineResponse20017) GetVar84Ok() (*string, bool)`

GetVar84Ok returns a tuple with the Var84 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar84

`func (o *InlineResponse20017) SetVar84(v string)`

SetVar84 sets Var84 field to given value.

### HasVar84

`func (o *InlineResponse20017) HasVar84() bool`

HasVar84 returns a boolean if a field has been set.

### GetVar85

`func (o *InlineResponse20017) GetVar85() string`

GetVar85 returns the Var85 field if non-nil, zero value otherwise.

### GetVar85Ok

`func (o *InlineResponse20017) GetVar85Ok() (*string, bool)`

GetVar85Ok returns a tuple with the Var85 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar85

`func (o *InlineResponse20017) SetVar85(v string)`

SetVar85 sets Var85 field to given value.

### HasVar85

`func (o *InlineResponse20017) HasVar85() bool`

HasVar85 returns a boolean if a field has been set.

### GetVar86

`func (o *InlineResponse20017) GetVar86() string`

GetVar86 returns the Var86 field if non-nil, zero value otherwise.

### GetVar86Ok

`func (o *InlineResponse20017) GetVar86Ok() (*string, bool)`

GetVar86Ok returns a tuple with the Var86 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar86

`func (o *InlineResponse20017) SetVar86(v string)`

SetVar86 sets Var86 field to given value.

### HasVar86

`func (o *InlineResponse20017) HasVar86() bool`

HasVar86 returns a boolean if a field has been set.

### GetVar87

`func (o *InlineResponse20017) GetVar87() string`

GetVar87 returns the Var87 field if non-nil, zero value otherwise.

### GetVar87Ok

`func (o *InlineResponse20017) GetVar87Ok() (*string, bool)`

GetVar87Ok returns a tuple with the Var87 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar87

`func (o *InlineResponse20017) SetVar87(v string)`

SetVar87 sets Var87 field to given value.

### HasVar87

`func (o *InlineResponse20017) HasVar87() bool`

HasVar87 returns a boolean if a field has been set.

### GetVar88

`func (o *InlineResponse20017) GetVar88() string`

GetVar88 returns the Var88 field if non-nil, zero value otherwise.

### GetVar88Ok

`func (o *InlineResponse20017) GetVar88Ok() (*string, bool)`

GetVar88Ok returns a tuple with the Var88 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar88

`func (o *InlineResponse20017) SetVar88(v string)`

SetVar88 sets Var88 field to given value.

### HasVar88

`func (o *InlineResponse20017) HasVar88() bool`

HasVar88 returns a boolean if a field has been set.

### GetVar6004

`func (o *InlineResponse20017) GetVar6004() string`

GetVar6004 returns the Var6004 field if non-nil, zero value otherwise.

### GetVar6004Ok

`func (o *InlineResponse20017) GetVar6004Ok() (*string, bool)`

GetVar6004Ok returns a tuple with the Var6004 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6004

`func (o *InlineResponse20017) SetVar6004(v string)`

SetVar6004 sets Var6004 field to given value.

### HasVar6004

`func (o *InlineResponse20017) HasVar6004() bool`

HasVar6004 returns a boolean if a field has been set.

### GetVar6008

`func (o *InlineResponse20017) GetVar6008() string`

GetVar6008 returns the Var6008 field if non-nil, zero value otherwise.

### GetVar6008Ok

`func (o *InlineResponse20017) GetVar6008Ok() (*string, bool)`

GetVar6008Ok returns a tuple with the Var6008 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6008

`func (o *InlineResponse20017) SetVar6008(v string)`

SetVar6008 sets Var6008 field to given value.

### HasVar6008

`func (o *InlineResponse20017) HasVar6008() bool`

HasVar6008 returns a boolean if a field has been set.

### GetVar6070

`func (o *InlineResponse20017) GetVar6070() string`

GetVar6070 returns the Var6070 field if non-nil, zero value otherwise.

### GetVar6070Ok

`func (o *InlineResponse20017) GetVar6070Ok() (*string, bool)`

GetVar6070Ok returns a tuple with the Var6070 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6070

`func (o *InlineResponse20017) SetVar6070(v string)`

SetVar6070 sets Var6070 field to given value.

### HasVar6070

`func (o *InlineResponse20017) HasVar6070() bool`

HasVar6070 returns a boolean if a field has been set.

### GetVar6072

`func (o *InlineResponse20017) GetVar6072() string`

GetVar6072 returns the Var6072 field if non-nil, zero value otherwise.

### GetVar6072Ok

`func (o *InlineResponse20017) GetVar6072Ok() (*string, bool)`

GetVar6072Ok returns a tuple with the Var6072 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6072

`func (o *InlineResponse20017) SetVar6072(v string)`

SetVar6072 sets Var6072 field to given value.

### HasVar6072

`func (o *InlineResponse20017) HasVar6072() bool`

HasVar6072 returns a boolean if a field has been set.

### GetVar6073

`func (o *InlineResponse20017) GetVar6073() string`

GetVar6073 returns the Var6073 field if non-nil, zero value otherwise.

### GetVar6073Ok

`func (o *InlineResponse20017) GetVar6073Ok() (*string, bool)`

GetVar6073Ok returns a tuple with the Var6073 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6073

`func (o *InlineResponse20017) SetVar6073(v string)`

SetVar6073 sets Var6073 field to given value.

### HasVar6073

`func (o *InlineResponse20017) HasVar6073() bool`

HasVar6073 returns a boolean if a field has been set.

### GetVar6119

`func (o *InlineResponse20017) GetVar6119() string`

GetVar6119 returns the Var6119 field if non-nil, zero value otherwise.

### GetVar6119Ok

`func (o *InlineResponse20017) GetVar6119Ok() (*string, bool)`

GetVar6119Ok returns a tuple with the Var6119 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6119

`func (o *InlineResponse20017) SetVar6119(v string)`

SetVar6119 sets Var6119 field to given value.

### HasVar6119

`func (o *InlineResponse20017) HasVar6119() bool`

HasVar6119 returns a boolean if a field has been set.

### GetVar6457

`func (o *InlineResponse20017) GetVar6457() string`

GetVar6457 returns the Var6457 field if non-nil, zero value otherwise.

### GetVar6457Ok

`func (o *InlineResponse20017) GetVar6457Ok() (*string, bool)`

GetVar6457Ok returns a tuple with the Var6457 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6457

`func (o *InlineResponse20017) SetVar6457(v string)`

SetVar6457 sets Var6457 field to given value.

### HasVar6457

`func (o *InlineResponse20017) HasVar6457() bool`

HasVar6457 returns a boolean if a field has been set.

### GetVar6509

`func (o *InlineResponse20017) GetVar6509() string`

GetVar6509 returns the Var6509 field if non-nil, zero value otherwise.

### GetVar6509Ok

`func (o *InlineResponse20017) GetVar6509Ok() (*string, bool)`

GetVar6509Ok returns a tuple with the Var6509 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6509

`func (o *InlineResponse20017) SetVar6509(v string)`

SetVar6509 sets Var6509 field to given value.

### HasVar6509

`func (o *InlineResponse20017) HasVar6509() bool`

HasVar6509 returns a boolean if a field has been set.

### GetVar7051

`func (o *InlineResponse20017) GetVar7051() string`

GetVar7051 returns the Var7051 field if non-nil, zero value otherwise.

### GetVar7051Ok

`func (o *InlineResponse20017) GetVar7051Ok() (*string, bool)`

GetVar7051Ok returns a tuple with the Var7051 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7051

`func (o *InlineResponse20017) SetVar7051(v string)`

SetVar7051 sets Var7051 field to given value.

### HasVar7051

`func (o *InlineResponse20017) HasVar7051() bool`

HasVar7051 returns a boolean if a field has been set.

### GetVar7094

`func (o *InlineResponse20017) GetVar7094() string`

GetVar7094 returns the Var7094 field if non-nil, zero value otherwise.

### GetVar7094Ok

`func (o *InlineResponse20017) GetVar7094Ok() (*string, bool)`

GetVar7094Ok returns a tuple with the Var7094 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7094

`func (o *InlineResponse20017) SetVar7094(v string)`

SetVar7094 sets Var7094 field to given value.

### HasVar7094

`func (o *InlineResponse20017) HasVar7094() bool`

HasVar7094 returns a boolean if a field has been set.

### GetVar7219

`func (o *InlineResponse20017) GetVar7219() string`

GetVar7219 returns the Var7219 field if non-nil, zero value otherwise.

### GetVar7219Ok

`func (o *InlineResponse20017) GetVar7219Ok() (*string, bool)`

GetVar7219Ok returns a tuple with the Var7219 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7219

`func (o *InlineResponse20017) SetVar7219(v string)`

SetVar7219 sets Var7219 field to given value.

### HasVar7219

`func (o *InlineResponse20017) HasVar7219() bool`

HasVar7219 returns a boolean if a field has been set.

### GetVar7220

`func (o *InlineResponse20017) GetVar7220() string`

GetVar7220 returns the Var7220 field if non-nil, zero value otherwise.

### GetVar7220Ok

`func (o *InlineResponse20017) GetVar7220Ok() (*string, bool)`

GetVar7220Ok returns a tuple with the Var7220 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7220

`func (o *InlineResponse20017) SetVar7220(v string)`

SetVar7220 sets Var7220 field to given value.

### HasVar7220

`func (o *InlineResponse20017) HasVar7220() bool`

HasVar7220 returns a boolean if a field has been set.

### GetVar7221

`func (o *InlineResponse20017) GetVar7221() string`

GetVar7221 returns the Var7221 field if non-nil, zero value otherwise.

### GetVar7221Ok

`func (o *InlineResponse20017) GetVar7221Ok() (*string, bool)`

GetVar7221Ok returns a tuple with the Var7221 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7221

`func (o *InlineResponse20017) SetVar7221(v string)`

SetVar7221 sets Var7221 field to given value.

### HasVar7221

`func (o *InlineResponse20017) HasVar7221() bool`

HasVar7221 returns a boolean if a field has been set.

### GetVar7280

`func (o *InlineResponse20017) GetVar7280() string`

GetVar7280 returns the Var7280 field if non-nil, zero value otherwise.

### GetVar7280Ok

`func (o *InlineResponse20017) GetVar7280Ok() (*string, bool)`

GetVar7280Ok returns a tuple with the Var7280 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7280

`func (o *InlineResponse20017) SetVar7280(v string)`

SetVar7280 sets Var7280 field to given value.

### HasVar7280

`func (o *InlineResponse20017) HasVar7280() bool`

HasVar7280 returns a boolean if a field has been set.

### GetVar7281

`func (o *InlineResponse20017) GetVar7281() string`

GetVar7281 returns the Var7281 field if non-nil, zero value otherwise.

### GetVar7281Ok

`func (o *InlineResponse20017) GetVar7281Ok() (*string, bool)`

GetVar7281Ok returns a tuple with the Var7281 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7281

`func (o *InlineResponse20017) SetVar7281(v string)`

SetVar7281 sets Var7281 field to given value.

### HasVar7281

`func (o *InlineResponse20017) HasVar7281() bool`

HasVar7281 returns a boolean if a field has been set.

### GetVar7282

`func (o *InlineResponse20017) GetVar7282() string`

GetVar7282 returns the Var7282 field if non-nil, zero value otherwise.

### GetVar7282Ok

`func (o *InlineResponse20017) GetVar7282Ok() (*string, bool)`

GetVar7282Ok returns a tuple with the Var7282 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7282

`func (o *InlineResponse20017) SetVar7282(v string)`

SetVar7282 sets Var7282 field to given value.

### HasVar7282

`func (o *InlineResponse20017) HasVar7282() bool`

HasVar7282 returns a boolean if a field has been set.

### GetVar7633

`func (o *InlineResponse20017) GetVar7633() string`

GetVar7633 returns the Var7633 field if non-nil, zero value otherwise.

### GetVar7633Ok

`func (o *InlineResponse20017) GetVar7633Ok() (*string, bool)`

GetVar7633Ok returns a tuple with the Var7633 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7633

`func (o *InlineResponse20017) SetVar7633(v string)`

SetVar7633 sets Var7633 field to given value.

### HasVar7633

`func (o *InlineResponse20017) HasVar7633() bool`

HasVar7633 returns a boolean if a field has been set.

### GetVar7284

`func (o *InlineResponse20017) GetVar7284() string`

GetVar7284 returns the Var7284 field if non-nil, zero value otherwise.

### GetVar7284Ok

`func (o *InlineResponse20017) GetVar7284Ok() (*string, bool)`

GetVar7284Ok returns a tuple with the Var7284 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7284

`func (o *InlineResponse20017) SetVar7284(v string)`

SetVar7284 sets Var7284 field to given value.

### HasVar7284

`func (o *InlineResponse20017) HasVar7284() bool`

HasVar7284 returns a boolean if a field has been set.

### GetVar7285

`func (o *InlineResponse20017) GetVar7285() string`

GetVar7285 returns the Var7285 field if non-nil, zero value otherwise.

### GetVar7285Ok

`func (o *InlineResponse20017) GetVar7285Ok() (*string, bool)`

GetVar7285Ok returns a tuple with the Var7285 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7285

`func (o *InlineResponse20017) SetVar7285(v string)`

SetVar7285 sets Var7285 field to given value.

### HasVar7285

`func (o *InlineResponse20017) HasVar7285() bool`

HasVar7285 returns a boolean if a field has been set.

### GetVar7286

`func (o *InlineResponse20017) GetVar7286() string`

GetVar7286 returns the Var7286 field if non-nil, zero value otherwise.

### GetVar7286Ok

`func (o *InlineResponse20017) GetVar7286Ok() (*string, bool)`

GetVar7286Ok returns a tuple with the Var7286 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7286

`func (o *InlineResponse20017) SetVar7286(v string)`

SetVar7286 sets Var7286 field to given value.

### HasVar7286

`func (o *InlineResponse20017) HasVar7286() bool`

HasVar7286 returns a boolean if a field has been set.

### GetVar7287

`func (o *InlineResponse20017) GetVar7287() string`

GetVar7287 returns the Var7287 field if non-nil, zero value otherwise.

### GetVar7287Ok

`func (o *InlineResponse20017) GetVar7287Ok() (*string, bool)`

GetVar7287Ok returns a tuple with the Var7287 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7287

`func (o *InlineResponse20017) SetVar7287(v string)`

SetVar7287 sets Var7287 field to given value.

### HasVar7287

`func (o *InlineResponse20017) HasVar7287() bool`

HasVar7287 returns a boolean if a field has been set.

### GetVar7288

`func (o *InlineResponse20017) GetVar7288() string`

GetVar7288 returns the Var7288 field if non-nil, zero value otherwise.

### GetVar7288Ok

`func (o *InlineResponse20017) GetVar7288Ok() (*string, bool)`

GetVar7288Ok returns a tuple with the Var7288 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7288

`func (o *InlineResponse20017) SetVar7288(v string)`

SetVar7288 sets Var7288 field to given value.

### HasVar7288

`func (o *InlineResponse20017) HasVar7288() bool`

HasVar7288 returns a boolean if a field has been set.

### GetVar7289

`func (o *InlineResponse20017) GetVar7289() string`

GetVar7289 returns the Var7289 field if non-nil, zero value otherwise.

### GetVar7289Ok

`func (o *InlineResponse20017) GetVar7289Ok() (*string, bool)`

GetVar7289Ok returns a tuple with the Var7289 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7289

`func (o *InlineResponse20017) SetVar7289(v string)`

SetVar7289 sets Var7289 field to given value.

### HasVar7289

`func (o *InlineResponse20017) HasVar7289() bool`

HasVar7289 returns a boolean if a field has been set.

### GetVar7290

`func (o *InlineResponse20017) GetVar7290() string`

GetVar7290 returns the Var7290 field if non-nil, zero value otherwise.

### GetVar7290Ok

`func (o *InlineResponse20017) GetVar7290Ok() (*string, bool)`

GetVar7290Ok returns a tuple with the Var7290 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7290

`func (o *InlineResponse20017) SetVar7290(v string)`

SetVar7290 sets Var7290 field to given value.

### HasVar7290

`func (o *InlineResponse20017) HasVar7290() bool`

HasVar7290 returns a boolean if a field has been set.

### GetVar7291

`func (o *InlineResponse20017) GetVar7291() string`

GetVar7291 returns the Var7291 field if non-nil, zero value otherwise.

### GetVar7291Ok

`func (o *InlineResponse20017) GetVar7291Ok() (*string, bool)`

GetVar7291Ok returns a tuple with the Var7291 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7291

`func (o *InlineResponse20017) SetVar7291(v string)`

SetVar7291 sets Var7291 field to given value.

### HasVar7291

`func (o *InlineResponse20017) HasVar7291() bool`

HasVar7291 returns a boolean if a field has been set.

### GetVar7292

`func (o *InlineResponse20017) GetVar7292() string`

GetVar7292 returns the Var7292 field if non-nil, zero value otherwise.

### GetVar7292Ok

`func (o *InlineResponse20017) GetVar7292Ok() (*string, bool)`

GetVar7292Ok returns a tuple with the Var7292 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7292

`func (o *InlineResponse20017) SetVar7292(v string)`

SetVar7292 sets Var7292 field to given value.

### HasVar7292

`func (o *InlineResponse20017) HasVar7292() bool`

HasVar7292 returns a boolean if a field has been set.

### GetVar7293

`func (o *InlineResponse20017) GetVar7293() string`

GetVar7293 returns the Var7293 field if non-nil, zero value otherwise.

### GetVar7293Ok

`func (o *InlineResponse20017) GetVar7293Ok() (*string, bool)`

GetVar7293Ok returns a tuple with the Var7293 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7293

`func (o *InlineResponse20017) SetVar7293(v string)`

SetVar7293 sets Var7293 field to given value.

### HasVar7293

`func (o *InlineResponse20017) HasVar7293() bool`

HasVar7293 returns a boolean if a field has been set.

### GetVar7294

`func (o *InlineResponse20017) GetVar7294() string`

GetVar7294 returns the Var7294 field if non-nil, zero value otherwise.

### GetVar7294Ok

`func (o *InlineResponse20017) GetVar7294Ok() (*string, bool)`

GetVar7294Ok returns a tuple with the Var7294 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7294

`func (o *InlineResponse20017) SetVar7294(v string)`

SetVar7294 sets Var7294 field to given value.

### HasVar7294

`func (o *InlineResponse20017) HasVar7294() bool`

HasVar7294 returns a boolean if a field has been set.

### GetVar7295

`func (o *InlineResponse20017) GetVar7295() string`

GetVar7295 returns the Var7295 field if non-nil, zero value otherwise.

### GetVar7295Ok

`func (o *InlineResponse20017) GetVar7295Ok() (*string, bool)`

GetVar7295Ok returns a tuple with the Var7295 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7295

`func (o *InlineResponse20017) SetVar7295(v string)`

SetVar7295 sets Var7295 field to given value.

### HasVar7295

`func (o *InlineResponse20017) HasVar7295() bool`

HasVar7295 returns a boolean if a field has been set.

### GetVar7296

`func (o *InlineResponse20017) GetVar7296() string`

GetVar7296 returns the Var7296 field if non-nil, zero value otherwise.

### GetVar7296Ok

`func (o *InlineResponse20017) GetVar7296Ok() (*string, bool)`

GetVar7296Ok returns a tuple with the Var7296 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7296

`func (o *InlineResponse20017) SetVar7296(v string)`

SetVar7296 sets Var7296 field to given value.

### HasVar7296

`func (o *InlineResponse20017) HasVar7296() bool`

HasVar7296 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


