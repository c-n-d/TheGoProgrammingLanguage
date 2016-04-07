/*
Dup2 prints the count and text of lines that appear more than once
in the input. It reads from stdin or from a list of named files

$ go run dup2.go
cat
cat
dog
dog
^D
2    dog
2    cat

$ go run dup2.go ../../data/ch1/dup_1.dat ../../data/ch1/dup_2.dat
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
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]

    if len(files) == 0 {
        countLines(os.Stdin, counts)
    } else {
        for _, file := range files {
            f, err := os.Open(file)

            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }

            countLines(f, counts)
            f.Close()
        }
    }

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}

func countLines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)

    // NOTE: Ignoring potential errors from input.Err
    for input.Scan() {
        counts[input.Text()]++
    }
}
