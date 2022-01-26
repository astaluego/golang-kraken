//go:build tools
// +build tools

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	urlAsset           = "https://api.kraken.com/0/public/Assets"
	urlAssetPairs      = "https://api.kraken.com/0/public/AssetPairs"
	pathAssetFile      = "/"
	pathAssetPairsFile = "/"
	filenameAsset      = "asset.go"
	filenameAssetPairs = "asset_pairs.go"
)

type ResponseAsset struct {
	Error  []string             `json:"error"`
	Result map[string]AssetInfo `json:"result"`
}

type AssetInfo struct {
	Altname string `json:"altname"`
}

type ResponseAssetPairs struct {
	Error  []string                  `json:"error"`
	Result map[string]AssetPairsInfo `json:"result"`
}

type AssetPairsInfo struct {
	Altname string `json:"altname"`
	WSname  string `json:"wsname"`
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Get data
	assets, err := getAssets()
	if err != nil {
		log.Fatal().Err(err)
	}
	assetPairs, err := getAssetPairs()
	if err != nil {
		log.Fatal().Err(err)
	}

	// Get the current path
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err)
	}

	// Generate asset.go file
	if err := generateAssetfile(currentPath, assets); err != nil {
		log.Error().Err(err).Msg("Failed to generate Asset file:")
	}

	// Generate asset_pairs.go file
	if err := generateAssetPairsfile(currentPath, assetPairs); err != nil {
		log.Error().Err(err).Msg("Failed to generate Asset Pairs file:")
	}
}

func getAssets() (*ResponseAsset, error) {
	// Get asset list from Kraken
	rsp, err := http.Get(urlAsset)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	// Read the json response
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	var responseAsset ResponseAsset

	// Unmarshal the response into a go structure
	err = json.Unmarshal(body, &responseAsset)
	if err != nil {
		return nil, err
	}

	// Clean data
	for key, value := range responseAsset.Result {
		value.Altname, err = rewriteAsset(value.Altname)
		if err != nil {
			log.Info().Msgf("Asset: key removed %s (%s: %s)", key, value.Altname, err.Error())
			delete(responseAsset.Result, key)
			continue
		}
		responseAsset.Result[key] = value
	}
	return &responseAsset, nil
}

func getAssetPairs() (*ResponseAssetPairs, error) {
	// Get asset pairs list from Kraken
	rsp, err := http.Get(urlAssetPairs)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	// Read the json response
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	var responseAssetPairs ResponseAssetPairs

	// Unmarshal the response into a go structure
	err = json.Unmarshal(body, &responseAssetPairs)
	if err != nil {
		return nil, err
	}

	// Clean data
	for key, value := range responseAssetPairs.Result {
		if value.WSname == "" {
			log.Info().Msgf("AssetPair: key removed %s (%s: %s)", key, value.Altname, err.Error())
			delete(responseAssetPairs.Result, key)
			continue
		}

		value.WSname = strings.ReplaceAll(value.WSname, "/", "_")
		value.WSname, err = rewriteAsset(value.WSname)
		if err != nil {
			log.Info().Msgf("AssetPair: key removed %s (%s: %s)", key, value.WSname, err.Error())
			delete(responseAssetPairs.Result, key)
			continue
		}
		responseAssetPairs.Result[key] = value
	}

	return &responseAssetPairs, nil
}

