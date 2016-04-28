/*
Netcat1 is a read only TCP client

$ go build ch8/clock1
$ go build ch8/netcat1
$ ./clock1 &

$ ./netcat1
16:55:34     $ ./netcat1
16:55:35
16:55:36
16:55:37
16:55:38
^C
             16:55:40
             16:55:41
             16:55:42
             16:55:43
             16:55:44
             ^C
$ killall clock1
*/

package main

import (
    "io"
    "log"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
    if _, err := io.Copy(dst, src); err != nil {
        log.Fatal(err)
    }
}
