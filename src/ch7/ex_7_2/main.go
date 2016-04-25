/*
Exercise 7.2 - Write a function CountingWriter; given an io.Writer return a new Writer that
               wraps the original and a pointer to an int64 variable that at any moment contains
               the number of bytes written to the new writer.

$ go run src/ch7/ex_7_2/main.go 
Wrapping os.Stdout as CountingWriter(os.Stdout)

Hello
Wrote [6] bytes so far
, there!
Wrote [15] bytes in total
*/

package main

import (
    "fmt"
    "io"
    "os"
)

type WriterWrapper struct {
    underlyingWriter *io.Writer
    count *int64
}

// Writes to WriterWrapper call the underlying Writer's Write method, update the count written by the
// underlying writer and propegate the results to the caller
func (wrapper WriterWrapper) Write(p []byte) (int, error) {
    uw := *(wrapper.underlyingWriter)
    n, err := uw.Write(p)
    *wrapper.count += int64(n)
    return n, err
}

func main() {
    // Wrapping os.Stdout in a counting WriterWrapper. As show below, writing to
    // os.Stdout directly (eg fmt.Println) won't change the # of bytes written.
    writer, bytesWritten := CountingWriter(os.Stdout)
    fmt.Println("Wrapping os.Stdout as CountingWriter(os.Stdout)\n")

    writer.Write([]byte("Hello\n"))
    fmt.Printf("Wrote [%d] bytes so far\n", *bytesWritten)

    fmt.Fprintf(writer, ", there!\n")
    fmt.Printf("Wrote [%d] bytes in total\n", *bytesWritten)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
    var count int64
    cw := WriterWrapper{&w, &count}
    return cw, &count
}
