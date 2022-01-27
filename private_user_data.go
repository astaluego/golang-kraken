package kraken

import (
	"math"
	"net/url"
	"strings"
	"time"

	"github.com/shopspring/decimal"
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
	// Asset is optional (default: USD)
	// Base asset used to determine balance
	Asset Asset
}

// TradeBalance
// Retrieve a summary of collateral balances, margin position valuations, equity and margin level.
// https://docs.kraken.com/rest/#operation/getTradeBalance
func (c *Client) TradeBalance(config TradeBalanceConfig) (*TradeBalance, error) {
	payload := Payload{}
	payload.OptAssets(config.Asset)

	response := TradeBalance{}
	err := c.doRequest("TradeBalance", true, url.Values(payload), &response)
	return &response, err
}

type OpenOrdersConfig struct {
	// Trades is optional (default: false)
	WithRelatedTrades bool
	// UserReferenceID is optional (Restrict results to given user reference id)
	UserReferenceID int64
}

// OpenOrders
// Retrieve information about currently open orders.
// https://docs.kraken.com/rest/#operation/getOpenOrders
func (c *Client) OpenOrders(config OpenOrdersConfig) (*map[string]OpenOrder, error) {
	payload := Payload{}
	payload.OptWithRelatedTrades(config.WithRelatedTrades)
	payload.OptUserReferenceID(config.UserReferenceID)

	var resp interface{}
	err := c.doRequest("OpenOrders", true, url.Values(payload), &resp)

	response := make(map[string]OpenOrder)

	for key, value := range resp.(map[string]interface{}) {
		if key == "open" {
			for k, v := range value.(map[string]interface{}) {
				openOrder := OpenOrder{}
				for k, v := range v.(map[string]interface{}) {
					switch k {
					case "refid":
						if v != nil {
							openOrder.ReferralOrderxID = v.(string)
						}
					case "userref":
						openOrder.UserReferenceID = int64(v.(float64))
					case "status":
						openOrder.Status = OrderStatus(v.(string))
					case "opentm":
						integer, decimal := math.Modf(v.(float64))
						t := time.Unix(int64(integer), int64(decimal*10000))
						openOrder.OpenAt = t
					case "starttm":
						if v.(float64) != 0 {
							integer, decimal := math.Modf(v.(float64))
							t := time.Unix(int64(integer), int64(decimal*10000))
							openOrder.StartAt = t
						}
					case "expiretm":
						if v.(float64) != 0 {
							integer, decimal := math.Modf(v.(float64))
							t := time.Unix(int64(integer), int64(decimal*10000))
							openOrder.ExpireAt = t
						}
					case "descr":
						for k, v := range v.(map[string]interface{}) {
							switch k {
							case "pair":
								openOrder.OrderDescription.AssetPair = AssetPair(v.(string))
							case "type":
								openOrder.OrderDescription.Type = Type(v.(string))
							case "ordertype":
								openOrder.OrderDescription.OrderType = OrderType(v.(string))
							case "price":
								price, err := decimal.NewFromString(v.(string))
								if err != nil {
									continue
								}
								openOrder.OrderDescription.Price = price
							case "price2":
								price, err := decimal.NewFromString(v.(string))
								if err != nil {
									continue
								}
								openOrder.OrderDescription.Price2 = price
							case "leverage":
								openOrder.OrderDescription.Leverage = v.(string)
							case "order":
								openOrder.OrderDescription.Description = v.(string)
							case "close":
								openOrder.OrderDescription.CloseCondition = v.(string)
							}

						}
					case "vol":
						vol, err := decimal.NewFromString(v.(string))
						if err != nil {
							continue
						}
						openOrder.Volume = vol
					case "vol_exec":
						vol, err := decimal.NewFromString(v.(string))
						if err != nil {
							continue
						}
						openOrder.VolumeExecuted = vol
					case "cost":
						cost, err := decimal.NewFromString(v.(string))
						if err != nil {
							continue
						}
						openOrder.Cost = cost
					case "fee":
						fee, err := decimal.NewFromString(v.(string))
						if err != nil {
							continue
						}
						openOrder.Fee = fee
					case "price":
						price, err := decimal.NewFromString(v.(string))
						if err != nil {
							continue
						}
						openOrder.Price = price
					case "stopprice":
						price, err := decimal.NewFromString(v.(string))
						if err != nil {
							continue
						}
						openOrder.StopPrice = price
					case "limitprice":
						price, err := decimal.NewFromString(v.(string))
						if err != nil {
							continue
						}
						openOrder.LimitPrice = price
					case "trigger":
						openOrder.Trigger = Trigger(v.(string))
					case "misc":
						if v.(string) == "" {
							continue
						}
						list := strings.Split(v.(string), ",")
						for _, elem := range list {
							openOrder.MiscellaneousInfos = append(openOrder.MiscellaneousInfos, MiscellaneousInfo(elem))
						}
					case "oflags":
						if v.(string) == "" {
							continue
						}
						list := strings.Split(v.(string), ",")
						for _, elem := range list {
							openOrder.OrderFlags = append(openOrder.OrderFlags, OrderFlag(elem))
						}
					case "trades":
						openOrder.RelatedTradesIDs = v.([]string)
					}
				}
				response[k] = openOrder
			}
		}
	}

	return &response, err
}
