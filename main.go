package main

import (
	"Patterns/adapter"
	"Patterns/strategy"
	"fmt"
)

func main() {

	fmt.Println("\nAdapter Pattern Example")
	oldGameConsole := &adapter.OldGameConsole{}
	newTelevision := &adapter.NewTelevision{}

	// Using adapter to connect old game console to new television
	adapter := &adapter.AVtoHDMIAdapter{AVDevice: oldGameConsole}
	adapter.InputHDMI("some data")

	// Direct connection to show how it works without adapter
	newTelevision.InputHDMI("Direct HDMI signal")

	//*---*****----------------------------------------------------------------

	fmt.Println("Strategy Pattern Example")
	strategyContext := &strategy.PaymentContext{}

	// Credit Card Payment
	creditCard := &strategy.CreditCardStrategy{CardNumber: "1234-5678-9012-3456"}
	strategyContext.SetStrategy(creditCard)
	strategyContext.Pay(100.00)

	// PayPal Payment
	payPal := &strategy.PayPalStrategy{Email: "user@example.com"}
	strategyContext.SetStrategy(payPal)
	strategyContext.Pay(200.00)

	// Bitcoin Payment
	bitcoin := &strategy.BitcoinStrategy{WalletAddress: "1BitcoinAddress"}
	strategyContext.SetStrategy(bitcoin)
	strategyContext.Pay(300.00)

}
