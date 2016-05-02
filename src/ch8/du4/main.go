/*
du4 commputes the disk usage of files in a directory

Includes termination, cancelles when the user hits return

$ go run src/ch8/du4/main.go -v /etc /usr /bin ~
21584 files 9.0 GB
52902 files 13.2 GB
63311 files 13.7 GB
74751 files 14.6 GB

$
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

var done = make(chan struct{})

func cancelled() bool {
    select {
        case <-done:
            return true
        default:
            return false
    }
}

func main() {
    flag.Parse()
    roots := flag.Args()
    if len(roots) == 0 {
        roots = []string{"."}
    }

    go func() {
        os.Stdin.Read(make([]byte, 1)) // read a single byte
        close(done)
    }()

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
                case <-done:
                    // Drain fileSizes to allow existing goroutines to finish.
                    for range fileSizes {
                        // do nothing
                    }
                    return
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
    if cancelled() {
        return
    }
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
    select {
        case sema <- struct{}{}: // aquire token
        case <-done:
            return nil // cancelled
    }
    defer func() { <- sema }() // release token

    entries, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Fprintf(os.Stderr, "du4: %v\n", err)
        return nil
    }
    return entries
}
