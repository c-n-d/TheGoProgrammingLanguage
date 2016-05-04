/*
Exercise 8.10 - Modify the web crawler in section 8.6 to support cancellation

$ go run src/ch8/ex_8_10/main.go https://golang.org
Press return to cancel http requests:
https://golang.org
https://golang.org/blog/
https://golang.org/pkg/
https://golang.org/
http://www.google.com/intl/en/policies/privacy/
...

2016/05/03 16:13:13 Get https://golang.org/doc/code.html: net/http: request canceled
https://golang.org/doc/codewalk/functions
2016/05/03 16:13:13 Get https://golang.org/ref/mem: net/http: request canceled
2016/05/03 16:13:13 Get https://code.google.com/p/go-tour/: net/http: request canceled
2016/05/03 16:13:13 Get https://golang.org/doc/codewalk/functions: net/http: request canceled
...
2016/05/03 16:13:13 Get https://golang.org/wiki/Learn: net/http: request canceled
*/

package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "ch5/links"
)

var cancel = make(chan struct{})

func cancelled() bool {
    select {
        case <-cancel:
            return true
        default:
            return false
    }
}

func init() {
    fmt.Println("Press return to cancel http requests:")
    time.Sleep(3 * time.Second)
}

func main() {
    worklist := make(chan []string)  // list of URLs, may have duplicates
    unseenLinks := make(chan string) // de-duplicate URLs

    go func() {
        os.Stdin.Read(make([]byte, 1)) // read a single byte
        close(cancel)
    }()

    // Starts with the command line argument
    go func() { worklist <- os.Args[1:] }()

    for i := 0; i < 20; i++ {
        go func() {
            for {
                select {
                    case <-cancel:
                        return
                    case link, ok := <-unseenLinks:
                        if !ok {
                            return
                        }
                        foundLinks := crawl(link)
                        go func() { worklist <- foundLinks }()
                }
            }
        }()
    }

   // The main goroutine de-duplicates worklist items
    seen := make(map[string]bool)
    for {
        select {
            case list, ok := <-worklist:
                if !ok {
                    close(unseenLinks)
                    return
                }
                for _, link := range list {
                    if !seen[link] && !cancelled() {
                        seen[link] = true
                        unseenLinks <- link
                    }
               }
            case <-cancel:
                return
        }
    }
}

func crawl(url string) []string {
    fmt.Println(url)
    list, err := links.ExtractCancel(url, cancel)
    if err != nil {
        log.Print(err)
    }
    return list
}
