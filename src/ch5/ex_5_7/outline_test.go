/*
Test for Exercise 5.7 - HTML pretty printer.
1. Calls outline for golang.org and storing the pretty printed HTML in a buffer
2. Attempts to parse the buffer
3. Reports html parse failures

$ cd src/ch5/ex_5_7/
$ go test -v -bench=.
=== RUN   TestOutlinePrettyPrint
--- PASS: TestOutlinePrettyPrint (0.14s)
PASS
ok  	ch5/ex_5_7	0.150s
*/

package main

import (
    "bytes"
    "bufio"
    "testing"
    "strings"

    "golang.org/x/net/html"
)

func TestOutlinePrettyPrint(t *testing.T) {
    var b bytes.Buffer
    writer := bufio.NewWriter(&b)

    outline("https://golang.org/", writer)

    writer.Flush()

    htmlStr := strings.NewReader(b.String())

    _, err := html.Parse(htmlStr)

    if err != nil {
        t.Errorf("Failed to parse HTML from outline: %v\n", err)
    }
}
