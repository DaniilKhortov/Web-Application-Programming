package main

/*
#include <stdio.h>

// Вбудована C-функція, яка приймає два int і повертає їхню суму
int add(int a, int b) {
    return a + b;
}
*/
import "C"

import "fmt"

func main() {
	fmt.Println("E-Queue")
	fmt.Println("Calling C-function via cgo\n")

	// Припустимо, що у черзі 3 клієнти, і обслуговується ще 2
	var queueCount int = 3
	var servingNow int = 2

	// Викликаємо C-функцію для підрахунку загальної кількості клієнтів
	total := C.add(C.int(queueCount), C.int(servingNow))

	// Виводимо результат (перетворюємо тип C.int у Go int)
	fmt.Printf("General amount of clients in queue: %d\n", int(total))
}
