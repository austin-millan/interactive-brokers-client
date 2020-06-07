/*
 * Client Portal Web API
 *
 * Production version of the Client Portal Web API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Order contains all the order related info
type Order struct {
	// account id
	Acct string `json:"acct,omitempty"`
	Conid int32 `json:"conid,omitempty"`
	OrderDesc string `json:"orderDesc,omitempty"`
	Description1 string `json:"description1,omitempty"`
	// for exmple FB
	Ticker string `json:"ticker,omitempty"`
	// for example STK
	SecType string `json:"secType,omitempty"`
	// for example NASDAQ.NMS
	ListingExchange string `json:"listingExchange,omitempty"`
	RemainingQuantity string `json:"remainingQuantity,omitempty"`
	FilledQuantity string `json:"filledQuantity,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
	// PendingSubmit - Indicates the order was sent, but confirmation has not been received that it has been received by the destination.                  Occurs most commonly if an exchange is closed. PendingCancel - Indicates that a request has been sent to cancel an order but confirmation has not been received of its cancellation. PreSubmitted - Indicates that a simulated order type has been accepted by the IBKR system and that this order has yet to be elected.                 The order is held in the IBKR system until the election criteria are met. At that time the order is transmitted to the order destination as specified.  Submitted - Indicates that the order has been accepted at the order destination and is working. Cancelled - Indicates that the balance of the order has been confirmed cancelled by the IB system.              This could occur unexpectedly when IB or the destination has rejected the order.   Filled - Indicates that the order has been completely filled.  Inactive - Indicates the order is not working, for instance if the order was invalid and triggered an error message,            or if the order was to short a security and shares have not yet been located.  
	Status string `json:"status,omitempty"`
	// for example Limit
	OrigOrderType string `json:"origOrderType,omitempty"`
	// BUY or SELL
	Side string `json:"side,omitempty"`
	Price float32 `json:"price,omitempty"`
	// back-ground color
	BgColor string `json:"bgColor,omitempty"`
	FgColor string `json:"fgColor,omitempty"`
	OrderId int32 `json:"orderId,omitempty"`
	// Only exists in child order of bracket
	ParentId int32 `json:"parentId,omitempty"`
}