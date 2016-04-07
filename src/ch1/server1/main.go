/*
Server1 is a minimal "echo" server

$ go run src/ch1/server1/main.go

$ curl localhost:8888
URL.Path = "/"
*/

package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    // Register the request handler for the root route
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

// handler echos the path component of the request URL r
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
