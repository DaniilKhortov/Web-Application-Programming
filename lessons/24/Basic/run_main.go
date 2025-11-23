package main

import (
	"fmt"
	queueapp "main/queueapp"
)

func main() {
	output := queueapp.SimulateQueueConcat(10)
	fmt.Println(output)
}
