package kraken

import (
	"fmt"
	"math"
	"net/url"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
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

// SystemStatus
// Get the last system status or trading mode.
// https://docs.kraken.com/rest/#operation/getSystemStatus
func (c *Client) SystemStatus() (*SystemStatus, error) {
	payload := Payload{}

	response := SystemStatus{}
	err := c.doRequest("SystemStatus", false, url.Values(payload), &response)
	return &response, err
}

type AssetsConfig struct {
	// AssetClass is optional
	// Default: currency
	AssetClass AssetClass
	// Assets is optional
	Assets []Asset
}

// Assets
// Get information about the assets that are available for deposit, withdrawal, trading and staking.
// https://docs.kraken.com/rest/#operation/getAssetInfo
func (c *Client) Assets(config AssetsConfig) (map[Asset]AssetInfo, error) {
	payload := Payload{}
	payload.OptAssetClass(config.AssetClass)
	payload.OptAssets(config.Assets...)

	response := make(map[Asset]AssetInfo)
	err := c.doRequest("Assets", false, url.Values(payload), &response)
	return response, err
}

type AssetPairsConfig struct {
	// AssetPairs is optional
	AssetPairs []AssetPair
	// Information is optional
	// Default: info
	Information Information
}

// AssetPairs
// Get tradable asset pairs
// https://docs.kraken.com/rest/#operation/getTradableAssetPairs
func (c *Client) AssetPairs(config AssetPairsConfig) (map[AssetPair]AssetPairsInfo, error) {
	payload := Payload{}
	payload.OptAssetPairs(config.AssetPairs...)
	payload.OptInformations(config.Information)

	var response map[AssetPair]AssetPairsInfo
	err := c.doRequest("AssetPairs", false, url.Values(payload), &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type TickerInformationConfig struct {
	// AssetPairs is optional
	AssetPairs []AssetPair
}

// TickerInformation
// Today's prices start at midnight UTC. Leaving the pair parameter blank will return tickers for all tradeable
// assets on Kraken.
// https://docs.kraken.com/rest/#tag/Market-Data/operation/getTradableAssetPairs
func (c *Client) TickerInformation(config TickerInformationConfig) (map[AssetPair]AssetTickerInfo, error) {
	payload := Payload{}
	payload.OptAssetPairs(config.AssetPairs...)

	var response map[AssetPair]AssetTickerInfo
	err := c.doRequest("Ticker", false, url.Values(payload), &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type OHLCConfig struct {
	// AssetPair is required
	AssetPair AssetPair
	// Interval is optional
	// Default: 1min
	Interval Interval
	// Since is optional
	Since time.Time
}

// OHLC
// Means "Open-high-low-close chart"
// Note: the last entry in the OHLC array is for the last, not-yet-committed frame and will always
// be present, regardless of the value of `since`.
// https://docs.kraken.com/rest/#operation/getOHLCData
func (c *Client) OHLC(config OHLCConfig) ([]OHLCData, time.Time, error) {
	if config.AssetPair == "" {
		return nil, time.Time{}, fmt.Errorf("AssetPair is required")
	}

	payload := Payload{}
	payload.OptAssetPairs(config.AssetPair)
	payload.OptInterval(config.Interval)
	payload.OptSince(config.Since)

	var resp interface{}
	err := c.doRequest("OHLC", false, url.Values(payload), &resp)
	if err != nil {
		return nil, time.Time{}, err
	}

	var last time.Time
	response := []OHLCData{}

	for key, value := range resp.(map[string]interface{}) {
		if key == "last" {
			last = time.Unix(int64(value.(float64)), 0)
			continue
		}
		for _, array := range value.([]interface{}) {
			a := array.([]interface{})

			if len(a) == 8 {
				t := time.Unix(int64(a[0].(float64)), 0)

				open, err := decimal.NewFromString(a[1].(string))
				if err != nil {
					continue
				}

				high, err := decimal.NewFromString(a[2].(string))
				if err != nil {
					continue
				}

				low, err := decimal.NewFromString(a[3].(string))
				if err != nil {
					continue
				}

				close, err := decimal.NewFromString(a[4].(string))
				if err != nil {
					continue
				}

				vwap, err := decimal.NewFromString(a[5].(string))
				if err != nil {
					continue
				}

				count, err := decimal.NewFromString(a[6].(string))
				if err != nil {
					continue
				}

				ohlcData := OHLCData{
					Time:   t,
					Open:   open,
					High:   high,
					Low:    low,
					Close:  close,
					VWAP:   vwap,
					Volume: count,
					Count:  int64(a[7].(float64)),
				}
				response = append(response, ohlcData)
			}
		}
	}

	return response, last, nil
}

type OrderBookConfig struct {
	// AssetPair is required
	AssetPair AssetPair
	// Count is optional
	// Default: 100
	Count int64
}

// OrderBook
// https://docs.kraken.com/rest/#operation/getOrderBook
func (c *Client) OrderBook(config OrderBookConfig) (*OrderBook, error) {
	if config.AssetPair == "" {
		return nil, fmt.Errorf("AssetPair is required")
	}

	payload := Payload{}
	payload.OptAssetPairs(config.AssetPair)
	payload.OptCount(config.Count)

	var resp interface{}
	err := c.doRequest("Depth", false, url.Values(payload), &resp)
	if err != nil {
		return nil, err
	}

	response := OrderBook{}

	for _, value := range resp.(map[string]interface{}) {
		for side, array := range value.(map[string]interface{}) {
			orderBookEntries := []OrderBookEntry{}
			for _, a := range array.([]interface{}) {
				obe := a.([]interface{})

				if len(obe) == 3 {
					t := time.Unix(int64(obe[2].(float64)), 0)

					price, err := decimal.NewFromString(obe[0].(string))
					if err != nil {
						continue
					}

					volume, err := decimal.NewFromString(obe[1].(string))
					if err != nil {
						continue
					}

					orderBookEntry := OrderBookEntry{
						Price:  price,
						Volume: volume,
						Time:   t,
					}
					orderBookEntries = append(orderBookEntries, orderBookEntry)
				}
			}

			if side == "asks" {
				response.Asks = orderBookEntries
			} else if side == "bids" {
				response.Bids = orderBookEntries
			}
		}
		break
	}

	return &response, nil
}

type RecentTradesConfig struct {
	// AssetPair is required
	AssetPair AssetPair
	// Since is optional
	Since time.Time
}

// RecentTrades
// Returns the last 1000 trades by default
// https://docs.kraken.com/rest/#operation/getRecentTrades
func (c *Client) RecentTrades(config RecentTradesConfig) ([]TradeData, time.Time, error) {
	if config.AssetPair == "" {
		return nil, time.Time{}, fmt.Errorf("AssetPair is required")
	}

	payload := Payload{}
	payload.OptAssetPairs(config.AssetPair)
	payload.OptSince(config.Since)

	var resp interface{}
	err := c.doRequest("Trades", false, url.Values(payload), &resp)
	if err != nil {
		return nil, time.Time{}, err
	}

	var last time.Time
	response := []TradeData{}

	for key, value := range resp.(map[string]interface{}) {
		if key == "last" {
			i, err := strconv.ParseInt(value.(string), 10, 64)
			if err != nil {
				continue
			}
			last = time.Unix(0, i)
			continue
		}
		for _, array := range value.([]interface{}) {
			a := array.([]interface{})

			if len(a) == 7 {
				price, err := decimal.NewFromString(a[0].(string))
				if err != nil {
					continue
				}

				volume, err := decimal.NewFromString(a[1].(string))
				if err != nil {
					continue
				}

				integer, decimal := math.Modf(a[2].(float64))
				t := time.Unix(int64(integer), int64(decimal*10000))

				var tradeType Type
				if a[3].(string) == "s" {
					tradeType = Sell
				} else if a[3].(string) == "b" {
					tradeType = Buy
				}

				var orderType OrderType
				if a[4].(string) == "m" {
					orderType = Market
				} else if a[4].(string) == "l" {
					orderType = Limit
				}

				tradeData := TradeData{
					Price:         price,
					Volume:        volume,
					Time:          t,
					Type:          tradeType,
					OrderType:     orderType,
					Miscellaneous: a[5].(string),
					TradeID:       int64(a[6].(float64)),
				}
				response = append(response, tradeData)
			}
		}
	}
	return response, last, err
}

type RecentSpreadsConfig struct {
	// AssetPair is required
	AssetPair AssetPair
	// Since is optional
	Since time.Time
}

// RecentSpreads
// Returns the last 1000 trades by default
// https://docs.kraken.com/rest/#operation/getRecentSpreads
func (c *Client) RecentSpreads(config RecentSpreadsConfig) ([]SpreadData, time.Time, error) {
	if config.AssetPair == "" {
		return nil, time.Time{}, fmt.Errorf("AssetPair is required")
	}

	payload := Payload{}
	payload.OptAssetPairs(config.AssetPair)
	payload.OptSince(config.Since)

	var resp interface{}
	err := c.doRequest("Spread", false, url.Values(payload), &resp)
	if err != nil {
		return nil, time.Time{}, err
	}

	var last time.Time
	response := []SpreadData{}

	for key, value := range resp.(map[string]interface{}) {
		if key == "last" {
			last = time.Unix(int64(value.(float64)), 0)
			continue
		}
		for _, array := range value.([]interface{}) {
			a := array.([]interface{})
			if len(a) == 3 {
				t := time.Unix(int64(a[0].(float64)), 0)

				bid, err := decimal.NewFromString(a[1].(string))
				if err != nil {
					continue
				}

				ask, err := decimal.NewFromString(a[2].(string))
				if err != nil {
					continue
				}

				spreadData := SpreadData{
					Time: t,
					Bid:  bid,
					Ask:  ask,
				}
				response = append(response, spreadData)
			}
		}
	}
	return response, last, nil
}
