package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Ініціалізація застосунку fyne
	myApp := app.New()
	myWindow := myApp.NewWindow("E-Queue")

	// Лічильник номерів у черзі
	queueNumber := 0

	// Мітка для відображення останнього номера
	label := widget.NewLabel("Press the button to get number to queue")

	// Кнопка для отримання номера
	button := widget.NewButton("Get number", func() {
		queueNumber++
		label.SetText("Your number in queue: " + strconv.Itoa(queueNumber))
		fmt.Println("User got number:", queueNumber)
	})

	// Розміщення кнопки та мітки по центру
	content := container.NewVBox(
		label,
		button,
	)

	myWindow.SetContent(container.NewCenter(content))
	myWindow.Resize(fyne.NewSize(400, 200))
	myWindow.ShowAndRun()
}
