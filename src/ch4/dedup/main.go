/*
Dedupe prints the line only the first time it has been seen

$ go run src/ch4/dedup/main.go
a
New: a
a
b
New: b
b
b
b
c
New: c
c
c
*/

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    seen := make(map[string]bool)
    input := bufio.NewScanner(os.Stdin)

    for input.Scan() {
        line := input.Text()
        if !seen[line] {
            seen[line] = true
            fmt.Println("New: " + line)
        }
    }

    if err := input.Err(); err !=nil {
        fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
        os.Exit(1)
    }
}
