/*
Netflag demonstrates bit operations on an integer type.
Use of iota to create a sequence of shifted 1's bits.

$ go run src/ch3/netflag/main.go
Flag is: 11111111
IsUp(11111111) = true
TurnDown(11111111) = 11111110
IsUp(11111110) = false

Flag is: 11101101
IsCast(11101101) = false
SetBroadcast(11101101) = 11101111
IsCast(11101111) = true
*/

package main

import 
    "fmt"

type Flags uint

const (
    FlagUp Flags = 1 << iota // is up
    FlagBroadcast            // supports broadcast access capability
    FlagLoopback             // is a loopback interface
    FlagPointToPoint         // belongs to a point-to-point link
    FlagMulticast            // supports multicast access capability
)

func main() {
    turnDownExample()
    castExample()
}

func turnDownExample() {
    // flag1 = 0b1111_1111
    flag1 := Flags(0xFF)

    fmt.Printf("Flag is: %b\n", flag1)
    fmt.Printf("IsUp(%b) = %t\n", flag1, IsUp(flag1))

    flag1Before := flag1
    TurnDown(&flag1)

    fmt.Printf("TurnDown(%b) = %b\n", flag1Before, flag1)
    fmt.Printf("IsUp(%b) = %t\n\n", flag1, IsUp(flag1))
}

func castExample() {
    // flag2 = 0b1110_1101
    flag2 := Flags(0xED)

    fmt.Printf("Flag is: %b\n", flag2)
    fmt.Printf("IsCast(%b) = %t\n", flag2, IsCast(flag2))

    flag2Before := flag2
    SetBroadcast(&flag2)

    fmt.Printf("SetBroadcast(%b) = %b\n", flag2Before, flag2)
    fmt.Printf("IsCast(%b) = %t\n", flag2, IsCast(flag2))
}

func IsUp(v Flags) bool { return v & FlagUp == FlagUp }
func TurnDown(v *Flags) { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool { return v & (FlagBroadcast | FlagMulticast) != 0 }
