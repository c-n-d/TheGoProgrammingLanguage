/*
Pipline1 demonstrates a infinite 3 stage pipeline

$ go run src/ch8/pipeline1/main.go
0
1
4
9
16
25
36
49
64
81
100
121
...
*/

package main

import "fmt"

func main() {
    naturals := make(chan int)
    squares  := make(chan int)

    // Counter
    go func() {
        for x := 0; ; x++ {
            naturals <- x
        }
    }()

    // Squarer
    go func() {
        for {
            x := <-naturals
            squares <- x * x
        }
    }()

    // Printer (in main goroutine)
    for {
        fmt.Println(<-squares)
    }
}
