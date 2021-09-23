package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

var URL string = "https://httpbin.org/post"

func main() {
    binance := Binance{}
    fmt.Printf(binance.systemStatus())
}

func PostTest() string {
    res, err  := http.Post(URL, "application/text", nil)
    data, err := ioutil.ReadAll(res.Body)


    if err != nil {
        fmt.Printf("error occurred")
    }
    return string(data)
}
