/*
 * Client Portal Web API
 *
 * Production version of the Client Portal Web API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineResponse20017 struct for InlineResponse20017
type InlineResponse20017 struct {
	ServerId string `json:"server_id,omitempty"`
	Conid int32 `json:"conid,omitempty"`
	Updated int32 `json:"_updated,omitempty"`
	// Last Price
	Var31 string `json:"31,omitempty"`
	// Symbol
	Var55 string `json:"55,omitempty"`
	// Text
	Var58 string `json:"58,omitempty"`
	// High
	Var70 string `json:"70,omitempty"`
	// Low
	Var71 string `json:"71,omitempty"`
	// Position
	Var72 string `json:"72,omitempty"`
	// Market Value
	Var73 string `json:"73,omitempty"`
	// Average Price
	Var74 string `json:"74,omitempty"`
	// Unrealized PnL
	Var75 string `json:"75,omitempty"`
	// Formatted position
	Var76 string `json:"76,omitempty"`
	// Formatted Unrealized PnL
	Var77 string `json:"77,omitempty"`
	// Daily PnL
	Var78 string `json:"78,omitempty"`
	// Change Price
	Var82 string `json:"82,omitempty"`
	// Change Percent
	Var83 string `json:"83,omitempty"`
	// Bid Price
	Var84 string `json:"84,omitempty"`
	// Ask Size
	Var85 string `json:"85,omitempty"`
	// Ask Price
	Var86 string `json:"86,omitempty"`
	// Volume
	Var87 string `json:"87,omitempty"`
	// Bid Size
	Var88 string `json:"88,omitempty"`
	// Exchange
	Var6004 string `json:"6004,omitempty"`
	// Conid
	Var6008 string `json:"6008,omitempty"`
	// Security Type
	Var6070 string `json:"6070,omitempty"`
	// Months
	Var6072 string `json:"6072,omitempty"`
	// Regular Expiry
	Var6073 string `json:"6073,omitempty"`
	// Marker for market data delivery method (similar to request id)
	Var6119 string `json:"6119,omitempty"`
	// Underlying Conid. Use /trsrv/secdef to get more information about the security
	Var6457 string `json:"6457,omitempty"`
	// Market Data Availability. The field may contain two chars. The first char is the primary code: R = Realtime, D = Delayed, Z = Frozen, Y = Frozen Delayed. The second char is the secondary code: P = Snapshot Available, p = Consolidated. 
	Var6509 string `json:"6509,omitempty"`
	// Company name
	Var7051 string `json:"7051,omitempty"`
	// Conid + Exchange
	Var7094 string `json:"7094,omitempty"`
	// Contract Description
	Var7219 string `json:"7219,omitempty"`
	// Contract Description
	Var7220 string `json:"7220,omitempty"`
	// Listing Exchange
	Var7221 string `json:"7221,omitempty"`
	// Industry
	Var7280 string `json:"7280,omitempty"`
	// Category
	Var7281 string `json:"7281,omitempty"`
	// Average Daily Volume
	Var7282 string `json:"7282,omitempty"`
	// Implied volatility of the option
	Var7633 string `json:"7633,omitempty"`
	// Historic Volume (30d)
	Var7284 string `json:"7284,omitempty"`
	// Put/Call Ratio
	Var7285 string `json:"7285,omitempty"`
	// Dividend Amount
	Var7286 string `json:"7286,omitempty"`
	// Dividend Yield %
	Var7287 string `json:"7287,omitempty"`
	// Ex-date of the dividend
	Var7288 string `json:"7288,omitempty"`
	// Market Cap
	Var7289 string `json:"7289,omitempty"`
	// P/E
	Var7290 string `json:"7290,omitempty"`
	// EPS
	Var7291 string `json:"7291,omitempty"`
	// Cost Basis
	Var7292 string `json:"7292,omitempty"`
	// 52 Week High
	Var7293 string `json:"7293,omitempty"`
	// 52 Week Low
	Var7294 string `json:"7294,omitempty"`
	// Open Price
	Var7295 string `json:"7295,omitempty"`
	// Close Price
	Var7296 string `json:"7296,omitempty"`
}
