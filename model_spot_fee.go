/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type SpotFee struct {
	// User ID
	UserId int64 `json:"user_id,omitempty"`
	// taker fee rate
	TakerFee string `json:"taker_fee,omitempty"`
	// maker fee rate
	MakerFee string `json:"maker_fee,omitempty"`
	// If GT deduction is enabled
	GtDiscount bool `json:"gt_discount,omitempty"`
	// Taker fee rate if using GT deduction. It will be 0 if GT deduction is disabled
	GtTakerFee string `json:"gt_taker_fee,omitempty"`
	// Maker fee rate if using GT deduction. It will be 0 if GT deduction is disabled
	GtMakerFee string `json:"gt_maker_fee,omitempty"`
	// Loan fee rate of margin lending
	LoanFee string `json:"loan_fee,omitempty"`
	// Point type. 0 - Initial version. 1 - new version since 202009
	PointType string `json:"point_type,omitempty"`
	// Currency pair
	CurrencyPair string `json:"currency_pair,omitempty"`
}
