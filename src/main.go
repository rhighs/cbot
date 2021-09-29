package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var TEST_API_KEY = "vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A"
var TEST_API_SECRET = "NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j"

type Keys struct {
	ApiKey    string `json:"apiKey"`
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
	balance := client.accountInfo()
	fmt.Printf("%+v\n", balance)

	priceOfAllCrypto := client.fetchTicker()
	fmt.Printf("%+v\n", priceOfAllCrypto)

}
