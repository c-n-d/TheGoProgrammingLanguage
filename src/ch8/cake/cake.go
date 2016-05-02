/*
Pacakge cake provides a simualation of a concurrent cake shop with numberous parameters

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

package cake

import (
    "fmt"
    "math/rand"
    "time"
)

type Shop struct {
    Verbose        bool
    Cakes          int           // number of cakes to bake
    BakeTime       time.Duration // time to bake one cake
    BakeStdDev     time.Duration // std dev of bake time
    BakeBuf        int           // buffer slots between baking and icing
    NumIcers       int           // number of cooks doing icing
    IceTime        time.Duration // time to ice one cake
    IceStdDev      time.Duration // std dev of icing time
    IceBuf         int           // buffer slots between icing and inscribing
    InscribeTime   time.Duration // time to inscribe one cake
    InscribeStdDev time.Duration // std dev of inscribing time
}

type cake int

func (s *Shop) baker(baked chan<- cake) {
    for i := 0; i < s.Cakes; i++ {
        c := cake(i)
        if s.Verbose {
            fmt.Println("baking", c)
        }
        work(s.BakeTime, s.BakeStdDev)
        baked <- c
    }
    close(baked)
}

func (s *Shop) icer(iced chan<- cake, baked <-chan cake) {
    for c := range baked {
        if s.Verbose {
            fmt.Println("icing", c)
        }
        work(s.IceTime, s.IceStdDev)
        iced <- c
    }
}

func (s *Shop) inscriber(iced <-chan cake) {
    for i := 0; i < s.Cakes; i++ {
        c := <- iced
        if s.Verbose {
            fmt.Println("inscribing", c)
        }
        work(s.BakeTime, s.BakeStdDev)
        if s.Verbose {
            fmt.Println("finished", c)
        }
    }
}

// Work runs the simulation 'runs' times
func (s *Shop) Work(runs int) {
    for run := 0; run < runs; run++ {
        baked := make(chan cake, s.BakeBuf)
        iced := make(chan cake, s.IceBuf)
        go s.baker(baked)
        for i := 0; i < s.NumIcers; i++ {
            go s.icer(iced, baked)
        }
        s.inscriber(iced)
    }
}

func work(d, stddev time.Duration) {
    delay := d + time.Duration(rand.NormFloat64() * float64(stddev))
    time.Sleep(delay)
}
