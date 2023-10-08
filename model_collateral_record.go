/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// 质押物记录
type CollateralRecord struct {
	// 订单id
	OrderId int64 `json:"order_id,omitempty"`
	// 质押物记录 id
	RecordId int64 `json:"record_id,omitempty"`
	// 借款币种
	BorrowCurrency string `json:"borrow_currency,omitempty"`
	// 借款数量
	BorrowAmount string `json:"borrow_amount,omitempty"`
	// 质押币种
	CollateralCurrency string `json:"collateral_currency,omitempty"`
	// 调整前质押数量
	BeforeCollateral string `json:"before_collateral,omitempty"`
	// 调整后质押数量
	AfterCollateral string `json:"after_collateral,omitempty"`
	// 调整前质押率
	BeforeLtv string `json:"before_ltv,omitempty"`
	// 调整后质押率
	AfterLtv string `json:"after_ltv,omitempty"`
	// 操作时间，时间戳，单位秒
	OperateTime int64 `json:"operate_time,omitempty"`
}
