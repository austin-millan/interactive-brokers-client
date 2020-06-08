# Order

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
**AuxPrice** | **float32** | Required Price to support Stop and Stop Limit orders | [optional] 
**TimeInForce** | [**TimeInForce**](timeInForce.md) |  | [optional] 
**OutsideRTH** | **float32** | Indicates if order is active outside regular trading hours, where defined. 0 &#x3D; no (default), 1 &#x3D; yes | [optional] 
**Side** | **float32** | Buy &#x3D; &#39;1&#39;, Sell &#x3D; &#39;2&#39; | [optional] 
**OrderRestrictions** | **float32** | MultiValueString representing the restrictions associated with an order. If more than one restriction is applicable to an order, this field can contain multiple instructions separated by space.  &#39;1&#39; Program Trade &#39;2&#39; Index Arbitrage  &#39;3&#39; Non-Index Arbitrage  | [optional] 
**GermanHftAlgo** | **bool** | By setting this bool to false the customer attests that the order is not subject to German HFT Act, was not generated using any automated algorithm, and no algorithm determined or changed financial instrument, side, quantity, order type, limit or other price, trading venue or timing of this order. Currently we cannot accept orders where this flag is set to true. Such orders will be rejected. | [optional] 
**Mifid2DecisionMaker** | **string** | This field permits specification of the user&#39;s preregistered (via account management) MiFID II short code for decision makers. | [optional] 
**Mifid2Algo** | **string** | This field permits specification of the user&#39;s preregistered (via account management) MiFID II short code for algos that are responsible for investment decisions | [optional] 
**Mifid2ExecutionTrader** | **string** | This field permits specification of the user&#39;s preregistered (via account management) MiFID II person responsible for handling/routing of the order | [optional] 
**Mifid2ExecutionAlgo** | **string** | This field permits specification of the user&#39;s preregistered (via account management) MiFID II short code for algos that are responsible for handling/routing of the order. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


