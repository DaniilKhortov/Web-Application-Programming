package main

import (
	"fmt"
	"strings"
	"time"
)

type Client struct {
	Name      string
	TicketNum int
}

func DataReader(out chan<- Client) {

	input := []Client{
		{"Michael", 1},
		{"Maria", 2},
		{"", 3},
		{"Dmytro", -1},
		{"Daryna", 4},
	}

	for _, client := range input {
		fmt.Println("[Reader] Recieved record:", client)
		out <- client
		time.Sleep(300 * time.Millisecond)
	}
	close(out)
}

func DataProcessor(in <-chan Client) {
	for client := range in {
		if validate(client) {
			fmt.Printf("[Processor] Client '%s' is registered as %d\n",
				client.Name, client.TicketNum)
		} else {
			fmt.Printf("[Processor] Invalid data: %+v\n", client)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func validate(c Client) bool {
	if strings.TrimSpace(c.Name) == "" {
		return false
	}
	if c.TicketNum <= 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println("E-Queue")

	dataChan := make(chan Client)

	go DataReader(dataChan)
	DataProcessor(dataChan)

	fmt.Println("Work is done!")
}
