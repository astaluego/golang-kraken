package kraken

import (
	"net/url"

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
	AssetClass AssetClass
	// Assets is optional
	Assets []Asset
}

// Assets
// Get information about the assets that are available for deposit, withdrawal, trading and staking.
// https://docs.kraken.com/rest/#operation/getAssetInfo
func (c *Client) Assets(config AssetsConfig) (*map[Asset]AssetInfo, error) {
	payload := Payload{}
	payload.OptAssetClass(config.AssetClass)
	payload.OptAssets(config.Assets...)

	response := make(map[Asset]AssetInfo)
	err := c.doRequest("Assets", false, url.Values(payload), &response)
	return &response, err
}

type AssetPairsConfig struct {
	// AssetPairs is optional
	AssetPairs []AssetPair
	// Information is optional
	Information Information
}

// AssetPairs
// Get tradable asset pairs
// https://docs.kraken.com/rest/#operation/getTradableAssetPairs
func (c *Client) AssetPairs(config AssetPairsConfig) (*map[AssetPair]AssetPairsInfo, error) {
	payload := Payload{}
	payload.OptAssetPairs(config.AssetPairs...)
	payload.OptInformations(config.Information)

	var resp interface{}
	err := c.doRequest("AssetPairs", false, url.Values(payload), &resp)
	if err != nil {
		return nil, err
	}

	response := make(map[AssetPair]AssetPairsInfo)

	for key, value := range resp.(map[string]interface{}) {
		assetPairsInfo := AssetPairsInfo{}

		for k, v := range value.(map[string]interface{}) {
			switch k {
			case "altname":
				assetPairsInfo.Altname = v.(string)
			case "wsname":
				assetPairsInfo.WSname = v.(string)
			case "aclass_base":
				assetPairsInfo.BaseAssetClass = AssetClass(v.(string))
			case "base":
				assetPairsInfo.BaseAsset = Asset(v.(string))
			case "aclass_quote":
				assetPairsInfo.QuoteAssetClass = AssetClass(v.(string))
			case "quote":
				assetPairsInfo.QuoteAsset = Asset(v.(string))
			case "pair_decimals":
				assetPairsInfo.PairDecimals = int64(v.(float64))
			case "lot_decimals":
				assetPairsInfo.LotDecimals = int64(v.(float64))
			case "lot_multiplier":
				assetPairsInfo.LotMultiplier = int64(v.(float64))
			case "leverage_buy":
				for _, elem := range v.([]interface{}) {
					assetPairsInfo.LeverageBuy = append(assetPairsInfo.LeverageBuy, int64(elem.(float64)))
				}
			case "leverage_sell":
				for _, elem := range v.([]interface{}) {
					assetPairsInfo.LeverageSell = append(assetPairsInfo.LeverageSell, int64(elem.(float64)))
				}
			case "fees":
				assetPairsInfo.Fees = []Fee{}
				for _, elem := range v.([]interface{}) {
					array := elem.([]interface{})
					assetPairsInfo.Fees = append(assetPairsInfo.Fees, Fee{
						Volume:  int64(array[0].(float64)),
						Percent: decimal.NewFromFloat(array[1].(float64)),
					})
				}
			case "fees_maker":
				assetPairsInfo.FeesMaker = []Fee{}
				for _, elem := range v.([]interface{}) {
					array := elem.([]interface{})
					assetPairsInfo.FeesMaker = append(assetPairsInfo.FeesMaker, Fee{
						Volume:  int64(array[0].(float64)),
						Percent: decimal.NewFromFloat(array[1].(float64)),
					})
				}
			case "fee_volume_currency":
				assetPairsInfo.FeeVolumeCurrency = Asset(v.(string))
			case "margin_call":
				assetPairsInfo.MarginCall = int64(v.(float64))
			case "margin_stop":
				assetPairsInfo.MarginStop = int64(v.(float64))
			case "ordermin":
				assetPairsInfo.OrderMin, _ = decimal.NewFromString(v.(string))
			}
		}
		response[AssetPair(key)] = assetPairsInfo
	}

	return &response, nil
}
