/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type UidPushWithdrawal struct {
	// Recipient UID
	ReceiveUid int64 `json:"receive_uid"`
	// Currency name
	Currency string `json:"currency"`
	// Transfer amount
	Amount string `json:"amount"`
}