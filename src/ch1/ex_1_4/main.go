/*
Exercise 1.4 prints the count text of lines that appear more than once
in the input. Prints name of file containting duplicates. Reads from a list of named files

$ go run src/ch1/ex_1_4/main.go data/ch1/dup_1.dat data/ch1/dup_2.dat
2    cat
3    dog
4    fish
2    b
3    c
4    d
5    e
data/ch1/dup_1.dat has duplicates
data/ch1/dup_2.dat has duplicates
*/

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    filesWithDup := make(map[string]bool)

    files := os.Args[1:]

    if len(files) == 0 {
        countLines(os.Stdin, counts, filesWithDup, "/dev/stdin")
    } else {
        for _, file := range files {
            f, err := os.Open(file)

            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }

            countLines(f, counts, filesWithDup, file)
            f.Close()
        }
    }

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }

    for fileName, hasDups := range filesWithDup {
        if hasDups {
            fmt.Printf("%s has duplicates\n",fileName)
        }
    }
}

func countLines(f *os.File, counts map[string]int, filesWithDup map[string]bool, fileName string) {
    input := bufio.NewScanner(f)

    // NOTE: Ignoring potential errors from input.Err
    for input.Scan() {
        var text = input.Text();
        counts[text]++

        if counts[text] > 1 {
            filesWithDup[fileName] = true
        }
    }
}
