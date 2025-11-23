package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Функція видалення елементу зрізу.
func removeAtIndex(slice []float64, index int) []float64 {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

// Функція злиття зрізів.
func mergeMaps(maps ...map[string]float64) map[string]float64 {
	merged := make(map[string]float64, 10)
	for _, m := range maps {
		for k, v := range m {
			merged[k] = v
		}
	}
	return merged
}

// Функція групування map за двома категоріями(ключами).
func groupByRegionAndCategory(data map[string]map[string]float64) {
	fmt.Println("\nGrouping by region and category:")
	for region, categories := range data {
		fmt.Printf("Region: %s\n", region)
		for category, value := range categories {
			fmt.Printf(" Category: %-10s || %.2f UAH/kW*hour\n", category, value)
		}
	}
}

func main() {
	fmt.Println("E-Queue")

	// Ініціалізація двовимірного масиву.
	grid := [][]float64{
		{12.5, 14.3, 11.7},
		{10.2, 13.8, 15.4},
		{9.9, 12.1, 14.8},
	}
	//Вивід споживання у вигляді таблиці.
	fmt.Println("\nEnergy Consumption by table:")
	for i, row := range grid {
		fmt.Printf("Day %d: %v\n", i+1, row)
	}

	//Створення зрізу monitoringData через make.
	monitoringData := make([]float64, 0, 10)

	//Заповнення зрізу monitoringData через цикл for та append.
	for i := 0; i < 6; i++ {
		monitoringData = append(monitoringData, rand.Float64()*10)
	}

	//Вивід зрізу monitoringData.
	fmt.Printf("\n Monitoring data: %.2f", monitoringData)

	//Видалення елемента зі зрізу.
	monitoringData = removeAtIndex(monitoringData, 2)
	fmt.Printf("Monitoring data after element [2] deleted: %.2f", monitoringData)

	//Сортування елементів зрізу за зростанням.
	sort.Float64s(monitoringData)
	fmt.Printf("\nMonitoring data after sort: %.2f", monitoringData)

	//Сортування елементів зрізу за спаданням.
	sort.Sort(sort.Reverse(sort.Float64Slice(monitoringData)))
	fmt.Printf("Monitoring data after reverse order sort: %.2f", monitoringData)

	//Підготовка 3-х map для злиття.
	mapA := map[string]float64{"Standart": 1.5, "Eco": 1.2}
	mapB := map[string]float64{"Premium": 3.8, "Comercial": 2.5}
	mapC := map[string]float64{"Student": 1.0}

	//Злиття 3-х map в один.
	merged := mergeMaps(mapA, mapB, mapC)

	//Вивід результату злиття.
	fmt.Println("\nTariff merge:")
	for k, v := range merged {
		fmt.Printf("%s: %.2f UAH/kW*hour.\n", k, v)
	}

	//Створення map для групування за ключами.
	regionalData := map[string]map[string]float64{
		"Kyiv": {
			"Standart":  220.5,
			"Comercial": 380.0,
		},
		"Lviv": {
			"Standart": 210.3,
			"Eco":      150.8,
		},
		"Odessa": {
			"Premium": 400.6,
			"Eco":     160.2,
		},
	}

	//Групування map за категоріями.
	groupByRegionAndCategory(regionalData)

	//Оптимізація пам’яті через pre-allocation
	preAllocated := make([]float64, 0, 50)
	for i := 0; i < 30; i++ {
		preAllocated = append(preAllocated, rand.Float64()*100)
	}
	fmt.Printf("\nPre-allocated slice: len=%d cap=%d\n", len(preAllocated), cap(preAllocated))
}
