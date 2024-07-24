package strategy

// PaymentContext is the context for payment strategies
type PaymentContext struct {
	strategy PaymentStrategy
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentContext) Pay(amount float64) {
	p.strategy.Pay(amount)
}
