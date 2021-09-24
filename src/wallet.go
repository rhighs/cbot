package main

import (
	"cbot/responses" //gotta specify the mod name
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"
    "fmt"
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

func (c *Client) accountInfo() responses.AccountInfo {
	var paths Paths
	var resp responses.AccountInfo
	file, err := ioutil.ReadFile("../paths.json")

	if err != nil {
		panic(err)
	}

    fmt.Printf(paths.Account.AccountInfo)

	err = json.Unmarshal([]byte(file), &paths)
	c.do("GET", "/api/v3/account"+"?timestamp="+strconv.FormatInt(time.Now().Unix(), 10), "", true, &resp)

	return resp
}
