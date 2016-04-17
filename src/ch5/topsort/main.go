/*
Toposort prints the nodes of a DAG in topological order

$ go run src/ch5/topsort/main.go
1:	intro to programming
2:	discrete math
3:	data structures
4:	algorithms
5:	linear algebra
6:	calculus
7:	formal languages
8:	computer organization
9:	compilers
10:	databases
11:	operating systems
12:	networks
13:	programming languages
*/

package main

import (
    "fmt"
    "sort"
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
    "programming languages": {"data structures", "computer organization"}}

func main() {
    for i, course := range topoSort(prereqs) {
        fmt.Printf("%d:\t%s\n", i+1, course)
    }
}

func topoSort(m map[string][]string) []string {
    var order []string
    seen := make(map[string]bool)
    var visitAll func(items []string)

    visitAll = func(items []string) {
        for _, item := range items {
            if !seen[item] {
                seen[item] = true
                visitAll(m[item])
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
