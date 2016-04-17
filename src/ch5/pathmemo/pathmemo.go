/*
PathMemo is a data structure for manintaining the state of the currently
expolred path in a graph.

Nodes from the graph are added and removed as exploration progresses.

Users may query if an element exists in the current path.
*/

package pathmemo

type PathMemo struct {
    // The list of elements added to the path (LIFO)
    path []string
    // Memo is queried to determine if a item is in the current path
    memo map[string]bool
}

// Add item to the path
func (pm *PathMemo) Push(item string) {
    pm.path = append(pm.path, item)
    pm.memo[item] = true
}

// Removes the last item from the path
func (pm *PathMemo) Pop() {
    item := pm.path[len(pm.path)-1]
    pm.memo[item] = false
    pm.path = pm.path[:len(pm.path)-1]
}

// Returns whether the item exists in the path and the path to this point
func (pm *PathMemo) InPath(item string) (bool, []string) {
    return pm.memo[item], pm.path
}

// Create a new PathMemo with an empty memo
func NewPathMemo() *PathMemo {
    return &PathMemo{memo: make(map[string]bool)}
}
