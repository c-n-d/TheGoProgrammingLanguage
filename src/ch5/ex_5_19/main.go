/*
Exercise 5.19 - Use panic and recover to write a function that contains no return statement
                yet returns a non-zero value.

$ go run src/ch5/ex_5_19/main.go
1460919345
1460919347
1460919349
*/

package main

import (
    "fmt"
    "time"
)

func main()  {
    fmt.Println(getUnixTime())
    time.Sleep(2 * time.Second)
    fmt.Println(getUnixTime())
    time.Sleep(2 * time.Second)
    fmt.Println(getUnixTime())
}

// Function using panic and recover to return the current time in
// seconds since epoch
func getUnixTime() (res int64) {
    type always struct{}

    defer func() {
        switch p:= recover(); p {
            case nil:
                res = -1
            case always{}:
                res = time.Now().Unix()
            default:
                panic(p)
        }
    }()

    panic(always{})
}
