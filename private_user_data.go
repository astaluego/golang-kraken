package kraken

import (
	"fmt"
	"net/url"
)

// AccountBalance
// Retrieve all cash balances, net of pending withdrawals.
// https://docs.kraken.com/rest/#tag/User-Data/operation/getAccountBalance
func (c *Client) AccountBalance() (*AccountBalance, error) {
	payload := Payload{}

	response := AccountBalance{}
	err := c.doRequest("Balance", true, url.Values(payload), &response)
	return &response, err
}

type TradeBalanceConfig struct {
	// Asset is required
	// Base asset used to determine balance
	// Default: USD
	Asset Asset
}

// TradeBalance
// Retrieve a summary of collateral balances, margin position valuations, equity and margin level.
// https://docs.kraken.com/rest/#tag/User-Data/operation/getTradeBalance
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

type OpenOrdersConfig struct {
	// Trades is optional
	// Whether or not to include trades related to position in output
	Trades bool

	// UserReferenceID is optional
	// Restrict results to given user reference id
	UserReferenceID int64
}

// OpenOrders
// Retrieve information about currently open orders.
// https://docs.kraken.com/rest/#tag/User-Data/operation/getTradeBalance
func (c *Client) OpenOrders(config OpenOrdersConfig) (*map[OrderStatus]map[string]Order, error) {

	payload := Payload{}
	payload.OptWithTrades(config.Trades)
	payload.OptWithUserReferenceID(config.UserReferenceID)

	response := make(map[OrderStatus]map[string]Order)
	err := c.doRequest("OpenOrders", true, url.Values(payload), &response)

	return &response, err
}
