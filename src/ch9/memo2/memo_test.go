/*
$ go test -v ch9/memo2
=== RUN   Test
https://golang.org, 173.022653ms, 7902 bytes
http://play.golang.org/, 100.785829ms, 5862 bytes
https://blog.golang.org/, 58.813403ms, 47304 bytes
https://tour.golang.org/, 55.008691ms, 2615 bytes
https://golang.org, 498ns, 7902 bytes
http://play.golang.org/, 187ns, 5862 bytes
https://blog.golang.org/, 100ns, 47304 bytes
https://tour.golang.org/, 227ns, 2615 bytes
--- PASS: Test (0.39s)
=== RUN   TestConcurrent
http://play.golang.org/, 89.391227ms, 5862 bytes
https://golang.org, 136.004515ms, 7902 bytes
https://tour.golang.org/, 186.641638ms, 2615 bytes
https://blog.golang.org/, 239.325878ms, 47304 bytes
http://play.golang.org/, 239.347783ms, 5862 bytes
https://golang.org, 239.369963ms, 7902 bytes
https://tour.golang.org/, 239.350884ms, 2615 bytes
https://blog.golang.org/, 239.359683ms, 47304 bytes
--- PASS: TestConcurrent (0.24s)
PASS
ok  	ch9/memo2	0.649s

$ go test -run=TestConcurrent -race ch9/memo2
ok  	ch9/memo2	1.693s
*/

package memo_test

import (
    "testing"

    "ch9/memo2"
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
