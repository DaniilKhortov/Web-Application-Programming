package main

import (
	"fmt"
	"time"
)

func main() {
	//Зберігаємо час початку роботи програми.
	start := time.Now()

	//Використовуємо defer для визначення часу роботи програми.
	defer func() {
		fmt.Printf("Program action time: %v\n", time.Since(start))
	}()

	fmt.Println("E-Queue of client Service")

	//Масив споживання енергії клієнтами.
	clients := []int{45, 60, 85, 30, 95, 70}

	//Валідація даних клієнтів.
	//Цикл for.
	for i, val := range clients {
		if val < 0 || val > 100 {
			fmt.Printf("Error! Client #%d energy consumption(%d) is not correct!\n", i+1, val)
			return
		}
	}

	//Виведення в консоль рівня споживання енергії клієнтами.
	//Класифікація рівня споживання.
	//Цикл for.
	for i, energy := range clients {
		fmt.Printf("\nClient #%d — energy consumption level: %d\n", i+1, energy)
		if energy < 50 {
			if energy < 30 {
				fmt.Println("Energy consumption on extremely-low level!")
			} else {
				fmt.Println("Energy consumption on low level!")
			}
		} else if energy <= 80 {
			if energy <= 65 {
				fmt.Println("Energy consumption on avarage level!")
			} else {
				fmt.Println("Energy consumption is on a bit higher than avarage level!")
			}
		} else {
			if energy <= 90 {
				fmt.Println("Energy consumption on high level!")
			} else {
				fmt.Println("Energy consumption on critical level! There will be emergency power outages!")
			}
		}
	}

	//Збереження клькості клієнтів у масиві.
	//len() повертає розмір масиву.
	totalClients := len(clients)
	fmt.Println("\nClient Queue")

	//Класифікація розміру черги.
	//Оператор умови switch.
	switch {
	case totalClients == 0:
		fmt.Println("Queue is empty!")
	case totalClients <= 3:
		fmt.Println("Queue is short! 1 consultant needed!")
	case totalClients <= 6:
		fmt.Println("Queue is avarage! 2 consultant needed!")
	default:
		fmt.Println("Queue is too long! Today is difficult to serve all the clients!")
	}

	//Проведення статистичного аналізу.
	//Оголошення сумарного, максимального та мінімального споживання.
	sum, max, min := 0, clients[0], clients[0]

	//Перебирання клієнтів для визначення максимуму та мінімуму.
	for _, val := range clients {
		sum += val
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	//Визначення серееднього рівня споживання.
	avg := float64(sum) / float64(len(clients))
	fmt.Printf("\nConsumed energy Statistics\nAvarage consumption: %.2f, Max consumption: %d, Min consumption:: %d\n", avg, max, min)

	//Виід клієнтів у консоль зі споживанням вище середнього.
	fmt.Println("\nClients with above-avarage consuption level:")
	for i, val := range clients {
		if val > 80 {
			fmt.Printf("Client #%d — %d\n", i+1, val)
		}
	}

	//Обслуговування клієнтів.
	fmt.Println("\nServing clients")
	for i := 1; i <= totalClients; i++ {
		fmt.Printf("Servicing client #%d...\n", i)
	}
	fmt.Println("All clients were served!")
}
