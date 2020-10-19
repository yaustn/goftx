package goftx

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

const (
	apiURL = "https://ftx.com/api"
)

// Client is a
type Client struct {
	client *http.Client
	apiKey string

	lastRequestTimestamp string
}

// NewClient returns a client to make requests to FTX's REST api
func NewClient() *Client {
	client := &http.Client{}

	return &Client{
		client: client,
	}
}

func (c *Client) get(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))

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

/*
func (c *Client) post() {

}

func (c *Client) sign() {

}
*/
