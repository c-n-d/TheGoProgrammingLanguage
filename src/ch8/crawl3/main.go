/*
Crawl3 crawls web links starting with the command line args.

Uses bounded parallelism, 20 active HTTP requests

$ go run src/ch8/crawl3/main.go https://golang.org
https://golang.org
https://golang.org/project/
https://golang.org/doc/
http://www.google.com/intl/en/policies/privacy/
https://golang.org/dl/
https://golang.org/pkg/
https://golang.org/help/
https://golang.org/
http://play.golang.org/
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
    worklist := make(chan []string)  // list of URLs, may have duplicates
    unseenLinks := make(chan string) // de-duplicate URLs

    // Starts with the command line argument
    go func() { worklist <- os.Args[1:] }()

    for i := 0; i < 20; i++ {
        go func() {
            for link := range unseenLinks {
                foundLinks := crawl(link)
                go func() { worklist <- foundLinks }()
            }
        }()
    }

   // The main goroutine de-duplicates worklist items
    seen := make(map[string]bool)
    for list := range worklist {
        for _, link := range list {
            if !seen[link] {
                seen[link] = true
                unseenLinks <- link
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
