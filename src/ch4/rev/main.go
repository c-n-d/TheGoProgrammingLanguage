/*
rev reverses a slice of ints in place

Additional samples from the reading on how to rotate a slice using reverse

$ go run src/ch4/rev/main.go
reverse([0 1 2 3 4 5]) = [5 4 3 2 1 0]
leftRotate([5 4 3 2 1 0], 3) = [2 1 0 5 4 3]
rightRotate([2 1 0 5 4 3], 2) = [4 3 2 1 0 5]
*/

package main

import "fmt"

func main() {
    a := [...]int{0, 1, 2, 3, 4, 5}

    fmt.Printf("reverse(%v) = ", a)
    reverse(a[:])
    fmt.Printf("%v\n", a)

    fmt.Printf("leftRotate(%v, %d) = ", a, 3)
    leftRotate(a[:], 3)
    fmt.Printf("%v\n", a)

    fmt.Printf("rightRotate(%v, %d) = ", a, 2)
    rightRotate(a[:], 2)
    fmt.Printf("%v\n", a)
}

func reverse(s []int) {
    for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
        s[i], s[j] = s[j], s[i]
    }
}

func leftRotate(s []int, n int) {
    reverse(s[:n])
    reverse(s[n:])
    reverse(s)
}

func rightRotate(s []int, n int) {
    reverse(s)
    reverse(s[:n])
    reverse(s[n:])
}
