/*
 * Client Portal Web API
 *
 * Production version of the Client Portal Web API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineResponse2004E device
type InlineResponse2004E struct {
	// device name
	NM string `json:"NM,omitempty"`
	// device id
	I string `json:"I,omitempty"`
	// unique device id
	UI string `json:"UI,omitempty"`
	// device is enabled or not 0-true, 1-false.
	A string `json:"A,omitempty"`
}
