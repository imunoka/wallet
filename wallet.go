package wallet

type Wallet struct {
	Balance Bitcoin 
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.Balance += amount 
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.Balance {
		return ErrInsufficientFunds 
	}

	w.Balance -= amount
	return nil
}
