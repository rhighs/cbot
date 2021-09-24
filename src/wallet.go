package main

import (
	"cbot/responses" //gotta specify the mod name
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"
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
	err = json.Unmarshal([]byte(file), paths)
	res := c.do("GET", paths.Account.AccountInfo+"?timestamp="+strconv.FormatInt(time.Now().Unix(), 10), "", true, &resp)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	if body != nil {

	}
	defer res.Body.Close()
	return resp
}
