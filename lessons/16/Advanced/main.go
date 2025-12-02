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
	//Ініціалізація застосунку fyne
	myApp := app.New()

	//Створення вікна графічного інтерфейсу
	myWindow := myApp.NewWindow("Застосунок 3")

	//Створення напису з часом
	timeLabel := widget.NewLabel("Час оновлюється...")

	//Створення опису поряд з годинником
	titleLabel := widget.NewLabel("Поточний час системи:")

	//Розміщення вмісту
	content := container.NewVBox(
		titleLabel,
		timeLabel,
	)
	myWindow.SetContent(content)

	//Задання розміру вікна
	myWindow.Resize(fyne.NewSize(300, 200))

	//Запуск графічного інтерфейсу
	myWindow.Show()

	//Створення лічильника часу
	ticker := time.NewTicker(1 * time.Second)

	//Створення горутини для оновлення часу
	go func() {
		for t := range ticker.C {

			currentTime := t.Format("15:04:05")

			timeLabel.SetText(fmt.Sprintf(" %s", currentTime))
		}
	}()

	//Запуск застосунку
	myApp.Run()
}
