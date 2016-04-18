/*
Exercise 6.3 - Implement IntersectWith, DifferenceWith, SymmetricDifferenceWith

$ go run src/ch6/ex_6_3/main.go
x={1 2 3}, x.Len()=3
y={3 4}, y.Len()=2
{1 2 3}.IntersectWith({3 4}) = {3}
{1 2 3}.DifferenceWith({3 4}) = {1 2}
{1 2 3}.SymmetricDifferenceWith({3 4}) = {1 2 4}
*/

package main

import (
    "fmt"

    "ch6/intset"
)

func main() {
    var x, y intset.IntSet
    x.Add(1)
    x.Add(2)
    x.Add(3)
    y.Add(3)
    y.Add(4)
    fmt.Printf("x=%s, x.Len()=%d\n", x.String(), x.Len())
    fmt.Printf("y=%s, y.Len()=%d\n", y.String(), y.Len())

    tmp := x.Copy()
    tmp.IntersectWith(&y)
    fmt.Printf("%s.IntersectWith(%s) = %s\n", x.String(), y.String(), tmp.String())

    tmp = x.Copy()
    tmp.DifferenceWith(&y)
    fmt.Printf("%s.DifferenceWith(%s) = %s\n", x.String(), y.String(), tmp.String())

    tmp = x.Copy()
    tmp.SymmetricDifferenceWith(&y)
    fmt.Printf("%s.SymmetricDifferenceWith(%s) = %s\n", x.String(), y.String(), tmp.String())
}
