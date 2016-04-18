/*
Exercise 6.1 - Implement Len() int, Remove(x int), Clear(), and Copy() *IntSet

$ go run src/ch6/ex_6_1/main.go
x={1 2 4}, x.Len()=3

x.Add(8)
x={1 2 4 8}, x.Len()=4

y := x.Copy()
y={1 2 4 8}, y.Len()=4

x.Remove(8)
x.Remove(16)
x={1 2 4}, x.Len()=3
y={1 2 4 8}, y.Len()=4

x.Clear()
x={}, x.Len()=0
y={1 2 4 8}, y.Len()=4
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

    x.Add(8)
    fmt.Println("\nx.Add(8)")
    fmt.Printf("x=%s, x.Len()=%d\n", x.String(), x.Len())

    y := x.Copy()
    fmt.Println("\ny := x.Copy()")
    fmt.Printf("y=%s, y.Len()=%d\n", y.String(), y.Len())

    x.Remove(8)
    x.Remove(16)
    fmt.Println("\nx.Remove(8)")
    fmt.Println("x.Remove(16)")
    fmt.Printf("x=%s, x.Len()=%d\n", x.String(), x.Len())
    fmt.Printf("y=%s, y.Len()=%d\n", y.String(), y.Len())

    x.Clear()
    fmt.Println("\nx.Clear()")
    fmt.Printf("x=%s, x.Len()=%d\n", x.String(), x.Len())
    fmt.Printf("y=%s, y.Len()=%d\n", y.String(), y.Len())
}
