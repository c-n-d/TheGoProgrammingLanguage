// +build ignore

/*
The thumbnail command produces thumbnails of JPEG files whose names are provided
on each line of the standard input
*/

package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "sync"

    "ch8/thumbnail"
)
var globalWG sync.WaitGroup

func main() {
    var filenames []string
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        filenames = append(filenames, input.Text())
    }

    fnChan := make(chan string, len(filenames))
    totalsChan := make(chan int64)
    for _, file := range filenames {
        fnChan <- file
    }
    close(fnChan)

    globalWG.Add(1)
    go makeThumbnails6(fnChan, totalsChan)
    fmt.Println(<-totalsChan)
    globalWG.Wait()
}

func makeThumbnails6(filenames <-chan string, totals chan<- int64) {
    defer globalWG.Done()

    sizes := make(chan int64)
    var wg sync.WaitGroup
    for f := range filenames {
        wg.Add(1)
        // worker
        go func(f string) {
            defer wg.Done()
            thumb, err := thumbnail.ImageFile(f)
            if err != nil {
                log.Print(err)
                return
            }
            info, _ := os.Stat(thumb) // ok to ignore error
            sizes <- info.Size()
        }(f)
    }

    // closer
    go func() {
        wg.Wait()
        close(sizes)
    }()

    var total int64
    for size := range sizes {
        total += size
    }
    totals <- total
}
