/*
Exercise 9.2 - Rewrite popcount (2.6.2) so that initialization of the lookup table is
               done once the first time it's needed.

$ go test -v -bench=. ch9/ex_9_2
*/

package popcount

import "sync"

// pc[i] is the population count of i
var pc [256]byte

var initOnce sync.Once

func initPC() {
    for i := range pc {
        /*
         * byte(i & 1) produces a {0, 1, 0, 1, 0 ...} sequence
         *
         *   i: 0  1  2  3 ... 100  101 ... 254  255
         * i/2: 0  0  1  1      50   50     127  127
         * i&1: 0  1  0  1       0    1       0    1
         */
        pc[i] = pc[i/2] + byte(i & 1)
    }
}

func PopCount(x uint64) int {
    // byte      7     6     5     4     3     2     1   0
    // uint64 |63-56|55-48|47-40|39-32|31-24|23-16|15-8|7-0|

    initOnce.Do(initPC)

    return int(pc[byte(x >> (0*8))] +
        pc[byte(x >> (1*8))] +
        pc[byte(x >> (2*8))] +
        pc[byte(x >> (3*8))] +
        pc[byte(x >> (4*8))] +
        pc[byte(x >> (5*8))] +
        pc[byte(x >> (6*8))] +
        pc[byte(x >> (7*8))])
}

// Exercise 2.3
func PopCount2(x uint64) int {
    var sum int

    for i := 0; i < 8; i++ {
        sum += int(pc[byte(x >> (uint(i*8)))])
    }

    return sum;
}

// Exercise 2.4
func PopCount3(x uint64) int {
    var sum int

    for ; x != 0; x = x >> 1 {
        sum += int(x & 1)
    }

    return sum;
}

// Exercise 2.5
func PopCount4(x uint64) int {
    var sum int

    for ;x != 0; x &= (x-1) {
        sum++
    }

    return sum
}
