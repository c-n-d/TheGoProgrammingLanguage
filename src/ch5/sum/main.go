/*
Sum demonstrates variadic functions (...).

$ go run src/ch5/sum/main.go
0
3
10
15
*/

package main

import "fmt"

func sum(vals...int) int {
    total := 0
    for _, val := range vals {
        total+=val
    }
    return total
}

func main() {
    fmt.Println(sum())
    fmt.Println(sum(3))
    fmt.Println(sum(1,2,3,4))

    values := []int{1,2,3,4,5}
    fmt.Println(sum(values...))
}
