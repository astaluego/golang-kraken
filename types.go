package kraken

import (
	"net/url"
	"time"
)

type Payload url.Values

type Status string

const (
	Online      Status = "online"
	Maintenance Status = "maintenance"
	CancelOnly  Status = "cancel_only"
	PostOnly    Status = "post_only"
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
