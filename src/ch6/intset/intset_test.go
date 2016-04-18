/*
Demonstrating IntSet

$ cd src/ch6/intset
$ go test 
{1 9 144}
{9 42}
{1 9 42 144}
true true false
{1 9 144}
{1 9 144}
{[514 0 65536]}
PASS
ok  	ch6/intset	0.008s
*/

package intset

import (
    "fmt"
    "testing"
)

func Test_one(t *testing.T) {
    var x, y IntSet
    x.Add(1)
    x.Add(144)
    x.Add(9)
    fmt.Println(x.String())

    y.Add(9)
    y.Add(42)
    fmt.Println(y.String())

    x.UnionWith(&y)

    fmt.Println(x.String())
    fmt.Println(x.Has(9), x.Has(42), x.Has(123))
}

func Test_two(t *testing.T) {
    var x IntSet
    x.Add(1)
    x.Add(144)
    x.Add(9)
    fmt.Println(&x)
    fmt.Println(x.String())
    fmt.Println(x)
}
