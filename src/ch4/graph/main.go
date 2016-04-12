/*
graph uses a map of maps to represent a node-to-node edge relationship

$ go run src/ch4/graph/main.go
addEdge(a, b)
addEdge(b, c)
addEdge(a, c)
hasEdge(a, b)=true
hasEdge(b, c)=true
hasEdge(a, c)=true
hasEdge(b, a)=false
hasEdge(c, d)=false
*/

package main

import "fmt"

type Edge struct {
    from, to string
}

var graph = make(map[string]map[string]bool)

func main() {
    es := []Edge{Edge{"a", "b"},
                 Edge{"b", "c"},
                 Edge{"a", "c"}}

    for _, e := range es {
        addEdge(e.from, e.to)
        fmt.Printf("addEdge(%s, %s)\n", e.from, e.to)
    }

     for _, e := range es {
        fmt.Printf("hasEdge(%s, %s)=%t\n", e.from, e.to, hasEdge(e.from, e.to))
    }

    fmt.Printf("hasEdge(%s, %s)=%t\n", "b", "a", hasEdge("b", "a"))
    fmt.Printf("hasEdge(%s, %s)=%t\n", "c", "d", hasEdge("c", "d"))
}

func addEdge(from, to string) {
    edges := graph[from]

    if edges == nil {
        edges = make(map[string]bool)
        graph[from] = edges
    }

    edges[to] = true
}

func hasEdge(from, to string) bool {
    return graph[from][to]
}
