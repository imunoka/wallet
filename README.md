# wallet

A simple Go package for managing wallet balances. Designed for educational and demonstration purposes, especially when learning Go, testing practices, and TDD (Test-Driven Development).

---

## 📦 Features

- Create and manage a wallet balance
- Deposit and withdraw funds
- Type-safe custom `Bitcoin` type
- Includes Go examples and tests
- Implements `fmt.Stringer` for pretty printing

---

## 📥 Installation

```bash
go get github.com/imunoka/wallet
```

---

## 💡 Example

```go
package main

import (
    "fmt"
    . "github.com/imunoka/wallet"
)

func main() {
    wallet := Wallet{}
    wallet.Deposit(Bitcoin(10))
    fmt.Println("Balance:", wallet.Balance())
}
```

---

## 🧪 Testing

Run tests with the following:

```bash
go test
```

To see full test output, use the following:
```bash
go test -v
```
