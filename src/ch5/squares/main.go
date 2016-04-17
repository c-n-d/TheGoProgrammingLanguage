/*
The squares program demontrates a function value with state

$ go run src/ch5/squares/main.go
1
4
9
16
25
36
*/

package main

import "fmt"

// squares returns a function that returns
// the next square number each time its called
func squares() func() int {
    var x int
    return func() int {
        x++
        return x * x
    }
}

func main() {
    f := squares()
    fmt.Println(f())
    fmt.Println(f())
    fmt.Println(f())
    fmt.Println(f())
    fmt.Println(f())
    fmt.Println(f())
}
