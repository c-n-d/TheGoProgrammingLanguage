/*
Testing various shop configurations

$ go test -bench=. ch8/cake
testing: warning: no tests to run
PASS
Benchmark-4                  	       5	 239208571 ns/op
BenchmarkBuffers-4           	       5	 241486531 ns/op
BenchmarkVariable-4          	       5	 272463426 ns/op
BenchmarkVariableBuffers-4   	       5	 249616082 ns/op
BenchmarkSlowIcing-4         	       1	1086056792 ns/op
BenchmarkSlowIcingManyIcers-4	       5	 270390708 ns/op
ok  	ch8/cake	11.091s
*/

package cake_test

import (
    "testing"
    "time"

    "ch8/cake"
)

var defaults = cake.Shop{
    Verbose:      testing.Verbose(),
    Cakes:        20,
    BakeTime:     10 * time.Millisecond,
    NumIcers:     1,
    IceTime:      10 * time.Millisecond,
    InscribeTime: 10 * time.Millisecond,
}

func Benchmark(b *testing.B) {
    // Baseline: one baker, icer and inscriber
    // Each takes 10ms, no buffers
    cakeshop := defaults
    cakeshop.Work(b.N) // 239ms
}

func BenchmarkBuffers(b *testing.B) {
    // Adding buffers has no effect
    cakeshop := defaults
    cakeshop.BakeBuf = 10
    cakeshop.IceBuf = 10
    cakeshop.Work(b.N) // 241ms
}

func BenchmarkVariable(b *testing.B) {
    // Adding variable to rate of each step
    // incrase total time due to channel delays
    cakeshop := defaults
    cakeshop.BakeStdDev = cakeshop.BakeTime / 4
    cakeshop.IceStdDev = cakeshop.IceTime / 4
    cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
    cakeshop.Work(b.N) // 272ms
}

func BenchmarkVariableBuffers(b *testing.B) {
    // Adding channel buffers reduced delays resulting from variablity
    cakeshop := defaults
    cakeshop.BakeStdDev = cakeshop.BakeTime / 4
    cakeshop.IceStdDev = cakeshop.IceTime / 4
    cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
    cakeshop.BakeBuf = 10
    cakeshop.IceBuf = 10
    cakeshop.Work(b.N) // 249ms
}

func BenchmarkSlowIcing(b *testing.B) {
    // Making middle stage slower adds directly to the critical path
    cakeshop := defaults
    cakeshop.IceTime = 50 * time.Millisecond
    cakeshop.Work(b.N) // 1086ms
}

func BenchmarkSlowIcingManyIcers(b *testing.B) {
    // Adding more icing cooks reduces the cost of icing
    // to its sequential component, following Amdahl's Law
    cakeshop := defaults
    cakeshop.IceTime = 50 * time.Millisecond
    cakeshop.NumIcers = 5
    cakeshop.Work(b.N) // 270ms
}
