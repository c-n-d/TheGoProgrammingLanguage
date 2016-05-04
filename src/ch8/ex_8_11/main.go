/*

$ go run src/ch8/ex_8_11/main.go \
https://www.amazon.com http://finance.yahoo.com http://stackoverflow.com https://netflix.com \
https://www.yahoo.com https://www.youtube.com https://en.wikipedia.org https://www.google.com \
https://maps.google.com https://news.google.com https://news.ycombinator.com https://twitter.com \
https://golang.org http://godoc.org https://www.google.com/finance http://www.hulu.com/ \
https://www.bbc.com http://www.reuters.com https://www.nytimes.com https://www.apple.com \
http://www.washingtonpost.com https://www.wolframalpha.com \
https://www.amazon.com http://finance.yahoo.com http://stackoverflow.com https://netflix.com \
https://www.yahoo.com https://www.youtube.com https://en.wikipedia.org https://www.google.com \
https://maps.google.com https://news.google.com https://news.ycombinator.com https://twitter.com \
https://golang.org http://godoc.org https://www.google.com/finance http://www.hulu.com/ \
https://www.bbc.com http://www.reuters.com https://www.nytimes.com https://www.apple.com \
http://www.washingtonpost.com https://www.wolframalpha.com
*/

package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

type FetchResult struct {
    Url  string
    Body []byte
    Err  error
}

func fetch(url string, cancel <-chan struct{}) FetchResult {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println(err)
        return FetchResult{Err: err}
    }
    req.Cancel = cancel

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println(err)
        return FetchResult{Err: err}
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return FetchResult{Err: err}
    }
    return FetchResult{url, body, err}
}

func getFirstResponse() (res FetchResult) {
    firstResponse := make(chan FetchResult, 1)
    cancel := make(chan struct{})

    for _, url := range os.Args[1:] {
        go func(url string) {
            res := fetch(url, cancel)
            if res.Err == nil {
                firstResponse <- res
            }
        }(url)
    }

    res = <-firstResponse
    close(cancel)
    return
}

func main() {
    first := getFirstResponse()
    fmt.Printf("%v %v\n", first.Url, first.Err)
}
