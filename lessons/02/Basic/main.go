package main

import "fmt"

func main() {

	//Оголошення константи tariff.
	//Константи задаються ключовим словом const і не можуть змінювати значення.
	const tariff float64 = 1.5

	//Динамічне оголошення змінної powerUsed.
	//Її тип визначається автоматично за присвоєним значенням.
	powerUsed := 1200.5

	//Статичне оголошення змінних із явним зазначенням типу.
	var bill float64 = 0
	var daysTillPayment int = 7

	//Змінні можна оголошувати без значень.
	//Тоді вони містять стандартне значення.
	var report string
	var needToPay bool

	//Обчислення рахунку.
	bill = tariff * powerUsed

	//Змінним можна задати або змінити значення.
	report = "Thank you for using our services!"
	needToPay = true

	//Вивід звіту у консоль.
	fmt.Println("Good afternoon!")
	fmt.Println(report)
	fmt.Println("Bill (UAH):")
	fmt.Println(bill)
	fmt.Println("Payment requirement:")
	fmt.Println(needToPay)
	fmt.Println("Payment must be made in (days):")
	fmt.Println(daysTillPayment)

}
