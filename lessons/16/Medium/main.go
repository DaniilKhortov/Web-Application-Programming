package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("E-Queue")

	// Поля вводу: Ім'я та Вік клієнта
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter name: ")

	surnameEntry := widget.NewEntry()
	surnameEntry.SetPlaceHolder("Enter surname: ")

	// Мітка для виводу результату
	resultLabel := widget.NewLabel("")

	// Кнопка для додавання у чергу
	button := widget.NewButton("Get number in queue", func() {
		name := nameEntry.Text
		surname := surnameEntry.Text

		// Базова валідація
		if name == "" {
			resultLabel.SetText("Error: name cannot be empty")
			return
		}

		if surname == "" {
			resultLabel.SetText("Error: name cannot be empty")
			return
		}

		// Формування номера у черзі на основі часу
		queueNumber := time.Now().Unix() % 1000
		result := fmt.Sprintf("Client %s %s got number in queue: %d", name, surname, queueNumber)
		resultLabel.SetText(result)

		// Вивід у консоль (додатково для перевірки)
		fmt.Println(result)
	})

	// Розміщення елементів інтерфейсу вертикально
	content := container.NewVBox(
		widget.NewLabel("Enter data to enter the queue"),
		nameEntry,
		surnameEntry,
		button,
		resultLabel,
	)

	myWindow.SetContent(container.NewCenter(content))
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}
