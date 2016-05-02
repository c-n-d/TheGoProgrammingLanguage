/*
Countdown implements a countdown for rocket launch

$ go run src/ch8/countdown1/main.go
Commencing countdown.
10
9
8
7
6
5
4
3
2
1
Lift off!
*/

package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Commencing countdown.")
    tick := time.Tick(1 * time.Second)
    for countdown := 10; countdown > 0; countdown-- {
        fmt.Println(countdown)
        <- tick
    }
    launch()
}

func launch() {
    fmt.Println("Lift off!")
}
