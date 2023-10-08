/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type PortfolioAccount struct {
	// User ID
	UserId int64 `json:"user_id,omitempty"`
	// Time of the most recent refresh
	RefreshTime int64 `json:"refresh_time,omitempty"`
	// Whether account is locked
	Locked   bool                              `json:"locked,omitempty"`
	Balances map[string]PortfolioMarginBalance `json:"balances,omitempty"`
	// Total account value in USDT, i.e., the sum of all currencies'
	Total string `json:"total,omitempty"`
	// Total borrowed value in USDT, i.e., the sum of all currencies
	Borrowed string `json:"borrowed,omitempty"`
	// Total unpaid interests in USDT, i.e., the sum of all currencies
	Interest string `json:"interest,omitempty"`
	// Total initial margin
	TotalInitialMargin string `json:"total_initial_margin,omitempty"`
	// Total margin balance
	TotalMarginBalance string `json:"total_margin_balance,omitempty"`
	// Total maintenance margin
	TotalMaintenanceMargin string `json:"total_maintenance_margin,omitempty"`
	// Total initial margin rate
	TotalInitialMarginRate string `json:"total_initial_margin_rate,omitempty"`
	// Total maintenance margin rate
	TotalMaintenanceMarginRate string `json:"total_maintenance_margin_rate,omitempty"`
	// Total available margin
	TotalAvailableMargin string `json:"total_available_margin,omitempty"`
	// Total amount of the portfolio margin account
	PortfolioMarginTotal string `json:"portfolio_margin_total,omitempty"`
	// Total liabilities of the portfolio margin account
	PortfolioMarginTotalLiab string `json:"portfolio_margin_total_liab,omitempty"`
	// Total equity of the portfolio margin account
	PortfolioMarginTotalEquity string `json:"portfolio_margin_total_equity,omitempty"`
}
