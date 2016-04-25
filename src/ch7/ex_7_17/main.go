/*
Exercise 7.17 - Extend xmlselect so that elements may also be selected by attribute.

$ go build -o ex_7_17 src/ch7/ex_7_17/main.go
$ go build -o fetch src/ch1/fetch/main.go
$ ./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | ./ex_7_17 html body div p a #rfc2119
*/

package main

import (
    "encoding/xml"
    "fmt"
    "io"
    "os"
    "strings"
)

func main() {
    dec := xml.NewDecoder(os.Stdin)
    var tokens []xml.StartElement
    for {
        tok, err := dec.Token()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
            os.Exit(1)
        }
        switch tok := tok.(type) {
            case xml.StartElement:
                tokens = append(tokens, tok) // push
            case xml.EndElement:
                tokens = tokens[:len(tokens)-1] // pop
            case xml.CharData:
                if containsAll(tokens, os.Args[1:]) {
                    fmt.Printf("%s: %s\n", strings.Join(stringTokens(tokens), " "), tok)
                }
        }
    }
}

func stringTokens(tokens []xml.StartElement) []string {
    var res []string
    for _, elem := range tokens {
        res = append(res, elem.Name.Local)
    }
    return res
}

// containsAll reports whether x contains the elements of y in order.
func containsAll(x []xml.StartElement, y []string) bool {
    for len(y) < len(x) {
        if len(y) == 0 {
            return true
        }
        if x[0].Name.Local == y[0] || containsAttr(x[0], y) {
            y = y[1:]
        }
        x = x[1:]
    }
    return false
}

func containsAttr(x xml.StartElement, y []string) bool {
    for _, attr := range x.Attr {
        for _, searchAttr := range y {
            if attr.Value == searchAttr {
                return true
            }
        }
    }
    return false
}
