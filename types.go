package kraken

import (
	"time"
)

type Status string

const (
	Online      Status = "online"
	Maintenance Status = "maintenance"
	CancelOnly  Status = "cancel_only"
	PostOnly    Status = "post_only"
)

type AssetClass string

const (
	Currency AssetClass = "currency"
)

type ServerTime struct {
	// Unixtime: Unix timestamp
	Unixtime int64 `json:"unixtime"`
	// RFC1123: RFC 1123 time format
	RFC1123 string `json:"rfc1123"`
}

type SystemStatus struct {
	// Status: Current system status
	Status Status `json:"status"`
	// Timestamp: Current timestamp (RFC3339)
	Timestamp *time.Time `json:"timestamp"`
}

type AssetInfo struct {
	// Altname: Alternative name
	Altname string `json:"altname"`
	// AssetClass: Asset class
	AssetClass AssetClass `json:"aclass"`
	// Decimals: Scaling decimal places for record keeping
	Decimals int64 `json:"decimals"`
	// DisplayDecimals: Scaling decimal places for output display
	DisplayDecimals int64 `json:"display_decimals"`
}
