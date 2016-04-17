/*
Exercise 5.15 - Write variadic min / max functions (and a variant that requires at least one argument)

$ go run src/ch5/ex_5_15/main.go
min([1 2 3 4 5])=1
max([1 2 3 4 5])=5
min([])=2147483647
max([])=-2147483648
min1([])=2147483647, please provide at least 1 value
max1([])=-2147483648, please provide at least 1 value
*/

package main

import (
    "fmt"
    "math"
)

func min(vals...int) int {
    min := math.MaxInt32
    for _, val := range vals {
        if val < min {
            min = val
        }
    }
    return min
}

func min1(vals...int) (int, error) {
    min := math.MaxInt32
    if len(vals) == 0 {
        return min, fmt.Errorf("please provide at least 1 value")
    }
    for _, val := range vals {
        if val < min {
            min = val
        }
    }
    return min, nil
}

func max(vals...int) int {
    max := math.MinInt32
    for _, val := range vals {
        if val > max {
            max = val
        }
    }
    return max
}

func max1(vals...int) (int, error) {
    max := math.MinInt32
    if len(vals) == 0 {
        return max, fmt.Errorf("please provide at least 1 value")
    }
    for _, val := range vals {
        if val > max {
            max = val
        }
    }
    return max, nil
}

func main() {
    values := []int{1,2,3,4,5}
    empty := []int{}

    fmt.Printf("min(%v)=%v\n", values, min(values...))
    fmt.Printf("max(%v)=%v\n", values, max(values...))
    fmt.Printf("min(%v)=%v\n", empty, min(empty...))
    fmt.Printf("max(%v)=%v\n", empty, max(empty...))

    res, err := min1(empty...)
    fmt.Printf("min1(%v)=%v, %v\n", empty,res, err)
    res, err = max1(empty...)
    fmt.Printf("max1(%v)=%v, %v\n", empty, res, err)
}
