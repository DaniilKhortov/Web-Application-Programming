package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//Ініціалізація застосунку fyne
	myApp := app.New()

	//Створення вікна графічного інтерфейсу
	myWindow := myApp.NewWindow("Застосунок 1")

	//Створення кнопки
	button := widget.NewButton("Натисни мене!", func() {
		//Тіло функції-тригера
		fmt.Println("Кнопку натиснуто!")
	})

	//Розміщення кнопки по-центру
	centered := container.NewCenter(button)
	myWindow.SetContent(centered)

	//Задання розміру вікна
	myWindow.Resize(fyne.NewSize(300, 200))

	//Запуск графічного інтерфейсу
	myWindow.ShowAndRun()
}
