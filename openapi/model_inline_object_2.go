/*
 * Client Portal Web API
 *
 * Production version of the Client Portal Web API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject2 struct for InlineObject2
type InlineObject2 struct {
	AcctIds []string `json:"acctIds,omitempty"`
	// Frequency of cumulative performance data points: 'D'aily, 'M'onthly,'Q'uarterly. 
	Freq string `json:"freq,omitempty"`
}
