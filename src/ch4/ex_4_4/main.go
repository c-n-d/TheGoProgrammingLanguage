/*
Exercise 4.4 - Rotate function that operates in a single pass.
The equivalent rotation distance is n % len(slice).

$ go run src/ch4/ex_4_4/main.go
leftRotate([0 1 2 3 4 5], 3) = [3 4 5 0 1 2]
leftRotate([0 1 2 3 4 5], 7) = [1 2 3 4 5 0]
rightRotate([0 1 2 3 4 5], 2) = [4 5 0 1 2 3]
rightRotate([0 1 2 3 4 5], 8) = [4 5 0 1 2 3]
*/

package main

import "fmt"

func main() {
    a := [...]int{0, 1, 2, 3, 4, 5}

    fmt.Printf("leftRotate(%v, %d) = %v\n", a, 3, leftRotate(a[:], 3))
    fmt.Printf("leftRotate(%v, %d) = %v\n", a, 7, leftRotate(a[:], 7))   // Same as leftRotate(a, 1)

    fmt.Printf("rightRotate(%v, %d) = %v\n", a, 2, rightRotate(a[:], 2))
    fmt.Printf("rightRotate(%v, %d) = %v\n", a, 8, rightRotate(a[:], 8)) // Same as rightRotate(a, 2)
}

func leftRotate(s []int, n int) []int {
    r := n % len(s)
    return append(s[r:], s[:r]...)
}

func rightRotate(s []int, n int) []int {
    r := n % len(s)
    return append(s[len(s) - r:], s[:len(s) - r]...)
}
