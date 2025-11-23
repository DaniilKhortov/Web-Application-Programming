package main

import "fmt"

func main() {
	//Оголошуємо defer.
	//Фенкція  з defer виконується останьою.
	defer fmt.Println("Queue system finished work. Have a nice day!")

	fmt.Println("E-Queue of client service")

	//Оголошення змінних для зчитування даних.
	var energy int
	var queueLength int

	//Зчитування рівня споживання енергії з консолі.
	fmt.Print("Enter current level of energy consumption (%): ")
	fmt.Scan(&energy)

	//Класифікація рівня споживання енергії.
	if energy < 50 {
		fmt.Println("Energy consumption on low level!")
	} else if energy <= 80 {
		fmt.Println("Energy consumption on avarage level!")
	} else if energy <= 100 {
		fmt.Println("Energy consumption on above-avarage level! There will be emergency power outages!")
	} else if energy > 100 || energy < 0 {
		fmt.Println("Entered energy consumption level is impossible!")
	}

	//Зчитування кількості клієнтів з консолі.
	fmt.Print("\nEnter client amount in queue: ")
	fmt.Scan(&queueLength)

	//Класифікація розміру черги.
	//Оператор умови switch.
	switch {
	//Умови switch називаються case.
	case queueLength == 0:
		fmt.Println("Queue is empty!")
	case queueLength <= 3:
		fmt.Println("Queue is short! 1 consultant needed!")
	case queueLength <= 6:
		fmt.Println("Queue is avarage! 2 consultant needed!")

		//Якщо жодна з умов не була виконана, застосовується default.
	default:
		fmt.Println("Queue is too long! Today is difficult to serve all the clients!")
	}

	//Обслуговування клієнтів.
	fmt.Println("\nServing clients")

	//Цикл for. Його аргументами є початкове значення, умова зупинки, інкремент.
	for i := 1; i <= queueLength; i++ {
		fmt.Printf("Servicing client #%d...\n", i)
	}

	fmt.Println("All clients were served!")
}
