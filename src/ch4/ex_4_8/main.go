/*
Charcount computes counts of unicode characters

$ go run src/ch4/ex_4_8/main.go
a b c £ go ¶
á ń
Ǭ ỳ
�
rune	count
'¶'	1
'á'	1
' '	7
'\n'	4
'c'	1
'£'	1
'g'	1
'ỳ'	1
'�'	1
'b'	1
'o'	1
'ń'	1
'Ǭ'	1
'a'	1

type	count
SYMBOL	2
CONTROL	4
GRAPHIC	19
LETTER	9
MARK	9
NUMBER	9
SPACE	11

len	count
1	16
2	5
3	2
4	0
*/

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "unicode"
    "unicode/utf8"
)

const (
    CONTROL = "CONTROL"
    DIGIT = "DIGIT"
    GRAPHIC = "GRAPHIC"
    LETTER = "LETTER"
    MARK = "MARK"
    NUMBER = "NUMBER"
    SPACE = "SPACE"
    SYMBOL = "SYMBOL"
)

func main() {
    counts := make(map[rune]int)    // counts of unicode characters
    var utflen [utf8.UTFMax + 1]int // count of lenghts of UTF-8 encodings
    invalid := 0                    // count of invalid UTF-8 characters

    typeCounts := make(map[string]int)    // counts of unicode characters by type

    in := bufio.NewReader(os.Stdin)
    for {
        r, n, err := in.ReadRune()  // returns rune, nbytes, error

        if err == io.EOF {
            break
        }
        if err !=nil {
            fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
            os.Exit(1)
        }
        if r == unicode.ReplacementChar && n == 1 {
            invalid++
            continue
        }

        collectTypeMetric(typeCounts, r)

        counts[r]++
        utflen[n]++
    }

    fmt.Printf("rune\tcount\n")
    for c, n := range counts {
        fmt.Printf("%q\t%d\n", c, n)
    }

    fmt.Printf("\ntype\tcount\n")
    for c, n := range typeCounts {
        fmt.Printf("%s\t%d\n", c, n)
    }

    fmt.Printf("\nlen\tcount\n")
    for i, n := range utflen {
        if i > 0 {
            fmt.Printf("%d\t%d\n", i, n)
        }
    }

    if invalid > 0 {
        fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
    }
}

func collectTypeMetric(typeCounts map[string]int, r rune ) {
    if unicode.IsControl(r) {
        typeCounts[CONTROL]++
    }
    if unicode.IsDigit(r) {
        typeCounts[DIGIT]++
    }
    if unicode.IsGraphic(r) {
        typeCounts[GRAPHIC]++
    }
    if unicode.IsLetter(r) {
        typeCounts[LETTER]++
    }
    if unicode.IsLetter(r) {
        typeCounts[MARK]++
    }
    if unicode.IsLetter(r) {
        typeCounts[NUMBER]++
    }
    if unicode.IsSpace(r) {
        typeCounts[SPACE]++
    }
    if unicode.IsSymbol(r) {
        typeCounts[SYMBOL]++
    }
}
