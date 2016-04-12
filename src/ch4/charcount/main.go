/*
Charcount computes counts of unicode characters

$ go run src/ch4/charcount/main.go
a b c £ go ¶
á ń
Ǭ ỳ
�
rune	count
'�'	1
'b'	1
'c'	1
'g'	1
'o'	1
'á'	1
'ń'	1
'ỳ'	1
'a'	1
' '	7
'£'	1
'¶'	1
'\n'	4
'Ǭ'	1

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

func main() {
    counts := make(map[rune]int)    // counts of unicode characters
    var utflen [utf8.UTFMax + 1]int // count of lenghts of UTF-8 encodings
    invalid := 0                    // count of invalid UTF-8 characters

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

        counts[r]++
        utflen[n]++
    }

    fmt.Printf("rune\tcount\n")
    for c, n := range counts {
        fmt.Printf("%q\t%d\n", c, n)
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
