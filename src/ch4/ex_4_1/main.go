/*
Exercise 4.1 counts the number of bits that are different between
two SHA256 hashes

$ go run src/ch4/ex_4_1/main.go
2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
125
ffffffffffffffffffffffffffffffff00000000000000000000000000000000
00000000000000000000000000000000ffffffffffffffffffffffffffffffff
256
*/

package main

import (
    "crypto/sha256"
    "fmt"
)

func main() {
    c1 := sha256.Sum256([]byte("x"))
    c2 := sha256.Sum256([]byte("X"))

    c3 := [32]uint8{0 :0xFF, 1 :0xFF, 2  :0xFF, 3  :0xFF, 4  :0xFF, 5  :0xFF, 6  :0xFF, 7  :0xFF ,
                    8 :0xFF, 9 :0xFF, 10 :0xFF, 11 :0xFF, 12 :0xFF, 13 :0xFF, 14 :0xFF, 15 :0xFF}
    c4 := [32]uint8{16 :0xFF, 17 :0xFF, 18 :0xFF, 19 :0xFF, 20 :0xFF, 21 :0xFF, 22 :0xFF, 23 :0xFF ,
                    24 :0xFF, 25 :0xFF, 26 :0xFF, 27 :0xFF, 28 :0xFF, 29 :0xFF, 30 :0xFF, 31 :0xFF}

    fmt.Printf("%x\n%x\n%d\n", c1, c2, countBitDiff(c1, c2))
    fmt.Printf("%x\n%x\n%d\n", c3, c4, countBitDiff(c3, c4))
}

func countBitDiff(c1, c2 [32]uint8) int {
    var diff int

    for i := range c1 {
        // xor: 0^0=0, 0^1=1, 1^0=1, 1^1=0
        c := c1[i] ^ c2[i]

        // count the number of high bits in the xor result
        for ; c != 0; c = c >> 1 {
            diff += int(c & 1)
        }
    }

    return diff
}
