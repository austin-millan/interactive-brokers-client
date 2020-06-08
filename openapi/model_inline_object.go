/*
 * IBKR 3rd Party Web API
 *
 * Interactive Brokers Web API for 3rd Party Companies
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject struct for InlineObject
type InlineObject struct {
	// The instrument type of the contract (CASH).
	Type string `json:"type,omitempty"`
	// The symbol that identifies the trading product.
	Symbol string `json:"symbol,omitempty"`
	// The currency in which the given pair trades.
	Currency string `json:"currency,omitempty"`
	// The exchange on which the trading product is listed (required for type=STK).
	Exchange string `json:"exchange,omitempty"`
	// The internal IB identifier for the trading product specified as an integer.
	Conid float32 `json:"conid,omitempty"`
}
