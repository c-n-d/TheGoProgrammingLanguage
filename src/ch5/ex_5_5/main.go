/*
Exercise 5.5 - Implement countWordsAndImages (using bare returns)

$ go build ch5/ex_5_5
$ ./ex_5_5 https://www.google.com http://stackoverflow.com
https://www.google.com: Words=243, Images=2
http://stackoverflow.com: Words=3365, Images=9
*/

package main

import (
    "bufio"
    "fmt"
    "net/http"
    "os"
    "strings"

    "golang.org/x/net/html"
)

func main() {
    for _, url := range os.Args[1:] {
        words, images, err := CountWordsAndImages(url)

        if err != nil {
            fmt.Fprintf(os.Stderr, "ex 5.5: %v\n", err)
            continue
        }
        fmt.Printf("%s: Words=%d, Images=%d\n", url, words, images)
    }
}

func CountWordsAndImages(url string) (words, images int, err error) {
    resp, err := http.Get(url)
    if err != nil {
        return
    }
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        err = fmt.Errorf("parsing HTML: %s", err)
        return
    }
    words, images = countWordsAndImages(doc)
    return
}

func countWordsAndImages(n *html.Node) (words, images int) {
    // If the current node is an image element, increament the imgage count
    if n.Type == html.ElementNode && n.Data == "img" {
        images++
    }

    // If the current node is a TextNode, split the Data string and count the words
    if n.Type == html.TextNode {
        scanner := bufio.NewScanner(strings.NewReader(n.Data))
        scanner.Split(bufio.ScanWords)
        for scanner.Scan() {
            scanner.Text()
            words++
        }
    }

    // Aggregate the count of words and images from the current node with the count from all of it's children
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        w, i := countWordsAndImages(c)
        words+=w
        images+=i
    }

    // Return the count of words and images
    return
}
