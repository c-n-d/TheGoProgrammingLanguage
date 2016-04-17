/*
Exercise 5.14 - Use breadthFirst to explore other data structures

$ go run src/ch5/ex_5_14/main.go
data structures
databases
discrete math
formal languages
algorithms
compilers
operating systems
programming languages
calculus
networks
intro to programming
computer organization
linear algebra
*/

package main 

import "fmt"

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
}

func breadthFirst(f func(item string) []string, worklist []string) {
    seen := make(map[string]bool)
    for len(worklist) > 0 {
        items := worklist
        worklist = nil
        for _, item := range items {
            if !seen[item] {
                seen[item] = true
                worklist = append(worklist, f(item)...)
            }
        }
    }
}

func getPrereqs(s string) []string {
    fmt.Println(s)
    return prereqs[s]
}

func keys(m map[string][]string) (keys []string) {
    for k, _ := range m {
        keys = append(keys, k)
    }
    return
}

func main() {
    breadthFirst(getPrereqs, keys(prereqs))
}
