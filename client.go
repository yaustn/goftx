package goftx

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// DefaultMaxRequestsPerSecond defines the rate at which FTX will accept requests.
	// Throughput beyond this will result in FTX returning HTTP 429 errors.
	DefaultMaxRequestsPerSecond = 30

	apiURL = "https://ftx.us/api"
)

// Response represents every FTX API response.
// Response.Result is an interface specific to the request and should be
// type-casted before use.
type Response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

// Client for interfacing with the FTX REST api.
type Client struct {
	client       *http.Client
	ftxAPIKey    string
	ftxAPISecret string
}

func NewClient(ftxAPIKey, ftxAPISecret string) *Client {
	client := &http.Client{}

	return &Client{
		client:       client,
		ftxAPIKey:    ftxAPIKey,
		ftxAPISecret: ftxAPISecret,
	}
}

func (c *Client) get(endpoint string, result interface{}) error {
	err := c.do("GET", endpoint, []byte{}, &result)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) post(endpoint string, request []byte, result interface{}) error {
	err := c.do("POST", endpoint, request, &result)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) delete(endpoint string, request []byte, result interface{}) error {
	err := c.do("DELETE", endpoint, request, &result)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) do(method, endpoint string, request []byte, result interface{}) error {
	response := new(Response)
	response.Result = result

	req := c.buildSignedRequest(method, endpoint, request)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respBody, response)
	if err != nil {
		return err
	} else if !response.Success {
		return errors.New(response.Error)
	}

	return nil
}

// See https://docs.ftx.com/#authentication
func (c *Client) buildSignedRequest(method, endpoint string, body []byte) *http.Request {
	// Get the current epoch ms timestamp
	nonce := c.timestamp()

	// SHA256 HMAC of concatenated string request encoded with the client's api secret
	signaturePayload := nonce + method + "/api" + endpoint + string(body)
	signature := c.sign(signaturePayload)

	// Create and sign the http request
	req, _ := http.NewRequest(method, apiURL+endpoint, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("FTXUS-KEY", c.ftxAPIKey)
	req.Header.Set("FTXUS-SIGN", signature)
	req.Header.Set("FTXUS-TS", nonce)

	return req
}

// sign takes in a string and returns the SHA256 HMAC using the client's api secret
func (c *Client) sign(payload string) string {
	mac := hmac.New(sha256.New, []byte(c.ftxAPISecret))
	mac.Write([]byte(payload))
	signature := hex.EncodeToString(mac.Sum(nil))

	return signature
}

// timestamp returns a string formatted epoch millisecond timestamp
func (c *Client) timestamp() string {
	epochMilli := int64(time.Now().UTC().UnixNano()) / int64(time.Millisecond)
	return fmt.Sprintf("%d", epochMilli)
}
