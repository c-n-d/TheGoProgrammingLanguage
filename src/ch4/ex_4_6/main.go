/*
Exercise 4.6 - In place fuunction to squash adjacent Unicode spaces in
a UTF-8 encoded []byte into a single ASCII space.

$ go run src/ch4/ex_4_6/main.go
squashSpaces("你好，   世界") = "你好， 世界"
squashSpaces("     Привет     мир     ") = " Привет мир "
*/

package main

import (
    "fmt"
    "unicode"
    "unicode/utf8"
)

func main() {
    inputs := [][]byte {
        []byte("你好，\u0020\u0020\u0020世界"),
        []byte("\u0020\u0020\u0020\u0020\u0020Привет\u0020\u0020\u0020\u0020\u0020мир\u0020\u0020\u0020\u0020\u0020")}

    for _, input := range inputs {
        tmp := make([]byte, len(input))
        copy(tmp, input)

        fmt.Printf("squashSpaces(%q) = %q\n", tmp, squashSpaces(input))
    }
}

/*
squashSpaces will remove adjacent, duplicate space characters.
Maintain two pointers; where we're reading from the string and where we're
copying runes back to the string. In the case of duplicate spaces, the reader is ahead of the writer.

Iterate over the runes in the string, always copying the current rune from the reader index. If the
reader is pointing to a unicode space ' ' rune, then we determine the index of the next non-space
rune to start reading from. Otherwise increament the reader and writer by the width of the last 
read rune. Writer is always updated by the width of the last written rune.
*/
func squashSpaces(s []byte) []byte {
    writePos := 0

    for readPos:= 0; readPos < len(s); {
        runeValue, width := utf8.DecodeRune(s[readPos:])

        copy(s[writePos:], string(runeValue))

        if unicode.IsSpace(runeValue) {
           readPos = nonSpaceIndex(s, readPos)
        } else {
            readPos += width
        }

        writePos += width
    }
    return s[:writePos]
}

// nonSpaceIndex will determine the index of the first unicode non-space rune by
// searching the byte array 's' starting from the 'start' index
func nonSpaceIndex(s []byte, start int) int {
    for i, w := start, 0; i < len(s); i+=w {
        runeValue, width := utf8.DecodeRune(s[i:])

        if !unicode.IsSpace(runeValue) {
            return i
        }

        w = width
    }
    return len(s)
}
