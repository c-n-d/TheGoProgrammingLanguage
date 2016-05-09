/*
Package memo provides a concurrency memoization of a function of type Func

Removes redundant work computing f(key). First goroutine computed f(key)
and all subsequent wait on the result to be ready
*/

package memo

import "sync"

type entry struct {
    res   result
    ready chan struct{} // close when res is ready
}

func New(f Func) *Memo {
    return &Memo{f: f, cache: make(map[string]*entry)}
}

// A Memo caches the results of calling a Func
type Memo struct {
    f     Func
    mu    sync.Mutex // guards cache
    cache map[string]*entry
}

// A Func is the type of the function to memoize
type Func func(key string) (interface{}, error)

type result struct {
    value interface{}
    err   error
}

func (memo *Memo) Get(key string) (interface{}, error) {
    memo.mu.Lock()
    e := memo.cache[key]
    if e == nil {
        // This is the first request for this key.
        // This goroutine is responsible for computing the
        // value and broadcasting the ready condition.
        e = &entry{ready: make(chan struct{})}
        memo.cache[key] = e
        memo.mu.Unlock()

        e.res.value, e.res.err = memo.f(key)

        close(e.ready) // broadcast the read condition
    } else {
        memo.mu.Unlock()

        <-e.ready // wait for the ready condition
    }
    return e.res.value, e.res.err
}
