package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func (t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		
		newBalance := Bitcoin(10)
		assertBalance(t, wallet, newBalance)
	})

	t.Run("withdraw", func (t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		newBalance := Bitcoin(10)
		assertBalance(t, wallet, newBalance)
	})

	t.Run("withdraw insufficient funds", func (t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(30))
		
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(20))
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	balance := wallet.Balance

	if balance != want {
		t.Errorf("got %s, want %s", balance, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("wanted an error but did not get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
