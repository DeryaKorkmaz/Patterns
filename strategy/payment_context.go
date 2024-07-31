package strategy

// PaymentContext is the context for payment strategies
type PaymentContext struct {
	strategy PaymentStrategy
}

// / SetStrategy, ödeme stratejisini değiştirmek için kullanılır.
func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

// Pay, belirlenen ödeme stratejisini kullanarak ödemeyi gerçekleştirir.
func (p *PaymentContext) Pay(amount float64) {
	p.strategy.Pay(amount)
}
