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
