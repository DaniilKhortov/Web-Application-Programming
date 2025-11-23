package main

import "fmt"

func main() {
	//Оголошення константи тарифу.
	const tariff float64 = 0.00005

	//Оголошення змінних для зчитування даних.
	var p float64 = 0
	var t int = 0

	//Зчитування значення з консолі.
	fmt.Printf("Enter avarage Power used (watt/s): ")

	// Для зчитування значення змінна має бути передана за адресою у Scan().
	fmt.Scan(&p)
	fmt.Printf("Enter duration (s): ")
	fmt.Scan(&t)

	//Обчислення рахунку.
	//Значення змінних можна конвертувати в інший тип.
	a := p * float64(t)
	var bill float64 = 0
	var daysTillPayment int = 7
	var report string
	var needToPay bool

	//Умовний оператор if.
	//Якщо енергія або тариф дорівнюють 0, виконується тіло if.
	//Інакше — перехід до блоку else.
	if a == 0 || tariff == 0 {
		if tariff != 0 {
			report = "No power was used!"
			needToPay = false
		} else {
			report = "Power is free of charge!"
			needToPay = false
		}
	} else {
		report = "Thank you for using our services!"
		//Обчтслення рахунку.
		bill = tariff * a
		report = fmt.Sprintf("Thank you for using our services! Payment must be done in %d days.", daysTillPayment)
		needToPay = true
	}
	//Вивід звіту у консоль.
	fmt.Printf("Good afternoon! \n%s \nTarrif(%.5f)*A(%.2f)=%.2f UAH. \nPayment requirement: %t", report, tariff, a, bill, needToPay)

}
