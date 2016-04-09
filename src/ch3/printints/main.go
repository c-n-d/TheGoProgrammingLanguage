/*
intsToString is like Sprintf(values) but adds commas

$ go run src/ch3/printints/main.go 
[1, 2, 3]
*/

package main

import (
    "bytes"
    "fmt"
)

func main() {
    fmt.Println(intsToString([]int{1, 2, 3}))
}

func intsToString(values []int) string {
    var buf bytes.Buffer
    buf.WriteByte('[')

    for i, v := range values {
        if i > 0 {
            buf.WriteString(", ")
        }
        fmt.Fprintf(&buf, "%d", v)
    }

    buf.WriteByte(']')
    return buf.String()
}
