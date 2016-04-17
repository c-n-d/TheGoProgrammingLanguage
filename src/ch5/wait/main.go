/*
wait program waits for a HTTP server to respond

$ go run src/ch5/wait/main.go https://google.com

$ go run src/ch5/wait/main.go
usage: go run src/ch5/wait/main.go <url>
exit status 1

$ go run src/ch5/wait/main.go https://google
2016/04/15 11:17:23 server not responding (Head https://google: dial tcp: lookup google: no such host); retrying...
2016/04/15 11:17:24 server not responding (Head https://google: dial tcp: lookup google: no such host); retrying...
2016/04/15 11:17:26 server not responding (Head https://google: dial tcp: lookup google: no such host); retrying...
2016/04/15 11:17:30 server not responding (Head https://google: dial tcp: lookup google: no such host); retrying...
2016/04/15 11:17:38 server not responding (Head https://google: dial tcp: lookup google: no such host); retrying...
2016/04/15 11:17:54 server not responding (Head https://google: dial tcp: lookup google: no such host); retrying...
2016/04/15 11:18:26 Site is down: server https://google failed to respond after 1m0s
exit status 1
*/

package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "usage: go run src/ch5/wait/main.go <url>\n")
        os.Exit(1)
    }
    url := os.Args[1]
    if err := WaitForServer(url); err != nil {
        log.Fatalf("Site is down: %v\n", err)
    }
}

// WaitForServer attempts to contact the server of a URL.
// Tries for 1 min using exponential backoff. Reports error if
// all attempts fail
func WaitForServer(url string) error {
    const timeout = 1 * time.Minute
    deadline := time.Now().Add(timeout)
    for tries:= 0; time.Now().Before(deadline); tries++ {
        _, err := http.Head(url)
        if err == nil {
            return nil
        }
        log.Printf("server not responding (%s); retrying...", err)
        time.Sleep(time.Second << uint(tries))
    }
    return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
