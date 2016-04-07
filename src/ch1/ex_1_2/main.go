/*
Exercise 1.2 prints all command line args and their index - one per line

$ go run src/ch1/ex_1_2/main.go pwd ls cd mkdir
Index: 0. Arg: {...}/ex_1_2.
Index: 1. Arg: pwd.
Index: 2. Arg: ls.
Index: 3. Arg: cd.
Index: 4. Arg: mkdir.
*/

package main

import (
    "os"
    "fmt"
)

func main() {
    // Range provides the index and the value
    for index, arg := range os.Args {
        fmt.Printf("Index: %d. Arg: %s.\n", index, arg)
    }
}