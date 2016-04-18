/*
Exercise 6.2 - Write a variadic (*IntSet).AddAll(...int) method

$ go run src/ch6/ex_6_2/main.go
x={1 2 4}, x.Len()=3

x.AddAll(8, 16, 32, 64)
x={1 2 4 8 16 32 64}, x.Len()=7

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
    fmt.Printf("x=%s, x.Len()=%d\n", x.String(), x.Len())

    x.AddAll(8, 16, 32, 64)
    fmt.Println("\nx.AddAll(8, 16, 32, 64)")
    fmt.Printf("x=%s, x.Len()=%d\n", x.String(), x.Len())
}
