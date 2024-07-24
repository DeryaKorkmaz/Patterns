package strategy

import "fmt"

// CreditCardStrategy is a concrete strategy for credit card payments
type CreditCardStrategy struct {
	CardNumber string
}

func (c *CreditCardStrategy) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Credit Card %s\n", amount, c.CardNumber)
}
