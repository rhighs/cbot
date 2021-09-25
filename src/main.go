package main

import (
	"fmt"
	"io/ioutil"
    "encoding/json"
	"net/http"
)

var URL string = "https://httpbin.org/post"
var TEST_API_KEY = "8jFR3cMWfCWTLcBf01eyD523F38qUaTunv12dax8YJS6cOhVQerN5sJXgZOSNxQL"
var TEST_API_SECRET = "sdYR2qyhhcNhMXy63K0ZohjwIQL7bCzkzsWT9PglcDaChHiDPzLY1hQRpfBgAPwT"

type Keys struct {
    ApiKey string `json:"apiKey"`
    ApiSecret string `json:"apiSecret"`
}

type Res struct {
	Mins  int64  `json:"mins"`
	Price string `json:"price"`
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

func PostTest() string {
	res, err := http.Post(URL, "application/text", nil)
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("error occurred")
	}
	return string(data)
}
