/*
Exercise 5.3 - Function to print all text nodes of an HTML document, skipping script and style elements.

$ go build ch1/fetch
$ go build ch5/ex_5_3
$ ./fetch https://golang.org | ./ex_5_3
The Go Programming Language
...
The Go Programming Language
Go
▽
Documents
Packages
The Project
Help
Blog
Play
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
Run
...
*/

package main

import (
    "fmt"
    "os"
    "strings"

    "golang.org/x/net/html"
)

func main() {
    doc, err := html.Parse(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
        os.Exit(1)
    }
    printText(doc)
}

func printText(n *html.Node) {
    if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
        return
    }

    if n.Type == html.TextNode {
        text := cleanText(n.Data)
        if text != "" {
            fmt.Printf("%s\n", text)
        }
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        printText(c)
    }
}

func cleanText(text string) string {
    return strings.TrimSpace(strings.Trim(text, "\n "))
}
