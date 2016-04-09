/*
Basename1 removes directory components and a .suffix

$ go run src/ch3/basename1/main.go
bash
install
*/

package main

import "fmt"

func main() {
    fmt.Println(basename("/usr/bin/bash"))
    fmt.Println(basename("/opt/prog/bin/install.sh"))
}

func basename(s string) string {
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '/' {
            s = s[i+1:]
            break
        }
    }

    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '.' {
            s = s[:i]
            break
        }
    }

    return s
}
