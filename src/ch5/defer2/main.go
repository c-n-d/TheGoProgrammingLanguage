/*
Defer2 demonstrates a deferred call to runtime.Stack during panic.

$ go run src/ch5/defer2/main.go
f(3)
f(2)
f(1)
defer 1
defer 2
defer 3
goroutine 1 [running]:
main.printStack()
	src/ch5/defer2/main.go:23
main.f(0)
	src/ch5/defer2/main.go:28
main.f(1)
	src/ch5/defer2/main.go:30
main.f(2)
	src/ch5/defer2/main.go:30
main.f(3)
	src/ch5/defer2/main.go:30
main.main()
	src/ch5/defer2/main.go:18
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.f(0)
	src/ch5/defer2/main.go:28
main.f(1)
	src/ch5/defer2/main.go:30
main.f(2)
	src/ch5/defer2/main.go:30
main.f(3)
	src/ch5/defer2/main.go:30
main.main()
	src/ch5/defer2/main.go:18
exit status 2
*/

package main

import (
    "fmt"
    "os"
    "runtime"
)

func main() {
    defer printStack()
    f(3)
}

func printStack() {
    var buf [4096]byte
    n := runtime.Stack(buf[:], false)
    os.Stdout.Write(buf[:n])
}

func f(x int) {
    fmt.Printf("f(%d)\n", x+0/x) // panics if x = 0
    defer fmt.Printf("defer %d\n", x) // panics if x = 0
    f(x-1)
}
