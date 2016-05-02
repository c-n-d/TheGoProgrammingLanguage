/*
Crawl1 crawls web links starting with the command line args.

Exhausts available file describtors with many concurrent calls to links.Extract

Never terminates due to worklist never being closed

$ go run src/ch8/crawl1/main.go https://golang.org
https://golang.org
https://golang.org/
http://www.google.com/intl/en/policies/privacy/
http://play.golang.org/
...
2016/04/29 10:55:49 Get ...: dial tcp: lookup github.com: no such host
2016/04/29 10:55:49 Get ...: lookup ...: no such host
2016/04/29 10:55:49 Get ...: dial tcp [...]:443: socket: too many open files
...
*/

package main

import (
    "fmt"
    "log"
    "os"

    "ch5/links"
)

func main() {
    worklist := make(chan []string)

    // Starts with the command line argument
    go func() { worklist <- os.Args[1:] }()

    seen := make(map[string]bool)
    for list := range worklist {
        for _, link := range list {
            if !seen[link] {
                seen[link] = true
                go func(link string) {
                    worklist <- crawl(link)
                }(link)
            }
        }
    }
}

func crawl(url string) []string {
    fmt.Println(url)
    list, err := links.Extract(url)
    if err != nil {
        log.Print(err)
    }
    return list
}
