/*
$ go test -v ch9/memo5
=== RUN   Test
https://golang.org, 250.238778ms, 7902 bytes
http://play.golang.org/, 103.412921ms, 5862 bytes
https://blog.golang.org/, 58.986788ms, 47304 bytes
https://tour.golang.org/, 55.624506ms, 2615 bytes
https://golang.org, 5.072µs, 7902 bytes
http://play.golang.org/, 7.25µs, 5862 bytes
https://blog.golang.org/, 1.979µs, 47304 bytes
https://tour.golang.org/, 1.716µs, 2615 bytes
--- PASS: Test (0.47s)
=== RUN   TestConcurrent
https://golang.org, 48.302079ms, 7902 bytes
https://golang.org, 48.350262ms, 7902 bytes
https://tour.golang.org/, 48.478313ms, 2615 bytes
https://tour.golang.org/, 48.488566ms, 2615 bytes
https://blog.golang.org/, 55.802067ms, 47304 bytes
https://blog.golang.org/, 55.825404ms, 47304 bytes
http://play.golang.org/, 91.611804ms, 5862 bytes
http://play.golang.org/, 91.609766ms, 5862 bytes
--- PASS: TestConcurrent (0.09s)
PASS
ok  	ch9/memo5	0.584s

$ go test -run=TestConcurrent -race ch9/memo5
ok  	ch9/memo5	1.362s
*/

package memo_test

import (
    "testing"

    "ch9/memo5"
    "ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
    m := memo.New(httpGetBody)
    defer m.Close()
    memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
    m := memo.New(httpGetBody)
    defer m.Close()
    memotest.Concurrent(t, m)
}
