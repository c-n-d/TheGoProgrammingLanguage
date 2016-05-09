/*
Exercise 9.4 - Construct a pipeline that connects an arbitrary number of goroutines with channels.
               Max number of stages w/o running out of mem? How long to transit the entire pipeline?

$ go run src/ch9/ex_9_4/main.go
Created 1874628 stages

32.208159996s
32.217069927s
32.217075814s
32.217076678s
32.217071293s
32.217072328s
32.217066332s
32.217066483s
32.217060495s
32.217060457s

$ go run src/ch9/ex_9_4/main.go
Created 1966977 stages

26.450919855s
26.458929248s
26.458944193s
26.45894737s
26.458944854s
26.458953914s
26.460605412s
26.460606783s
26.460612721s
26.460613634s
*/

package main

import (
    "fmt"
    "sync"
    "time"
)

var wg sync.WaitGroup

type TimeTravler struct {
    travler int
    created time.Time
}

func main() {
    wg.Add(1)

    startChan := make(chan TimeTravler)
    wait := make(chan struct{})

    go start(startChan, wait)
    lastCreate := time.Now()

    prevOut := startChan
    stageDepth := 0

    for {
        out := make(chan TimeTravler)
        go stage(out, prevOut, stageDepth)
        prevOut = out
        // If it took > 20 seconds to create the goroutine, probably getting close to
        // memory limit. Naive termination for now...
        if time.Since(lastCreate) > (20 * time.Second) {
            break
        }

        lastCreate = time.Now()
        stageDepth++ 
    }

    fmt.Printf("Created %d stages\n\n", stageDepth)

    go end(prevOut)
    wait<- struct{}{}
    close(wait)

    wg.Wait()
}

func start(out chan<- TimeTravler, wait <-chan struct{}) {
    // Blocks until the entire pipeline is set up
    <- wait

    for x := 0; x < 10; x++ {
        out <- TimeTravler{x, time.Now()}
    }
    close(out)
}

func stage(out chan<- TimeTravler, in <-chan TimeTravler, id int) {
    for input := range in {
        out <- input
    }
    close(out)
}

func end(in <-chan TimeTravler) {
    for input := range in {
        fmt.Printf("%s\n", time.Since(input.created))
    }
    wg.Done()
}
