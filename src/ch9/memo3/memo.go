/*
Package memo provides a concurrency memoization of a function of type Func

Inprove performace, only access to cache are in a CS
*/

package memo

import "sync"

// A Memo caches the results of calling a Func
type Memo struct {
    f     Func
    mu    sync.Mutex // guards cache
    cache map[string]result
}

// A Func is the type of the function to memoize
type Func func(key string) (interface{}, error)

type result struct {
    value interface{}
    err   error
}

func New(f Func) *Memo {
    return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
    memo.mu.Lock()
    res, ok := memo.cache[key]
    memo.mu.Unlock()
    if !ok {
        res.value, res.err = memo.f(key)
        // Between the two critical sections, serveral goroutines
        // may race to compute f(key) and update the map
        memo.mu.Lock()
        memo.cache[key] = res
        memo.mu.Unlock()
    }
    return res.value, res.err
}
