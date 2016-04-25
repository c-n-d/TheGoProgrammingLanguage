/*
Http2 is an e-commerce server with /list and /price endpoints

$ go run src/ch7/http2/main.go
$ curl localhost:8000
no such page: /
$ curl localhost:8000/list
shoes: $50.00
socks: $5.00
$ curl localhost:8000/price?item=shoes
$50.00
$ curl localhost:8000/price?item=hat
no such items: "hat"
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
    switch r.URL.Path {
        case "/list":
            for item, price := range db {
                fmt.Fprintf(w, "%s: %s\n", item, price)
            }
        case "/price":
            item := r.URL.Query().Get("item")
            price, ok := db[item]
            if !ok {
                w.WriteHeader(http.StatusNotFound) // 404
                fmt.Fprintf(w, "no such items: %q\n", item)
                return
            }
            fmt.Fprintf(w, "%s\n", price)
        default:
            w.WriteHeader(http.StatusNotFound) // 404
            fmt.Fprintf(w, "no such page: %s\n", r.URL)
    }
}

func main() {
    db := database{"shoes": 50, "socks":5}
    log.Fatal(http.ListenAndServe("localhost:8000", db))
}
