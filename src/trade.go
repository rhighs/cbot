package main

import (
    "cbot/responses"
    "io/ioutil"
)

func (c *Client) PlaceLimitOrder() responses.NewOrderFull {
    var result responses.NewOrderFull
    resp := c.do("POST", paths.Spot.NewOrder, true, &result)
    ioutil.ReadAll(resp.Body)
    return result
}

func (c *Client) fetchTicker() responses.PriceTicker {
    var result responses.PriceTicker
    c.do("GET", "/fapi/v1/ticker/price", false, &result)
    return result
}
