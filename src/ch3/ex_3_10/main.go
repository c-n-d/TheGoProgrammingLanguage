/*
Exercise 3.10 is an iterative version of Comma and uses bytes.Buffer

$ go run src/ch3/ex_3_10/main.go
comma(123) = 123
comma(1234) = 1,234
comma(123456) = 123,456
comma(1234567) = 1,234,567
comma(1234567890) = 1,234,567,890
*/

package main

import (
    "bytes"
    "fmt"
)

func main() {
    inputs := []string{"123", "1234", "123456", "1234567", "1234567890"}

    for _, val := range inputs {
        fmt.Printf("comma(%s) = %s\n", val, comma(val))
    }
}

func comma(s string) string {
    var buf bytes.Buffer
    start := len(s) % 3

    if start != 0 {
        buf.WriteString(s[0:start] + ",")
    }

    for i := start; i < len(s); i+=3 {
        buf.WriteString(s[i:i+3])

        if i + 3 < len(s) {
            buf.WriteString(",")
        }
    }

    return buf.String()
}
