package kraken

import (
	"fmt"
	"net/url"
)

// AccountBalance
// Retrieve all cash balances, net of pending withdrawals.
// https://docs.kraken.com/rest/#operation/getAccountBalance
func (c *Client) AccountBalance() (*AccountBalance, error) {
	payload := Payload{}

	response := AccountBalance{}
	err := c.doRequest("Balance", true, url.Values(payload), &response)
	return &response, err
}

type TradeBalanceConfig struct {
	// Asset is optional
	// Base asset used to determine balance
	// Default: USD
	Asset Asset
}

// TradeBalance
// Retrieve a summary of collateral balances, margin position valuations, equity and margin level.
// https://docs.kraken.com/rest/#operation/getTradeBalance
func (c *Client) TradeBalance(config TradeBalanceConfig) (*TradeBalance, error) {
	if config.Asset == "" {
		return nil, fmt.Errorf("Asset is required")
	}

	payload := Payload{}
	payload.OptAssets(config.Asset)

	response := TradeBalance{}
	err := c.doRequest("TradeBalance", true, url.Values(payload), &response)
	return &response, err
}
