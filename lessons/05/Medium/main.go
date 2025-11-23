package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// Функція generateClients створює масив клієнтів
// На вхід функція приймає масив імен
// На вихід отримуємо масив клієнтів або повідомлення помилки
func generateClients(names []string) (map[string]float64, error) {
	if len(names) == 0 {
		return nil, errors.New("no clients provided")
	}

	clients := make(map[string]float64)
	for _, n := range names {
		clients[n] = rand.Float64() * 1000
	}
	return clients, nil
}

// Функція applyDiscount застосовує знижку на платіж
// На вхід функція приймає платіж та знижку
// Поточна функція нічого не повертає, проте модифікує значення значення вказівника
func applyDiscount(bill *float64, discount float64) {
	if discount > 0 && discount < 1000 {
		*bill = *bill - discount
		if *bill < 0 {
			*bill = 0
		}

	}

}

// Функція calcBill рахує платіж клієнта
// На вхід функція приймає tariff (тариф), client (спожита енергія клієнтом), k (коефіцієнт К)
// На вихід отримуємо обчислений платіж з типом float64
func calcBill(tariff, client, k float64) float64 {
	return tariff * client * k
}

// Функція validateParams перевіряє параметри
// На вхід функція приймає param (параметр), name (назва параметру)
// На вихід отримуємо помилку, яка має оброблятись або nil (нічого)
func validateParams(param float64, name string) error {
	if param < 0 {
		message := fmt.Sprintf("Argument %s must be higher than 0! Current value: %v", name, param)
		return errors.New(message)
	}
	return nil
}

// Функція sliceStat обчислює статистичні характеристики вибірки
// На вхід функція приймає arr (масив, що буде використана як статистична вибірка)
// На вихід отримуємо характеристики:
// sum - загальна сума
// avg - середньостатистичне значення
// min - мінімальне значення
// max - максимальне значення
// count - розмір вибірки
func sliceStat(arr []float64) (sum, avg, min, max float64, count int) {
	if len(arr) == 0 {
		return 0, 0, 0, 0, 0
	}

	min, max = arr[0], arr[0]
	for _, v := range arr {
		sum += v
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	count = len(arr)
	avg = sum / float64(count)
	return
}

// Функція sumAll сумує вибірку значень.
// Ця функція є варіадичною. [...] означає, що змінних може бути на вході 1 та більше
// На вхід функція приймає вибірку values.
// На вихід отримуємо суму
func sumAll(values ...float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum
}

// Функція processBills використовується для запуску іншої функції
// На вхід функція приймає платежі та функцію handler
// Функція нічого не повертає. Проте, ітеративно викликає іншу функцію.
// Це дозволяє спростити користування імпортованої функції
func processBills(bills map[string]float64, handler func(string, float64)) {
	for name, bill := range bills {
		handler(name, bill)
	}
}

func main() {
	defer fmt.Println("\nProgram finished successfully!")

	fmt.Println("E-Queue")

	//Оголошення констант
	const tariff float64 = 3.74
	const k float64 = 1
	const discount float64 = 750

	//Масив імен клієнтів
	clientNames := []string{"Taras", "Vitaliy", "Daniel"}

	//Створення клієнтів
	clients, err := generateClients(clientNames)

	//Перевірка на очікувану поведінку програми при утворені клієнтів
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nConsumer data")
	fmt.Println("---------------------------------------------")

	//Вивід списку клієнтів
	for client, energy := range clients {
		name := fmt.Sprintf("%s`s consumed energy", client)

		//Перевірка правильності параметрів клієнтів
		if err := validateParams(energy, name); err != nil {
			fmt.Println("Error!", err)
			return
		}
		fmt.Printf("%s: %.2f kW*hour.\n", client, energy)
	}

	//Утворення списку платежів клієнтів
	clientsBill := map[string]float64{}

	//Обчислення платежів клієнтів
	for client, energyUsed := range clients {
		bill := calcBill(tariff, energyUsed, k)

		applyDiscount(&bill, discount)

		clientsBill[client] = bill
	}

	fmt.Println("\nConsumer bill data")
	fmt.Println("---------------------------------------------")

	//Вивід списку клієнтів з платежами
	for client, bill := range clientsBill {
		fmt.Printf("%s: %.2f UAH.\n", client, bill)
		name := fmt.Sprintf("%s`s energy price", client)
		if err := validateParams(bill, name); err != nil {
			fmt.Println("Error!", err)
			return
		}

	}

	//Вивід списку клієнтів з платежами через замикання
	processBills(clientsBill, func(name string, bill float64) {
		fmt.Printf("%s: %.2f UAH.\n", name, bill)
	})

	//Обчислення та вивід статистичних характеристик
	clientsData := []float64{}
	for _, bill := range clientsBill {
		clientsData = append(clientsData, bill)
	}
	sum, avg, min, max, count := sliceStat(clientsData)

	//Варіадична функція
	totalByVar := sumAll(clientsData...)
	fmt.Println("\nConsumer statistics")
	fmt.Println("---------------------------------------------")
	fmt.Printf("Consumers stats: \nGeneral price paid: %.2f UAH(%.2f UAH with VarFunc)\nAvarage spendings: %.2f UAH\nMin price: %.2f UAH\nMax price: %.2f UAH\nConsumer amount: %d", sum, totalByVar, avg, min, max, count)

}
