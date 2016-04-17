/*
Exercise 5.9 - Write a function expand that replaces each substring 
               "$foo" within s bt the text returned by f("foo").

$ go run src/ch5/ex_5_9/main.go
expand(car zoo cat dog, reverse)= rac ooz tac god
expand(car zoo cat dog, double)=car car zoo zoo cat cat dog dog
*/

package main

import (
    "fmt"
    "strings"
)

func main() {
    orig := "car zoo cat dog"
    fmt.Printf("expand(%s, reverse)=%s\n", orig, expand(orig, reverse))
    fmt.Printf("expand(%s, double)=%s\n", orig, expand(orig, double))
}

func expand(s string, f func(string) string) string {
    m :=  4 - len(s) % 4
    s = rightPad(s, m)
    var res string

    for i := 0; i < len(s); i += 4 {
        res += f(s[i:i+4])
    }

    return res
}

func reverse(str string) string {
    s := strings.Split(str, "")
    for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
        s[i], s[j] = s[j], s[i]
    }
    return strings.Join(s, "")
}

func double(str string) string {
    return strings.Repeat(str, 2)
}

func rightPad(s string, nc int) string {
    return s + strings.Repeat(" ", nc)
}
