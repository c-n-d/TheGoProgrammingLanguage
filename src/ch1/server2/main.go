/*
Server2 is a minimal "echo" server and counter server

$ go run src/ch1/server2/main.go

$ curl localhost:8888
URL.Path = "/"
$ curl localhost:8888
URL.Path = "/"
$ curl localhost:8888
URL.Path = "/"

$ curl localhost:8888/count
Count 3
*/

package main

import (
    "fmt"
    "log"
    "net/http"
    "sync"
)

var mu sync.Mutex
var count int

func main() {
    // Register the request handler for the root route
    http.HandleFunc("/", handler)
    http.HandleFunc("/count", counter)
    log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

// handler echos the path component of the request URL r
func handler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    count++
    mu.Unlock()

    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echos the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    fmt.Fprintf(w, "Count %d\n", count)
    mu.Unlock()
}
