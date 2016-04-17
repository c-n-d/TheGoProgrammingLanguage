/*
Findlinks2 does a HTTP GET on each provided URL, parses the resulting HTML
and prints the links.

$ go build ch5/findlinks2
$ ./findlinks2 https://www.google.com http://stackoverflow.com/
/images/branding/product/ico/googleg_lodp.ico
https://www.google.com/imghp?hl=en&tab=wi
https://maps.google.com/maps?hl=en&tab=wl
https://play.google.com/?hl=en&tab=w8
https://www.youtube.com/?tab=w1
...
/questions/tagged/sitecore
/questions/tagged/web-forms-for-marketers
/questions/36611395/sitecore-wffm-sendemailmessage-smtp-password-authentication-error
/users/3443730/adam
/questions/36611394/how-can-i-display-a-variable-in-a-dynamic-text-field
/questions/tagged/actionscript-3
/questions/tagged/variables
/questions/36611394/how-can-i-display-a-variable-in-a-dynamic-text-field
...

$ ./findlinks2 https://www.google.com http://stackoverflow
/images/branding/product/ico/googleg_lodp.ico
https://www.google.com/imghp?hl=en&tab=wi
https://maps.google.com/maps?hl=en&tab=wl
https://play.google.com/?hl=en&tab=w8
https://www.youtube.com/?tab=w1
...
findlinks2: Get http://stackoverflow: dial tcp: lookup stackoverflow: no such host
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
        links, err := findLinks(url)

        if err != nil {
            fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
            continue
        }

        for _, link := range links {
            fmt.Println(link)
        }
    }
}

func findLinks(url string) ([]string, error) {
    resp, err := http.Get(url)

    if err != nil {
        return nil, err
    }

    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
    }

    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
        return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
    }

    return visit(nil, doc), nil
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
