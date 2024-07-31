package main

import (
	"Patterns/adapter"
	"Patterns/strategy"
	"fmt"
	"os"
)

func main() {

	// Örneğin, eski bir video oyun konsolunu (AV çıkışlı) yeni bir televizyona (HDMI girişli) bağlamak istiyoruz. Bunun için bir adaptör kullanıyoruz.
	//Adaptör, AV sinyalini HDMI sinyaline dönüştürür ve televizyona iletir. Ayrıca, doğrudan HDMI sinyali göndermenin nasıl çalıştığını da gösterir.

	fmt.Println("\nAdapter Pattern Example")
	oldGameConsole := &adapter.OldGameConsole{}
	newTelevision := &adapter.NewTelevision{}
	//newGameConsole := &adapter.NewGameConsole{}

	// Using adapter to connect old game console to new television
	adapter := &adapter.AVtoHDMIAdapter{AVDevice: oldGameConsole}
	adapter.InputHDMI("some data")

	// Direct connection to show how it works without adapter
	newTelevision.InputHDMI("Direct HDMI signal")

	os.Exit(0)
	//*---*****----------------------------------------------------------------

	//Örneğin, PaymentStrategy arayüzü ile farklı ödeme stratejileri (CreditCardStrategy, PayPalStrategy, CashStrategy) tanımladık.
	//PaymentContext sınıfı, bu stratejileri dinamik olarak değiştirebilmemizi ve belirli bir stratejiye göre ödeme yapabilmemizi sağlar.
	// Bu sayede, ödeme yöntemi değiştiğinde sadece strateji sınıfını değiştirerek esnek bir yapı elde etmiş oluyoruz.
	fmt.Println("Strategy Pattern Example")
	strategyContext := &strategy.PaymentContext{}

	// Credit Card Payment
	creditCard := &strategy.CreditCardStrategy{CardNumber: "1234-5678-9012-3456"}
	strategyContext.SetStrategy(creditCard)
	strategyContext.Pay(100.00)

	// PayPal Payment
	payPal := &strategy.PayPalStrategy{Email: "user@outlook.com"}
	strategyContext.SetStrategy(payPal)
	strategyContext.Pay(200.00)

	// Bitcoin Payment
	bitcoin := &strategy.BitcoinStrategy{WalletAddress: "1BitcoinAddress"}
	strategyContext.SetStrategy(bitcoin)
	strategyContext.Pay(300.00)

}
