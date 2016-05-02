/*
du3 commputes the disk usage of files in a directory

Traverses all directories in parallel, uses concurrency limiting semaphore
to avoid too many open files at once.

$ go run src/ch8/du3/main.go -v /etc /usr /bin ~
30102 files 11.9 GB
48055 files 12.7 GB
56373 files 13.4 GB
64841 files 13.5 GB
72952 files 14.3 GB
79385 files 14.5 GB
86140 files 15.0 GB
93787 files 15.4 GB
*/

package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "sync"
    "time"
)

var verbose = flag.Bool("v", false, "show verbose progress message")

func main() {
    flag.Parse()
    roots := flag.Args()
    if len(roots) == 0 {
        roots = []string{"."}
    }

    fileSizes := make (chan int64)
    var n sync.WaitGroup
    for _, root := range roots {
        n.Add(1)
        go walkDir(root, &n, fileSizes)
    }
    go func() {
        n.Wait()
        close(fileSizes)
    }()

    // Print the results periodically
    var tick <-chan time.Time
    if *verbose {
        tick = time.Tick(500 * time.Millisecond)
    }
    var nfiles, nbytes int64
    loop:
        for {
            select {
                case size, ok := <-fileSizes:
                    if !ok {
                        break loop // fileSizes was closed
                    }
                    nfiles++
                    nbytes += size
                case <-tick:
                    printDiskUsage(nfiles, nbytes)
            }
        }
    printDiskUsage(nfiles, nbytes) // final totals
}

func printDiskUsage(nfile, nbytes int64) {
    fmt.Printf("%d files %.1f GB\n", nfile, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
    defer n.Done()
    for _, entry := range dirents(dir) {
        if entry.IsDir() {
            n.Add(1)
            subdir := filepath.Join(dir, entry.Name())
            go walkDir(subdir, n, fileSizes)
        } else {
            fileSizes <- entry.Size()
        }
    }
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
    sema <- struct{}{} // aquire token
    defer func() { <- sema }() // release token

    entries, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Fprintf(os.Stderr, "du3: %v\n", err)
        return nil
    }
    return entries
}
