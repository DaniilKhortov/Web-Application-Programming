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

// Функція makeBillTracker акумулює загальну суму платежів
// На вхід функція нічого не потребує
// На вихід отримуємо нову функцію
func makeBillTracker() func(float64) float64 {
	var total float64
	return func(amount float64) float64 {
		total += amount
		return total
	}
}

// Функція makeDiscountFunc ініціалізовує функцію знижки на платіж
// На вхід функція приймає платіж та знижку
// Поточна функція повертає функціюю застосування знижки
func makeDiscountFunc(rate float64) func(float64) float64 {
	return func(bill float64) float64 {
		discounted := bill * (1 - rate)
		if discounted < 0 {
			return 0
		}
		return discounted
	}
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

// Функція calcBill рахує платіж клієнта
// На вхід функція приймає tariff (тариф), client (спожита енергія клієнтом), k (коефіцієнт К)
// На вихід отримуємо обчислений платіж з типом float64
func calcBill(tariff, client, k float64) float64 {
	return tariff * client * k
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

// Функція printHierarchy використовується для виведення складу працівників
// На вхід функція приймає дерево, рівень за ієрархією та назву
// Функція нічого не повертає
// Функція рекурсивна, оскільки викликає саму себе
// Хоча підхід є ефективним, проте в подальших роботах може бути замінена зрозумілішим ітеративним методом
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

// Функція safeRun використовується для обробки виключень
// defer відкладає спрацювання виключення функції до закінчення роботи основної функції
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

	//Оголошення констант
	const tariff float64 = 3.74
	const k float64 = 1
	const discountRate float64 = 0.05

	//Масив імен клієнтів
	clientNames := []string{"Taras", "Vitaliy", "Daniel"}

	//Створення клієнтів
	clients, err := generateClients(clientNames)

	//Перевірка на очікувану поведінку програми при утворені клієнтів
	if err != nil {
		panic(err)
	}

	//Ініціалізація функція застосування знижок
	applyDiscount := makeDiscountFunc(discountRate)

	fmt.Println("\nConsumer data")
	fmt.Println("---------------------------------------------")

	//Вивід списку клієнтів
	for client, energy := range clients {
		name := fmt.Sprintf("%s`s consumed energy", client)

		//Перевірка правильності параметрів клієнтів
		if err := validateParams(energy, name); err != nil {
			panic(err)
		}
		fmt.Printf("%s: %.2f kW*hour.\n", client, energy)
	}

	//Утворення лічильника платежу
	billTracker := makeBillTracker()

	//Вивід списку клієнтів з платежами
	clientsBill := map[string]float64{}
	for client, energyUsed := range clients {
		bill := calcBill(tariff, energyUsed, k)
		bill = applyDiscount(bill)
		clientsBill[client] = bill
		total := billTracker(bill)
		fmt.Printf("General price paid after %s : total: %.2f UAH\n", client, total)
	}

	//Вивід списку клієнтів з платежами через замикання
	fmt.Println("\nConsumer bill data")
	fmt.Println("---------------------------------------------")
	processBills(clientsBill, func(name string, bill float64) {
		fmt.Printf("%s: %.2f UAH.\n", name, bill)
	})

	//Обчислення та вивід статистичних характеристик
	clientsData := []float64{}
	for _, bill := range clientsBill {
		clientsData = append(clientsData, bill)
	}
	sum, avg, min, max, count := sliceStat(clientsData)

	fmt.Println("\nConsumer statistics")
	fmt.Println("---------------------------------------------")
	fmt.Printf("Total paid: %.2f UAH\nAverage: %.2f UAH\nMin: %.2f\nMax: %.2f UAH\nCount: %d\n",
		sum, avg, min, max, count)

	//Вивід ієрархії компанії як дерево
	fmt.Println("\nCompany hierarchy:")
	fmt.Println("---------------------------------------------")
	hierarchy := map[string][]string{
		"Head Office": {"Taras", "Vitaliy"},
		"Vitaliy":     {"Daniel"},
	}
	printHierarchy(hierarchy, 0, "Head Office")

	//Варіадична функція
	totalByVar := sumAll(clientsData...)
	fmt.Println("\nConsumer statistics")
	fmt.Println("---------------------------------------------")
	fmt.Printf("VarFunc result of general bill: %.2f UAH", totalByVar)

}

func main() {
	safeRun(mainLogic)
}
