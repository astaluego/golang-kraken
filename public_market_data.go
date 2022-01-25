package kraken

import (
	"net/url"
)

// ServerTime
// Get the server's time.
// https://docs.kraken.com/rest/#operation/getServerTime
func (c *Client) ServerTime() (*ServerTime, error) {
	payload := Payload{}

	response := ServerTime{}
	err := c.doRequest("Time", false, url.Values(payload), &response)
	return &response, err
}
