package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Client struct {
	apiKey     string
	apiSecret  string
	httpClient *http.Client
}

func NewClient(apiKey string, apiSecret string) (c *Client, err error) {
	if len(apiKey) == 0 || len(apiSecret) == 0 {
		err := errors.New("One of your keys is not valid for use, len = 0")
		return nil, err
	}
	return &Client{
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		httpClient: &http.Client{},
	}, nil
}

/*
* TODO
* Basically we want this "binance client" to call any given path and any given method, via this go
* "member function"
* inspired by: https://github.com/pdepip/go-binance/blob/master/binance/client.go
 */
func (c *Client) do(method string, path string, isAuth bool, result interface{}) (res *http.Response) {

	fullUrl := "https://api1.binance.com" + path //path should include the query string
	req, err := http.NewRequest(method, fullUrl, nil)
	if err != nil {
		return
	}
	req.Header.Add("Accept", "application/json")

	if isAuth {
		req.Header.Add("X-MBX-APIKEY", c.apiKey)
		queryString := req.URL.Query()

		timestamp := time.Now().Unix() * 1000
		queryString.Set("timestamp", strconv.FormatInt(timestamp, 10))

		mac := hmac.New(sha256.New, []byte(c.apiSecret)) //secret is the key of the cryptographic message
		_, err = mac.Write([]byte(queryString.Encode())) //write totalParams
		if err != nil {
			return
		}

		signature := hex.EncodeToString(mac.Sum(nil))
		req.URL.RawQuery = queryString.Encode() + "&signature=" + signature
	}

	res, err = c.httpClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		str, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf(string(str))
		_, err = fmt.Printf("Status code " + strconv.FormatInt(int64(res.StatusCode), 10) + " on client request...\n")
	}

	if res != nil {
		decoder := json.NewDecoder(res.Body)
		str, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf(string(str))
		err = decoder.Decode(result)
	}
	return
}
