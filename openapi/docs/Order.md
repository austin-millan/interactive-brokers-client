# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acct** | **string** | account id | [optional] 
**Conid** | **int32** |  | [optional] 
**OrderDesc** | **string** |  | [optional] 
**Description1** | **string** |  | [optional] 
**Ticker** | **string** | for exmple FB | [optional] 
**SecType** | **string** | for example STK | [optional] 
**ListingExchange** | **string** | for example NASDAQ.NMS | [optional] 
**RemainingQuantity** | **string** |  | [optional] 
**FilledQuantity** | **string** |  | [optional] 
**CompanyName** | **string** |  | [optional] 
**Status** | **string** | PendingSubmit - Indicates the order was sent, but confirmation has not been received that it has been received by the destination.                  Occurs most commonly if an exchange is closed. PendingCancel - Indicates that a request has been sent to cancel an order but confirmation has not been received of its cancellation. PreSubmitted - Indicates that a simulated order type has been accepted by the IBKR system and that this order has yet to be elected.                 The order is held in the IBKR system until the election criteria are met. At that time the order is transmitted to the order destination as specified.  Submitted - Indicates that the order has been accepted at the order destination and is working. Cancelled - Indicates that the balance of the order has been confirmed cancelled by the IB system.              This could occur unexpectedly when IB or the destination has rejected the order.   Filled - Indicates that the order has been completely filled.  Inactive - Indicates the order is not working, for instance if the order was invalid and triggered an error message,            or if the order was to short a security and shares have not yet been located.   | [optional] 
**OrigOrderType** | **string** | for example Limit | [optional] 
**Side** | **string** | BUY or SELL | [optional] 
**Price** | **float32** |  | [optional] 
**BgColor** | **string** | back-ground color | [optional] 
**FgColor** | **string** |  | [optional] 
**OrderId** | **int32** |  | [optional] 
**ParentId** | **int32** | Only exists in child order of bracket | [optional] 
**OrderRef** | **string** | User defined string used to identify the order. Value is set using \&quot;cOID\&quot; field while placing an order. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


