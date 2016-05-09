/*
Exercise 9.1 - Add a function Withdraw(amount int) bool to bank1
*/

package bank

type WithdrawalSlip struct {
    amount  int
    success chan<- bool
}

var deposits   = make(chan int) // send amount to deposit
var balances   = make(chan int) // receive balance
var withdrawals = make(chan WithdrawalSlip) // send amount to withdrawl

func Deposit(amount int)   { deposits <- amount }
func Balance() int         { return <-balances }

func Withdraw(amount int) bool {
    success := make(chan bool)
    withdrawals <- WithdrawalSlip{ amount, success }
    return <-success
}

func teller() {
    var balance int // balance is confined to the teller goroutine
    for {
        select {
            case amount := <-deposits:
                balance += amount
            case balances <- balance:
            case slip := <-withdrawals:
                if slip.amount > balance {
                    slip.success <- false
                    continue
                }
                balance -= slip.amount
                slip.success <- true
        }
    }
}

func init() {
    go teller() // start the monitor goroutine
}
