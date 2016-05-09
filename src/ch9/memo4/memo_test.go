/*
$ go test -v ch9/memo4
# ch9/memo4
src/ch9/memo4/memo.go:53: undefined: res in res.value
src/ch9/memo4/memo.go:53: undefined: res in res.err
FAIL	ch9/memo4 [build failed]
216-243-38-56:TheGoPL computer$ go test -v ch9/memo4
=== RUN   Test
https://golang.org, 315.859676ms, 7902 bytes
http://play.golang.org/, 98.990563ms, 5862 bytes
https://blog.golang.org/, 68.70798ms, 47304 bytes
https://tour.golang.org/, 60.110911ms, 2615 bytes
https://golang.org, 5.941µs, 7902 bytes
http://play.golang.org/, 1.492µs, 5862 bytes
https://blog.golang.org/, 1.385µs, 47304 bytes
https://tour.golang.org/, 906ns, 2615 bytes
--- PASS: Test (0.54s)
=== RUN   TestConcurrent
https://tour.golang.org/, 46.603411ms, 2615 bytes
https://tour.golang.org/, 46.074585ms, 2615 bytes
https://golang.org, 49.630541ms, 7902 bytes
https://golang.org, 48.829993ms, 7902 bytes
https://blog.golang.org/, 54.21721ms, 47304 bytes
https://blog.golang.org/, 54.052301ms, 47304 bytes
http://play.golang.org/, 89.690759ms, 5862 bytes
http://play.golang.org/, 88.544597ms, 5862 bytes
--- PASS: TestConcurrent (0.09s)
PASS
ok  	ch9/memo4	0.662s

$ go test -run=TestConcurrent -race ch9/memo4
ok  	ch9/memo4	1.364s
*/

package memo_test

import (
    "testing"

    "ch9/memo4"
    "ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
    m := memo.New(httpGetBody)
    memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
    m := memo.New(httpGetBody)
    memotest.Concurrent(t, m)
}
