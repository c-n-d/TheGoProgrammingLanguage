
/*
Exercise 8.6 -

$ go run src/ch8/ex_8_6/main.go https://golang.org

*/

package main

import (
    "flag"
    "fmt"
    "log"
    "sync"

    "ch5/links"
)

var depth = flag.Int("depth", 2, "URLs reachable from `depth` links will be fetched")

type LinkDepth struct {
    Link  string
    Depth int
}

type LinkListDepth struct {
    Links []string
    Depth int
}

func init() {
    flag.Parse()
}

var wg sync.WaitGroup

func main() {
    worklist    := make(chan LinkListDepth)  // list of URLs, may have duplicates
    unseenLinks := make(chan LinkDepth) // de-duplicate URLs

     wg.Add(len(flag.Args()))

    // Starts with the command line argument
    go func() { worklist <- LinkListDepth{flag.Args(), 0} }()

    for i := 0; i < 20; i++ {
        go func() {
            for ld := range unseenLinks {
                cd := ld.Depth
                foundLinks := crawl(ld.Link, cd)
                go func() {
                    if cd <= *depth {
                        worklist <- LinkListDepth{ foundLinks, cd + 1 }
                    }
                }()
            }
        }()
    }

    go func() {
        wg.Wait()
        fmt.Println("Done ")
        //close(unseenLinks)
    }()

   // The main goroutine de-duplicates worklist items
    seen := make(map[string]bool)
    for lld := range worklist {
        for _, link := range lld.Links {
            if !seen[link] {
                seen[link] = true
                unseenLinks <- LinkDepth{ link, lld.Depth }
            }
        }
    }
}

func crawl(url string, depth int) []string {
    fmt.Println(url, depth)
    list, err := links.Extract(url)
    if err != nil {
        log.Print(err)
    }
    return list
}
