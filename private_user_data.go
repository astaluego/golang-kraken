package kraken

import "net/url"

// AccountBalance
// Retrieve all cash balances, net of pending withdrawals.
// https://docs.kraken.com/rest/#operation/getAccountBalance
func (c *Client) AccountBalance() (*AccountBalance, error) {
	payload := Payload{}

	response := AccountBalance{}
	err := c.doRequest("Balance", true, url.Values(payload), &response)
	return &response, err
}
