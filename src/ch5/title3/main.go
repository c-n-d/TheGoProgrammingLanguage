/*
Title3 prints the title element of an HTML document specified by a URL. Panics if multiple title occur

$ go run src/ch5/title3/main.go https://golang.org http://localhost:8888/
The Go Programming Language
multiple title elements

$ curl http://localhost:8888/
<html>
  <head>
    <title>Title 1</title>
    <title>Title 2</title>
  </head>
  <body>
  </body>
</html>
*/

package main

import (
    "fmt"
    "net/http"
    "os"
    "strings"

    "golang.org/x/net/html"
)

func soleTitle(doc *html.Node) (title string, err error) {
    type bailout struct{}

    defer func() {
        switch p := recover(); p {
            case nil:
                // no panic
            case bailout{}:
                // "expected" panic
                err = fmt.Errorf("multiple title elements")
            default:
                panic(p) // unexpected panic; carry on panicing
        }
    }()

    // Bail out of recursion if we find more than one non-empty title
    forEachNode(doc, func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "title" &&
            n.FirstChild != nil {
                if title != "" {
                    panic(bailout{}) // multiple title elements
                }
                title = n.FirstChild.Data
            }
    }, nil)
    if title == "" {
        return "", fmt.Errorf("no title element")
    }
    return title, nil
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
    title, err := soleTitle(doc)
    if err != nil {
        return err
    }
    fmt.Println(title)
    return nil
}

func main() {
    for _, url := range os.Args[1:] {
        if err := title(url); err != nil {
            fmt.Println(err)
        }
    }
}
