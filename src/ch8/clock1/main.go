/*
Clock1 is a TCP server that periodically writes the time.

$ go run src/ch8/clock1/main.go &
[1] PID
$ nc localhost 8000
16:46:12
16:46:13
16:46:14
16:46:15
16:46:16
16:46:17
*/

package main

import (
    "io"
    "log"
    "net"
    "time"
)

func main() {
    listener, err := net.Listen("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Print(err) // e.g. conncetion aborted
            continue
        }
        handleConn(conn) // handle one connection at a time
    }
}

func handleConn(c net.Conn) {
    defer c.Close()
    for {
        _, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
        if err != nil {
            return // e.g. client disconnected
        }
        time.Sleep(1 * time.Second)
    }
}
