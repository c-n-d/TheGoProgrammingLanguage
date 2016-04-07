/*
Echo2 prints all of the command line arguments

$ go run src/ch1/echo2/main.go pwd ls cd mkdir
pwd ls cd mkdir
*/

package main

import (
    "os"
    "fmt"
)

func main() {
    s, sep := "", ""

    // The '_' here is the blank identifier
    // Go does not allow unused variable, so we cannot declare a tmp variable for the index

    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }

    fmt.Println(s)
}