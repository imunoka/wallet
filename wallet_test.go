package wallet

import "testing"

func TestDeposit(t *testing.T) {
	wallet := Wallet{}
	newBalance := Bitcoin(10)

	wallet.Deposit(newBalance)

	assertBalance(t, wallet, newBalance)
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	balance := wallet.Balance

	if balance != want {
		t.Errorf("got %s, want %s", balance, want)
	}
}
