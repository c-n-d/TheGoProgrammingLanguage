/*
Exercise 8.3 - Modify netcat3 to close only the write half of the TCP connection so the program
               may continue to print the final echos after stdin has been closed.

$ go build -o reverb1 ch8/reverb1
$ ./reverb1
$ go run src/ch8/ex_8_3/main.go
Hello,
	 HELLO,
World!
	 Hello,
^D       hello,
	 WORLD!
	 World!
	 world!
2016/04/28 17:32:33 done
$killall reverb1
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
    done := make(chan struct{})
    go func() {
        io.Copy(os.Stdout, conn) // Note: Ignoring errors
        log.Println("done")
        done <- struct{}{} // signal to the main goroutine
    }()
    mustCopy(conn, os.Stdin)
    closeWrite(conn)
    <- done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
    if _, err := io.Copy(dst, src); err != nil {
        log.Fatal(err)
    }
}

func closeWrite(conn net.Conn) {
    if conn, ok := conn.(*net.TCPConn); ok {
        conn.CloseWrite()
    } else {
        log.Println("Unable to close writer using conn as net.TCPConn.")
    }
}
