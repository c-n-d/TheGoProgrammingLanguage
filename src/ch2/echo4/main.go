/*
Echo4 prints its command line arguments

$ go run src/ch2/echo4/main.go pwd ls cd mkdir
pwd ls cd mkdir

$ go run src/ch2/echo4/main.go -s ", " pwd ls cd mkdir
pwd, ls, cd, mkdir

$ go run src/ch2/echo4/main.go -n -s ", " pwd ls cd mkdir
pwd, ls, cd, mkdir$

$ go run src/ch2/echo4/main.go --help
Usage of main:
  -n    Omit trailing newline
  -s string
        Separator (default " ")
*/

package main

import (
    "flag"
    "fmt"
    "strings"
)

var n = flag.Bool("n", false, "Omit trailing newline")
var sep = flag.String("s", " ", "Separator")

func main() {
    flag.Parse()
    fmt.Print(strings.Join(flag.Args(), *sep))

    if !*n {
        fmt.Println()
    }
}
