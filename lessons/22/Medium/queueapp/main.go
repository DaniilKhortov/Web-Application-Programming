package queueapp

import (
	"fmt"
	"strings"
)

// SimulateQueueConcat — базова функція, що імітує формування повідомлень
// для клієнтів електронної черги шляхом конкатенації рядків через оператор +
func SimulateQueueConcat(n int) string {
	var result string
	for i := 1; i <= n; i++ {
		result += fmt.Sprintf("Client #%d is waiting...\n", i)
	}
	return result
}

// SimulateQueueConcatOptimized — оптимізована версія, що використовує strings.Builder
// замість конкатенації. Це зменшує кількість алокацій пам’яті та прискорює виконання.
func SimulateQueueConcatOptimized(n int) string {
	var sb strings.Builder
	sb.Grow(n * 30) // невелике попереднє виділення пам’яті
	for i := 1; i <= n; i++ {
		sb.WriteString(fmt.Sprintf("Client #%d is waiting...\n", i))
	}
	return sb.String()
}
