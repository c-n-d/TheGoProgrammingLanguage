/*
Http4 is an e-commerce server that registers the /list and /price endpoints with http.HandleFunc

$ go run src/ch7/http4/main.go
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

func main() {
    db := database{"shoes": 50, "socks":5}
    http.HandleFunc("/list", db.list)
    http.HandleFunc("/price", db.price)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, r *http.Request) {
    for item, price := range db {
        fmt.Fprintf(w, "%s: %s\n", item, price)
    }
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
    item := r.URL.Query().Get("item")
    price, ok := db[item]
    if !ok {
        w.WriteHeader(http.StatusNotFound) // 404
        fmt.Fprintf(w, "no such items: %q\n", item)
        return
    }
    fmt.Fprintf(w, "%s\n", price)
}
