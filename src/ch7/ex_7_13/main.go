/*
Exercise 7.13 - Add Strings method to Expr interface to pretty-print the syntax tree

$ go run src/ch7/surface/main.go
$ open "http://localhost:8000/echo?expr=sin(-x)*pow(1.5,-r)"
$ open "http://localhost:8000/echo?expr=pow(2, sin(y))*pow(2,sin(x))/12"
$ open "http://localhost:8000/echo?expr=sin(x*y/10)/10"

> Your expression is: sin(-x)*pow(1.500000, -r)
> Your expression is: pow(2.000000, sin(y))*pow(2.000000, sin(x))/12.000000
> Your expression is: sin(x*y/10.000000)/10.000000
*/

package main

import (
    "fmt"
    "log"
    "net/http"

    "ch7/eval"
)

func parseAndCheck(s string) (eval.Expr, error) {
    if s == "" {
        return nil, fmt.Errorf("empty expression")
    }
    expr, err := eval.Parse(s)
    if err != nil {
        return nil, err
    }
    vars := make(map[eval.Var]bool)
    if err := expr.Check(vars); err != nil {
        return nil, err
    }
    for v := range vars {
        if v != "x" && v != "y" && v != "r" {
            return nil, fmt.Errorf("undefined variable: %s", v)
        }
    }
    return expr, nil
}

func echo(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    expr, err := parseAndCheck(r.Form.Get("expr"))
    if err != nil {
        http.Error(w, "bad expr: " + err.Error(), http.StatusBadRequest)
        return
    }
    w.Header().Set("Content-Type", "text/html")

    fmt.Fprintf(w, "Your expression is: %s", expr.String())
}

func main() {
    http.HandleFunc("/echo", echo)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
