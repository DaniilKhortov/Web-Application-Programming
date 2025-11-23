package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("E-Queue")

	//Створення й заповнення масиву через  цикл for.
	//Для заповнення використовується rand.Float64() з пакету math/rand.
	var hourlyConsumption [24]float64
	for i := 0; i < 24; i++ {
		hourlyConsumption[i] = rand.Float64() * 100
	}

	//Вивід споживання по-годинно.
	fmt.Println("\nEnergy Consumption by hour:")
	for i, v := range hourlyConsumption {
		fmt.Printf("%02d:00: %.2f  \n", i, v)
	}

	//Створення зрізу monitoringData з довжиною 3 та ємністю 6.
	monitoringData := make([]float64, 3, 6)
	monitoringData[0], monitoringData[1], monitoringData[2] = 2.4, 3.1, 4.7

	//Вивід зрізу monitoringData, його довжини та ємності.
	fmt.Println("\nMonitoring data:", monitoringData)
	fmt.Printf("Array length: %d\n", len(monitoringData))
	fmt.Printf("Array capacity: %d\n", cap(monitoringData))

	//Вивід зрізу monitoringData, його довжини та ємності після додавання 3-х елеменів.
	monitoringData = append(monitoringData, 5.2, 6.8, 4.4)
	fmt.Println("Adding new data:", monitoringData)
	fmt.Printf("Array length: %d\n", len(monitoringData))
	fmt.Printf("Array capacity: %d\n", cap(monitoringData))

	//Створення підмасиву.
	subset := monitoringData[1:4]
	fmt.Println("\nSlice:", subset)

	//Копіювання зрізу.
	copySlice := make([]float64, len(monitoringData))
	copy(copySlice, monitoringData)
	//Змінюємо перший елемент, щоб перевірити незалежність утвореного зрізу.
	copySlice[0] = 999

	//Порівняння зрізів.
	fmt.Println("\nSlice copy:", copySlice)
	fmt.Println("Original array:", monitoringData)

	//Створення двовимірного map користувачів.
	clientData := map[string][]float64{
		"Alice": {2.4, 3.1, 4.7},
		"Bob":   {5.2, 6.8},
	}
	//Додавання користувача до map.
	clientData["Charlie"] = []float64{3.3, 3.8, 4.2}

	//Вивід map користувачів.
	fmt.Println("\nClient data:")
	for name, data := range clientData {
		fmt.Printf("%s: %v\n", name, data)
	}

	//Створення map тарифів.
	//Вказано альтернативний варіант його заповнення.
	tariffs := map[string]float64{
		"Standart":  1.5,
		"Comercial": 2.3,
		"Premium":   3.8,
		"Eco":       1.1,
	}

	//Вивід елементів map з категоріями та значеннями.
	fmt.Println("\nTariffs by category:")
	for category, price := range tariffs {
		fmt.Printf("%s: %.2f UAH/kW*hour.\n", category, price)

	}

	//Перевірка на існування категорії.
	//Категорія Premium існує, а отже має вивестися результат пошуку.
	categoryToCheck := "Eco"
	if val, exists := tariffs[categoryToCheck]; exists {
		fmt.Printf("\nCategory %s found! Tariff: %.2f UAH/kW*hour.\n", categoryToCheck, val)

	} else {
		fmt.Printf("\nCategory %s not found!", categoryToCheck)
	}
}
