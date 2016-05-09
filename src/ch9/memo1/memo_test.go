/*
$ go test -v ch9/memo1
=== RUN   Test
https://golang.org, 179.074262ms, 7902 bytes
http://play.golang.org/, 99.161096ms, 5862 bytes
https://blog.golang.org/, 59.891137ms, 47304 bytes
https://tour.golang.org/, 53.34476ms, 2615 bytes
https://golang.org, 401ns, 7902 bytes
http://play.golang.org/, 125ns, 5862 bytes
https://blog.golang.org/, 117ns, 47304 bytes
https://tour.golang.org/, 165ns, 2615 bytes
--- PASS: Test (0.39s)
=== RUN   TestConcurrent
https://tour.golang.org/, 45.723253ms, 2615 bytes
https://blog.golang.org/, 51.402625ms, 47304 bytes
https://golang.org, 52.532685ms, 7902 bytes
https://golang.org, 62.710747ms, 7902 bytes
https://tour.golang.org/, 63.268598ms, 2615 bytes
https://blog.golang.org/, 69.895931ms, 47304 bytes
http://play.golang.org/, 91.171119ms, 5862 bytes
http://play.golang.org/, 111.532072ms, 5862 bytes
--- PASS: TestConcurrent (0.11s)
PASS
ok  	ch9/memo1	0.524s

$ go test -run=TestConcurrent -race ch9/memo1
https://golang.org, 420.189858ms, 7902 bytes
==================
WARNING: DATA RACE
Write by goroutine 22:
  runtime.mapassign1()
      ~/go/src/runtime/hashmap.go:429 +0x0
  ch9/memo1.(*Memo).Get()
      ...src/ch9/memo1/memo.go:30 +0x205
  ch9/memotest.Concurrent.func1()
      ...src/ch9/memotest/memotest.go:72 +0xbd

Previous write by goroutine 12:
  runtime.mapassign1()
      /usr/local/go/src/runtime/hashmap.go:429 +0x0
  ch9/memo1.(*Memo).Get()
      ...src/ch9/memo1/memo.go:30 +0x205
  ch9/memotest.Concurrent.func1()
      ...src/ch9/memotest/memotest.go:72 +0xbd
...
==================
...
Found 1 data race(s)
FAIL	ch9/memo1	1.466s
*/

package memo_test

import (
    "testing"

    "ch9/memo1"
    "ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
    m := memo.New(httpGetBody)
    memotest.Sequential(t, m)
}

// Note: not concurrent safe, test fails
func TestConcurrent(t *testing.T) {
    m := memo.New(httpGetBody)
    memotest.Concurrent(t, m)
}
