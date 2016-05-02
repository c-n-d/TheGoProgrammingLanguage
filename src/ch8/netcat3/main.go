/*
Netcat is a simple read/write for TCP servers

$ go build -o reverb2 ch8/reverb2
$ ./reverb2
$ go run src/ch8/netcat3/main.go
Hello, 
	 HELLO,
World!
	 WORLD!
	 Hello,
	 World!
	 hello,
	 world!
^D
2016/04/28 16:07:51 done
$killall reverb2
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
    conn.Close()
    <- done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
    if _, err := io.Copy(dst, src); err != nil {
        log.Fatal(err)
    }
}
