// Package bank provides a concurrency-safe bank with one account.
package bank

type request struct {
	amount int
	result chan bool
}

var deposits = make(chan int)         // send amount to deposit
var withdrawals = make(chan *request) // send amount to withdraw
var balances = make(chan int)         // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	result := make(chan bool)
	withdrawals <- &request{amount, result}
	return <-result
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case req := <-withdrawals:
			if balance-req.amount < 0 {
				req.result <- false
			} else {
				balance -= req.amount
				req.result <- true
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
