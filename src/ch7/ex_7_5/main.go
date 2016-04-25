/*
Exercise 7.5 - LimitReader function in the io package accepts a reader and n bytes, reporting EOF after n bytes. Implement it.

$ go run src/ch7/ex_7_5/main.go -l 55 -b 10
Limit: [55], Buffer Size: [10], Read String: [abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz]

abcdefghij
klmnopqrst
uvwxyzabcd
efghijklmn
opqrstuvwx
yzabc
*/

package main

import (
    "flag"
    "fmt"
    "io"
    "strings"
)

type LReader struct {
    reader *io.Reader
    limit *int64
    progress *int64
}

func (lr LReader) Read(p []byte) (n int, err error) {
    if *lr.progress >= *lr.limit {
        return 0, io.EOF
    }
    if int64(len(p)) + *lr.progress <= *lr.limit {
        n, err = (*lr.reader).Read(p)
    } else {
        bufSize := *lr.limit - *lr.progress
        tmp := make([]byte, bufSize)
        n, err = (*lr.reader).Read(tmp)
        copy(p, tmp)
    }
    *lr.progress += int64(n)
    return
}

var limit   = flag.Int64("l", 8, "Limit: the total number of bytes to read")
var bufSize = flag.Int("b", 1, "BufferSize: size of buffer to which bytes are read")

func init() {
    flag.Parse()
}

func main() {
    buf   := make([]byte, *bufSize)
    empty := make([]byte, *bufSize)
    str   := strings.Repeat("abcdefghijklmnopqrstuvwxyz", 1 + int(*limit) / 26)

    stringReader := strings.NewReader(str)
    limitReader  := LimitReader(stringReader, *limit)

    fmt.Printf("Limit: [%d], Buffer Size: [%d], Read String: [%s]\n\n", *limit, *bufSize, str)

    for  _, err := limitReader.Read(buf); err != io.EOF; _, err = limitReader.Read(buf) {
        fmt.Println(string(buf))
        copy(buf, empty)
    }
}

func LimitReader(r io.Reader, n int64) io.Reader {
    lim := n
    var prog int64
    return LReader{&r, &lim, &prog}
}
