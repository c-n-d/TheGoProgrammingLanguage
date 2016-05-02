/*
Countdown implements a countdown for rocket launch

$ go run src/ch8/countdown2/main.go
Commencing countdown. Press return to abort.
Lift off!
$ go run src/ch8/countdown2/main.go
Commencing countdown. Press return to abort.

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
    select {
        case <-time.After(10 * time.Second):
            // do nothing
        case <-abort:
            fmt.Println("Launch aborted!")
            return
    }
    launch()
}

func launch() {
    fmt.Println("Lift off!")
}
