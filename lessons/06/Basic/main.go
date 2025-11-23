package main

import (
	"fmt"
)

// Функція registerClient додає нового клієнта у чергу через вказівник на зріз
// Оператор * біля типу даних вказує, що очікується значення вказівника
func registerClient(queue *[]string, count *int, name string) {

	//Перевірка на наявність адрес зрізу черги та кількості клієнтів
	if queue == nil || count == nil {
		fmt.Println("Pointer error! Nil pointer detected!")
		return
	}

	//Додавання клієнта у зріз
	*queue = append(*queue, name)

	//Оновлення кількості клієнтів
	*count = *count + 1
	fmt.Printf("Client %s registered. Clients in queue: %d\n", name, *count)
}

// Функція serveClient обслуговує клієнтів
func serveClient(queue *[]string, count *int) {

	//Перевірка на наявність адрес зрізу черги та кількості клієнтів
	if queue == nil || count == nil {
		fmt.Println("Pointer error! Nil pointer detected!")
		return
	}

	//Перевірка на наповнення черги
	if len(*queue) == 0 {
		fmt.Println("Queue is empty")
		return
	}

	//Зберіаємо першого клієнта в черзі для виводу
	served := (*queue)[0]
	//Скорочуємо зріз. Перезаписуємо значення queue від другого клієнта в минулій черзі
	*queue = (*queue)[1:]
	//Оновлення кількості клієнтів
	*count = *count - 1
	fmt.Printf("Serving client %s ... Clients left: %d\n", served, *count)
}

func main() {
	fmt.Println("E-Queue - Register desk")

	//Створюємо змінну підрахунку клієнтів
	var totalClients int = 0

	//pTotal - вказівник на змінну totalClients
	//Вказівники зберігають адресу пам`ятті на значення
	//Для отримання адреси пам`ятті на змінну використовуємо оператор &
	pTotal := &totalClients

	//Цей фрагмент виводить адресу, що зберігає вказівник
	fmt.Println("totalClients address:", pTotal)

	//Цей фрагмент виводить значення, на яке посилається вказівник
	//Для отримання значення даних, на яких посилається вказівник, використовуємо оператор *
	//Якщо оператор * знаходиться біля вказівника, то отримаємо значення на яке посилається вказівник
	fmt.Println("totalClients value:", *pTotal)

	//Ініціалізація зрізу списку клієнтів
	queue := []string{}

	fmt.Println("\nRegistering clients")
	fmt.Println("-----------------------")

	//Використання функції registerClient для додавання значень у масив
	//Функція отримує значення масиву через адресу пам`ятті
	registerClient(&queue, pTotal, "Dmytro")
	registerClient(&queue, pTotal, "Valriy")
	registerClient(&queue, pTotal, "Fedir")

	fmt.Println("Current queue:", queue)

	fmt.Println("\n Servicing clients")
	fmt.Println("-----------------------")

	//Використання функції для обслуговування клієнтів
	serveClient(&queue, pTotal)
	serveClient(&queue, pTotal)

	fmt.Println("Queue after service:", queue)

	fmt.Println("\nUsing new()")
	fmt.Println("-----------------------")

	//Використання оператора new
	//new виділяє місце пам'ятті та повертає її адресу
	//new часто використовується для ініціалізації порожніх змінних
	pAverage := new(float64)

	//Перевірка на наявність значення pAverage
	//pAverage вже зберігає адресу у пам'ятті, тому тіло умови завжди виконуватиметься
	if pAverage != nil {

		//Присвоєння значення вказівника pAverage
		*pAverage = float64(*pTotal) / 2.0
		fmt.Println("Average client served:", *pAverage)
	}

	fmt.Println("\nTesting nil-pointer")
	fmt.Println("-----------------------")

	//Перевірка на наявність значення pNil
	//Хоча pNil був ініціалізований, проте він не посилається на фрагмент пам'ятті
	//В результаті адреса вказвника pNil завжди nil (тобто адреса відсутня)
	var pNil *int
	if pNil == nil {
		fmt.Println("Nil pointer exception!")
	}

}
