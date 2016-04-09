/*
Exercise 3.11 is an iteration on Exercise 3.10. Adds support for 
handling floating point numbers and an optional sign

$ go run src/ch3/ex_3_11/main.go 
enhancedComma(-123) = -123
enhancedComma(1234) = 1,234
enhancedComma(+1234) = +1,234
enhancedComma(-12345) = -12,345
enhancedComma(1234567.7654321) = 1,234,567.7654321
enhancedComma(-123456789.987654321) = -123,456,789.987654321
*/

package main

import (
    "fmt"
    "strings"
)

func main() {
    inputs := []string{"-123", "1234", "+1234", "-12345", "1234567.7654321", "-123456789.987654321"}

    for _, val := range inputs {
        fmt.Printf("enhancedComma(%s) = %s\n", val, enhancedComma(val))
    }
}

func enhancedComma(s string) string {
    s, sign := parseSign(s)
    s, fp := parseFloatingPoint(s)

    return sign + comma(s) + fp
}

// For a given numeric string, returns the sign and the remaining numeric portion
// ex. "-1234" => ("1234", "-")
func parseSign(s string) (string, string) {
    if s[0:1] == "+" || s[0:1] == "-" {
        return s[1:], s[0:1]
    }

    return s, ""
}

// For a given numeric string, returns the whole number and floating point separatly
// ex. "1234.567" => ("1234", ".567")
func parseFloatingPoint(s string) (string, string) {
    dot := strings.LastIndex(s, ".")

    if dot > 0 {
        return s[0:dot], s[dot:]
    }

    return s, ""
}

func comma(s string) string {
    n := len(s)

    if n <= 3 {
        return s
    }

    return comma(s[:n-3]) + ","  + s[n-3:]
}
