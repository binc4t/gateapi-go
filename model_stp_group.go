/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type StpGroup struct {
	// STP Group ID
	Id int64 `json:"id,omitempty"`
	// STP Group name
	Name string `json:"name"`
	// Creator ID
	CreatorId int64 `json:"creator_id,omitempty"`
	// Creation time
	CreateTime int64 `json:"create_time,omitempty"`
}