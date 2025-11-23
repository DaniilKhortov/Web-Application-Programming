package queueapp

import (
	"fmt"
	"strings"
)

func SimulateQueueConcat(n int) string {
	var result string
	for i := 1; i <= n; i++ {
		result += fmt.Sprintf("Client #%d is waiting...\n", i)
	}
	return result
}

func SimulateQueueConcatOptimized(n int) string {
	var sb strings.Builder
	for i := 1; i <= n; i++ {
		sb.WriteString(fmt.Sprintf("Client #%d is waiting...\n", i))
	}
	return sb.String()
}
