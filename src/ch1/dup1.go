/*
Dup1 prints the text of each line that appears more than
once in the standard input, preceded by its count

$ go run dup1.go
bird
cat
cat
dog
dog
dog
fish
fish
fish
fish
^D
2    cat
3    dog
4    fish
*/
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)

    // NOTE: Ignoring potential errors from input.Err
    for input.Scan() {
        counts[input.Text()]++
    }

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
