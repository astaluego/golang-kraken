package kraken

import (
	"net/url"
	"strings"
)

type Payload url.Values

func (payload Payload) OptAssets(assets ...Asset) {
	if len(assets) == 0 {
		return
	}

	list := []string{}
	for _, asset := range assets {
		list = append(list, string(asset))
	}
	payload["asset"] = []string{strings.Join(list, ",")}
}

func (payload Payload) OptAssetClass(assetClass AssetClass) {
	if string(assetClass) == "" {
		return
	}

	payload["aclass"] = []string{string(assetClass)}
}
