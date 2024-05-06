/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type UnifiedSettings struct {
	// USDT contract switch. This parameter is required when the mode is multi-currency margin mode
	UsdtFutures bool `json:"usdt_futures,omitempty"`
	// Spot hedging switch. This parameter is required when the mode is portfolio margin mode
	SpotHedge bool `json:"spot_hedge,omitempty"`
}