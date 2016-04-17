/*
Title2 prints the title of an html document specified by a url. Defers closing the response body until the end of title

$ go run src/ch5/title2/main.go https://golang.org https://golang.org/doc/effective_go.html https://golang.org/doc/gopher/frontpage.png
The Go Programming Language
Effective Go - The Go Programming Language
https://golang.org/doc/gopher/frontpage.png has type image/png, not text/html
*/

package main

import (
    "fmt"
    "net/http"
    "os"
    "strings"

    "golang.org/x/net/html"
)

func title(url string) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }

    defer resp.Body.Close()

    // Check Content-Type is HTML (e.g. "text/html; charset=utf-8")
    ct := resp.Header.Get("Content-Type")
    if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
        return fmt.Errorf("%s has type %s, not text/html", url, ct)
    }
    doc, err := html.Parse(resp.Body)

    if err != nil {
        return fmt.Errorf("parsing %s as HTML: %v", url, err)
    }
    visitNode := func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "title" &&
           n.FirstChild != nil {
               fmt.Println(n.FirstChild.Data)
           }
    }
    forEachNode(doc, visitNode, nil)
    return nil
}

func forEachNode(n *html.Node, pre, post func(*html.Node)) {
    if pre != nil {
        pre(n)
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(c, pre, post)
    }

    if post != nil {
        post(n)
    }
}

func main() {
    for _, url := range os.Args[1:] {
        if err := title(url); err != nil {
            fmt.Println(err)
        }
    }
}
