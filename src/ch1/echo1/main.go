/*
Echo1 prints all of the command line arguments

$ go run src/ch1/echo1/main.go pwd ls cd mkdir
pwd ls cd mkdir
*/

package main

import (
    "fmt"
    "os"
)

func main() {
    var s, sep string

    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }

    fmt.Println(s)
}
