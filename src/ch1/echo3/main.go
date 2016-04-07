/*
Echo3 prints all of the command line arguments

$ go run src/ch1/echo3/main.go pwd ls cd mkdir
pwd ls cd mkdir
*/

package main

import (
    "os"
    "fmt"
    "strings"
)

func main() {

    fmt.Println(strings.Join(os.Args[1:], " "))

}