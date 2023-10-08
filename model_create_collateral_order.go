/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type CreateCollateralOrder struct {
	// 质押数量
	CollateralAmount string `json:"collateral_amount"`
	// 质押币种
	CollateralCurrency string `json:"collateral_currency"`
	// 借款数量
	BorrowAmount string `json:"borrow_amount"`
	// 借款币种
	BorrowCurrency string `json:"borrow_currency"`
}
