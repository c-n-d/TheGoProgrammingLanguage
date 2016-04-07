/*
Fetchall fetches URLs in parallel and reports their times and sizes

$ go run src/ch1/fetchall/main.go https://golang.org https://www.google.com http://stackoverflow.com
0.13s  247051 http://stackoverflow.com
0.13s    7856 https://golang.org
0.14s   19113 https://www.google.com
0.14s elapsed
*/

package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {
    start := time.Now()
    // Each go routine communicates with the main routine via a channel
    ch := make(chan string)

    for _, url := range os.Args[1:] {
        // Start a go routine to execute fetch in paralle
        go fetch(url, ch)
    }

    for range os.Args[1:] {
        // Receive from the channge ch
        fmt.Println(<- ch)
    }

    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)

    if err != nil {
        // Send the error to the channel
        ch <- fmt.Sprint(err)
        return
    }

    nbytes, err := io.Copy(ioutil.Discard, resp.Body)

    // Don't leak resources
    resp.Body.Close()

    if err != nil {
        ch <- fmt.Sprintf("error while reading %s: %v", url, err);
        return
    }

    secs := time.Since(start).Seconds()

    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
