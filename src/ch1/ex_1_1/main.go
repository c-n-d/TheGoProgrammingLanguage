/*
Exercise 1.1 prints all command line args, including the program name

$ go run src/ch1/ex_1_1/main.go pwd ls cd mkdir
{...}/ex_1_1 pwd ls cd mkdir
*/

package main

import (
    "os"
    "fmt"
    "strings"
)

func main() {
    // os.Args[0] contains the program name.
    // Unlike echo1-3, we don't need to take a portion of the slice.
    fmt.Println(strings.Join(os.Args, " "))
}