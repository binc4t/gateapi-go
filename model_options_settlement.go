/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type OptionsSettlement struct {
	// Last changed time of configuration
	Time float64 `json:"time,omitempty"`
	// Options contract name
	Contract string `json:"contract,omitempty"`
	// Settlement profit per size (quote currency)
	Profit string `json:"profit,omitempty"`
	// Settlement fee per size (quote currency)
	Fee string `json:"fee,omitempty"`
	// Strike price (quote currency)
	StrikePrice string `json:"strike_price,omitempty"`
	// Settlement price (quote currency)
	SettlePrice string `json:"settle_price,omitempty"`
}