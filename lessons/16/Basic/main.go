package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Створення нового додатку
	myApp := app.New()

	// Створення головного вікна
	myWindow := myApp.NewWindow("Застосунок 1")

	// Створення кнопки з текстом
	button := widget.NewButton("Натисни мене!", func() {
		fmt.Println("Кнопку натиснуто!")
	})

	// Центруємо кнопку у вікні за допомогою контейнера
	centered := container.NewCenter(button)

	// Встановлюємо контейнер як вміст вікна
	myWindow.SetContent(centered)

	// Встановлюємо розмір вікна
	myWindow.Resize(fyne.NewSize(300, 200))

	// Показуємо вікно
	myWindow.ShowAndRun()
}
