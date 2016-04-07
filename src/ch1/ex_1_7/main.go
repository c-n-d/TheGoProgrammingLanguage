/*
Extercise 1.7 prints the content found at the URL. Uses io.Copy(dst, src)

$ go run src/ch1/ex_1_7/main.go https://golang.org
...
Read 7856 bytes from https://golang.org

$ go run src/ch1/ex_1_7/main.go https://golang
fetch: Get https://golang: dial tcp: lookup golang: no such host
*/
package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

func main() {
    for _, url := range os.Args[1:] {
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
