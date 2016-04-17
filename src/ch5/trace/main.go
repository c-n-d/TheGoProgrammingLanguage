/*
Trace uses defer to add entry/exit diagnostics to a function

$ go run src/ch5/trace/main.go
2016/04/17 01:12:34 enter bigSlowOperation
2016/04/17 01:12:44 exit bigSlowOperation, (10.001853263s)
*/

package main

import (
    "log"
    "time"
)

func bigSlowOperation() {
    defer trace("bigSlowOperation")()
    //...lots of work...
    time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
    start := time.Now()
    log.Printf("enter %s", msg)
    return func() { log.Printf("exit %s, (%s)", msg, time.Since(start)) }
}

func main() {
    bigSlowOperation()
}
