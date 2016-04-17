/*
Exercise 5.7 - Extends outline2 to pretty print the HTML

$ go run src/ch5/ex_5_7/main.go https://golang.org/ > src/ch5/ex_5_7/golang.html
$ open src/ch5/ex_5_7/golang.html

$ cd src/ch5/ex_5_7/
$ go test -v -bench=.
=== RUN   TestOutlinePrettyPrint
--- PASS: TestOutlinePrettyPrint (0.14s)
PASS
ok  	ch5/ex_5_7	0.150s
*/

package main

import (
    "fmt"
    "io"
    "net/http"
    "os"

    "golang.org/x/net/html"
)

func main() {
    for _, url := range os.Args[1:] {
        outline(url, os.Stdout)
    }
}

func outline(url string, w io.Writer) error {
    resp, err := http.Get(url)

    if err != nil {
        return err
    }

    defer resp.Body.Close()

    doc, err := html.Parse(resp.Body)

    if err != nil {
        return err
    }

    forEachNode(doc, w, startElement, endElement)

    return nil
}

/*
forEachNode calls the function pre(x) and post(x) for each node x 
in the tree rooted at n. Both fn are optional. pre is called
before the children are visited (preorder) and post after the children
are visited (postorder)
*/
func forEachNode(n *html.Node, w io.Writer, pre, post func(*html.Node, io.Writer)) {
    if pre != nil {
        pre(n, w)
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(c, w, pre, post)
    }

    if post != nil {
        post(n, w)
    }
}

var depth int

// StartElement prints comments in <!-- --> form, text blocks
// and element nodes with their attributes.
func startElement(n *html.Node, w io.Writer) {
    if n.Type == html.CommentNode {
        fmt.Fprintf(w, "%*s<!--%s-->\n", depth, "", n.Data)
    }

    if n.Type == html.TextNode {
        fmt.Fprintf(w, "%*s%s\n", depth, "", n.Data)
    }

    if n.Type == html.ElementNode {
        fmt.Fprintf(w, "%*s<%s", depth, "", n.Data)
        for _, attr := range n.Attr {
            fmt.Fprintf(w, " %s=\"%s\"", attr.Key, attr.Val)
        }

        if hasClosingElement(n) {
            fmt.Fprintf(w, ">\n")
            depth++
        } else {
            fmt.Fprintf(w, "/>\n")
        }
    }
}

// Only prints closing tag for elements with children
func endElement(n *html.Node, w io.Writer) {
    if n.Type == html.ElementNode  && hasClosingElement(n) {
        depth--
        fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
    }
}

// If element has no children, use short version (i.e. <img/> vs <img></img>)
func hasClosingElement(n *html.Node) bool {
    return n.FirstChild != nil
}
