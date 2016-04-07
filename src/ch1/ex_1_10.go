/*
Exercise 1.10 builds on fetchall to write the times and sizes to a file.

Observe the effect of caching on the load time of large sites

$ go run ex_1_10.go https://www.google.com http://stackoverflow.com
0.19s   19106 https://www.google.com
0.54s  245604 http://stackoverflow.com
0.54s elapsed

$ go run ex_1_10.go https://www.google.com http://stackoverflow.com
...

 cat ex_1_10_report.dat
0.19s   19106 https://www.google.com
0.54s  245604 http://stackoverflow.com
0.54s elapsed
0.13s  245552 http://stackoverflow.com
0.17s   19104 https://www.google.com
0.17s elapsed
0.13s  245552 http://stackoverflow.com
0.18s   19113 https://www.google.com
0.18s elapsed
0.17s   19127 https://www.google.com
0.20s  245552 http://stackoverflow.com
0.20s elapsed

*/
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
    "time"
)

var reportFileName = "ex_1_10_report.dat"

func main() {
    start := time.Now()
    ch := make(chan string)

    reportFile, err := os.OpenFile(reportFileName, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0644)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to open file [%s] with error [%v]. Not reporting.", reportFileName, err)
    }

    for _, url := range os.Args[1:] {
        // Starts a go routine
        go fetch(url, ch)
    }

    for range os.Args[1:] {
        // Receive from the channge ch
        recv := <- ch
        fmt.Println(recv)

        if reportFile != nil {
            reportFile.WriteString(recv + "\n")
        }
    }

    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

    if reportFile != nil {
            fmt.Fprintf(reportFile, "%.2fs elapsed\n", time.Since(start).Seconds())
            reportFile.Close()
    }
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    url = appendProtocol(url)

    resp, err := http.Get(url)

    if err != nil {
        // Send the error to the channel
        ch <- fmt.Sprint(err)
        return
    }

    nbytes, err := io.Copy(ioutil.Discard, resp.Body)

    // Don't leak resources
    resp.Body.Close()

    if err != nil {
        ch <- fmt.Sprintf("error while reading %s: %v", url, err);
        return
    }

    secs := time.Since(start).Seconds()

    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func appendProtocol(url string) string {
    if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
        return "http://" + url
    }

    return url
}