package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// Функція generateClients створює масив клієнтів
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

// Функція makeBillTracker створює замикання. Використовується для акумулювання платежу
func makeBillTracker() func(float64) float64 {
	var total float64
	return func(amount float64) float64 {
		total += amount
		return total
	}
}

// Функція makeDiscountFunc через замикання призначає знижку
func makeDiscountFunc(rate float64) func(float64) float64 {
	return func(bill float64) float64 {
		discounted := bill * (1 - rate)
		if discounted < 0 {
			return 0
		}
		return discounted
	}
}

// Функція validateParams перевіряє артибути об'єктів
func validateParams(param float64, name string) error {
	if param < 0 {
		message := fmt.Sprintf("Argument %s must be higher than 0! Current value: %v", name, param)
		return errors.New(message)
	}
	return nil
}

// Функція calcBill обчислює рахунок
func calcBill(tariff, client, k float64) float64 {
	return tariff * client * k
}

// Функція processBills використовується при виведені рахунків
// Використовує іншу функцію за аргумент
func processBills(bills map[string]float64, handler func(string, float64)) {
	for name, bill := range bills {
		handler(name, bill)
	}
}

// Функція printHierarchy виводить дерево з клієнтів
func printHierarchy(tree map[string][]string, level int, name string) {
	prefix := ""
	for i := 0; i < level; i++ {
		prefix += "  "
	}
	fmt.Println(prefix + name)
	for _, child := range tree[name] {
		printHierarchy(tree, level+1, child)
	}
}

// Функція sliceStat обчислює статистичні характеристики масиву
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

// Функція safeRun використовується для обробки виключень
// defer переносить спрацювання виключення функції на початок
func safeRun(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\n Panic caught:", r)
		}
	}()
	fn()
}

// Функція mainLogic використоаується для управління програмою
func mainLogic() {
	defer fmt.Println("\nProgram finished successfully!")

	fmt.Println("E-Queue")
	//Ініціалізація констант для підрахунку платежів
	//Тариф
	const tariff float64 = 3.74
	//Коефіцієнт К
	const k float64 = 1
	//Знижка
	const discountRate float64 = 0.05

	//Імена клієнтів
	clientNames := []string{"Taras", "Vitaliy", "Daniel"}

	//Утворення клієнтів
	clients, err := generateClients(clientNames)

	//Тестування recover
	if err != nil {
		panic(err)
	}

	//Застосування знижки
	applyDiscount := makeDiscountFunc(discountRate)

	fmt.Println("\nConsumer data")
	fmt.Println("---------------------------------------------")

	//Вивід черги з клієнтів та спожитої енергії
	for client, energy := range clients {
		name := fmt.Sprintf("%s`s consumed energy", client)
		if err := validateParams(energy, name); err != nil {
			panic(err)
		}
		fmt.Printf("%s: %.2f kW*hour.\n", client, energy)
	}

	//Утворення акумулятора платежу
	billTracker := makeBillTracker()

	//Підрахунок платежу клієнтів
	clientsBill := map[string]float64{}
	for client, energyUsed := range clients {
		bill := calcBill(tariff, energyUsed, k)
		bill = applyDiscount(bill)
		clientsBill[client] = bill
		total := billTracker(bill)
		fmt.Printf("General price paid after %s | total: %.2f UAH\n", client, total)
	}

	fmt.Println("\nConsumer bill data")
	fmt.Println("---------------------------------------------")

	//Вивід платежів
	processBills(clientsBill, func(name string, bill float64) {
		fmt.Printf("%s: %.2f UAH.\n", name, bill)
	})

	//Обчислення статистичних харектиристик черги клієнтів як вибірки
	clientsData := []float64{}
	for _, bill := range clientsBill {
		clientsData = append(clientsData, bill)
	}
	sum, avg, min, max, count := sliceStat(clientsData)

	//Вивід статистичних харектиристик вибірки
	fmt.Println("\nConsumer statistics")
	fmt.Println("---------------------------------------------")
	fmt.Printf("Total paid: %.2f UAH\nAverage: %.2f UAH\nMin: %.2f\nMax: %.2f UAH\nCount: %d\n",
		sum, avg, min, max, count)

	//Вивід ієрархії компанії як дерева
	fmt.Println("\nCompany hierarchy:")
	fmt.Println("---------------------------------------------")
	hierarchy := map[string][]string{
		"Head Office": {"Taras", "Vitaliy"},
		"Vitaliy":     {"Daniel"},
	}

	printHierarchy(hierarchy, 0, "Head Office")
}

// Виклик safeRun, що керує обрахунком платежів, характеристик та виводить склад компанії
func main() {
	safeRun(mainLogic)
}
