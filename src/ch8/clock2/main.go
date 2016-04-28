/*
Clock2 is a TCP server that periodically writes the time. (concurrent connections)

$ go build ch8/clock2
$ go build ch8/netcat1
$ ./netcat1 
17:05:59
17:06:00        $ ./netcat1
17:06:01        17:06:01
17:06:02        17:06:02
17:06:03        17:06:03
17:06:04        17:06:04
17:06:05        17:06:05
17:06:06        ^C
17:06:07        $ ./netcat1
17:06:08        17:06:08
17:06:09        17:06:09
17:06:10        17:06:10
17:06:11        17:06:11
17:06:12        ^C
17:06:13
^C
$ killall clock2
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
