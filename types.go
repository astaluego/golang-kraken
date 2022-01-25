package kraken

import "net/url"

type Payload url.Values

type ServerTime struct {
	// Unixtime: Unix timestamp
	Unixtime int64 `json:"unixtime"`
	// RFC1123: RFC 1123 time format
	RFC1123 string `json:"rfc1123"`
}
