/*
Exercise 8.8 -

$ go build ch8/ex_8_8
$ ./ex_8_8 &
$ go run src/ch8/netcat2/main.go
2016/05/01 16:20:53 dial tcp [::1]:8000: getsockopt: connection refused
exit status 1
$ killall ex_8_8
*/

package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "time"
    "strings"
)

func echo(c net.Conn, shout string, delay time.Duration) {
    fmt.Fprintln(c, "\t", strings.ToUpper(shout))
    time.Sleep(delay)
    fmt.Fprintln(c, "\t", shout)
    time.Sleep(delay)
    fmt.Fprintln(c, "\t", strings.ToLower(shout))
    time.Sleep(delay)
}

func handleConn(c net.Conn) {
    input := bufio.NewScanner(c)

    read := make(chan string)
    go func() {
        for input.Scan() {
            read <- input.Text()
        }
        if err := input.Err(); err != nil {
            log.Fatal(err)
        }
    }()

    select {
        case <-time.After(10 * time.Second):
            c.Close()
            return
        case text := <-read:
            go echo(c, text, 1*time.Second)
    }
}

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
