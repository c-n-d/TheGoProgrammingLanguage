/*
du2 commputes the disk usage of files in a directory

Uses select and time.Tick to print totals periodically if -v is set

$ go run src/ch8/du2/main.go -v ~/Desktop/spark-1.6.0
7600 files 0.1 GB
15662 files 0.2 GB
23574 files 0.2 GB
30591 files 0.3 GB
40727 files 0.3 GB
44049 files 0.3 GB
*/

package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
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
    go func() {
        for _, root := range roots {
            walkDir(root, fileSizes)
        }
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
func walkDir(dir string, fileSizes chan<- int64) {
    for _, entry := range dirents(dir) {
        if entry.IsDir() {
            subdir := filepath.Join(dir, entry.Name())
            walkDir(subdir, fileSizes)
        } else {
            fileSizes <- entry.Size()
        }
    }
}

func dirents(dir string) []os.FileInfo {
    entries, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Fprintf(os.Stderr, "du2: %v\n", err)
        return nil
    }
    return entries
}
