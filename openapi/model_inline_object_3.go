/*
 * IBKR 3rd Party Web API
 *
 * Interactive Brokers Web API for 3rd Party Companies
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject3 struct for InlineObject3
type InlineObject3 struct {
	// The order ID assigned by the customer.
	CustomerOrderId string `json:"CustomerOrderId,omitempty"`
	// The symbol that identifies the trading product
	Ticker string `json:"Ticker,omitempty"`
	// The exchange on which the trading product is listed (only for InstrumentType=STK)
	ListingExchange string `json:"ListingExchange,omitempty"`
	// The currency in which the FX pair trades (only for InstrumentType=CASH)
	Currency string `json:"Currency,omitempty"`
	// The instrument type of the contract
	InstrumentType string `json:"InstrumentType,omitempty"`
	// The internal IB identifier for the trading product specified as an integer (can be obtained in response to /secdef request)
	ContractId float32 `json:"ContractId,omitempty"`
	// The number of units in the order; contracts or shares
	Quantity float32 `json:"Quantity,omitempty"`
	// The order price
	Price float32 `json:"Price,omitempty"`
	OrderType OrderType `json:"Order Type,omitempty"`
	// Required price to support Stop and Stop Limit orders
	AuxPrice float32 `json:"Aux Price,omitempty"`
	TimeInForce TimeInForce `json:"Time in Force,omitempty"`
	// Buy = '1', Sell = '2'
	Side float32 `json:"Side,omitempty"`
}
