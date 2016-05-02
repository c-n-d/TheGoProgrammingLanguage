/*
Pipline3 demonstates a 3 stage pipeline with range, close, and unidirectional channel types.

$ go run src/ch8/pipeline3/main.go
$ go run src/ch8/pipeline3/main.go
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
8649
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

    go counter(naturals)
    go squarer(squares, naturals)
    printer(squares)
}

func counter(out chan<- int) {
    for x := 0; x < 100 ; x++ {
        out <- x
    }
    close(out)
}

func squarer(out chan<- int, in <-chan int) {
    for x := range in {
        out <- x * x
    }
    close(out)
}

func printer(in <-chan int) {
    for x := range in{
        fmt.Println(x)
    }
}
