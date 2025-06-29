package wallet

// ErrInsufficientFunds is returned when a user attempts to withdraw more BTC from a wallet that contained within
const ErrInsufficientFunds = WalletErr("cannot withdraw more Bitcoin than in wallet")

// WalletErr is used by the wallet package for the underlying type of all errors
type WalletErr string

// Error allows WalletErr to meet requirements of the error interface
func (e WalletErr) Error() string {
	return string(e)
}
