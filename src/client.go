package main

import (
    "net/http"
)

type Client struct {
    apiKey string
    apiSecret string
    httpClient *httpClient
}

func NewClient(apiKey string, apiSecret string) (c *Client) {
    return &Client {
        apiKey: apiKey,
        apiSecrect: apiSecret,
        httpClient: &http.Client{}
    }
}

/*
* TODO
* Basically we want this "binance client" to call any given path and any given method, via this go
* "member function"
* inspired by: https://github.com/pdepip/go-binance/blob/master/binance/client.go
*/
func (c *Client) do(method string, fullUrl string, payload string, isAuth bool, result interface{}) (resp *http.Response) {
    return nil
}

