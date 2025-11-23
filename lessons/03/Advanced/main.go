package main

import (
	"fmt"
	"time"
)

// Функція перевірки значень клієнтів.
func checkCriticalError(value int) {
	if value > 100 {
		panic(fmt.Sprintf("Critical Error! Value %d is invalid.", value))
	}
}

func main() {
	//Зберігаємо час початку роботи програми.
	start := time.Now()

	//Використовуємо defer для визначення часу роботи програми.
	defer func() {
		fmt.Printf("Program action time: %v\n", time.Since(start))
	}()

	//Використовуємо defer для обробки виключень panic за допомогою recover.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic cought. Recovering...", r)
		}
	}()

	fmt.Println("E-Queue of client Service")

	//Демонстрація роботи кількох вкладених функцій.
	defer fmt.Println("Debug:: Defer 1 — executed last")
	defer fmt.Println("Debug:: Defer 2 — executed before last")

	//Масив споживання енергії клієнтами.
	clientsConsumption := []int{5, 10, 15, 8}
	//Масив стресу клієнтів.
	clientsStress := []int{20, 50, 70, 40}
	//Масив терпіння клієнтів.
	clientsPatience := []int{45, 85, 95, 65}

	//Мітка ProcessClients
ProcessClients:

	//Перебирання клієнтів для визначення максимуму та мінімуму.
	for i := 0; i < len(clientsConsumption); i++ {
		fmt.Printf("\nClient #%d:\n", i+1)

		energy := clientsConsumption[i]
		stress := clientsStress[i]
		patience := clientsPatience[i]

		//Валідація даних клієнтів.
		checkCriticalError(patience)

		//Класифікація рівня споживання.
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

		//Класифікація стресу клієнтів.
		switch {
		case stress < 30:
			fmt.Println("Client is calm.")
			fallthrough
		case stress < 60:
			fmt.Println("Client is worried.")
		default:
			fmt.Println("Client is stressed.")
		}

		//Класифікація терпіння клієнтів.
		if patience > 100 {
			fmt.Println("Client left!")
			break ProcessClients
		} else if patience > 80 {
			fmt.Println("Client is waiting too long!")
		} else {
			fmt.Println("Client is patient.")
		}
	}

	//Оголошення сумарного, максимального та мінімального споживання.
	fmt.Println("\nStatistics:")
	totalSum := 0
	maxVal := clientsConsumption[0]
	minVal := clientsConsumption[0]

	//Перебирання клієнтів для визначення максимуму та мінімуму.
	for _, val := range clientsConsumption {
		totalSum += val
		if val > maxVal {
			maxVal = val
		}
		if val < minVal {
			minVal = val
		}
	}
	//Визначення серееднього рівня споживання.
	avg := float64(totalSum) / float64(len(clientsConsumption))

	//Виід статистики у консоль.
	fmt.Printf("Total energy sum: %d | Average: %.2f | Max: %d | Min: %d\n",
		totalSum, avg, maxVal, minVal)
}
