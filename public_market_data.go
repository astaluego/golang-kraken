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
