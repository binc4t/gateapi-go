/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type CrossMarginRepayment struct {
	// Loan record ID
	Id string `json:"id,omitempty"`
	// Repayment time
	CreateTime int64 `json:"create_time,omitempty"`
	// Loan record ID
	LoanId string `json:"loan_id,omitempty"`
	// Currency name
	Currency string `json:"currency,omitempty"`
	// Repaid principal
	Principal string `json:"principal,omitempty"`
	// Repaid interest
	Interest string `json:"interest,omitempty"`
	// 还款类型 , none - 无还款类型, manual_repay - 手动还款 , auto_repay - 自动还款, cancel_auto_repay - 撤单后自动还款
	RepaymentType string `json:"repayment_type,omitempty"`
}
