/*
Exercise 7.1 - Implement a counter for words and for lines.

$ go run src/ch7/ex_7_1/main.go
wc.Write(eightWords), wc=8
2 * fmt.Fprintf(&wc, eightWords), wc=24

lc.Write(fiveLines), lc=5
2 * fmt.Fprintf(&lc, fiveLines), lc=15

*/

package main

import (
    "bufio"
    "fmt"
    "strings"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
    var count int
    reader  := strings.NewReader(string(p))
    scanner := bufio.NewScanner(reader)

    scanner.Split(bufio.ScanWords)

    for ;scanner.Scan(); scanner.Text() {
        if scanner.Err() != nil {
            return 0, scanner.Err()
        }
        count++
    }

    *c += WordCounter(count)
    return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
    var count int
    reader  := strings.NewReader(string(p))
    scanner := bufio.NewScanner(reader)

    for ;scanner.Scan(); scanner.Text() {
        if scanner.Err() != nil {
            return 0, scanner.Err()
        }
        count++
    }

    *c += LineCounter(count)
    return len(p), nil
}

var eightWords = `One Two Three Four
Five Six Seven Eight
`

var fiveLines = `Line one
Line two
Line three
Line four
Line five
`

func main() {
    var wc WordCounter
    var lc LineCounter

    wc.Write([]byte(eightWords))
    fmt.Printf("wc.Write(eightWords), wc=%v\n", wc)
    fmt.Fprintf(&wc, eightWords)
    fmt.Fprintf(&wc, eightWords)
    fmt.Printf("2 * fmt.Fprintf(&wc, eightWords), wc=%v\n\n", wc)

    lc.Write([]byte(fiveLines))
    fmt.Printf("lc.Write(fiveLines), lc=%v\n", lc)
    fmt.Fprintf(&lc, fiveLines)
    fmt.Fprintf(&lc, fiveLines)
    fmt.Printf("2 * fmt.Fprintf(&lc, fiveLines), lc=%v\n", lc)
}
