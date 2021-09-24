package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var URL string = "https://httpbin.org/post"
var TEST_API_KEY = "vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A"
var TEST_API_SECRET = "NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j"

type Res struct {
	Mins  int64  `json:"mins"`
	Price string `json:"price"`
}

func main() {
	client, err := NewClient(TEST_API_KEY, TEST_API_SECRET)
	if err != nil {
		panic(err)
	}

	resp := client.accountInfo()
	fmt.Printf("%+v\n", resp)
	/*
	   client.do("GET", "/api/v3/avgPrice?symbol=BTCEUR", "", false, &res)
	   fmt.Printf("mins: " +  strconv.FormatInt(res.Mins, 10) + " | symbol: " + res.Price + "\n")
	*/
}

func PostTest() string {
	res, err := http.Post(URL, "application/text", nil)
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("error occurred")
	}
	return string(data)
}
