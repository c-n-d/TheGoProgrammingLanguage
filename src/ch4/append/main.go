/*
AppendInt is a function append an integer value to the end of a slice. Capactiy is ensured on 
each addition; doubling when the capacity of the underlying array is surpassed.

$ go run src/ch4/append/main.go
0  cap=1	[0]
1  cap=2	[0 1]
2  cap=4	[0 1 2]
3  cap=4	[0 1 2 3]
4  cap=8	[0 1 2 3 4]
5  cap=8	[0 1 2 3 4 5]
6  cap=8	[0 1 2 3 4 5 6]
7  cap=8	[0 1 2 3 4 5 6 7]
8  cap=16	[0 1 2 3 4 5 6 7 8]
9  cap=16	[0 1 2 3 4 5 6 7 8 9]
*/

package main

import "fmt"

func main() {
    var x, y []int

    for i := 0; i < 10; i++ {
        y = appendInt(x, i)
        fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
        x = y
    }
}

func appendInt(x []int, y...int) []int {
    var z []int
    zlen := len(x) + len(y)

    if zlen <= cap(x) {
        // There is room to grow. Extend the slice
        z = x[:zlen]
    } else {
        // There is insufficent room to grow. Allocate a new array.
        // Grow by doubling, for amortized linear complexity.
        zcap := zlen
        if zcap < 2*len(x) {
            zcap = 2*len(x)
        }
        z = make([]int, zlen, zcap)
        copy(z, x)
    }

    copy(z[len(x):], y)
    return z
}
