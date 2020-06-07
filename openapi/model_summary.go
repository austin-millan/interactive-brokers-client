/*
 * Client Portal Web API
 *
 * Production version of the Client Portal Web API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Summary account information
type Summary struct {
	Total SummaryTotal `json:"total,omitempty"`
	// date format-- yyyy-MM-dd
	StartDate string `json:"startDate,omitempty"`
	ExcludedAccounts []SummaryExcludedAccounts `json:"excludedAccounts,omitempty"`
	LastSuccessfulUpdate string `json:"lastSuccessfulUpdate,omitempty"`
	AccountSummaries []SummaryAccountSummaries `json:"accountSummaries,omitempty"`
	EndDate string `json:"endDate,omitempty"`
	// indicator of user having configured any external accounts
	HasExternalAccounts bool `json:"hasExternalAccounts,omitempty"`
	Rc int32 `json:"rc,omitempty"`
	Currency string `json:"currency,omitempty"`
	UserId string `json:"userId,omitempty"`
	Pm string `json:"pm,omitempty"`
	View string `json:"view,omitempty"`
	BalanceByDate SummaryBalanceByDate `json:"balanceByDate,omitempty"`
}