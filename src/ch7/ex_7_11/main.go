/*
Exercise 7.11 - Add additional handlers for create/read/update/delete

$ curl http://localhost:8000/list
shoes: $50.00
socks: $5.00
$ curl "http://localhost:8000/create?item=hat&price=3"
$ curl "http://localhost:8000/read?item=hat"
$3.00
$ curl http://localhost:8000/list
shoes: $50.00
socks: $5.00
hat: $3.00
$ curl "http://localhost:8000/update?item=socks&price=10"
$ curl http://localhost:8000/list
shoes: $50.00
socks: $10.00
hat: $3.00
$ curl "http://localhost:8000/delete?item=shoes"
$ curl http://localhost:8000/list
socks: $10.00
hat: $3.00
*/

package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func main() {
    db := database{ "shoes": 50, "socks":5 }

    var routeTable = map[string]func (w http.ResponseWriter, r *http.Request){
        "/":       db.list,
        "/list":   db.list,
        "/price":  db.price,

        "/create": db.create,
        "/read":   db.price,
        "/update": db.update,
        "/delete": db.deleteItem,
    }

    for route, handler := range routeTable {
        http.HandleFunc(route, handler)
    }

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

func (db database) create(w http.ResponseWriter, r *http.Request) {
    item := r.URL.Query().Get("item")
    price := r.URL.Query().Get("price")
    _, ok := db[item]
    if ok {
        w.WriteHeader(http.StatusBadRequest) // 400 bad request
        fmt.Fprintf(w, "entry alread exists for item, please use /update: %q\n", item)
        return
    }
    dollarPrice, err := strconv.ParseFloat(price, 32)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest) // 400 bad request
        fmt.Fprintf(w, "invalid price for item: %q\n", price)
        return
    }
    db[item] = dollars(dollarPrice)
    w.WriteHeader(http.StatusCreated) // 201 created
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
    item := r.URL.Query().Get("item")
    price := r.URL.Query().Get("price")
    _, ok := db[item]
    if !ok {
        w.WriteHeader(http.StatusNotFound) // 404
        fmt.Fprintf(w, "item does not exists, please use /create: %q\n", item)
        return
    }
    dollarPrice, err := strconv.ParseFloat(price, 32)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest) // 400 bad request
        fmt.Fprintf(w, "invalid price for item: %q\n", price)
        return
    }
    db[item] = dollars(dollarPrice)
}

func (db database) deleteItem(w http.ResponseWriter, r *http.Request) {
    item := r.URL.Query().Get("item")
    _, ok := db[item]
    if !ok {
        w.WriteHeader(http.StatusBadRequest) // 400 bad request
        fmt.Fprintf(w, "item does not exists: %q\n", item)
        return
    }
    delete(db, item)
}
