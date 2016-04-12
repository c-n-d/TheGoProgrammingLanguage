/*
Exercise 4.9 - wordfreq reports the frequency of words in an input tect file

$ go run src/ch4/ex_4_9/main.go data/ch4/wordfreq.dat
word           |   count
---------------|-------------
officia        |    1
do             |    1
enim           |    1
cillum         |    1
quis           |    1
ea             |    1
amet,          |    1
consectetur    |    1
magna          |    1
sed            |    1
...
*/

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    filename := os.Args[1]
    wordfreq := make(map[string]int)

    file, err := os.Open(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "freqcount: %v\n", err)
        os.Exit(1)
    }

    scanner := bufio.NewScanner(bufio.NewReader(file))
    scanner.Split(bufio.ScanWords)

    for scanner.Scan() {
        word := scanner.Text()
        wordfreq[word]++
    }

    fmt.Printf("%-15s|   %5s\n", "word", "count")
    fmt.Printf("---------------|-------------\n")

    for k, v := range wordfreq {
        fmt.Printf("%-15s|%5d\n", k, v)
    }
}
