/*
Reverb1 is a TCP server that simulates an echo

$ go build -o reverb1 ch8/reverb1
$ ./reverb1 &
$ go run src/ch8/netcat2/main.go
Hello!
	 HELLO!
There!
	 Hello!
	 hello!
	 THERE!
	 There!
	 there!
^D
$ killall reverb1
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
    for input.Scan() {
        echo(c, input.Text(), 1*time.Second)
    }
    // Note: Ignoring potential errors from input.Err()
    c.Close()
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
