/*
Server3 is a minimal server that "echos" the HTTP request info

$ go run src/ch1/server3/main.go

$ curl localhost:8888
GET / HTTP/1.1
Header[User-Agent] = [curl]
Header[Accept] = [*/*]
Host = [localhost:8888]
RemoteAddr = 127.0.0.1
Form[q] = [value]
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

// handler echos the HTTP request
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

    for k, v := range r.Header {
        fmt.Fprintf(w, "Header[%v] = %v\n", k, v)
    }

    fmt.Fprintf(w, "Host = [%v]\n", r.Host)
    fmt.Fprintf(w, "RemoteAddr = %v\n", r.RemoteAddr)

    if err := r.ParseForm(); err != nil {
        log.Print(err)
    }

    for k, v := range r.Form {
        fmt.Fprintf(w, "Form[%v] = %v\n", k, v)
    }
}
