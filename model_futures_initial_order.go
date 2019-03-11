/*
 * Gate API v4
 *
 * APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * API version: 4.5.1
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type FuturesInitialOrder struct {
	// Futures contract
	Contract string `json:"contract"`
	// Order size. Positive size means to buy, while negative one means to sell. Set to 0 to close the position
	Size int64 `json:"size,omitempty"`
	// Order price. Set to 0 to use market price
	Price string `json:"price"`
	// Set to true if trying to close the position
	Close bool `json:"close,omitempty"`
	// Time in force. If using market price, only `ioc` is supported.  - gtc: GoodTillCancelled - ioc: ImmediateOrCancelled
	Tif string `json:"tif,omitempty"`
	// How the order is created. Possible values are: web, api and app
	Text string `json:"text,omitempty"`
	// Set to true to create an post-only order
	ReduceOnly bool `json:"reduce_only,omitempty"`
	// Is the order post-only
	IsReduceOnly bool `json:"is_reduce_only,omitempty"`
	// Is the order to close position
	IsClose bool `json:"is_close,omitempty"`
}