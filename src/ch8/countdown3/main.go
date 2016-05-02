/*
Countdown implements a countdown for rocket launch

$ go run src/ch8/countdown3/main.go
Commencing countdown. Press return to abort.
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
$ go run src/ch8/countdown3/main.go
Commencing countdown. Press return to abort.
10
9
8
7

Launch aborted!
*/

package main

import (
    "fmt"
    "time"
    "os"
)

func main() {
    abort := make(chan struct{})
    go func() {
        os.Stdin.Read(make([]byte, 1)) // read a single byte
        abort <- struct{}{}
    }()

    fmt.Println("Commencing countdown. Press return to abort.")
    tick := time.Tick(1 * time.Second)
    for countdown := 10; countdown > 0; countdown-- {
        fmt.Println(countdown)
        select {
            case <- tick:
                // do nothing
            case <-abort:
                fmt.Println("Launch aborted!")
                return
        }
    }
    launch()
}

func launch() {
    fmt.Println("Lift off!")
}
