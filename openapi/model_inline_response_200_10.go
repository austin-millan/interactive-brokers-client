/*
 * Client Portal Web API
 *
 * Production version of the Client Portal Web API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineResponse20010 struct for InlineResponse20010
type InlineResponse20010 struct {
	ServerId string `json:"server_id,omitempty"`
	ColumnName string `json:"column_name,omitempty"`
	Symbol string `json:"symbol,omitempty"`
	Conidex string `json:"conidex,omitempty"`
	ConId float32 `json:"con_id,omitempty"`
	AvailableChartPeriods string `json:"available_chart_periods,omitempty"`
	CompanyName string `json:"company_name,omitempty"`
	ContractDescription1 string `json:"contract_description_1,omitempty"`
	ListingExchange string `json:"listing_exchange,omitempty"`
	SecType string `json:"sec_type,omitempty"`
}