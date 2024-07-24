package strategy

import "fmt"

// PayPalStrategy is a concrete strategy for PayPal payments
type PayPalStrategy struct {
	Email string
}

func (p *PayPalStrategy) Pay(amount float64) {
	fmt.Printf("Paid %.2f using PayPal %s\n", amount, p.Email)
}
