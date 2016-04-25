/*
ByteCounter demonstrates an implementation of io.Writer that counts bytes.

$ go run src/ch7/bytecounter/main.go
5
12
*/

package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
    *c += ByteCounter(len(p))
    return len(p), nil
}

func main() {
    var c ByteCounter
    c.Write([]byte("hello"))
    fmt.Println(c)

    c = 0
    var name = "Dolly"
    fmt.Fprintf(&c, "hello, %s", name)
    fmt.Println(c)
}
