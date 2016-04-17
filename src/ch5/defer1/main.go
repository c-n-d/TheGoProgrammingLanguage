/*
Defer1 demonstrates a deferred call during panic

$ go run src/ch5/defer1/main.go
f(3)
f(2)
f(1)
defer 1
defer 2
defer 3
panic: runtime error: integer divide by zero
main.f(0)
	src/ch5/defer1/main.go:11
main.f(1)
	src/ch5/defer1/main.go:13
main.f(2)
	src/ch5/defer1/main.go:13
main.f(3)
	src/ch5/defer1/main.go:13
main.main()
	src/ch5/defer1/main.go:16
*/

package main

import "fmt"

func main() {
    f(3)
}

func f(x int) {
    fmt.Printf("f(%d)\n", x+0/x) // panics if x = 0
    defer fmt.Printf("defer %d\n", x) // panics if x = 0
    f(x-1)
}
