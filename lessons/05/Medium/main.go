package main

import (
	"errors"
	"fmt"
	"math/rand"
)

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

func applyDiscount(bill *float64, discount float64) {
	if discount > 0 && discount < 1000 {
		*bill = *bill - discount
		if *bill < 0 {
			*bill = 0
		}

	}

}

func calcBill(tariff, client, k float64) float64 {
	return tariff * client * k
}
func validateParams(param float64, name string) error {
	if param < 0 {
		message := fmt.Sprintf("Argument %s must be higher than 0! Current value: %v", name, param)
		return errors.New(message)
	}
	return nil
}

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

func sumAll(values ...float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum
}

func processBills(bills map[string]float64, handler func(string, float64)) {
	for name, bill := range bills {
		handler(name, bill)
	}
}

func main() {
	defer fmt.Println("\nProgram finished successfully!")

	fmt.Println("E-Queue")
	const tariff float64 = 3.74
	const k float64 = 1
	const discount float64 = 750

	clientNames := []string{"Taras", "Vitaliy", "Daniel"}
	clients, err := generateClients(clientNames)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("\nConsumer data")
	fmt.Println("---------------------------------------------")
	for client, energy := range clients {
		name := fmt.Sprintf("%s`s consumed energy", client)
		if err := validateParams(energy, name); err != nil {
			fmt.Println("Error!", err)
			return
		}
		fmt.Printf("%s: %.2f kW*hour.\n", client, energy)
	}

	clientsBill := map[string]float64{}
	for client, energyUsed := range clients {
		bill := calcBill(tariff, energyUsed, k)

		applyDiscount(&bill, discount)

		clientsBill[client] = bill
	}

	fmt.Println("\nConsumer bill data")
	fmt.Println("---------------------------------------------")
	for client, bill := range clientsBill {
		fmt.Printf("%s: %.2f UAH.\n", client, bill)
		name := fmt.Sprintf("%s`s energy price", client)
		if err := validateParams(bill, name); err != nil {
			fmt.Println("Error!", err)
			return
		}

	}
	processBills(clientsBill, func(name string, bill float64) {
		fmt.Printf("%s: %.2f UAH.\n", name, bill)
	})

	// --- 4. Статистика ---
	clientsData := []float64{}
	for _, bill := range clientsBill {
		clientsData = append(clientsData, bill)
	}
	sum, avg, min, max, count := sliceStat(clientsData)

	// --- 5. Варіадична функція ---
	totalByVar := sumAll(clientsData...)
	fmt.Println("\nConsumer statistics")
	fmt.Println("---------------------------------------------")
	fmt.Printf("Consumers stats: \nGeneral price paid: %.2f UAH(%.2f UAH with VarFunc)\nAvarage spendings: %.2f UAH\nMin price: %.2f UAH\nMax price: %.2f UAH\nConsumer amount: %d", sum, totalByVar, avg, min, max, count)

}
