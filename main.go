/*
	func main() {
		var name1 = "Alif"
		//var name1 string = "Alif"
		name := "Alif"
		year := 1986
		float := 243.90
		println(name)
		println(year)
		println(float)
		println(name1)
		//var name string = "Alif"
		//var amount float64
		//	println("Hello World")
		fmt.Println("welcome to Alif academy")
	}
*/
/*
func calculateRate(currency string) float64 {
	return 13500
}

*/
/*
func main() {
	rateEuro := 14079.98
	rateUsd := 12222.27
	rateRub := 155.51
	var choice int
	fmt.Println("1 - USD")
	fmt.Println("2 - EUR")
	fmt.Println("3 - RUB")
	fmt.Scan(&choice)

	var amount float64
	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)
	if choice == 2 {
		result := rateEuro * amount
		fmt.Printf("%.2f EUR = %.2f UZS\n", amount, result)
	} else if choice == 3 {
		result := rateRub * amount
		fmt.Printf("%.2f RUB = %.2f UZS\n", amount, result)
	} else if choice == 1 {
		result := rateUsd * amount
		fmt.Printf("%.2f USD = %.2f UZS\n", amount, result)
	} else {
		fmt.Println("можно выбрать только 1, 2, 3")
	}

}
*/
//var amount float64
//fmt.Print("введите сумму в EUR : ")
//fmt.Scan(&amount)
//result := rateEuro * amount
//fmt.Printf("%.2f EUR = %.2f UZS\n", amount, result)

// -------------SECOND LESSON---------
// ///------CLASSWOORK------FIRST TASK----------
//package main
//
//import "fmt"
//
//func main() {
//	var a int = 234
//	var b int8 = 126
//	var c int64 = 12343
//	d := "Hello teacher"
//	e := true
//	var f uint = 12345678
//	fmt.Println(a, b, c, d, e, f)
//	fmt.Println(c)
//	fmt.Printf("Тип переменной c: %T\n", c)
//}

// -----CLASSWOORK----SECOND TASK--------------
// package main
//
// import "fmt"
//
//	func main() {
//		// int 8 := 120
//
//		//var a, b int64
//		//fmt.Print("Введите первое число: ")
//		//fmt.Print("Введите второе число: ")
//		//fmt.Scan(&a, &b)
//		var a int64
//		fmt.Print("Введите первое число: ")
//		fmt.Scan(&a)
//		var b int64
//		fmt.Print("Введите второе число: ")
//		fmt.Scan(&b)
//		if a != b {
//			fmt.Println("a не равно b")
//			if a > b {
//				fmt.Println("a больше b на ", a-b)
//				fmt.Println("a + b = ", a+b)
//				fmt.Println("a - b = ", a-b)
//				fmt.Println("a : b = ", a/b)
//				fmt.Println("a * b = ", a*b)
//				fmt.Println("при делении а на b остаток получается", a%b)
//			} else if b > a {
//				fmt.Println("b больше a на ", b-a)
//				fmt.Println("a + b = ", a+b)
//				fmt.Println("b : a  = ", b/a)
//				fmt.Println("a * b = ", a*b)
//				fmt.Println("при делении b на a остаток получается", b%a)
//			} else if a == b {
//				fmt.Println("a равен b")
//				fmt.Println("a - b =", a-b)
//				fmt.Println("a + b = ", a+b)
//				fmt.Println("a : b = ", a/b)
//				fmt.Println("a * b = ", a*b)
//				fmt.Println("при делении а на b остаток получается", a%b)
//			}
//		}
//
// }
// ------HOMEWORK------ FIRST TASK------
// package main
//
// import "fmt"
//
//	func main() {
//		var a int64
//		fmt.Print("Введите сумму оплаты: ")
//		fmt.Scan(&a)
//		c := (a * 22) / 1000
//		fmt.Println("Кешбэк за сервисы (2.2%): ", c)
//		fmt.Println("Кешбек за рассрочку (0.1%): ", (a*1)/1000)
//	}
//
// ------HOMEWORK------ SECOND TASK------
//package main
//
//import "fmt"
//
//func main() {
//	var a int64
//	fmt.Print("Введите сумму перевода: ", a)
//}
//-------FOURTH LESSON---------
