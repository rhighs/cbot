package main

import (
	"cbot/responses" //gotta specify the mod name
	"fmt"
)

func (c *Client) systemStatus() responses.AccountStatus {
	var resp responses.AccountStatus
	c.do("GET", paths.Account.AccountStatus, true, &resp)
	return resp
}

func (c *Client) accountInfo() responses.AccountInfo {
	var resp responses.AccountInfo
	fmt.Printf(paths.Account.AccountInfo)
	c.do("GET", "/api/v3/account", true, &resp)
	return resp
}
