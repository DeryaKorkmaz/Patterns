package strategy

// PaymentStrategy is the strategy interface
type PaymentStrategy interface {
	Pay(amount float64)
}
