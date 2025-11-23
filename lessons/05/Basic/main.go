package main

import (
	"errors"
	"fmt"
	"math/rand"
)

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
	if param <= 0 {
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

	//Без вказування параметрів можна повернути всі значення, що імпортувалися до функції
	return
}

func main() {
	fmt.Println("E-Queue")
	//Константи для обчислення платежу
	//Тариф
	const tariff float64 = 3.74
	//Коефіцієнт К
	const k float64 = 1

	fmt.Println("\nConsumer data")
	fmt.Println("---------------------------------------------")

	//Створення клієнтів
	clients := map[string]float64{"Taras": rand.Float64() * 1000, "Vitaliy": rand.Float64() * 1000, "Daniel": rand.Float64() * 1000}

	for client, energy := range clients {
		name := fmt.Sprintf("%s`s consumed energy", client)

		//Перевірка параметрів на помилки
		if err := validateParams(energy, name); err != nil {
			fmt.Println("Error!", err)
			return
		}

		fmt.Printf("%s: %.2f kW*hour.\n", client, energy)

	}

	//Утворення списку платежів та інформації клієнтів
	clientsBill := map[string]float64{}
	clientsData := []float64{}
	for client, energyUsed := range clients {
		bill := calcBill(tariff, energyUsed, k)

		clientsBill[client] = bill
		clientsData = append(clientsData, bill)

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

	//Вивід статистичних характеристик
	fmt.Println("\nConsumer statistics")
	fmt.Println("---------------------------------------------")

	//Можна отримувати кілька значень з функцій
	sum, avg, min, max, count := sliceStat(clientsData)
	fmt.Printf("Consumers stats: \nGeneral price paid: %.2f UAH\nAvarage spendings: %.2f UAH\nMin price: %.2f UAH\nMax price: %.2f UAH\nConsumer amount: %d", sum, avg, min, max, count)

}
