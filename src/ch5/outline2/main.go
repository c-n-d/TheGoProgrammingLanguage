/*
Outline2 prints the outline of a HTML tree

$ go run src/ch5/outline2/main.go https://google.com
<html>
  <head>
    <meta>
    </meta>
    <meta>
    </meta>
    <meta>
    </meta>
    <meta>
    </meta>
    <title>
    </title>
    <script>
    </script>
...
*/

package main

import (
    "fmt"
    "net/http"
    "os"

    "golang.org/x/net/html"
)

func main() {
    for _, url := range os.Args[1:] {
        outline(url)
    }
}

func outline(url string) error {
    resp, err := http.Get(url)

    if err != nil {
        return err
    }

    defer resp.Body.Close()

    doc, err := html.Parse(resp.Body)

    if err != nil {
        return err
    }

    forEachNode(doc, startElement, endElement)

    return nil
}

/*
forEachNode calls the function pre(x) and post(x) for each node x 
in the tree rooted at n. Both fn are optional. pre is called
before the children are visited (preorder) and post after the children
are visited (postorder)
*/
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
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

var depth int

func startElement(n *html.Node) {
    if n.Type == html.ElementNode {
        fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
        depth++
    }
}

func endElement(n *html.Node) {
    if n.Type == html.ElementNode {
        depth--
        fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
    }
}
