/*
Test deposit, balance, and withdrawals functions of the bank

$ go test -bench=. ch9/ex_9_1
= 300
= 300
withdraw [175] success? = true. Balance = 125
withdraw [500] success? = false. Balance = 125
PASS
ok  	ch9/ex_9_1	0.011s
*/

package bank_test

import (
    "fmt"
    "testing"

    "ch9/ex_9_1"
)

func TestBank(t *testing.T) {
    done := make(chan struct{})

    // +200
    go func() {
        bank.Deposit(200)
        fmt.Println("=", bank.Balance())
        done <- struct{}{}
    }()

    // +100
    go func() {
        bank.Deposit(100)
        fmt.Println("=", bank.Balance())
        done <- struct{}{}
    }()

    <-done
    <-done

    // -500
    go func() {
        success := bank.Withdraw(500)
        fmt.Printf("withdraw [500] success? = %t. Balance = %d\n", success, bank.Balance())
        done <- struct{}{}
    }()

    // -175
    go func() {
        success := bank.Withdraw(175)
        fmt.Printf("withdraw [175] success? = %t. Balance = %d\n", success, bank.Balance())
        done <- struct{}{}
    }()

    <-done
    <-done

    if got, want := bank.Balance(), 125; got != want {
        t.Errorf("Balance = %d, want %d", got, want)
    }
}
