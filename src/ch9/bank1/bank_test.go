/*
Test deposit and balance functions of the bank

$ go test -bench=. ch9/bank1
= 100
= 300
PASS
ok  	ch9/bank1	0.008s
*/

package bank_test

import (
    "fmt"
    "testing"

    "ch9/bank1"
)

func TestBank(t *testing.T) {
    done := make(chan struct{})

    // Alice
    go func() {
        bank.Deposit(200)
        fmt.Println("=", bank.Balance())
        done <- struct{}{}
    }()

    // Bob
    go func() {
        bank.Deposit(100)
        fmt.Println("=", bank.Balance())
        done <- struct{}{}
    }()

    <-done
    <-done

    if got, want := bank.Balance(), 300; got != want {
        t.Errorf("Balance = %d, want %d", got, want)
    }
}
