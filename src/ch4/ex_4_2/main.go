/*
Exercise 4.2 prints the SHA256 hash of the value from stdin. Allows for -d flag to specify sha256/384/512 

$ go run src/ch4/ex_4_2/main.go -d 512
The quick brown fox jumps over the lazy dog
07e547d9586f6a73f73fbac0435ed76951218fb7d0c8d788a309d785436bbb642e93a252a954f23912547d1e8a3b5ed6e1bfd7097821233fa0538f3db854fee6

$ go run src/ch4/ex_4_2/main.go
The quick brown fox jumps over the lazy dog
d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592
*/

package main

import (
    "bufio"
    "crypto/sha256"
    "crypto/sha512"
    "flag"
    "fmt"
    "os"
)

var digest = flag.String("d", "256", "Whether to use sha256, sha384, or sha512")

func init() {
    flag.Parse()
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        input := scanner.Text()

        switch *digest {
            case "256":
                fmt.Printf("%x\n", sha256.Sum256([]byte(input)))
            case "384":
                fmt.Printf("%x\n", sha512.Sum384([]byte(input)))
            case "512":
                fmt.Printf("%x\n", sha512.Sum512([]byte(input)))
            default:
                fmt.Printf("using sha256: %x\n", sha256.Sum256([]byte(input)))
        }
    }
}
