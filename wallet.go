package wallet

type Wallet struct {
	Balance Bitcoin 
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.Balance += amount 
}
