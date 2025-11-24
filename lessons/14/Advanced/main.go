package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

type Client struct {
	ID        int
	Name      string
	TicketNum int
}

type Result struct {
	Client  Client
	Message string
}

func DataGenerator(out chan<- Client, errCh chan<- error, done chan struct{}) {
	defer close(out)

	input := []Client{
		{1, "Michael", 1},
		{2, "Maria", 2},
		{3, "", 3},
		{4, "Dmytro", -1},
		{5, "Daryna", 4},
	}

	for _, c := range input {
		select {
		case <-done:
			fmt.Println("[Generator] Finnished computing by command")
			return
		default:
			fmt.Println("[Generator] Sent:", c)
			out <- c
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func Validator(in <-chan Client, out chan<- Client, errCh chan<- error, done chan struct{}) {
	defer close(out)

	for client := range in {
		select {
		case <-done:
			fmt.Println("[Validator] Finnished computing by command")
			return
		default:
			if !validate(client) {
				err := errors.New(fmt.Sprintf("Invalid clients data: %+v", client))
				errCh <- err
				return //
			}
			fmt.Printf("[Validator] Data is correct: %+v\n", client)
			out <- client
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func Aggregator(in <-chan Client, results chan<- Result, errCh chan<- error, done chan struct{}) {
	defer close(results)

	for client := range in {
		select {
		case <-done:
			fmt.Println("[Aggregator] Finnished computing by command")
			return
		default:
			message := fmt.Sprintf("Client %s with number %d served.", client.Name, client.TicketNum)
			results <- Result{Client: client, Message: message}
			fmt.Println("[Aggregator] Added result:", message)
			time.Sleep(300 * time.Millisecond)
		}
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

	dataCh := make(chan Client)
	validCh := make(chan Client)
	resultsCh := make(chan Result)
	errCh := make(chan error)
	done := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		DataGenerator(dataCh, errCh, done)
	}()

	go func() {
		defer wg.Done()
		Validator(dataCh, validCh, errCh, done)
	}()

	go func() {
		defer wg.Done()
		Aggregator(validCh, resultsCh, errCh, done)
	}()

	go func() {
		for res := range resultsCh {
			fmt.Printf("[Result] %s\n", res.Message)
		}
	}()

	select {
	case err := <-errCh:
		fmt.Println("\n[Main] Recieved error:", err)
		close(done)
	case <-time.After(5 * time.Second):
		fmt.Println("\n[Main] Work is done without error!")
		close(done)
	}

	wg.Wait()
	fmt.Println("Work is done!")
}
