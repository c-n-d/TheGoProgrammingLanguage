/*
$ cd src/ch2/popcount

$ go test -v -bench=.
=== RUN   TestPopCountVersions
PopCount(1234567898765431) = 26, PopCount2(1234567898765431) = 26, PopCount3(1234567898765431) = 26, PopCount4(1234567898765431) = 26
PopCount(664652156845612141) = 36, PopCount2(664652156845612141) = 36, PopCount3(664652156845612141) = 36, PopCount4(664652156845612141) = 36
--- PASS: TestPopCountVersions (0.00s)
PASS
BenchmarkPopCount-4     200000000             6.18 ns/op
BenchmarkPopCount2-4    100000000            16.8 ns/op
BenchmarkPopCount3-4    100000000            22.0 ns/op
BenchmarkPopCount4-4    100000000            14.7 ns/op
ok      ch2/popcount    7.310s
*/

package popcount

import (
    "fmt"
    "testing"
)

func BenchmarkPopCount(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PopCount(uint64(i))
    }
}

func BenchmarkPopCount2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PopCount2(uint64(i))
    }
}

func BenchmarkPopCount3(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PopCount3(uint64(i))
    }
}

func BenchmarkPopCount4(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PopCount4(uint64(i))
    }
}

func TestPopCountVersions(t *testing.T) {
    runPopCount(t, 1234567898765431)
    runPopCount(t, 664652156845612141)
}

func runPopCount(t *testing.T, input uint64) {
    fmt.Printf("PopCount(%[1]d) = %d, PopCount2(%[1]d) = %d, PopCount3(%[1]d) = %d, PopCount4(%[1]d) = %d\n",
        input, PopCount(input), PopCount2(input), PopCount3(input), PopCount4(input))

    if PopCount(input) != PopCount2(input) || PopCount2(input) != PopCount3(input) || PopCount3(input) != PopCount4(input) {
        t.Errorf("PopCount(%[1]d) = %d, PopCount2(%[1]d) = %d, PopCount3(%[1]d) = %d, PopCount4(%[1]d) = %d\n",
            input, PopCount(input), PopCount2(input), PopCount3(input), PopCount4(input))
    }
}
