package transaction

import "fmt"

func Printer(sender, receiver string, amount, commission int, isAzo bool) {
	fmt.Printf("========= Чek =========\n")
	fmt.Printf("Отправитель:%s\n", sender)
	fmt.Printf("Полуачатель:%s\n", receiver)
	fmt.Printf("Сумма:%d\n", amount)
	fmt.Printf("Статус Аъзо:%t\n", isAzo)
	fmt.Printf("Комиссия: %d\n", commission)
	fmt.Printf("Итого: %d\n", amount+commission)
	fmt.Printf("========================\n")
}
