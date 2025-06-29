package wallet

import "testing"

func TestDeposit(t *testing.T) {
	wallet := Wallet{}
	newBalance := Bitcoin(10)

	wallet.Deposit(newBalance)

	got := wallet.Balance
	want := newBalance 

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
