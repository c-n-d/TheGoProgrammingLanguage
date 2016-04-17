/*
Exercise 5.10 - Expands on Topsort to use maps instead of slices and eliminates
                the initial sorting.

$ go run src/ch5/ex_5_10/main.go
1:	intro to programming
2:	discrete math
3:	formal languages
4:	data structures
5:	algorithms
6:	linear algebra
7:	calculus
8:	computer organization
9:	compilers
10:	programming languages
11:	databases
12:	operating systems
13:	networks
*/

package main

import (
    "fmt"
)

/*
         comporg
       ↗   ↑   ↖
net → os   pl  compi → flang
       ↘   ↓  ↙       ↓
algo →   dstruct   →  dmath → intro
          ↑
          db

calc → linalg
*/

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

    var visitAll func(m map[string][]string)

    visitAll = func(m map[string][]string) {
        if len(m) == 0 {
            return
        }

        // for each course in the map, build a map of the course's prerequs to a list of the each values prereq
        // and expolre the subtree. After entire subtree is explored, mark the key as seen and append it to the order
        for k, v := range m {
            tmp := make(map[string][]string)

            for _, p := range v {
                if !seen[p] {
                    tmp[p] = prereqs[p]
                }
            }

            visitAll(tmp)

            if !seen[k] {
                seen[k] = true
                order = append(order, k)
            }
        }
    }

    visitAll(m)
    return order
}
