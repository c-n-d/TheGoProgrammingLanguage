/*
Crawl2 crawls web links starting with the command line args.

Uses a buffered channel as a counting semaphore to limit the number of
concurrent calls to links.Extract

$ go run src/ch8/crawl2/main.go https://golang.org
https://golang.org
https://golang.org/help/
https://golang.org/doc/
https://golang.org/
http://www.google.com/intl/en/policies/privacy/
https://golang.org/pkg/
https://golang.org/project/
https://golang.org/dl/
https://developers.google.com/site-policies#restrictions
https://golang.org/blog/
http://play.golang.org/
https://blog.golang.org/
https://tour.golang.org/
https://golang.org/LICENSE
https://golang.org/doc/tos.html
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
    var n int // number of pending sends to worklist

    // Starts with the command line argument
    n++
    go func() { worklist <- os.Args[1:] }()

    seen := make(map[string]bool)
    for ; n > 0; n-- {
        list := <-worklist
        for _, link := range list {
            if !seen[link] {
                seen[link] = true
                n++
                go func(link string) {
                    worklist <- crawl(link)
                }(link)
            }
        }
    }
}

//tokens is a counting semaphore used to enforce a limit or 20 concurrent requests
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
    fmt.Println(url)
    tokens <- struct{}{} // aquire a token
    list, err := links.Extract(url)
    <- tokens // release the token
    if err != nil {
        log.Print(err)
    }
    return list
}
