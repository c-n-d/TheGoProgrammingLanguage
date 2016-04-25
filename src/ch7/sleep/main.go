/*
The sleep program sleeps for a specified period of time

$ go run src/ch7/sleep/main.go -period "5s"
Sleeping for 5s...
$ go run src/ch7/sleep/main.go -period "1m5s"
Sleeping for 1m5s...
*/

package main

import (
    "flag"
    "fmt"
    "time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
    flag.Parse()
    fmt.Printf("Sleeping for %v...", *period)
    time.Sleep(*period)
    fmt.Println()
}
