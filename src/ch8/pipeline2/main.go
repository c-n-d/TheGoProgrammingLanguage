/*
Pipline2 demonstrates a finite 3 stage pipeline

$ go run src/ch8/pipeline2/main.go
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
...
8836
9025
9216
9409
9604
9801
*/

package main

import "fmt"

func main() {
    naturals := make(chan int)
    squares  := make(chan int)

    // Counter
    go func() {
        for x := 0; x < 100 ; x++ {
            naturals <- x
        }
        close(naturals)
    }()

    // Squarer
    go func() {
        for x := range naturals {
            squares <- x * x
        }
        close(squares)
    }()

    // Printer (in main goroutine)
    for x := range squares{
        fmt.Println(x)
    }
}
