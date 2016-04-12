/*
Exercise 4.3 - Reverse using array pointers

$ go run src/ch4/ex_4_3/main.go
reverse([0 1 1 2 3 5]) = [5 3 2 1 1 0]
*/

package main

import "fmt"

func main() {
    a := [6]int{0, 1, 1, 2, 3, 5}

    fmt.Printf("reverse(%v) = ", a)
    reverse(&a)
    fmt.Printf("%v\n", a)
}

func reverse(s *[6]int) {
    for i, j := 0, len(*s) - 1; i < j; i, j = i + 1, j - 1 {
        (*s)[i], (*s)[j] = (*s)[j], (*s)[i]
    }
}
