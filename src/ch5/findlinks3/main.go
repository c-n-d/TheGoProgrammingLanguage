/*
findlinks3 crawls the web, starting with the URLs on the command line

$ go run src/ch5/findlinks3/main.go https://golang.org/
https://golang.org/
https://golang.org/doc/
https://golang.org/pkg/
https://golang.org/project/
https://golang.org/help/
https://golang.org/blog/
http://play.golang.org/
https://tour.golang.org/
https://golang.org/dl/
...
*/

package main

import (
    "fmt"
    "log"
    "os"

    "ch5/links"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
    seen := make(map[string]bool)
    for len(worklist) > 0 {
        items := worklist
        worklist = nil
        for _, item := range items {
            if !seen[item] {
                seen[item] = true
                worklist = append(worklist, f(item)...)
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

func main() {
    // Crawl the web breadth first
    // starting from the command line arguments
    breadthFirst(crawl, os.Args[1:])
}
