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

	// Мітка для поточного часу
	timeLabel := widget.NewLabel("Time: ")

	// Кнопка для отримання номера у черзі
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
		fmt.Println(result)
	})

	// Компонування інтерфейсу
	content := container.NewVBox(
		widget.NewLabel("Enter data to enter the queue"),
		nameEntry,
		surnameEntry,
		button,
		resultLabel,
		timeLabel,
	)

	myWindow.SetContent(container.NewCenter(content))
	myWindow.Resize(fyne.NewSize(450, 350))

	// Горрутина для оновлення часу кожну секунду
	// У залежності від версії може викликати помилки
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for t := range ticker.C {
			// Форматування часу
			currentTime := fmt.Sprintf("Time: %s", t.Format("15:04:05"))
			// Оновлення GUI через SetText
			timeLabel.SetText(currentTime)
		}
	}()

	myWindow.ShowAndRun()
}
