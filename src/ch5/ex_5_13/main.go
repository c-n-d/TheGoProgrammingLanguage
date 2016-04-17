/*
Exercise 5.13 - Modify crawl to make local copies of the pages it finds,
                creating directories. Don't make copies of pages from different
                domains.

$ go run src/ch5/ex_5_13/main.go https://golang.org/
https://golang.org/
Writing html from [https://golang.org/] to tmp/golang.org/index.html
https://golang.org/doc/
Writing html from [https://golang.org/doc/] to tmp/golang.org/doc/index.html
https://golang.org/pkg/
Writing html from [https://golang.org/pkg/] to tmp/golang.org/pkg/index.html
https://golang.org/project/
Writing html from [https://golang.org/project/] to tmp/golang.org/project/index.html
https://golang.org/help/
Writing html from [https://golang.org/help/] to tmp/golang.org/help/index.html
https://golang.org/blog/
Writing html from [https://golang.org/blog/] to tmp/golang.org/blog/index.html
http://play.golang.org/
Writing html from [http://play.golang.org/] to tmp/play.golang.org/index.html
https://tour.golang.org/
https://golang.org/dl/
Writing html from [https://golang.org/dl/] to tmp/golang.org/dl/index.html
https://blog.golang.org/
https://developers.google.com/site-policies#restrictions
https://golang.org/LICENSE

$ ls tmp/
golang.org    index.html    play.golang.org
$ ls tmp/golang.org/
LICENSE        blog        cmd        dl        doc        help        index.html    pkg        project        ref        wiki
$ ls tmp/golang.org/blog/
index.html
*/

package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/url"
    "os"
    "strings"

    "ch5/links"
)

const (
    OUT_DIR = "tmp"
)

var domainWhitelist = map[string]bool{
    "golang.org": true,
    "play.golang.org": true,
}

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
    seen := make(map[string]bool)
    for len(worklist) > 0 {
        items := worklist
        worklist = nil
        for _, item := range items {
            if !seen[item] {
                seen[item] = true
                worklist = append(worklist, f(item)...)
            }
        }
    }
}

func crawl(url string) []string {
    fmt.Println(url)
    err := download(url)
    if err != nil {
        fmt.Printf("skipping %s: %v\n", url, err)
    }
    list, err := links.Extract(url)
    if err != nil {
        log.Print(err)
    }
    return list
}

func download(urlStr string) error {
    u, err := url.Parse(urlStr)
    if err != nil {
        return fmt.Errorf("invalid url %s, %v\n", urlStr, err)
    }

    if domainWhitelist[u.Host] {
        destDir := dirFromURL(u, OUT_DIR)
        fileName := fileNameFromURL(u)

        err := createDir(destDir)
        if err != nil {
            return err
        }

        htmlData, err := links.DownloadHTML(urlStr)
        if err != nil {
            return err
        }
        fmt.Printf("Writing html from [%s] to %s\n", urlStr, destDir + "/" + fileName)

        err = ioutil.WriteFile(destDir + "/" + fileName, []byte(htmlData), os.FileMode(int(0777)))
        if err != nil {
            fmt.Printf("error writefile: %v\n", err)
            return err
        }
    }
    return nil
}

func dirFromURL(u *url.URL, base string) string {
    dest := base + "/" + u.Host + u.Path
    if len(dest) == 0 {
        return base
    }
    lastSlash := strings.LastIndex(dest, "/")
    if lastSlash != -1 {
        dest = dest[:lastSlash]
    }
    return dest
}

func fileNameFromURL(u *url.URL) string {
    path := u.Path
    lastSlash := strings.LastIndex(path, "/")
    if lastSlash != -1 {
        path = path[lastSlash:]
    }
    if len(path) == 0 || path == "/" {
        return "index.html"
    }
    return path
}

func createDir(dir string) error {
    if err := os.Mkdir(dir, os.ModeDir | os.FileMode(int(0777))); err != nil {
        if !os.IsExist(err) {
            return err
        }
    }
    return nil
}

func main() {
    // Crawl the web breadth first
    // starting from the command line arguments
    breadthFirst(crawl, os.Args[1:])
}
