package main

import (
	"encoding/json"
	"io/ioutil"
)

func ParseJson(filePath string) Paths {
	var paths Paths
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(file), paths)
	return paths
}

var paths Paths

func init() {
	paths = ParseJson("../paths.json")
}

type Paths struct {
	APIURL  []string `json:"api_url"`
	Account struct {
		AccountStatus           string `json:"accountStatus"`
		AccountAPITradingStatus string `json:"accountApiTradingStatus"`
		APIKeyPermission        string `json:"apiKeyPermission"`
		AccountInfo             string `json:"accountInfo"`
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
	Spot struct {
		NewOrder    string `json:"newOrder"`
		CancelOrder string `json:"cancelOrder"`
		OpenOrders  string `json:"openOrders"`
		QueryOrder  string `json:"queryOrder"`
		AllOrders   string `json:"allOrders"`
	}
	Ticker struct {
		PriceTicker string `json:"priceTicker"`
	} `json:"ticker"`
}
