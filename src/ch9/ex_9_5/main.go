/*
Exercise 9.5 - Write a program w/ two goroutines that send messages back and forth over two
               unbuffered channels. How many communications per second?

$ go run src/ch9/ex_9_5/main.go
msg / second:
276216.15115

$ go run src/ch9/ex_9_5/main.go
msg / second:
276216.15115

*/

package main

import (
    "fmt"
    "sync"
    "time"
)

type Message struct {
    last time.Time
}

var wg sync.WaitGroup

func main() {
    wg.Add(1)

    ping   := make(chan int)
    pong   := make(chan int)
    notify := make(chan struct{})

    // reads from pong, writes to ping
    go stage(ping, pong, notify)
    // reads from ping, writes to pong
    go stage(pong, ping, notify)

    go timekeeper(notify)

    ping <- 42

    wg.Wait()
}

func stage(out chan<- int, in <-chan int, notify chan<- struct{}) {
    for input := range in {
        notify <- struct{}{}
        out <- input
    }
    close(out)
}

func timekeeper(notify <-chan struct{}) {
    <-notify

    start := time.Now()
    count := float64(1)

    fmt.Println("msg / second:")
    for {
        fmt.Printf("\r%5.5f", (count / time.Since(start).Seconds()))
        <-notify
        count++
    }
    wg.Done()
}
