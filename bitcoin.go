package wallet

import "fmt"

// Bitcoin represents the unit of currency stored within a Wallet
type Bitcoin int

// String allows Bitcoin to implement the Stringer interface
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
