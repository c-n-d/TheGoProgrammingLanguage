/*
Exercise 5.2 - Populate a mapping from element names - p, div, span, and so on - to
               the number of elements with that name in the HTML document.

$ go build ch1/fetch
$ go build ch5/ex_5_2
$ ./fetch https://golang.org | ./ex_5_2
Element    | # Occur
-----------|-----------
meta       | 3
body       | 1
span       | 3
input      | 1
select     | 1
option     | 8
head       | 1
link       | 3
textarea   | 2
iframe     | 1
title      | 1
div        | 33
br         | 3
html       | 1
script     | 9
a          | 22
form       | 1
pre        | 1
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

    ec := make(map[string]int)

   fmt.Printf("%-10s | %s\n", "Element", "# Occur")
   fmt.Println("-----------|-----------")

    for element, count := range countElements(ec, doc) {
        fmt.Printf("%-10s | %d\n", element, count)
    }
}

func countElements(ec map[string]int, n *html.Node) map[string]int {
    if n.Type == html.ElementNode {
        ec[n.Data]++
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        countElements(ec, c)
    }

    return ec
}
