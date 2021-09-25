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
