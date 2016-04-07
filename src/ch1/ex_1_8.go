/*
Extercise 1.8 builds on exercise 1.7. http:// is prepended to the URL
if it is not provided

$ go run ex_1_8.go golang.org
...
Read 7856 bytes from http://golang.org

$ go run ex_1_8.go golang
fetch: Get http://golang: dial tcp: lookup golang: no such host
*/

package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "strings"
)

func main() {
    for _, url := range os.Args[1:] {

        url = appendProtocol(url)

        resp, err := http.Get(url)

        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
            continue
        }

        // Using io.Copy, copy data from the response body directly to stdout
        written, err := io.Copy(os.Stdout, resp.Body)

        resp.Body.Close()

        if err != nil {
            fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
            os.Exit(1)
        }

        fmt.Printf("\nRead %d bytes from %s\n", written, url)
    }
}

func appendProtocol(url string) string {
    if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
        return "http://" + url
    }

    return url
}
