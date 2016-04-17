/*
Exercise 5.1 - Recursively traverse the n.FirstChild linked list using calls to vist.

html
  |
head - body
  |      |
  |      p - p - p - div
  |
meta - script

$ go build ch5/findlinks1
$ go build ch1/fetch
$ go build ch5/ex_5_1
$ ./fetch https://golang.org | ./ex_5_1
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

$ ./fetch https://golang.org | ./ex_5_1 | wc -l
      21
$ ./fetch https://golang.org | ./findlinks1 | wc -l
      21
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

    if n.FirstChild != nil {
        links = visit(links, n.FirstChild)
    }

    if n.NextSibling != nil {
        links = visit(links, n.NextSibling)
    }

    return links
}
