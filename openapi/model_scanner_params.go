/*
 * Client Portal Web API
 *
 * Production version of the Client Portal Web API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ScannerParams struct for ScannerParams
type ScannerParams struct {
	// for example-STK
	Instrument string `json:"instrument,omitempty"`
	// for example-TOP_PERC_GAIN
	Type string `json:"type,omitempty"`
	Filter []ScannerParamsFilter `json:"filter,omitempty"`
	Location string `json:"location,omitempty"`
	Size string `json:"size,omitempty"`
}