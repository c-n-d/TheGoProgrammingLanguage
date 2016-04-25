/*
Exercise 7.3 - Write a String method for the *tree type

$ go run src/ch7/ex_7_3/main.go
t.String() =  0 3 5 6 7 8 9
*/

package main

import (
    "fmt"
)

type tree struct {
    value int
    left, right *tree
}

func (t *tree) String() string {
    return t.collectValues("")
}

func (t *tree) collectValues(values string) string {
    if t != nil {
        values = t.left.collectValues(values)
        values += fmt.Sprintf("%d ", t.value)
        values = t.right.collectValues(values)
    }
    return values
}

func (t *tree) addAll(values...int) *tree {
    for _, v := range values {
        t = t.add(v)
    }
    return t
}

func (t *tree) add(value int) *tree {
    if t == nil {
        t = new(tree)
        t.value = value
        return t
    }
    if value < t.value {
        t.left = t.left.add(value)
    } else {
        t.right = t.right.add(value)
    }
    return t
}

func main() {
    var t *tree = nil
    t = t.addAll(8, 7, 6, 5, 3, 0, 9)
    fmt.Println("t.String() = ", t.String())
}
