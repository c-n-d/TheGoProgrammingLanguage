/*
Findlinks1 prints the links in an HTML document read from standard input

$ go build ch1/fetch
$ go build ch5/findlinks1
$ ./fetch https://golang.org | ./findlinks1
/lib/godoc/style.css
/opensearch.xml
/lib/godoc/jquery.treeview.css
/
/
#
/doc/
/pkg/
/project/
/help/
/blog/
http://play.golang.org/
#
#
//tour.golang.org/
https://golang.org/dl/
//blog.golang.org/
https://developers.google.com/site-policies#restrictions
/LICENSE
/doc/tos.html
http://www.google.com/intl/en/policies/privacy/
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
    if n.Type == html.ElementNode || n.Data == "a" {
        for _, a := range n.Attr {
            if a.Key == "href" {
                links = append(links, a.Val)
            }
        }
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        links = visit(links, c)
    }

    return links
}
