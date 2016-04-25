/*
Http1 is a rudimetary e-commerce server

$ go run src/ch7/http1/main.go
$ curl localhost:8000
shoes: $50.00
socks: $5.00
*/

package main

import (
    "fmt"
    "log"
    "net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    for item, price := range db {
        fmt.Fprintf(w, "%s: %s\n", item, price)
    }
}

func main() {
    db := database{"shoes": 50, "socks":5}
    log.Fatal(http.ListenAndServe("localhost:8000", db))
}
