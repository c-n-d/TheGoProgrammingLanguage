/*
Boiling prints the boiling point of water.

$ go run src/ch2/boiling/main.go
boiling point = 212°F or 100°C
*/

package main

import "fmt"

const boilingF = 212.0

func main() {
    var f = boilingF
    var c = (f - 32) * 5 / 9

    fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
