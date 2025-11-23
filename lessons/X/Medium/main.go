package main

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Створення додатку
	myApp := app.New()

	// Створення головного вікна
	myWindow := myApp.NewWindow("Застосунок 2")

	// Поля вводу для двох чисел
	entry1 := widget.NewEntry()
	entry1.SetPlaceHolder("Введіть перше число")

	entry2 := widget.NewEntry()
	entry2.SetPlaceHolder("Введіть друге число")

	// Мітка для відображення результату
	resultLabel := widget.NewLabel("Місце для результату")

	// Кнопка для виконання обчислення
	calcButton := widget.NewButton("Обчислити суму", func() {
		// Отримання тексту з полів
		num1Text := entry1.Text
		num2Text := entry2.Text

		// Спроба перетворення тексту у числа
		num1, err1 := strconv.ParseFloat(num1Text, 64)
		num2, err2 := strconv.ParseFloat(num2Text, 64)

		// Перевірка на помилки
		if err1 != nil || err2 != nil {
			resultLabel.SetText(fmt.Sprintf("[%s] Помилка: введіть лише числа!", time.Now().Format("15:04:05")))
			return
		}

		// Обчислення суми
		sum := num1 + num2

		// Виведення результату
		resultLabel.SetText(fmt.Sprintf("[%s] Сума: %.2f", time.Now().Format("15:04:05"), sum))
	})

	// Розташування елементів у вертикальному контейнері
	content := container.NewVBox(
		widget.NewLabel("Введіть два числа для обчислення суми:"),
		entry1,
		entry2,
		calcButton,
		resultLabel,
	)

	// Встановлення вмісту вікна
	myWindow.SetContent(content)

	// Встановлення розміру вікна
	myWindow.Resize(fyne.NewSize(400, 300))

	// Запуск програми
	myWindow.ShowAndRun()
}
