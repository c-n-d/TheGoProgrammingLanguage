/*
urlvalues demonstrates a map type with methods

$ go run src/ch6/urlvalues/main.go
en

1
[1 2]

panic: assignment to entry in nil map

goroutine 1 [running]:
panic
	src/runtime/panic.go:464
main.main()
	src/ch6/urlvalues/main.go:24
exit status 2
*/

package main

import (
    "fmt"
    "net/url"
)

func main() {
    m := url.Values{"lang": {"en"}} // direct constructor
    m.Add("item", "1")
    m.Add("item", "2")

    fmt.Println(m.Get("lang"))
    fmt.Println(m.Get("q"))
    fmt.Println(m.Get("item"))
    fmt.Println(m["item"])

    m = nil
    fmt.Println(m.Get("item"))
    m.Add("item", "3")
}
