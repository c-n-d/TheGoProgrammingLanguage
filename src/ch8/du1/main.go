/*
du1 commputes the disk usage of files in a directory

$ go run src/ch8/du1/main.go ./src
1505 files 0.0 GB
$ go run src/ch8/du1/main.go ./data
8 files 0.0 GB
$ go run src/ch8/du1/main.go ./data ./src .
3530 files 0.1 GB
*/

package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
)

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
    // Print the results
    var nfiles, nbytes int64
    for size := range fileSizes {
        nfiles++
        nbytes += size
    }
    printDiskUsage(nfiles, nbytes)
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
        fmt.Fprintf(os.Stderr, "du1: %v\n", err)
        return nil
    }
    return entries
}
