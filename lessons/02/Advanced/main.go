package main

import "fmt"

func main() {
	//Оголошення константи тарифу.
	const tariff float64 = 0.05

	//Оголошення змінних.
	var daysTillPayment int = 7
	const k float64 = 1
	var report string
	var needToPay bool

	//Оголошення змінних для зчитування даних.
	var power float64
	var duration int

	//Зчитування значення з консолі.
	//Помилку можна зберегти у змінну err1.
	//Для перевірки використовують порівняння err1 з nil.
	fmt.Printf("Enter avarage Power used (W): ")
	_, err1 := fmt.Scan(&power)
	if err1 != nil || power < 0 {
		fmt.Println("Error! Invalid power value.")
		return
	}

	fmt.Printf("Enter duration (seconds): ")
	_, err2 := fmt.Scan(&duration)
	if err2 != nil || duration <= 0 {
		fmt.Println("Error! Invalid duration value.")
		return
	}

	// Енергія в форматі кВт·год.
	power = power / 1000
	duration = duration / 3600
	energy := power * float64(duration)

	//Обчислення рахунку.
	bill := k * tariff * energy

	//Перевірка на використання енергії та тариф.
	if energy == 0 || tariff == 0 {
		if tariff != 0 {
			report = "No power was used!"
			needToPay = false
		} else {
			report = "Power is free of charge!"
			needToPay = false
		}

	} else if bill < 0.1 {
		report = fmt.Sprintf("Your bill is negligible (%.4f UAH). Payment not required.", bill)
		needToPay = false

	} else {

		report = fmt.Sprintf("Thank you for using our services! Payment must be done in %d days.", daysTillPayment)
		needToPay = true
	}

	//Формування звіту за допомогою Sprintf().
	//Sprintf() створює форматований рядок, не виводячи його у консоль.
	finalReport := fmt.Sprintf(
		"\n--- Energy Usage Report ---\n"+
			"Power used: %.2f kW\n"+
			"Duration: %d h\n"+
			"Energy total: %.2f J\n"+
			"Tariff: %v UAH/kW\n"+
			"Coefficient k: %.1f\n"+
			"Bill: %.2f UAH\n"+
			"Need to pay: %t\n"+
			"Status: %s\n",
		power, duration, energy, tariff, k, bill, needToPay, report,
	)

	//Вивід звіту у консоль.
	fmt.Println(finalReport)
}
