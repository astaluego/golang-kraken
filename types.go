package kraken

import (
	"time"

	"github.com/shopspring/decimal"
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

type Information string

const (
	AllInformations      Information = "info"
	LeverageInformations Information = "leverage"
	FeesInformations     Information = "fees"
	MarginInformations   Information = "margin"
)

type Interval string

const (
	Interval1min  Interval = "1"
	Interval5min  Interval = "5"
	Interval15min Interval = "15"
	Interval30min Interval = "30"
	Interval1h    Interval = "60"
	Interval4h    Interval = "240"
	Interval1d    Interval = "1440"
	Interval7d    Interval = "10080"
	Interval15d   Interval = "21600"
)

type OrderType string

const (
	Market OrderType = "market"
	Limit  OrderType = "limit"
)

type Type string

const (
	Buy  Type = "buy"
	Sell Type = "sell"
)

type ServerTime struct {
	// Unix timestamp
	Unixtime int64 `json:"unixtime"`
	// RFC 1123 time format
	RFC1123 string `json:"rfc1123"`
}

type SystemStatus struct {
	// Current system status
	Status Status `json:"status"`
	// Current timestamp (RFC3339)
	Timestamp time.Time `json:"timestamp"`
}

type AssetInfo struct {
	// Alternative name
	Altname string `json:"altname"`
	// Asset class
	AssetClass AssetClass `json:"aclass"`
	// Scaling decimal places for record keeping
	Decimals int64 `json:"decimals"`
	// Scaling decimal places for output display
	DisplayDecimals int64 `json:"display_decimals"`
}

type AssetPairsInfo struct {
	// Alternative name
	Altname string `json:"altname"`
	// WebSocket name
	WSname string `json:"wsname"`
	// AssetClass of base
	BaseAssetClass AssetClass `json:"aclass_base"`
	// Base asset
	BaseAsset Asset `json:"base"`
	// AssetClass of quote
	QuoteAssetClass AssetClass `json:"aclass_quote"`
	// Quote Asset
	QuoteAsset Asset `json:"quote"`
	// Scaling decimal places for pair
	PairDecimals int64 `json:"pair_decimals"`
	// Scaling decimal places for volume
	LotDecimals int64 `json:"lot_decimals"`
	// Amount to multiply lot volume by to get currency volume
	LotMultiplier int64 `json:"lot_multiplier"`
	// Array of leverage amounts available when buying
	LeverageBuy []int64 `json:"leverage_buy"`
	// Array of leverage amounts available when selling
	LeverageSell []int64 `json:"leverage_sell"`
	// Fees (Volume/Percent)
	Fees []Fee `json:"fees"`
	// Fees Maker (Volume/Percent)
	FeesMaker []Fee `json:"fees_maker"`
	// Volume discount currency
	FeeVolumeCurrency Asset `json:"fee_volume_currency"`
	// Margin call level
	MarginCall int64 `json:"margin_call"`
	// Stop-out/liquidation margin level
	MarginStop int64 `json:"margin_stop"`
	// Minimum order size (in terms of base currency)
	OrderMin decimal.Decimal `json:"ordermin"`
}

type Fee struct {
	// Volume
	Volume int64 `json:"volume"`
	// Percent
	Percent decimal.Decimal `json:"percent"`
}

type OHLCData struct {
	// Time
	Time time.Time `json:"time"`
	// Open
	Open decimal.Decimal `json:"open"`
	// High
	High decimal.Decimal `json:"high"`
	// Low
	Low decimal.Decimal `json:"low"`
	// Close
	Close decimal.Decimal `json:"close"`
	// Volume Weighted Average Price
	VWAP decimal.Decimal `json:"vwap"`
	// Volume
	Volume decimal.Decimal `json:"volume"`
	// Count
	Count int64 `json:"count"`
}

type OrderBook struct {
	// Asks
	Asks []OrderBookEntry `json:"asks"`
	// Bids
	Bids []OrderBookEntry `json:"bids"`
}

type OrderBookEntry struct {
	// Price
	Price decimal.Decimal `json:"price"`
	// Volume
	Volume decimal.Decimal `json:"volume"`
	// Time
	Time time.Time `json:"time"`
}

type TradeData struct {
	// Price
	Price decimal.Decimal `json:"price"`
	// Volume
	Volume decimal.Decimal `json:"volume"`
	// Time
	Time time.Time `json:"time"`
	// Type (buy or sell)
	Type Type `json:"type"`
	// OrderType (market or limit)
	OrderType OrderType `json:"order_type"`
	// Miscellaneous
	Miscellaneous string `json:"miscellanous"`
}

type SpreadData struct {
	// Time
	Time time.Time `json:"time"`
	// Bid
	Bid decimal.Decimal `json:"bid"`
	// Ask
	Ask decimal.Decimal `json:"ask"`
}

type AccountBalance struct {
	Assets
}

type TradeBalance struct {
	// Equivalent balance (combined balance of all currencies)
	EquivalentBalance decimal.Decimal `json:"eb,string"`
	// Trade balance (combined balance of all equity currencies)
	TradeBalance decimal.Decimal `json:"tb,string"`
	// Margin amount of open positions
	MarginOP decimal.Decimal `json:"m,string"`
	// Unrealized net profit/loss of open positions
	UnrealizedNetProfitLossOP decimal.Decimal `json:"n,string"`
	// Cost basis of open positions
	CostBasisOP decimal.Decimal `json:"c,string"`
	// Current floating valuation of open positions
	CurrentValuationOP decimal.Decimal `json:"v,string"`
	// Equity = trade balance + unrealized net profit/loss
	Equity decimal.Decimal `json:"e,string"`
	// FreeMargin = equity - initial margin (maximum margin available to open new positions)
	FreeMargin decimal.Decimal `json:"mf,string"`
	// MargimLevel = (equity / initial margin) * 100
	MarginLevel decimal.Decimal `json:"ml,string"`
}
