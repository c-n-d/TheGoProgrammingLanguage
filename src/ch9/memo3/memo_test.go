/*
$ go test -v ch9/memo3
=== RUN   Test
https://golang.org, 293.954451ms, 7902 bytes
http://play.golang.org/, 108.768483ms, 5862 bytes
https://blog.golang.org/, 60.438633ms, 47304 bytes
https://tour.golang.org/, 55.852706ms, 2615 bytes
https://golang.org, 479ns, 7902 bytes
http://play.golang.org/, 186ns, 5862 bytes
https://blog.golang.org/, 147ns, 47304 bytes
https://tour.golang.org/, 240ns, 2615 bytes
--- PASS: Test (0.52s)
=== RUN   TestConcurrent
https://tour.golang.org/, 46.675867ms, 2615 bytes
https://golang.org, 49.568827ms, 7902 bytes
https://blog.golang.org/, 54.811206ms, 47304 bytes
https://tour.golang.org/, 57.327206ms, 2615 bytes
https://golang.org, 65.351174ms, 7902 bytes
https://blog.golang.org/, 84.893847ms, 47304 bytes
http://play.golang.org/, 88.776182ms, 5862 bytes
http://play.golang.org/, 112.313743ms, 5862 bytes
--- PASS: TestConcurrent (0.11s)
PASS
ok  	ch9/memo3	0.661s

$ go test -run=TestConcurrent -race ch9/memo3
ok  	ch9/memo3	1.446s
*/

package memo_test

import (
    "testing"

    "ch9/memo3"
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
