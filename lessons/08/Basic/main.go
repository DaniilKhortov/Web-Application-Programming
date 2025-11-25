package main

import (
	"electronic-queue/utils"
	"fmt"
)

func main() {
	//Створення циклічного текстового меню
	for {
		fmt.Println("\n  E-QUEUE  ")
		fmt.Println("1. Add client")
		fmt.Println("2. Show queue")
		fmt.Println("3. Serve client")
		fmt.Println("4. Quit")
		fmt.Print("Choose action: ")

		//Зчитування вводу користувача
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			//Додавання клієнта
			var name string
			fmt.Print("Enter client`s name: ")
			fmt.Scan(&name)
			utils.AddClient(name)

		case 2:
			//Вивід поточної черги
			utils.ShowQueue()

		case 3:
			//Обслуговування першого клієнта в черзі
			utils.ServeNext()

		case 4:
			//Завершення роботи
			fmt.Println("Finishing ...")
			return

		default:
			fmt.Println("Impossible action!")
		}
	}
}
