/*
Exercise 5.8 - Extends outline2 so pre/post return a boolean result indicating
               whether to continue traversal. Use it to write a function ElementById
               that stops as soon as a match is found.

$ go run src/ch5/ex_5_8/main.go https://golang.org/ home
Unable to find node with id [home] at url [https://golang.org/]

$ go run src/ch5/ex_5_8/main.go https://golang.org/ about
<div id="about">
*/

package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "golang.org/x/net/html"
)

func main() {
    if len(os.Args) != 3 {
        log.Fatalf("usage: go run src/ch5/ex_5_8/main.go <url> <id>")
    }

    url := os.Args[1]
    id := os.Args[2]

    resp, err := http.Get(url)

    if err != nil {
        log.Fatalf("%v", err)
    }

    defer resp.Body.Close()

    doc, err := html.Parse(resp.Body)

    if err != nil {
        log.Fatalf("parser err -> %v", err)
    }

    n := ElementByID(doc, id)

    if n != nil {
        printNode(n, 0)
    } else {
        fmt.Printf("Unable to find node with id [%s] at url [%s]\n", id, url)
    }
}

func ElementByID(doc *html.Node, id string) *html.Node {
    return findFirstNode(doc, matchingId(id), nil)
}

/*
findFirstNode traverses the HTML DOM until a node satisfies the condition of pre() or post()
Terminates on finding the first satisfactory node
*/
func findFirstNode(n *html.Node, pre, post func(*html.Node) bool) *html.Node {
    if pre != nil && pre(n) {
        return n
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        match := findFirstNode(c, pre, post)

        if match != nil {
            return match
        }
    }

    if post != nil && post(n) {
        return n
    }

    return nil
}

// matchingId accepts a string 'id' and returns a function
// that determines if a html.Node contains the requested 'id'
func matchingId(id string) func (n *html.Node) bool {
    return func (n *html.Node) bool {
        for _, attr := range n.Attr {
            if attr.Key == "id" {
                return attr.Val == id
            }
        }
        return false
    }
}

func printNode(n *html.Node, depth int) {
        fmt.Printf("%*s<%s", depth, "", n.Data)
        for _, attr := range n.Attr {
            fmt.Printf(" %s=\"%s\"", attr.Key, attr.Val)
        }
        fmt.Printf(">\n")
}
