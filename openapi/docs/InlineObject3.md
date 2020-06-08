# InlineObject3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerOrderId** | **string** | The order ID assigned by the customer. | [optional] 
**Ticker** | **string** | The symbol that identifies the trading product | [optional] 
**ListingExchange** | **string** | The exchange on which the trading product is listed (only for InstrumentType&#x3D;STK) | [optional] 
**Currency** | **string** | The currency in which the FX pair trades (only for InstrumentType&#x3D;CASH) | [optional] 
**InstrumentType** | **string** | The instrument type of the contract | [optional] 
**ContractId** | **float32** | The internal IB identifier for the trading product specified as an integer (can be obtained in response to /secdef request) | [optional] 
**Quantity** | **float32** | The number of units in the order; contracts or shares | [optional] 
**Price** | **float32** | The order price | [optional] 
**OrderType** | [**OrderType**](orderType.md) |  | [optional] 
**AuxPrice** | **float32** | Required price to support Stop and Stop Limit orders | [optional] 
**TimeInForce** | [**TimeInForce**](timeInForce.md) |  | [optional] 
**Side** | **float32** | Buy &#x3D; &#39;1&#39;, Sell &#x3D; &#39;2&#39; | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


