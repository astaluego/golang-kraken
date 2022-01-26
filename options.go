package kraken

import (
	"net/url"
	"strconv"
	"strings"
	"time"
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

func (payload Payload) OptAssetPairs(assetPairs ...AssetPair) {
	if len(assetPairs) == 0 {
		return
	}

	list := []string{}
	for _, assetPair := range assetPairs {
		list = append(list, string(assetPair))
	}
	payload["pair"] = []string{strings.Join(list, ",")}
}

func (payload Payload) OptCount(count int64) {
	payload["count"] = []string{strconv.FormatInt(count, 10)}
}

func (payload Payload) OptInformations(information Information) {
	if string(information) == "" {
		return
	}

	payload["info"] = []string{string(information)}
}

func (payload Payload) OptSince(time time.Time) {
	if time.IsZero() {
		return
	}

	payload["since"] = []string{strconv.FormatInt(time.Unix(), 10)}
}
