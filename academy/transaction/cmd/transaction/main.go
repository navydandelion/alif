package main

import "awesomeProject/internal/transaction"

func main() {
	sender := "Sherali"
	receiver := "Alisher"
	amount := int(500_000)
	commission := amount * 10 / 1000
	isAzo := true
	transaction.Printer(sender, receiver, amount, commission, isAzo)
}
