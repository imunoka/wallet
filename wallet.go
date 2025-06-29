package wallet

// Wallet represents a Bitcoin wallet with a balance
type Wallet struct {
	Balance Bitcoin 
}

// Deposit adds `amount` Bitcoin to a wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.Balance += amount 
}

// Withdraw removes `amount` Bitcoin from a wallet and returns an error if the requested amount is greater than the amount of Bitcoin in the wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.Balance {
		return ErrInsufficientFunds 
	}

	w.Balance -= amount
	return nil
}
