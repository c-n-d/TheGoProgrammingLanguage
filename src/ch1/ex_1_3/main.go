/*
Exercise 1.3 runs a naive benchmark on two methods of concatenating the command line args

We'll skip Go's build in testing & benchmarking for now, but it will make an appearance in an upcoming exercise

$ go run src/ch1/ex_1_3/main.go pwd ls cd mkdir
String concat - 10000000 iterations - 3083 ms elapsed
String join - 10000000 iterations - 1851 ms elapsed
*/
package main

import (
    "os"
    "strings"
    "time"
    "fmt"
)

func main() {
    var times = 10000000

    var startNaive = time.Now()

    for i := 0; i < times; i++ {
        naiveArgsEcho()
    }

    fmt.Printf("String concat - %d iterations - %d ms elapsed\n",
                times,
                time.Since(startNaive).Nanoseconds() / 1000 / 1000)

    var startStringJoin = time.Now()

    for i := 0; i < times; i++ {
        joinArgsEcho()
    }

    fmt.Printf("String join - %d iterations - %d ms elapsed\n",
                times,
                time.Since(startStringJoin).Nanoseconds() / 1000 / 1000)
}

func naiveArgsEcho() {
    var s, sep string

    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
}

func joinArgsEcho() {
    strings.Join(os.Args[1:], " ")
}