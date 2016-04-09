/*
Basename2 also removes directory components and a .suffix, but uses string operations

$ go run src/ch3/basename2/main.go
bash
install
*/

package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(basename("/usr/bin/bash"))
    fmt.Println(basename("/opt/prog/bin/install.sh"))
}

func basename(s string) string {
    slash := strings.LastIndex(s, "/")
    s = s[slash+1:]

    if dot := strings.LastIndex(s, "."); dot >=0 {
        s = s[:dot]
    }

    return s
}
