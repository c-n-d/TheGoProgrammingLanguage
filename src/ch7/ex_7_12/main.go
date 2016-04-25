/*
Exercise 7.12 - Change the hadler for /list to print its output as an HTML table

$ curl http://localhost:8000/list
<table style="border: solid black;border-collapse:collapse">
    <thead>
        <tr>
            <th style="border: solid black">Item</th>
            <th style="border: solid black">Price</th>
        </tr>
    </thead>
    <tbody>
        
        <tr>
            <td style="border: thin solid black">shoes</td>
            <td style="border: thin solid black">$50.00</td>
        </tr>
        
        <tr>
            <td style="border: thin solid black">socks</td>
            <td style="border: thin solid black">$5.00</td>
        </tr>
        
    </tbody>
</table>
*/

package main

import (
    "fmt"
    "html/template"
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

const listTempl = `<table style="border: solid black;border-collapse:collapse">
    <thead>
        <tr>
            <th style="border: solid black">Item</th>
            <th style="border: solid black">Price</th>
        </tr>
    </thead>
    <tbody>
        {{range $item, $price := .}}
        <tr>
            <td style="border: thin solid black">{{$item}}</td>
            <td style="border: thin solid black">{{$price}}</td>
        </tr>
        {{end}}
    </tbody>
</table>
`

var list = template.Must(template.New("listtable").
    Parse(listTempl))

type database map[string]dollars

func (db database) list(w http.ResponseWriter, r *http.Request) {
    if err := list.Execute(w, db); err != nil {
        w.WriteHeader(http.StatusInternalServerError) // 500
        fmt.Fprintf(w, "unable to print list: %q\n", err)
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
