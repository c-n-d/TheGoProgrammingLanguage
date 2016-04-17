/*
Exercise 5.17 - Write a variadic function ElementsByTagName. For a HTML node tree, and zero or more
tags, returns all the elements that match one of those names.

$ go run src/ch5/ex_5_17/main.go https://golang.org
ElementsByTagName(doc, "link") matched [3] elements on [https://golang.org]
ElementsByTagName(doc, "script", "a", "div", "p") matched [64] elements on [https://golang.org]
*/

package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, names...string) []*html.Node {
    return findNodes(doc, &[]*html.Node{}, matchesTag(names...), nil)
}

func fetchDoc(url string) (*html.Node, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    doc, err := html.Parse(resp.Body)
    if err != nil {
        return nil, err
    }
    return doc, nil
}

func findNodes(n *html.Node, matches *[]*html.Node, pre, post func(n *html.Node) bool) []*html.Node {
    if pre != nil && pre(n){
        *matches = append(*matches, n)
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        findNodes(c, matches, pre, post)
    }

    if post != nil {
        post(n)
    }
    return *matches
}

func matchesTag(tags...string) func(n *html.Node) bool {
    return func(n *html.Node) bool {
        if n.Type == html.ElementNode {
            for _, tag := range tags {
                if tag == n.Data {
                    return true
                }
            }
        }
        return false
    }
}

func main() {
    for _, url := range os.Args[1:] {
        doc, err := fetchDoc(url)
        if err != nil {
            log.Fatal(err)
        }

        links := ElementsByTagName(doc, "link")
        fmt.Printf("ElementsByTagName(doc, \"link\") matched [%d] elements on [%s]\n", len(links), url)
        tags := ElementsByTagName(doc, "script", "a", "div", "p")
        fmt.Printf("ElementsByTagName(doc, \"script\", \"a\", \"div\", \"p\") matched [%d] elements on [%s]\n", len(tags), url)
    }
}
