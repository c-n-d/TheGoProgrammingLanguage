/*
Exercise 5.4 - Extend the visit function so that it extracts ofther kinds of links from
               the document, such as images, scripts, and style sheets

Below is a random sample of .js, .css, and .png files produced by this exercise

$ go build ch1/fetch
$ go build ch5/ex_5_4
$ ./fetch https://golang.org/doc | ./ex_5_4
/lib/godoc/style.css
/opensearch.xml
/
#
/doc/
/pkg/
/project/
/blog/
http://play.golang.org/
/doc/install
/doc/gopher/doc.png
//tour.golang.org/
code.html
//www.youtube.com/watch?v=XCsL89YtqCs
/cmd/go/
//blog.golang.org/
/doc/codewalk/functions
/pkg/encoding/gob/
/blog/laws-of-reflection
/doc/articles/go_command.html
https://ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js
/lib/godoc/playground.js
/lib/godoc/godocs.js
...
*/

package main

import (
    "fmt"
    "os"

    "golang.org/x/net/html"
)

func main() {
    doc, err := html.Parse(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
        os.Exit(1)
    }

    for _, link := range visit(nil, doc) {
        fmt.Println(link)
    }
}

func visit(links []string, n *html.Node) []string {
    if n.Type == html.ElementNode {
        for _, attr := range n.Attr {
            // Check href (a, link) and src (script, img) attributes for links
            if attr.Key == "href" || attr.Key == "src"{
                links = append(links, attr.Val)
            }
        }
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        links = visit(links, c)
    }

    return links
}
