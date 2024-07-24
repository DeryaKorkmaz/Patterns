package strategy

import "fmt"

// BitcoinStrategy is a concrete strategy for Bitcoin payments
type BitcoinStrategy struct {
	WalletAddress string
}

func (b *BitcoinStrategy) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Bitcoin wallet %s\n", amount, b.WalletAddress)
}
