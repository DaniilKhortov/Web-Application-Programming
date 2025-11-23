package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("E-Queue")

	//Масив зберігання споживання енергії за 24 години.
	var hourlyConsumption [24]float64

	//Заповнення масиву через  цикл for.
	for i := 0; i < 24; i++ {
		//Для заповнення використовується rand.Float64() з пакету math/rand.
		//rand.Float64() повертає випадкове число на проміжку [0, 1).
		hourlyConsumption[i] = rand.Float64() * 100
	}

	//Вивід споживання по-годинно.
	fmt.Println("\nEnergy Consumption by hour:")
	for i, v := range hourlyConsumption {
		fmt.Printf("%02d:00: %.2f  \n", i, v)
	}

	//Створення зрізу monitoringData.
	monitoringData := []float64{1.2, 2.5, 3.7, 4.1, 3.9, 2.8}

	//Вивід зрізу monitoringData, його довжини та ємності.
	fmt.Println("\nMonitoring data:", monitoringData)
	fmt.Printf("Array length: %d\n", len(monitoringData))
	fmt.Printf("Array capacity: %d\n", cap(monitoringData))

	//Створення map тарифів.
	tariffs := make(map[string]float64)

	//Додавання елементів до map.
	tariffs["Standart"] = 1.5
	tariffs["Comercial"] = 2.3
	tariffs["Premium"] = 3.8

	//Вивід елементів map з категоріями та значеннями.
	fmt.Println("\nTariffs by category:")
	for category, price := range tariffs {
		fmt.Printf("%s: %.2f UAH/kW*hour.\n", category, price)
	}

	//Перевірка на існування категорії.
	//Категорія Premium існує, а отже має вивестися результат пошуку.
	categoryToCheck := "Premium"
	value, exists := tariffs[categoryToCheck]
	fmt.Printf("\nCategory check: \"%s\":\n", categoryToCheck)
	if exists {
		fmt.Printf("\nCategory found! Tariff: %.2f UAH/kW*hour.\n", value)
	} else {
		fmt.Println("Category not found!")
	}

	//Перевірка на існування категорії.
	//Категорія Eco не існує.
	// Результатом має бути повідомлення про невдалий пошук.
	categoryToCheck = "Eco"
	value, exists = tariffs[categoryToCheck]
	fmt.Printf("\nCategory check: \"%s\":\n", categoryToCheck)
	if exists {
		fmt.Printf("\nCategory found! Tariff: %.2f UAH/kW*hour.\n", value)
	} else {
		fmt.Println("Category not found!")
	}
}
