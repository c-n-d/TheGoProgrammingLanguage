/*
Exercise 4.7 - Modify reverse to reverse the characters of a []byte slice
that represents a UTF-8 encoded string, in place. w/o new memory allocation?

1. reverse - Allocate new output slice. Read the original slice in reverse, copying to output
2. reverse2 - Moves the last charcter of the slice to the beginning, by shifting the remaining slice
              by the width of the last rune. Each iteration works on one fewer runes; leaving the first
              rune in place.

$ go run src/ch4/ex_4_7/main.go
reverse("你好，世界") = "界世，好你"
reverse2("你好，世界") = "界世，好你"

reverse("Привет мир") = "рим тевирП"
reverse2("Привет мир") = "рим тевирП"
*/

package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    inputs := [][]byte {
        []byte("你好，世界"),
        []byte("Привет мир")}

    for _, input := range inputs {
        tmp := make([]byte, len(input))
        copy(tmp, input)

        fmt.Printf("reverse(%q) = %q\n", tmp, reverse(input))
        fmt.Printf("reverse2(%q) = %q\n\n", tmp, reverse2(input))
    }
}

// The 'reverse' function reverses the runes in a byte array. A new byte array is allocated for the output.
func reverse(s []byte) []byte {
    output   := make([]byte, len(s))
    writePos := 0

    for len(s) > 0 {
        runeValue, width := utf8.DecodeLastRune(s)

        copy(output[writePos:], string(runeValue))

        writePos += width

        s = s[:len(s) - width]
    }

    return output
}
/*
reverse2 will reverse the string w/o allocating additional memory.

1. Read the last rune in the string
2. Shift the slice to the right by the width of the rune
3. After shifting, append the rune to the beginning of the slice
4. Update the slice to not include the first character

orig: "你好，世界"

i | Whole Str. | Shifted view
--|------------|-------------
0 | "界你好，世" | "界你好，世"
1 | "界世你好，" | "世你好，"
2 | "界世，你好" | "，你好"
3 | "界世，好你" | "好你"
4 | "界世，好你" | "你"
*/
func reverse2(s []byte) [] byte{
    ret := s
    n := utf8.RuneCount(s)

    for i := 0; i < n; i++ {
        runeValue, width := utf8.DecodeLastRune(s)
        shift(s, width)
        copy(s, string(runeValue))
        s = s[width:]
    }
    return ret
}

/*
shift will move the bytes 'width' positions to the right

ex.
original = [0, 1, 2, 3, 4, 5]

shift(original, 3)

 i | j | s[i] | s[j]
---|---|------|-----
 5 | 2 |  5   |  2
 4 | 1 |  4   |  1
 3 | 0 |  3   |  0

 => [0 1 2 0 1 2]
*/
func shift(s []byte, width int) {
    fmt.Printf(" i | j | s[i] | s[j]\n")
    fmt.Printf("---|---|------|-----\n")
    for i, j := len(s) - 1, len(s) - width - 1; j >= 0; i, j = i - 1, j - 1 {
        fmt.Printf(" %d | %d | %d | %d \n", i, j, s[i], s[j])
        s[i] = s[j]
    }
}
