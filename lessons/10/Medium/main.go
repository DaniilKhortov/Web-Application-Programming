package main

/*
#include <string.h> // Потрібно для strlen
#include <stdlib.h> // Потрібно для free
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("E-Queue")
	fmt.Println("C-Function with String transfer\n")

	// Go-рядок, який представляє ідентифікатор клієнта в черзі
	clientName := "client_1024"

	// Конвертація Go-рядка у C-рядок
	cStr := C.CString(clientName)

	// Обов’язкове звільнення пам’яті після використання
	defer C.free(unsafe.Pointer(cStr))

	// Виклик стандартної C-функції strlen
	length := C.strlen(cStr)

	// Вивід результату в Go
	fmt.Printf("Client: %s\n", clientName)
	fmt.Printf("Symbol amount in ID: %d\n", int(length))
}
