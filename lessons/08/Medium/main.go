package main

import (
	"Medium/internal/logger"
	"Medium/services"
	"fmt"
)

func main() {
	logger.LogInfo("Запуск програми електронної черги")

	for {
		fmt.Println("\n  E-QUEUE  ")
		fmt.Println("1. Add client")
		fmt.Println("2. Show queue")
		fmt.Println("3. Serve client")
		fmt.Println("4. Quit")
		fmt.Print("Choose action: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name string
			fmt.Print("Enter client`s name: ")
			fmt.Scan(&name)
			services.AddClient(name)

		case 2:
			services.ShowQueue()

		case 3:
			services.ServeNext()

		case 4:
			fmt.Println("Finishing ...")
			return

		default:
			fmt.Println("Impossible action!")
		}
	}
}
