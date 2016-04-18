/*
Exercise 6.4 - Add a method Elems that returns a slice containing elements of the set

$ go run src/ch6/ex_6_4/main.go
*/

package main

import (
    "fmt"

    "ch6/intset"
)

func main() {
    var x intset.IntSet
    x.Add(1)
    x.Add(2)
    x.Add(4)
    x.AddAll(8, 16, 32, 64)
    fmt.Printf("x=%s, x.Len()=%d\n", x.String(), x.Len())

    fmt.Println("range x.Elem()")
    for _, e := range x.Elem() {
        fmt.Println(e)
    }
}
