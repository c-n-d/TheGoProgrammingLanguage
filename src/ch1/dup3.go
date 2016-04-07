/*
Dup3 prints the count and text of lines that appear more than once
in the input. Reads from a list of named files

$ go run dup3.go ../../data/ch1/dup_1.dat ../../data/ch1/dup_2.dat
4    d
5    e
2    cat
3    dog
4    fish
2    b
3    c
*/

package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    // Makes a new map of {string => int}
    counts := make(map[string]int)

    for _, filename := range os.Args[1:] {
        // Reads the whole file as a byte slice
        data, err := ioutil.ReadFile(filename)

        if err != nil {
            fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
            continue
        }

        for _, line := range strings.Split(string(data), "\n") {
            counts[line]++
        }

    }

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
