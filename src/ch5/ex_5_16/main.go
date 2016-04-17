/*
Exercise 5.16 - Write a veriadic version of strings.Join.

$ go run src/ch5/ex_5_16/main.go
join(", ", [cat dog fish bird])="cat, dog, fish, bird"
join(", ", [])=""
*/

package main

import "fmt"

// Join logic from std lib https://golang.org/src/strings/strings.go?s=9002:9042#L341
func join(sep string, values...string) string {
    if len(values) == 0 {
        return ""
    }

    if len(values) == 1 {
        return values[0]
    }

    n := len(sep) * (len(values) - 1)
    for i := 0; i < len(values); i++ {
        n += len(values[i])
    }

    res := make([]byte, n)
    resp := copy(res, values[0])
    for _, s := range values[1:] {
        resp += copy(res[resp:], sep)
        resp += copy(res[resp:], s)
    }
    return string(res)
}

func main() {
    sep := ", "
    values := []string{"cat", "dog", "fish", "bird"}
    empty := []string{}

    fmt.Printf("join(%q, %v)=%q\n", sep, values, join(sep, values...))
    fmt.Printf("join(%q, %v)=%q\n", sep, empty, join(sep, empty...))
}
