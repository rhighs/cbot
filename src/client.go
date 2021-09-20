package main

import (
    "net/http"
)

type Client struct {
    apiKey string
    apiSecret string
    httpClient *httpClient
}

func NewClient(apiKey string, apiSecret string) *Client {
    return &Client {
        apiKey: apiKey,
        apiSecrect: apiSecret,
        httpClient: &http.Client{}
    }
}

