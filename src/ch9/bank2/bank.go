/*
Package bank provides a concurrency safe bank with one account

Uses a buffered channel as a semaphore
*/

package bank

var (
    sema    = make(chan struct{}, 1) // a binary semaphore guarding balance
    balance int
)

func Deposit(amount int) {
    sema <- struct{}{} // aquire a token
    balance = balance + amount
    <-sema // release token
}

func Balance() int {
    sema <- struct{}{} // aquire a token
    b := balance
    <-sema // release token
    return b
}
