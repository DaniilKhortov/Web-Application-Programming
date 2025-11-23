package main

import (
	"electronic-queue/utils"
	"fmt"
)

func main() {
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
			utils.AddClient(name)

		case 2:
			utils.ShowQueue()

		case 3:
			utils.ServeNext()

		case 4:
			fmt.Println("Finishing ...")
			return

		default:
			fmt.Println("Impossible action!")
		}
	}
}
