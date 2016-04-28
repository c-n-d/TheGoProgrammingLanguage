/*
Exercise 8.1 - Modify clock2 to accept a port. Write `clockwall` that acts as a client of multiple
               clock servers and displays the result in a table.

$ go build -o clock_8_1 src/ch8/ex_8_1/main.go
$ TZ=US/Eastern    ./clock_8_1 -port 8010 &
$ TZ=Europe/London ./clock_8_1 -port 8020 &
$ TZ=Asia/Tokyo    ./clock_8_1 -port 8030 &
...
$ killall clock_8_1
*/

package main

import (
    "flag"
    "fmt"
    "io"
    "log"
    "net"
    "time"
)

var port = flag.Int("port", 8000, "tcp port to listen on")

func init() {
    flag.Parse()
}

func main() {
    address := fmt.Sprintf("localhost:%d", *port)
    listener, err := net.Listen("tcp", address)
    if err != nil {
        log.Fatal(err)
    }
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Print(err) // e.g. conncetion aborted
            continue
        }
        go handleConn(conn) // handle connections concurrently
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
