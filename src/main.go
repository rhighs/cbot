package main

import (
	"fmt"
	"io/ioutil"
    "encoding/json"
)

var TEST_API_KEY = "8jFR3cMWfCWTLcBf01eyD523F38qUaTunv12dax8YJS6cOhVQerN5sJXgZOSNxQL"
var TEST_API_SECRET = "sdYR2qyhhcNhMXy63K0ZohjwIQL7bCzkzsWT9PglcDaChHiDPzLY1hQRpfBgAPwT"

type Keys struct {
    ApiKey string `json:"apiKey"`
    ApiSecret string `json:"apiSecret"`
}

func ParseKeysFile() Keys {
    var keys Keys
    file, err := ioutil.ReadFile("../keys.json")
    if err != nil {
        keys.ApiKey = TEST_API_KEY
        keys.ApiSecret = TEST_API_SECRET
    } else {
        err = json.Unmarshal([]byte(file), &keys)
    }
    return keys
}

func main() {
    keys := ParseKeysFile()
	client, err := NewClient(keys.ApiKey, keys.ApiSecret)
	if err != nil {
		panic(err)
	}

	resp := client.accountInfo()
	fmt.Printf("%+v\n", resp)
}
