/*
Exercise 5.11 - Expands on TopSort to report cycles

$ go run src/ch5/ex_5_11/main.go
2016/04/17 12:17:13 Cycle scaused by adding [calculus] to [calculus linear algebra]
exit status 1
*/

package main

import (
    "fmt"
    "log"
    "sort"

    "ch5/pathmemo"
)

// prereqs  maps computer science courses to their prerequisites
var prereqs = map[string][]string {
    "algorithms": {"data structures"},
    "calculus": {"linear algebra"},
    "compilers": {
        "data structures",
        "formal languages",
        "computer organization",
    },
    "data structures": {"discrete math"},
    "databases": {"data structures"},
    "discrete math": {"intro to programming"},
    "formal languages": {"discrete math"},
    "networks": {"operating systems"},
    "operating systems": {"data structures", "computer organization"},
    "programming languages": {"data structures", "computer organization"},
    "linear algebra": {"calculus"},                                        // CYCLE!
}

func main() {
    for i, course := range topoSort(prereqs) {
        fmt.Printf("%d:\t%s\n", i+1, course)
    }
}

func topoSort(m map[string][]string) []string {
    var order []string
    seen := make(map[string]bool)
    memo := pathmemo.NewPathMemo()

    var visitAll func(items []string)

    visitAll = func(items []string) {
        for _, item := range items {
            // A cycle exists if the item being explored already exists in the
            // previously explored path
            if isCycle, cycle := memo.InPath(item); isCycle {
                log.Fatalf("Cycle scaused by adding [%s] to %v\n", item, cycle)
            }

            if !seen[item] {
                seen[item] = true

                // Add the item to the path before exploring the subgraph. Remove afterwards.
                memo.Push(item)
                visitAll(m[item])
                memo.Pop()

                order = append(order, item)
            }
        }
    }

    var keys []string
    for key := range m {
        keys = append(keys, key)
    }

    sort.Strings(keys)
    visitAll(keys)
    return order
}
