package goftx

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	apiURL = "https://ftx.us/api"
)

// Client for interfacing with the FTX REST api.
type Client struct {
	client    *http.Client
	apiKey    string
	apiSecret string
}

// NewClient returns a client given an api key and secret.
func NewClient(key, secret string) *Client {
	client := &http.Client{}

	return &Client{
		client:    client,
		apiKey:    key,
		apiSecret: secret,
	}
}

// func NewClient(filepath string) *Client {}

func (c *Client) get(endpoint string) ([]byte, error) {
	req := c.buildSignedRequest("GET", endpoint, []byte{})

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

// See https://docs.ftx.com/#authentication
func (c *Client) buildSignedRequest(method, endpoint string, body []byte) *http.Request {
	// Get the current epoch ms timestamp
	ts := strconv.FormatInt(time.Now().UTC().Unix()*1000, 10)

	// SHA256 HMAC of concatenated string request encoded with the client's api secret
	signaturePayload := ts + method + endpoint + string(body)
	mac := hmac.New(sha256.New, []byte(c.apiSecret))
	mac.Write([]byte(signaturePayload))
	signature := hex.EncodeToString(mac.Sum(nil))

	// Create and sign the http request
	req, _ := http.NewRequest(method, apiURL+endpoint, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("FTX-KEY", c.apiKey)
	req.Header.Set("FTX-SIGN", signature)
	req.Header.Set("FTX-TS", ts)

	return req
}

/*
func (c *Client) post() {

}

func (c *Client) sign() {

}
*/
