/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// 抵押借币订单
type CollateralOrder struct {
	// 订单id
	OrderId int64 `json:"order_id,omitempty"`
	// 质押币种
	CollateralCurrency string `json:"collateral_currency,omitempty"`
	// 质押数量
	CollateralAmount string `json:"collateral_amount,omitempty"`
	// 借款币种
	BorrowCurrency string `json:"borrow_currency,omitempty"`
	// 借款数量
	BorrowAmount string `json:"borrow_amount,omitempty"`
	// 已还款数量
	RepaidAmount string `json:"repaid_amount,omitempty"`
	// 已还本金
	RepaidPrincipal string `json:"repaid_principal,omitempty"`
	// 已还利息
	RepaidInterest string `json:"repaid_interest,omitempty"`
	// 初始质押率
	InitLtv string `json:"init_ltv,omitempty"`
	// 当前质押率
	CurrentLtv string `json:"current_ltv,omitempty"`
	// 平仓质押率
	LiquidateLtv string `json:"liquidate_ltv,omitempty"`
	// 订单状态: - initial: 下单初始状态 - collateral_deducted: 扣除质押物成功 - collateral_returning: 放款失败-待退回质押物 - lent: 放款成功 - repaying: 还款中 - liquidating: 平仓中 - finished: 已完成 - closed_liquidated: 已结束-平仓还款结束
	Status string `json:"status,omitempty"`
	// 借款时间，时间戳，单位秒
	BorrowTime int64 `json:"borrow_time,omitempty"`
	// 待还本息（待还本金+待还利息）
	LeftRepayTotal string `json:"left_repay_total,omitempty"`
	// 待还本金
	LeftRepayPrincipal string `json:"left_repay_principal,omitempty"`
	// 待还利息
	LeftRepayInterest string `json:"left_repay_interest,omitempty"`
}
