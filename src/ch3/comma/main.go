/*
Comma inserts commas in a non-negative decimal integer string

$ go run src/ch3/comma/main.go
1,234
1,234,567
1,234,567,890
*/

package main

import "fmt"

func main() {
    fmt.Println(comma("1234"))
    fmt.Println(comma("1234567"))
    fmt.Println(comma("1234567890"))
}

func comma(s string) string {
    n := len(s)

    if n <= 3 {
        return s
    }

    return comma(s[:n-3]) + ","  + s[n-3:]
}
