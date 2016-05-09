/*
Exercise 9.6 - Measure the performance of a compute bound parallel program w/ varying GOMAXPROCS.
               What's optimal for your computer? How many CPUs does your computer have?

$ GOMAXPROCS=1 go run src/ch9/ex_9_6/main.go
There are [4] logical CPUs usable by the current process.
∫ 5 + 5*Cos(x*π) + Sin(-x*π) over [0, 100000] ≈ 500000.00001812604
30.114347986s

*/

package main

import (
    "fmt"
    "math"
    "runtime"
    "sync"
    "time"
)

var wg sync.WaitGroup

func main() {
        fmt.Printf("There are [%d] logical CPUs usable by the current process.\n", runtime.NumCPU())
        start := time.Now()
        do()
        fmt.Println(time.Since(start))

}

func do() {
    begin, end := float64(0), float64(100000)
    delta := .05

    accChan  := make(chan float64)
    evalChan := make(chan float64)

    for x := begin; x < end; x += delta {
        wg.Add(1)
        go evaluate(x, delta, evalChan)
    }

    go accumulate(accChan, evalChan)

    wg.Wait()
    close(evalChan)

    fmt.Println("∫ 5 + 5*Cos(x*π) + Sin(-x*π) over [0, 100000] ≈", <- accChan)
}

func evaluate(x, delta float64, y chan<- float64) {
        y <- f(x) * delta
        wg.Done()
}

func f(x float64) float64 {
    return 5 + 5 * math.Cos(x*math.Pi) + math.Sin(-x*math.Pi)
}

func accumulate(result chan<- float64, part <-chan float64) {
    var total float64
    for val := range part {
        total += val
    }
    result <- total
}
