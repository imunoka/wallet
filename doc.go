/*
Package wallet provides an API for creating and managing mock Bitcoin wallets


To get the most out of wallet, use dot import syntax and create/manage wallets with the following:

	package main

	import (
		"fmt"
		. "github.com/imunoka/wallet"
	)

	func main() {
		wallet := Wallet{Bitcoin(100)}

		wallet.Deposit(Bitcoin(15))
		
		err := wallet.Withdraw(Bitcoin(200))
		fmt.Println("Error:", err)

		balance := wallet.Balance
		fmt.Println("Balance:", balance)
	}
*/
package wallet
