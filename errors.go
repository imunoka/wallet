package wallet

const ErrInsufficientFunds = WalletErr("cannot withdraw more Bitcoin than in wallet")

type WalletErr string

func (e WalletErr) Error() string {
	return string(e)
}
