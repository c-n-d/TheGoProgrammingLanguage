/*
treesort defines a tree stuct contain a value and reference to sub trees

1. Sort adds all the provided values to a new tree, and returns the sorted values
2. appendValues does an pre-order traversal of the b-tree
3. add places the value at the correct location in the tree

$ go run src/ch4/treesort/main.go
Original values: [2 4 6 8 1 3 5 7 9]
Sorted values: [1 2 3 4 5 6 7 8 9]
*/

package main

import "fmt"

type tree struct {
    value int
    left, right *tree
}

func main() {
    arr := []int{2,4,6,8,1,3,5,7,9}

    fmt.Printf("Original values: %v\n", arr)
    Sort(arr)
    fmt.Printf("Sorted values: %v\n", arr)
}

// Sort sorts values in place
func Sort(values []int) {
    var root *tree
    for _, v := range values {
        root = add(root, v)
    }
    appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice
func appendValues(values []int, t*tree) []int {
    if t != nil {
        values = appendValues(values, t.left)
        values = append(values, t.value)
        values = appendValues(values, t.right)
    }
    return values
}

func add(t *tree, value int) *tree {
    if t == nil {
        // Equivalent to return &tree{value: value}.
        t = new(tree)
        t.value = value
        return t
    }
    if value < t.value {
        t.left = add(t.left, value)
    } else {
        t.right = add(t.right, value)
    }
    return t
}
