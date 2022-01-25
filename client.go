//go:generate go run generate/generate.go generate/template_asset.go

package kraken

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// APIUrl is the Kraken API Endpoint
	apiURL = "https://api.kraken.com"

	// APIVersion is the Kraken API version number
	apiVersion = "0"
)

type Response struct {
	Error  []string    `json:"error"`
	Result interface{} `json:"result"`
}

type Client struct {
	httpClient *http.Client
	apiKey     string
	apiSecret  string
}

// New inits a new Client
func New() *Client {
	return &Client{
		httpClient: http.DefaultClient,
	}
}

// WithAuthentification allows to specify a custom secret in order to execute private requests
func (c *Client) WithAuthentification(key, secret string) {
	c.apiKey = key
	c.apiSecret = secret
}

func (c *Client) doRequest(endpoint string, isPrivate bool, data url.Values, respType interface{}) error {
	var (
		req *http.Request
		err error
	)

	if isPrivate {
		req, err = c.buildPrivateRequest(endpoint, data)
		if err != nil {
			return err
		}
	} else {
		req, err = c.buildPublicRequest(endpoint, data)
		if err != nil {
			return err
		}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return errors.New("failed to make http request")
	}
	defer resp.Body.Close()

	return c.parseResponse(resp, respType)
}

func (c *Client) buildPublicRequest(endpoint string, data url.Values) (*http.Request, error) {
	if data == nil {
		data = url.Values{}
	}

	URL := fmt.Sprintf("%s/%s/public/%s", apiURL, apiVersion, endpoint)
	req, err := http.NewRequest("POST", URL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create public request: %s", err.Error())
	}

	return req, nil
}

func (c *Client) buildPrivateRequest(endpoint string, data url.Values) (*http.Request, error) {
	if c.apiKey == "" || c.apiSecret == "" {
		return nil, fmt.Errorf("failed to create private request: key or secret is empty")
	}

	if data == nil {
		data = url.Values{}
	}
	data.Set("nonce", fmt.Sprintf("%d", time.Now().UTC().UnixMilli()))

	URL := fmt.Sprintf("%s/%s/private/%s", apiURL, apiVersion, endpoint)
	req, err := http.NewRequest("POST", URL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create private request: %s", err.Error())
	}

	req.Header.Add("API-Key", c.apiKey)

	signature, err := c.getSignature(fmt.Sprintf("/%s/private/%s", apiVersion, endpoint), data)
	if err != nil {
		return nil, fmt.Errorf("failed to sign private request: %s", err.Error())
	}
	req.Header.Add("API-Sign", signature)

	return req, nil
}

func (c *Client) getSignature(requestURL string, data url.Values) (string, error) {
	sha := sha256.New()
	_, err := sha.Write([]byte(data.Get("nonce") + data.Encode()))
	if err != nil {
		return "", err
	}
	shasum := sha.Sum(nil)

	secret, err := base64.StdEncoding.DecodeString(c.apiSecret)
	if err != nil {
		return "", err
	}

	mac := hmac.New(sha512.New, secret)
	_, err = mac.Write(append([]byte(requestURL), shasum...))
	if err != nil {
		return "", err
	}
	macsum := mac.Sum(nil)

	return base64.StdEncoding.EncodeToString(macsum), nil
}

func (c *Client) parseResponse(response *http.Response, respType interface{}) error {
	if response.StatusCode != 200 {
		return fmt.Errorf("failed to get a successful response. status %d", response.StatusCode)
	}

	if response.Body == nil {
		return fmt.Errorf("failed to get a response body")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %+v", err)
	}

	var resp Response
	if respType != nil {
		resp.Result = respType
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %+v", err)
	}

	if len(resp.Error) > 0 {
		return fmt.Errorf("got server errors: %+v", resp.Error)
	}

	return nil
}
