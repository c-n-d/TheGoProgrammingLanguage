/*
The popcount package defines various functions for determining the number of
set bits in a uint64 value - the population count.

The init function computes a lookup table for the number of 1 bits in the values from 0 - 255

1. PopCount uses each byte from the vaule to look up the corresponding number of set bits from pc
2. PopCount2 uses the same technique as PopCount, but uses a loop instead of a single expression
3. PopCount3 shifts the uint64 value 1-bit right on each iteration and adds the lowest bit value (0/1) to the sum
4. PopCount4 uses the fact that x & (x-1) clears the right most non-zero bit. For an interger with n 1-bits, after
   n iterations of x = x & (x-1) all 1-bits will have been cleared.
*/

package popcount

// pc[i] is the population count of i
var pc [256]byte

func init() {
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
