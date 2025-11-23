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
	// Створюємо новий Fyne-додаток
	myApp := app.New()

	// Головне вікно
	myWindow := myApp.NewWindow("Застосунок 3")

	// Мітка, що показує час
	timeLabel := widget.NewLabel("Час оновлюється...")

	// Додаткова мітка для опису
	titleLabel := widget.NewLabel("Поточний час системи:")

	// Контейнер для компонування
	content := container.NewVBox(
		titleLabel,
		timeLabel,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.Show()

	// Ticker генерує подію щосекунди
	ticker := time.NewTicker(1 * time.Second)

	// Запускаємо горутину для оновлення часу
	go func() {
		for t := range ticker.C {
			// Форматуємо поточний час
			currentTime := t.Format("15:04:05")

			// Оновлення тексту мітки у GUI
			timeLabel.SetText(fmt.Sprintf(" %s", currentTime))
		}
	}()

	// Запуск основного циклу програми
	myApp.Run()
}