func generateAssetfile(currentPath string, assets *ResponseAsset) error {
	// Create a temporary file
	f, err := os.Create(filenameAsset + ".tmp")
	if err != nil {
		return err
	}
	defer f.Close()

	var buf bytes.Buffer

	// Generate the file with the template below the main function
	err = packageTemplateAsset.Execute(&buf, struct {
		URL    string
		Assets map[string]AssetInfo
	}{
		URL:    urlAsset,
		Assets: assets.Result,
	})
	if err != nil {
		return err
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	f.Write(p)

	// Open the file generated
	tmp, err := ioutil.ReadFile(filenameAsset + ".tmp")
	if err != nil {
		return err
	}

	// Check if there is a file
	if _, err := os.Stat(currentPath + pathAssetFile + filenameAsset); os.IsNotExist(err) {
		err := os.Rename(filenameAsset+".tmp", currentPath+pathAssetFile+filenameAsset)
		if err != nil {
			return err
		}
		fmt.Print("Asset file has been generated successfully <3\n")
		return nil
	}

	// Open the current file
	current, err := ioutil.ReadFile(currentPath + pathAssetFile + filenameAsset)
	if err != nil {
		return err
	}

	// Make a diff between the 2 files, and exit if they are differents
	if !bytes.Equal(tmp, current) {
		// If the 2 files are different
		fmt.Print("/!\\ The Asset structure has been updated\n")
		err := os.Rename(filenameAsset+".tmp", currentPath+pathAssetFile+filenameAsset)
		if err != nil {
			return err
		}
	} else {
		// If the 2 files contains the same content, remove the temporary file
		fmt.Print("\\o/ Same Asset list\n")
		err = os.Remove(filenameAsset + ".tmp")
		if err != nil {
			return err
		}
	}

	return nil
}

func generateAssetPairsfile(currentPath string, assetPairs *ResponseAssetPairs) error {
	// Create a temporary file
	f, err := os.Create(filenameAssetPairs + ".tmp")
	if err != nil {
		return err
	}
	defer f.Close()

	var buf bytes.Buffer

	// Generate the file with the template below the main function
	err = packageTemplateAssetPairs.Execute(&buf, struct {
		URL        string
		AssetPairs map[string]AssetPairsInfo
	}{
		URL:        urlAssetPairs,
		AssetPairs: assetPairs.Result,
	})
	if err != nil {
		return err
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	f.Write(p)

	// Open the file generated
	tmp, err := ioutil.ReadFile(filenameAssetPairs + ".tmp")
	if err != nil {
		return err
	}

	// Check if there is a file
	if _, err := os.Stat(currentPath + pathAssetPairsFile + filenameAssetPairs); os.IsNotExist(err) {
		err := os.Rename(filenameAssetPairs+".tmp", currentPath+pathAssetPairsFile+filenameAssetPairs)
		if err != nil {
			return err
		}
		fmt.Print("Asset Pairs file has been generated successfully <3\n")
		return nil
	}

	// Open the current file
	current, err := ioutil.ReadFile(currentPath + pathAssetPairsFile + filenameAssetPairs)
	if err != nil {
		return err
	}

	// Make a diff between the 2 files, and exit if they are differents
	if !bytes.Equal(tmp, current) {
		// If the 2 files are different
		fmt.Print("/!\\ The Asset Pairs structure has been updated\n")
		err := os.Rename(filenameAssetPairs+".tmp", currentPath+pathAssetPairsFile+filenameAssetPairs)
		if err != nil {
			return err
		}
	} else {
		// If the 2 files contains the same content, remove the temporary file
		fmt.Print("\\o/ Same Asset Pairs list\n")
		err = os.Remove(filenameAssetPairs + ".tmp")
		if err != nil {
			return err
		}
	}

	return nil
}

func rewriteAsset(asset string) (string, error) {
	switch {
	case hasDigitPrefix(asset):
		return asset, errors.New("begins with a ciffer")
	case strings.Contains(asset, ".HOLD"):
		return strings.ReplaceAll(asset, ".HOLD", "h"), nil
	case strings.Contains(asset, ".S"):
		return strings.ReplaceAll(asset, ".S", "s"), nil
	case strings.Contains(asset, ".M"):
		return strings.ReplaceAll(asset, ".M", "m"), nil
	case strings.Contains(asset, ".P"):
		return strings.ReplaceAll(asset, ".P", "p"), nil
	}

	return asset, nil
}

func hasDigitPrefix(str string) bool {
	for _, c := range str {
		if c > '0' && c < '9' {
			return true
		}
		break
	}

	return false
}
