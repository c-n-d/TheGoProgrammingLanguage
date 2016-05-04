/*
Exercise 8.9 - Write a version of du that computes and periodically displays separate totals for each root

$ go run src/ch8/ex_8_9/main.go -v /etc /usr /bin
Root [/etc]: 182 files 0.0 GB
Root [/usr]: 13663 files 1.0 GB
Root [/bin]: 37 files 0.0 GB

Root [/etc]: 259 files 0.0 GB
Root [/usr]: 28033 files 1.3 GB
Root [/bin]: 37 files 0.0 GB

Root [/etc]: 259 files 0.0 GB
Root [/usr]: 28033 files 1.3 GB
Root [/bin]: 37 files 0.0 GB

Root [/etc]: 259 files 0.0 GB
Root [/usr]: 41144 files 1.4 GB
Root [/bin]: 37 files 0.0 GB

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

type DirSize struct {
    Root           string
    Nfiles, Nbytes int64
}

func (ds *DirSize) Add(other DirSize) {
      ds.Nfiles += other.Nfiles
      ds.Nbytes += other.Nbytes
}

func (ds *DirSize) String() string {
    return fmt.Sprintf("Root [%s]: %d files %.1f GB\n", ds.Root, ds.Nfiles, float64(ds.Nbytes)/1e9)
}

// Maps root to it's accumulated directory size
var accumulated = make(map[string]*DirSize)

var verbose = flag.Bool("v", false, "show verbose progress message")

func main() {
    flag.Parse()
    roots := flag.Args()
    if len(roots) == 0 {
        roots = []string{"."}
    }

    fileSizes := make (chan DirSize)
    var n sync.WaitGroup
    for _, root := range roots {
        n.Add(1)
        go walkDir(root, root, &n, fileSizes)
        accumulated[root] = &DirSize{root, 0, 0}
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

    loop:
        for {
            select {
                case ds, ok := <-fileSizes:
                    if !ok {
                        break loop // fileSizes was closed
                    }
                    accumulated[ds.Root].Add(ds)
                case <-tick:
                    printDiskUsage()
            }
        }
    printDiskUsage() // final totals
}

func printDiskUsage() {
    for _, ds := range accumulated {
        fmt.Printf(ds.String())
    }
    fmt.Println()
}

// walkDir recursively walks the file tree rooted at dir
func walkDir(root ,dir string, n *sync.WaitGroup, fileSizes chan<- DirSize) {
    defer n.Done()
    for _, entry := range dirents(dir) {
        if entry.IsDir() {
            n.Add(1)
            subdir := filepath.Join(dir, entry.Name())
            go walkDir(root, subdir, n, fileSizes)
        } else {
            fileSizes <- DirSize{root, 1, entry.Size()}
        }
    }
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
    sema <- struct{}{} // aquire token
    defer func() { <- sema }() // release token

    entries, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Fprintf(os.Stderr, "ex8.9: %v\n", err)
        return nil
    }
    return entries
}
