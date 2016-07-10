package bank_test

import (
	"fmt"
	"testing"

	"./bank"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	var balance int
	go func() {
		if bank.Withdraw(300) {
			balance = 0
		} else {
			balance = 300
		}
		fmt.Println("=", balance)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done
	<-done

	if got, want := bank.Balance(), balance; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
