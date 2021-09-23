package main

import (
    "encoding/json"
    "io/ioutil"
    "cbot/responses" //gotta specify the mod name
)

func (c *Client) systemStatus() string {
    var paths Paths
    var resp responses.AccountStatus
    file, err := ioutil.ReadFile("../paths.json")
    if err != nil {
        panic(err)
    }
    err = json.Unmarshal([]byte(file), paths)
    res := c.do("GET", paths.Account.AccountStatus, "", true, &resp)
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()
    return string(body)
}

type Paths struct {
	APIURL  []string `json:"api_url"`
	Account struct {
		AccountStatus           string `json:"accountStatus"`
		AccountAPITradingStatus string `json:"accountApiTradingStatus"`
		APIKeyPermission        string `json:"apiKeyPermission"`
	} `json:"account"`
	Asset struct {
		DustTransfer          string `json:"dustTransfer"`
		AssetDividendRecord   string `json:"assetDividendRecord"`
		AssetDetail           string `json:"assetDetail"`
		TradeFee              string `json:"tradeFee"`
		UserUniversalTransfer string `json:"userUniversalTransfer"`
		TransferHistoryQuery  string `json:"transferHistoryQuery"`
		FundingWallet         string `json:"fundingWallet"`
	} `json:"asset"`
	Wallet struct {
		Withdraw        string `json:"withdraw"`
		Deposit         string `json:"deposit"`
		WithdrawHistory string `json:"withdrawHistory"`
		DepositAddress  string `json:"depositAddress"`
	} `json:"wallet"`
}
