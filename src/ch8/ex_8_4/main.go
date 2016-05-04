/*
Exercise 8.4 - Modify reverb2 to use a sync.WaitGroup per connection to count the number
               of active echo commands. When it falls to zero close the write half of the TCP
               connection. Ensure that netcat from ex8.3 waits for the final echo from concurrent
               shouts, even after stdin has been closed

$ go build ch8/ex_8_4
$ ./ex_8_4 &
$ go run src/ch8/ex_8_3/main.go
Hello 
	 HELLO
There!
	 THERE!
	 Hello
^D       There!

	 hello
	 there!
2016/04/29 09:51:52 done
$ killall ex_8_4
*/

package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "time"
    "strings"
    "sync"
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
            defer wg.Done()
            fmt.Fprintln(c, "\t", strings.ToUpper(shout))
            time.Sleep(delay)
            fmt.Fprintln(c, "\t", shout)
            time.Sleep(delay)
            fmt.Fprintln(c, "\t", strings.ToLower(shout))
            time.Sleep(delay)
}

// WaitGroup is used to count the number of open echo goroutines,
// allowing the last one to complete even after stdin is closed
func handleConn(c net.Conn) {
    var wg sync.WaitGroup
    input := bufio.NewScanner(c)
    for input.Scan() {
        wg.Add(1)
        go echo(c, input.Text(), 1*time.Second, &wg)
    }
    // Note: Ignoring potential errors from input.Err()
    wg.Wait()
    c.(*net.TCPConn).CloseWrite()
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
