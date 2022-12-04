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

type AssetStatus string

const (
	AssetEnabled                    AssetStatus = "enabled"
	AssetDepositOnly                AssetStatus = "deposit_only"
	AssetWithdrawalOnly             AssetStatus = "withdrawal_only"
	AssetFundingTemporarilyDisabled AssetStatus = "funding_temporarily_disabled"
)

type AssetPairStatus string

const (
	AssetPairOnline     AssetPairStatus = "online"
	AssetPairCancelOnly AssetPairStatus = "cancel_only"
	AssetPairPostOnly   AssetPairStatus = "post_only"
	AssetPairLimitOnly  AssetPairStatus = "limit_only"
	AssetPairReduceOnly AssetPairStatus = "reduce_only"
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

type OrderStatus string

const (
	Pending  OrderStatus = "pending"
	Open     OrderStatus = "open"
	Closed   OrderStatus = "closed"
	Canceled OrderStatus = "canceled"
	Expired  OrderStatus = "expired"
)

type OrderType string

const (
	Market          OrderType = "market"
	Limit           OrderType = "limit"
	StopLoss        OrderType = "stop-loss"
	TakeProfit      OrderType = "take-profit"
	StopLossLimit   OrderType = "stop-loss-limit"
	TakeProfitLimit OrderType = "take-profit-limit"
	SettlePosition  OrderType = "settle-position"
)

type Type string

const (
	Buy  Type = "buy"
	Sell Type = "sell"
)

type TriggerType string

const (
	Last  TriggerType = "last"
	Index TriggerType = "index"
)

type OrderFlag string

const (
	// Post only order (available when ordertype = limit)
	Post OrderFlag = "post"
	// Fcib prefer fee in base currency (default if selling)
	Fcib OrderFlag = "fcib"
	// Fciq prefer fee in quote currency (default if buying, mutually exclusive with fcib)
	Fciq OrderFlag = "fciq"
	// Nompp disable market price protection for market orders
	Nompp OrderFlag = "nompp"
	// Viqc order volume expressed in quote currency. This is supported only for market orders
	Viqc OrderFlag = "viqc"
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
	// Alternate name
	Altname string `json:"altname"`
	// Asset class
	AssetClass AssetClass `json:"aclass"`
	// Scaling decimal places for record keeping
	Decimals int64 `json:"decimals"`
	// Scaling decimal places for output display
	DisplayDecimals int64 `json:"display_decimals"`
	// Valuation as margin collateral (if applicable)
	CollateralValue decimal.Decimal `json:"collateral_value"`
	// Status of asset
	Status AssetStatus `json:"status"`
}

type AssetPairsInfo struct {
	// Alternate name
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
	// Scaling decimal places for cost
	CostDecimals int64 `json:"cost_decimals"`
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
	// Minimum order cost (in terms of quote currency)
	CostMin decimal.Decimal `json:"costmin"`
	// Minimum increment between valid price levels
	TickSize decimal.Decimal `json:"tick_size"`
	// Status of asset pair
	Status AssetPairStatus `json:"status"`
	// Long position limit
	LongPositionLimit int64 `json:"long_position_limit"`
	// Short position limit
	ShortPositionLimit int64 `json:"short_position_limit"`
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
	// Trade ID
	TradeID int64 `json:"trade_id"`
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
	EquivalentBalance decimal.Decimal `json:"eb"`
	// Trade balance (combined balance of all equity currencies)
	TradeBalance decimal.Decimal `json:"tb"`
	// Margin amount of open positions
	MarginOP decimal.Decimal `json:"m"`
	// Unrealized net profit/loss of open positions
	UnrealizedNetProfitLossOP decimal.Decimal `json:"n"`
	// Cost basis of open positions
	CostBasisOP decimal.Decimal `json:"c"`
	// Current floating valuation of open positions
	CurrentValuationOP decimal.Decimal `json:"v"`
	// Equity = trade balance + unrealized net profit/loss
	Equity decimal.Decimal `json:"e"`
	// FreeMargin = equity - initial margin (maximum margin available to open new positions)
	FreeMargin decimal.Decimal `json:"mf"`
	// MargimLevel = (equity / initial margin) * 100
	MarginLevel decimal.Decimal `json:"ml"`
	// Unexecuted value: Value of unfilled and partially filled orders
	Unexecuted decimal.Decimal `json:"uv"`
}

type Order struct {
	// Referral order transaction ID that created this order
	ReferralOrderTxID string `json:"refid"`
	// User reference id
	UserReferenceID int64 `json:"userref"`
	// Status of order
	Status OrderStatus `json:"status"`
	// Unix timestamp of when order was placed
	OpenedAt time.Time `json:"opentm"`
	// Unix timestamp of order start time (or 0 if not set)
	StartAt time.Time `json:"starttm"`
	// Unix timestamp of order end time (or 0 if not set)
	ExpireAt time.Time `json:"expiretm"`

	// Order description info
	OrderDescription struct {
		// Asset pair
		Pair AssetPair `json:"pair"`
		// Type of order (buy/sell)
		Type Type `json:"type"`
		// Order type
		Ordertype OrderType `json:"ordertype"`
		// Primary price
		Price decimal.Decimal `json:"price"`
		// Secondary price
		Price2 decimal.Decimal `json:"price2"`
		// Amount of leverage
		Leverage string `json:"leverage"`
		// Order description
		Order string `json:"order"`
		// Conditional close order description (if conditional close set)
		Close string `json:"close"`
	} `json:"descr"`

	// Volume of order (base currency)
	Volume decimal.Decimal `json:"vol"`
	// Volume executed (base currency)
	VolumeExecuted decimal.Decimal `json:"vol_exec"`
	// Total cost (quote currency unless)
	Cost decimal.Decimal `json:"cost"`
	// Total fee (quote currency)
	Fee decimal.Decimal `json:"fee"`
	// Average price (quote currency)
	Price decimal.Decimal `json:"price"`
	// Stop price (quote currency)
	StopPrice decimal.Decimal `json:"stopprice"`
	// Triggered limit price (quote currency, when limit based order type triggered)
	LimitPrice decimal.Decimal `json:"limitprice"`
	// Price signal used to trigger "stop-loss" "take-profit" "stop-loss-limit" "take-profit-limit" orders.
	Trigger TriggerType `json:"trigger"`
	// Comma delimited list of miscellaneous info
	// - stopped triggered by stop price
	// - touched triggered by touch price
	// - liquidated liquidation
	// - partial partial fill
	Miscellaneous string `json:"misc"`
	// List of order flags
	Flags []OrderFlag `json:"oflags"`
	// List of trade IDs related to order (if trades info requested and data available)
	TradesIDs []string `json:"trades"`
}
