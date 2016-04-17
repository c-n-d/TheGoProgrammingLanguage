/*
Fetch stores a file on the local system after fetching the URL

$ go run src/ch5/fetch/main.go https://golang.org
fetch(https://golang.org) = index.html, 7858, <nil>
*/

package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path"
)
func fetch(url string) (filename string, n int64, err error) {
    resp, err := http.Get(url)
    if err != nil {
        return "", 0, err
    }
    defer resp.Body.Close()

    local := path.Base(resp.Request.URL.Path)
    if local == "/" || local == "." {
        local = "index.html"
    }
    f, err := os.Create(local)
    if err != nil {
        return "", 0, err
    }
    n, err = io.Copy(f, resp.Body)
    if closeErr := f.Close(); err == nil {
        err = closeErr
    }
    return local, n, err
}

func main() {
    for _, url := range os.Args[1:] {
        fileName, n, err := fetch(url);
        if err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Printf("fetch(%s) = %s, %d, %v\n",url, fileName, n, err)
    }
}
