package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lcalculator
#include "calculator.h"
*/

import "C"

import "fmt"

func main() {
	fmt.Println("E-Queue")
	fmt.Println("Using Cgo with custom library")

	a := 7
	b := 5

	// Виклик функції multiply з бібліотеки libcalculator.a
	result := C.multiply(C.int(a), C.int(b))

	fmt.Printf("Multiplying %d * %d = %d\n", a, b, int(result))
}

//Компіляція
// gcc -c calculator.c -o calculator.o         #Компіляція calculator.c в об'єктний файл
//ar rcs libcalculator.a calculator.o		   #Переведення calculator.o в статичну бібліотеку
