/*
Nonempty is an example of an in-place slice algorithm

$ go run src/ch4/nonempty/main.go
nonempty(["one" "three" "three"]) = ["one" "three"]
nonempty(["four" "six" "six"]) = ["four" "six"]
*/

package main

import "fmt"

func main() {
    data := []string{"one", "", "three"}
    data1 := []string{"four", "", "six"}

    // The nonempty functions change the underlying array; that is why
    // data and data1 appear to contain 'duplicate' elements
    fmt.Printf("nonempty(%q) = %q\n", data, nonempty(data))
    fmt.Printf("nonempty(%q) = %q\n", data1, nonempty(data1))
}

// nonempty returns a slice holding all the non-empty strings.
// The underlying array is modified during the call
func nonempty(strings []string) []string {
    i := 0
    for _, s := range strings {
        if s != "" {
            strings[i] = s
            i++
        }
    }
    return strings[:i]
}

// nonempty using append
func nonempty2(strings []string) []string {
    out := strings[:0]
    for _, s := range strings {
        if s != "" {
            out = append(out, s)
        }
    }
    return out
}
